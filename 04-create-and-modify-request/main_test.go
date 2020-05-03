package main

import (
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/request_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
	"github.com/stretchr/testify/require"
	"testing"
)

// TestRequestCRUD is a complete demonstration of how to
// - Create an Object
// - Retrieve an Object
// - Update an Object
// - Delete an Object
func TestRequestCRUD(t *testing.T) {

	api, err := conquest_api.New()
	require.NoError(t, err)

	createRequestCommand := request_service.NewCreateRequestParams().WithBody(&models.ConquestAPICreateRequestCommand{
		ChangeSet: &models.ConquestAPIRequestRecordChangeSet{
			Changes: []string{
				"RequestDetail",
				"RequestorName",
			},
			Updated: &models.ConquestAPIRequestRecord{
				RequestDetail: "There is an issue with ...",
				RequestorName: "John Smith",
			},
		},
	})

	createRequestResult, err := api.RequestService.CreateRequest(createRequestCommand, nil)
	require.NoError(t, err)
	arqId := createRequestResult.GetPayload()

	// Retrieve the object for viewing or editing

	getRequestRequest := request_service.NewGetRequestParams().WithValue(arqId)
	_, err = api.RequestService.GetRequest(getRequestRequest, nil)
	require.NoError(t, err) // // in API 2.0.10 this returns 404?
}
