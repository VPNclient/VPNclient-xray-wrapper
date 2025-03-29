package core

import (
    "github.com/xtls/xray-core/features/stats"
)

// StatsManager управляет статистикой
type StatsManager struct {
    manager stats.Manager
}

// NewStatsManager создает новый менеджер статистики
func NewStatsManager(instance *core.Instance) *StatsManager {
    return &StatsManager{
	manager: instance.GetFeature(stats.ManagerType()).(stats.Manager),
    }
}

// GetTraffic возвращает статистику трафика
func (m *StatsManager) GetTraffic(tag string) (up, down uint64) {
    upCounter := m.manager.GetCounter("user>>>" + tag + ">>>traffic>>>uplink")
    downCounter := m.manager.GetCounter("user>>>" + tag + ">>>traffic>>>downlink")

    if upCounter != nil {
	up = upCounter.Value()
    }
    if downCounter != nil {
	down = downCounter.Value()
    }

    return
}
