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

// checks if the ShipmentDomestic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentDomestic{}

// ShipmentDomestic struct for ShipmentDomestic
type ShipmentDomestic struct {
	// This defines the label size of the Shipment, e.g., Shipping Label having Doc Size (4' X 6' or 8.5' X 11'). For NORATE carrier- only supported is `DOC_4X6`
	Size string `json:"size"`
	// This defines the type of the Shipment, e.g., Shipping Label.
	Type string `json:"type"`
	// This defines the type of the shipment which is printed. For example Shipping label prints in PDF form.
	Format *string `json:"format,omitempty"`
	// This defines the date of the Shipment in the format YYYY:MM:DD.
	DateOfShipment *string `json:"dateOfShipment,omitempty"`
	FromAddress ShipmentDomesticFromAddress `json:"fromAddress"`
	Parcel ShipmentDomesticParcel `json:"parcel"`
	//  A unique identifier associated with the Carrier account used by client users during shipment process.
	CarrierAccountId string `json:"carrierAccountId"`
	// >-Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection. It has categories like Package, Envelopes, Paks, Boxes, Tube, defined per specific carrier and used in abbreviated form, e.g., FRPKG, LGENV, TUBE,PKG.
	ParcelType string `json:"parcelType"`
	// >-Parcel Id is optional and required to be given in case of branded parcels which have brandedDimension and/or brandedWeight. If parcel has brandedDimension, in that case user need not to pass dimensionUnit and dimension details(length, width and height) in the parcel object. And if brandedWeight is also available for the parcel then in that case weightUnit and weight need not to be passed  in parcel object
	ParcelId *string `json:"parcelId,omitempty"`
	// >-A unique identifier given to the carrier-specific service. This is required for creating a shipment, while it is optional for rating a parcel.
	ServiceId string `json:"serviceId"`
	//  This provides a carrier-service based special or extra sevice.
	SpecialServices []SpecialService `json:"specialServices,omitempty"`
	ShipmentOptions *ShipmentOptionsV2 `json:"shipmentOptions,omitempty"`
	// Additional metadata that needs to be stored for this shipment can be added here.
	Metadata []ShipmentDomesticMetadataInner `json:"metadata,omitempty"`
	ToAddress ShipmentDomesticToAddress `json:"toAddress"`
}

type _ShipmentDomestic ShipmentDomestic

// NewShipmentDomestic instantiates a new ShipmentDomestic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentDomestic(size string, type_ string, fromAddress ShipmentDomesticFromAddress, parcel ShipmentDomesticParcel, carrierAccountId string, parcelType string, serviceId string, toAddress ShipmentDomesticToAddress) *ShipmentDomestic {
	this := ShipmentDomestic{}
	this.Size = size
	this.Type = type_
	this.FromAddress = fromAddress
	this.Parcel = parcel
	this.CarrierAccountId = carrierAccountId
	this.ParcelType = parcelType
	this.ServiceId = serviceId
	this.ToAddress = toAddress
	return &this
}

// NewShipmentDomesticWithDefaults instantiates a new ShipmentDomestic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentDomesticWithDefaults() *ShipmentDomestic {
	this := ShipmentDomestic{}
	return &this
}

// GetSize returns the Size field value
func (o *ShipmentDomestic) GetSize() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetSizeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *ShipmentDomestic) SetSize(v string) {
	o.Size = v
}

// GetType returns the Type field value
func (o *ShipmentDomestic) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ShipmentDomestic) SetType(v string) {
	o.Type = v
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *ShipmentDomestic) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *ShipmentDomestic) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *ShipmentDomestic) SetFormat(v string) {
	o.Format = &v
}

// GetDateOfShipment returns the DateOfShipment field value if set, zero value otherwise.
func (o *ShipmentDomestic) GetDateOfShipment() string {
	if o == nil || IsNil(o.DateOfShipment) {
		var ret string
		return ret
	}
	return *o.DateOfShipment
}

// GetDateOfShipmentOk returns a tuple with the DateOfShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetDateOfShipmentOk() (*string, bool) {
	if o == nil || IsNil(o.DateOfShipment) {
		return nil, false
	}
	return o.DateOfShipment, true
}

