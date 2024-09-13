# MultipieceShipment200Response

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

### NewMultipieceShipment200Response

`func NewMultipieceShipment200Response() *MultipieceShipment200Response`

NewMultipieceShipment200Response instantiates a new MultipieceShipment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMultipieceShipment200ResponseWithDefaults

`func NewMultipieceShipment200ResponseWithDefaults() *MultipieceShipment200Response`

NewMultipieceShipment200ResponseWithDefaults instantiates a new MultipieceShipment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCorrelationId

`func (o *MultipieceShipment200Response) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *MultipieceShipment200Response) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *MultipieceShipment200Response) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.

### HasCorrelationId

`func (o *MultipieceShipment200Response) HasCorrelationId() bool`

HasCorrelationId returns a boolean if a field has been set.

### GetLabelLayout

`func (o *MultipieceShipment200Response) GetLabelLayout() []MultipieceDomesticShipmentResponseLabelLayoutInner`

GetLabelLayout returns the LabelLayout field if non-nil, zero value otherwise.

### GetLabelLayoutOk

`func (o *MultipieceShipment200Response) GetLabelLayoutOk() (*[]MultipieceDomesticShipmentResponseLabelLayoutInner, bool)`

GetLabelLayoutOk returns a tuple with the LabelLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelLayout

`func (o *MultipieceShipment200Response) SetLabelLayout(v []MultipieceDomesticShipmentResponseLabelLayoutInner)`

SetLabelLayout sets LabelLayout field to given value.

### HasLabelLayout

`func (o *MultipieceShipment200Response) HasLabelLayout() bool`

HasLabelLayout returns a boolean if a field has been set.

### GetFromAddress

`func (o *MultipieceShipment200Response) GetFromAddress() MultipieceDomesticShipmentRequestFromAddress`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *MultipieceShipment200Response) GetFromAddressOk() (*MultipieceDomesticShipmentRequestFromAddress, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *MultipieceShipment200Response) SetFromAddress(v MultipieceDomesticShipmentRequestFromAddress)`

SetFromAddress sets FromAddress field to given value.

### HasFromAddress

`func (o *MultipieceShipment200Response) HasFromAddress() bool`

HasFromAddress returns a boolean if a field has been set.

### GetMultiPieceRates

`func (o *MultipieceShipment200Response) GetMultiPieceRates() []MultipieceDomesticShipmentResponseMultiPieceRatesInner`

GetMultiPieceRates returns the MultiPieceRates field if non-nil, zero value otherwise.

### GetMultiPieceRatesOk

`func (o *MultipieceShipment200Response) GetMultiPieceRatesOk() (*[]MultipieceDomesticShipmentResponseMultiPieceRatesInner, bool)`

GetMultiPieceRatesOk returns a tuple with the MultiPieceRates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiPieceRates

`func (o *MultipieceShipment200Response) SetMultiPieceRates(v []MultipieceDomesticShipmentResponseMultiPieceRatesInner)`

SetMultiPieceRates sets MultiPieceRates field to given value.

### HasMultiPieceRates

`func (o *MultipieceShipment200Response) HasMultiPieceRates() bool`

HasMultiPieceRates returns a boolean if a field has been set.

### GetParcelTrackingNumber

`func (o *MultipieceShipment200Response) GetParcelTrackingNumber() string`

GetParcelTrackingNumber returns the ParcelTrackingNumber field if non-nil, zero value otherwise.

### GetParcelTrackingNumberOk

`func (o *MultipieceShipment200Response) GetParcelTrackingNumberOk() (*string, bool)`

GetParcelTrackingNumberOk returns a tuple with the ParcelTrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelTrackingNumber

`func (o *MultipieceShipment200Response) SetParcelTrackingNumber(v string)`

SetParcelTrackingNumber sets ParcelTrackingNumber field to given value.

### HasParcelTrackingNumber

`func (o *MultipieceShipment200Response) HasParcelTrackingNumber() bool`

HasParcelTrackingNumber returns a boolean if a field has been set.

### GetShipmentId

`func (o *MultipieceShipment200Response) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *MultipieceShipment200Response) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *MultipieceShipment200Response) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.

### HasShipmentId

`func (o *MultipieceShipment200Response) HasShipmentId() bool`

HasShipmentId returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *MultipieceShipment200Response) GetShipmentOptions() MultipieceDomesticShipmentRequestShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *MultipieceShipment200Response) GetShipmentOptionsOk() (*MultipieceDomesticShipmentRequestShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *MultipieceShipment200Response) SetShipmentOptions(v MultipieceDomesticShipmentRequestShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *MultipieceShipment200Response) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetToAddress

`func (o *MultipieceShipment200Response) GetToAddress() MultipieceDomesticShipmentRequestToAddress`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *MultipieceShipment200Response) GetToAddressOk() (*MultipieceDomesticShipmentRequestToAddress, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *MultipieceShipment200Response) SetToAddress(v MultipieceDomesticShipmentRequestToAddress)`

SetToAddress sets ToAddress field to given value.

### HasToAddress

`func (o *MultipieceShipment200Response) HasToAddress() bool`

HasToAddress returns a boolean if a field has been set.

### GetCustoms

`func (o *MultipieceShipment200Response) GetCustoms() MultipieceInternationalShipmentResponseCustoms`

GetCustoms returns the Customs field if non-nil, zero value otherwise.

### GetCustomsOk

`func (o *MultipieceShipment200Response) GetCustomsOk() (*MultipieceInternationalShipmentResponseCustoms, bool)`

GetCustomsOk returns a tuple with the Customs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustoms

`func (o *MultipieceShipment200Response) SetCustoms(v MultipieceInternationalShipmentResponseCustoms)`

SetCustoms sets Customs field to given value.

### HasCustoms

`func (o *MultipieceShipment200Response) HasCustoms() bool`

HasCustoms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


