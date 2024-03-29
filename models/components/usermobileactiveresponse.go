// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/test-collab/models/sdkerrors"
)

type UserMobileActiveResponse struct {
	Error *sdkerrors.Error `json:"error,omitempty"`
	// Success is true if the operation was successful.
	Success bool `json:"success"`
}

func (o *UserMobileActiveResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *UserMobileActiveResponse) GetSuccess() bool {
	if o == nil {
		return false
	}
	return o.Success
}
