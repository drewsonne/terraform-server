package restapi

import (
	"encoding/json"
	"github.com/go-openapi/swag"
	"github.com/zeebox/terraform-server/server/restapi/operations"
)

var cliconfig = ConfigFileFlags{}

// Configuration describes the structure of options in the server config file
type Configuration struct {
	Identity CfgIdentity `json:"identity"`
	Backend  CfgBackend  `json:"backend"`
}

// CfgIdentity describes the structure of options for Identity Providers
type CfgIdentity struct {
	Defaults []CfgIdentityDefault `json:"defaults"`
}

// CfgIdentityDefault allows the setting of a username and password for a default user. This value will only be used
// when initialising a new managed Identity Provider, and will be ignored on subsequent boots.
// @TODO Restrict this to only be for the `admin` user
// @TODO Allow a force cli option when booting to reset the password
type CfgIdentityDefault struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

// CfgBackend describes how the server can load the backend database and the primary managed Identity Provider
type CfgBackend struct {
	DatabaseType string      `json:"database_type"`
	Database     interface{} `json:"database"`
	IdentityType string      `json:"identity_type"`
	Identity     interface{} `json:"identity"`
}

func (cfg *Configuration) loadFromFile(path string) {
	j, err := swag.YAMLDoc(path)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(j, cfg)
	if err != nil {
		panic(err)
	}
}

//ConfigFileFlags for loading settings for the server
type ConfigFileFlags struct {
	ConfigFile string `short:"c" long:"config" description:"Path to configuration on disk"`
}

func parseCliConfiguration() *Configuration {
	cfg := newDefaultConfiguration()
	if cliconfig.ConfigFile != "" {
		cfg.loadFromFile(cliconfig.ConfigFile)
	}
	return cfg
}

func newDefaultConfiguration() *Configuration {
	return &Configuration{
		Identity: CfgIdentity{
			Defaults: []CfgIdentityDefault{
				{User: "admin", Password: "password"},
			},
		},
		Backend: CfgBackend{
			IdentityType: "memory",
			DatabaseType: "memory",
		},
	}
}

func configureFlags(api *operations.TerraformServerAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "Config",
			LongDescription:  "Configuration",
			Options:          &cliconfig,
		},
	}
}