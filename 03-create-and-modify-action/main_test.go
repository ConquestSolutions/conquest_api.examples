package main

import (
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/action_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestActionCRUD is a complete demonstration of how to
// - Create an Object
// - Retrieve an Object
// - Update an Object
// - Delete an Object
func TestActionCRUD(t *testing.T) {

	api, err := conquest_api.New()
	require.NoError(t, err)

	// Create a Facility (ParentID = 0)

	createAssetCommand := asset_service.NewCreateAssetParams().WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: "TestAssetWithAction",
	})
	createAssetResult, err := api.AssetService.CreateAsset(createAssetCommand, nil)
	require.NoError(t, err)
	facilityId := createAssetResult.GetPayload()

	// Create an Action for that Facility
	newDescription := "TestActionCRUD"

	createActionCommand := asset_service.NewCreateActionForAssetParams().WithBody(&models.ConquestAPICreateActionForAssetCommand{
		AssetID: facilityId,
		ActionDescription: newDescription,
	})
	createActionResult, err := api.AssetService.CreateActionForAsset(createActionCommand, nil)
	require.NoError(t, err)
	actionId := createActionResult.GetPayload()

	// Retrieve the object for viewing or editing

	getActionRequest := action_service.NewGetActionParams().WithValue(actionId)
	getActionResponse, err := api.ActionService.GetAction(getActionRequest, nil)
	require.NoError(t, err)

	lastEdit := getActionResponse.GetPayload().EditDate
	originalAction := getActionResponse.GetPayload().Record
	require.Equal(t, newDescription, originalAction.ActionDescription)

	// Make an edit

	now := strfmt.DateTime(time.Now().UTC())
	newDescription += " (updated!)"
	updatedAction := *originalAction // shallow copy
	updatedAction.ActionDescription = newDescription
	updatedAction.StartDate = &models.ConquestAPITimestampValue{Value: now}

	updateActionCommand := action_service.NewUpdateActionParams().WithBody(&models.ConquestAPIActionRecordChangeSet{
		ActionID: actionId,
		Changes: []string{
			"ActionDescription",
			"StartDate",
		},
		LastEdit: lastEdit,
		Original: originalAction,
		Updated:  &updatedAction,
	})
	updateActionResult, err := api.ActionService.UpdateAction(updateActionCommand, nil)
	require.NoError(t, err)
	require.True(t, updateActionResult.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getActionRequest = action_service.NewGetActionParams().WithValue(actionId)
	getActionResponse, err = api.ActionService.GetAction(getActionRequest, nil)
	require.NoError(t, err)
	lastEdit = getActionResponse.GetPayload().EditDate
	retrievedAction := getActionResponse.GetPayload().Record
	require.Equal(t, newDescription, retrievedAction.ActionDescription)

	// Clean up our test data

	deleteActionCommand := action_service.NewDeleteActionParams().WithBody(&models.ConquestAPIDeleteActionCommand{
		ActionID: actionId,
	})
	_, err = api.ActionService.DeleteAction(deleteActionCommand, nil)
	require.NoError(t, err)

	// Confirm it's gone

	getActionRequest = action_service.NewGetActionParams().WithValue(actionId)
	_, err = api.ActionService.GetAction(getActionRequest, nil)
	require.Error(t, err)
	require.Equal(t, conquest_api.ErrStatusCode(err), 404)
}
