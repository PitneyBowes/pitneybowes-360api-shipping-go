# GetDefaultsById

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the created Default. For example: Domestic. | [optional] 
**DefaultID** | Pointer to **string** | Unique identifier assigned to the Default while its creation using CreateDefaults API. | [optional] 
**CreatedDate** | Pointer to **time.Time** | The timestamp when the Default is created. | [optional] 
**UpdatedDate** | Pointer to **time.Time** | The timestamp when the Default is updated. | [optional] 
**SendingOptions** | Pointer to [**SendingOptions**](SendingOptions.md) |  | [optional] 

## Methods

### NewGetDefaultsById

`func NewGetDefaultsById() *GetDefaultsById`

NewGetDefaultsById instantiates a new GetDefaultsById object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDefaultsByIdWithDefaults

`func NewGetDefaultsByIdWithDefaults() *GetDefaultsById`

NewGetDefaultsByIdWithDefaults instantiates a new GetDefaultsById object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetDefaultsById) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetDefaultsById) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetDefaultsById) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetDefaultsById) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDefaultID

`func (o *GetDefaultsById) GetDefaultID() string`

GetDefaultID returns the DefaultID field if non-nil, zero value otherwise.

### GetDefaultIDOk

`func (o *GetDefaultsById) GetDefaultIDOk() (*string, bool)`

GetDefaultIDOk returns a tuple with the DefaultID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultID

`func (o *GetDefaultsById) SetDefaultID(v string)`

SetDefaultID sets DefaultID field to given value.

### HasDefaultID

`func (o *GetDefaultsById) HasDefaultID() bool`

HasDefaultID returns a boolean if a field has been set.

### GetCreatedDate

`func (o *GetDefaultsById) GetCreatedDate() time.Time`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *GetDefaultsById) GetCreatedDateOk() (*time.Time, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *GetDefaultsById) SetCreatedDate(v time.Time)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *GetDefaultsById) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetUpdatedDate

`func (o *GetDefaultsById) GetUpdatedDate() time.Time`

GetUpdatedDate returns the UpdatedDate field if non-nil, zero value otherwise.

### GetUpdatedDateOk

`func (o *GetDefaultsById) GetUpdatedDateOk() (*time.Time, bool)`

GetUpdatedDateOk returns a tuple with the UpdatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedDate

`func (o *GetDefaultsById) SetUpdatedDate(v time.Time)`

SetUpdatedDate sets UpdatedDate field to given value.

### HasUpdatedDate

`func (o *GetDefaultsById) HasUpdatedDate() bool`

HasUpdatedDate returns a boolean if a field has been set.

### GetSendingOptions

`func (o *GetDefaultsById) GetSendingOptions() SendingOptions`

GetSendingOptions returns the SendingOptions field if non-nil, zero value otherwise.

### GetSendingOptionsOk

`func (o *GetDefaultsById) GetSendingOptionsOk() (*SendingOptions, bool)`

GetSendingOptionsOk returns a tuple with the SendingOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendingOptions

`func (o *GetDefaultsById) SetSendingOptions(v SendingOptions)`

SetSendingOptions sets SendingOptions field to given value.

### HasSendingOptions

`func (o *GetDefaultsById) HasSendingOptions() bool`

HasSendingOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


