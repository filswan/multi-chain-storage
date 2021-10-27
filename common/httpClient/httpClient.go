package httpClient

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"payment-bridge/logs"
	"time"
)

var (
	httpClient *http.Client
	duration   = 30 * time.Second
	caFile     = flag.String("CA", "conf/tls/ca-chain.cert.pem", "A PEM eoncoded CA's certificate file.")
)

func initHttpClient() {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	httpClient = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(duration),
	}
}

func InitHttpsClient() {
	os.Setenv("GODEBUG", os.Getenv("GODEBUG")+",tls13=1")
	flag.Parse()

	caCert, err := ioutil.ReadFile(*caFile)
	if err != nil {
		logs.GetLogger().Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	tlsConfig := &tls.Config{
		RootCAs:    caCertPool,
		MinVersion: tls.VersionTLS13,
	}
	tlsConfig.BuildNameToCertificate()

	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		TLSClientConfig:       tlsConfig,
	}

	httpClient = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(duration),
	}
}

func GetHttpClient() *http.Client {
	if httpClient == nil {
		initHttpClient()
	}
	return httpClient
}

func SendRequestAndGetBytes(_httpMethod, _url string, _requestData []byte, _header http.Header) (responseData []byte, err error) {

	req, err := http.NewRequest(_httpMethod, _url, bytes.NewReader(_requestData))

	if err != nil {
		logs.GetLogger().Error(err)
		return
	}

	if len(_header) > 0 {
		req.Header = _header
	}

	resp, err := GetHttpClient().Do(req)
	defer func() {
		if resp != nil {
			if err = resp.Body.Close(); err != nil {
				logs.GetLogger().Error(err)
			}
		}
	}()
	if err != nil {
		logs.GetLogger().Error(err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprint("response contains a status code ", resp.StatusCode))
		logs.GetLogger().Error(err)
		return
	}

	responseData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logs.GetLogger().Error(err)
	}
	return
}
