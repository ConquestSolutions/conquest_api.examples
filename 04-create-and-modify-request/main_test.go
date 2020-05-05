package main

import (
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/request_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/view_service"
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

	// Find all top-level orgs (parent = 0)
	listTopLevelOrgUnitsQuery := view_service.NewListHierarchyNodesParams().WithBody(&models.ConquestAPIListHierarchyNodesQuery{
		ObjectKey: &models.ConquestAPIObjectKey{
			ObjectType: models.ConquestAPIObjectTypeObjectTypeOrganisationUnit,
			Int32Value: 0,
		},
	})
	listTopLevelOrgUnitsResult, err := api.ViewService.ListHierarchyNodes(listTopLevelOrgUnitsQuery, nil)
	require.NoError(t, err)

	tree := listTopLevelOrgUnitsResult.GetPayload().Headers
	require.GreaterOrEqual(t, len(tree), 1)

	orgUnitId := tree[0].ObjectKey.Int32Value

	createRequestCommand := request_service.NewCreateRequestParams().WithBody(&models.ConquestAPICreateRequestCommand{
		ChangeSet: &models.ConquestAPIRequestRecordChangeSet{
			Changes: []string{
				"RequestDetail",
				"RequestorName",
				"OrganisationUnitID",
			},
			Updated: &models.ConquestAPIRequestRecord{
				RequestDetail:      "There is an issue with ...",
				RequestorName:      "John Smith",
				OrganisationUnitID: orgUnitId,
			},
		},
	})

	createRequestResult, err := api.RequestService.CreateRequest(createRequestCommand, nil)
	require.NoError(t, err)
	arqId := createRequestResult.GetPayload()

	// Retrieve the object for viewing or editing

	getRequestRequest := request_service.NewGetRequestParams().WithValue(arqId)
	getRequestResponse, err := api.RequestService.GetRequest(getRequestRequest, nil)
	require.NoError(t, err) // // in API 2.0.10 this returns 404?

	lastEdit := getRequestResponse.GetPayload().EditDate
	originalRequest := getRequestResponse.GetPayload().Record
	require.Equal(t, "John Smith", originalRequest.RequestorName)

	// Make an edit

	newDetail := originalRequest.RequestDetail + " (updated!)"
	updatedRequest := *originalRequest // shallow copy
	updatedRequest.RequestDetail = newDetail

	updateRequestCommand := request_service.NewUpdateRequestParams().WithBody(&models.ConquestAPIRequestRecordChangeSet{
		RequestID: arqId,
		Changes: []string{
			"RequestDetail",
		},
		LastEdit: lastEdit,
		Original: originalRequest,
		Updated:  &updatedRequest,
	})
	updateRequestResult, err := api.RequestService.UpdateRequest(updateRequestCommand, nil)
	require.NoError(t, err)
	require.True(t, updateRequestResult.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getRequestRequest = request_service.NewGetRequestParams().WithValue(arqId)
	getRequestResponse, err = api.RequestService.GetRequest(getRequestRequest, nil)
	require.NoError(t, err)
	lastEdit = getRequestResponse.GetPayload().EditDate
	retrievedRequest := getRequestResponse.GetPayload().Record
	require.Equal(t, newDetail, retrievedRequest.RequestDetail)

	// Clean up our test data

	deleteRequestCommand := request_service.NewDeleteRequestParams().WithBody(&models.ConquestAPIDeleteRequestCommand{
		RequestID: arqId,
	})
	_, err = api.RequestService.DeleteRequest(deleteRequestCommand, nil)
	require.NoError(t, err)

	// Confirm it's gone

	getRequestRequest = request_service.NewGetRequestParams().WithValue(arqId)
	_, err = api.RequestService.GetRequest(getRequestRequest, nil)
	require.Error(t, err)
	require.Equal(t, conquest_api.ErrStatusCode(err), 404)
}
