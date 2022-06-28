package main

import (
	"github.com/ConquestSolutions/conquest_api.examples/conquest_api"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/request_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/view_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
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
	listTopLevelOrgUnitsQuery := view_service.NewViewServiceListHierarchyNodesParams()
	objectType := models.ConquestAPIObjectTypeObjectTypeOrganisationUnit
	listTopLevelOrgUnitsQuery.WithBody(&models.ConquestAPIListHierarchyNodesQuery{
		ObjectKey: &models.ConquestAPIObjectKey{
			ObjectType: &objectType,
			Int32Value: 0,
		},
	})
	listTopLevelOrgUnitsResult, err := api.ViewService.ViewServiceListHierarchyNodes(listTopLevelOrgUnitsQuery, nil)
	require.NoError(t, err)

	tree := listTopLevelOrgUnitsResult.GetPayload().Headers
	require.GreaterOrEqual(t, len(tree), 1)

	orgUnitId := tree[0].ObjectKey.Int32Value

	createRequestCommand := request_service.NewRequestServiceCreateRequestParams()
	createRequestCommand.WithBody(&models.ConquestAPICreateRequestCommand{
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

	createRequestResult, err := api.RequestService.RequestServiceCreateRequest(createRequestCommand, nil)
	require.NoError(t, err)
	arqId := createRequestResult.GetPayload()

	// Retrieve the object for viewing or editing

	getRequestRequest1 := request_service.NewRequestServiceGetRequestParams()
	getRequestRequest1.WithValue(arqId)
	getRequestResponse, err := api.RequestService.RequestServiceGetRequest(getRequestRequest1, nil)
	require.NoError(t, err) // // in API 2.0.10 this returns 404?

	lastEdit := getRequestResponse.GetPayload().EditDate
	originalRequest := getRequestResponse.GetPayload().Record
	require.Equal(t, "John Smith", originalRequest.RequestorName)

	// Make an edit

	newDetail := originalRequest.RequestDetail + " (updated!)"
	updatedRequest := *originalRequest // shallow copy
	updatedRequest.RequestDetail = newDetail

	updateRequestCommand1 := request_service.NewRequestServiceUpdateRequestParams()
	updateRequestCommand1.WithBody(&models.ConquestAPIRequestRecordChangeSet{
		RequestID: arqId,
		Changes: []string{
			"RequestDetail",
		},
		LastEdit: lastEdit,
		Original: originalRequest,
		Updated:  &updatedRequest,
	})
	updateRequestResult1, err := api.RequestService.RequestServiceUpdateRequest(updateRequestCommand1, nil)
	require.NoError(t, err)
	require.True(t, updateRequestResult1.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getRequestRequest2 := request_service.NewRequestServiceGetRequestParams()
	getRequestRequest2.WithValue(arqId)
	getRequestResponse2, err := api.RequestService.RequestServiceGetRequest(getRequestRequest2, nil)
	require.NoError(t, err)
	lastEdit = getRequestResponse2.GetPayload().EditDate
	retrievedRequest := getRequestResponse2.GetPayload().Record
	require.Equal(t, newDetail, retrievedRequest.RequestDetail)

	// Clean up our test data

	deleteRequestCommand := request_service.NewRequestServiceDeleteRequestParams()
	deleteRequestCommand.WithBody(&models.ConquestAPIDeleteRequestCommand{
		RequestID: arqId,
	})
	_, err = api.RequestService.RequestServiceDeleteRequest(deleteRequestCommand, nil)
	require.NoError(t, err)

	// Confirm it's gone

	getRequestRequest3 := request_service.NewRequestServiceGetRequestParams()
	getRequestRequest3.WithValue(arqId)
	_, err = api.RequestService.RequestServiceGetRequest(getRequestRequest3, nil)
	rerr, ok := err.(conquest_api.RequestError)
	require.True(t, ok)
	require.Equal(t, 404, rerr.Code())
}
