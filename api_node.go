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


// NodeAPIService NodeAPI service
type NodeAPIService service

type ApiNodesGetRequest struct {
	ctx context.Context
	ApiService *NodeAPIService
	token *string
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiNodesGetRequest) Token(token string) ApiNodesGetRequest {
	r.token = &token
	return r
}

func (r ApiNodesGetRequest) Execute() ([]Node, *http.Response, error) {
	return r.ApiService.NodesGetExecute(r)
}

/*
NodesGet Get node list

Return node list of the region identified by region token

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiNodesGetRequest
*/
func (a *NodeAPIService) NodesGet(ctx context.Context) ApiNodesGetRequest {
	return ApiNodesGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []Node
func (a *NodeAPIService) NodesGetExecute(r ApiNodesGetRequest) ([]Node, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Node
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAPIService.NodesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiNodesNodeIDGetRequest struct {
	ctx context.Context
	ApiService *NodeAPIService
	nodeID int32
	token *string
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiNodesNodeIDGetRequest) Token(token string) ApiNodesNodeIDGetRequest {
	r.token = &token
	return r
}

func (r ApiNodesNodeIDGetRequest) Execute() (*Node, *http.Response, error) {
	return r.ApiService.NodesNodeIDGetExecute(r)
}

/*
NodesNodeIDGet Get node by ID

Return details of the node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeID ID of the node
 @return ApiNodesNodeIDGetRequest
*/
func (a *NodeAPIService) NodesNodeIDGet(ctx context.Context, nodeID int32) ApiNodesNodeIDGetRequest {
	return ApiNodesNodeIDGetRequest{
		ApiService: a,
		ctx: ctx,
		nodeID: nodeID,
	}
}

// Execute executes the request
//  @return Node
func (a *NodeAPIService) NodesNodeIDGetExecute(r ApiNodesNodeIDGetRequest) (*Node, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Node
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAPIService.NodesNodeIDGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeID"+"}", url.PathEscape(parameterValueToString(r.nodeID, "nodeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nodeID < 1 {
		return localVarReturnValue, nil, reportError("nodeID must be greater than 1")
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

type ApiNodesNodeIDPatchRequest struct {
	ctx context.Context
	ApiService *NodeAPIService
	nodeID int32
	token *string
	node *Node
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiNodesNodeIDPatchRequest) Token(token string) ApiNodesNodeIDPatchRequest {
	r.token = &token
	return r
}

func (r ApiNodesNodeIDPatchRequest) Node(node Node) ApiNodesNodeIDPatchRequest {
	r.node = &node
	return r
}

func (r ApiNodesNodeIDPatchRequest) Execute() (*Node, *http.Response, error) {
	return r.ApiService.NodesNodeIDPatchExecute(r)
}

/*
NodesNodeIDPatch Method for NodesNodeIDPatch

Update node properties

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeID ID of the node
 @return ApiNodesNodeIDPatchRequest
*/
func (a *NodeAPIService) NodesNodeIDPatch(ctx context.Context, nodeID int32) ApiNodesNodeIDPatchRequest {
	return ApiNodesNodeIDPatchRequest{
		ApiService: a,
		ctx: ctx,
		nodeID: nodeID,
	}
}

// Execute executes the request
//  @return Node
func (a *NodeAPIService) NodesNodeIDPatchExecute(r ApiNodesNodeIDPatchRequest) (*Node, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Node
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAPIService.NodesNodeIDPatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeID"+"}", url.PathEscape(parameterValueToString(r.nodeID, "nodeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nodeID < 1 {
		return localVarReturnValue, nil, reportError("nodeID must be greater than 1")
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
	localVarPostBody = r.node
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

type ApiNodesNodeIDPutRequest struct {
	ctx context.Context
	ApiService *NodeAPIService
	nodeID int32
	token *string
	node *Node
}

// Token of the region, it&#39;s once used for authentication, now deprecated. please use access key/secret key instead.
// Deprecated
func (r ApiNodesNodeIDPutRequest) Token(token string) ApiNodesNodeIDPutRequest {
	r.token = &token
	return r
}

func (r ApiNodesNodeIDPutRequest) Node(node Node) ApiNodesNodeIDPutRequest {
	r.node = &node
	return r
}

func (r ApiNodesNodeIDPutRequest) Execute() (*Node, *http.Response, error) {
	return r.ApiService.NodesNodeIDPutExecute(r)
}

/*
NodesNodeIDPut Update node by ID

Update node

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nodeID ID of the node
 @return ApiNodesNodeIDPutRequest
*/
func (a *NodeAPIService) NodesNodeIDPut(ctx context.Context, nodeID int32) ApiNodesNodeIDPutRequest {
	return ApiNodesNodeIDPutRequest{
		ApiService: a,
		ctx: ctx,
		nodeID: nodeID,
	}
}

// Execute executes the request
//  @return Node
func (a *NodeAPIService) NodesNodeIDPutExecute(r ApiNodesNodeIDPutRequest) (*Node, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Node
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NodeAPIService.NodesNodeIDPut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/nodes/{nodeID}"
	localVarPath = strings.Replace(localVarPath, "{"+"nodeID"+"}", url.PathEscape(parameterValueToString(r.nodeID, "nodeID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nodeID < 1 {
		return localVarReturnValue, nil, reportError("nodeID must be greater than 1")
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
	localVarPostBody = r.node
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
