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

// checks if the CreateBulkShipmentsERRLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBulkShipmentsERRLabel{}

// CreateBulkShipmentsERRLabel This ShipmentBatch contains the schema information for ERR Label.
type CreateBulkShipmentsERRLabel struct {
	// Name of the of ERR Batch which consists of multiple shipments (shipments in bulk) for Label printing, e.g. ERR-Bulk05.
	Name string `json:"name"`
	// Indicates the name of the group of batches, which consists of multiple Batch groups.
	GroupName *string `json:"groupName,omitempty"`
	// This indicates the label size of the Bulk Shipment when it gets printed,i.e., DocSize. This has two options 8' X 11' or 4' X 6'.
	Size string `json:"size"`
	// Indicates the type of the Batch Shipment, e.g., Shipping Label.
	Type string `json:"type"`
	// Defines the type of the shipment which is printed, e.g., Shipping label gets printed in PDF form.
	Format *string `json:"format,omitempty"`
	// A unique identifier associated with the Carrier account used by client users during shipment process. Default CarrierAccountID for this batch will be user's registered USPS account. User can override this value by defining it at Shipment level.
	CarrierAccountId string `json:"carrierAccountId"`
	// A unique identifier given to the carrier-specific service. User can override this value by defining it at Shipment level.
	ServiceId string `json:"serviceId"`
	// Parcel Type is required for creating a shipment while rating a parcel. And it varies as per USPS selected services, e.g. FRPKG, LGENV, TUBE, and PKG. User can override this value by defining it at Shipment level.
	ParcelType string `json:"parcelType"`
	// A unique identifier given to the parcel or package corresponding to the selected service. This is optional field, but is used in few cases. Examples include BLM10, B1095, MT1098, etc. User can override this value by defining it at Shipment level.
	ParcelID *string `json:"parcelID,omitempty"`
	SpecialServices []SpecialServiceERRInner `json:"specialServices,omitempty"`
	Shipments []ShipmentERR `json:"shipments"`
}

type _CreateBulkShipmentsERRLabel CreateBulkShipmentsERRLabel

// NewCreateBulkShipmentsERRLabel instantiates a new CreateBulkShipmentsERRLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBulkShipmentsERRLabel(name string, size string, type_ string, carrierAccountId string, serviceId string, parcelType string, shipments []ShipmentERR) *CreateBulkShipmentsERRLabel {
	this := CreateBulkShipmentsERRLabel{}
	this.Name = name
	this.Size = size
	this.Type = type_
	this.CarrierAccountId = carrierAccountId
	this.ServiceId = serviceId
	this.ParcelType = parcelType
	this.Shipments = shipments
	return &this
}

// NewCreateBulkShipmentsERRLabelWithDefaults instantiates a new CreateBulkShipmentsERRLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBulkShipmentsERRLabelWithDefaults() *CreateBulkShipmentsERRLabel {
	this := CreateBulkShipmentsERRLabel{}
	return &this
}

// GetName returns the Name field value
func (o *CreateBulkShipmentsERRLabel) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateBulkShipmentsERRLabel) SetName(v string) {
	o.Name = v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRLabel) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRLabel) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *CreateBulkShipmentsERRLabel) SetGroupName(v string) {
	o.GroupName = &v
}

// GetSize returns the Size field value
func (o *CreateBulkShipmentsERRLabel) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *CreateBulkShipmentsERRLabel) SetSize(v string) {
	o.Size = v
}

// GetType returns the Type field value
func (o *CreateBulkShipmentsERRLabel) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CreateBulkShipmentsERRLabel) SetType(v string) {
	o.Type = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRLabel) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRLabel) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *CreateBulkShipmentsERRLabel) SetFormat(v string) {
	o.Format = &v
}

// GetCarrierAccountId returns the CarrierAccountId field value
func (o *CreateBulkShipmentsERRLabel) GetCarrierAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierAccountId, true
}

// SetCarrierAccountId sets field value
func (o *CreateBulkShipmentsERRLabel) SetCarrierAccountId(v string) {
	o.CarrierAccountId = v
}

