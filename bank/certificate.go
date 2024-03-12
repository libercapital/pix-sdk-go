package bank

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
)

type Certificate struct {
	clientCrt    string
	clientCrtKey string
}

func (c Certificate) Load() (*x509.CertPool, tls.Certificate) {
	caCert, _ := base64.StdEncoding.DecodeString(c.clientCrt)
	caKey, _ := base64.StdEncoding.DecodeString(c.clientCrtKey)

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(caCert)

	cert, _ := tls.X509KeyPair(caCert, caKey)

	return rootCAs, cert
}

func NewCertificate(crt, crtKey string) Certificate {
	return Certificate{
		clientCrt:    crt,
		clientCrtKey: crtKey,
	}
}
