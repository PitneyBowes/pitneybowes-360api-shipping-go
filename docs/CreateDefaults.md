# CreateDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the Default, e.g., FedEx-SAKS. | [optional] 
**DefaultID** | Pointer to **string** | A unique identifier to be assigned to the Default.  | [optional] 
**SendingOptions** | Pointer to [**SendingOptions**](SendingOptions.md) |  | [optional] 

## Methods

### NewCreateDefaults

`func NewCreateDefaults() *CreateDefaults`

NewCreateDefaults instantiates a new CreateDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDefaultsWithDefaults

`func NewCreateDefaultsWithDefaults() *CreateDefaults`

NewCreateDefaultsWithDefaults instantiates a new CreateDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateDefaults) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateDefaults) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateDefaults) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateDefaults) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultID

`func (o *CreateDefaults) GetDefaultID() string`

GetDefaultID returns the DefaultID field if non-nil, zero value otherwise.

### GetDefaultIDOk

`func (o *CreateDefaults) GetDefaultIDOk() (*string, bool)`

GetDefaultIDOk returns a tuple with the DefaultID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultID

`func (o *CreateDefaults) SetDefaultID(v string)`

SetDefaultID sets DefaultID field to given value.

### HasDefaultID

`func (o *CreateDefaults) HasDefaultID() bool`

HasDefaultID returns a boolean if a field has been set.

### GetSendingOptions

`func (o *CreateDefaults) GetSendingOptions() SendingOptions`

GetSendingOptions returns the SendingOptions field if non-nil, zero value otherwise.

### GetSendingOptionsOk

`func (o *CreateDefaults) GetSendingOptionsOk() (*SendingOptions, bool)`

GetSendingOptionsOk returns a tuple with the SendingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingOptions

`func (o *CreateDefaults) SetSendingOptions(v SendingOptions)`

SetSendingOptions sets SendingOptions field to given value.

### HasSendingOptions

`func (o *CreateDefaults) HasSendingOptions() bool`

HasSendingOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


