// Code generated by go-swagger; DO NOT EDIT.

package harborcli

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AimAlex/harbor-client/harborcli/artifact"
	"github.com/AimAlex/harbor-client/harborcli/auditlog"
	"github.com/AimAlex/harbor-client/harborcli/chart_repository"
	"github.com/AimAlex/harbor-client/harborcli/icon"
	"github.com/AimAlex/harbor-client/harborcli/preheat"
	"github.com/AimAlex/harbor-client/harborcli/products"
	"github.com/AimAlex/harbor-client/harborcli/project"
	"github.com/AimAlex/harbor-client/harborcli/repository"
	"github.com/AimAlex/harbor-client/harborcli/scan"
	"github.com/AimAlex/harbor-client/harborcli/scanners"
	"github.com/AimAlex/harbor-client/harborcli/version"
)

// Default harbor API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v2.0"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new harbor API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *HarborAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new harbor API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *HarborAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new harbor API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *HarborAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(HarborAPI)
	cli.Transport = transport
	cli.Artifact = artifact.New(transport, formats)
	cli.Auditlog = auditlog.New(transport, formats)
	cli.ChartRepository = chart_repository.New(transport, formats)
	cli.Icon = icon.New(transport, formats)
	cli.Preheat = preheat.New(transport, formats)
	cli.Products = products.New(transport, formats)
	cli.Project = project.New(transport, formats)
	cli.Repository = repository.New(transport, formats)
	cli.Scan = scan.New(transport, formats)
	cli.Scanners = scanners.New(transport, formats)
	cli.Version = version.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// HarborAPI is a client for harbor API
type HarborAPI struct {
	Artifact artifact.ClientService

	Auditlog auditlog.ClientService

	ChartRepository chart_repository.ClientService

	Icon icon.ClientService

	Preheat preheat.ClientService

	Products products.ClientService

	Project project.ClientService

	Repository repository.ClientService

	Scan scan.ClientService

	Scanners scanners.ClientService

	Version version.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *HarborAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Artifact.SetTransport(transport)
	c.Auditlog.SetTransport(transport)
	c.ChartRepository.SetTransport(transport)
	c.Icon.SetTransport(transport)
	c.Preheat.SetTransport(transport)
	c.Products.SetTransport(transport)
	c.Project.SetTransport(transport)
	c.Repository.SetTransport(transport)
	c.Scan.SetTransport(transport)
	c.Scanners.SetTransport(transport)
	c.Version.SetTransport(transport)
}
