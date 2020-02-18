// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package urlgroup

import (
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

type QueriesOperations struct{}

// ArrayStringCsvEmptyCreateRequest creates the ArrayStringCsvEmpty request.
func (QueriesOperations) ArrayStringCsvEmptyCreateRequest(u url.URL, arrayQuery []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/array/csv/string/empty")
	query := u.Query()
	query.Set("arrayQuery", strings.Join(arrayQuery, ","))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayStringCsvEmptyHandleResponse handles the ArrayStringCsvEmpty response.
func (QueriesOperations) ArrayStringCsvEmptyHandleResponse(resp *azcore.Response) (*QueriesArrayStringCsvEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesArrayStringCsvEmptyResponse{StatusCode: resp.StatusCode}, nil
}

// ArrayStringCsvNullCreateRequest creates the ArrayStringCsvNull request.
func (QueriesOperations) ArrayStringCsvNullCreateRequest(u url.URL, arrayQuery []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/array/csv/string/null")
	query := u.Query()
	query.Set("arrayQuery", strings.Join(arrayQuery, ","))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayStringCsvNullHandleResponse handles the ArrayStringCsvNull response.
func (QueriesOperations) ArrayStringCsvNullHandleResponse(resp *azcore.Response) (*QueriesArrayStringCsvNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesArrayStringCsvNullResponse{StatusCode: resp.StatusCode}, nil
}

// ArrayStringCsvValidCreateRequest creates the ArrayStringCsvValid request.
func (QueriesOperations) ArrayStringCsvValidCreateRequest(u url.URL, arrayQuery []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/array/csv/string/valid")
	query := u.Query()
	query.Set("arrayQuery", strings.Join(arrayQuery, ","))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayStringCsvValidHandleResponse handles the ArrayStringCsvValid response.
func (QueriesOperations) ArrayStringCsvValidHandleResponse(resp *azcore.Response) (*QueriesArrayStringCsvValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesArrayStringCsvValidResponse{StatusCode: resp.StatusCode}, nil
}

// ArrayStringPipesValidCreateRequest creates the ArrayStringPipesValid request.
func (QueriesOperations) ArrayStringPipesValidCreateRequest(u url.URL, arrayQuery []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/array/pipes/string/valid")
	query := u.Query()
	query.Set("arrayQuery", strings.Join(arrayQuery, "|"))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayStringPipesValidHandleResponse handles the ArrayStringPipesValid response.
func (QueriesOperations) ArrayStringPipesValidHandleResponse(resp *azcore.Response) (*QueriesArrayStringPipesValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesArrayStringPipesValidResponse{StatusCode: resp.StatusCode}, nil
}

// ArrayStringSsvValidCreateRequest creates the ArrayStringSsvValid request.
func (QueriesOperations) ArrayStringSsvValidCreateRequest(u url.URL, arrayQuery []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/array/ssv/string/valid")
	query := u.Query()
	query.Set("arrayQuery", strings.Join(arrayQuery, " "))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayStringSsvValidHandleResponse handles the ArrayStringSsvValid response.
func (QueriesOperations) ArrayStringSsvValidHandleResponse(resp *azcore.Response) (*QueriesArrayStringSsvValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesArrayStringSsvValidResponse{StatusCode: resp.StatusCode}, nil
}

// ArrayStringTsvValidCreateRequest creates the ArrayStringTsvValid request.
func (QueriesOperations) ArrayStringTsvValidCreateRequest(u url.URL, arrayQuery []string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/array/tsv/string/valid")
	query := u.Query()
	query.Set("arrayQuery", strings.Join(arrayQuery, "\t"))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ArrayStringTsvValidHandleResponse handles the ArrayStringTsvValid response.
func (QueriesOperations) ArrayStringTsvValidHandleResponse(resp *azcore.Response) (*QueriesArrayStringTsvValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesArrayStringTsvValidResponse{StatusCode: resp.StatusCode}, nil
}

// ByteEmptyCreateRequest creates the ByteEmpty request.
func (QueriesOperations) ByteEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/byte/empty")
	query := u.Query()
	query.Set("byteQuery", "")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ByteEmptyHandleResponse handles the ByteEmpty response.
func (QueriesOperations) ByteEmptyHandleResponse(resp *azcore.Response) (*QueriesByteEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesByteEmptyResponse{StatusCode: resp.StatusCode}, nil
}

// ByteMultiByteCreateRequest creates the ByteMultiByte request.
func (QueriesOperations) ByteMultiByteCreateRequest(u url.URL, byteQuery []byte) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/byte/multibyte")
	query := u.Query()
	query.Set("byteQuery", string(byteQuery))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ByteMultiByteHandleResponse handles the ByteMultiByte response.
func (QueriesOperations) ByteMultiByteHandleResponse(resp *azcore.Response) (*QueriesByteMultiByteResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesByteMultiByteResponse{StatusCode: resp.StatusCode}, nil
}

// ByteNullCreateRequest creates the ByteNull request.
func (QueriesOperations) ByteNullCreateRequest(u url.URL, byteQuery []byte) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/byte/null")
	query := u.Query()
	query.Set("byteQuery", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(byteQuery), "[]")), ","))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// ByteNullHandleResponse handles the ByteNull response.
func (QueriesOperations) ByteNullHandleResponse(resp *azcore.Response) (*QueriesByteNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesByteNullResponse{StatusCode: resp.StatusCode}, nil
}

// DateNullCreateRequest creates the DateNull request.
func (QueriesOperations) DateNullCreateRequest(u url.URL, dateQuery time.Time) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/date/null")
	query := u.Query()
	query.Set("dateQuery", dateQuery.String())
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateNullHandleResponse handles the DateNull response.
func (QueriesOperations) DateNullHandleResponse(resp *azcore.Response) (*QueriesDateNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDateNullResponse{StatusCode: resp.StatusCode}, nil
}

// DateTimeNullCreateRequest creates the DateTimeNull request.
func (QueriesOperations) DateTimeNullCreateRequest(u url.URL, dateTimeQuery time.Time) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/datetime/null")
	query := u.Query()
	query.Set("dateTimeQuery", dateTimeQuery.String())
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateTimeNullHandleResponse handles the DateTimeNull response.
func (QueriesOperations) DateTimeNullHandleResponse(resp *azcore.Response) (*QueriesDateTimeNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDateTimeNullResponse{StatusCode: resp.StatusCode}, nil
}

// DateTimeValidCreateRequest creates the DateTimeValid request.
func (QueriesOperations) DateTimeValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/datetime/2012-01-01T01%3A01%3A01Z")
	query := u.Query()
	query.Set("dateTimeQuery", "2012-01-01T01:01:01Z")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateTimeValidHandleResponse handles the DateTimeValid response.
func (QueriesOperations) DateTimeValidHandleResponse(resp *azcore.Response) (*QueriesDateTimeValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDateTimeValidResponse{StatusCode: resp.StatusCode}, nil
}

// DateValidCreateRequest creates the DateValid request.
func (QueriesOperations) DateValidCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/date/2012-01-01")
	query := u.Query()
	query.Set("dateQuery", "2012-01-01")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DateValidHandleResponse handles the DateValid response.
func (QueriesOperations) DateValidHandleResponse(resp *azcore.Response) (*QueriesDateValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDateValidResponse{StatusCode: resp.StatusCode}, nil
}

// DoubleDecimalNegativeCreateRequest creates the DoubleDecimalNegative request.
func (QueriesOperations) DoubleDecimalNegativeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/double/-9999999.999")
	query := u.Query()
	query.Set("doubleQuery", "-9999999.999")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DoubleDecimalNegativeHandleResponse handles the DoubleDecimalNegative response.
func (QueriesOperations) DoubleDecimalNegativeHandleResponse(resp *azcore.Response) (*QueriesDoubleDecimalNegativeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDoubleDecimalNegativeResponse{StatusCode: resp.StatusCode}, nil
}

// DoubleDecimalPositiveCreateRequest creates the DoubleDecimalPositive request.
func (QueriesOperations) DoubleDecimalPositiveCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/double/9999999.999")
	query := u.Query()
	query.Set("doubleQuery", "9999999.999")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DoubleDecimalPositiveHandleResponse handles the DoubleDecimalPositive response.
func (QueriesOperations) DoubleDecimalPositiveHandleResponse(resp *azcore.Response) (*QueriesDoubleDecimalPositiveResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDoubleDecimalPositiveResponse{StatusCode: resp.StatusCode}, nil
}

// DoubleNullCreateRequest creates the DoubleNull request.
func (QueriesOperations) DoubleNullCreateRequest(u url.URL, doubleQuery float64) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/double/null")
	query := u.Query()
	query.Set("doubleQuery", strconv.FormatFloat(doubleQuery, 'f', -1, 64))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// DoubleNullHandleResponse handles the DoubleNull response.
func (QueriesOperations) DoubleNullHandleResponse(resp *azcore.Response) (*QueriesDoubleNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesDoubleNullResponse{StatusCode: resp.StatusCode}, nil
}

// EnumNullCreateRequest creates the EnumNull request.
func (QueriesOperations) EnumNullCreateRequest(u url.URL, enumQuery UriColor) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/enum/null")
	query := u.Query()
	query.Set("enumQuery", string(enumQuery))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// EnumNullHandleResponse handles the EnumNull response.
func (QueriesOperations) EnumNullHandleResponse(resp *azcore.Response) (*QueriesEnumNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesEnumNullResponse{StatusCode: resp.StatusCode}, nil
}

// EnumValidCreateRequest creates the EnumValid request.
func (QueriesOperations) EnumValidCreateRequest(u url.URL, enumQuery UriColor) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/enum/green%20color")
	query := u.Query()
	query.Set("enumQuery", string(enumQuery))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// EnumValidHandleResponse handles the EnumValid response.
func (QueriesOperations) EnumValidHandleResponse(resp *azcore.Response) (*QueriesEnumValidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesEnumValidResponse{StatusCode: resp.StatusCode}, nil
}

// FloatNullCreateRequest creates the FloatNull request.
func (QueriesOperations) FloatNullCreateRequest(u url.URL, floatQuery float32) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/float/null")
	query := u.Query()
	query.Set("floatQuery", strconv.FormatFloat(float64(floatQuery), 'f', -1, 32))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// FloatNullHandleResponse handles the FloatNull response.
func (QueriesOperations) FloatNullHandleResponse(resp *azcore.Response) (*QueriesFloatNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesFloatNullResponse{StatusCode: resp.StatusCode}, nil
}

// FloatScientificNegativeCreateRequest creates the FloatScientificNegative request.
func (QueriesOperations) FloatScientificNegativeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/float/-1.034E-20")
	query := u.Query()
	query.Set("floatQuery", "-1.034e-20")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// FloatScientificNegativeHandleResponse handles the FloatScientificNegative response.
func (QueriesOperations) FloatScientificNegativeHandleResponse(resp *azcore.Response) (*QueriesFloatScientificNegativeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesFloatScientificNegativeResponse{StatusCode: resp.StatusCode}, nil
}

// FloatScientificPositiveCreateRequest creates the FloatScientificPositive request.
func (QueriesOperations) FloatScientificPositiveCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/float/1.034E+20")
	query := u.Query()
	query.Set("floatQuery", "103400000000000000000")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// FloatScientificPositiveHandleResponse handles the FloatScientificPositive response.
func (QueriesOperations) FloatScientificPositiveHandleResponse(resp *azcore.Response) (*QueriesFloatScientificPositiveResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesFloatScientificPositiveResponse{StatusCode: resp.StatusCode}, nil
}

// GetBooleanFalseCreateRequest creates the GetBooleanFalse request.
func (QueriesOperations) GetBooleanFalseCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/bool/false")
	query := u.Query()
	query.Set("boolQuery", "false")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBooleanFalseHandleResponse handles the GetBooleanFalse response.
func (QueriesOperations) GetBooleanFalseHandleResponse(resp *azcore.Response) (*QueriesGetBooleanFalseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetBooleanFalseResponse{StatusCode: resp.StatusCode}, nil
}

// GetBooleanNullCreateRequest creates the GetBooleanNull request.
func (QueriesOperations) GetBooleanNullCreateRequest(u url.URL, boolQuery bool) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/bool/null")
	query := u.Query()
	query.Set("boolQuery", strconv.FormatBool(boolQuery))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBooleanNullHandleResponse handles the GetBooleanNull response.
func (QueriesOperations) GetBooleanNullHandleResponse(resp *azcore.Response) (*QueriesGetBooleanNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetBooleanNullResponse{StatusCode: resp.StatusCode}, nil
}

// GetBooleanTrueCreateRequest creates the GetBooleanTrue request.
func (QueriesOperations) GetBooleanTrueCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/bool/true")
	query := u.Query()
	query.Set("boolQuery", "true")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetBooleanTrueHandleResponse handles the GetBooleanTrue response.
func (QueriesOperations) GetBooleanTrueHandleResponse(resp *azcore.Response) (*QueriesGetBooleanTrueResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetBooleanTrueResponse{StatusCode: resp.StatusCode}, nil
}

// GetIntNegativeOneMillionCreateRequest creates the GetIntNegativeOneMillion request.
func (QueriesOperations) GetIntNegativeOneMillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/int/-1000000")
	query := u.Query()
	query.Set("intQuery", "-1000000")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntNegativeOneMillionHandleResponse handles the GetIntNegativeOneMillion response.
func (QueriesOperations) GetIntNegativeOneMillionHandleResponse(resp *azcore.Response) (*QueriesGetIntNegativeOneMillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetIntNegativeOneMillionResponse{StatusCode: resp.StatusCode}, nil
}

// GetIntNullCreateRequest creates the GetIntNull request.
func (QueriesOperations) GetIntNullCreateRequest(u url.URL, intQuery int32) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/int/null")
	query := u.Query()
	query.Set("intQuery", strconv.FormatInt(int64(intQuery), 10))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntNullHandleResponse handles the GetIntNull response.
func (QueriesOperations) GetIntNullHandleResponse(resp *azcore.Response) (*QueriesGetIntNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetIntNullResponse{StatusCode: resp.StatusCode}, nil
}

// GetIntOneMillionCreateRequest creates the GetIntOneMillion request.
func (QueriesOperations) GetIntOneMillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/int/1000000")
	query := u.Query()
	query.Set("intQuery", "1000000")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetIntOneMillionHandleResponse handles the GetIntOneMillion response.
func (QueriesOperations) GetIntOneMillionHandleResponse(resp *azcore.Response) (*QueriesGetIntOneMillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetIntOneMillionResponse{StatusCode: resp.StatusCode}, nil
}

// GetLongNullCreateRequest creates the GetLongNull request.
func (QueriesOperations) GetLongNullCreateRequest(u url.URL, longQuery int64) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/long/null")
	query := u.Query()
	query.Set("longQuery", strconv.FormatInt(longQuery, 10))
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetLongNullHandleResponse handles the GetLongNull response.
func (QueriesOperations) GetLongNullHandleResponse(resp *azcore.Response) (*QueriesGetLongNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetLongNullResponse{StatusCode: resp.StatusCode}, nil
}

// GetNegativeTenBillionCreateRequest creates the GetNegativeTenBillion request.
func (QueriesOperations) GetNegativeTenBillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/long/-10000000000")
	query := u.Query()
	query.Set("longQuery", "-10000000000")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetNegativeTenBillionHandleResponse handles the GetNegativeTenBillion response.
func (QueriesOperations) GetNegativeTenBillionHandleResponse(resp *azcore.Response) (*QueriesGetNegativeTenBillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetNegativeTenBillionResponse{StatusCode: resp.StatusCode}, nil
}

// GetTenBillionCreateRequest creates the GetTenBillion request.
func (QueriesOperations) GetTenBillionCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/long/10000000000")
	query := u.Query()
	query.Set("longQuery", "10000000000")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// GetTenBillionHandleResponse handles the GetTenBillion response.
func (QueriesOperations) GetTenBillionHandleResponse(resp *azcore.Response) (*QueriesGetTenBillionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesGetTenBillionResponse{StatusCode: resp.StatusCode}, nil
}

// StringEmptyCreateRequest creates the StringEmpty request.
func (QueriesOperations) StringEmptyCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/string/empty")
	query := u.Query()
	query.Set("stringQuery", "")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringEmptyHandleResponse handles the StringEmpty response.
func (QueriesOperations) StringEmptyHandleResponse(resp *azcore.Response) (*QueriesStringEmptyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesStringEmptyResponse{StatusCode: resp.StatusCode}, nil
}

// StringNullCreateRequest creates the StringNull request.
func (QueriesOperations) StringNullCreateRequest(u url.URL, stringQuery string) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/string/null")
	query := u.Query()
	query.Set("stringQuery", stringQuery)
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringNullHandleResponse handles the StringNull response.
func (QueriesOperations) StringNullHandleResponse(resp *azcore.Response) (*QueriesStringNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesStringNullResponse{StatusCode: resp.StatusCode}, nil
}

// StringURLEncodedCreateRequest creates the StringURLEncoded request.
func (QueriesOperations) StringURLEncodedCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/string/begin%21%2A%27%28%29%3B%3A%40%20%26%3D%2B%24%2C%2F%3F%23%5B%5Dend")
	query := u.Query()
	query.Set("stringQuery", "begin!*'();:@ &=+$,/?#[]end")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringURLEncodedHandleResponse handles the StringURLEncoded response.
func (QueriesOperations) StringURLEncodedHandleResponse(resp *azcore.Response) (*QueriesStringURLEncodedResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesStringURLEncodedResponse{StatusCode: resp.StatusCode}, nil
}

// StringUnicodeCreateRequest creates the StringUnicode request.
func (QueriesOperations) StringUnicodeCreateRequest(u url.URL) (*azcore.Request, error) {
	u.Path = path.Join(u.Path, "/queries/string/unicode/")
	query := u.Query()
	query.Set("stringQuery", "啊齄丂狛狜隣郎隣兀﨩")
	u.RawQuery = query.Encode()
	return azcore.NewRequest(http.MethodGet, u), nil
}

// StringUnicodeHandleResponse handles the StringUnicode response.
func (QueriesOperations) StringUnicodeHandleResponse(resp *azcore.Response) (*QueriesStringUnicodeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &QueriesStringUnicodeResponse{StatusCode: resp.StatusCode}, nil
}
