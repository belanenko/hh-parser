package proxy

type schema string

const (
	HTTP   schema = "http"
	SOCKS5 schema = "socks5"
)

type Proxy struct {
	Schema schema
	Host   string
	Port   int32
	UserInfo
}

type UserInfo struct {
	Username string
	Password string
}
