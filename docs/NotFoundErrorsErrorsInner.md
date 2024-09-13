# NotFoundErrorsErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **string** | Error code(s) that appear due to HTTP 404 Page or File not found | [optional] 
**ErrorDescription** | Pointer to **string** | The HTTP 404 Not Found response status code indicates that the server cannot find the requested resource. | [optional] 
**AdditionalCode** | Pointer to **string** | A unique identifier for the error, for example 1101055, 0100008, or 1021126. | [optional] 
**AdditionalInfo** | Pointer to **string** | This is an additional information about the error. This error &#39;Invalid Request&#39; might appear due to invalid dimension, weight, or serviceid, or if the information is missing. | [optional] 
**AdditionalParameters** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNotFoundErrorsErrorsInner

`func NewNotFoundErrorsErrorsInner() *NotFoundErrorsErrorsInner`

NewNotFoundErrorsErrorsInner instantiates a new NotFoundErrorsErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotFoundErrorsErrorsInnerWithDefaults

`func NewNotFoundErrorsErrorsInnerWithDefaults() *NotFoundErrorsErrorsInner`

NewNotFoundErrorsErrorsInnerWithDefaults instantiates a new NotFoundErrorsErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *NotFoundErrorsErrorsInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *NotFoundErrorsErrorsInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *NotFoundErrorsErrorsInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *NotFoundErrorsErrorsInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDescription

`func (o *NotFoundErrorsErrorsInner) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *NotFoundErrorsErrorsInner) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *NotFoundErrorsErrorsInner) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *NotFoundErrorsErrorsInner) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.

### GetAdditionalCode

`func (o *NotFoundErrorsErrorsInner) GetAdditionalCode() string`

GetAdditionalCode returns the AdditionalCode field if non-nil, zero value otherwise.

### GetAdditionalCodeOk

`func (o *NotFoundErrorsErrorsInner) GetAdditionalCodeOk() (*string, bool)`

GetAdditionalCodeOk returns a tuple with the AdditionalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalCode

`func (o *NotFoundErrorsErrorsInner) SetAdditionalCode(v string)`

SetAdditionalCode sets AdditionalCode field to given value.

### HasAdditionalCode

`func (o *NotFoundErrorsErrorsInner) HasAdditionalCode() bool`

HasAdditionalCode returns a boolean if a field has been set.

### GetAdditionalInfo

`func (o *NotFoundErrorsErrorsInner) GetAdditionalInfo() string`

GetAdditionalInfo returns the AdditionalInfo field if non-nil, zero value otherwise.

### GetAdditionalInfoOk

`func (o *NotFoundErrorsErrorsInner) GetAdditionalInfoOk() (*string, bool)`

GetAdditionalInfoOk returns a tuple with the AdditionalInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInfo

`func (o *NotFoundErrorsErrorsInner) SetAdditionalInfo(v string)`

SetAdditionalInfo sets AdditionalInfo field to given value.

### HasAdditionalInfo

`func (o *NotFoundErrorsErrorsInner) HasAdditionalInfo() bool`

HasAdditionalInfo returns a boolean if a field has been set.

### GetAdditionalParameters

`func (o *NotFoundErrorsErrorsInner) GetAdditionalParameters() []string`

GetAdditionalParameters returns the AdditionalParameters field if non-nil, zero value otherwise.

### GetAdditionalParametersOk

`func (o *NotFoundErrorsErrorsInner) GetAdditionalParametersOk() (*[]string, bool)`

GetAdditionalParametersOk returns a tuple with the AdditionalParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalParameters

`func (o *NotFoundErrorsErrorsInner) SetAdditionalParameters(v []string)`

SetAdditionalParameters sets AdditionalParameters field to given value.

### HasAdditionalParameters

`func (o *NotFoundErrorsErrorsInner) HasAdditionalParameters() bool`

HasAdditionalParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


