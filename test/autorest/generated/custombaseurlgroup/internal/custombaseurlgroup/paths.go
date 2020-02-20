// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package custombaseurlgroup

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
)

type PathsOperations struct{}

// GetEmptyCreateRequest creates the GetEmpty request.
func (PathsOperations) GetEmptyCreateRequest(u url.URL, accountName string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/customuri")
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetEmptyHandleResponse handles the GetEmpty response.
func (PathsOperations) GetEmptyHandleResponse(resp *azcore.Response) (*PathsGetEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &PathsGetEmptyResponse{StatusCode: resp.StatusCode}, nil
}

