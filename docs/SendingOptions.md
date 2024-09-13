# SendingOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** | Name of the carrier. E.g., FedEx. | [optional] 
**Service** | Pointer to **string** | Name of the carrier-based service. E.g., 2DA. | [optional] 
**CarrierAccounts** | Pointer to [**SendingOptionsCarrierAccounts**](SendingOptionsCarrierAccounts.md) |  | [optional] 
**LabelSize** | Pointer to **string** | Size of the label, e.g., DOC_4X6. | [optional] 
**LabelType** | Pointer to **string** | Type of the Label, e.g., Shipping_Label. | [optional] 
**LabelFormat** | Pointer to **string** | Format of the Label, e.g., PDF. | [optional] 
**FromAddress** | Pointer to [**FromAddressV2**](FromAddressV2.md) |  | [optional] 
**Parcel** | Pointer to [**ParcelV2**](ParcelV2.md) |  | [optional] 

## Methods

### NewSendingOptions

`func NewSendingOptions() *SendingOptions`

NewSendingOptions instantiates a new SendingOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingOptionsWithDefaults

`func NewSendingOptionsWithDefaults() *SendingOptions`

NewSendingOptionsWithDefaults instantiates a new SendingOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *SendingOptions) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SendingOptions) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SendingOptions) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SendingOptions) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetService

`func (o *SendingOptions) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SendingOptions) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SendingOptions) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SendingOptions) HasService() bool`

HasService returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *SendingOptions) GetCarrierAccounts() SendingOptionsCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *SendingOptions) GetCarrierAccountsOk() (*SendingOptionsCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *SendingOptions) SetCarrierAccounts(v SendingOptionsCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *SendingOptions) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetLabelSize

`func (o *SendingOptions) GetLabelSize() string`

GetLabelSize returns the LabelSize field if non-nil, zero value otherwise.

### GetLabelSizeOk

`func (o *SendingOptions) GetLabelSizeOk() (*string, bool)`

GetLabelSizeOk returns a tuple with the LabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSize

`func (o *SendingOptions) SetLabelSize(v string)`

SetLabelSize sets LabelSize field to given value.

### HasLabelSize

`func (o *SendingOptions) HasLabelSize() bool`

HasLabelSize returns a boolean if a field has been set.

### GetLabelType

`func (o *SendingOptions) GetLabelType() string`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *SendingOptions) GetLabelTypeOk() (*string, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *SendingOptions) SetLabelType(v string)`

SetLabelType sets LabelType field to given value.

### HasLabelType

`func (o *SendingOptions) HasLabelType() bool`

HasLabelType returns a boolean if a field has been set.

### GetLabelFormat

`func (o *SendingOptions) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *SendingOptions) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *SendingOptions) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *SendingOptions) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetFromAddress

`func (o *SendingOptions) GetFromAddress() FromAddressV2`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *SendingOptions) GetFromAddressOk() (*FromAddressV2, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *SendingOptions) SetFromAddress(v FromAddressV2)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *SendingOptions) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *SendingOptions) GetParcel() ParcelV2`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *SendingOptions) GetParcelOk() (*ParcelV2, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *SendingOptions) SetParcel(v ParcelV2)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *SendingOptions) HasParcel() bool`

HasParcel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