// HasDateOfShipment returns a boolean if a field has been set.
func (o *ShipmentDomestic) HasDateOfShipment() bool {
	if o != nil && !IsNil(o.DateOfShipment) {
		return true
	}

	return false
}

// SetDateOfShipment gets a reference to the given string and assigns it to the DateOfShipment field.
func (o *ShipmentDomestic) SetDateOfShipment(v string) {
	o.DateOfShipment = &v
}

// GetFromAddress returns the FromAddress field value
func (o *ShipmentDomestic) GetFromAddress() ShipmentDomesticFromAddress {
	if o == nil {
		var ret ShipmentDomesticFromAddress
		return ret
	}

	return o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetFromAddressOk() (*ShipmentDomesticFromAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromAddress, true
}

// SetFromAddress sets field value
func (o *ShipmentDomestic) SetFromAddress(v ShipmentDomesticFromAddress) {
	o.FromAddress = v
}

// GetParcel returns the Parcel field value
func (o *ShipmentDomestic) GetParcel() ShipmentDomesticParcel {
	if o == nil {
		var ret ShipmentDomesticParcel
		return ret
	}

	return o.Parcel
}

// GetParcelOk returns a tuple with the Parcel field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetParcelOk() (*ShipmentDomesticParcel, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parcel, true
}

// SetParcel sets field value
func (o *ShipmentDomestic) SetParcel(v ShipmentDomesticParcel) {
	o.Parcel = v
}

// GetCarrierAccountId returns the CarrierAccountId field value
func (o *ShipmentDomestic) GetCarrierAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CarrierAccountId, true
}

// SetCarrierAccountId sets field value
func (o *ShipmentDomestic) SetCarrierAccountId(v string) {
	o.CarrierAccountId = v
}

// GetParcelType returns the ParcelType field value
func (o *ShipmentDomestic) GetParcelType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ParcelType
}

// GetParcelTypeOk returns a tuple with the ParcelType field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetParcelTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParcelType, true
}

// SetParcelType sets field value
func (o *ShipmentDomestic) SetParcelType(v string) {
	o.ParcelType = v
}

// GetParcelId returns the ParcelId field value if set, zero value otherwise.
func (o *ShipmentDomestic) GetParcelId() string {
	if o == nil || IsNil(o.ParcelId) {
		var ret string
		return ret
	}
	return *o.ParcelId
}

// GetParcelIdOk returns a tuple with the ParcelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetParcelIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParcelId) {
		return nil, false
	}
	return o.ParcelId, true
}

// HasParcelId returns a boolean if a field has been set.
func (o *ShipmentDomestic) HasParcelId() bool {
	if o != nil && !IsNil(o.ParcelId) {
		return true
	}

	return false
}

// SetParcelId gets a reference to the given string and assigns it to the ParcelId field.
func (o *ShipmentDomestic) SetParcelId(v string) {
	o.ParcelId = &v
}

// GetServiceId returns the ServiceId field value
func (o *ShipmentDomestic) GetServiceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceId
}

// GetServiceIdOk returns a tuple with the ServiceId field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetServiceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceId, true
}

// SetServiceId sets field value
func (o *ShipmentDomestic) SetServiceId(v string) {
	o.ServiceId = v
}

// GetSpecialServices returns the SpecialServices field value if set, zero value otherwise.
func (o *ShipmentDomestic) GetSpecialServices() []SpecialService {
	if o == nil || IsNil(o.SpecialServices) {
		var ret []SpecialService
		return ret
	}
	return o.SpecialServices
}

// GetSpecialServicesOk returns a tuple with the SpecialServices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetSpecialServicesOk() ([]SpecialService, bool) {
	if o == nil || IsNil(o.SpecialServices) {
		return nil, false
	}
	return o.SpecialServices, true
}

// HasSpecialServices returns a boolean if a field has been set.
func (o *ShipmentDomestic) HasSpecialServices() bool {
	if o != nil && !IsNil(o.SpecialServices) {
		return true
	}

	return false
}

