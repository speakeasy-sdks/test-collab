// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-collab/models/components"
	"net/http"
)

type DeviceUnregisterRequestResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Result of the operation.
	DeviceUnregisterResponse *components.DeviceUnregisterResponse
}

func (o *DeviceUnregisterRequestResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeviceUnregisterRequestResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeviceUnregisterRequestResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeviceUnregisterRequestResponse) GetDeviceUnregisterResponse() *components.DeviceUnregisterResponse {
	if o == nil {
		return nil
	}
	return o.DeviceUnregisterResponse
}