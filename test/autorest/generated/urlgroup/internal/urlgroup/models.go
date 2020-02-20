// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

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

// PathItemsGetAllWithValuesResponse contains the response from method PathItems.GetAllWithValues.
type PathItemsGetAllWithValuesResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathItemsGetGlobalAndLocalQueryNullResponse contains the response from method PathItems.GetGlobalAndLocalQueryNull.
type PathItemsGetGlobalAndLocalQueryNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathItemsGetGlobalQueryNullResponse contains the response from method PathItems.GetGlobalQueryNull.
type PathItemsGetGlobalQueryNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathItemsGetLocalPathItemQueryNullResponse contains the response from method PathItems.GetLocalPathItemQueryNull.
type PathItemsGetLocalPathItemQueryNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsArrayCsvInPathResponse contains the response from method Paths.ArrayCsvInPath.
type PathsArrayCsvInPathResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsBase64URLResponse contains the response from method Paths.Base64URL.
type PathsBase64URLResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsByteEmptyResponse contains the response from method Paths.ByteEmpty.
type PathsByteEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsByteMultiByteResponse contains the response from method Paths.ByteMultiByte.
type PathsByteMultiByteResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsByteNullResponse contains the response from method Paths.ByteNull.
type PathsByteNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsDateNullResponse contains the response from method Paths.DateNull.
type PathsDateNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsDateTimeNullResponse contains the response from method Paths.DateTimeNull.
type PathsDateTimeNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsDateTimeValidResponse contains the response from method Paths.DateTimeValid.
type PathsDateTimeValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsDateValidResponse contains the response from method Paths.DateValid.
type PathsDateValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsDoubleDecimalNegativeResponse contains the response from method Paths.DoubleDecimalNegative.
type PathsDoubleDecimalNegativeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsDoubleDecimalPositiveResponse contains the response from method Paths.DoubleDecimalPositive.
type PathsDoubleDecimalPositiveResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsEnumNullResponse contains the response from method Paths.EnumNull.
type PathsEnumNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsEnumValidResponse contains the response from method Paths.EnumValid.
type PathsEnumValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsFloatScientificNegativeResponse contains the response from method Paths.FloatScientificNegative.
type PathsFloatScientificNegativeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsFloatScientificPositiveResponse contains the response from method Paths.FloatScientificPositive.
type PathsFloatScientificPositiveResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsGetBooleanFalseResponse contains the response from method Paths.GetBooleanFalse.
type PathsGetBooleanFalseResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsGetBooleanTrueResponse contains the response from method Paths.GetBooleanTrue.
type PathsGetBooleanTrueResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsGetIntNegativeOneMillionResponse contains the response from method Paths.GetIntNegativeOneMillion.
type PathsGetIntNegativeOneMillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsGetIntOneMillionResponse contains the response from method Paths.GetIntOneMillion.
type PathsGetIntOneMillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsGetNegativeTenBillionResponse contains the response from method Paths.GetNegativeTenBillion.
type PathsGetNegativeTenBillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsGetTenBillionResponse contains the response from method Paths.GetTenBillion.
type PathsGetTenBillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsStringEmptyResponse contains the response from method Paths.StringEmpty.
type PathsStringEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsStringNullResponse contains the response from method Paths.StringNull.
type PathsStringNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsStringURLEncodedResponse contains the response from method Paths.StringURLEncoded.
type PathsStringURLEncodedResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsStringURLNonEncodedResponse contains the response from method Paths.StringURLNonEncoded.
type PathsStringURLNonEncodedResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsStringUnicodeResponse contains the response from method Paths.StringUnicode.
type PathsStringUnicodeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// PathsUnixTimeURLResponse contains the response from method Paths.UnixTimeURL.
type PathsUnixTimeURLResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesArrayStringCsvEmptyResponse contains the response from method Queries.ArrayStringCsvEmpty.
type QueriesArrayStringCsvEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesArrayStringCsvNullResponse contains the response from method Queries.ArrayStringCsvNull.
type QueriesArrayStringCsvNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesArrayStringCsvValidResponse contains the response from method Queries.ArrayStringCsvValid.
type QueriesArrayStringCsvValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesArrayStringPipesValidResponse contains the response from method Queries.ArrayStringPipesValid.
type QueriesArrayStringPipesValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesArrayStringSsvValidResponse contains the response from method Queries.ArrayStringSsvValid.
type QueriesArrayStringSsvValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesArrayStringTsvValidResponse contains the response from method Queries.ArrayStringTsvValid.
type QueriesArrayStringTsvValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesByteEmptyResponse contains the response from method Queries.ByteEmpty.
type QueriesByteEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesByteMultiByteResponse contains the response from method Queries.ByteMultiByte.
type QueriesByteMultiByteResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesByteNullResponse contains the response from method Queries.ByteNull.
type QueriesByteNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDateNullResponse contains the response from method Queries.DateNull.
type QueriesDateNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDateTimeNullResponse contains the response from method Queries.DateTimeNull.
type QueriesDateTimeNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDateTimeValidResponse contains the response from method Queries.DateTimeValid.
type QueriesDateTimeValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDateValidResponse contains the response from method Queries.DateValid.
type QueriesDateValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDoubleDecimalNegativeResponse contains the response from method Queries.DoubleDecimalNegative.
type QueriesDoubleDecimalNegativeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDoubleDecimalPositiveResponse contains the response from method Queries.DoubleDecimalPositive.
type QueriesDoubleDecimalPositiveResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesDoubleNullResponse contains the response from method Queries.DoubleNull.
type QueriesDoubleNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesEnumNullResponse contains the response from method Queries.EnumNull.
type QueriesEnumNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesEnumValidResponse contains the response from method Queries.EnumValid.
type QueriesEnumValidResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesFloatNullResponse contains the response from method Queries.FloatNull.
type QueriesFloatNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesFloatScientificNegativeResponse contains the response from method Queries.FloatScientificNegative.
type QueriesFloatScientificNegativeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesFloatScientificPositiveResponse contains the response from method Queries.FloatScientificPositive.
type QueriesFloatScientificPositiveResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetBooleanFalseResponse contains the response from method Queries.GetBooleanFalse.
type QueriesGetBooleanFalseResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetBooleanNullResponse contains the response from method Queries.GetBooleanNull.
type QueriesGetBooleanNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetBooleanTrueResponse contains the response from method Queries.GetBooleanTrue.
type QueriesGetBooleanTrueResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetIntNegativeOneMillionResponse contains the response from method Queries.GetIntNegativeOneMillion.
type QueriesGetIntNegativeOneMillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetIntNullResponse contains the response from method Queries.GetIntNull.
type QueriesGetIntNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetIntOneMillionResponse contains the response from method Queries.GetIntOneMillion.
type QueriesGetIntOneMillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetLongNullResponse contains the response from method Queries.GetLongNull.
type QueriesGetLongNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetNegativeTenBillionResponse contains the response from method Queries.GetNegativeTenBillion.
type QueriesGetNegativeTenBillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesGetTenBillionResponse contains the response from method Queries.GetTenBillion.
type QueriesGetTenBillionResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesStringEmptyResponse contains the response from method Queries.StringEmpty.
type QueriesStringEmptyResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesStringNullResponse contains the response from method Queries.StringNull.
type QueriesStringNullResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesStringURLEncodedResponse contains the response from method Queries.StringURLEncoded.
type QueriesStringURLEncodedResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

// QueriesStringUnicodeResponse contains the response from method Queries.StringUnicode.
type QueriesStringUnicodeResponse struct {
	// StatusCode contains the HTTP status code.
	StatusCode int
}

