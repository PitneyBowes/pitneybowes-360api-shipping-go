# SchedulePickupUSPSRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | **string** | It specifies the location from where packages would be collected. Applicable values are &#x60;Front Door&#x60;,&#x60;Back Door&#x60;,&#x60;Side Door&#x60;,&#x60;Knock on Door/Ring Bell&#x60;,&#x60;Mail Room&#x60;,&#x60;Office&#x60;,&#x60;Reception,In/At Mailbox&#x60;,&#x60;Other&#x60; | 
**CarrierAccountId** | **string** | It specifies the carrier account id, its value can be referenced from the &#x60;Get Carrier Accounts&#x60; API. | 
**PickupAddress** | [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | 
**ShipmentIds** | Pointer to **[]string** | It indicates the shipmentIds for which pickup to be scheduled. | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupUSPSRequestPickupSummaryInner**](SchedulePickupUSPSRequestPickupSummaryInner.md) | This can be used to add package details for which labels are not created yet but would want to schedule pickup in advance. | [optional] 
**Additionalnotes** | Pointer to **string** | It can be used to provide any additional comments or remarks, it would be printed on the scheduled pickup document It is required to provide when packageLocation is set to &#x60;Other&#x60; | [optional] 
**Reference** | Pointer to **string** | It is used for any reference purpose | [optional] 

## Methods

### NewSchedulePickupUSPSRequest

`func NewSchedulePickupUSPSRequest(packageLocation string, carrierAccountId string, pickupAddress SchedulePickupDHLEXPRequestPickupAddress, ) *SchedulePickupUSPSRequest`

NewSchedulePickupUSPSRequest instantiates a new SchedulePickupUSPSRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupUSPSRequestWithDefaults

`func NewSchedulePickupUSPSRequestWithDefaults() *SchedulePickupUSPSRequest`

NewSchedulePickupUSPSRequestWithDefaults instantiates a new SchedulePickupUSPSRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupUSPSRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupUSPSRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupUSPSRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.


### GetCarrierAccountId

`func (o *SchedulePickupUSPSRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupUSPSRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupUSPSRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetPickupAddress

`func (o *SchedulePickupUSPSRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupUSPSRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupUSPSRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.


### GetShipmentIds

`func (o *SchedulePickupUSPSRequest) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickupUSPSRequest) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickupUSPSRequest) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickupUSPSRequest) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupUSPSRequest) GetPickupSummary() []SchedulePickupUSPSRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupUSPSRequest) GetPickupSummaryOk() (*[]SchedulePickupUSPSRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupUSPSRequest) SetPickupSummary(v []SchedulePickupUSPSRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupUSPSRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupUSPSRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupUSPSRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupUSPSRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupUSPSRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupUSPSRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupUSPSRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupUSPSRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupUSPSRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


