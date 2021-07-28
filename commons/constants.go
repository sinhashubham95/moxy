package commons

import "fmt"

// Common Constants
const (
	ActuatorPrefix           = "/actuator"
	CertFilePath             = "cert-file"
	CertFilePathDefaultValue = ""
	CertFilePathUsage        = "cert file path is the path of the certificate file"
	KeyFilePath              = "key-file"
	KeyFilePathDefaultValue  = ""
	KeyFilePathUsage         = "key file path is the path of the key file"
	MoxyPrefix               = "/moxy"
	PersistencePath          = "persistence-path"
	PersistenceDefaultValue  = "persistence.db"
	PersistenceUsage         = "persistence path is the path of the file which acts as the persistence layer for this application"
	Port                     = "port"
	PortDefaultValue         = 8080
	PortUsage                = "port is the port number on which the application is running"
	TLSEnabled               = "tls"
	TLSEnabledDefaultValue   = false
	TLSEnabledUsage          = "tls enabled tells whether the application is running on HTTPS or not"
)

// Endpoint Paths
var (
	BasePath           = "/"
	MockEndpointPath   = fmt.Sprintf("%smock", MoxyPrefix)
	UnMockEndpointPath = fmt.Sprintf("%sunMock", MoxyPrefix)
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
