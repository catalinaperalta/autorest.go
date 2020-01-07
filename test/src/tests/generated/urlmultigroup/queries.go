package urlmultigroup

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

// QueriesClient is the test Infrastructure for AutoRest
type QueriesClient struct {
	BaseClient
}

// NewQueriesClient creates an instance of the QueriesClient client.
func NewQueriesClient() QueriesClient {
	return NewQueriesClientWithBaseURI(DefaultBaseURI)
}

// NewQueriesClientWithBaseURI creates an instance of the QueriesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewQueriesClientWithBaseURI(baseURI string) QueriesClient {
	return QueriesClient{NewWithBaseURI(baseURI)}
}

// ArrayStringMultiEmpty get an empty array [] of string using the multi-array format
// Parameters:
// arrayQuery - an empty array [] of string using the multi-array format
func (client QueriesClient) ArrayStringMultiEmpty(ctx context.Context, arrayQuery []string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.ArrayStringMultiEmpty")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ArrayStringMultiEmptyPreparer(ctx, arrayQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.ArrayStringMultiEmptySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.ArrayStringMultiEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiEmpty", resp, "Failure responding to request")
	}

	return
}

// ArrayStringMultiEmptyPreparer prepares the ArrayStringMultiEmpty request.
func (client QueriesClient) ArrayStringMultiEmptyPreparer(ctx context.Context, arrayQuery []string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if arrayQuery != nil && len(arrayQuery) > 0 {
		queryParameters["arrayQuery"] = arrayQuery
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/queries/array/multi/string/empty"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ArrayStringMultiEmptySender sends the ArrayStringMultiEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) ArrayStringMultiEmptySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ArrayStringMultiEmptyResponder handles the response to the ArrayStringMultiEmpty request. The method always
// closes the http.Response Body.
func (client QueriesClient) ArrayStringMultiEmptyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ArrayStringMultiNull get a null array of string using the multi-array format
// Parameters:
// arrayQuery - a null array of string using the multi-array format
func (client QueriesClient) ArrayStringMultiNull(ctx context.Context, arrayQuery []string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.ArrayStringMultiNull")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ArrayStringMultiNullPreparer(ctx, arrayQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.ArrayStringMultiNullSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiNull", resp, "Failure sending request")
		return
	}

	result, err = client.ArrayStringMultiNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiNull", resp, "Failure responding to request")
	}

	return
}

// ArrayStringMultiNullPreparer prepares the ArrayStringMultiNull request.
func (client QueriesClient) ArrayStringMultiNullPreparer(ctx context.Context, arrayQuery []string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if arrayQuery != nil && len(arrayQuery) > 0 {
		queryParameters["arrayQuery"] = arrayQuery
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/queries/array/multi/string/null"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ArrayStringMultiNullSender sends the ArrayStringMultiNull request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) ArrayStringMultiNullSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ArrayStringMultiNullResponder handles the response to the ArrayStringMultiNull request. The method always
// closes the http.Response Body.
func (client QueriesClient) ArrayStringMultiNullResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// ArrayStringMultiValid get an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the
// mult-array format
// Parameters:
// arrayQuery - an array of string ['ArrayQuery1', 'begin!*'();:@ &=+$,/?#[]end' , null, ''] using the
// mult-array format
func (client QueriesClient) ArrayStringMultiValid(ctx context.Context, arrayQuery []string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/QueriesClient.ArrayStringMultiValid")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ArrayStringMultiValidPreparer(ctx, arrayQuery)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiValid", nil, "Failure preparing request")
		return
	}

	resp, err := client.ArrayStringMultiValidSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiValid", resp, "Failure sending request")
		return
	}

	result, err = client.ArrayStringMultiValidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "urlmultigroup.QueriesClient", "ArrayStringMultiValid", resp, "Failure responding to request")
	}

	return
}

// ArrayStringMultiValidPreparer prepares the ArrayStringMultiValid request.
func (client QueriesClient) ArrayStringMultiValidPreparer(ctx context.Context, arrayQuery []string) (*http.Request, error) {
	queryParameters := map[string]interface{}{}
	if arrayQuery != nil && len(arrayQuery) > 0 {
		queryParameters["arrayQuery"] = arrayQuery
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/queries/array/multi/string/valid"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ArrayStringMultiValidSender sends the ArrayStringMultiValid request. The method will close the
// http.Response Body if it receives an error.
func (client QueriesClient) ArrayStringMultiValidSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// ArrayStringMultiValidResponder handles the response to the ArrayStringMultiValid request. The method always
// closes the http.Response Body.
func (client QueriesClient) ArrayStringMultiValidResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
