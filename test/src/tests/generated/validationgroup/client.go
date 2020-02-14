// Package validationgroup implements the Azure ARM Validationgroup service API version 1.0.0.
//
// Test Infrastructure for AutoRest. No server backend exists for these tests.
package validationgroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

const (
	// DefaultBaseURI is the default URI used for the service Validationgroup
	DefaultBaseURI = "http://localhost:3000"
)

// BaseClient is the base client for Validationgroup.
type BaseClient struct {
	autorest.Client
	BaseURI        string
	SubscriptionID string
}

// New creates an instance of the BaseClient client.
func New(subscriptionID string) BaseClient {
	return NewWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewWithBaseURI creates an instance of the BaseClient client using a custom endpoint.  Use this when interacting with
// an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return BaseClient{
		Client:         autorest.NewClientWithUserAgent(UserAgent()),
		BaseURI:        baseURI,
		SubscriptionID: subscriptionID,
	}
}

// GetWithConstantInPath sends the get with constant in path request.
func (client BaseClient) GetWithConstantInPath(ctx context.Context) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.GetWithConstantInPath")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetWithConstantInPathPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "GetWithConstantInPath", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetWithConstantInPathSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "GetWithConstantInPath", resp, "Failure sending request")
		return
	}

	result, err = client.GetWithConstantInPathResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "GetWithConstantInPath", resp, "Failure responding to request")
	}

	return
}

// GetWithConstantInPathPreparer prepares the GetWithConstantInPath request.
func (client BaseClient) GetWithConstantInPathPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"constantParam": autorest.Encode("path", "constant"),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/validation/constantsInPath/{constantParam}/value", pathParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetWithConstantInPathSender sends the GetWithConstantInPath request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) GetWithConstantInPathSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetWithConstantInPathResponder handles the response to the GetWithConstantInPath request. The method always
// closes the http.Response Body.
func (client BaseClient) GetWithConstantInPathResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// PostWithConstantInBody sends the post with constant in body request.
func (client BaseClient) PostWithConstantInBody(ctx context.Context, body *Product) (result Product, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.PostWithConstantInBody")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.DisplayNames", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.DisplayNames", Name: validation.MaxItems, Rule: 6, Chain: nil},
						{Target: "body.DisplayNames", Name: validation.MinItems, Rule: 0, Chain: nil},
						{Target: "body.DisplayNames", Name: validation.UniqueItems, Rule: true, Chain: nil},
					}},
					{Target: "body.Capacity", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Capacity", Name: validation.ExclusiveMaximum, Rule: 100, Chain: nil},
							{Target: "body.Capacity", Name: validation.ExclusiveMinimum, Rule: 0, Chain: nil},
						}},
					{Target: "body.Image", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Image", Name: validation.Pattern, Rule: `http://\w+`, Chain: nil}}},
					{Target: "body.Child", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "body.Child.ConstProperty", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "body.ConstChild", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "body.ConstChild.ConstProperty", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.ConstChild.ConstProperty2", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "body.ConstInt", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "body.ConstString", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("validationgroup.BaseClient", "PostWithConstantInBody", err.Error())
	}

	req, err := client.PostWithConstantInBodyPreparer(ctx, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "PostWithConstantInBody", nil, "Failure preparing request")
		return
	}

	resp, err := client.PostWithConstantInBodySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "PostWithConstantInBody", resp, "Failure sending request")
		return
	}

	result, err = client.PostWithConstantInBodyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "PostWithConstantInBody", resp, "Failure responding to request")
	}

	return
}

