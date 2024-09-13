# SchedulePickupDHLEXPRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | **string** | It specifies the location from where packages would be collected. Applicable values are &#x60;NONE&#x60;, &#x60;FRONT&#x60;,&#x60;RECEPTION&#x60;,&#x60;BACK&#x60;,&#x60;LOADING DOCK&#x60; | 
**CarrierAccountId** | **string** | It specifies the carrier account id, its value can be referenced from the &#x60;Get Carrier Accounts&#x60; API. | 
**PickupAddress** | [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | 
**ShipmentIds** | Pointer to **[]string** | It indicates the shipmentIds for which pickup to be scheduled. | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupDHLEXPRequestPickupSummaryInner**](SchedulePickupDHLEXPRequestPickupSummaryInner.md) | This can be used to add package details for which labels are not created yet but would want to schedule pickup in advance. | [optional] 
**Additionalnotes** | Pointer to **string** | It can be used to provide any additional comments or remarks, it would be printed on the scheduled pickup document | [optional] 
**Reference** | Pointer to **string** | It is used for any reference purpose | [optional] 
**PickupOptions** | [**SchedulePickupDHLEXPRequestPickupOptions**](SchedulePickupDHLEXPRequestPickupOptions.md) |  | 

## Methods

### NewSchedulePickupDHLEXPRequest

`func NewSchedulePickupDHLEXPRequest(packageLocation string, carrierAccountId string, pickupAddress SchedulePickupDHLEXPRequestPickupAddress, pickupOptions SchedulePickupDHLEXPRequestPickupOptions, ) *SchedulePickupDHLEXPRequest`

NewSchedulePickupDHLEXPRequest instantiates a new SchedulePickupDHLEXPRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupDHLEXPRequestWithDefaults

`func NewSchedulePickupDHLEXPRequestWithDefaults() *SchedulePickupDHLEXPRequest`

NewSchedulePickupDHLEXPRequestWithDefaults instantiates a new SchedulePickupDHLEXPRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupDHLEXPRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupDHLEXPRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupDHLEXPRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.


### GetCarrierAccountId

`func (o *SchedulePickupDHLEXPRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupDHLEXPRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupDHLEXPRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetPickupAddress

`func (o *SchedulePickupDHLEXPRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupDHLEXPRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.


### GetShipmentIds

`func (o *SchedulePickupDHLEXPRequest) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickupDHLEXPRequest) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickupDHLEXPRequest) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickupDHLEXPRequest) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupDHLEXPRequest) GetPickupSummary() []SchedulePickupDHLEXPRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupSummaryOk() (*[]SchedulePickupDHLEXPRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupDHLEXPRequest) SetPickupSummary(v []SchedulePickupDHLEXPRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupDHLEXPRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupDHLEXPRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupDHLEXPRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupDHLEXPRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupDHLEXPRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupDHLEXPRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupDHLEXPRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupDHLEXPRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupDHLEXPRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupDHLEXPRequest) GetPickupOptions() SchedulePickupDHLEXPRequestPickupOptions`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupDHLEXPRequest) GetPickupOptionsOk() (*SchedulePickupDHLEXPRequestPickupOptions, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupDHLEXPRequest) SetPickupOptions(v SchedulePickupDHLEXPRequestPickupOptions)`

SetPickupOptions sets PickupOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


