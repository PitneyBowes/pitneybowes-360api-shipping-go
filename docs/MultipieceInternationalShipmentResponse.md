# MultipieceInternationalShipmentResponse

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
**Customs** | Pointer to [**MultipieceInternationalShipmentResponseCustoms**](MultipieceInternationalShipmentResponseCustoms.md) |  | [optional] 

## Methods

### NewMultipieceInternationalShipmentResponse

`func NewMultipieceInternationalShipmentResponse() *MultipieceInternationalShipmentResponse`

NewMultipieceInternationalShipmentResponse instantiates a new MultipieceInternationalShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceInternationalShipmentResponseWithDefaults

`func NewMultipieceInternationalShipmentResponseWithDefaults() *MultipieceInternationalShipmentResponse`

NewMultipieceInternationalShipmentResponseWithDefaults instantiates a new MultipieceInternationalShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *MultipieceInternationalShipmentResponse) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MultipieceInternationalShipmentResponse) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MultipieceInternationalShipmentResponse) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MultipieceInternationalShipmentResponse) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetLabelLayout

`func (o *MultipieceInternationalShipmentResponse) GetLabelLayout() []MultipieceDomesticShipmentResponseLabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *MultipieceInternationalShipmentResponse) GetLabelLayoutOk() (*[]MultipieceDomesticShipmentResponseLabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *MultipieceInternationalShipmentResponse) SetLabelLayout(v []MultipieceDomesticShipmentResponseLabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *MultipieceInternationalShipmentResponse) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceInternationalShipmentResponse) GetFromAddress() MultipieceDomesticShipmentRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceInternationalShipmentResponse) GetFromAddressOk() (*MultipieceDomesticShipmentRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceInternationalShipmentResponse) SetFromAddress(v MultipieceDomesticShipmentRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceInternationalShipmentResponse) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetMultiPieceRates

`func (o *MultipieceInternationalShipmentResponse) GetMultiPieceRates() []MultipieceDomesticShipmentResponseMultiPieceRatesInner`

GetMultiPieceRates returns the MultiPieceRates field if non-nil, zero value otherwise.

### GetMultiPieceRatesOk

`func (o *MultipieceInternationalShipmentResponse) GetMultiPieceRatesOk() (*[]MultipieceDomesticShipmentResponseMultiPieceRatesInner, bool)`

GetMultiPieceRatesOk returns a tuple with the MultiPieceRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceRates

`func (o *MultipieceInternationalShipmentResponse) SetMultiPieceRates(v []MultipieceDomesticShipmentResponseMultiPieceRatesInner)`

SetMultiPieceRates sets MultiPieceRates field to given value.

### HasMultiPieceRates

`func (o *MultipieceInternationalShipmentResponse) HasMultiPieceRates() bool`

HasMultiPieceRates returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *MultipieceInternationalShipmentResponse) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *MultipieceInternationalShipmentResponse) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *MultipieceInternationalShipmentResponse) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *MultipieceInternationalShipmentResponse) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetShipmentId

`func (o *MultipieceInternationalShipmentResponse) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *MultipieceInternationalShipmentResponse) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *MultipieceInternationalShipmentResponse) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *MultipieceInternationalShipmentResponse) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceInternationalShipmentResponse) GetShipmentOptions() MultipieceDomesticShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceInternationalShipmentResponse) GetShipmentOptionsOk() (*MultipieceDomesticShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceInternationalShipmentResponse) SetShipmentOptions(v MultipieceDomesticShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceInternationalShipmentResponse) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceInternationalShipmentResponse) GetToAddress() MultipieceDomesticShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceInternationalShipmentResponse) GetToAddressOk() (*MultipieceDomesticShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceInternationalShipmentResponse) SetToAddress(v MultipieceDomesticShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceInternationalShipmentResponse) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetCustoms

`func (o *MultipieceInternationalShipmentResponse) GetCustoms() MultipieceInternationalShipmentResponseCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *MultipieceInternationalShipmentResponse) GetCustomsOk() (*MultipieceInternationalShipmentResponseCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *MultipieceInternationalShipmentResponse) SetCustoms(v MultipieceInternationalShipmentResponseCustoms)`

SetCustoms sets Customs field to given value.

### HasCustoms

`func (o *MultipieceInternationalShipmentResponse) HasCustoms() bool`

HasCustoms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


