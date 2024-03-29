// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/speakeasy-sdks/test-collab/internal/utils"
)

type UserInfoRequest struct {
	// RequestID is a UUID generated on the customer side to be associated with the unique request for tracking.
	// Acceptable characters are: alphanumeric with symbols '-._+=/'.
	RequestID *string `default:"eba12f3a-5555-47bc-b85d-21c0cbc4b973" json:"requestId"`
	// UserID is a user identifier generated on the customer side to be associated with a device.
	// Acceptable characters are: alphanumeric with symbols '-._+=/:'.
	UserID string `json:"userId"`
}

func (u UserInfoRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UserInfoRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UserInfoRequest) GetRequestID() *string {
	if o == nil {
		return nil
	}
	return o.RequestID
}

func (o *UserInfoRequest) GetUserID() string {
	if o == nil {
		return ""
	}
	return o.UserID
}
