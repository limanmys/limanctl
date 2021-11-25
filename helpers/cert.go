package helpers

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"net"
	"time"
)

func GetCertificate(address string) (string, error) {
	conn, err := tls.DialWithDialer(&net.Dialer{
		Timeout: 3 * time.Second,
	}, "tcp", address, &tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		return "", err
	}
	defer conn.Close()
	var b bytes.Buffer
	for _, cert := range conn.ConnectionState().PeerCertificates {
		err := pem.Encode(&b, &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: cert.Raw,
		})
		if err != nil {
			return "", err
		}
	}
	return b.String(), nil
}
