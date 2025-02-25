// Code generated by go-swagger; DO NOT EDIT.

package version_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models2 "github.com/percona/percona-server-mysql-operator/pkg/version/service/client/models"
)

// VersionServiceApplyReader is a Reader for the VersionServiceApply structure.
type VersionServiceApplyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VersionServiceApplyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVersionServiceApplyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVersionServiceApplyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVersionServiceApplyOK creates a VersionServiceApplyOK with default headers values
func NewVersionServiceApplyOK() *VersionServiceApplyOK {
	return &VersionServiceApplyOK{}
}

/*
VersionServiceApplyOK describes a response with status code 200, with default header values.

A successful response.
*/
type VersionServiceApplyOK struct {
	Payload *models2.VersionVersionResponse
}

func (o *VersionServiceApplyOK) Error() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}/{apply}][%d] versionServiceApplyOK  %+v", 200, o.Payload)
}
func (o *VersionServiceApplyOK) GetPayload() *models2.VersionVersionResponse {
	return o.Payload
}

func (o *VersionServiceApplyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.VersionVersionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVersionServiceApplyDefault creates a VersionServiceApplyDefault with default headers values
func NewVersionServiceApplyDefault(code int) *VersionServiceApplyDefault {
	return &VersionServiceApplyDefault{
		_statusCode: code,
	}
}

/*
VersionServiceApplyDefault describes a response with status code -1, with default header values.

An unexpected error response
*/
type VersionServiceApplyDefault struct {
	_statusCode int

	Payload *models2.GooglerpcStatus
}

// Code gets the status code for the version service apply default response
func (o *VersionServiceApplyDefault) Code() int {
	return o._statusCode
}

func (o *VersionServiceApplyDefault) Error() string {
	return fmt.Sprintf("[GET /versions/v1/{product}/{operatorVersion}/{apply}][%d] VersionService_Apply default  %+v", o._statusCode, o.Payload)
}
func (o *VersionServiceApplyDefault) GetPayload() *models2.GooglerpcStatus {
	return o.Payload
}

func (o *VersionServiceApplyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models2.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
