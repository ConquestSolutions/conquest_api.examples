// Code generated by go-swagger; DO NOT EDIT.

package geo_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new geo service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for geo service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	GeoServiceAddGeoPackage(params *GeoServiceAddGeoPackageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GeoServiceAddGeoPackageOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GeoServiceAddGeoPackage geo service add geo package API
*/
func (a *Client) GeoServiceAddGeoPackage(params *GeoServiceAddGeoPackageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GeoServiceAddGeoPackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGeoServiceAddGeoPackageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GeoService_AddGeoPackage",
		Method:             "POST",
		PathPattern:        "/api/geo/add_geo_package",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GeoServiceAddGeoPackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GeoServiceAddGeoPackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GeoServiceAddGeoPackageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
