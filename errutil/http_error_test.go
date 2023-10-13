package errutil

import (
	"context"
	"errors"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestParseRPCErrorToHTTP(t *testing.T) {
	ctx := context.Background()
	uCode := EmailAlreadyRegistered
	grpcCode := errorMap[uCode].GrpcCode
	uDesc := "Email already registered."
	sysErr := errors.New("some error")

	t.Run("status", func(t *testing.T) {
		_status := status.New(grpcCode, sysErr.Error())
		httpError := ParseRPCErrorToHTTP(ctx, _status.Err())
		assertHttpError(t, *httpError, HTTPError{
			StatusCode: http.StatusInternalServerError,
			Code:       InternalServerError,
			Desc:       "Internal Service Error",
		})
	})
	t.Run("status with detail", func(t *testing.T) {
		errDetail := &ErrorDetail{
			Code:       uCode,
			GrpcCode:   grpcCode,
			Desc:       uDesc,
			Error:      sysErr.Error(),
			CallerInfo: callerInfo(ctx),
		}
		_status, _ := status.New(grpcCode, sysErr.Error()).WithDetails(errDetail.ToPB())
		httpError := ParseRPCErrorToHTTP(ctx, _status.Err())
		assertHttpError(t, *httpError, HTTPError{
			StatusCode: grpcCodeToHTTPStatusCode[errorMap[uCode].GrpcCode],
			Code:       uCode,
			Desc:       uDesc,
		})
	})
	t.Run("status with detail and different grpc codes", func(t *testing.T) {
		diffGrpcCode := codes.Unknown
		assert.NotEqual(t, grpcCode, diffGrpcCode)

		errDetail := &ErrorDetail{
			Code:       uCode,
			GrpcCode:   grpcCode,
			Desc:       uDesc,
			Error:      sysErr.Error(),
			CallerInfo: callerInfo(ctx),
		}
		_status, _ := status.New(diffGrpcCode, sysErr.Error()).WithDetails(errDetail.ToPB())

		httpError := ParseRPCErrorToHTTP(ctx, _status.Err())
		assertHttpError(t, *httpError, HTTPError{
			StatusCode: grpcCodeToHTTPStatusCode[errorMap[uCode].GrpcCode],
			Code:       uCode,
			Desc:       uDesc,
		})
	})
	t.Run("status with detail and different system error", func(t *testing.T) {
		barErr := errors.New("bar error")
		assert.NotEqual(t, sysErr.Error(), barErr.Error())

		errDetail := &ErrorDetail{
			Code:       uCode,
			GrpcCode:   grpcCode,
			Desc:       uDesc,
			Error:      sysErr.Error(),
			CallerInfo: callerInfo(ctx),
		}
		_status, _ := status.New(grpcCode, barErr.Error()).WithDetails(errDetail.ToPB())

		httpError := ParseRPCErrorToHTTP(ctx, _status.Err())
		assertHttpError(t, *httpError, HTTPError{
			StatusCode: grpcCodeToHTTPStatusCode[errorMap[uCode].GrpcCode],
			Code:       uCode,
			Desc:       uDesc,
		})
	})
}

/*
Added this test case to assert a fix that will prevent the caller of "GetHTTPError"
from sharing and editing the default HTTP Errors as defined in "defaultHttpErrorMap".

Bug "Editing shared HTTP Errors":

	After 1st call: Internal Service Error: require query parameter \"YNumber\"
	After 2nd call: Internal Service Error: require query parameter \"YNumber\": require query parameter \"YNumber\"
	etc..
*/
func TestGetHTTPErrorExpectUniqueReference(t *testing.T) {
	ctx := context.Background()

	// get and assert 1st HTTP Error "InternalServerError"
	httpError1, ok := GetHTTPError(ctx, InternalServerError)
	if !ok {
		t.Fatalf("http error %q should exist", InternalServerError)
	}
	assertHttpError(t, *httpError1, HTTPError{
		StatusCode: 500,
		RequestID:  "",
		Code:       InternalServerError,
		Desc:       "Internal Service Error",
	})

	// get and assert 2nd HTTP Error "InternalServerError"
	httpError2, ok := GetHTTPError(ctx, InternalServerError)
	if !ok {
		t.Fatalf("http error %q should exist", InternalServerError)
	}
	assertHttpError(t, *httpError2, HTTPError{
		StatusCode: 500,
		RequestID:  "",
		Code:       InternalServerError,
		Desc:       "Internal Service Error",
	})

	if httpError1 == httpError2 || reflect.ValueOf(httpError1).Kind() != reflect.Ptr {
		t.Errorf("GetHTTPError does not return unique reference (&httpError1 != &httError2)")
	}
}

func TestAddHttpErrorDefsShouldOverrideExistingErrors(t *testing.T) {
	ctx := context.Background()

	// get and assert 1st HTTP Error "InternalServerError"
	oldError, ok := GetHTTPError(ctx, InternalServerError)
	if !ok {
		t.Fatalf("http error %q should exist", InternalServerError)
	}
	assertHttpError(t, *oldError, HTTPError{
		StatusCode: 500,
		RequestID:  "",
		Code:       InternalServerError,
		Desc:       "Internal Service Error",
	})

	// create new HTTP Error which will override HTTPError with ErrorCode "InternalServerError"
	newError := HTTPError{
		StatusCode: 555,
		RequestID:  "",
		Code:       InternalServerError,
		Desc:       "foo",
	}
	AddHttpErrorDefs([]*HTTPError{&newError})

	// assert that we have overridden InternalServerError
	httpError, ok := GetHTTPError(ctx, InternalServerError)
	if !ok {
		t.Fatalf("http error %q should exist", InternalServerError)
	}
	if httpError.StatusCode != newError.StatusCode || httpError.StatusCode == oldError.StatusCode {
		t.Errorf("StatusCode: got %d, want %d", httpError.StatusCode, newError.StatusCode)
	}
	if httpError.Desc != newError.Desc || httpError.Desc == oldError.Desc {
		t.Errorf("Desc: got %q, want %q", httpError.Desc, newError.Desc)
	}
}

func assertHttpError(t *testing.T, got, want HTTPError) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
