# RateShop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateOfShipment** | Pointer to **string** | This defines the date of the Shipment in the format YYYY:MM:DD.,required for future shipment rating | [optional] 
**FromAddress** | [**RateShopFromAddress**](RateShopFromAddress.md) |  | 
**Parcel** | [**RateShopParcel**](RateShopParcel.md) |  | 
**CarrierAccounts** | Pointer to **[]string** | This provide one or more carrier accounts Ids for rate shop. It can be referred from &#x60;Get Carrier Accounts&#x60; API | [optional] 
**ParcelType** | Pointer to **string** | Parcel Type required for rating, it&#39;s value can be referred from response of &#x60;Get Parcel Types&#x60; API | [optional] 
**ParcelId** | Pointer to **string** | &gt;-Parcel Id is optional and required to be given in case of branded parcels which have brandedDimension and/or brandedWeight. If parcel has brandedDimension, in that case user need not to pass dimensionUnit and dimension details(length, width and height) in the parcel object. And if brandedWeight is also available for the parcel then in that case weightUnit and weight need not to be passed  in parcel object | [optional] 
**ToAddress** | [**SingleRateToAddress**](SingleRateToAddress.md) |  | 
**IsHazmat** | Pointer to **bool** | isHazmat if set to true,only services which support Hazardous material shipment would be visible in the response | [optional] 

## Methods

### NewRateShop

`func NewRateShop(fromAddress RateShopFromAddress, parcel RateShopParcel, toAddress SingleRateToAddress, ) *RateShop`

NewRateShop instantiates a new RateShop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateShopWithDefaults

`func NewRateShopWithDefaults() *RateShop`

NewRateShopWithDefaults instantiates a new RateShop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateOfShipment

`func (o *RateShop) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *RateShop) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *RateShop) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *RateShop) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetFromAddress

`func (o *RateShop) GetFromAddress() RateShopFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *RateShop) GetFromAddressOk() (*RateShopFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *RateShop) SetFromAddress(v RateShopFromAddress)`

SetFromAddress sets FromAddress field to given value.


### GetParcel

`func (o *RateShop) GetParcel() RateShopParcel`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *RateShop) GetParcelOk() (*RateShopParcel, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *RateShop) SetParcel(v RateShopParcel)`

SetParcel sets Parcel field to given value.


### GetCarrierAccounts

`func (o *RateShop) GetCarrierAccounts() []string`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *RateShop) GetCarrierAccountsOk() (*[]string, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *RateShop) SetCarrierAccounts(v []string)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *RateShop) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetParcelType

`func (o *RateShop) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *RateShop) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *RateShop) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.

### HasParcelType

`func (o *RateShop) HasParcelType() bool`

HasParcelType returns a boolean if a field has been set.

### GetParcelId

`func (o *RateShop) GetParcelId() string`

GetParcelId returns the ParcelId field if non-nil, zero value otherwise.

### GetParcelIdOk

`func (o *RateShop) GetParcelIdOk() (*string, bool)`

GetParcelIdOk returns a tuple with the ParcelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelId

`func (o *RateShop) SetParcelId(v string)`

SetParcelId sets ParcelId field to given value.

### HasParcelId

`func (o *RateShop) HasParcelId() bool`

HasParcelId returns a boolean if a field has been set.

### GetToAddress

`func (o *RateShop) GetToAddress() SingleRateToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *RateShop) GetToAddressOk() (*SingleRateToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *RateShop) SetToAddress(v SingleRateToAddress)`

SetToAddress sets ToAddress field to given value.


### GetIsHazmat

`func (o *RateShop) GetIsHazmat() bool`

GetIsHazmat returns the IsHazmat field if non-nil, zero value otherwise.

### GetIsHazmatOk

`func (o *RateShop) GetIsHazmatOk() (*bool, bool)`

GetIsHazmatOk returns a tuple with the IsHazmat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHazmat

`func (o *RateShop) SetIsHazmat(v bool)`

SetIsHazmat sets IsHazmat field to given value.

### HasIsHazmat

`func (o *RateShop) HasIsHazmat() bool`

HasIsHazmat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


