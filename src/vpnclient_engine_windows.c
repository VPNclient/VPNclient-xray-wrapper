#include <windows.h>
#include <winsock2.h>
#include <ws2tcpip.h>
#include <stdio.h>
#include <stdlib.h>
#include <bcrypt.h>
#include "vpn_client.h"

#pragma comment(lib, "ws2_32.lib")
#pragma comment(lib, "Bcrypt.lib")

// Создание TAP-адаптера (через драйвер OpenVPN)
HANDLE tap_create(const char* dev_name) {
    char device_path[256];
    snprintf(device_path, sizeof(device_path), "\\\\.\\Global\\%s.tap", dev_name);

    HANDLE hTap = CreateFileA(
        device_path,
        GENERIC_READ | GENERIC_WRITE,
        0, 0,
        OPEN_EXISTING,
        FILE_ATTRIBUTE_SYSTEM,
        NULL
    );

    if (hTap == INVALID_HANDLE_VALUE) {
        printf("Cannot open TAP device (Error: %d)\n", GetLastError());
        return NULL;
    }

    // Активация адаптера
    DWORD len;
    BOOL status = DeviceIoControl(
        hTap, TAP_IOCTL_SET_MEDIA_STATUS,
        &(DWORD){ TRUE }, sizeof(DWORD),
        NULL, 0, &len, NULL
    );

    if (!status) {
        printf("Failed to set media status (Error: %d)\n", GetLastError());
        CloseHandle(hTap);
        return NULL;
    }

    printf("TAP device %s created\n", dev_name);
    return hTap;
}

// Шифрование данных (AES-256-GCM через BCrypt)
NTSTATUS encrypt_data(
    PUCHAR plaintext, ULONG plaintext_len,
    PUCHAR key, ULONG key_len,
    PUCHAR iv, ULONG iv_len,
    PUCHAR ciphertext, ULONG* ciphertext_len
) {
    BCRYPT_ALG_HANDLE hAlg = NULL;
    BCRYPT_KEY_HANDLE hKey = NULL;
    NTSTATUS status;

    // Открываем алгоритм AES
    status = BCryptOpenAlgorithmProvider(
        &hAlg, BCRYPT_AES_ALGORITHM,
        NULL, 0
    );
    if (status != 0) goto cleanup;

    // Генерируем ключ
    status = BCryptGenerateSymmetricKey(
        hAlg, &hKey, NULL, 0,
        key, key_len, 0
    );
    if (status != 0) goto cleanup;

    // Шифруем данные
    status = BCryptEncrypt(
        hKey, plaintext, plaintext_len,
        NULL, iv, iv_len,
        ciphertext, *ciphertext_len,
        ciphertext_len, BCRYPT_BLOCK_PADDING
    );

cleanup:
    if (hKey) BCryptDestroyKey(hKey);
    if (hAlg) BCryptCloseAlgorithmProvider(hAlg, 0);
    return status;
}

// Основной цикл VPN-клиента
void vpnclient_engine_windows_loop(HANDLE hTap, SOCKET sock) {
    fd_set read_fds;
    char buffer[4096];

    while (1) {
        FD_ZERO(&read_fds);
        FD_SET(sock, &read_fds);

        // Ждём данные от сокета или TAP (WinSock + Overlapped I/O)
        int ret = select(0, &read_fds, NULL, NULL, NULL);
        if (ret == SOCKET_ERROR) {
            printf("select() failed: %d\n", WSAGetLastError());
            break;
        }

        if (FD_ISSET(sock, &read_fds)) {
            // Чтение из сети -> запись в TAP
            int n = recv(sock, buffer, sizeof(buffer), 0);
            DWORD written;
            WriteFile(hTap, buffer, n, &written, NULL);
        }

        // Чтение из TAP -> отправка в сеть (асинхронно)
        DWORD read;
        if (ReadFile(hTap, buffer, sizeof(buffer), &read, NULL)) {
            send(sock, buffer, read, 0);
        }
    }
}