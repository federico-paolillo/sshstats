package app

import "github.com/federico-paolillo/ssh-attempts/internal/loki"

type Config struct {
	Port    int
	Address string
	Loki    *loki.Config
}
