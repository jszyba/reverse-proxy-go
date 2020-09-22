package main

const (
	YarnResourceManager = "http://ip-10-42-24-247.ec2.internal:8088/cluster"
	SparkHistoryServer = "http://ip-10-42-24-247.ec2.internal:18080"
	SocksProxyURI = "socks5://localhost:8157"
	HTTP          = "http"
)

func main() {

	StartReverseProxy(YarnResourceManager)

}