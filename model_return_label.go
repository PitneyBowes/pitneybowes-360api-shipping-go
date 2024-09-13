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

// checks if the ReturnLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReturnLabel{}

// ReturnLabel struct for ReturnLabel
type ReturnLabel struct {
	// This defines the label size of the Shipment, e.g., Shipping Label having Doc Size (8' X 11').
	Size *string `json:"size,omitempty"`
	// This defines the type of the Shipment, e.g., Shipping Label.
	Type *string `json:"type,omitempty"`
	FromAddress *ReturnLabelFromAddress `json:"fromAddress,omitempty"`
	Parcel *ShipmentDomesticParcel `json:"parcel,omitempty"`
	// The unique identifier associated with the Carrier account used by client users during shipment process.
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	// >-Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. It has categories like Package, Envelopes, Paks, Boxes, Tube, defined per specific carrier and used in abbreviated form, e.g., FRPKG, LGENV, TUBE,PKG.
	ParcelType *string `json:"parcelType,omitempty"`
	// A unique identifier associated with the Parcel.
	ParcelId *string `json:"parcelId,omitempty"`
	// >-A unique identifier given to the carrier-specific service. This is required for creating a shipment, while it is optional for rating a parcel.
	ServiceId *string `json:"serviceId,omitempty"`
	SpecialServices []ReturnLabelSpecialServicesInner `json:"specialServices,omitempty"`
	ShipmentOptions *ShipmentOptionsV2 `json:"shipmentOptions,omitempty"`
	// Additional metadata that needs to be stored for this shipment can be added here. For now, `costAccountName` is supported.
	Metadata []ShipmentInternationalMetadataInner `json:"metadata,omitempty"`
	ToAddress *ReturnLabelToAddress `json:"toAddress,omitempty"`
}

// NewReturnLabel instantiates a new ReturnLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReturnLabel() *ReturnLabel {
	this := ReturnLabel{}
	return &this
}

// NewReturnLabelWithDefaults instantiates a new ReturnLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReturnLabelWithDefaults() *ReturnLabel {
	this := ReturnLabel{}
	return &this
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *ReturnLabel) GetSize() string {
	if o == nil || IsNil(o.Size) {
		var ret string
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetSizeOk() (*string, bool) {
	if o == nil || IsNil(o.Size) {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *ReturnLabel) HasSize() bool {
	if o != nil && !IsNil(o.Size) {
		return true
	}

	return false
}

// SetSize gets a reference to the given string and assigns it to the Size field.
func (o *ReturnLabel) SetSize(v string) {
	o.Size = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ReturnLabel) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ReturnLabel) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ReturnLabel) SetType(v string) {
	o.Type = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *ReturnLabel) GetFromAddress() ReturnLabelFromAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret ReturnLabelFromAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetFromAddressOk() (*ReturnLabelFromAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *ReturnLabel) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given ReturnLabelFromAddress and assigns it to the FromAddress field.
func (o *ReturnLabel) SetFromAddress(v ReturnLabelFromAddress) {
	o.FromAddress = &v
}

// GetParcel returns the Parcel field value if set, zero value otherwise.
func (o *ReturnLabel) GetParcel() ShipmentDomesticParcel {
	if o == nil || IsNil(o.Parcel) {
		var ret ShipmentDomesticParcel
		return ret
	}
	return *o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetParcelOk() (*ShipmentDomesticParcel, bool) {
	if o == nil || IsNil(o.Parcel) {
		return nil, false
	}
	return o.Parcel, true
}

// HasParcel returns a boolean if a field has been set.
func (o *ReturnLabel) HasParcel() bool {
	if o != nil && !IsNil(o.Parcel) {
		return true
	}

	return false
}

// SetParcel gets a reference to the given ShipmentDomesticParcel and assigns it to the Parcel field.
func (o *ReturnLabel) SetParcel(v ShipmentDomesticParcel) {
	o.Parcel = &v
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *ReturnLabel) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *ReturnLabel) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *ReturnLabel) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetParcelType returns the ParcelType field value if set, zero value otherwise.
func (o *ReturnLabel) GetParcelType() string {
	if o == nil || IsNil(o.ParcelType) {
		var ret string
		return ret
	}
	return *o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetParcelTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelType) {
		return nil, false
	}
	return o.ParcelType, true
}

// HasParcelType returns a boolean if a field has been set.
func (o *ReturnLabel) HasParcelType() bool {
	if o != nil && !IsNil(o.ParcelType) {
		return true
	}

	return false
}

// SetParcelType gets a reference to the given string and assigns it to the ParcelType field.
func (o *ReturnLabel) SetParcelType(v string) {
	o.ParcelType = &v
}

// GetParcelId returns the ParcelId field value if set, zero value otherwise.
func (o *ReturnLabel) GetParcelId() string {
	if o == nil || IsNil(o.ParcelId) {
		var ret string
		return ret
	}
	return *o.ParcelId
}

// GetParcelIdOk returns a tuple with the ParcelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetParcelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelId) {
		return nil, false
	}
	return o.ParcelId, true
}

// HasParcelId returns a boolean if a field has been set.
func (o *ReturnLabel) HasParcelId() bool {
	if o != nil && !IsNil(o.ParcelId) {
		return true
	}

	return false
}

