# AddressSuggestResponseSuggestions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SuggestionType** | Pointer to **string** | The part of the address that is changed in the accompanying suggestions. Possible values are: - street - streetPrimaryRange (e.g., street number, PO Box number) - secondaryRange (e.g., suite number, apartment number) - PO BoxPrimaryRange - ruralRouteUnitRange - highwayContractUnitRange - city - company - puertoricanUrbanization | [optional] 
**Addresses** | Pointer to [**[]AddressSuggestResponseSuggestionsAddressesInner**](AddressSuggestResponseSuggestionsAddressesInner.md) |  | [optional] 

## Methods

### NewAddressSuggestResponseSuggestions

`func NewAddressSuggestResponseSuggestions() *AddressSuggestResponseSuggestions`

NewAddressSuggestResponseSuggestions instantiates a new AddressSuggestResponseSuggestions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressSuggestResponseSuggestionsWithDefaults

`func NewAddressSuggestResponseSuggestionsWithDefaults() *AddressSuggestResponseSuggestions`

NewAddressSuggestResponseSuggestionsWithDefaults instantiates a new AddressSuggestResponseSuggestions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggestionType

`func (o *AddressSuggestResponseSuggestions) GetSuggestionType() string`

GetSuggestionType returns the SuggestionType field if non-nil, zero value otherwise.

### GetSuggestionTypeOk

`func (o *AddressSuggestResponseSuggestions) GetSuggestionTypeOk() (*string, bool)`

GetSuggestionTypeOk returns a tuple with the SuggestionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionType

`func (o *AddressSuggestResponseSuggestions) SetSuggestionType(v string)`

SetSuggestionType sets SuggestionType field to given value.

### HasSuggestionType

`func (o *AddressSuggestResponseSuggestions) HasSuggestionType() bool`

HasSuggestionType returns a boolean if a field has been set.

### GetAddresses

`func (o *AddressSuggestResponseSuggestions) GetAddresses() []AddressSuggestResponseSuggestionsAddressesInner`

GetAddresses returns the Addresses field if non-nil, zero value otherwise.

### GetAddressesOk

`func (o *AddressSuggestResponseSuggestions) GetAddressesOk() (*[]AddressSuggestResponseSuggestionsAddressesInner, bool)`

GetAddressesOk returns a tuple with the Addresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddresses

`func (o *AddressSuggestResponseSuggestions) SetAddresses(v []AddressSuggestResponseSuggestionsAddressesInner)`

SetAddresses sets Addresses field to given value.

### HasAddresses

`func (o *AddressSuggestResponseSuggestions) HasAddresses() bool`

HasAddresses returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


