// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/test-collab/internal/utils"
)

type AuthExchangeRequest struct {
	// AuthID is the unique ID of the auth flow to exchange for a token.
	AuthID *string `default:"eba12f3a-5555-47bc-b85d-21c0cbc4b973" json:"authId"`
}

func (a AuthExchangeRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthExchangeRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthExchangeRequest) GetAuthID() *string {
	if o == nil {
		return nil
	}
	return o.AuthID
}