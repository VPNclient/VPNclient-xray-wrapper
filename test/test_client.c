#include "../include/vpnclient_engine_windows.h"
#include <assert.h>

void test_tap_creation() {
    HANDLE hTap = tap_create("tap0");
    assert(hTap != NULL && hTap != INVALID_HANDLE_VALUE);
    CloseHandle(hTap);
}

int main() {
    WSADATA wsa;
    WSAStartup(MAKEWORD(2, 2), &wsa);

    test_tap_creation();
    printf("All tests passed!\n");

    WSACleanup();
    return 0;
}