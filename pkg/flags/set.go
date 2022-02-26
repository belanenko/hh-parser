package flags

import "flag"

type Config struct {
	NumGorutines  int
	ProxyFilePath string
	IndexStart    int
	Count         int
}

func Set() *Config {
	numGorutines := flag.Int("threads", 1, "count of threads")
	proxyFilePath := flag.String("pfp", "", "proxy file path")
	indexStart := flag.Int("startid", 0, "Start parsing index")
	indexCount := flag.Int("countid", 1, "Count vacancy to parsing")
	flag.Parse()

	return &Config{
		NumGorutines:  *numGorutines,
		ProxyFilePath: *proxyFilePath,
		IndexStart:    *indexStart,
		Count:         *indexCount,
	}
}
