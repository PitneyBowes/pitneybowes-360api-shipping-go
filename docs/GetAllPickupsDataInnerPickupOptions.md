# GetAllPickupsDataInnerPickupOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PickupStartDateTime** | Pointer to **string** | It specifies the earliest date and time that your parcels will be available for pickup (UTC time) | [optional] 
**PickupEndDateTime** | Pointer to **string** | It specifies the latest date and time that your parcels will be available for pickup (UTC time) | [optional] 
**Overweight** | Pointer to **float32** | It specifies the count of packages which are overwight (weight&gt; 70 lbs) | [optional] 
**CarrierType** | Pointer to **string** | It specifies the type of pickup service - &#x60;Ground&#x60; or &#x60;Express&#x60;. Choose &#x60;Ground&#x60; to schedule pickup for next business day or up to two weeks later for Ground packages only (i.e. GRD,HOM,STD,SP_MEDIA,SP_PRCLSEL,SP_PRE_PRINT,SP_PRE_STD,SP_RETURNS). Choose &#x60;Express&#x60; to schedule pickup a same day or a next day pickup for express packages only (i.e. NDA_AM,NDA,NDA_SVR,NDA_AM_EH,NDA_EH,NDA_SVR_EH,2DA_AM,2DA,3DA,XPP,EXP,EXP_X,EXP_CP,XPD). | [optional] 

## Methods

### NewGetAllPickupsDataInnerPickupOptions

`func NewGetAllPickupsDataInnerPickupOptions() *GetAllPickupsDataInnerPickupOptions`

NewGetAllPickupsDataInnerPickupOptions instantiates a new GetAllPickupsDataInnerPickupOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllPickupsDataInnerPickupOptionsWithDefaults

`func NewGetAllPickupsDataInnerPickupOptionsWithDefaults() *GetAllPickupsDataInnerPickupOptions`

NewGetAllPickupsDataInnerPickupOptionsWithDefaults instantiates a new GetAllPickupsDataInnerPickupOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPickupStartDateTime

`func (o *GetAllPickupsDataInnerPickupOptions) GetPickupStartDateTime() string`

GetPickupStartDateTime returns the PickupStartDateTime field if non-nil, zero value otherwise.

### GetPickupStartDateTimeOk

`func (o *GetAllPickupsDataInnerPickupOptions) GetPickupStartDateTimeOk() (*string, bool)`

GetPickupStartDateTimeOk returns a tuple with the PickupStartDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupStartDateTime

`func (o *GetAllPickupsDataInnerPickupOptions) SetPickupStartDateTime(v string)`

SetPickupStartDateTime sets PickupStartDateTime field to given value.

### HasPickupStartDateTime

`func (o *GetAllPickupsDataInnerPickupOptions) HasPickupStartDateTime() bool`

HasPickupStartDateTime returns a boolean if a field has been set.

### GetPickupEndDateTime

`func (o *GetAllPickupsDataInnerPickupOptions) GetPickupEndDateTime() string`

GetPickupEndDateTime returns the PickupEndDateTime field if non-nil, zero value otherwise.

### GetPickupEndDateTimeOk

`func (o *GetAllPickupsDataInnerPickupOptions) GetPickupEndDateTimeOk() (*string, bool)`

GetPickupEndDateTimeOk returns a tuple with the PickupEndDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupEndDateTime

`func (o *GetAllPickupsDataInnerPickupOptions) SetPickupEndDateTime(v string)`

SetPickupEndDateTime sets PickupEndDateTime field to given value.

### HasPickupEndDateTime

`func (o *GetAllPickupsDataInnerPickupOptions) HasPickupEndDateTime() bool`

HasPickupEndDateTime returns a boolean if a field has been set.

### GetOverweight

`func (o *GetAllPickupsDataInnerPickupOptions) GetOverweight() float32`

GetOverweight returns the Overweight field if non-nil, zero value otherwise.

### GetOverweightOk

`func (o *GetAllPickupsDataInnerPickupOptions) GetOverweightOk() (*float32, bool)`

GetOverweightOk returns a tuple with the Overweight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverweight

`func (o *GetAllPickupsDataInnerPickupOptions) SetOverweight(v float32)`

SetOverweight sets Overweight field to given value.

### HasOverweight

`func (o *GetAllPickupsDataInnerPickupOptions) HasOverweight() bool`

HasOverweight returns a boolean if a field has been set.

### GetCarrierType

`func (o *GetAllPickupsDataInnerPickupOptions) GetCarrierType() string`

GetCarrierType returns the CarrierType field if non-nil, zero value otherwise.

### GetCarrierTypeOk

`func (o *GetAllPickupsDataInnerPickupOptions) GetCarrierTypeOk() (*string, bool)`

GetCarrierTypeOk returns a tuple with the CarrierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierType

`func (o *GetAllPickupsDataInnerPickupOptions) SetCarrierType(v string)`

SetCarrierType sets CarrierType field to given value.

### HasCarrierType

`func (o *GetAllPickupsDataInnerPickupOptions) HasCarrierType() bool`

HasCarrierType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


