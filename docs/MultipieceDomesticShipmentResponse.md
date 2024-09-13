# MultipieceDomesticShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CorrelationId** | Pointer to **string** | description | [optional] 
**LabelLayout** | Pointer to [**[]MultipieceDomesticShipmentResponseLabelLayoutInner**](MultipieceDomesticShipmentResponseLabelLayoutInner.md) | description | [optional] 
**FromAddress** | Pointer to [**MultipieceDomesticShipmentRequestFromAddress**](MultipieceDomesticShipmentRequestFromAddress.md) |  | [optional] 
**MultiPieceRates** | Pointer to [**[]MultipieceDomesticShipmentResponseMultiPieceRatesInner**](MultipieceDomesticShipmentResponseMultiPieceRatesInner.md) | description | [optional] 
**ParcelTrackingNumber** | Pointer to **string** | description | [optional] 
**ShipmentId** | Pointer to **string** | description | [optional] 
**ShipmentOptions** | Pointer to [**MultipieceDomesticShipmentRequestShipmentOptions**](MultipieceDomesticShipmentRequestShipmentOptions.md) |  | [optional] 
**ToAddress** | Pointer to [**MultipieceDomesticShipmentRequestToAddress**](MultipieceDomesticShipmentRequestToAddress.md) |  | [optional] 

## Methods

### NewMultipieceDomesticShipmentResponse

`func NewMultipieceDomesticShipmentResponse() *MultipieceDomesticShipmentResponse`

NewMultipieceDomesticShipmentResponse instantiates a new MultipieceDomesticShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceDomesticShipmentResponseWithDefaults

`func NewMultipieceDomesticShipmentResponseWithDefaults() *MultipieceDomesticShipmentResponse`

NewMultipieceDomesticShipmentResponseWithDefaults instantiates a new MultipieceDomesticShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *MultipieceDomesticShipmentResponse) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MultipieceDomesticShipmentResponse) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MultipieceDomesticShipmentResponse) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MultipieceDomesticShipmentResponse) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetLabelLayout

`func (o *MultipieceDomesticShipmentResponse) GetLabelLayout() []MultipieceDomesticShipmentResponseLabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *MultipieceDomesticShipmentResponse) GetLabelLayoutOk() (*[]MultipieceDomesticShipmentResponseLabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *MultipieceDomesticShipmentResponse) SetLabelLayout(v []MultipieceDomesticShipmentResponseLabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *MultipieceDomesticShipmentResponse) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceDomesticShipmentResponse) GetFromAddress() MultipieceDomesticShipmentRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceDomesticShipmentResponse) GetFromAddressOk() (*MultipieceDomesticShipmentRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceDomesticShipmentResponse) SetFromAddress(v MultipieceDomesticShipmentRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceDomesticShipmentResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetMultiPieceRates

`func (o *MultipieceDomesticShipmentResponse) GetMultiPieceRates() []MultipieceDomesticShipmentResponseMultiPieceRatesInner`

GetMultiPieceRates returns the MultiPieceRates field if non-nil, zero value otherwise.

### GetMultiPieceRatesOk

`func (o *MultipieceDomesticShipmentResponse) GetMultiPieceRatesOk() (*[]MultipieceDomesticShipmentResponseMultiPieceRatesInner, bool)`

GetMultiPieceRatesOk returns a tuple with the MultiPieceRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceRates

`func (o *MultipieceDomesticShipmentResponse) SetMultiPieceRates(v []MultipieceDomesticShipmentResponseMultiPieceRatesInner)`

SetMultiPieceRates sets MultiPieceRates field to given value.

### HasMultiPieceRates

`func (o *MultipieceDomesticShipmentResponse) HasMultiPieceRates() bool`

HasMultiPieceRates returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *MultipieceDomesticShipmentResponse) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *MultipieceDomesticShipmentResponse) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *MultipieceDomesticShipmentResponse) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *MultipieceDomesticShipmentResponse) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetShipmentId

`func (o *MultipieceDomesticShipmentResponse) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *MultipieceDomesticShipmentResponse) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *MultipieceDomesticShipmentResponse) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *MultipieceDomesticShipmentResponse) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceDomesticShipmentResponse) GetShipmentOptions() MultipieceDomesticShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceDomesticShipmentResponse) GetShipmentOptionsOk() (*MultipieceDomesticShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceDomesticShipmentResponse) SetShipmentOptions(v MultipieceDomesticShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceDomesticShipmentResponse) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceDomesticShipmentResponse) GetToAddress() MultipieceDomesticShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceDomesticShipmentResponse) GetToAddressOk() (*MultipieceDomesticShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceDomesticShipmentResponse) SetToAddress(v MultipieceDomesticShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceDomesticShipmentResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


