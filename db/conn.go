package db

import (
	"fmt"
	"strings"
	"time"
)

type Connection struct {
	Host                        string
	Port                        string
	Database                    string
	User                        string
	Password                    string
	SSLMode                     SSLMode
	SSLCertAuthorityCertificate string
	SSLPublicCertificate        string
	SSLPrivateKey               string
	FallbackConnections         []FallbackConnection
	MaxOpenConnections          int
	MaxIdleConnections          int
	ConnectionMaxIdleTime       time.Duration
	ConnectionMaxLifeTime       time.Duration
	ConnectionTimeout           time.Duration
	ConnectionOptions           []ConnectionOption
}

func (c Connection) CombineInstance() string {
	i := []string{fmt.Sprintf("%s:%s", c.Host, c.Port)}
	for _, v := range c.FallbackConnections {
		i = append(i, fmt.Sprintf("%s:%s", v.Host, v.Port))
	}

	return strings.Trim(strings.Join(i, ","), ",")
}

func (c Connection) ToConnectionString() string {
	s := fmt.Sprintf("postgresql://%s:%s@%s/%s?TimeZone=Asia/Ho_Chi_Minh&sslmode=%s",
		c.User,
		c.Password,
		c.CombineInstance(),
		c.Database,
		c.SSLMode)

	if c.SSLMode != Disable {
		if c.SSLCertAuthorityCertificate != "" {
			s += fmt.Sprintf("&sslrootcert=%s", c.SSLCertAuthorityCertificate)
		}

		if c.SSLPublicCertificate != "" {
			s += fmt.Sprintf("&sslcert=%s", c.SSLPublicCertificate)
		}

		if c.SSLPrivateKey != "" {
			s += fmt.Sprintf("&sslkey=%s", c.SSLPrivateKey)
		}
	}

	return s
}

type FallbackConnection struct {
	Host string
	Port string
}

type ConnectionOption func(*Connection)

func SetConnection(host string, port string) ConnectionOption {
	return func(c *Connection) {
		c.Host = host
		c.Port = port
	}
}

func SetFallbackConnection(host string, port string) ConnectionOption {
	return func(c *Connection) {
		if host != "" && port != "" {
			c.FallbackConnections = append(c.FallbackConnections, FallbackConnection{
				Host: host,
				Port: port,
			})
		}
	}
}

func SetSSL(mode SSLMode, caCertificate, publicCertificate, privateKey string) ConnectionOption {
	return func(c *Connection) {
		c.SSLMode = mode
		c.SSLCertAuthorityCertificate = caCertificate
		c.SSLPublicCertificate = publicCertificate
		c.SSLPrivateKey = privateKey
	}
}

func SetMaxOpenConnections(max int) ConnectionOption {
	return func(c *Connection) {
		c.MaxOpenConnections = max
	}
}

func SetMaxIdleConnections(max int) ConnectionOption {
	return func(c *Connection) {
		c.MaxIdleConnections = max
	}
}

func SetConnectionMaxIdleTime(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionMaxIdleTime = max
	}
}

func SetConnectionMaxLifeTime(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionMaxLifeTime = max
	}
}

func SetConnectionTimeout(max time.Duration) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionTimeout = max
	}
}

func SetLoginCredentials(user, password string) ConnectionOption {
	return func(c *Connection) {
		c.User = user
		c.Password = password
	}
}

func SetDatabase(database string) ConnectionOption {
	return func(c *Connection) {
		c.Database = database
	}
}

func AddChainConnectionOptions(opts ...ConnectionOption) ConnectionOption {
	return func(c *Connection) {
		c.ConnectionOptions = opts
	}
}
