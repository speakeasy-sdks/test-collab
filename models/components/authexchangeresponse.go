// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/test-collab/models/sdkerrors"
)

type AuthExchangeResponse struct {
	// AuthToken is a bearer token for use by the Prove Auth SDK.
	AuthToken string           `json:"authToken"`
	Error     *sdkerrors.Error `json:"error,omitempty"`
}

func (o *AuthExchangeResponse) GetAuthToken() string {
	if o == nil {
		return ""
	}
	return o.AuthToken
}

func (o *AuthExchangeResponse) GetError() *sdkerrors.Error {
	if o == nil {
		return nil
	}
	return o.Error
}