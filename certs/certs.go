// certs/certs.go

package certs

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func LoadCertificate(certPath, keyPath string) (tls.Certificate, error) {
	return tls.LoadX509KeyPair(certPath, keyPath)
}

func LoadCA(certPath string) (*x509.CertPool, error) {
	caCert, err := ioutil.ReadFile(certPath)
	if err != nil {
		return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)
	return caCertPool, nil
}
