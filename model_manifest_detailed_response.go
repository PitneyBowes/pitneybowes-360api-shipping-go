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

// checks if the ManifestDetailedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ManifestDetailedResponse{}

// ManifestDetailedResponse struct for ManifestDetailedResponse
type ManifestDetailedResponse struct {
	// A unique identifier associated with the Carrier account which is used while creating Manifest.
	CarrierAccountId *string `json:"carrierAccountId,omitempty"`
	// Name of the Carrier.
	CarrierName *string `json:"carrierName,omitempty"`
	// The electronically generated document that has manifest (end-of-day) records of all shipments of the day.
	ManifestDocuments []ManifestDetailedResponseManifestDocumentsInner `json:"manifestDocuments,omitempty"`
	// The unique manifest ID. This field is not returned for APAC Services.
	ManifestID *string `json:"manifestID,omitempty"`
	// The manifest tracking number. This is returned only if the carrier has a pre-defined valid value, e.g., UPS, FedEX, or USPS.
	ManifestTrackingNumber *string `json:"manifestTrackingNumber,omitempty"`
	FromAddress *ManifestDetailedResponseFromAddress `json:"fromAddress,omitempty"`
	ParcelTrackingNumbers []interface{} `json:"parcelTrackingNumbers,omitempty"`
	// The date the shipments are to be tendered to the carrier, entered as YYYY-MM-DD.
	SubmissionDate *string `json:"submissionDate,omitempty"`
}

// NewManifestDetailedResponse instantiates a new ManifestDetailedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewManifestDetailedResponse() *ManifestDetailedResponse {
	this := ManifestDetailedResponse{}
	return &this
}

// NewManifestDetailedResponseWithDefaults instantiates a new ManifestDetailedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewManifestDetailedResponseWithDefaults() *ManifestDetailedResponse {
	this := ManifestDetailedResponse{}
	return &this
}

// GetCarrierAccountId returns the CarrierAccountId field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetCarrierAccountId() string {
	if o == nil || IsNil(o.CarrierAccountId) {
		var ret string
		return ret
	}
	return *o.CarrierAccountId
}

// GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetCarrierAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierAccountId) {
		return nil, false
	}
	return o.CarrierAccountId, true
}

// HasCarrierAccountId returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasCarrierAccountId() bool {
	if o != nil && !IsNil(o.CarrierAccountId) {
		return true
	}

	return false
}

// SetCarrierAccountId gets a reference to the given string and assigns it to the CarrierAccountId field.
func (o *ManifestDetailedResponse) SetCarrierAccountId(v string) {
	o.CarrierAccountId = &v
}

// GetCarrierName returns the CarrierName field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetCarrierName() string {
	if o == nil || IsNil(o.CarrierName) {
		var ret string
		return ret
	}
	return *o.CarrierName
}

// GetCarrierNameOk returns a tuple with the CarrierName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetCarrierNameOk() (*string, bool) {
	if o == nil || IsNil(o.CarrierName) {
		return nil, false
	}
	return o.CarrierName, true
}

// HasCarrierName returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasCarrierName() bool {
	if o != nil && !IsNil(o.CarrierName) {
		return true
	}

	return false
}

// SetCarrierName gets a reference to the given string and assigns it to the CarrierName field.
func (o *ManifestDetailedResponse) SetCarrierName(v string) {
	o.CarrierName = &v
}

// GetManifestDocuments returns the ManifestDocuments field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetManifestDocuments() []ManifestDetailedResponseManifestDocumentsInner {
	if o == nil || IsNil(o.ManifestDocuments) {
		var ret []ManifestDetailedResponseManifestDocumentsInner
		return ret
	}
	return o.ManifestDocuments
}

// GetManifestDocumentsOk returns a tuple with the ManifestDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetManifestDocumentsOk() ([]ManifestDetailedResponseManifestDocumentsInner, bool) {
	if o == nil || IsNil(o.ManifestDocuments) {
		return nil, false
	}
	return o.ManifestDocuments, true
}

// HasManifestDocuments returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasManifestDocuments() bool {
	if o != nil && !IsNil(o.ManifestDocuments) {
		return true
	}

	return false
}

// SetManifestDocuments gets a reference to the given []ManifestDetailedResponseManifestDocumentsInner and assigns it to the ManifestDocuments field.
func (o *ManifestDetailedResponse) SetManifestDocuments(v []ManifestDetailedResponseManifestDocumentsInner) {
	o.ManifestDocuments = v
}

// GetManifestID returns the ManifestID field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetManifestID() string {
	if o == nil || IsNil(o.ManifestID) {
		var ret string
		return ret
	}
	return *o.ManifestID
}

// GetManifestIDOk returns a tuple with the ManifestID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetManifestIDOk() (*string, bool) {
	if o == nil || IsNil(o.ManifestID) {
		return nil, false
	}
	return o.ManifestID, true
}

// HasManifestID returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasManifestID() bool {
	if o != nil && !IsNil(o.ManifestID) {
		return true
	}

	return false
}

// SetManifestID gets a reference to the given string and assigns it to the ManifestID field.
func (o *ManifestDetailedResponse) SetManifestID(v string) {
	o.ManifestID = &v
}

