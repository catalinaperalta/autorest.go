// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	azinternal "generatortests/autorest/generated/complexgroup/internal/complexgroup"
)

// PrimitiveOperations contains the methods for the Primitive group.
type PrimitiveOperations interface {
	// GetBool - Get complex types with bool properties 
	GetBool(ctx context.Context) (*PrimitiveGetBoolResponse, error)
	// GetByte - Get complex types with byte properties 
	GetByte(ctx context.Context) (*PrimitiveGetByteResponse, error)
	// GetDate - Get complex types with date properties 
	GetDate(ctx context.Context) (*PrimitiveGetDateResponse, error)
	// GetDateTime - Get complex types with datetime properties 
	GetDateTime(ctx context.Context) (*PrimitiveGetDateTimeResponse, error)
	// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties 
	GetDateTimeRFC1123(ctx context.Context) (*PrimitiveGetDateTimeRFC1123Response, error)
	// GetDouble - Get complex types with double properties 
	GetDouble(ctx context.Context) (*PrimitiveGetDoubleResponse, error)
	// GetDuration - Get complex types with duration properties 
	GetDuration(ctx context.Context) (*PrimitiveGetDurationResponse, error)
	// GetFloat - Get complex types with float properties 
	GetFloat(ctx context.Context) (*PrimitiveGetFloatResponse, error)
	// GetInt - Get complex types with integer properties 
	GetInt(ctx context.Context) (*PrimitiveGetIntResponse, error)
	// GetLong - Get complex types with long properties 
	GetLong(ctx context.Context) (*PrimitiveGetLongResponse, error)
	// GetString - Get complex types with string properties 
	GetString(ctx context.Context) (*PrimitiveGetStringResponse, error)
	// PutBool - Put complex types with bool properties 
	PutBool(ctx context.Context, complexBody BooleanWrapper) (*PrimitivePutBoolResponse, error)
	// PutByte - Put complex types with byte properties 
	PutByte(ctx context.Context, complexBody ByteWrapper) (*PrimitivePutByteResponse, error)
	// PutDate - Put complex types with date properties 
	PutDate(ctx context.Context, complexBody DateWrapper) (*PrimitivePutDateResponse, error)
	// PutDateTime - Put complex types with datetime properties 
	PutDateTime(ctx context.Context, complexBody DatetimeWrapper) (*PrimitivePutDateTimeResponse, error)
	// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties 
	PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper) (*PrimitivePutDateTimeRFC1123Response, error)
	// PutDouble - Put complex types with double properties 
	PutDouble(ctx context.Context, complexBody DoubleWrapper) (*PrimitivePutDoubleResponse, error)
	// PutDuration - Put complex types with duration properties 
	PutDuration(ctx context.Context, complexBody DurationWrapper) (*PrimitivePutDurationResponse, error)
	// PutFloat - Put complex types with float properties 
	PutFloat(ctx context.Context, complexBody FloatWrapper) (*PrimitivePutFloatResponse, error)
	// PutInt - Put complex types with integer properties 
	PutInt(ctx context.Context, complexBody IntWrapper) (*PrimitivePutIntResponse, error)
	// PutLong - Put complex types with long properties 
	PutLong(ctx context.Context, complexBody LongWrapper) (*PrimitivePutLongResponse, error)
	// PutString - Put complex types with string properties 
	PutString(ctx context.Context, complexBody StringWrapper) (*PrimitivePutStringResponse, error)
}

type primitiveOperations struct {
	*Client
	azinternal.PrimitiveOperations
}

