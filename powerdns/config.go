package powerdns

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log"

	"github.com/hashicorp/terraform/helper/pathorcontents"
)

type Config struct {
	ServerUrl     string
	ApiKey        string
	InsecureHTTPS bool
	CACertificate string
}

// Client returns a new client for accessing PowerDNS
func (c *Config) Client() (*Client, error) {

	tlsConfig := &tls.Config{}

	if c.CACertificate != "" {

		caCert, _, err := pathorcontents.Read(c.CACertificate)
		if err != nil {
			return nil, fmt.Errorf("Error reading CA Cert: %s", err)
		}

		caCertPool := x509.NewCertPool()
		caCertPool.AppendCertsFromPEM([]byte(caCert))
		tlsConfig.RootCAs = caCertPool
	}

	tlsConfig.InsecureSkipVerify = c.InsecureHTTPS

	client, err := NewClient(c.ServerUrl, c.ApiKey, tlsConfig)

	if err != nil {
		return nil, fmt.Errorf("Error setting up PowerDNS client: %s", err)
	}

	log.Printf("[INFO] PowerDNS Client configured for server %s", c.ServerUrl)

	return client, nil
}
