package flags

import "flag"

func Set() (*int, *string, *int, *int) {
	numGorutines := flag.Int("threads", 1, "count of threads")
	proxyFilePath := flag.String("pfp", "", "proxy file path")
	indexStart := flag.Int("startid", 0, "Start parsing index")
	indexCount := flag.Int("countid", 1, "Count vacancy to parsing")
	flag.Parse()

	return numGorutines, proxyFilePath, indexStart, indexCount
}
