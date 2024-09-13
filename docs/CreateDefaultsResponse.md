# CreateDefaultsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier used in Database for Default, for mapping purpose. | [optional] 
**DefaultId** | Pointer to **string** | Assigned unique identifier for the created Default. | [optional] 

## Methods

### NewCreateDefaultsResponse

`func NewCreateDefaultsResponse() *CreateDefaultsResponse`

NewCreateDefaultsResponse instantiates a new CreateDefaultsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDefaultsResponseWithDefaults

`func NewCreateDefaultsResponseWithDefaults() *CreateDefaultsResponse`

NewCreateDefaultsResponseWithDefaults instantiates a new CreateDefaultsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateDefaultsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateDefaultsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateDefaultsResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateDefaultsResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDefaultId

`func (o *CreateDefaultsResponse) GetDefaultId() string`

GetDefaultId returns the DefaultId field if non-nil, zero value otherwise.

### GetDefaultIdOk

`func (o *CreateDefaultsResponse) GetDefaultIdOk() (*string, bool)`

GetDefaultIdOk returns a tuple with the DefaultId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultId

`func (o *CreateDefaultsResponse) SetDefaultId(v string)`

SetDefaultId sets DefaultId field to given value.

### HasDefaultId

`func (o *CreateDefaultsResponse) HasDefaultId() bool`

HasDefaultId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


