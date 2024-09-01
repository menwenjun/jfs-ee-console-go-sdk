/*
JuiceFS console API

API of the JuiceFS console (https://juicefs.com/api/v1)

API version: 0.0.1
Contact: team@juicedata.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// VolumeACLAPIService VolumeACLAPI service
type VolumeACLAPIService service

type ApiVolumesVolumeIDExportsGetRequest struct {
	ctx context.Context
	ApiService *VolumeACLAPIService
	volumeID int32
	token *string
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiVolumesVolumeIDExportsGetRequest) Token(token string) ApiVolumesVolumeIDExportsGetRequest {
	r.token = &token
	return r
}

func (r ApiVolumesVolumeIDExportsGetRequest) Execute() ([]GettableVolumeExport, *http.Response, error) {
	return r.ApiService.VolumesVolumeIDExportsGetExecute(r)
}

/*
VolumesVolumeIDExportsGet Get ACL list

List all exports(ACL rules) of this volume

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param volumeID ID of the JuiceFS volume
 @return ApiVolumesVolumeIDExportsGetRequest
*/
func (a *VolumeACLAPIService) VolumesVolumeIDExportsGet(ctx context.Context, volumeID int32) ApiVolumesVolumeIDExportsGetRequest {
	return ApiVolumesVolumeIDExportsGetRequest{
		ApiService: a,
		ctx: ctx,
		volumeID: volumeID,
	}
}

