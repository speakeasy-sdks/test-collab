// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponseSubjectDeviceResult struct {
	Error   *AuthFinishResponseError                      `json:"error,omitempty"`
	Success *AuthFinishResponseSubjectDeviceResultSuccess `json:"success,omitempty"`
}

func (o *AuthFinishResponseSubjectDeviceResult) GetError() *AuthFinishResponseError {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *AuthFinishResponseSubjectDeviceResult) GetSuccess() *AuthFinishResponseSubjectDeviceResultSuccess {
	if o == nil {
		return nil
	}
	return o.Success
}