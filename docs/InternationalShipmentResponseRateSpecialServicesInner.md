# InternationalShipmentResponseRateSpecialServicesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fee** | Pointer to **float32** | The amount of the special service. | [optional] 
**InputParameters** | Pointer to [**[]InternationalShipmentResponseRateSpecialServicesInnerInputParametersInner**](InternationalShipmentResponseRateSpecialServicesInnerInputParametersInner.md) | &gt;-The parameters to set for the special service, such as an insurance value or a receipt-number format. This is required if the specialservice requires input parameters. If a special service does not require input parameters, you can either leave out the array or pass an empty array. | [optional] 
**SpecialServiceId** | Pointer to **string** | This is the unique identifier given to various specialservice, which is used while Rating. | [optional] 

## Methods

### NewInternationalShipmentResponseRateSpecialServicesInner

`func NewInternationalShipmentResponseRateSpecialServicesInner() *InternationalShipmentResponseRateSpecialServicesInner`

NewInternationalShipmentResponseRateSpecialServicesInner instantiates a new InternationalShipmentResponseRateSpecialServicesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalShipmentResponseRateSpecialServicesInnerWithDefaults

`func NewInternationalShipmentResponseRateSpecialServicesInnerWithDefaults() *InternationalShipmentResponseRateSpecialServicesInner`

NewInternationalShipmentResponseRateSpecialServicesInnerWithDefaults instantiates a new InternationalShipmentResponseRateSpecialServicesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFee

`func (o *InternationalShipmentResponseRateSpecialServicesInner) GetFee() float32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *InternationalShipmentResponseRateSpecialServicesInner) GetFeeOk() (*float32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *InternationalShipmentResponseRateSpecialServicesInner) SetFee(v float32)`

SetFee sets Fee field to given value.

### HasFee

`func (o *InternationalShipmentResponseRateSpecialServicesInner) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetInputParameters

`func (o *InternationalShipmentResponseRateSpecialServicesInner) GetInputParameters() []InternationalShipmentResponseRateSpecialServicesInnerInputParametersInner`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *InternationalShipmentResponseRateSpecialServicesInner) GetInputParametersOk() (*[]InternationalShipmentResponseRateSpecialServicesInnerInputParametersInner, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *InternationalShipmentResponseRateSpecialServicesInner) SetInputParameters(v []InternationalShipmentResponseRateSpecialServicesInnerInputParametersInner)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *InternationalShipmentResponseRateSpecialServicesInner) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetSpecialServiceId

`func (o *InternationalShipmentResponseRateSpecialServicesInner) GetSpecialServiceId() string`

GetSpecialServiceId returns the SpecialServiceId field if non-nil, zero value otherwise.

### GetSpecialServiceIdOk

`func (o *InternationalShipmentResponseRateSpecialServicesInner) GetSpecialServiceIdOk() (*string, bool)`

GetSpecialServiceIdOk returns a tuple with the SpecialServiceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialServiceId

`func (o *InternationalShipmentResponseRateSpecialServicesInner) SetSpecialServiceId(v string)`

SetSpecialServiceId sets SpecialServiceId field to given value.

### HasSpecialServiceId

`func (o *InternationalShipmentResponseRateSpecialServicesInner) HasSpecialServiceId() bool`

HasSpecialServiceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