// GetServiceId returns the ServiceId field value
func (o *CreateBulkShipmentsERRLabel) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *CreateBulkShipmentsERRLabel) SetServiceId(v string) {
	o.ServiceId = v
}

// GetParcelType returns the ParcelType field value
func (o *CreateBulkShipmentsERRLabel) GetParcelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetParcelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParcelType, true
}

// SetParcelType sets field value
func (o *CreateBulkShipmentsERRLabel) SetParcelType(v string) {
	o.ParcelType = v
}

// GetParcelID returns the ParcelID field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRLabel) GetParcelID() string {
	if o == nil || IsNil(o.ParcelID) {
		var ret string
		return ret
	}
	return *o.ParcelID
}

// GetParcelIDOk returns a tuple with the ParcelID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetParcelIDOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelID) {
		return nil, false
	}
	return o.ParcelID, true
}

// HasParcelID returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRLabel) HasParcelID() bool {
	if o != nil && !IsNil(o.ParcelID) {
		return true
	}

	return false
}

// SetParcelID gets a reference to the given string and assigns it to the ParcelID field.
func (o *CreateBulkShipmentsERRLabel) SetParcelID(v string) {
	o.ParcelID = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *CreateBulkShipmentsERRLabel) GetSpecialServices() []SpecialServiceERRInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialServiceERRInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetSpecialServicesOk() ([]SpecialServiceERRInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *CreateBulkShipmentsERRLabel) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialServiceERRInner and assigns it to the SpecialServices field.
func (o *CreateBulkShipmentsERRLabel) SetSpecialServices(v []SpecialServiceERRInner) {
	o.SpecialServices = v
}

// GetShipments returns the Shipments field value
func (o *CreateBulkShipmentsERRLabel) GetShipments() []ShipmentERR {
	if o == nil {
		var ret []ShipmentERR
		return ret
	}

	return o.Shipments
}

// GetShipmentsOk returns a tuple with the Shipments field value
// and a boolean to check if the value has been set.
func (o *CreateBulkShipmentsERRLabel) GetShipmentsOk() ([]ShipmentERR, bool) {
	if o == nil {
		return nil, false
	}
	return o.Shipments, true
}

// SetShipments sets field value
func (o *CreateBulkShipmentsERRLabel) SetShipments(v []ShipmentERR) {
	o.Shipments = v
}

func (o CreateBulkShipmentsERRLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBulkShipmentsERRLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	toSerialize["size"] = o.Size
	toSerialize["type"] = o.Type
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	toSerialize["carrierAccountId"] = o.CarrierAccountId
	toSerialize["serviceId"] = o.ServiceId
	toSerialize["parcelType"] = o.ParcelType
	if !IsNil(o.ParcelID) {
		toSerialize["parcelID"] = o.ParcelID
	}
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	toSerialize["shipments"] = o.Shipments
	return toSerialize, nil
}

func (o *CreateBulkShipmentsERRLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"size",
		"type",
		"carrierAccountId",
		"serviceId",
		"parcelType",
		"shipments",
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

	varCreateBulkShipmentsERRLabel := _CreateBulkShipmentsERRLabel{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateBulkShipmentsERRLabel)

	if err != nil {
		return err
	}

	*o = CreateBulkShipmentsERRLabel(varCreateBulkShipmentsERRLabel)

	return err
}

type NullableCreateBulkShipmentsERRLabel struct {
	value *CreateBulkShipmentsERRLabel
	isSet bool
}

func (v NullableCreateBulkShipmentsERRLabel) Get() *CreateBulkShipmentsERRLabel {
	return v.value
}

func (v *NullableCreateBulkShipmentsERRLabel) Set(val *CreateBulkShipmentsERRLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBulkShipmentsERRLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBulkShipmentsERRLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBulkShipmentsERRLabel(val *CreateBulkShipmentsERRLabel) *NullableCreateBulkShipmentsERRLabel {
	return &NullableCreateBulkShipmentsERRLabel{value: val, isSet: true}
}

func (v NullableCreateBulkShipmentsERRLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBulkShipmentsERRLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


