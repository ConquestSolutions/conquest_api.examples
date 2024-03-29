// Code generated by go-swagger; DO NOT EDIT.

package document_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new document service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for document service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DocumentServiceAddDocument(params *DocumentServiceAddDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceAddDocumentOK, error)

	DocumentServiceChangeDocumentContent(params *DocumentServiceChangeDocumentContentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceChangeDocumentContentOK, error)

	DocumentServiceChangeDocumentDescription(params *DocumentServiceChangeDocumentDescriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceChangeDocumentDescriptionOK, error)

	DocumentServiceGenerateDocumentLink(params *DocumentServiceGenerateDocumentLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceGenerateDocumentLinkOK, error)

	DocumentServiceGetDocuments(params *DocumentServiceGetDocumentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceGetDocumentsOK, error)

	DocumentServiceListDocumentLocations(params *DocumentServiceListDocumentLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceListDocumentLocationsOK, error)

	DocumentServiceRemoveDocument(params *DocumentServiceRemoveDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceRemoveDocumentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DocumentServiceAddDocument document service add document API
*/
func (a *Client) DocumentServiceAddDocument(params *DocumentServiceAddDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceAddDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceAddDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_AddDocument",
		Method:             "POST",
		PathPattern:        "/api/documents/add_document",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceAddDocumentReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceAddDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceAddDocumentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DocumentServiceChangeDocumentContent document service change document content API
*/
func (a *Client) DocumentServiceChangeDocumentContent(params *DocumentServiceChangeDocumentContentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceChangeDocumentContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceChangeDocumentContentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_ChangeDocumentContent",
		Method:             "POST",
		PathPattern:        "/api/documents/change_document_content",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceChangeDocumentContentReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceChangeDocumentContentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceChangeDocumentContentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DocumentServiceChangeDocumentDescription document service change document description API
*/
func (a *Client) DocumentServiceChangeDocumentDescription(params *DocumentServiceChangeDocumentDescriptionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceChangeDocumentDescriptionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceChangeDocumentDescriptionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_ChangeDocumentDescription",
		Method:             "POST",
		PathPattern:        "/api/documents/change_document_description",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceChangeDocumentDescriptionReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceChangeDocumentDescriptionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceChangeDocumentDescriptionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DocumentServiceGenerateDocumentLink document service generate document link API
*/
func (a *Client) DocumentServiceGenerateDocumentLink(params *DocumentServiceGenerateDocumentLinkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceGenerateDocumentLinkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceGenerateDocumentLinkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_GenerateDocumentLink",
		Method:             "POST",
		PathPattern:        "/api/documents/generate_document_link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceGenerateDocumentLinkReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceGenerateDocumentLinkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceGenerateDocumentLinkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DocumentServiceGetDocuments document service get documents API
*/
func (a *Client) DocumentServiceGetDocuments(params *DocumentServiceGetDocumentsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceGetDocumentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceGetDocumentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_GetDocuments",
		Method:             "POST",
		PathPattern:        "/api/documents/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceGetDocumentsReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceGetDocumentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceGetDocumentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DocumentServiceListDocumentLocations document service list document locations API
*/
func (a *Client) DocumentServiceListDocumentLocations(params *DocumentServiceListDocumentLocationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceListDocumentLocationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceListDocumentLocationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_ListDocumentLocations",
		Method:             "POST",
		PathPattern:        "/api/documents/list_locations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceListDocumentLocationsReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceListDocumentLocationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceListDocumentLocationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DocumentServiceRemoveDocument document service remove document API
*/
func (a *Client) DocumentServiceRemoveDocument(params *DocumentServiceRemoveDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DocumentServiceRemoveDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentServiceRemoveDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DocumentService_RemoveDocument",
		Method:             "POST",
		PathPattern:        "/api/documents/remove_document",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentServiceRemoveDocumentReader{formats: a.formats},
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
	success, ok := result.(*DocumentServiceRemoveDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DocumentServiceRemoveDocumentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
