# CarriersCarriersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierID** | Pointer to **string** | A unique identifier associated with the specific carrier. | [optional] 
**Name** | Pointer to **string** | This is the name of the Carrier, a shipping service which is used to carry shipment (e.g., parcels/packages/envelopes) from one place to other. | [optional] 
**OriginCountry** | Pointer to **string** | This indicates the two-character ISO code of the source country from the ISO country list. | [optional] 
**Properties** | Pointer to [**CarriersCarriersInnerProperties**](CarriersCarriersInnerProperties.md) |  | [optional] 

## Methods

### NewCarriersCarriersInner

`func NewCarriersCarriersInner() *CarriersCarriersInner`

NewCarriersCarriersInner instantiates a new CarriersCarriersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarriersCarriersInnerWithDefaults

`func NewCarriersCarriersInnerWithDefaults() *CarriersCarriersInner`

NewCarriersCarriersInnerWithDefaults instantiates a new CarriersCarriersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierID

`func (o *CarriersCarriersInner) GetCarrierID() string`

GetCarrierID returns the CarrierID field if non-nil, zero value otherwise.

### GetCarrierIDOk

`func (o *CarriersCarriersInner) GetCarrierIDOk() (*string, bool)`

GetCarrierIDOk returns a tuple with the CarrierID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierID

`func (o *CarriersCarriersInner) SetCarrierID(v string)`

SetCarrierID sets CarrierID field to given value.

### HasCarrierID

`func (o *CarriersCarriersInner) HasCarrierID() bool`

HasCarrierID returns a boolean if a field has been set.

### GetName

`func (o *CarriersCarriersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CarriersCarriersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CarriersCarriersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CarriersCarriersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginCountry

`func (o *CarriersCarriersInner) GetOriginCountry() string`

GetOriginCountry returns the OriginCountry field if non-nil, zero value otherwise.

### GetOriginCountryOk

`func (o *CarriersCarriersInner) GetOriginCountryOk() (*string, bool)`

GetOriginCountryOk returns a tuple with the OriginCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginCountry

`func (o *CarriersCarriersInner) SetOriginCountry(v string)`

SetOriginCountry sets OriginCountry field to given value.

### HasOriginCountry

`func (o *CarriersCarriersInner) HasOriginCountry() bool`

HasOriginCountry returns a boolean if a field has been set.

### GetProperties

`func (o *CarriersCarriersInner) GetProperties() CarriersCarriersInnerProperties`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *CarriersCarriersInner) GetPropertiesOk() (*CarriersCarriersInnerProperties, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *CarriersCarriersInner) SetProperties(v CarriersCarriersInnerProperties)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *CarriersCarriersInner) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


