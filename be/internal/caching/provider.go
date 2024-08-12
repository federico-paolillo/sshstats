package caching

import (
	"fmt"
	"log"
	"time"

	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
	"github.com/patrickmn/go-cache"
)

const defaultExpiry = 24 * time.Hour
const cleanInterval = 12 * time.Hour

type Provider struct {
	cache    *cache.Cache
	provider stats.Provider
	log      *log.Logger
}

func NewProvider(log *log.Logger, provider stats.Provider) *Provider {
	cache := cache.New(defaultExpiry, cleanInterval)

	return &Provider{cache, provider, log}
}

func (p *Provider) Top15LoginAttempts(nodeName string) (stats.Attempts, error) {
	cacheKey := top15LoginAttemptsCacheKey(nodeName)

	maybeAttempsInCache, found := p.cache.Get(cacheKey)

	if found {
		if attempsInCache, ok := maybeAttempsInCache.(stats.Attempts); ok {
			log.Printf("hitting cache for node '%s'. cache key %s", nodeName, cacheKey)

			return attempsInCache, nil
		}
	}

	log.Printf("refreshing cache for node '%s'. cache key %s", nodeName, cacheKey)

	attempts, err := p.provider.Top15LoginAttempts(nodeName)

	if err != nil {
		return nil, fmt.Errorf("caching: could not refresh top 15 login attempts cache. %w", err)
	}

	p.cache.Set(cacheKey, attempts, cache.DefaultExpiration)

	return attempts, nil
}

func (p *Provider) Last10Attackers() (stats.Attackers, error) {
	panic("not implemented")
}

func top15LoginAttemptsCacheKey(nodeName string) string {
	return fmt.Sprintf("__t15la_%s", nodeName)
}

var _ stats.Provider = (*Provider)(nil)
