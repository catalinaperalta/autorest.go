// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type BoolOperations struct{}

// GetFalseCreateRequest creates the GetFalse request.
func (BoolOperations) GetFalseCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/bool/false")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetFalseHandleResponse handles the GetFalse response.
func (BoolOperations) GetFalseHandleResponse(resp *azcore.Response) (*BoolGetFalseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetFalseResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetInvalidCreateRequest creates the GetInvalid request.
func (BoolOperations) GetInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/bool/invalid")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetInvalidHandleResponse handles the GetInvalid response.
func (BoolOperations) GetInvalidHandleResponse(resp *azcore.Response) (*BoolGetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetInvalidResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNullCreateRequest creates the GetNull request.
func (BoolOperations) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/bool/null")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNullHandleResponse handles the GetNull response.
func (BoolOperations) GetNullHandleResponse(resp *azcore.Response) (*BoolGetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetNullResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetTrueCreateRequest creates the GetTrue request.
func (BoolOperations) GetTrueCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/bool/true")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetTrueHandleResponse handles the GetTrue response.
func (BoolOperations) GetTrueHandleResponse(resp *azcore.Response) (*BoolGetTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := BoolGetTrueResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// PutFalseCreateRequest creates the PutFalse request.
func (BoolOperations) PutFalseCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/bool/false")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(false)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutFalseHandleResponse handles the PutFalse response.
func (BoolOperations) PutFalseHandleResponse(resp *azcore.Response) (*BoolPutFalseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &BoolPutFalseResponse{RawResponse: resp.Response}, nil
}

// PutTrueCreateRequest creates the PutTrue request.
func (BoolOperations) PutTrueCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/bool/true")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(true)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutTrueHandleResponse handles the PutTrue response.
func (BoolOperations) PutTrueHandleResponse(resp *azcore.Response) (*BoolPutTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &BoolPutTrueResponse{RawResponse: resp.Response}, nil
}