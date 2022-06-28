package conquest_api

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ConquestSolutions/conquest_api.examples/go-swagger/models"
)

// Implements an upload strategy based on parameters provided by the API
//
// The API will supply a method, an endpoint and required headers that an upload is parameterised.
func (config *Config) UploadFile(body io.ReadCloser, contentLength int, uploadInfo *models.ConquestAPIAddDocumentResult) error {
	switch strings.ToUpper(uploadInfo.UploadMethod) {
	// Preferred
	case "PUT":
		return config.uploadFileViaPut(body, contentLength, uploadInfo)
	case "POST":
		return config.uploadFileViaPost(body, contentLength, uploadInfo)
	}
	return fmt.Errorf("unknown upload method: " + uploadInfo.UploadMethod)
}

// curl --upload-file "${ApiHost}${UploadUri}"
func (config *Config) uploadFileViaPut(file io.ReadCloser, contentLength int, uploadInfo *models.ConquestAPIAddDocumentResult) error {

	// Instantiate a request with the provided parameters.
	httpReq, err := http.NewRequest(http.MethodPut, uploadInfo.UploadURI, file)
	if err != nil {
		return err
	}
	for k, v := range uploadInfo.UploadHeaders {
		// Don't use Add, header is case sensitive for
		httpReq.Header[k] = []string{v}
	}
	httpReq.Header.Set("Content-Type", uploadInfo.Document.ContentType)
	httpReq.Header.Set("Content-Length", strconv.Itoa(contentLength))

	// Send
	return executeHttpRequest(config.InsecureSkipVerify, httpReq, &Empty{})
}

// curl -i -X POST -H "Authorization: Bearer ACCESS_TOKEN" -H "Content-Type: multipart/form-data" -F "document=@inspection-photo.png" "${ApiHost}${UploadUri}"
func (config *Config) uploadFileViaPost(file io.ReadCloser, contentLength int, uploadInfo *models.ConquestAPIAddDocumentResult) error {

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
		httpReq.Header[k] = []string{v}
	}
	if useCredentials {
		httpReq.Header.Add("Authorization", "Bearer "+config.AccessToken)
	}
	httpReq.Header.Set("Content-Type", writer.FormDataContentType())
	httpReq.Header.Set("Content-Length", strconv.Itoa(contentLength))

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
	if httpResp.StatusCode != http.StatusOK {
		content := ""
		if httpResp.ContentLength > 0 {
			if data, err := ioutil.ReadAll(httpResp.Body); err == nil {
				content = string(data)
			}
		}
		return UnknownResponse{
			Status:     httpResp.Status,
			StatusCode: httpResp.StatusCode,
			Content:    content,
		}
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
		Status:      httpResp.Status,
		StatusCode:  httpResp.StatusCode,
		ContentType: contentType,
		Content:     respData.String(),
	}
}