// GetBool - Get complex types with bool properties 
func (client *primitiveOperations) GetBool(ctx context.Context) (*PrimitiveGetBoolResponse, error) {
	req, err := client.GetBoolCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetByte - Get complex types with byte properties 
func (client *primitiveOperations) GetByte(ctx context.Context) (*PrimitiveGetByteResponse, error) {
	req, err := client.GetByteCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDate - Get complex types with date properties 
func (client *primitiveOperations) GetDate(ctx context.Context) (*PrimitiveGetDateResponse, error) {
	req, err := client.GetDateCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDateTime - Get complex types with datetime properties 
func (client *primitiveOperations) GetDateTime(ctx context.Context) (*PrimitiveGetDateTimeResponse, error) {
	req, err := client.GetDateTimeCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties 
func (client *primitiveOperations) GetDateTimeRFC1123(ctx context.Context) (*PrimitiveGetDateTimeRFC1123Response, error) {
	req, err := client.GetDateTimeRFC1123CreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetDateTimeRFC1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDouble - Get complex types with double properties 
func (client *primitiveOperations) GetDouble(ctx context.Context) (*PrimitiveGetDoubleResponse, error) {
	req, err := client.GetDoubleCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetDuration - Get complex types with duration properties 
func (client *primitiveOperations) GetDuration(ctx context.Context) (*PrimitiveGetDurationResponse, error) {
	req, err := client.GetDurationCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetFloat - Get complex types with float properties 
func (client *primitiveOperations) GetFloat(ctx context.Context) (*PrimitiveGetFloatResponse, error) {
	req, err := client.GetFloatCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetInt - Get complex types with integer properties 
func (client *primitiveOperations) GetInt(ctx context.Context) (*PrimitiveGetIntResponse, error) {
	req, err := client.GetIntCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetLong - Get complex types with long properties 
func (client *primitiveOperations) GetLong(ctx context.Context) (*PrimitiveGetLongResponse, error) {
	req, err := client.GetLongCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetString - Get complex types with string properties 
func (client *primitiveOperations) GetString(ctx context.Context) (*PrimitiveGetStringResponse, error) {
	req, err := client.GetStringCreateRequest(*client.u)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutBool - Put complex types with bool properties 
func (client *primitiveOperations) PutBool(ctx context.Context, complexBody BooleanWrapper) (*PrimitivePutBoolResponse, error) {
	req, err := client.PutBoolCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutBoolHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutByte - Put complex types with byte properties 
func (client *primitiveOperations) PutByte(ctx context.Context, complexBody ByteWrapper) (*PrimitivePutByteResponse, error) {
	req, err := client.PutByteCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutByteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutDate - Put complex types with date properties 
func (client *primitiveOperations) PutDate(ctx context.Context, complexBody DateWrapper) (*PrimitivePutDateResponse, error) {
	req, err := client.PutDateCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutDateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutDateTime - Put complex types with datetime properties 
func (client *primitiveOperations) PutDateTime(ctx context.Context, complexBody DatetimeWrapper) (*PrimitivePutDateTimeResponse, error) {
	req, err := client.PutDateTimeCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutDateTimeHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties 
func (client *primitiveOperations) PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper) (*PrimitivePutDateTimeRFC1123Response, error) {
	req, err := client.PutDateTimeRFC1123CreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutDateTimeRFC1123HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutDouble - Put complex types with double properties 
func (client *primitiveOperations) PutDouble(ctx context.Context, complexBody DoubleWrapper) (*PrimitivePutDoubleResponse, error) {
	req, err := client.PutDoubleCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutDoubleHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutDuration - Put complex types with duration properties 
func (client *primitiveOperations) PutDuration(ctx context.Context, complexBody DurationWrapper) (*PrimitivePutDurationResponse, error) {
	req, err := client.PutDurationCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutDurationHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutFloat - Put complex types with float properties 
func (client *primitiveOperations) PutFloat(ctx context.Context, complexBody FloatWrapper) (*PrimitivePutFloatResponse, error) {
	req, err := client.PutFloatCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutFloatHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutInt - Put complex types with integer properties 
func (client *primitiveOperations) PutInt(ctx context.Context, complexBody IntWrapper) (*PrimitivePutIntResponse, error) {
	req, err := client.PutIntCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutIntHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutLong - Put complex types with long properties 
func (client *primitiveOperations) PutLong(ctx context.Context, complexBody LongWrapper) (*PrimitivePutLongResponse, error) {
	req, err := client.PutLongCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutLongHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutString - Put complex types with string properties 
func (client *primitiveOperations) PutString(ctx context.Context, complexBody StringWrapper) (*PrimitivePutStringResponse, error) {
	req, err := client.PutStringCreateRequest(*client.u, complexBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.p.Do(ctx, req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutStringHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

var _ PrimitiveOperations = (*primitiveOperations)(nil)