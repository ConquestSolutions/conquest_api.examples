// Code generated by go-swagger; DO NOT EDIT.

package asset_inspection_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new asset inspection service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for asset inspection service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AssetInspectionServiceCreateAssetInspection(params *AssetInspectionServiceCreateAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceCreateAssetInspectionOK, error)

	AssetInspectionServiceCreateDefectForAssetInspection(params *AssetInspectionServiceCreateDefectForAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceCreateDefectForAssetInspectionOK, error)

	AssetInspectionServiceDeleteAssetInspection(params *AssetInspectionServiceDeleteAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceDeleteAssetInspectionOK, error)

	AssetInspectionServiceGetAssetInspection(params *AssetInspectionServiceGetAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceGetAssetInspectionOK, error)

	AssetInspectionServiceUpdateAssetInspection(params *AssetInspectionServiceUpdateAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceUpdateAssetInspectionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  AssetInspectionServiceCreateAssetInspection asset inspection service create asset inspection API
*/
func (a *Client) AssetInspectionServiceCreateAssetInspection(params *AssetInspectionServiceCreateAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceCreateAssetInspectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetInspectionServiceCreateAssetInspectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AssetInspectionService_CreateAssetInspection",
		Method:             "POST",
		PathPattern:        "/api/asset_inspections/create_asset_inspection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetInspectionServiceCreateAssetInspectionReader{formats: a.formats},
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
	success, ok := result.(*AssetInspectionServiceCreateAssetInspectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssetInspectionServiceCreateAssetInspectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AssetInspectionServiceCreateDefectForAssetInspection asset inspection service create defect for asset inspection API
*/
func (a *Client) AssetInspectionServiceCreateDefectForAssetInspection(params *AssetInspectionServiceCreateDefectForAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceCreateDefectForAssetInspectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetInspectionServiceCreateDefectForAssetInspectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AssetInspectionService_CreateDefectForAssetInspection",
		Method:             "POST",
		PathPattern:        "/api/asset_inspections/create_defect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetInspectionServiceCreateDefectForAssetInspectionReader{formats: a.formats},
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
	success, ok := result.(*AssetInspectionServiceCreateDefectForAssetInspectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssetInspectionServiceCreateDefectForAssetInspectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AssetInspectionServiceDeleteAssetInspection asset inspection service delete asset inspection API
*/
func (a *Client) AssetInspectionServiceDeleteAssetInspection(params *AssetInspectionServiceDeleteAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceDeleteAssetInspectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetInspectionServiceDeleteAssetInspectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AssetInspectionService_DeleteAssetInspection",
		Method:             "POST",
		PathPattern:        "/api/asset_inspections/delete_asset_inspection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetInspectionServiceDeleteAssetInspectionReader{formats: a.formats},
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
	success, ok := result.(*AssetInspectionServiceDeleteAssetInspectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssetInspectionServiceDeleteAssetInspectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AssetInspectionServiceGetAssetInspection asset inspection service get asset inspection API
*/
func (a *Client) AssetInspectionServiceGetAssetInspection(params *AssetInspectionServiceGetAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceGetAssetInspectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetInspectionServiceGetAssetInspectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AssetInspectionService_GetAssetInspection",
		Method:             "GET",
		PathPattern:        "/api/asset_inspections/{value}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetInspectionServiceGetAssetInspectionReader{formats: a.formats},
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
	success, ok := result.(*AssetInspectionServiceGetAssetInspectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssetInspectionServiceGetAssetInspectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  AssetInspectionServiceUpdateAssetInspection asset inspection service update asset inspection API
*/
func (a *Client) AssetInspectionServiceUpdateAssetInspection(params *AssetInspectionServiceUpdateAssetInspectionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AssetInspectionServiceUpdateAssetInspectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAssetInspectionServiceUpdateAssetInspectionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AssetInspectionService_UpdateAssetInspection",
		Method:             "POST",
		PathPattern:        "/api/asset_inspections/update_asset_inspection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AssetInspectionServiceUpdateAssetInspectionReader{formats: a.formats},
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
	success, ok := result.(*AssetInspectionServiceUpdateAssetInspectionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AssetInspectionServiceUpdateAssetInspectionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
