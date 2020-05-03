package conquest_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/ConquestSolutions/apiv2.examples/go-swagger/models"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"reflect"
	"strings"
)

// Implements an upload strategy based on parameters provided by the API
//
// The API will supply a method, an endpoint and required headers that an upload is parameterised.
func (config *Config) UploadFile(body io.ReadCloser, uploadInfo *models.ConquestAPIAddDocumentResult) error {
	switch strings.ToUpper(uploadInfo.UploadMethod) {
	// Preferred
	case "PUT":
		return config.uploadFileViaPut(body, uploadInfo)
	case "POST":
		return config.uploadFileViaPost(body, uploadInfo)
	}
	return fmt.Errorf("unknown upload method: " + uploadInfo.UploadMethod)
}

// curl --upload-file "${ApiHost}${UploadUri}"
func (config *Config) uploadFileViaPut(file io.ReadCloser, uploadInfo *models.ConquestAPIAddDocumentResult) error {

	// Instantiate a request with the provided parameters.
	httpReq, err := http.NewRequest(http.MethodPut, uploadInfo.UploadURI, file)
	if err != nil {
		return err
	}
	for k, v := range uploadInfo.UploadHeaders {
		httpReq.Header.Add(k, v)
	}
	httpReq.Header.Add("Content-Type", uploadInfo.Document.ContentType)

	// Send
	return executeHttpRequest(config.InsecureSkipVerify, httpReq, &Empty{})
}

// curl -i -X POST -H "Authorization: Bearer ACCESS_TOKEN" -H "Content-Type: multipart/form-data" -F "document=@inspection-photo.png" "${ApiHost}${UploadUri}"
func (config *Config) uploadFileViaPost(file io.ReadCloser, uploadInfo *models.ConquestAPIAddDocumentResult) error {

	useCredentials := false

	// "${ApiHost}${UploadUri}"
	uploadEndpoint := uploadInfo.UploadURI
	if strings.HasPrefix(uploadEndpoint, "/") {
		useCredentials = true
		uploadEndpoint = strings.TrimSuffix(config.Address, "/") + uploadInfo.UploadURI
	}

	// Build a multi-part request body
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("document", filepath.Base("my-file"))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)
	writer.Close()

	// Instantiate a request with the provided parameters.
	httpReq, err := http.NewRequest(http.MethodPost, uploadEndpoint, body)
	if err != nil {
		return err
	}
	for k, v := range uploadInfo.UploadHeaders {
		httpReq.Header.Add(k, v)
	}
	if useCredentials {
		httpReq.Header.Add("Authorization", "Bearer "+config.AccessToken)
	}
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())

	// Send
	return executeHttpRequest(config.InsecureSkipVerify, httpReq, &Empty{})
}

func executeHttpRequest(insecureSkipVerify bool, httpReq *http.Request, resp interface{}) error {
	cli := http.DefaultClient
	if insecureSkipVerify {
		cli.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: insecureSkipVerify,
			},
		}
	}
	httpResp, err := cli.Do(httpReq)
	if err != nil {
		return err
	}

	// Handle the case where nothing was returned
	respType := reflect.TypeOf(resp)
	if httpResp.ContentLength == 0 ||
		respType.Elem() == reflect.TypeOf(Empty{}) {
		if httpResp.StatusCode != http.StatusOK {
			return UnknownResponse{
				Status: httpResp.StatusCode,
			}
		}
		return nil
	}

	// Read the data
	var respData bytes.Buffer
	_, err = io.Copy(&respData, httpResp.Body)
	if err != nil {
		return err
	}

	// Handle json, if we understand the response
	contentType := httpResp.Header.Get("Content-Type")
	if strings.Contains(contentType, "application/json") {
		switch httpResp.StatusCode {
		case http.StatusOK:
			return json.Unmarshal(respData.Bytes(), resp)
		case http.StatusBadRequest, http.StatusForbidden, http.StatusUnauthorized:
			var errResp ErrorResponse
			jsonErr := json.Unmarshal(respData.Bytes(), errResp)
			if jsonErr == nil && errResp.ErrorCode != "" {
				return errResp
			}
		}
	}

	// Report anything that we cannot handle as a general error.
	// That could be handled later.

	return UnknownResponse{
		Status:      httpResp.StatusCode,
		ContentType: contentType,
		Content:     respData.String(),
	}
}