// PostWithConstantInBodyPreparer prepares the PostWithConstantInBody request.
func (client BaseClient) PostWithConstantInBodyPreparer(ctx context.Context, body *Product) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"constantParam": autorest.Encode("path", "constant"),
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/validation/constantsInPath/{constantParam}/value", pathParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostWithConstantInBodySender sends the PostWithConstantInBody request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) PostWithConstantInBodySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// PostWithConstantInBodyResponder handles the response to the PostWithConstantInBody request. The method always
// closes the http.Response Body.
func (client BaseClient) PostWithConstantInBodyResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidationOfBody validates body parameters on the method. See swagger for details.
// Parameters:
// resourceGroupName - required string between 3 and 10 chars with pattern [a-zA-Z0-9]+.
// ID - required int multiple of 10 from 100 to 1000.
func (client BaseClient) ValidationOfBody(ctx context.Context, resourceGroupName string, ID int32, body *Product) (result Product, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ValidationOfBody")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 10, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: ID,
			Constraints: []validation.Constraint{{Target: "ID", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
				{Target: "ID", Name: validation.InclusiveMinimum, Rule: int64(100), Chain: nil},
				{Target: "ID", Name: validation.MultipleOf, Rule: 10, Chain: nil}}},
		{TargetValue: body,
			Constraints: []validation.Constraint{{Target: "body", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "body.DisplayNames", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "body.DisplayNames", Name: validation.MaxItems, Rule: 6, Chain: nil},
						{Target: "body.DisplayNames", Name: validation.MinItems, Rule: 0, Chain: nil},
						{Target: "body.DisplayNames", Name: validation.UniqueItems, Rule: true, Chain: nil},
					}},
					{Target: "body.Capacity", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Capacity", Name: validation.ExclusiveMaximum, Rule: 100, Chain: nil},
							{Target: "body.Capacity", Name: validation.ExclusiveMinimum, Rule: 0, Chain: nil},
						}},
					{Target: "body.Image", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "body.Image", Name: validation.Pattern, Rule: `http://\w+`, Chain: nil}}},
					{Target: "body.Child", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "body.Child.ConstProperty", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "body.ConstChild", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "body.ConstChild.ConstProperty", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "body.ConstChild.ConstProperty2", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "body.ConstInt", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "body.ConstString", Name: validation.Null, Rule: true, Chain: nil},
				}}}}}); err != nil {
		return result, validation.NewError("validationgroup.BaseClient", "ValidationOfBody", err.Error())
	}

	req, err := client.ValidationOfBodyPreparer(ctx, resourceGroupName, ID, body)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "ValidationOfBody", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidationOfBodySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "ValidationOfBody", resp, "Failure sending request")
		return
	}

	result, err = client.ValidationOfBodyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "ValidationOfBody", resp, "Failure responding to request")
	}

	return
}

// ValidationOfBodyPreparer prepares the ValidationOfBody request.
func (client BaseClient) ValidationOfBodyPreparer(ctx context.Context, resourceGroupName string, ID int32, body *Product) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"apiVersion": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/fakepath/{subscriptionId}/{resourceGroupName}/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if body != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithJSON(body))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidationOfBodySender sends the ValidationOfBody request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ValidationOfBodySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ValidationOfBodyResponder handles the response to the ValidationOfBody request. The method always
// closes the http.Response Body.
func (client BaseClient) ValidationOfBodyResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ValidationOfMethodParameters validates input parameters on the method. See swagger for details.
// Parameters:
// resourceGroupName - required string between 3 and 10 chars with pattern [a-zA-Z0-9]+.
// ID - required int multiple of 10 from 100 to 1000.
func (client BaseClient) ValidationOfMethodParameters(ctx context.Context, resourceGroupName string, ID int32) (result Product, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/BaseClient.ValidationOfMethodParameters")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 10, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9]+`, Chain: nil}}},
		{TargetValue: ID,
			Constraints: []validation.Constraint{{Target: "ID", Name: validation.InclusiveMaximum, Rule: int64(1000), Chain: nil},
				{Target: "ID", Name: validation.InclusiveMinimum, Rule: int64(100), Chain: nil},
				{Target: "ID", Name: validation.MultipleOf, Rule: 10, Chain: nil}}}}); err != nil {
		return result, validation.NewError("validationgroup.BaseClient", "ValidationOfMethodParameters", err.Error())
	}

	req, err := client.ValidationOfMethodParametersPreparer(ctx, resourceGroupName, ID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "ValidationOfMethodParameters", nil, "Failure preparing request")
		return
	}

	resp, err := client.ValidationOfMethodParametersSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "ValidationOfMethodParameters", resp, "Failure sending request")
		return
	}

	result, err = client.ValidationOfMethodParametersResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "validationgroup.BaseClient", "ValidationOfMethodParameters", resp, "Failure responding to request")
	}

	return
}

// ValidationOfMethodParametersPreparer prepares the ValidationOfMethodParameters request.
func (client BaseClient) ValidationOfMethodParametersPreparer(ctx context.Context, resourceGroupName string, ID int32) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"id":                autorest.Encode("path", ID),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "1.0.0"
	queryParameters := map[string]interface{}{
		"apiVersion": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/fakepath/{subscriptionId}/{resourceGroupName}/{id}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ValidationOfMethodParametersSender sends the ValidationOfMethodParameters request. The method will close the
// http.Response Body if it receives an error.
func (client BaseClient) ValidationOfMethodParametersSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ValidationOfMethodParametersResponder handles the response to the ValidationOfMethodParameters request. The method always
// closes the http.Response Body.
func (client BaseClient) ValidationOfMethodParametersResponder(resp *http.Response) (result Product, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
