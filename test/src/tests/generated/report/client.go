package report

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// ServiceVersion specifies the version of the operations used in this package.
	ServiceVersion = "1.0.0"
	// DefaultBaseURL is the default URL used for the service Report
	DefaultBaseURL = "http://localhost:3000"
)

// ManagementClient is the base client for Report.
type ManagementClient struct {
	url url.URL
	p   pipeline.Pipeline
}

// NewManagementClient creates an instance of the ManagementClient client.
func NewManagementClient(p pipeline.Pipeline) ManagementClient {
	u, err := url.Parse(DefaultBaseURL)
	if err != nil {
		panic(err)
	}
	return NewManagementClientWithURL(*u, p)
}

// NewManagementClientWithURL creates an instance of the ManagementClient client.
func NewManagementClientWithURL(url url.URL, p pipeline.Pipeline) ManagementClient {
	return ManagementClient{
		url: url,
		p:   p,
	}
}

// URL returns a copy of the URL for this client.
func (mc ManagementClient) URL() url.URL {
	return mc.url
}

// Pipeline returns the pipeline for this client.
func (mc ManagementClient) Pipeline() pipeline.Pipeline {
	return mc.p
}

// GetReport get test coverage report
//
// qualifier is if specified, qualifies the generated report further (e.g. '2.7' vs '3.5' in for Python). The only
// effect is, that generators that run all tests several times, can distinguish the generated reports.
func (client ManagementClient) GetReport(ctx context.Context, qualifier *string) (*GetReportResponse, error) {
	req, err := client.getReportPreparer(qualifier)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getReportResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*GetReportResponse), err
}

// getReportPreparer prepares the GetReport request.
func (client ManagementClient) getReportPreparer(qualifier *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/report"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	if qualifier != nil && len(*qualifier) > 0 {
		params.Set("qualifier", *qualifier)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getReportResponder handles the response to the GetReport request.
func (client ManagementClient) getReportResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &GetReportResponse{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, &result.Value)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
