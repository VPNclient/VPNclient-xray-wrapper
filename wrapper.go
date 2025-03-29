package vpnclient_xray_wrapper

import (
    "context"
    "errors"
    "fmt"
    "time"

    "github.com/xtls/xray-core/core"
    "github.com/xtls/xray-core/features/dns"
    "github.com/xtls/xray-core/features/stats"
    "github.com/xtls/xray-core/infra/conf/serial"
    "github.com/xtls/xray-core/main"
)

// Client представляет собой VPN клиент на основе Xray
type Client struct {
    instance *core.Instance
    config   string
    stats    stats.Manager
    dns      dns.Client
}

// Stats представляет статистику соединения
type Stats struct {
    UploadBytes   uint64
    DownloadBytes uint64
    UploadRate    float64
    DownloadRate  float64
    Connections   uint32
}

// NewClient создает новый экземпляр VPN клиента
func NewClient() *Client {
    return &Client{}
}

// Start запускает VPN клиент с указанной конфигурацией
func (c *Client) Start(config string) error {
    if c.instance != nil {
	return errors.New("client is already running")
    }

    c.config = config

    // Инициализация Xray core
    configPb, err := serial.LoadJSONConfig([]byte(config))
    if err != nil {
	return fmt.Errorf("failed to parse config: %v", err)
    }

    instance, err := core.New(configPb)
    if err != nil {
	return fmt.Errorf("failed to create instance: %v", err)
    }

    if err := instance.Start(); err != nil {
	return fmt.Errorf("failed to start instance: %v", err)
    }

    c.instance = instance

    // Инициализация менеджера статистики
    c.stats = instance.GetFeature(stats.ManagerType()).(stats.Manager)

    // Инициализация DNS клиента
    c.dns = instance.GetFeature(dns.ClientType()).(dns.Client)

    return nil
}

// Stop останавливает VPN клиент
func (c *Client) Stop() error {
    if c.instance == nil {
	return errors.New("client is not running")
    }

    err := c.instance.Close()
    c.instance = nil
    return err
}

// GetStats возвращает текущую статистику соединения
func (c *Client) GetStats() (*Stats, error) {
    if c.instance == nil {
	return nil, errors.New("client is not running")
    }

    // Получаем статистику загрузки/выгрузки
    upCounter := c.stats.GetCounter("user>>>"+c.getTag()+">>>traffic>>>uplink")
    downCounter := c.stats.GetCounter("user>>>"+c.getTag()+">>>traffic>>>downlink")

    var upBytes, downBytes uint64
    if upCounter != nil {
	upBytes = upCounter.Value()
    }
    if downCounter != nil {
	downBytes = downCounter.Value()
    }

    // TODO: Реализовать расчет скорости (требуется хранение предыдущих значений и времени)

    return &Stats{
	UploadBytes:   upBytes,
	DownloadBytes: downBytes,
	Connections:   c.getActiveConnections(),
    }, nil
}

// QueryDNS выполняет DNS запрос
func (c *Client) QueryDNS(domain string) ([]string, error) {
    if c.instance == nil {
	return nil, errors.New("client is not running")
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    ips, err := c.dns.LookupIP(ctx, domain)
    if err != nil {
	return nil, err
    }

    var result []string
    for _, ip := range ips {
	result = append(result, ip.String())
    }

    return result, nil
}

// getTag извлекает тег из конфигурации
func (c *Client) getTag() string {
    // TODO: Реализовать парсинг тега из конфигурации
    // Временно возвращаем дефолтное значение
    return "proxy"
}

// getActiveConnections возвращает количество активных соединений
func (c *Client) getActiveConnections() uint32 {
    // TODO: Реализовать подсчет активных соединений
    // Временно возвращаем 0
    return 0
}