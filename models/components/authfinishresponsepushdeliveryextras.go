// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

type AuthFinishResponsePushDeliveryExtras struct {
	DistanceFromOriginatingIP *int64  `json:"distanceFromOriginatingIp,omitempty"`
	OriginatingIPGeolocation  *string `json:"originatingIpGeolocation,omitempty"`
}

func (o *AuthFinishResponsePushDeliveryExtras) GetDistanceFromOriginatingIP() *int64 {
	if o == nil {
		return nil
	}
	return o.DistanceFromOriginatingIP
}

func (o *AuthFinishResponsePushDeliveryExtras) GetOriginatingIPGeolocation() *string {
	if o == nil {
		return nil
	}
	return o.OriginatingIPGeolocation
}