// SetParcelId gets a reference to the given string and assigns it to the ParcelId field.
func (o *ReturnLabel) SetParcelId(v string) {
	o.ParcelId = &v
}

// GetServiceId returns the ServiceId field value if set, zero value otherwise.
func (o *ReturnLabel) GetServiceId() string {
	if o == nil || IsNil(o.ServiceId) {
		var ret string
		return ret
	}
	return *o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetServiceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceId) {
		return nil, false
	}
	return o.ServiceId, true
}

// HasServiceId returns a boolean if a field has been set.
func (o *ReturnLabel) HasServiceId() bool {
	if o != nil && !IsNil(o.ServiceId) {
		return true
	}

	return false
}

// SetServiceId gets a reference to the given string and assigns it to the ServiceId field.
func (o *ReturnLabel) SetServiceId(v string) {
	o.ServiceId = &v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *ReturnLabel) GetSpecialServices() []ReturnLabelSpecialServicesInner {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []ReturnLabelSpecialServicesInner
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetSpecialServicesOk() ([]ReturnLabelSpecialServicesInner, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *ReturnLabel) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []ReturnLabelSpecialServicesInner and assigns it to the SpecialServices field.
func (o *ReturnLabel) SetSpecialServices(v []ReturnLabelSpecialServicesInner) {
	o.SpecialServices = v
}

// GetShipmentOptions returns the ShipmentOptions field value if set, zero value otherwise.
func (o *ReturnLabel) GetShipmentOptions() ShipmentOptionsV2 {
	if o == nil || IsNil(o.ShipmentOptions) {
		var ret ShipmentOptionsV2
		return ret
	}
	return *o.ShipmentOptions
}

// GetShipmentOptionsOk returns a tuple with the ShipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetShipmentOptionsOk() (*ShipmentOptionsV2, bool) {
	if o == nil || IsNil(o.ShipmentOptions) {
		return nil, false
	}
	return o.ShipmentOptions, true
}

// HasShipmentOptions returns a boolean if a field has been set.
func (o *ReturnLabel) HasShipmentOptions() bool {
	if o != nil && !IsNil(o.ShipmentOptions) {
		return true
	}

	return false
}

// SetShipmentOptions gets a reference to the given ShipmentOptionsV2 and assigns it to the ShipmentOptions field.
func (o *ReturnLabel) SetShipmentOptions(v ShipmentOptionsV2) {
	o.ShipmentOptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ReturnLabel) GetMetadata() []ShipmentInternationalMetadataInner {
	if o == nil || IsNil(o.Metadata) {
		var ret []ShipmentInternationalMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetMetadataOk() ([]ShipmentInternationalMetadataInner, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ReturnLabel) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ShipmentInternationalMetadataInner and assigns it to the Metadata field.
func (o *ReturnLabel) SetMetadata(v []ShipmentInternationalMetadataInner) {
	o.Metadata = v
}

// GetToAddress returns the ToAddress field value if set, zero value otherwise.
func (o *ReturnLabel) GetToAddress() ReturnLabelToAddress {
	if o == nil || IsNil(o.ToAddress) {
		var ret ReturnLabelToAddress
		return ret
	}
	return *o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReturnLabel) GetToAddressOk() (*ReturnLabelToAddress, bool) {
	if o == nil || IsNil(o.ToAddress) {
		return nil, false
	}
	return o.ToAddress, true
}

// HasToAddress returns a boolean if a field has been set.
func (o *ReturnLabel) HasToAddress() bool {
	if o != nil && !IsNil(o.ToAddress) {
		return true
	}

	return false
}

// SetToAddress gets a reference to the given ReturnLabelToAddress and assigns it to the ToAddress field.
func (o *ReturnLabel) SetToAddress(v ReturnLabelToAddress) {
	o.ToAddress = &v
}

func (o ReturnLabel) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReturnLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Size) {
		toSerialize["size"] = o.Size
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.Parcel) {
		toSerialize["parcel"] = o.Parcel
	}
	if !IsNil(o.CarrierAccountId) {
		toSerialize["carrierAccountId"] = o.CarrierAccountId
	}
	if !IsNil(o.ParcelType) {
		toSerialize["parcelType"] = o.ParcelType
	}
	if !IsNil(o.ParcelId) {
		toSerialize["parcelId"] = o.ParcelId
	}
	if !IsNil(o.ServiceId) {
		toSerialize["serviceId"] = o.ServiceId
	}
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	if !IsNil(o.ShipmentOptions) {
		toSerialize["shipmentOptions"] = o.ShipmentOptions
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.ToAddress) {
		toSerialize["toAddress"] = o.ToAddress
	}
	return toSerialize, nil
}

type NullableReturnLabel struct {
	value *ReturnLabel
	isSet bool
}

func (v NullableReturnLabel) Get() *ReturnLabel {
	return v.value
}

func (v *NullableReturnLabel) Set(val *ReturnLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableReturnLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableReturnLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReturnLabel(val *ReturnLabel) *NullableReturnLabel {
	return &NullableReturnLabel{value: val, isSet: true}
}

func (v NullableReturnLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReturnLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


