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

// checks if the RateShopResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RateShopResponse{}

// RateShopResponse struct for RateShopResponse
type RateShopResponse struct {
	FromAddress *RateShopResponseFromAddress `json:"fromAddress,omitempty"`
	Parcel *RateShopResponseParcel `json:"parcel,omitempty"`
	// It displays all available rates for each services
	Rates []RateShopResponseRatesInner `json:"rates,omitempty"`
	ToAddress *RateShopResponseToAddress `json:"toAddress,omitempty"`
	// isHazmat if set to true,only services which support Hazardous material shipment would be visible in the response
	IsHazmat *bool `json:"isHazmat,omitempty"`
	// It display any error while getting rates
	Errors []RateShopResponseErrorsInner `json:"errors,omitempty"`
}

// NewRateShopResponse instantiates a new RateShopResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateShopResponse() *RateShopResponse {
	this := RateShopResponse{}
	return &this
}

// NewRateShopResponseWithDefaults instantiates a new RateShopResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateShopResponseWithDefaults() *RateShopResponse {
	this := RateShopResponse{}
	return &this
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *RateShopResponse) GetFromAddress() RateShopResponseFromAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret RateShopResponseFromAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponse) GetFromAddressOk() (*RateShopResponseFromAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *RateShopResponse) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given RateShopResponseFromAddress and assigns it to the FromAddress field.
func (o *RateShopResponse) SetFromAddress(v RateShopResponseFromAddress) {
	o.FromAddress = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *RateShopResponse) GetParcel() RateShopResponseParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret RateShopResponseParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponse) GetParcelOk() (*RateShopResponseParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *RateShopResponse) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given RateShopResponseParcel and assigns it to the Parcel field.
func (o *RateShopResponse) SetParcel(v RateShopResponseParcel) {
	o.Parcel = &v
}

// GetRates returns the Rates field value if set, zero value otherwise.
func (o *RateShopResponse) GetRates() []RateShopResponseRatesInner {
	if o == nil || IsNil(o.Rates) {
		var ret []RateShopResponseRatesInner
		return ret
	}
	return o.Rates
}

// GetRatesOk returns a tuple with the Rates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponse) GetRatesOk() ([]RateShopResponseRatesInner, bool) {
	if o == nil || IsNil(o.Rates) {
		return nil, false
	}
	return o.Rates, true
}

// HasRates returns a boolean if a field has been set.
func (o *RateShopResponse) HasRates() bool {
	if o != nil && !IsNil(o.Rates) {
		return true
	}

	return false
}

// SetRates gets a reference to the given []RateShopResponseRatesInner and assigns it to the Rates field.
func (o *RateShopResponse) SetRates(v []RateShopResponseRatesInner) {
	o.Rates = v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *RateShopResponse) GetToAddress() RateShopResponseToAddress {
	if o == nil || IsNil(o.ToAddress) {
		var ret RateShopResponseToAddress
		return ret
	}
	return *o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponse) GetToAddressOk() (*RateShopResponseToAddress, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *RateShopResponse) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given RateShopResponseToAddress and assigns it to the ToAddress field.
func (o *RateShopResponse) SetToAddress(v RateShopResponseToAddress) {
	o.ToAddress = &v
}

// GetIsHazmat returns the IsHazmat field value if set, zero value otherwise.
func (o *RateShopResponse) GetIsHazmat() bool {
	if o == nil || IsNil(o.IsHazmat) {
		var ret bool
		return ret
	}
	return *o.IsHazmat
}

// GetIsHazmatOk returns a tuple with the IsHazmat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponse) GetIsHazmatOk() (*bool, bool) {
	if o == nil || IsNil(o.IsHazmat) {
		return nil, false
	}
	return o.IsHazmat, true
}

// HasIsHazmat returns a boolean if a field has been set.
func (o *RateShopResponse) HasIsHazmat() bool {
	if o != nil && !IsNil(o.IsHazmat) {
		return true
	}

	return false
}

// SetIsHazmat gets a reference to the given bool and assigns it to the IsHazmat field.
func (o *RateShopResponse) SetIsHazmat(v bool) {
	o.IsHazmat = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *RateShopResponse) GetErrors() []RateShopResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []RateShopResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateShopResponse) GetErrorsOk() ([]RateShopResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *RateShopResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []RateShopResponseErrorsInner and assigns it to the Errors field.
func (o *RateShopResponse) SetErrors(v []RateShopResponseErrorsInner) {
	o.Errors = v
}

func (o RateShopResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RateShopResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.Rates) {
		toSerialize["rates"] = o.Rates
	}
	if !IsNil(o.ToAddress) {
		toSerialize["toAddress"] = o.ToAddress
	}
	if !IsNil(o.IsHazmat) {
		toSerialize["isHazmat"] = o.IsHazmat
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableRateShopResponse struct {
	value *RateShopResponse
	isSet bool
}

func (v NullableRateShopResponse) Get() *RateShopResponse {
	return v.value
}

func (v *NullableRateShopResponse) Set(val *RateShopResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRateShopResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRateShopResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateShopResponse(val *RateShopResponse) *NullableRateShopResponse {
	return &NullableRateShopResponse{value: val, isSet: true}
}

func (v NullableRateShopResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateShopResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