// Execute executes the request
//  @return []GettableVolumeExport
func (a *VolumeACLAPIService) VolumesVolumeIDExportsGetExecute(r ApiVolumesVolumeIDExportsGetRequest) ([]GettableVolumeExport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GettableVolumeExport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumeACLAPIService.VolumesVolumeIDExportsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeID}/exports"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeID"+"}", url.PathEscape(parameterValueToString(r.volumeID, "volumeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.volumeID < 1 {
		return localVarReturnValue, nil, reportError("volumeID must be greater than 1")
	}

	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVolumesVolumeIDExportsPostRequest struct {
	ctx context.Context
	ApiService *VolumeACLAPIService
	volumeID int32
	token *string
	postableVolumeExport *PostableVolumeExport
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiVolumesVolumeIDExportsPostRequest) Token(token string) ApiVolumesVolumeIDExportsPostRequest {
	r.token = &token
	return r
}

func (r ApiVolumesVolumeIDExportsPostRequest) PostableVolumeExport(postableVolumeExport PostableVolumeExport) ApiVolumesVolumeIDExportsPostRequest {
	r.postableVolumeExport = &postableVolumeExport
	return r
}

func (r ApiVolumesVolumeIDExportsPostRequest) Execute() ([]GettableVolumeExport, *http.Response, error) {
	return r.ApiService.VolumesVolumeIDExportsPostExecute(r)
}

/*
VolumesVolumeIDExportsPost Create ACL

Create an export(ACL rule)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param volumeID ID of the JuiceFS volume
 @return ApiVolumesVolumeIDExportsPostRequest
*/
func (a *VolumeACLAPIService) VolumesVolumeIDExportsPost(ctx context.Context, volumeID int32) ApiVolumesVolumeIDExportsPostRequest {
	return ApiVolumesVolumeIDExportsPostRequest{
		ApiService: a,
		ctx: ctx,
		volumeID: volumeID,
	}
}

// Execute executes the request
//  @return []GettableVolumeExport
func (a *VolumeACLAPIService) VolumesVolumeIDExportsPostExecute(r ApiVolumesVolumeIDExportsPostRequest) ([]GettableVolumeExport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GettableVolumeExport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumeACLAPIService.VolumesVolumeIDExportsPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeID}/exports"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeID"+"}", url.PathEscape(parameterValueToString(r.volumeID, "volumeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.volumeID < 1 {
		return localVarReturnValue, nil, reportError("volumeID must be greater than 1")
	}

	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postableVolumeExport
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest struct {
	ctx context.Context
	ApiService *VolumeACLAPIService
	volumeID int32
	volumeExportID int32
	token *string
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest) Token(token string) ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest {
	r.token = &token
	return r
}

func (r ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.VolumesVolumeIDExportsVolumeExportIDDeleteExecute(r)
}

/*
VolumesVolumeIDExportsVolumeExportIDDelete Delete ACL

Delete export(ACL rule) by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param volumeID ID of the JuiceFS volume
 @param volumeExportID ID of the JuiceFS volume export
 @return ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest
*/
func (a *VolumeACLAPIService) VolumesVolumeIDExportsVolumeExportIDDelete(ctx context.Context, volumeID int32, volumeExportID int32) ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest {
	return ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest{
		ApiService: a,
		ctx: ctx,
		volumeID: volumeID,
		volumeExportID: volumeExportID,
	}
}

// Execute executes the request
func (a *VolumeACLAPIService) VolumesVolumeIDExportsVolumeExportIDDeleteExecute(r ApiVolumesVolumeIDExportsVolumeExportIDDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumeACLAPIService.VolumesVolumeIDExportsVolumeExportIDDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeID}/exports/{volumeExportID}"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeID"+"}", url.PathEscape(parameterValueToString(r.volumeID, "volumeID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"volumeExportID"+"}", url.PathEscape(parameterValueToString(r.volumeExportID, "volumeExportID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.volumeID < 1 {
		return nil, reportError("volumeID must be greater than 1")
	}
	if r.volumeExportID < 1 {
		return nil, reportError("volumeExportID must be greater than 1")
	}

	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiVolumesVolumeIDExportsVolumeExportIDGetRequest struct {
	ctx context.Context
	ApiService *VolumeACLAPIService
	volumeID int32
	volumeExportID int32
	token *string
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiVolumesVolumeIDExportsVolumeExportIDGetRequest) Token(token string) ApiVolumesVolumeIDExportsVolumeExportIDGetRequest {
	r.token = &token
	return r
}

func (r ApiVolumesVolumeIDExportsVolumeExportIDGetRequest) Execute() (*GettableVolumeExport, *http.Response, error) {
	return r.ApiService.VolumesVolumeIDExportsVolumeExportIDGetExecute(r)
}

/*
VolumesVolumeIDExportsVolumeExportIDGet Get ACL by ID

Get export(ACL rule) by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param volumeID ID of the JuiceFS volume
 @param volumeExportID ID of the JuiceFS volume export
 @return ApiVolumesVolumeIDExportsVolumeExportIDGetRequest
*/
func (a *VolumeACLAPIService) VolumesVolumeIDExportsVolumeExportIDGet(ctx context.Context, volumeID int32, volumeExportID int32) ApiVolumesVolumeIDExportsVolumeExportIDGetRequest {
	return ApiVolumesVolumeIDExportsVolumeExportIDGetRequest{
		ApiService: a,
		ctx: ctx,
		volumeID: volumeID,
		volumeExportID: volumeExportID,
	}
}

// Execute executes the request
//  @return GettableVolumeExport
func (a *VolumeACLAPIService) VolumesVolumeIDExportsVolumeExportIDGetExecute(r ApiVolumesVolumeIDExportsVolumeExportIDGetRequest) (*GettableVolumeExport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GettableVolumeExport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumeACLAPIService.VolumesVolumeIDExportsVolumeExportIDGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeID}/exports/{volumeExportID}"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeID"+"}", url.PathEscape(parameterValueToString(r.volumeID, "volumeID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"volumeExportID"+"}", url.PathEscape(parameterValueToString(r.volumeExportID, "volumeExportID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.volumeID < 1 {
		return localVarReturnValue, nil, reportError("volumeID must be greater than 1")
	}
	if r.volumeExportID < 1 {
		return localVarReturnValue, nil, reportError("volumeExportID must be greater than 1")
	}

	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiVolumesVolumeIDExportsVolumeExportIDPutRequest struct {
	ctx context.Context
	ApiService *VolumeACLAPIService
	volumeID int32
	volumeExportID int32
	token *string
	postableVolumeExport *PostableVolumeExport
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiVolumesVolumeIDExportsVolumeExportIDPutRequest) Token(token string) ApiVolumesVolumeIDExportsVolumeExportIDPutRequest {
	r.token = &token
	return r
}

func (r ApiVolumesVolumeIDExportsVolumeExportIDPutRequest) PostableVolumeExport(postableVolumeExport PostableVolumeExport) ApiVolumesVolumeIDExportsVolumeExportIDPutRequest {
	r.postableVolumeExport = &postableVolumeExport
	return r
}

func (r ApiVolumesVolumeIDExportsVolumeExportIDPutRequest) Execute() (*GettableVolumeExport, *http.Response, error) {
	return r.ApiService.VolumesVolumeIDExportsVolumeExportIDPutExecute(r)
}

/*
VolumesVolumeIDExportsVolumeExportIDPut Update ACL by ID

Update export(ACL rule) by ID

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param volumeID ID of the JuiceFS volume
 @param volumeExportID ID of the JuiceFS volume export
 @return ApiVolumesVolumeIDExportsVolumeExportIDPutRequest
*/
func (a *VolumeACLAPIService) VolumesVolumeIDExportsVolumeExportIDPut(ctx context.Context, volumeID int32, volumeExportID int32) ApiVolumesVolumeIDExportsVolumeExportIDPutRequest {
	return ApiVolumesVolumeIDExportsVolumeExportIDPutRequest{
		ApiService: a,
		ctx: ctx,
		volumeID: volumeID,
		volumeExportID: volumeExportID,
	}
}

// Execute executes the request
//  @return GettableVolumeExport
func (a *VolumeACLAPIService) VolumesVolumeIDExportsVolumeExportIDPutExecute(r ApiVolumesVolumeIDExportsVolumeExportIDPutRequest) (*GettableVolumeExport, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GettableVolumeExport
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "VolumeACLAPIService.VolumesVolumeIDExportsVolumeExportIDPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/volumes/{volumeID}/exports/{volumeExportID}"
	localVarPath = strings.Replace(localVarPath, "{"+"volumeID"+"}", url.PathEscape(parameterValueToString(r.volumeID, "volumeID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"volumeExportID"+"}", url.PathEscape(parameterValueToString(r.volumeExportID, "volumeExportID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.volumeID < 1 {
		return localVarReturnValue, nil, reportError("volumeID must be greater than 1")
	}
	if r.volumeExportID < 1 {
		return localVarReturnValue, nil, reportError("volumeExportID must be greater than 1")
	}

	if r.token != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token", r.token, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.postableVolumeExport
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}