package bytegroup

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

// ByteClient is the test Infrastructure for AutoRest Swagger BAT
type ByteClient struct {
	BaseClient
}

// NewByteClient creates an instance of the ByteClient client.
func NewByteClient() ByteClient {
	return NewByteClientWithBaseURI(DefaultBaseURI)
}

// NewByteClientWithBaseURI creates an instance of the ByteClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewByteClientWithBaseURI(baseURI string) ByteClient {
	return ByteClient{NewWithBaseURI(baseURI)}
}

// GetEmpty get empty byte value ''
func (client ByteClient) GetEmpty(ctx context.Context) (result ByteArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ByteClient.GetEmpty")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetEmptyPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetEmpty", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetEmptySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetEmpty", resp, "Failure sending request")
		return
	}

	result, err = client.GetEmptyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetEmpty", resp, "Failure responding to request")
	}

	return
}

// GetEmptyPreparer prepares the GetEmpty request.
func (client ByteClient) GetEmptyPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/byte/empty"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetEmptySender sends the GetEmpty request. The method will close the
// http.Response Body if it receives an error.
func (client ByteClient) GetEmptySender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetEmptyResponder handles the response to the GetEmpty request. The method always
// closes the http.Response Body.
func (client ByteClient) GetEmptyResponder(resp *http.Response) (result ByteArray, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetInvalid get invalid byte value ':::SWAGGER::::'
func (client ByteClient) GetInvalid(ctx context.Context) (result ByteArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ByteClient.GetInvalid")
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
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetInvalid", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetInvalidSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetInvalid", resp, "Failure sending request")
		return
	}

	result, err = client.GetInvalidResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetInvalid", resp, "Failure responding to request")
	}

	return
}

// GetInvalidPreparer prepares the GetInvalid request.
func (client ByteClient) GetInvalidPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/byte/invalid"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetInvalidSender sends the GetInvalid request. The method will close the
// http.Response Body if it receives an error.
func (client ByteClient) GetInvalidSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetInvalidResponder handles the response to the GetInvalid request. The method always
// closes the http.Response Body.
func (client ByteClient) GetInvalidResponder(resp *http.Response) (result ByteArray, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNonASCII get non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client ByteClient) GetNonASCII(ctx context.Context) (result ByteArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ByteClient.GetNonASCII")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetNonASCIIPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetNonASCII", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNonASCIISender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetNonASCII", resp, "Failure sending request")
		return
	}

	result, err = client.GetNonASCIIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetNonASCII", resp, "Failure responding to request")
	}

	return
}

// GetNonASCIIPreparer prepares the GetNonASCII request.
func (client ByteClient) GetNonASCIIPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/byte/nonAscii"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNonASCIISender sends the GetNonASCII request. The method will close the
// http.Response Body if it receives an error.
func (client ByteClient) GetNonASCIISender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetNonASCIIResponder handles the response to the GetNonASCII request. The method always
// closes the http.Response Body.
func (client ByteClient) GetNonASCIIResponder(resp *http.Response) (result ByteArray, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetNull get null byte value
func (client ByteClient) GetNull(ctx context.Context) (result ByteArray, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ByteClient.GetNull")
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
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetNull", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetNullSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetNull", resp, "Failure sending request")
		return
	}

	result, err = client.GetNullResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "GetNull", resp, "Failure responding to request")
	}

	return
}

// GetNullPreparer prepares the GetNull request.
func (client ByteClient) GetNullPreparer(ctx context.Context) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/byte/null"))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetNullSender sends the GetNull request. The method will close the
// http.Response Body if it receives an error.
func (client ByteClient) GetNullSender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// GetNullResponder handles the response to the GetNull request. The method always
// closes the http.Response Body.
func (client ByteClient) GetNullResponder(resp *http.Response) (result ByteArray, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutNonASCII put non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
// Parameters:
// byteBody - base64-encoded non-ascii byte string hex(FF FE FD FC FB FA F9 F8 F7 F6)
func (client ByteClient) PutNonASCII(ctx context.Context, byteBody []byte) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ByteClient.PutNonASCII")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: byteBody,
			Constraints: []validation.Constraint{{Target: "byteBody", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("bytegroup.ByteClient", "PutNonASCII", err.Error())
	}

	req, err := client.PutNonASCIIPreparer(ctx, byteBody)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "PutNonASCII", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutNonASCIISender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "PutNonASCII", resp, "Failure sending request")
		return
	}

	result, err = client.PutNonASCIIResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bytegroup.ByteClient", "PutNonASCII", resp, "Failure responding to request")
	}

	return
}

// PutNonASCIIPreparer prepares the PutNonASCII request.
func (client ByteClient) PutNonASCIIPreparer(ctx context.Context, byteBody []byte) (*http.Request, error) {
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPath("/byte/nonAscii"),
		autorest.WithJSON(byteBody))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutNonASCIISender sends the PutNonASCII request. The method will close the
// http.Response Body if it receives an error.
func (client ByteClient) PutNonASCIISender(req *http.Request) (*http.Response, error) {
	sd := autorest.GetSendDecorators(req.Context(), autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
	return autorest.SendWithSender(client, req, sd...)
}

// PutNonASCIIResponder handles the response to the PutNonASCII request. The method always
// closes the http.Response Body.
func (client ByteClient) PutNonASCIIResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}
