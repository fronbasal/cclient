package cclient

import (
	"net/http"

	"golang.org/x/net/proxy"

	utls "gitlab.com/yawning/utls.git"
)

func NewClient(clientHello utls.ClientHelloID, proxyUrl ...string) (http.Client, error) {
	if len(proxyUrl) > 0 && len(proxyUrl) > 0 {
		dialer, err := newConnectDialer(proxyUrl[0])
		if err != nil {
			return http.Client{}, err
		}
		return http.Client{
			Transport: newRoundTripper(clientHello, dialer),
		}, nil
	} else {
		return http.Client{
			Transport: newRoundTripper(clientHello, proxy.Direct),
		}, nil
	}
}
