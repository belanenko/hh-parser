package proxy

type Schema string

const (
	HTTP   Schema = "http"
	SOCKS5 Schema = "socks5"
)

type Proxy struct {
	Schema Schema
	Host   string
	Port   int32
	UserInfo
}

type UserInfo struct {
	Username string
	Password string
}
