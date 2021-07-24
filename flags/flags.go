package flags

import (
	actuatorCommons "github.com/sinhashubham95/go-actuator/commons"
	flag "github.com/spf13/pflag"

	"github.com/sinhashubham95/moxy/commons"
)

var (
	persistencePath = flag.String(commons.PersistencePath, commons.PersistenceDefaultValue, commons.PersistenceUsage)
	port            = flag.Int(commons.Port, commons.PortDefaultValue, commons.PortUsage)

	// dummy flags
	_ = flag.String(actuatorCommons.Env, actuatorCommons.EnvDefaultValue, actuatorCommons.EnvUsage)
	_ = flag.String(actuatorCommons.Name, actuatorCommons.NameDefaultValue, actuatorCommons.NameUsage)
	_ = flag.String(actuatorCommons.Version, actuatorCommons.VersionDefaultValue, actuatorCommons.VersionUsage)
)

func init() {
	flag.Parse()
}

// PersistencePath is the path of the persistence database file
func PersistencePath() string {
	return *persistencePath
}

// Port is the port number where the application is running
func Port() int {
	return *port
}
