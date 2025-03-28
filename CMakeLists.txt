cmake_minimum_required(VERSION 3.50)
project(VPNClientEngineWindows)

set(CMAKE_C_STANDARD 11)


# Основная библиотека
add_library(vpnclient_engine_windows STATIC
    src/vpnclient_engine_windows.c
#    src/tun_device.c
#    src/crypto.c
)

target_include_directories(vpnclient_engine_windows PUBLIC include)
target_link_libraries(vpnclient_engine_windows ws2_32 bcrypt)

# Тесты
add_executable(test_client test/test_client.c)
target_link_libraries(test_client vpnclient_engine_windows)
