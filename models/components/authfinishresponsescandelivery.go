// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponseScanDelivery struct {
	Extras   *AuthFinishResponseScanDeliveryExtras `json:"extras,omitempty"`
	Response *string                               `json:"response,omitempty"`
}

func (o *AuthFinishResponseScanDelivery) GetExtras() *AuthFinishResponseScanDeliveryExtras {
	if o == nil {
		return nil
	}
	return o.Extras
}

func (o *AuthFinishResponseScanDelivery) GetResponse() *string {
	if o == nil {
		return nil
	}
	return o.Response
}