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

// checks if the SchedulePickupDHLEXPResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchedulePickupDHLEXPResponse{}

// SchedulePickupDHLEXPResponse struct for SchedulePickupDHLEXPResponse
type SchedulePickupDHLEXPResponse struct {
	// It specifies the location from where packages would be collected.
	PackageLocation *string `json:"packageLocation,omitempty"`
	// It displays the unique confirmation number of the pickup
	PickupConfirmationNumber *string `json:"pickupConfirmationNumber,omitempty"`
	// It displays the unique pickup Id which can be further used to get scheduled PDF and cancel pdf if required.
	PickupId *string `json:"pickupId,omitempty"`
	// It dispays the carrier
	Carrier *string `json:"carrier,omitempty"`
	// It displays the carrier acount id which is used to create the shipment
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	PickupAddress *SchedulePickupDHLEXPResponsePickupAddress `json:"pickupAddress,omitempty"`
	// It indicates the shipmentIds for which pickup is scheduled.
	ShipmentIds []string `json:"shipmentIds,omitempty"`
	// It displays the package details provided in the request.
	PickupSummary []SchedulePickupDHLEXPResponsePickupSummaryInner `json:"pickupSummary,omitempty"`
	// It displays additional comments or remarks provided in the request, it would be printed on the scheduled pickup document
	Additionalnotes *string `json:"additionalnotes,omitempty"`
	// It displays any reference information provided in the request.
	Reference *string `json:"reference,omitempty"`
	// It displays the scheduled pickup date and time (in UTC)
	PickupDateTime *string `json:"pickupDateTime,omitempty"`
	// It displays the total package weight.
	PickupTotalWeight *float32 `json:"pickupTotalWeight,omitempty"`
	// It displays the weight unit.
	PickupTotalWeightUnit *string `json:"pickupTotalWeightUnit,omitempty"`
	PickupOptions *SchedulePickupDHLEXPResponsePickupOptions `json:"pickupOptions,omitempty"`
}

// NewSchedulePickupDHLEXPResponse instantiates a new SchedulePickupDHLEXPResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchedulePickupDHLEXPResponse() *SchedulePickupDHLEXPResponse {
	this := SchedulePickupDHLEXPResponse{}
	return &this
}

// NewSchedulePickupDHLEXPResponseWithDefaults instantiates a new SchedulePickupDHLEXPResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchedulePickupDHLEXPResponseWithDefaults() *SchedulePickupDHLEXPResponse {
	this := SchedulePickupDHLEXPResponse{}
	return &this
}

// GetPackageLocation returns the PackageLocation field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPackageLocation() string {
	if o == nil || IsNil(o.PackageLocation) {
		var ret string
		return ret
	}
	return *o.PackageLocation
}

// GetPackageLocationOk returns a tuple with the PackageLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPackageLocationOk() (*string, bool) {
	if o == nil || IsNil(o.PackageLocation) {
		return nil, false
	}
	return o.PackageLocation, true
}

// HasPackageLocation returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPackageLocation() bool {
	if o != nil && !IsNil(o.PackageLocation) {
		return true
	}

	return false
}

// SetPackageLocation gets a reference to the given string and assigns it to the PackageLocation field.
func (o *SchedulePickupDHLEXPResponse) SetPackageLocation(v string) {
	o.PackageLocation = &v
}

// GetPickupConfirmationNumber returns the PickupConfirmationNumber field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupConfirmationNumber() string {
	if o == nil || IsNil(o.PickupConfirmationNumber) {
		var ret string
		return ret
	}
	return *o.PickupConfirmationNumber
}

// GetPickupConfirmationNumberOk returns a tuple with the PickupConfirmationNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupConfirmationNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PickupConfirmationNumber) {
		return nil, false
	}
	return o.PickupConfirmationNumber, true
}

// HasPickupConfirmationNumber returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupConfirmationNumber() bool {
	if o != nil && !IsNil(o.PickupConfirmationNumber) {
		return true
	}

	return false
}

