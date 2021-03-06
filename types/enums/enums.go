package enums

type RedisNameSpace = string

const (
	UrlNS  RedisNameSpace = "url"
	UserNS RedisNameSpace = "user"
)

type Environment = string

const (
	Docker     Environment = "docker"
	Local      Environment = "local"
	Testing    Environment = "testing"
	Sandbox    Environment = "sandbox"
	Production Environment = "production"
)

type DatabaseName = string
