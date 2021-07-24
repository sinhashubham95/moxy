package commons

// Common Constants
const (
	ActuatorPrefix          = "/actuator"
	MoxyPrefix              = "/moxy"
	PersistencePath         = "persistence-path"
	PersistenceDefaultValue = "persistence.db"
	PersistenceUsage        = "persistence path is the path of the file which acts as the persistence layer for this application"
	Port                    = "port"
	PortDefaultValue        = 8080
	PortUsage               = "port is the port number on which the application is running"
)

// Endpoint Paths
const (
	MockEndpointPath   = "/moxy/mock"
	UnMockEndpointPath = "/moxy/unMock"
)

// MockEntityName is the mock entity name
const (
	MockEntityName = "mock"
)

// Response constants
const (
	TagHeader                  = "x-tag"
	URLHeader                  = "x-url"
	ApplicationJSONContentType = "application/json"
	TextStringContentType      = "text/string"
)
