# GetSingleShipmentRateSpecialServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **float32** | The amount of the specialSevice. | [optional] 
**InputParameters** | Pointer to [**[]GetSingleShipmentRateSpecialServicesInnerInputParametersInner**](GetSingleShipmentRateSpecialServicesInnerInputParametersInner.md) | &gt;-The parameters to set for the special service, such as an insurance value or a receipt-number format. This is required if the specialservice requires input parameters. If a special service does not require input parameters, you can either leave out the array or pass an empty array. | [optional] 
**SpecialServiceId** | Pointer to **string** | A unique identifier associated to the Special Service , which depends on the carrier based service. | [optional] 

## Methods

### NewGetSingleShipmentRateSpecialServicesInner

`func NewGetSingleShipmentRateSpecialServicesInner() *GetSingleShipmentRateSpecialServicesInner`

NewGetSingleShipmentRateSpecialServicesInner instantiates a new GetSingleShipmentRateSpecialServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSingleShipmentRateSpecialServicesInnerWithDefaults

`func NewGetSingleShipmentRateSpecialServicesInnerWithDefaults() *GetSingleShipmentRateSpecialServicesInner`

NewGetSingleShipmentRateSpecialServicesInnerWithDefaults instantiates a new GetSingleShipmentRateSpecialServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *GetSingleShipmentRateSpecialServicesInner) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *GetSingleShipmentRateSpecialServicesInner) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *GetSingleShipmentRateSpecialServicesInner) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *GetSingleShipmentRateSpecialServicesInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetInputParameters

`func (o *GetSingleShipmentRateSpecialServicesInner) GetInputParameters() []GetSingleShipmentRateSpecialServicesInnerInputParametersInner`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *GetSingleShipmentRateSpecialServicesInner) GetInputParametersOk() (*[]GetSingleShipmentRateSpecialServicesInnerInputParametersInner, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *GetSingleShipmentRateSpecialServicesInner) SetInputParameters(v []GetSingleShipmentRateSpecialServicesInnerInputParametersInner)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *GetSingleShipmentRateSpecialServicesInner) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetSpecialServiceId

`func (o *GetSingleShipmentRateSpecialServicesInner) GetSpecialServiceId() string`

GetSpecialServiceId returns the SpecialServiceId field if non-nil, zero value otherwise.

### GetSpecialServiceIdOk

`func (o *GetSingleShipmentRateSpecialServicesInner) GetSpecialServiceIdOk() (*string, bool)`

GetSpecialServiceIdOk returns a tuple with the SpecialServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServiceId

`func (o *GetSingleShipmentRateSpecialServicesInner) SetSpecialServiceId(v string)`

SetSpecialServiceId sets SpecialServiceId field to given value.

### HasSpecialServiceId

`func (o *GetSingleShipmentRateSpecialServicesInner) HasSpecialServiceId() bool`

HasSpecialServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


