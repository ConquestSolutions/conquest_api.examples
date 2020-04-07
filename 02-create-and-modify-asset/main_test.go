package main

import (
	"github.com/ConquestSolutions/apiv2.examples/conquest_api"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
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

	createAssetCommand := asset_service.NewCreateAssetParams().WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: newDescription,
	})
	createAssetResult, err := api.AssetService.CreateAsset(createAssetCommand, nil)
	require.NoError(t, err)
	facilityId := createAssetResult.GetPayload()

	// Retrieve the object for viewing or editing

	getAssetRequest := asset_service.NewGetAssetParams().WithValue(facilityId)
	getAssetResponse, err := api.AssetService.GetAsset(getAssetRequest, nil)
	require.NoError(t, err)

	lastEdit := getAssetResponse.GetPayload().EditDate
	originalAsset := getAssetResponse.GetPayload().Record
	require.Equal(t, newDescription, originalAsset.AssetDescription)

	// Make an edit

	now := strfmt.DateTime(time.Now().UTC())
	newDescription += " (updated!)"
	updatedAsset := *originalAsset // shallow copy
	updatedAsset.AssetDescription = newDescription
	updatedAsset.YearCreated = &models.ConquestAPITimestampValue{Value: now}

	updateAssetCommand := asset_service.NewUpdateAssetParams().WithBody(&models.ConquestAPIAssetRecordChangeSet{
		AssetID: facilityId,
		Changes: []string{
			"AssetDescription",
			"YearCreated",
		},
		LastEdit: lastEdit,
		Original: originalAsset,
		Updated:  &updatedAsset,
	})
	updateAssetResult, err := api.AssetService.UpdateAsset(updateAssetCommand, nil)
	require.NoError(t, err)
	require.True(t, updateAssetResult.GetPayload().Accepted)

	// Retrieve the updated object to verify the change

	getAssetRequest = asset_service.NewGetAssetParams().WithValue(facilityId)
	getAssetResponse, err = api.AssetService.GetAsset(getAssetRequest, nil)
	require.NoError(t, err)
	lastEdit = getAssetResponse.GetPayload().EditDate
	retrievedAsset := getAssetResponse.GetPayload().Record
	require.Equal(t, newDescription, retrievedAsset.AssetDescription)
	require.NotNil(t, retrievedAsset.YearCreated)

	// BUG: There will be a few millisecond difference. Be sure to compare with a rounded value until fixed.
	roundedNow := time.Time(now).Round(1 * time.Second)
	roundedYearCreated := time.Time(retrievedAsset.YearCreated.Value).Round(1 * time.Second)
	require.Equal(t, roundedNow, roundedYearCreated)

	// Unset the year created
	updatedAsset = *retrievedAsset // shallow copy
	updatedAsset.YearCreated = nil
	updateAssetCommand = asset_service.NewUpdateAssetParams().WithBody(&models.ConquestAPIAssetRecordChangeSet{
		AssetID: facilityId,
		Changes: []string{
			"YearCreated",
		},
		LastEdit: lastEdit,
		Original: retrievedAsset,
		Updated:  &updatedAsset,
	})
	updateAssetResult, err = api.AssetService.UpdateAsset(updateAssetCommand, nil)
	require.NoError(t, err)
	require.True(t, updateAssetResult.GetPayload().Accepted)

	// Retrieve the updated object and verify the change

	getAssetRequest = asset_service.NewGetAssetParams().WithValue(facilityId)
	getAssetResponse, err = api.AssetService.GetAsset(getAssetRequest, nil)
	require.NoError(t, err)
	retrievedAsset = getAssetResponse.GetPayload().Record
	require.Nil(t, retrievedAsset.YearCreated)

	// Clean up our test data

	deleteAssetCommand := asset_service.NewDeleteAssetParams().WithBody(&models.ConquestAPIDeleteAssetCommand{
		AssetID: facilityId,
	})
	_, err = api.AssetService.DeleteAsset(deleteAssetCommand, nil)
	require.NoError(t, err)

	// Confirm it's gone

	getAssetRequest = asset_service.NewGetAssetParams().WithValue(facilityId)
	_, err = api.AssetService.GetAsset(getAssetRequest, nil)
	require.Error(t, err)
	require.Equal(t, conquest_api.ErrStatusCode(err), 404)
}
