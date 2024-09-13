# SchedulePickupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PackageLocation** | **string** | It specifies the location from where packages would be collected. Applicable values are &#x60;FRONT&#x60;,&#x60;NONE&#x60;,&#x60;SIDE&#x60;,&#x60;REAR&#x60; | 
**CarrierAccountId** | **string** | It specifies the carrier account id, its value can be referenced from the &#x60;Get Carrier Accounts&#x60; API. | 
**PickupAddress** | [**SchedulePickupDHLEXPRequestPickupAddress**](SchedulePickupDHLEXPRequestPickupAddress.md) |  | 
**ShipmentIds** | Pointer to **[]string** | It indicates the shipmentIds for which pickup to be scheduled. | [optional] 
**PickupSummary** | Pointer to [**[]SchedulePickupFedexRequestPickupSummaryInner**](SchedulePickupFedexRequestPickupSummaryInner.md) | This can be used to add package details for which labels are not created yet but would want to schedule pickup in advance. | [optional] 
**Additionalnotes** | Pointer to **string** | It can be used to provide any additional comments or remarks, it would be printed on the scheduled pickup document. | [optional] 
**Reference** | Pointer to **string** | It is used for any reference purpose | [optional] 
**PickupOptions** | [**SchedulePickupFedexRequestPickupOptions**](SchedulePickupFedexRequestPickupOptions.md) |  | 

## Methods

### NewSchedulePickupRequest

`func NewSchedulePickupRequest(packageLocation string, carrierAccountId string, pickupAddress SchedulePickupDHLEXPRequestPickupAddress, pickupOptions SchedulePickupFedexRequestPickupOptions, ) *SchedulePickupRequest`

NewSchedulePickupRequest instantiates a new SchedulePickupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchedulePickupRequestWithDefaults

`func NewSchedulePickupRequestWithDefaults() *SchedulePickupRequest`

NewSchedulePickupRequestWithDefaults instantiates a new SchedulePickupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPackageLocation

`func (o *SchedulePickupRequest) GetPackageLocation() string`

GetPackageLocation returns the PackageLocation field if non-nil, zero value otherwise.

### GetPackageLocationOk

`func (o *SchedulePickupRequest) GetPackageLocationOk() (*string, bool)`

GetPackageLocationOk returns a tuple with the PackageLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackageLocation

`func (o *SchedulePickupRequest) SetPackageLocation(v string)`

SetPackageLocation sets PackageLocation field to given value.


### GetCarrierAccountId

`func (o *SchedulePickupRequest) GetCarrierAccountId() string`

GetCarrierAccountId returns the CarrierAccountId field if non-nil, zero value otherwise.

### GetCarrierAccountIdOk

`func (o *SchedulePickupRequest) GetCarrierAccountIdOk() (*string, bool)`

GetCarrierAccountIdOk returns a tuple with the CarrierAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccountId

`func (o *SchedulePickupRequest) SetCarrierAccountId(v string)`

SetCarrierAccountId sets CarrierAccountId field to given value.


### GetPickupAddress

`func (o *SchedulePickupRequest) GetPickupAddress() SchedulePickupDHLEXPRequestPickupAddress`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *SchedulePickupRequest) GetPickupAddressOk() (*SchedulePickupDHLEXPRequestPickupAddress, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *SchedulePickupRequest) SetPickupAddress(v SchedulePickupDHLEXPRequestPickupAddress)`

SetPickupAddress sets PickupAddress field to given value.


### GetShipmentIds

`func (o *SchedulePickupRequest) GetShipmentIds() []string`

GetShipmentIds returns the ShipmentIds field if non-nil, zero value otherwise.

### GetShipmentIdsOk

`func (o *SchedulePickupRequest) GetShipmentIdsOk() (*[]string, bool)`

GetShipmentIdsOk returns a tuple with the ShipmentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentIds

`func (o *SchedulePickupRequest) SetShipmentIds(v []string)`

SetShipmentIds sets ShipmentIds field to given value.

### HasShipmentIds

`func (o *SchedulePickupRequest) HasShipmentIds() bool`

HasShipmentIds returns a boolean if a field has been set.

### GetPickupSummary

`func (o *SchedulePickupRequest) GetPickupSummary() []SchedulePickupFedexRequestPickupSummaryInner`

GetPickupSummary returns the PickupSummary field if non-nil, zero value otherwise.

### GetPickupSummaryOk

`func (o *SchedulePickupRequest) GetPickupSummaryOk() (*[]SchedulePickupFedexRequestPickupSummaryInner, bool)`

GetPickupSummaryOk returns a tuple with the PickupSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupSummary

`func (o *SchedulePickupRequest) SetPickupSummary(v []SchedulePickupFedexRequestPickupSummaryInner)`

SetPickupSummary sets PickupSummary field to given value.

### HasPickupSummary

`func (o *SchedulePickupRequest) HasPickupSummary() bool`

HasPickupSummary returns a boolean if a field has been set.

### GetAdditionalnotes

`func (o *SchedulePickupRequest) GetAdditionalnotes() string`

GetAdditionalnotes returns the Additionalnotes field if non-nil, zero value otherwise.

### GetAdditionalnotesOk

`func (o *SchedulePickupRequest) GetAdditionalnotesOk() (*string, bool)`

GetAdditionalnotesOk returns a tuple with the Additionalnotes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalnotes

`func (o *SchedulePickupRequest) SetAdditionalnotes(v string)`

SetAdditionalnotes sets Additionalnotes field to given value.

### HasAdditionalnotes

`func (o *SchedulePickupRequest) HasAdditionalnotes() bool`

HasAdditionalnotes returns a boolean if a field has been set.

### GetReference

`func (o *SchedulePickupRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *SchedulePickupRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *SchedulePickupRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *SchedulePickupRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetPickupOptions

`func (o *SchedulePickupRequest) GetPickupOptions() SchedulePickupFedexRequestPickupOptions`

GetPickupOptions returns the PickupOptions field if non-nil, zero value otherwise.

### GetPickupOptionsOk

`func (o *SchedulePickupRequest) GetPickupOptionsOk() (*SchedulePickupFedexRequestPickupOptions, bool)`

GetPickupOptionsOk returns a tuple with the PickupOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupOptions

`func (o *SchedulePickupRequest) SetPickupOptions(v SchedulePickupFedexRequestPickupOptions)`

SetPickupOptions sets PickupOptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


