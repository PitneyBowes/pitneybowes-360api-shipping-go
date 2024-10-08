/*
Shipping APIs

### Introduction  The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.   Some of the key API operations available in the Shipping API includes: ### Shipment API  | Operation      | Description | | ----------- | ----------- |  | Get Carriers    | This operation fetches all onboarded carriers. Typically, user will use this service to get list of onboarded carriers and supported properties for those carriers.  |  | Get Countries | This operation fetches list of supported destination countries for a provided carrier and origin country.  |  | Get Services | This operation fetches a list of supported services for a carrier with respect to specific origin and destination country. |  | Get ParcelTypes| This operation fetches ParcelTypes based on carrier, origin and destination country. |  | Get Special Services| This operation fetches Special Services for a given carrier, service, origin and destination country. |  | Get Carrier Accounts| This operation retrieves onboarded Carriers with their Carrier Account Ids which uniquely identify multiple accounts of same carrier.  |  | Rate Shop and Get Single Rate| This API contains 2 operations, rate shop and single rate. Rate shop will fetch rates for all carrier services based on the given addresses (From and To), weight, and dimension for given parcelType. Single rate will get rate for specific service and special service (if requested) based on the given addresses (From and To), weight, and dimension, parcelType and serviceId with or without specialServices. Single rate will be used mainly to a rate a shipment before creating shipment.  |  | Create Shipment| This operation creates a new Shipment or Shipment Label. This is for both Domestic and International. | | Get All Shipments| This operation fetches all created Shipments. |  | Get Shipment by Id| Retrieves single shipment using Shipment Id. |  | Reprint Shipment| This operation reprints Shipment by the shipmentId. It retrieves an existing shipping label to reprint. The API sends the shipmentId returned by the original Created Shipment request. Use this only if the shipping label in the Create Shipment response was spoilt or lost. |  | Cancel Shipment| This operation cancels previously created shipment. |  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipping

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SpecialService type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SpecialService{}

// SpecialService struct for SpecialService
type SpecialService struct {
	// >- The parameters to set for the special service, such as an insurance value or a receipt-number format. This is required if the special service requires input parameters. If a special service does not require input parameters, you can either leave out the array or pass an empty array.
	InputParameters []Parameter `json:"inputParameters,omitempty"`
	// A unique identifier associate to the special service, which is to be applied.
	SpecialserviceId string `json:"specialserviceId"`
}

type _SpecialService SpecialService

// NewSpecialService instantiates a new SpecialService object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpecialService(specialserviceId string) *SpecialService {
	this := SpecialService{}
	this.SpecialserviceId = specialserviceId
	return &this
}

// NewSpecialServiceWithDefaults instantiates a new SpecialService object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpecialServiceWithDefaults() *SpecialService {
	this := SpecialService{}
	return &this
}

// GetInputParameters returns the InputParameters field value if set, zero value otherwise.
func (o *SpecialService) GetInputParameters() []Parameter {
	if o == nil || IsNil(o.InputParameters) {
		var ret []Parameter
		return ret
	}
	return o.InputParameters
}

// GetInputParametersOk returns a tuple with the InputParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpecialService) GetInputParametersOk() ([]Parameter, bool) {
	if o == nil || IsNil(o.InputParameters) {
		return nil, false
	}
	return o.InputParameters, true
}

// HasInputParameters returns a boolean if a field has been set.
func (o *SpecialService) HasInputParameters() bool {
	if o != nil && !IsNil(o.InputParameters) {
		return true
	}

	return false
}

// SetInputParameters gets a reference to the given []Parameter and assigns it to the InputParameters field.
func (o *SpecialService) SetInputParameters(v []Parameter) {
	o.InputParameters = v
}

// GetSpecialserviceId returns the SpecialserviceId field value
func (o *SpecialService) GetSpecialserviceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpecialserviceId
}

// GetSpecialserviceIdOk returns a tuple with the SpecialserviceId field value
// and a boolean to check if the value has been set.
func (o *SpecialService) GetSpecialserviceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SpecialserviceId, true
}

// SetSpecialserviceId sets field value
func (o *SpecialService) SetSpecialserviceId(v string) {
	o.SpecialserviceId = v
}

func (o SpecialService) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SpecialService) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InputParameters) {
		toSerialize["inputParameters"] = o.InputParameters
	}
	toSerialize["specialserviceId"] = o.SpecialserviceId
	return toSerialize, nil
}

func (o *SpecialService) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"specialserviceId",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSpecialService := _SpecialService{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSpecialService)

	if err != nil {
		return err
	}

	*o = SpecialService(varSpecialService)

	return err
}

type NullableSpecialService struct {
	value *SpecialService
	isSet bool
}

func (v NullableSpecialService) Get() *SpecialService {
	return v.value
}

func (v *NullableSpecialService) Set(val *SpecialService) {
	v.value = val
	v.isSet = true
}

func (v NullableSpecialService) IsSet() bool {
	return v.isSet
}

func (v *NullableSpecialService) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpecialService(val *SpecialService) *NullableSpecialService {
	return &NullableSpecialService{value: val, isSet: true}
}

func (v NullableSpecialService) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpecialService) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


