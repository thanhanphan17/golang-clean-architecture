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

// CombineInstance combines the host and port of the connection
// along with the host and port of any fallback connections
// and returns a comma-separated string of all the combined instances.
func (c Connection) CombineInstance() string {
	// Create a slice to store the combined instances
	instances := []string{fmt.Sprintf("%s:%s", c.Host, c.Port)}

	// Iterate over the fallback connections
	for _, v := range c.FallbackConnections {
		// Append the host and port of each fallback connection to the slice
		instances = append(instances, fmt.Sprintf("%s:%s", v.Host, v.Port))
	}

	// Join all the instances with a comma and remove any leading/trailing commas
	return strings.Trim(strings.Join(instances, ","), ",")
}

// ToConnectionString returns the PostgreSQL connection string based on the Connection struct.
func (c Connection) ToPostgresConnectionString() string {
	// Construct the base connection string with user, password, instance, database, and SSL mode.
	// Use fmt.Sprintf to format the string.
	s := fmt.Sprintf("postgresql://%s:%s@%s/%s?TimeZone=Asia/Ho_Chi_Minh&sslmode=%s",
		c.User,
		c.Password,
		c.CombineInstance(),
		c.Database,
		c.SSLMode)

	// Check if SSL mode is not Disable.
	if c.SSLMode != Disable {
		// Append the SSL root certificate to the connection string if provided.
		if c.SSLCertAuthorityCertificate != "" {
			s += fmt.Sprintf("&sslrootcert=%s", c.SSLCertAuthorityCertificate)
		}

		// Append the SSL public certificate to the connection string if provided.
		if c.SSLPublicCertificate != "" {
			s += fmt.Sprintf("&sslcert=%s", c.SSLPublicCertificate)
		}

		// Append the SSL private key to the connection string if provided.
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

// SetConnection returns a ConnectionOption function that sets the host and port of a Connection struct.
//
// Parameters:
//   - host: the host address to set
//   - port: the port number to set
//
// Returns:
//   - ConnectionOption: a function that sets the host and port of a Connection struct
func SetConnection(host string, port string) ConnectionOption {
	return func(c *Connection) {
		c.Host = host
		c.Port = port
	}
}

// SetFallbackConnection is a function that returns a ConnectionOption.
// It sets the fallback connection with the given host and port.
// If both the host and port are provided, the fallback connection is added to the list of fallback connections.
// The returned ConnectionOption function can be used to modify a Connection object.
// The modified Connection object will have the fallback connection added if the host and port are provided.
func SetFallbackConnection(host string, port string) ConnectionOption {
	return func(c *Connection) {
		// Check if both host and port are provided
		if host != "" && port != "" {
			// Add the fallback connection to the list of fallback connections
			c.FallbackConnections = append(c.FallbackConnections, FallbackConnection{
				Host: host,
				Port: port,
			})
		}
	}
}

// SetSSL is a function that returns a ConnectionOption function.
// The returned function sets the SSL mode and certificates for a Connection.
// It takes in the SSL mode, CA certificate, public certificate, and private key as arguments.
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
