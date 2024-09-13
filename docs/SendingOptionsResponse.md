# SendingOptionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Carrier** | Pointer to **string** | Name of the carrier. E.g., FedEx. | [optional] 
**Service** | Pointer to **string** | Name of the carrier-based service. E.g., 2DA. | [optional] 
**CarrierAccounts** | Pointer to [**SendingOptionsResponseCarrierAccounts**](SendingOptionsResponseCarrierAccounts.md) |  | [optional] 
**LabelSize** | Pointer to **string** | Size of the label, e.g., DOC_4X6. | [optional] 
**LabelType** | Pointer to **string** | Type of the Label, e.g., Shipping_Label. | [optional] 
**LabelFormat** | Pointer to **string** | Format of the Label, e.g., PDF. | [optional] 
**FromAddress** | Pointer to [**FromAddressV2Response**](FromAddressV2Response.md) |  | [optional] 
**Parcel** | Pointer to [**ParcelV2Response**](ParcelV2Response.md) |  | [optional] 

## Methods

### NewSendingOptionsResponse

`func NewSendingOptionsResponse() *SendingOptionsResponse`

NewSendingOptionsResponse instantiates a new SendingOptionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendingOptionsResponseWithDefaults

`func NewSendingOptionsResponseWithDefaults() *SendingOptionsResponse`

NewSendingOptionsResponseWithDefaults instantiates a new SendingOptionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrier

`func (o *SendingOptionsResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *SendingOptionsResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *SendingOptionsResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *SendingOptionsResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### GetService

`func (o *SendingOptionsResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *SendingOptionsResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *SendingOptionsResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *SendingOptionsResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### GetCarrierAccounts

`func (o *SendingOptionsResponse) GetCarrierAccounts() SendingOptionsResponseCarrierAccounts`

GetCarrierAccounts returns the CarrierAccounts field if non-nil, zero value otherwise.

### GetCarrierAccountsOk

`func (o *SendingOptionsResponse) GetCarrierAccountsOk() (*SendingOptionsResponseCarrierAccounts, bool)`

GetCarrierAccountsOk returns a tuple with the CarrierAccounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierAccounts

`func (o *SendingOptionsResponse) SetCarrierAccounts(v SendingOptionsResponseCarrierAccounts)`

SetCarrierAccounts sets CarrierAccounts field to given value.

### HasCarrierAccounts

`func (o *SendingOptionsResponse) HasCarrierAccounts() bool`

HasCarrierAccounts returns a boolean if a field has been set.

### GetLabelSize

`func (o *SendingOptionsResponse) GetLabelSize() string`

GetLabelSize returns the LabelSize field if non-nil, zero value otherwise.

### GetLabelSizeOk

`func (o *SendingOptionsResponse) GetLabelSizeOk() (*string, bool)`

GetLabelSizeOk returns a tuple with the LabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSize

`func (o *SendingOptionsResponse) SetLabelSize(v string)`

SetLabelSize sets LabelSize field to given value.

### HasLabelSize

`func (o *SendingOptionsResponse) HasLabelSize() bool`

HasLabelSize returns a boolean if a field has been set.

### GetLabelType

`func (o *SendingOptionsResponse) GetLabelType() string`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *SendingOptionsResponse) GetLabelTypeOk() (*string, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *SendingOptionsResponse) SetLabelType(v string)`

SetLabelType sets LabelType field to given value.

### HasLabelType

`func (o *SendingOptionsResponse) HasLabelType() bool`

HasLabelType returns a boolean if a field has been set.

### GetLabelFormat

`func (o *SendingOptionsResponse) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *SendingOptionsResponse) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *SendingOptionsResponse) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.

### HasLabelFormat

`func (o *SendingOptionsResponse) HasLabelFormat() bool`

HasLabelFormat returns a boolean if a field has been set.

### GetFromAddress

`func (o *SendingOptionsResponse) GetFromAddress() FromAddressV2Response`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *SendingOptionsResponse) GetFromAddressOk() (*FromAddressV2Response, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *SendingOptionsResponse) SetFromAddress(v FromAddressV2Response)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *SendingOptionsResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetParcel

`func (o *SendingOptionsResponse) GetParcel() ParcelV2Response`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *SendingOptionsResponse) GetParcelOk() (*ParcelV2Response, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *SendingOptionsResponse) SetParcel(v ParcelV2Response)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *SendingOptionsResponse) HasParcel() bool`

HasParcel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