// SetPickupConfirmationNumber gets a reference to the given string and assigns it to the PickupConfirmationNumber field.
func (o *SchedulePickupDHLEXPResponse) SetPickupConfirmationNumber(v string) {
	o.PickupConfirmationNumber = &v
}

// GetPickupId returns the PickupId field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupId() string {
	if o == nil || IsNil(o.PickupId) {
		var ret string
		return ret
	}
	return *o.PickupId
}

// GetPickupIdOk returns a tuple with the PickupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupIdOk() (*string, bool) {
	if o == nil || IsNil(o.PickupId) {
		return nil, false
	}
	return o.PickupId, true
}

// HasPickupId returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupId() bool {
	if o != nil && !IsNil(o.PickupId) {
		return true
	}

	return false
}

// SetPickupId gets a reference to the given string and assigns it to the PickupId field.
func (o *SchedulePickupDHLEXPResponse) SetPickupId(v string) {
	o.PickupId = &v
}

// GetCarrier returns the Carrier field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetCarrier() string {
	if o == nil || IsNil(o.Carrier) {
		var ret string
		return ret
	}
	return *o.Carrier
}

// GetCarrierOk returns a tuple with the Carrier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetCarrierOk() (*string, bool) {
	if o == nil || IsNil(o.Carrier) {
		return nil, false
	}
	return o.Carrier, true
}

// HasCarrier returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasCarrier() bool {
	if o != nil && !IsNil(o.Carrier) {
		return true
	}

	return false
}

// SetCarrier gets a reference to the given string and assigns it to the Carrier field.
func (o *SchedulePickupDHLEXPResponse) SetCarrier(v string) {
	o.Carrier = &v
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *SchedulePickupDHLEXPResponse) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetPickupAddress returns the PickupAddress field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupAddress() SchedulePickupDHLEXPResponsePickupAddress {
	if o == nil || IsNil(o.PickupAddress) {
		var ret SchedulePickupDHLEXPResponsePickupAddress
		return ret
	}
	return *o.PickupAddress
}

// GetPickupAddressOk returns a tuple with the PickupAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupAddressOk() (*SchedulePickupDHLEXPResponsePickupAddress, bool) {
	if o == nil || IsNil(o.PickupAddress) {
		return nil, false
	}
	return o.PickupAddress, true
}

// HasPickupAddress returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupAddress() bool {
	if o != nil && !IsNil(o.PickupAddress) {
		return true
	}

	return false
}

// SetPickupAddress gets a reference to the given SchedulePickupDHLEXPResponsePickupAddress and assigns it to the PickupAddress field.
func (o *SchedulePickupDHLEXPResponse) SetPickupAddress(v SchedulePickupDHLEXPResponsePickupAddress) {
	o.PickupAddress = &v
}

// GetShipmentIds returns the ShipmentIds field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetShipmentIds() []string {
	if o == nil || IsNil(o.ShipmentIds) {
		var ret []string
		return ret
	}
	return o.ShipmentIds
}

// GetShipmentIdsOk returns a tuple with the ShipmentIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetShipmentIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ShipmentIds) {
		return nil, false
	}
	return o.ShipmentIds, true
}

// HasShipmentIds returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasShipmentIds() bool {
	if o != nil && !IsNil(o.ShipmentIds) {
		return true
	}

	return false
}

// SetShipmentIds gets a reference to the given []string and assigns it to the ShipmentIds field.
func (o *SchedulePickupDHLEXPResponse) SetShipmentIds(v []string) {
	o.ShipmentIds = v
}

// GetPickupSummary returns the PickupSummary field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupSummary() []SchedulePickupDHLEXPResponsePickupSummaryInner {
	if o == nil || IsNil(o.PickupSummary) {
		var ret []SchedulePickupDHLEXPResponsePickupSummaryInner
		return ret
	}
	return o.PickupSummary
}

// GetPickupSummaryOk returns a tuple with the PickupSummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupSummaryOk() ([]SchedulePickupDHLEXPResponsePickupSummaryInner, bool) {
	if o == nil || IsNil(o.PickupSummary) {
		return nil, false
	}
	return o.PickupSummary, true
}

