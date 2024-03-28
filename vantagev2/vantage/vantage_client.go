// Code generated by go-swagger; DO NOT EDIT.

package vantage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/vantage-sh/vantage-go/vantagev2/vantage/access_grants"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/business_metrics"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/costs"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/dashboards"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/filters"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/folders"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/notifications"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/ping"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/prices"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/resource_reports"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/segments"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/teams"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/users"
	"github.com/vantage-sh/vantage-go/vantagev2/vantage/workspaces"
)

// Default vantage HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "api.vantage.sh"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v2"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new vantage HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Vantage {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new vantage HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Vantage {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new vantage client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Vantage {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Vantage)
	cli.Transport = transport
	cli.AccessGrants = access_grants.New(transport, formats)
	cli.BusinessMetrics = business_metrics.New(transport, formats)
	cli.Costs = costs.New(transport, formats)
	cli.Dashboards = dashboards.New(transport, formats)
	cli.Filters = filters.New(transport, formats)
	cli.Folders = folders.New(transport, formats)
	cli.Notifications = notifications.New(transport, formats)
	cli.Ping = ping.New(transport, formats)
	cli.Prices = prices.New(transport, formats)
	cli.ResourceReports = resource_reports.New(transport, formats)
	cli.Segments = segments.New(transport, formats)
	cli.Teams = teams.New(transport, formats)
	cli.Users = users.New(transport, formats)
	cli.Workspaces = workspaces.New(transport, formats)
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

// Vantage is a client for vantage
type Vantage struct {
	AccessGrants access_grants.ClientService

	BusinessMetrics business_metrics.ClientService

	Costs costs.ClientService

	Dashboards dashboards.ClientService

	Filters filters.ClientService

	Folders folders.ClientService

	Notifications notifications.ClientService

	Ping ping.ClientService

	Prices prices.ClientService

	ResourceReports resource_reports.ClientService

	Segments segments.ClientService

	Teams teams.ClientService

	Users users.ClientService

	Workspaces workspaces.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Vantage) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccessGrants.SetTransport(transport)
	c.BusinessMetrics.SetTransport(transport)
	c.Costs.SetTransport(transport)
	c.Dashboards.SetTransport(transport)
	c.Filters.SetTransport(transport)
	c.Folders.SetTransport(transport)
	c.Notifications.SetTransport(transport)
	c.Ping.SetTransport(transport)
	c.Prices.SetTransport(transport)
	c.ResourceReports.SetTransport(transport)
	c.Segments.SetTransport(transport)
	c.Teams.SetTransport(transport)
	c.Users.SetTransport(transport)
	c.Workspaces.SetTransport(transport)
}
