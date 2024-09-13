# SpecialServiceBatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InputParameters** | Pointer to [**[]Parameter**](Parameter.md) | &gt;- The parameters to set for the special service, such as an insurance value or a receipt-number format. This is required if the special service requires input parameters. If a special service does not require input parameters, you can either leave out the array or pass an empty array. | [optional] 
**SpecialserviceId** | **string** | A unique identifier associated with the special service to be applied. | 

## Methods

### NewSpecialServiceBatch

`func NewSpecialServiceBatch(specialserviceId string, ) *SpecialServiceBatch`

NewSpecialServiceBatch instantiates a new SpecialServiceBatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpecialServiceBatchWithDefaults

`func NewSpecialServiceBatchWithDefaults() *SpecialServiceBatch`

NewSpecialServiceBatchWithDefaults instantiates a new SpecialServiceBatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInputParameters

`func (o *SpecialServiceBatch) GetInputParameters() []Parameter`

GetInputParameters returns the InputParameters field if non-nil, zero value otherwise.

### GetInputParametersOk

`func (o *SpecialServiceBatch) GetInputParametersOk() (*[]Parameter, bool)`

GetInputParametersOk returns a tuple with the InputParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputParameters

`func (o *SpecialServiceBatch) SetInputParameters(v []Parameter)`

SetInputParameters sets InputParameters field to given value.

### HasInputParameters

`func (o *SpecialServiceBatch) HasInputParameters() bool`

HasInputParameters returns a boolean if a field has been set.

### GetSpecialserviceId

`func (o *SpecialServiceBatch) GetSpecialserviceId() string`

GetSpecialserviceId returns the SpecialserviceId field if non-nil, zero value otherwise.

### GetSpecialserviceIdOk

`func (o *SpecialServiceBatch) GetSpecialserviceIdOk() (*string, bool)`

GetSpecialserviceIdOk returns a tuple with the SpecialserviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialserviceId

`func (o *SpecialServiceBatch) SetSpecialserviceId(v string)`

SetSpecialserviceId sets SpecialserviceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


