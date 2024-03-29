// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/test-collab/internal/utils"
)

type AuthFinishRequest struct {
	// AuthID is the unique ID of the auth flow to retrieve the result.
	AuthID *string `default:"eba12f3a-5555-47bc-b85d-21c0cbc4b973" json:"authId"`
}

func (a AuthFinishRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AuthFinishRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AuthFinishRequest) GetAuthID() *string {
	if o == nil {
		return nil
	}
	return o.AuthID
}
