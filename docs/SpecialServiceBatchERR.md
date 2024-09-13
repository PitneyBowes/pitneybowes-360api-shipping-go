# SpecialServiceBatchERR

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputParameters** | Pointer to [**[]Parameter**](Parameter.md) | &gt;- These parameters are required to set for the special service, such as an Insurance value or a Receipt-number format. This is applicable only when Special Service requires input parameters. If a Special Service does not require input parameters, user can or pass an empty array. | [optional] 
**SpecialserviceId** | **string** | A unique identifier associated with the special service which varies based on selected USPS Service and ParcelTypes/PackageTypes. | 

## Methods

### NewSpecialServiceBatchERR

`func NewSpecialServiceBatchERR(specialserviceId string, ) *SpecialServiceBatchERR`

NewSpecialServiceBatchERR instantiates a new SpecialServiceBatchERR object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialServiceBatchERRWithDefaults

`func NewSpecialServiceBatchERRWithDefaults() *SpecialServiceBatchERR`

NewSpecialServiceBatchERRWithDefaults instantiates a new SpecialServiceBatchERR object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputParameters

`func (o *SpecialServiceBatchERR) GetInputParameters() []Parameter`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *SpecialServiceBatchERR) GetInputParametersOk() (*[]Parameter, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *SpecialServiceBatchERR) SetInputParameters(v []Parameter)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *SpecialServiceBatchERR) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetSpecialserviceId

`func (o *SpecialServiceBatchERR) GetSpecialserviceId() string`

GetSpecialserviceId returns the SpecialserviceId field if non-nil, zero value otherwise.

### GetSpecialserviceIdOk

`func (o *SpecialServiceBatchERR) GetSpecialserviceIdOk() (*string, bool)`

GetSpecialserviceIdOk returns a tuple with the SpecialserviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialserviceId

`func (o *SpecialServiceBatchERR) SetSpecialserviceId(v string)`

SetSpecialserviceId sets SpecialserviceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


