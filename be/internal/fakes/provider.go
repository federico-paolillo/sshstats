package fakes

import (
	"testing"

	"github.com/federico-paolillo/ssh-attempts/pkg/stats"
)

type Provider struct {
	Calls int
	Data  map[string]stats.Attempts
	T     *testing.T
}

func (f *Provider) Top15LoginAttempts(nodename string) (stats.Attempts, error) {
	var result stats.Attempts

	f.Calls++

	if result, ok := f.Data[nodename]; ok {
		return result, nil
	}

	return result, nil
}

func (f *Provider) Last10Attackers() (stats.Attackers, error) {
	panic("not implemented")
}

var _ stats.Provider = (*Provider)(nil)
