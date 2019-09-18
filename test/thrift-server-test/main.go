package thrift_server_test

func main() {
	addr := ":9898"
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	runServer(transportFactory, protocolFactory, addr, false)
}
