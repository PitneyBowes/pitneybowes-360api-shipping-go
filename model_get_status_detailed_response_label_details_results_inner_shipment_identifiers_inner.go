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

// checks if the GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner{}

// GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner struct for GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner
type GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner struct {
	//  This indicates the identifier at DB level mapped with Shipment ID.
	Id *string `json:"id,omitempty"`
	//  Shipment ID is a unique identifier for an individual shipment.
	ShipmentId *string `json:"shipmentId,omitempty"`
	//  This is a user-defined value provided by users just for their reference. This is for mapping purpose against each shipment.
	ExternalId *string `json:"externalId,omitempty"`
}

// NewGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner instantiates a new GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner() *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner {
	this := GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner{}
	return &this
}

// NewGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInnerWithDefaults instantiates a new GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInnerWithDefaults() *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner {
	this := GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) SetId(v string) {
	o.Id = &v
}

// GetShipmentId returns the ShipmentId field value if set, zero value otherwise.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) GetShipmentId() string {
	if o == nil || IsNil(o.ShipmentId) {
		var ret string
		return ret
	}
	return *o.ShipmentId
}

// GetShipmentIdOk returns a tuple with the ShipmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) GetShipmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ShipmentId) {
		return nil, false
	}
	return o.ShipmentId, true
}

// HasShipmentId returns a boolean if a field has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) HasShipmentId() bool {
	if o != nil && !IsNil(o.ShipmentId) {
		return true
	}

	return false
}

// SetShipmentId gets a reference to the given string and assigns it to the ShipmentId field.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) SetShipmentId(v string) {
	o.ShipmentId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.ShipmentId) {
		toSerialize["shipmentId"] = o.ShipmentId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	return toSerialize, nil
}

type NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner struct {
	value *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner
	isSet bool
}

func (v NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) Get() *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner {
	return v.value
}

func (v *NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) Set(val *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner(val *GetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) *NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner {
	return &NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner{value: val, isSet: true}
}

func (v NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatusDetailedResponseLabelDetailsResultsInnerShipmentIdentifiersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


