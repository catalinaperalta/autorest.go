package httpinfrastructuregroup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// HTTPFailureClient is the test Infrastructure for AutoRest
type HTTPFailureClient struct {
	BaseClient
}

// NewHTTPFailureClient creates an instance of the HTTPFailureClient client.
func NewHTTPFailureClient() HTTPFailureClient {
	return NewHTTPFailureClientWithBaseURI(DefaultBaseURI)
}

// NewHTTPFailureClientWithBaseURI creates an instance of the HTTPFailureClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewHTTPFailureClientWithBaseURI(baseURI string) HTTPFailureClient {
	return HTTPFailureClient{NewWithBaseURI(baseURI)}
}

// GetEmptyError get empty error form server
func (client HTTPFailureClient) GetEmptyError(ctx context.Context) (result Bool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPFailureClient.GetEmptyError")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEmptyErrorPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetEmptyError", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptyErrorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetEmptyError", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyErrorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetEmptyError", resp, "Failure responding to request")
	}

	return
}

// GetEmptyErrorPreparer prepares the GetEmptyError request.
func (client HTTPFailureClient) GetEmptyErrorPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/emptybody/error"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptyErrorSender sends the GetEmptyError request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetEmptyErrorSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetEmptyErrorResponder handles the response to the GetEmptyError request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetEmptyErrorResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNoModelEmpty get empty response from server
func (client HTTPFailureClient) GetNoModelEmpty(ctx context.Context) (result Bool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPFailureClient.GetNoModelEmpty")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNoModelEmptyPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNoModelEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetNoModelEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelEmpty", resp, "Failure responding to request")
	}

	return
}

// GetNoModelEmptyPreparer prepares the GetNoModelEmpty request.
func (client HTTPFailureClient) GetNoModelEmptyPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/nomodel/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNoModelEmptySender sends the GetNoModelEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetNoModelEmptySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNoModelEmptyResponder handles the response to the GetNoModelEmpty request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetNoModelEmptyResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNoModelError get empty error form server
func (client HTTPFailureClient) GetNoModelError(ctx context.Context) (result Bool, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/HTTPFailureClient.GetNoModelError")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNoModelErrorPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelError", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNoModelErrorSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelError", resp, "Failure sending request")
		return
	}

	result, err = client.GetNoModelErrorResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "httpinfrastructuregroup.HTTPFailureClient", "GetNoModelError", resp, "Failure responding to request")
	}

	return
}

// GetNoModelErrorPreparer prepares the GetNoModelError request.
func (client HTTPFailureClient) GetNoModelErrorPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/http/failure/nomodel/error"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNoModelErrorSender sends the GetNoModelError request. The method will close the
// http.Response Body if it receives an error.
func (client HTTPFailureClient) GetNoModelErrorSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetNoModelErrorResponder handles the response to the GetNoModelError request. The method always
// closes the http.Response Body.
func (client HTTPFailureClient) GetNoModelErrorResponder(resp *http.Response) (result Bool, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
