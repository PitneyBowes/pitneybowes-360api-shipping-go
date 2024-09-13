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

// checks if the BulkShipmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkShipmentResponse{}

// BulkShipmentResponse struct for BulkShipmentResponse
type BulkShipmentResponse struct {
	//  This is the system generated Batch ID.
	BatchId *string `json:"batchId,omitempty"`
	// This indicates group name for the Batch.
	GroupName *string `json:"groupName,omitempty"`
	//  The name of the Batch.
	Name *string `json:"name,omitempty"`
	Status *JobStatus `json:"status,omitempty"`
	AddressValidation CounterStatus `json:"addressValidation,omitempty"`
	Rating CounterStatus `json:"rating,omitempty"`
	LabelGeneration CounterStatus `json:"labelGeneration,omitempty"`
}

// NewBulkShipmentResponse instantiates a new BulkShipmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkShipmentResponse() *BulkShipmentResponse {
	this := BulkShipmentResponse{}
	return &this
}

// NewBulkShipmentResponseWithDefaults instantiates a new BulkShipmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkShipmentResponseWithDefaults() *BulkShipmentResponse {
	this := BulkShipmentResponse{}
	return &this
}

// GetBatchId returns the BatchId field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetBatchId() string {
	if o == nil || IsNil(o.BatchId) {
		var ret string
		return ret
	}
	return *o.BatchId
}

// GetBatchIdOk returns a tuple with the BatchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetBatchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BatchId) {
		return nil, false
	}
	return o.BatchId, true
}

// HasBatchId returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasBatchId() bool {
	if o != nil && !IsNil(o.BatchId) {
		return true
	}

	return false
}

// SetBatchId gets a reference to the given string and assigns it to the BatchId field.
func (o *BulkShipmentResponse) SetBatchId(v string) {
	o.BatchId = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetGroupName() string {
	if o == nil || IsNil(o.GroupName) {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetGroupNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroupName) {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasGroupName() bool {
	if o != nil && !IsNil(o.GroupName) {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *BulkShipmentResponse) SetGroupName(v string) {
	o.GroupName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *BulkShipmentResponse) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetStatus() JobStatus {
	if o == nil || IsNil(o.Status) {
		var ret JobStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetStatusOk() (*JobStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given JobStatus and assigns it to the Status field.
func (o *BulkShipmentResponse) SetStatus(v JobStatus) {
	o.Status = &v
}

// GetAddressValidation returns the AddressValidation field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetAddressValidation() CounterStatus {
	if o == nil || IsNil(o.AddressValidation) {
		var ret CounterStatus
		return ret
	}
	return o.AddressValidation
}

// GetAddressValidationOk returns a tuple with the AddressValidation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetAddressValidationOk() (CounterStatus, bool) {
	if o == nil || IsNil(o.AddressValidation) {
		return CounterStatus{}, false
	}
	return o.AddressValidation, true
}

// HasAddressValidation returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasAddressValidation() bool {
	if o != nil && !IsNil(o.AddressValidation) {
		return true
	}

	return false
}

// SetAddressValidation gets a reference to the given CounterStatus and assigns it to the AddressValidation field.
func (o *BulkShipmentResponse) SetAddressValidation(v CounterStatus) {
	o.AddressValidation = v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetRating() CounterStatus {
	if o == nil || IsNil(o.Rating) {
		var ret CounterStatus
		return ret
	}
	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetRatingOk() (CounterStatus, bool) {
	if o == nil || IsNil(o.Rating) {
		return CounterStatus{}, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given CounterStatus and assigns it to the Rating field.
func (o *BulkShipmentResponse) SetRating(v CounterStatus) {
	o.Rating = v
}

// GetLabelGeneration returns the LabelGeneration field value if set, zero value otherwise.
func (o *BulkShipmentResponse) GetLabelGeneration() CounterStatus {
	if o == nil || IsNil(o.LabelGeneration) {
		var ret CounterStatus
		return ret
	}
	return o.LabelGeneration
}

// GetLabelGenerationOk returns a tuple with the LabelGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkShipmentResponse) GetLabelGenerationOk() (CounterStatus, bool) {
	if o == nil || IsNil(o.LabelGeneration) {
		return CounterStatus{}, false
	}
	return o.LabelGeneration, true
}

// HasLabelGeneration returns a boolean if a field has been set.
func (o *BulkShipmentResponse) HasLabelGeneration() bool {
	if o != nil && !IsNil(o.LabelGeneration) {
		return true
	}

	return false
}

// SetLabelGeneration gets a reference to the given CounterStatus and assigns it to the LabelGeneration field.
func (o *BulkShipmentResponse) SetLabelGeneration(v CounterStatus) {
	o.LabelGeneration = v
}

func (o BulkShipmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkShipmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BatchId) {
		toSerialize["batchId"] = o.BatchId
	}
	if !IsNil(o.GroupName) {
		toSerialize["groupName"] = o.GroupName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.AddressValidation) {
		toSerialize["addressValidation"] = o.AddressValidation
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.LabelGeneration) {
		toSerialize["labelGeneration"] = o.LabelGeneration
	}
	return toSerialize, nil
}

type NullableBulkShipmentResponse struct {
	value *BulkShipmentResponse
	isSet bool
}

func (v NullableBulkShipmentResponse) Get() *BulkShipmentResponse {
	return v.value
}

func (v *NullableBulkShipmentResponse) Set(val *BulkShipmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkShipmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkShipmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkShipmentResponse(val *BulkShipmentResponse) *NullableBulkShipmentResponse {
	return &NullableBulkShipmentResponse{value: val, isSet: true}
}

func (v NullableBulkShipmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkShipmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


