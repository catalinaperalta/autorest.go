package durationgroup

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

// DurationClient is the test Infrastructure for AutoRest
type DurationClient struct {
	BaseClient
}

// NewDurationClient creates an instance of the DurationClient client.
func NewDurationClient() DurationClient {
	return NewDurationClientWithBaseURI(DefaultBaseURI)
}

// NewDurationClientWithBaseURI creates an instance of the DurationClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDurationClientWithBaseURI(baseURI string) DurationClient {
	return DurationClient{NewWithBaseURI(baseURI)}
}

// GetInvalid get an invalid duration value
func (client DurationClient) GetInvalid(ctx context.Context) (result TimeSpan, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DurationClient.GetInvalid")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetInvalidPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client DurationClient) GetInvalidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/duration/invalid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client DurationClient) GetInvalidSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client DurationClient) GetInvalidResponder(resp *http.Response) (result TimeSpan, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null duration value
func (client DurationClient) GetNull(ctx context.Context) (result TimeSpan, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DurationClient.GetNull")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNullPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client DurationClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/duration/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client DurationClient) GetNullSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client DurationClient) GetNullResponder(resp *http.Response) (result TimeSpan, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetPositiveDuration get a positive duration value
func (client DurationClient) GetPositiveDuration(ctx context.Context) (result TimeSpan, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DurationClient.GetPositiveDuration")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPositiveDurationPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetPositiveDuration", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetPositiveDurationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetPositiveDuration", resp, "Failure sending request")
		return
	}

	result, err = client.GetPositiveDurationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "GetPositiveDuration", resp, "Failure responding to request")
	}

	return
}

// GetPositiveDurationPreparer prepares the GetPositiveDuration request.
func (client DurationClient) GetPositiveDurationPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/duration/positiveduration"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetPositiveDurationSender sends the GetPositiveDuration request. The method will close the
// http.Response Body if it receives an error.
func (client DurationClient) GetPositiveDurationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetPositiveDurationResponder handles the response to the GetPositiveDuration request. The method always
// closes the http.Response Body.
func (client DurationClient) GetPositiveDurationResponder(resp *http.Response) (result TimeSpan, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutPositiveDuration put a positive duration value
func (client DurationClient) PutPositiveDuration(ctx context.Context, durationBody string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DurationClient.PutPositiveDuration")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutPositiveDurationPreparer(ctx, durationBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "PutPositiveDuration", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutPositiveDurationSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "PutPositiveDuration", resp, "Failure sending request")
		return
	}

	result, err = client.PutPositiveDurationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "durationgroup.DurationClient", "PutPositiveDuration", resp, "Failure responding to request")
	}

	return
}

// PutPositiveDurationPreparer prepares the PutPositiveDuration request.
func (client DurationClient) PutPositiveDurationPreparer(ctx context.Context, durationBody string) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/duration/positiveduration"),
		autorest.WithJSON(durationBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutPositiveDurationSender sends the PutPositiveDuration request. The method will close the
// http.Response Body if it receives an error.
func (client DurationClient) PutPositiveDurationSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PutPositiveDurationResponder handles the response to the PutPositiveDuration request. The method always
// closes the http.Response Body.
func (client DurationClient) PutPositiveDurationResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
