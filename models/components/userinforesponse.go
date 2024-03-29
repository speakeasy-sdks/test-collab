// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/test-collab/models/sdkerrors"
)

type UserInfoResponse struct {
	Error *sdkerrors.Error `json:"error,omitempty"`
	// UserDevices contains the list of user device records.
	UserDevices []UserInfoDeviceResponse `json:"userDevices,omitempty"`
}

func (o *UserInfoResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *UserInfoResponse) GetUserDevices() []UserInfoDeviceResponse {
	if o == nil {
		return nil
	}
	return o.UserDevices
}