// GetManifestTrackingNumber returns the ManifestTrackingNumber field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetManifestTrackingNumber() string {
	if o == nil || IsNil(o.ManifestTrackingNumber) {
		var ret string
		return ret
	}
	return *o.ManifestTrackingNumber
}

// GetManifestTrackingNumberOk returns a tuple with the ManifestTrackingNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetManifestTrackingNumberOk() (*string, bool) {
	if o == nil || IsNil(o.ManifestTrackingNumber) {
		return nil, false
	}
	return o.ManifestTrackingNumber, true
}

// HasManifestTrackingNumber returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasManifestTrackingNumber() bool {
	if o != nil && !IsNil(o.ManifestTrackingNumber) {
		return true
	}

	return false
}

// SetManifestTrackingNumber gets a reference to the given string and assigns it to the ManifestTrackingNumber field.
func (o *ManifestDetailedResponse) SetManifestTrackingNumber(v string) {
	o.ManifestTrackingNumber = &v
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetFromAddress() ManifestDetailedResponseFromAddress {
	if o == nil || IsNil(o.FromAddress) {
		var ret ManifestDetailedResponseFromAddress
		return ret
	}
	return *o.FromAddress
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetFromAddressOk() (*ManifestDetailedResponseFromAddress, bool) {
	if o == nil || IsNil(o.FromAddress) {
		return nil, false
	}
	return o.FromAddress, true
}

// HasFromAddress returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasFromAddress() bool {
	if o != nil && !IsNil(o.FromAddress) {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given ManifestDetailedResponseFromAddress and assigns it to the FromAddress field.
func (o *ManifestDetailedResponse) SetFromAddress(v ManifestDetailedResponseFromAddress) {
	o.FromAddress = &v
}

// GetParcelTrackingNumbers returns the ParcelTrackingNumbers field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetParcelTrackingNumbers() []interface{} {
	if o == nil || IsNil(o.ParcelTrackingNumbers) {
		var ret []interface{}
		return ret
	}
	return o.ParcelTrackingNumbers
}

// GetParcelTrackingNumbersOk returns a tuple with the ParcelTrackingNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetParcelTrackingNumbersOk() ([]interface{}, bool) {
	if o == nil || IsNil(o.ParcelTrackingNumbers) {
		return nil, false
	}
	return o.ParcelTrackingNumbers, true
}

// HasParcelTrackingNumbers returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasParcelTrackingNumbers() bool {
	if o != nil && !IsNil(o.ParcelTrackingNumbers) {
		return true
	}

	return false
}

// SetParcelTrackingNumbers gets a reference to the given []interface{} and assigns it to the ParcelTrackingNumbers field.
func (o *ManifestDetailedResponse) SetParcelTrackingNumbers(v []interface{}) {
	o.ParcelTrackingNumbers = v
}

// GetSubmissionDate returns the SubmissionDate field value if set, zero value otherwise.
func (o *ManifestDetailedResponse) GetSubmissionDate() string {
	if o == nil || IsNil(o.SubmissionDate) {
		var ret string
		return ret
	}
	return *o.SubmissionDate
}

// GetSubmissionDateOk returns a tuple with the SubmissionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ManifestDetailedResponse) GetSubmissionDateOk() (*string, bool) {
	if o == nil || IsNil(o.SubmissionDate) {
		return nil, false
	}
	return o.SubmissionDate, true
}

// HasSubmissionDate returns a boolean if a field has been set.
func (o *ManifestDetailedResponse) HasSubmissionDate() bool {
	if o != nil && !IsNil(o.SubmissionDate) {
		return true
	}

	return false
}

// SetSubmissionDate gets a reference to the given string and assigns it to the SubmissionDate field.
func (o *ManifestDetailedResponse) SetSubmissionDate(v string) {
	o.SubmissionDate = &v
}

func (o ManifestDetailedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ManifestDetailedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CarrierAccountId) {
		toSerialize["carrierAccountId"] = o.CarrierAccountId
	}
	if !IsNil(o.CarrierName) {
		toSerialize["carrierName"] = o.CarrierName
	}
	if !IsNil(o.ManifestDocuments) {
		toSerialize["manifestDocuments"] = o.ManifestDocuments
	}
	if !IsNil(o.ManifestID) {
		toSerialize["manifestID"] = o.ManifestID
	}
	if !IsNil(o.ManifestTrackingNumber) {
		toSerialize["manifestTrackingNumber"] = o.ManifestTrackingNumber
	}
	if !IsNil(o.FromAddress) {
		toSerialize["fromAddress"] = o.FromAddress
	}
	if !IsNil(o.ParcelTrackingNumbers) {
		toSerialize["parcelTrackingNumbers"] = o.ParcelTrackingNumbers
	}
	if !IsNil(o.SubmissionDate) {
		toSerialize["submissionDate"] = o.SubmissionDate
	}
	return toSerialize, nil
}

type NullableManifestDetailedResponse struct {
	value *ManifestDetailedResponse
	isSet bool
}

func (v NullableManifestDetailedResponse) Get() *ManifestDetailedResponse {
	return v.value
}

func (v *NullableManifestDetailedResponse) Set(val *ManifestDetailedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableManifestDetailedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableManifestDetailedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManifestDetailedResponse(val *ManifestDetailedResponse) *NullableManifestDetailedResponse {
	return &NullableManifestDetailedResponse{value: val, isSet: true}
}

func (v NullableManifestDetailedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManifestDetailedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