// SetSpecialServices gets a reference to the given []SpecialService and assigns it to the SpecialServices field.
func (o *ShipmentDomestic) SetSpecialServices(v []SpecialService) {
	o.SpecialServices = v
}

// GetShipmentOptions returns the ShipmentOptions field value if set, zero value otherwise.
func (o *ShipmentDomestic) GetShipmentOptions() ShipmentOptionsV2 {
	if o == nil || IsNil(o.ShipmentOptions) {
		var ret ShipmentOptionsV2
		return ret
	}
	return *o.ShipmentOptions
}

// GetShipmentOptionsOk returns a tuple with the ShipmentOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetShipmentOptionsOk() (*ShipmentOptionsV2, bool) {
	if o == nil || IsNil(o.ShipmentOptions) {
		return nil, false
	}
	return o.ShipmentOptions, true
}

// HasShipmentOptions returns a boolean if a field has been set.
func (o *ShipmentDomestic) HasShipmentOptions() bool {
	if o != nil && !IsNil(o.ShipmentOptions) {
		return true
	}

	return false
}

// SetShipmentOptions gets a reference to the given ShipmentOptionsV2 and assigns it to the ShipmentOptions field.
func (o *ShipmentDomestic) SetShipmentOptions(v ShipmentOptionsV2) {
	o.ShipmentOptions = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *ShipmentDomestic) GetMetadata() []ShipmentDomesticMetadataInner {
	if o == nil || IsNil(o.Metadata) {
		var ret []ShipmentDomesticMetadataInner
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetMetadataOk() ([]ShipmentDomesticMetadataInner, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *ShipmentDomestic) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []ShipmentDomesticMetadataInner and assigns it to the Metadata field.
func (o *ShipmentDomestic) SetMetadata(v []ShipmentDomesticMetadataInner) {
	o.Metadata = v
}

// GetToAddress returns the ToAddress field value
func (o *ShipmentDomestic) GetToAddress() ShipmentDomesticToAddress {
	if o == nil {
		var ret ShipmentDomesticToAddress
		return ret
	}

	return o.ToAddress
}

// GetToAddressOk returns a tuple with the ToAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentDomestic) GetToAddressOk() (*ShipmentDomesticToAddress, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToAddress, true
}

// SetToAddress sets field value
func (o *ShipmentDomestic) SetToAddress(v ShipmentDomesticToAddress) {
	o.ToAddress = v
}

func (o ShipmentDomestic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentDomestic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["size"] = o.Size
	toSerialize["type"] = o.Type
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.DateOfShipment) {
		toSerialize["dateOfShipment"] = o.DateOfShipment
	}
	toSerialize["fromAddress"] = o.FromAddress
	toSerialize["parcel"] = o.Parcel
	toSerialize["carrierAccountId"] = o.CarrierAccountId
	toSerialize["parcelType"] = o.ParcelType
	if !IsNil(o.ParcelId) {
		toSerialize["parcelId"] = o.ParcelId
	}
	toSerialize["serviceId"] = o.ServiceId
	if !IsNil(o.SpecialServices) {
		toSerialize["specialServices"] = o.SpecialServices
	}
	if !IsNil(o.ShipmentOptions) {
		toSerialize["shipmentOptions"] = o.ShipmentOptions
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	toSerialize["toAddress"] = o.ToAddress
	return toSerialize, nil
}

func (o *ShipmentDomestic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"size",
		"type",
		"fromAddress",
		"parcel",
		"carrierAccountId",
		"parcelType",
		"serviceId",
		"toAddress",
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

	varShipmentDomestic := _ShipmentDomestic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentDomestic)

	if err != nil {
		return err
	}

	*o = ShipmentDomestic(varShipmentDomestic)

	return err
}

type NullableShipmentDomestic struct {
	value *ShipmentDomestic
	isSet bool
}

func (v NullableShipmentDomestic) Get() *ShipmentDomestic {
	return v.value
}

func (v *NullableShipmentDomestic) Set(val *ShipmentDomestic) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentDomestic) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentDomestic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentDomestic(val *ShipmentDomestic) *NullableShipmentDomestic {
	return &NullableShipmentDomestic{value: val, isSet: true}
}

func (v NullableShipmentDomestic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentDomestic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


