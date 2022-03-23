package client

func ConnectProxy(proxy *string) (proxyString interface{}) {
	if len(*proxy) > 0 {
		var proxyString interface{} = proxy
		return proxyString
	} else {
		return nil
	}
}