// HasPickupSummary returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupSummary() bool {
	if o != nil && !IsNil(o.PickupSummary) {
		return true
	}

	return false
}

// SetPickupSummary gets a reference to the given []SchedulePickupDHLEXPResponsePickupSummaryInner and assigns it to the PickupSummary field.
func (o *SchedulePickupDHLEXPResponse) SetPickupSummary(v []SchedulePickupDHLEXPResponsePickupSummaryInner) {
	o.PickupSummary = v
}

// GetAdditionalnotes returns the Additionalnotes field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetAdditionalnotes() string {
	if o == nil || IsNil(o.Additionalnotes) {
		var ret string
		return ret
	}
	return *o.Additionalnotes
}

// GetAdditionalnotesOk returns a tuple with the Additionalnotes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetAdditionalnotesOk() (*string, bool) {
	if o == nil || IsNil(o.Additionalnotes) {
		return nil, false
	}
	return o.Additionalnotes, true
}

// HasAdditionalnotes returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasAdditionalnotes() bool {
	if o != nil && !IsNil(o.Additionalnotes) {
		return true
	}

	return false
}

// SetAdditionalnotes gets a reference to the given string and assigns it to the Additionalnotes field.
func (o *SchedulePickupDHLEXPResponse) SetAdditionalnotes(v string) {
	o.Additionalnotes = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetReference() string {
	if o == nil || IsNil(o.Reference) {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Reference) {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasReference() bool {
	if o != nil && !IsNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *SchedulePickupDHLEXPResponse) SetReference(v string) {
	o.Reference = &v
}

// GetPickupDateTime returns the PickupDateTime field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupDateTime() string {
	if o == nil || IsNil(o.PickupDateTime) {
		var ret string
		return ret
	}
	return *o.PickupDateTime
}

// GetPickupDateTimeOk returns a tuple with the PickupDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupDateTimeOk() (*string, bool) {
	if o == nil || IsNil(o.PickupDateTime) {
		return nil, false
	}
	return o.PickupDateTime, true
}

// HasPickupDateTime returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupDateTime() bool {
	if o != nil && !IsNil(o.PickupDateTime) {
		return true
	}

	return false
}

// SetPickupDateTime gets a reference to the given string and assigns it to the PickupDateTime field.
func (o *SchedulePickupDHLEXPResponse) SetPickupDateTime(v string) {
	o.PickupDateTime = &v
}

// GetPickupTotalWeight returns the PickupTotalWeight field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeight() float32 {
	if o == nil || IsNil(o.PickupTotalWeight) {
		var ret float32
		return ret
	}
	return *o.PickupTotalWeight
}

// GetPickupTotalWeightOk returns a tuple with the PickupTotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.PickupTotalWeight) {
		return nil, false
	}
	return o.PickupTotalWeight, true
}

// HasPickupTotalWeight returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupTotalWeight() bool {
	if o != nil && !IsNil(o.PickupTotalWeight) {
		return true
	}

	return false
}

// SetPickupTotalWeight gets a reference to the given float32 and assigns it to the PickupTotalWeight field.
func (o *SchedulePickupDHLEXPResponse) SetPickupTotalWeight(v float32) {
	o.PickupTotalWeight = &v
}

// GetPickupTotalWeightUnit returns the PickupTotalWeightUnit field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeightUnit() string {
	if o == nil || IsNil(o.PickupTotalWeightUnit) {
		var ret string
		return ret
	}
	return *o.PickupTotalWeightUnit
}

// GetPickupTotalWeightUnitOk returns a tuple with the PickupTotalWeightUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupTotalWeightUnitOk() (*string, bool) {
	if o == nil || IsNil(o.PickupTotalWeightUnit) {
		return nil, false
	}
	return o.PickupTotalWeightUnit, true
}

// HasPickupTotalWeightUnit returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupTotalWeightUnit() bool {
	if o != nil && !IsNil(o.PickupTotalWeightUnit) {
		return true
	}

	return false
}

