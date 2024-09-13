# SingleRate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfShipment** | Pointer to **string** | This defines the date of the Shipment in the format YYYY:MM:DD, required for future shipment rating | [optional] 
**FromAddress** | [**SingleRateFromAddress**](SingleRateFromAddress.md) |  | 
**Parcel** | [**SingleRateParcel**](SingleRateParcel.md) |  | 
**CarrierAccounts** | **[]string** |  This provides a single carrier account Id in case of single rate request. It can be referred from response of &#x60;Get Carrier Accounts&#x60; API. | 
**ParcelType** | **string** | Parcel Type its value can be referred from response of &#x60;Get Parcel Types&#x60; API. | 
**ParcelId** | Pointer to **string** | &gt;-Parcel Id is optional and required to be given in case of branded parcels which have brandedDimension and/or brandedWeight. If parcel has brandedDimension, in that case user need not to pass dimensionUnit and dimension details(length, width and height) in the parcel object. And if brandedWeight is also available for the parcel then in that case weightUnit and weight need not to be passed  in parcel object   | [optional] 
**ServiceId** | **string** | Service to be used for rating, it can be referred from response of &#x60;Get Services&#x60; API | 
**SpecialServices** | Pointer to [**[]SpecialService**](SpecialService.md) | Special services to be used for rating, it can be referred from response of &#x60;Get Special Services&#x60; API | [optional] 
**ToAddress** | [**SingleRateToAddress**](SingleRateToAddress.md) |  | 

## Methods

### NewSingleRate

`func NewSingleRate(fromAddress SingleRateFromAddress, parcel SingleRateParcel, carrierAccounts []string, parcelType string, serviceId string, toAddress SingleRateToAddress, ) *SingleRate`

NewSingleRate instantiates a new SingleRate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleRateWithDefaults

`func NewSingleRateWithDefaults() *SingleRate`

NewSingleRateWithDefaults instantiates a new SingleRate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfShipment

`func (o *SingleRate) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *SingleRate) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *SingleRate) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *SingleRate) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetFromAddress

`func (o *SingleRate) GetFromAddress() SingleRateFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *SingleRate) GetFromAddressOk() (*SingleRateFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *SingleRate) SetFromAddress(v SingleRateFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *SingleRate) GetParcel() SingleRateParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *SingleRate) GetParcelOk() (*SingleRateParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *SingleRate) SetParcel(v SingleRateParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccounts

`func (o *SingleRate) GetCarrierAccounts() []string`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *SingleRate) GetCarrierAccountsOk() (*[]string, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *SingleRate) SetCarrierAccounts(v []string)`

SetCarrierAccounts sets CarrierAccounts field to given value.


### GetParcelType

`func (o *SingleRate) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *SingleRate) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *SingleRate) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetParcelId

`func (o *SingleRate) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *SingleRate) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *SingleRate) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *SingleRate) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetServiceId

`func (o *SingleRate) GetServiceId() string`

GetServiceId returns the ServiceId field if non-nil, zero value otherwise.

### GetServiceIdOk

`func (o *SingleRate) GetServiceIdOk() (*string, bool)`

GetServiceIdOk returns a tuple with the ServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceId

`func (o *SingleRate) SetServiceId(v string)`

SetServiceId sets ServiceId field to given value.


### GetSpecialServices

`func (o *SingleRate) GetSpecialServices() []SpecialService`

GetSpecialServices returns the SpecialServices field if non-nil, zero value otherwise.

### GetSpecialServicesOk

`func (o *SingleRate) GetSpecialServicesOk() (*[]SpecialService, bool)`

GetSpecialServicesOk returns a tuple with the SpecialServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServices

`func (o *SingleRate) SetSpecialServices(v []SpecialService)`

SetSpecialServices sets SpecialServices field to given value.

### HasSpecialServices

`func (o *SingleRate) HasSpecialServices() bool`

HasSpecialServices returns a boolean if a field has been set.

### GetToAddress

`func (o *SingleRate) GetToAddress() SingleRateToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *SingleRate) GetToAddressOk() (*SingleRateToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *SingleRate) SetToAddress(v SingleRateToAddress)`

SetToAddress sets ToAddress field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


