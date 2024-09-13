/*
Shipping APIs

### Introduction  The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.   Some of the key API operations available in the Shipping API includes: ### Shipment API  | Operation      | Description | | ----------- | ----------- |  | Get Carriers    | This operation fetches all onboarded carriers. Typically, user will use this service to get list of onboarded carriers and supported properties for those carriers.  |  | Get Countries | This operation fetches list of supported destination countries for a provided carrier and origin country.  |  | Get Services | This operation fetches a list of supported services for a carrier with respect to specific origin and destination country. |  | Get ParcelTypes| This operation fetches ParcelTypes based on carrier, origin and destination country. |  | Get Special Services| This operation fetches Special Services for a given carrier, service, origin and destination country. |  | Get Carrier Accounts| This operation retrieves onboarded Carriers with their Carrier Account Ids which uniquely identify multiple accounts of same carrier.  |  | Rate Shop and Get Single Rate| This API contains 2 operations, rate shop and single rate. Rate shop will fetch rates for all carrier services based on the given addresses (From and To), weight, and dimension for given parcelType. Single rate will get rate for specific service and special service (if requested) based on the given addresses (From and To), weight, and dimension, parcelType and serviceId with or without specialServices. Single rate will be used mainly to a rate a shipment before creating shipment.  |  | Create Shipment| This operation creates a new Shipment or Shipment Label. This is for both Domestic and International. | | Get All Shipments| This operation fetches all created Shipments. |  | Get Shipment by Id| Retrieves single shipment using Shipment Id. |  | Reprint Shipment| This operation reprints Shipment by the shipmentId. It retrieves an existing shipping label to reprint. The API sends the shipmentId returned by the original Created Shipment request. Use this only if the shipping label in the Create Shipment response was spoilt or lost. |  | Cancel Shipment| This operation cancels previously created shipment. |  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipping

import (
	"encoding/json"
)

// checks if the MultipieceRateShopResponseRatesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MultipieceRateShopResponseRatesInner{}

// MultipieceRateShopResponseRatesInner description
type MultipieceRateShopResponseRatesInner struct {
	// description
	TotalCarrierCharge *float32 `json:"totalCarrierCharge,omitempty"`
	// description
	Carrier *string `json:"carrier,omitempty"`
	// description
	CurrencyCode *string `json:"currencyCode,omitempty"`
	// description
	RateTypeId *string `json:"rateTypeId,omitempty"`
	// description
	ServiceId *string `json:"serviceId,omitempty"`
	DeliveryCommitment *RateShopResponseRatesInnerDeliveryCommitment `json:"deliveryCommitment,omitempty"`
	// description
	MultiPieceParcels []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner `json:"multiPieceParcels,omitempty"`
	// description
	Surcharges []MultipieceRatesResponseRatesInnerSurchargesInner `json:"surcharges,omitempty"`
}

// NewMultipieceRateShopResponseRatesInner instantiates a new MultipieceRateShopResponseRatesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMultipieceRateShopResponseRatesInner() *MultipieceRateShopResponseRatesInner {
	this := MultipieceRateShopResponseRatesInner{}
	return &this
}

// NewMultipieceRateShopResponseRatesInnerWithDefaults instantiates a new MultipieceRateShopResponseRatesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMultipieceRateShopResponseRatesInnerWithDefaults() *MultipieceRateShopResponseRatesInner {
	this := MultipieceRateShopResponseRatesInner{}
	return &this
}

// GetTotalCarrierCharge returns the TotalCarrierCharge field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetTotalCarrierCharge() float32 {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		var ret float32
		return ret
	}
	return *o.TotalCarrierCharge
}

// GetTotalCarrierChargeOk returns a tuple with the TotalCarrierCharge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetTotalCarrierChargeOk() (*float32, bool) {
	if o == nil || IsNil(o.TotalCarrierCharge) {
		return nil, false
	}
	return o.TotalCarrierCharge, true
}

// HasTotalCarrierCharge returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasTotalCarrierCharge() bool {
	if o != nil && !IsNil(o.TotalCarrierCharge) {
		return true
	}

	return false
}

// SetTotalCarrierCharge gets a reference to the given float32 and assigns it to the TotalCarrierCharge field.
func (o *MultipieceRateShopResponseRatesInner) SetTotalCarrierCharge(v float32) {
	o.TotalCarrierCharge = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *MultipieceRateShopResponseRatesInner) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCurrencyCode returns the CurrencyCode field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetCurrencyCode() string {
	if o == nil || IsNil(o.CurrencyCode) {
		var ret string
		return ret
	}
	return *o.CurrencyCode
}

// GetCurrencyCodeOk returns a tuple with the CurrencyCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetCurrencyCodeOk() (*string, bool) {
	if o == nil || IsNil(o.CurrencyCode) {
		return nil, false
	}
	return o.CurrencyCode, true
}

// HasCurrencyCode returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasCurrencyCode() bool {
	if o != nil && !IsNil(o.CurrencyCode) {
		return true
	}

	return false
}

// SetCurrencyCode gets a reference to the given string and assigns it to the CurrencyCode field.
func (o *MultipieceRateShopResponseRatesInner) SetCurrencyCode(v string) {
	o.CurrencyCode = &v
}

// GetRateTypeId returns the RateTypeId field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetRateTypeId() string {
	if o == nil || IsNil(o.RateTypeId) {
		var ret string
		return ret
	}
	return *o.RateTypeId
}

// GetRateTypeIdOk returns a tuple with the RateTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetRateTypeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RateTypeId) {
		return nil, false
	}
	return o.RateTypeId, true
}

// HasRateTypeId returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasRateTypeId() bool {
	if o != nil && !IsNil(o.RateTypeId) {
		return true
	}

	return false
}

// SetRateTypeId gets a reference to the given string and assigns it to the RateTypeId field.
func (o *MultipieceRateShopResponseRatesInner) SetRateTypeId(v string) {
	o.RateTypeId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *MultipieceRateShopResponseRatesInner) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetDeliveryCommitment returns the DeliveryCommitment field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetDeliveryCommitment() RateShopResponseRatesInnerDeliveryCommitment {
	if o == nil || IsNil(o.DeliveryCommitment) {
		var ret RateShopResponseRatesInnerDeliveryCommitment
		return ret
	}
	return *o.DeliveryCommitment
}

// GetDeliveryCommitmentOk returns a tuple with the DeliveryCommitment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetDeliveryCommitmentOk() (*RateShopResponseRatesInnerDeliveryCommitment, bool) {
	if o == nil || IsNil(o.DeliveryCommitment) {
		return nil, false
	}
	return o.DeliveryCommitment, true
}

// HasDeliveryCommitment returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasDeliveryCommitment() bool {
	if o != nil && !IsNil(o.DeliveryCommitment) {
		return true
	}

	return false
}

// SetDeliveryCommitment gets a reference to the given RateShopResponseRatesInnerDeliveryCommitment and assigns it to the DeliveryCommitment field.
func (o *MultipieceRateShopResponseRatesInner) SetDeliveryCommitment(v RateShopResponseRatesInnerDeliveryCommitment) {
	o.DeliveryCommitment = &v
}

// GetMultiPieceParcels returns the MultiPieceParcels field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetMultiPieceParcels() []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner {
	if o == nil || IsNil(o.MultiPieceParcels) {
		var ret []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner
		return ret
	}
	return o.MultiPieceParcels
}

// GetMultiPieceParcelsOk returns a tuple with the MultiPieceParcels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetMultiPieceParcelsOk() ([]MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner, bool) {
	if o == nil || IsNil(o.MultiPieceParcels) {
		return nil, false
	}
	return o.MultiPieceParcels, true
}

// HasMultiPieceParcels returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasMultiPieceParcels() bool {
	if o != nil && !IsNil(o.MultiPieceParcels) {
		return true
	}

	return false
}

// SetMultiPieceParcels gets a reference to the given []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner and assigns it to the MultiPieceParcels field.
func (o *MultipieceRateShopResponseRatesInner) SetMultiPieceParcels(v []MultipieceRateShopResponseRatesInnerMultiPieceParcelsInner) {
	o.MultiPieceParcels = v
}

// GetSurcharges returns the Surcharges field value if set, zero value otherwise.
func (o *MultipieceRateShopResponseRatesInner) GetSurcharges() []MultipieceRatesResponseRatesInnerSurchargesInner {
	if o == nil || IsNil(o.Surcharges) {
		var ret []MultipieceRatesResponseRatesInnerSurchargesInner
		return ret
	}
	return o.Surcharges
}

// GetSurchargesOk returns a tuple with the Surcharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MultipieceRateShopResponseRatesInner) GetSurchargesOk() ([]MultipieceRatesResponseRatesInnerSurchargesInner, bool) {
	if o == nil || IsNil(o.Surcharges) {
		return nil, false
	}
	return o.Surcharges, true
}

// HasSurcharges returns a boolean if a field has been set.
func (o *MultipieceRateShopResponseRatesInner) HasSurcharges() bool {
	if o != nil && !IsNil(o.Surcharges) {
		return true
	}

	return false
}

// SetSurcharges gets a reference to the given []MultipieceRatesResponseRatesInnerSurchargesInner and assigns it to the Surcharges field.
func (o *MultipieceRateShopResponseRatesInner) SetSurcharges(v []MultipieceRatesResponseRatesInnerSurchargesInner) {
	o.Surcharges = v
}

func (o MultipieceRateShopResponseRatesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MultipieceRateShopResponseRatesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TotalCarrierCharge) {
		toSerialize["totalCarrierCharge"] = o.TotalCarrierCharge
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.CurrencyCode) {
		toSerialize["currencyCode"] = o.CurrencyCode
	}
	if !IsNil(o.RateTypeId) {
		toSerialize["rateTypeId"] = o.RateTypeId
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.DeliveryCommitment) {
		toSerialize["deliveryCommitment"] = o.DeliveryCommitment
	}
	if !IsNil(o.MultiPieceParcels) {
		toSerialize["multiPieceParcels"] = o.MultiPieceParcels
	}
	if !IsNil(o.Surcharges) {
		toSerialize["surcharges"] = o.Surcharges
	}
	return toSerialize, nil
}

type NullableMultipieceRateShopResponseRatesInner struct {
	value *MultipieceRateShopResponseRatesInner
	isSet bool
}

func (v NullableMultipieceRateShopResponseRatesInner) Get() *MultipieceRateShopResponseRatesInner {
	return v.value
}

func (v *NullableMultipieceRateShopResponseRatesInner) Set(val *MultipieceRateShopResponseRatesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMultipieceRateShopResponseRatesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMultipieceRateShopResponseRatesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMultipieceRateShopResponseRatesInner(val *MultipieceRateShopResponseRatesInner) *NullableMultipieceRateShopResponseRatesInner {
	return &NullableMultipieceRateShopResponseRatesInner{value: val, isSet: true}
}

func (v NullableMultipieceRateShopResponseRatesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMultipieceRateShopResponseRatesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


