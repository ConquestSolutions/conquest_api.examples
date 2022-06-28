package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"github.com/ConquestSolutions/conquest_api.examples/conquest_api"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/asset_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/client/document_service"
	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"testing"
)

// TestAddDocument is a complete demonstration of how to add a document to an object
func TestAddDocument(t *testing.T) {

	cfg, err := conquest_api.DefaultConfig()
	require.NoError(t, err)
	api, err := conquest_api.NewClient(*cfg)
	require.NoError(t, err)

	// Create a Facility (ParentID = 0)

	createAssetCommand := asset_service.NewAssetServiceCreateAssetParams()
	createAssetCommand.WithBody(&models.ConquestAPICreateAssetCommand{
		Proposed:         true,
		AssetDescription: "TestAddDocumentToAsset",
	})
	createAssetResult, err := api.AssetService.AssetServiceCreateAsset(createAssetCommand, nil)
	require.NoError(t, err)
	assetId := createAssetResult.GetPayload()

	fmt.Printf("Created asset with id %d", assetId)

	// Create a document record and get an upload link

	contentType := "image/png"
	addDocumentCommand := document_service.NewDocumentServiceAddDocumentParams()
	objectType := models.ConquestAPIObjectTypeObjectTypeAsset
	addDocumentCommand.WithBody(&models.ConquestAPIAddDocumentCommand{
		ObjectKey: &models.ConquestAPIObjectKey{
			ObjectType: &objectType,
			Int32Value: assetId,
		},
		DocumentDescription: "Test document",
		Address:             fmt.Sprintf("Asset/%d/TestAddDocument3.png", assetId),
		ContentType:         contentType,
	})

	addDocumentResult, err := api.DocumentService.DocumentServiceAddDocument(addDocumentCommand, nil)
	require.NoError(t, err)
	uploadInfo := addDocumentResult.GetPayload()

	f, err := os.Open("logo.png")
	require.NoError(t, err)

	// TODO why does a simple HTTP request not work
	// require.NoError(t, cfg.UploadFile(fileData, len(data), uploadInfo))

	// We upload to blob storage, use the azblob library to perform the upload
	cli, err := azblob.NewBlockBlobClientWithNoCredential(uploadInfo.UploadURI, nil)
	require.NoError(t, err)

	_, err = cli.UploadFile(context.TODO(), f, azblob.UploadOption{
		HTTPHeaders: &azblob.BlobHTTPHeaders{
			BlobContentType: &contentType,
		},
	})
	require.NoError(t, err)

	// get the document thumbnail

	generateThumbnailCommand := document_service.NewDocumentServiceGenerateDocumentLinkParams()
	generateThumbnailCommand.WithBody(&models.ConquestAPIGenerateDocumentLinkCommand{
		DocumentID:          uploadInfo.Document.DocumentID,
		ObjectKey:           uploadInfo.Document.ObjectKey,
		XThumbnailParameter: "medium",
	})
	generateThumbnailResult, err := api.DocumentService.DocumentServiceGenerateDocumentLink(generateThumbnailCommand, nil)
	require.NoError(t, err)

	linkResult := generateThumbnailResult.GetPayload()

	println("document link: " + linkResult.Link)

	ioutil.WriteFile("document-link.txt", []byte(linkResult.Link), os.ModePerm)

	println("wrote: " + "document-link.txt")
}
