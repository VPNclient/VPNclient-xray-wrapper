package core

import (
    "context"
    "time"

    "github.com/xtls/xray-core/features/dns"
)

// DNSClient представляет DNS клиент
type DNSClient struct {
    client dns.Client
}

// NewDNSClient создает новый DNS клиент
func NewDNSClient(instance *core.Instance) *DNSClient {
    return &DNSClient{
	client: instance.GetFeature(dns.ClientType()).(dns.Client),
    }
}

// Lookup выполняет DNS запрос
func (c *DNSClient) Lookup(domain string) ([]string, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    ips, err := c.client.LookupIP(ctx, domain)
    if err != nil {
	return nil, err
    }

    var result []string
    for _, ip := range ips {
	result = append(result, ip.String())
    }

    return result, nil
}