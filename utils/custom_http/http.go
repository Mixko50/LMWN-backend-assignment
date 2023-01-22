package custom_http

import (
	"crypto/tls"
	"net/http"
	"time"
)

var Transport = &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
}

var HttpClient = &http.Client{
	Timeout:   10 * time.Second,
	Transport: Transport,
}
