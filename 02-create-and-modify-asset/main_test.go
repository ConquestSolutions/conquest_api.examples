package main

import (
	"github.com/ConquestSolutions/conquest_api.examples/conquest_api"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

// TestAssetCRUD is a complete demonstration of how to
// - Create an Object
// - Retrieve an Object
// - Update an Object
// - Delete an Object
func TestAssetCRUD(t *testing.T) {

	api, err := conquest_api.New()
	require.NoError(t, err)

	// Create a Facility (ParentID = 0)

	newDescription := "TestAssetCRUD"

	createAssetCommand := asset_service.NewAssetServiceCreateAssetParams()
	createAssetCommand.WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: newDescription,
	})
	createAssetResult, err := api.AssetService.AssetServiceCreateAsset(createAssetCommand, nil)
	require.NoError(t, err)
	facilityId := createAssetResult.GetPayload()

	// Retrieve the object for viewing or editing

	getAssetRequest1 := asset_service.NewAssetServiceGetAssetParams()
	getAssetRequest1.WithValue(facilityId)
	getAssetResponse1, err := api.AssetService.AssetServiceGetAsset(getAssetRequest1, nil)
	require.NoError(t, err)

	lastEdit := getAssetResponse1.GetPayload().EditDate
	originalAsset := getAssetResponse1.GetPayload().Record
	require.Equal(t, newDescription, originalAsset.AssetDescription)

	// Make an edit

	now := strfmt.DateTime(time.Now().UTC())
	newDescription += " (updated!)"
	updatedAsset := *originalAsset // shallow copy
	updatedAsset.AssetDescription = newDescription
	updatedAsset.YearCreated = &models.ConquestAPITimestampValue{Value: now}

	updateAssetCommand1 := asset_service.NewAssetServiceUpdateAssetParams()
	updateAssetCommand1.WithBody(&models.ConquestAPIAssetRecordChangeSet{
		AssetID: facilityId,
		Changes: []string{
			"AssetDescription",
			"YearCreated",
		},
		LastEdit: lastEdit,
		Original: originalAsset,
		Updated:  &updatedAsset,
	})
	updateAssetResult1, err := api.AssetService.AssetServiceUpdateAsset(updateAssetCommand1, nil)
	require.NoError(t, err)
	require.True(t, updateAssetResult1.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getAssetRequest2 := asset_service.NewAssetServiceGetAssetParams()
	getAssetRequest2 = getAssetRequest1.WithValue(facilityId)
	getAssetResponse2, err := api.AssetService.AssetServiceGetAsset(getAssetRequest2, nil)
	require.NoError(t, err)
	lastEdit = getAssetResponse2.GetPayload().EditDate
	retrievedAsset := getAssetResponse2.GetPayload().Record
	require.Equal(t, newDescription, retrievedAsset.AssetDescription)
	require.NotNil(t, retrievedAsset.YearCreated)

	// BUG: There will be a few millisecond difference. Be sure to compare with a rounded value until fixed.
	roundedNow := time.Time(now).Round(1 * time.Second)
	roundedYearCreated := time.Time(retrievedAsset.YearCreated.Value).Round(1 * time.Second)
	require.Equal(t, roundedNow, roundedYearCreated)

	// Unset the year created
	updatedAsset = *retrievedAsset // shallow copy
	updatedAsset.YearCreated = nil

	updateAssetCommand2 := asset_service.NewAssetServiceUpdateAssetParams()
	updateAssetCommand2 = updateAssetCommand2.WithBody(&models.ConquestAPIAssetRecordChangeSet{
		AssetID: facilityId,
		Changes: []string{
			"YearCreated",
		},
		LastEdit: lastEdit,
		Original: retrievedAsset,
		Updated:  &updatedAsset,
	})
	updateAssetResult2, err := api.AssetService.AssetServiceUpdateAsset(updateAssetCommand2, nil)
	require.NoError(t, err)
	require.True(t, updateAssetResult2.GetPayload().Accepted)

	// Retrieve the updated object and verify the change

	getAssetRequest3 := asset_service.NewAssetServiceGetAssetParams()
	getAssetRequest3 = getAssetRequest3.WithValue(facilityId)
	getAssetResponse3, err := api.AssetService.AssetServiceGetAsset(getAssetRequest3, nil)
	require.NoError(t, err)
	retrievedAsset = getAssetResponse3.GetPayload().Record
	require.Nil(t, retrievedAsset.YearCreated)

	// Clean up our test data

	deleteAssetCommand := asset_service.NewAssetServiceDeleteAssetParams()
	deleteAssetCommand.WithBody(&models.ConquestAPIDeleteAssetCommand{
		AssetID: facilityId,
	})
	_, err = api.AssetService.AssetServiceDeleteAsset(deleteAssetCommand, nil)
	require.NoError(t, err)

	// Confirm it's gone

	getAssetRequest4 := asset_service.NewAssetServiceGetAssetParams()
	getAssetRequest4.WithValue(facilityId)
	_, err = api.AssetService.AssetServiceGetAsset(getAssetRequest4, nil)
	rerr, ok := err.(conquest_api.RequestError)
	require.True(t, ok)
	require.Equal(t, 404, rerr.Code())
}
