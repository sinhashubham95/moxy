package flags

import (
	"github.com/sinhashubham95/moxy/commons"
	flag "github.com/spf13/pflag"
	"os"
	"strconv"
	"strings"

	actuatorCommons "github.com/sinhashubham95/go-actuator/commons"
)

var (
	certFilePath    = flag.String(commons.CertFilePath, commons.CertFilePathDefaultValue, commons.CertFilePathUsage)
	keyFilePath     = flag.String(commons.KeyFilePath, commons.KeyFilePathDefaultValue, commons.KeyFilePathUsage)
	persistencePath = flag.String(commons.PersistencePath, commons.PersistenceDefaultValue, commons.PersistenceUsage)
	port            = flag.Int(commons.Port, commons.PortDefaultValue, commons.PortUsage)
	tlsEnabled      = flag.Bool(commons.TLSEnabled, commons.TLSEnabledDefaultValue, commons.TLSEnabledUsage)

	// dummy flags
	_ = flag.String(actuatorCommons.Env, actuatorCommons.EnvDefaultValue, actuatorCommons.EnvUsage)
	_ = flag.String(actuatorCommons.Name, actuatorCommons.NameDefaultValue, actuatorCommons.NameUsage)
	_ = flag.String(actuatorCommons.Version, actuatorCommons.VersionDefaultValue, actuatorCommons.VersionUsage)
)

func init() {
	flag.Parse()
}

// CertFilePath is the path to the certificate file
func CertFilePath() string {
	return *certFilePath
}

// KeyFilePath is the path to the key file
func KeyFilePath() string {
	return *keyFilePath
}

// PersistencePath is the path of the persistence database file
func PersistencePath() string {
	return *persistencePath
}

// Port is the port number where the application is running
func Port() int {
	// see if the port is available from the environment variable
	p, err := strconv.Atoi(os.Getenv(strings.ToUpper(commons.Port)))
	if err == nil {
		return p
	}
	// otherwise use the flag
	return *port
}

// TLSEnabled tells whether the application should start on HTTPS
func TLSEnabled() bool {
	return *tlsEnabled
}
