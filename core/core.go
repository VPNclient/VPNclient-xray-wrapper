package core

import (
    "github.com/xtls/xray-core/core"
    "github.com/xtls/xray-core/main"
)

// Initialize инициализирует Xray core
func Initialize() {
    // Устанавливаем уровень лога
    main.SetConfigLevel("warning")
}

// CreateInstance создает новый экземпляр Xray
func CreateInstance(config []byte) (*core.Instance, error) {
    configPb, err := serial.LoadJSONConfig(config)
    if err != nil {
	return nil, err
    }

    return core.New(configPb)
}