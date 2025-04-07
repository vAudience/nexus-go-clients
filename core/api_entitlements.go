/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.7
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// EntitlementsAPIService EntitlementsAPI service
type EntitlementsAPIService service

type ApiGetMyEntitlementsRequest struct {
	ctx context.Context
	ApiService *EntitlementsAPIService
	organizationId *string
}

// Organization ID
func (r ApiGetMyEntitlementsRequest) OrganizationId(organizationId string) ApiGetMyEntitlementsRequest {
	r.organizationId = &organizationId
	return r
}

func (r ApiGetMyEntitlementsRequest) Execute() (*map[string]EntitlementResponse, *http.Response, error) {
	return r.ApiService.GetMyEntitlementsExecute(r)
}

/*
GetMyEntitlements Get my entitlements

Get entitlements for the current user

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetMyEntitlementsRequest
*/
func (a *EntitlementsAPIService) GetMyEntitlements(ctx context.Context) ApiGetMyEntitlementsRequest {
	return ApiGetMyEntitlementsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]EntitlementResponse
func (a *EntitlementsAPIService) GetMyEntitlementsExecute(r ApiGetMyEntitlementsRequest) (*map[string]EntitlementResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]EntitlementResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "EntitlementsAPIService.GetMyEntitlements")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/entitlements/me"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.organizationId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "organizationId", r.organizationId, "", "")
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
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKey"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["x-api-key"] = key
			}
		}
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
