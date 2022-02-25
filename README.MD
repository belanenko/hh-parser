# hh.ru parse tool 

---

## Start flags

1. `-threads` sets parsing-gorutines count. <br/>Example: `-threads 4`.
2. <del>`-config` path to *.env config file. <br/>Example: `-config "~/app/config.env"`
3. `-startIndex` start vacancy id position  <br/>Example: `-startIndex 50000000`
4. `-conut` count vacancies which will be parsed<br/>Example: `-conut 10000`
5. `-proxyfile` path to *.txt proxy file. <br/>Proxy format should be: `ip:port:login:password`<br/>Example: `-proxyfile "~/app/proxy.txt"` 

**NOTE: you can use only `socks5` <del>and `https`</del> proxy with auth.**
### Example flags: 
> `./app.go --threads 5 --pfp=/home/tim/code/github.com/belanenko/hh-parser/assets/socks5.txt --startid=500000 --countid=1000`

