package main

import (
	"github.com/ConquestSolutions/conquest_api.examples/conquest_api"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/action_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
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

	createAssetCommand := asset_service.NewAssetServiceCreateAssetParams()
	createAssetCommand.WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: "TestAssetWithAction",
	})
	createAssetResult, err := api.AssetService.AssetServiceCreateAsset(createAssetCommand, nil)
	require.NoError(t, err)
	facilityId := createAssetResult.GetPayload()

	// Create an Action for that Facility
	newDescription := "TestActionCRUD"

	createActionCommand := asset_service.NewAssetServiceCreateActionForAssetParams()
	createActionCommand.WithBody(&models.ConquestAPICreateActionForAssetCommand{
		AssetID:           facilityId,
		ActionDescription: newDescription,
	})
	createActionResult, err := api.AssetService.AssetServiceCreateActionForAsset(createActionCommand, nil)
	require.NoError(t, err)
	actionId := createActionResult.GetPayload()

	// Retrieve the object for viewing or editing

	getActionRequest1 := action_service.NewActionServiceGetActionParams()
	getActionRequest1.WithValue(actionId)
	getActionResponse1, err := api.ActionService.ActionServiceGetAction(getActionRequest1, nil)
	require.NoError(t, err)

	lastEdit := getActionResponse1.GetPayload().EditDate
	originalAction := getActionResponse1.GetPayload().Record
	require.Equal(t, newDescription, originalAction.ActionDescription)

	// Make an edit

	now := strfmt.DateTime(time.Now().UTC())
	newDescription += " (updated!)"
	updatedAction := *originalAction // shallow copy
	updatedAction.ActionDescription = newDescription
	updatedAction.StartDate = &models.ConquestAPITimestampValue{Value: now}

	updateActionCommand := action_service.NewActionServiceUpdateActionParams()
	updateActionCommand.WithBody(&models.ConquestAPIActionRecordChangeSet{
		ActionID: actionId,
		Changes: []string{
			"ActionDescription",
			"StartDate",
		},
		LastEdit: lastEdit,
		Original: originalAction,
		Updated:  &updatedAction,
	})
	updateActionResult, err := api.ActionService.ActionServiceUpdateAction(updateActionCommand, nil)
	require.NoError(t, err)
	require.True(t, updateActionResult.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getActionRequest2 := action_service.NewActionServiceGetActionParams()
	getActionRequest2.WithValue(actionId)
	getActionResponse2, err := api.ActionService.ActionServiceGetAction(getActionRequest2, nil)
	require.NoError(t, err)
	lastEdit = getActionResponse2.GetPayload().EditDate
	retrievedAction := getActionResponse2.GetPayload().Record
	require.Equal(t, newDescription, retrievedAction.ActionDescription)

	// Clean up our test data

	deleteActionCommand := action_service.NewActionServiceDeleteActionParams()
	deleteActionCommand.WithBody(&models.ConquestAPIDeleteActionCommand{
		ActionID: actionId,
	})
	_, err = api.ActionService.ActionServiceDeleteAction(deleteActionCommand, nil)
	require.NoError(t, err)

	// Confirm it's gone

	getActionRequest3 := action_service.NewActionServiceGetActionParams()
	getActionRequest3.WithValue(actionId)
	_, err = api.ActionService.ActionServiceGetAction(getActionRequest3, nil)
	rerr, ok := err.(conquest_api.RequestError)
	require.True(t, ok)
	require.Equal(t, 404, rerr.Code())
}
