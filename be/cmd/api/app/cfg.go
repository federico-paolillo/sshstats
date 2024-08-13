package app

type LokiSettings struct {
	Endpoint string
	Username string
	Password string
}

type AuthSettings struct {
	HeaderKey   string
	HeaderValue string
}

type ServerSettings struct {
	Address string
}

type Config struct {
	Server *ServerSettings
	Loki   *LokiSettings
	Auth   *AuthSettings
}
