# AddressSuggestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**AddressSuggestResponseAddress**](AddressSuggestResponseAddress.md) |  | [optional] 
**Suggestions** | Pointer to [**AddressSuggestResponseSuggestions**](AddressSuggestResponseSuggestions.md) |  | [optional] 

## Methods

### NewAddressSuggestResponse

`func NewAddressSuggestResponse() *AddressSuggestResponse`

NewAddressSuggestResponse instantiates a new AddressSuggestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressSuggestResponseWithDefaults

`func NewAddressSuggestResponseWithDefaults() *AddressSuggestResponse`

NewAddressSuggestResponseWithDefaults instantiates a new AddressSuggestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AddressSuggestResponse) GetAddress() AddressSuggestResponseAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AddressSuggestResponse) GetAddressOk() (*AddressSuggestResponseAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AddressSuggestResponse) SetAddress(v AddressSuggestResponseAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AddressSuggestResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetSuggestions

`func (o *AddressSuggestResponse) GetSuggestions() AddressSuggestResponseSuggestions`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *AddressSuggestResponse) GetSuggestionsOk() (*AddressSuggestResponseSuggestions, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *AddressSuggestResponse) SetSuggestions(v AddressSuggestResponseSuggestions)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *AddressSuggestResponse) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


