// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package booleangroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// BoolGetFalseResponse contains the response from method Bool.GetFalse.
type BoolGetFalseResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
	// simple boolean
	Value *bool
}

// BoolGetInvalidResponse contains the response from method Bool.GetInvalid.
type BoolGetInvalidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
	Value *bool
}

// BoolGetNullResponse contains the response from method Bool.GetNull.
type BoolGetNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
	Value *bool
}

// BoolGetTrueResponse contains the response from method Bool.GetTrue.
type BoolGetTrueResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
	// simple boolean
	Value *bool
}

// BoolPutFalseResponse contains the response from method Bool.PutFalse.
type BoolPutFalseResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// BoolPutTrueResponse contains the response from method Bool.PutTrue.
type BoolPutTrueResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status *int32 `json:"status,omitempty"`
}

func newError(resp *azcore.Response) error {
	err := Error{}
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

func (e Error) Error() string {
	msg := ""
	if e.Message != nil {
		msg += fmt.Sprintf("Message: %v\n", *e.Message)
	}
	if e.Status != nil {
		msg += fmt.Sprintf("Status: %v\n", *e.Status)
	}
	if msg == "" {
		msg = "missing error info"
	}
	return msg
}

