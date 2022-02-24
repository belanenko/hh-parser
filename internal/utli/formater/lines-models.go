package formater

import (
	"strconv"
	"strings"

	"github.com/hh-parser/internal/models/proxy"
)

func StringToProxy(schema proxy.Schema, strProxy ...string) []proxy.Proxy {
	out := make([]proxy.Proxy, len(strProxy))
	for i, p := range strProxy {
		data := strings.Split(p, ":")
		host := data[0]
		port, _ := strconv.Atoi(data[1])
		username := data[3]
		password := data[3]

		out[i] = proxy.Proxy{
			Schema: schema,
			Host:   host,
			Port:   int32(port),
			UserInfo: proxy.UserInfo{
				Username: username,
				Password: password,
			},
		}
	}
	return out
}
