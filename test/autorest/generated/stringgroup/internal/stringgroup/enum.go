// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type EnumOperations struct{}

// GetNotExpandableCreateRequest creates the GetNotExpandable request.
func (EnumOperations) GetNotExpandableCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/enum/notExpandable")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNotExpandableHandleResponse handles the GetNotExpandable response.
func (EnumOperations) GetNotExpandableHandleResponse(resp *azcore.Response) (*EnumGetNotExpandableResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := EnumGetNotExpandableResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetReferencedCreateRequest creates the GetReferenced request.
func (EnumOperations) GetReferencedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/enum/Referenced")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetReferencedHandleResponse handles the GetReferenced response.
func (EnumOperations) GetReferencedHandleResponse(resp *azcore.Response) (*EnumGetReferencedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := EnumGetReferencedResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (EnumOperations) GetReferencedConstantCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/enum/ReferencedConstant")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (EnumOperations) GetReferencedConstantHandleResponse(resp *azcore.Response) (*EnumGetReferencedConstantResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	result := EnumGetReferencedConstantResponse{StatusCode: resp.StatusCode}
	return &result, resp.UnmarshalAsJSON(&result.RefColorConstant)
}

// PutNotExpandableCreateRequest creates the PutNotExpandable request.
func (EnumOperations) PutNotExpandableCreateRequest(u url.URL, stringBody Colors) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/enum/notExpandable")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(stringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutNotExpandableHandleResponse handles the PutNotExpandable response.
func (EnumOperations) PutNotExpandableHandleResponse(resp *azcore.Response) (*EnumPutNotExpandableResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &EnumPutNotExpandableResponse{StatusCode: resp.StatusCode}, nil
}

// PutReferencedCreateRequest creates the PutReferenced request.
func (EnumOperations) PutReferencedCreateRequest(u url.URL, enumStringBody Colors) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/enum/Referenced")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(enumStringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutReferencedHandleResponse handles the PutReferenced response.
func (EnumOperations) PutReferencedHandleResponse(resp *azcore.Response) (*EnumPutReferencedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &EnumPutReferencedResponse{StatusCode: resp.StatusCode}, nil
}

// PutReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (EnumOperations) PutReferencedConstantCreateRequest(u url.URL, enumStringBody RefColorConstant) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/string/enum/ReferencedConstant")
	req := azcore.NewRequest(http.MethodPut, u)
	err := req.MarshalAsJSON(enumStringBody)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutReferencedConstantHandleResponse handles the PutReferencedConstant response.
func (EnumOperations) PutReferencedConstantHandleResponse(resp *azcore.Response) (*EnumPutReferencedConstantResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &EnumPutReferencedConstantResponse{StatusCode: resp.StatusCode}, nil
}

