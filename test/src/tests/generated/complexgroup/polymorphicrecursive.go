package complexgroup

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

// PolymorphicrecursiveClient is the test Infrastructure for AutoRest
type PolymorphicrecursiveClient struct {
	BaseClient
}

// NewPolymorphicrecursiveClient creates an instance of the PolymorphicrecursiveClient client.
func NewPolymorphicrecursiveClient() PolymorphicrecursiveClient {
	return NewPolymorphicrecursiveClientWithBaseURI(DefaultBaseURI)
}

// NewPolymorphicrecursiveClientWithBaseURI creates an instance of the PolymorphicrecursiveClient client using a custom
// endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure
// stack).
func NewPolymorphicrecursiveClientWithBaseURI(baseURI string) PolymorphicrecursiveClient {
	return PolymorphicrecursiveClient{NewWithBaseURI(baseURI)}
}

// GetValid get complex types that are polymorphic and have recursive references
func (client PolymorphicrecursiveClient) GetValid(ctx context.Context) (result FishModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolymorphicrecursiveClient.GetValid")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetValidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphicrecursiveClient", "GetValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetValidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphicrecursiveClient", "GetValid", resp, "Failure sending request")
		return
	}

	result, err = client.GetValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphicrecursiveClient", "GetValid", resp, "Failure responding to request")
	}

	return
}

// GetValidPreparer prepares the GetValid request.
func (client PolymorphicrecursiveClient) GetValidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphicrecursive/valid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetValidSender sends the GetValid request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphicrecursiveClient) GetValidSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetValidResponder handles the response to the GetValid request. The method always
// closes the http.Response Body.
func (client PolymorphicrecursiveClient) GetValidResponder(resp *http.Response) (result FishModel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutValid put complex types that are polymorphic and have recursive references
// Parameters:
// complexBody - please put a salmon that looks like this:
// {
// "fishtype": "salmon",
// "species": "king",
// "length": 1,
// "age": 1,
// "location": "alaska",
// "iswild": true,
// "siblings": [
// {
// "fishtype": "shark",
// "species": "predator",
// "length": 20,
// "age": 6,
// "siblings": [
// {
// "fishtype": "salmon",
// "species": "coho",
// "length": 2,
// "age": 2,
// "location": "atlantic",
// "iswild": true,
// "siblings": [
// {
// "fishtype": "shark",
// "species": "predator",
// "length": 20,
// "age": 6
// },
// {
// "fishtype": "sawshark",
// "species": "dangerous",
// "length": 10,
// "age": 105
// }
// ]
// },
// {
// "fishtype": "sawshark",
// "species": "dangerous",
// "length": 10,
// "age": 105
// }
// ]
// },
// {
// "fishtype": "sawshark",
// "species": "dangerous",
// "length": 10,
// "age": 105
// }
// ]
// }
func (client PolymorphicrecursiveClient) PutValid(ctx context.Context, complexBody BasicFish) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/PolymorphicrecursiveClient.PutValid")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: complexBody,
			Constraints: []validation.Constraint{{Target: "complexBody.Length", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("complexgroup.PolymorphicrecursiveClient", "PutValid", err.Error())
	}

	req, err := client.PutValidPreparer(ctx, complexBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphicrecursiveClient", "PutValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphicrecursiveClient", "PutValid", resp, "Failure sending request")
		return
	}

	result, err = client.PutValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "complexgroup.PolymorphicrecursiveClient", "PutValid", resp, "Failure responding to request")
	}

	return
}

// PutValidPreparer prepares the PutValid request.
func (client PolymorphicrecursiveClient) PutValidPreparer(ctx context.Context, complexBody BasicFish) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/complex/polymorphicrecursive/valid"),
		autorest.WithJSON(complexBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutValidSender sends the PutValid request. The method will close the
// http.Response Body if it receives an error.
func (client PolymorphicrecursiveClient) PutValidSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PutValidResponder handles the response to the PutValid request. The method always
// closes the http.Response Body.
func (client PolymorphicrecursiveClient) PutValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