// SetPickupTotalWeightUnit gets a reference to the given string and assigns it to the PickupTotalWeightUnit field.
func (o *SchedulePickupDHLEXPResponse) SetPickupTotalWeightUnit(v string) {
	o.PickupTotalWeightUnit = &v
}

// GetPickupOptions returns the PickupOptions field value if set, zero value otherwise.
func (o *SchedulePickupDHLEXPResponse) GetPickupOptions() SchedulePickupDHLEXPResponsePickupOptions {
	if o == nil || IsNil(o.PickupOptions) {
		var ret SchedulePickupDHLEXPResponsePickupOptions
		return ret
	}
	return *o.PickupOptions
}

// GetPickupOptionsOk returns a tuple with the PickupOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchedulePickupDHLEXPResponse) GetPickupOptionsOk() (*SchedulePickupDHLEXPResponsePickupOptions, bool) {
	if o == nil || IsNil(o.PickupOptions) {
		return nil, false
	}
	return o.PickupOptions, true
}

// HasPickupOptions returns a boolean if a field has been set.
func (o *SchedulePickupDHLEXPResponse) HasPickupOptions() bool {
	if o != nil && !IsNil(o.PickupOptions) {
		return true
	}

	return false
}

// SetPickupOptions gets a reference to the given SchedulePickupDHLEXPResponsePickupOptions and assigns it to the PickupOptions field.
func (o *SchedulePickupDHLEXPResponse) SetPickupOptions(v SchedulePickupDHLEXPResponsePickupOptions) {
	o.PickupOptions = &v
}

func (o SchedulePickupDHLEXPResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchedulePickupDHLEXPResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PackageLocation) {
		toSerialize["packageLocation"] = o.PackageLocation
	}
	if !IsNil(o.PickupConfirmationNumber) {
		toSerialize["pickupConfirmationNumber"] = o.PickupConfirmationNumber
	}
	if !IsNil(o.PickupId) {
		toSerialize["pickupId"] = o.PickupId
	}
	if !IsNil(o.Carrier) {
		toSerialize["carrier"] = o.Carrier
	}
	if !IsNil(o.CarrierAccountId) {
		toSerialize["carrierAccountId"] = o.CarrierAccountId
	}
	if !IsNil(o.PickupAddress) {
		toSerialize["pickupAddress"] = o.PickupAddress
	}
	if !IsNil(o.ShipmentIds) {
		toSerialize["shipmentIds"] = o.ShipmentIds
	}
	if !IsNil(o.PickupSummary) {
		toSerialize["pickupSummary"] = o.PickupSummary
	}
	if !IsNil(o.Additionalnotes) {
		toSerialize["additionalnotes"] = o.Additionalnotes
	}
	if !IsNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !IsNil(o.PickupDateTime) {
		toSerialize["pickupDateTime"] = o.PickupDateTime
	}
	if !IsNil(o.PickupTotalWeight) {
		toSerialize["pickupTotalWeight"] = o.PickupTotalWeight
	}
	if !IsNil(o.PickupTotalWeightUnit) {
		toSerialize["pickupTotalWeightUnit"] = o.PickupTotalWeightUnit
	}
	if !IsNil(o.PickupOptions) {
		toSerialize["pickupOptions"] = o.PickupOptions
	}
	return toSerialize, nil
}

type NullableSchedulePickupDHLEXPResponse struct {
	value *SchedulePickupDHLEXPResponse
	isSet bool
}

func (v NullableSchedulePickupDHLEXPResponse) Get() *SchedulePickupDHLEXPResponse {
	return v.value
}

func (v *NullableSchedulePickupDHLEXPResponse) Set(val *SchedulePickupDHLEXPResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupDHLEXPResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupDHLEXPResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupDHLEXPResponse(val *SchedulePickupDHLEXPResponse) *NullableSchedulePickupDHLEXPResponse {
	return &NullableSchedulePickupDHLEXPResponse{value: val, isSet: true}
}

func (v NullableSchedulePickupDHLEXPResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupDHLEXPResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


