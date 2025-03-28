#ifndef VPNCLIENT_ENGINE_WINDOWS_H
#define VPNCLIENT_ENGINE_WINDOWS_H


#include <windows.h>
#include <winsock2.h>

#define TAP_IOCTL_SET_MEDIA_STATUS 0x4000003C  // Для TAP-Windows драйвера

HANDLE tap_create(const char* dev_name);
NTSTATUS encrypt_data(
    PUCHAR plaintext, ULONG plaintext_len,
    PUCHAR key, ULONG key_len,
    PUCHAR iv, ULONG iv_len,
    PUCHAR ciphertext, ULONG* ciphertext_len
);
void vpnclient_engine_windows_loop(HANDLE hTap, SOCKET sock);

#endif