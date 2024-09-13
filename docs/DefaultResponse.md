# DefaultResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the created Default. For example: Domestic. | [optional] 
**DefaultID** | Pointer to **string** | Unique identifier assigned to the Default while its creation using CreateDefaults API. | [optional] 
**CreatedDate** | Pointer to **string** | The timestamp when the Default is created. | [optional] 
**UpdatedDate** | Pointer to **string** | The timestamp when the Default is updated. | [optional] 
**SendingOptions** | Pointer to [**SendingOptionsResponse**](SendingOptionsResponse.md) |  | [optional] 

## Methods

### NewDefaultResponse

`func NewDefaultResponse() *DefaultResponse`

NewDefaultResponse instantiates a new DefaultResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultResponseWithDefaults

`func NewDefaultResponseWithDefaults() *DefaultResponse`

NewDefaultResponseWithDefaults instantiates a new DefaultResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DefaultResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DefaultResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DefaultResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DefaultResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultID

`func (o *DefaultResponse) GetDefaultID() string`

GetDefaultID returns the DefaultID field if non-nil, zero value otherwise.

### GetDefaultIDOk

`func (o *DefaultResponse) GetDefaultIDOk() (*string, bool)`

GetDefaultIDOk returns a tuple with the DefaultID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultID

`func (o *DefaultResponse) SetDefaultID(v string)`

SetDefaultID sets DefaultID field to given value.

### HasDefaultID

`func (o *DefaultResponse) HasDefaultID() bool`

HasDefaultID returns a boolean if a field has been set.

### GetCreatedDate

`func (o *DefaultResponse) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *DefaultResponse) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *DefaultResponse) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *DefaultResponse) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *DefaultResponse) GetUpdatedDate() string`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *DefaultResponse) GetUpdatedDateOk() (*string, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *DefaultResponse) SetUpdatedDate(v string)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *DefaultResponse) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetSendingOptions

`func (o *DefaultResponse) GetSendingOptions() SendingOptionsResponse`

GetSendingOptions returns the SendingOptions field if non-nil, zero value otherwise.

### GetSendingOptionsOk

`func (o *DefaultResponse) GetSendingOptionsOk() (*SendingOptionsResponse, bool)`

GetSendingOptionsOk returns a tuple with the SendingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingOptions

`func (o *DefaultResponse) SetSendingOptions(v SendingOptionsResponse)`

SetSendingOptions sets SendingOptions field to given value.

### HasSendingOptions

`func (o *DefaultResponse) HasSendingOptions() bool`

HasSendingOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


