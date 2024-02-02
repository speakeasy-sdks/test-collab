// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponseDelivery struct {
	Push *AuthFinishResponsePushDelivery `json:"push,omitempty"`
	Scan *AuthFinishResponseScanDelivery `json:"scan,omitempty"`
}

func (o *AuthFinishResponseDelivery) GetPush() *AuthFinishResponsePushDelivery {
	if o == nil {
		return nil
	}
	return o.Push
}

func (o *AuthFinishResponseDelivery) GetScan() *AuthFinishResponseScanDelivery {
	if o == nil {
		return nil
	}
	return o.Scan
}
