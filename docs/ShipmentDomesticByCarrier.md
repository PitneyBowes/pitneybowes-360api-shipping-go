# ShipmentDomesticByCarrier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromAddress** | [**FromAddressV2**](FromAddressV2.md) |  | 
**ToAddress** | [**ToAddressV2**](ToAddressV2.md) |  | 
**Parcel** | Pointer to [**ParcelV2**](ParcelV2.md) |  | [optional] 
**ParcelType** | **string** | Parcel Type is required for creating a shipment while rating a parcel, which varies as per Carrier selection.&lt;br /&gt; ParcelType can have categories like Package, Envelopes, Paks, Boxes, Tube, etc. &lt;br /&gt; &#x60;Max length &#x3D; 30&#x60;  | 
**RateShopBy** | **string** | RateShop, which is attached to an Enterprise or Location, is done through three approaches: by Carrier, by RateGroup, and by Ruleset. &lt;br /&gt;  Through Carrier, customers can choose the carriers as per requirement, based on which services, parcel types, and special services can be selected, and RateShop is done. &lt;br /&gt; Through RateGroup, customers can select the RateGroup, which has been divided into two categories: Cheapest (w.r.t. price) and Fastest (w.r.t. delivery hours). &lt;br /&gt; Through Ruleset, customers can define the Condition/rule for selecting carriers and their services, so they do not need to worry for Rate Shopping every time they create Shipment. For example, For a particular location, they can set one definite carrier, or apply RateGroup - Cheapest/Fastest. Similarly, for a particular amount like below $1000 Dollars, they can select a definite carrier service, based on RateGroup. | 
**ByCarrier** | Pointer to [**ByCarrierV2**](ByCarrierV2.md) |  | [optional] 
**ShipmentOptions** | Pointer to [**ShipmentDomesticByCarrierShipmentOptions**](ShipmentDomesticByCarrierShipmentOptions.md) |  | [optional] 
**DocTab** | Pointer to [**ShipmentDomesticByCarrierDocTab**](ShipmentDomesticByCarrierDocTab.md) |  | [optional] 
**DeliveryConfirmation** | Pointer to **string** | Indicates the supporting special service or document as an evidence of shipment delivery.  For the delivery confirmation, user can select any of the following special services, but they may vary as per the carrier selection. &lt;br /&gt;   - Signature Required/ Indirect Signature Required : SIG   - Signed Hard Copy: SIGHC   - Delivery confirmation: DEL_CON   - Proof of age required (18 years) Adult Signature Required: ADULT_SIG   - Proof of age required (19 years): ADULT_SIG_19   - No Signature Required: NO_SIG   - Direct Signature Required: DIRECT_SIG   - Chain of Signature: COS       Carrier specific options:   - UPS supports *SIG and ADULT_SIG*.    - FedEx supports *SIG, ADULT_SIG, NO_SIG, and DIRECT_SIG*.   - Purolator supports *ADULT_SIG, NO_SIG, and COS*.   - GoFor supports *SIG*.   - CPC supports *SIG, SIGHC, DEL_CON, ADULT_SIG, ADULT_SIG_19, and NO_SIG*.      | [optional] 
**Handling** | Pointer to **string** | Few shipments need a special handling, and the reason can be fragile items or highly secured shipments. There might be other case scenarios. In a simple term, this field defines shipment handling, which provides users a capability to select handling options. &lt;br /&gt; User can select any of the following handling options (special services), but they may vary as per the carrier selection.   - Hold For Pickup: HOLD   - Saturday Delivery: SAT_DELIVERY   - UPS Premium Care: PREM_CARE   - Direct Delivery Only: DIRECT   - Additional Handling: ADD_HDL       Carrier specific options:   - UPS supports all handling options mentioned above.    - FedEx supports *HOLD, SAT_DELIVERY, and ADD_HDL*.   - Purolator supports *HOLD, SAT_DELIVERY, and ADD_HDL*.     | [optional] 
**Insurance** | Pointer to **string** | Indicates the insurance coverage, which is selected by users while create shipment - rate shopping. User can select below-mentioned special service for insurance:    - Declared Value Surcharge: INS      Carrier specific options:   - UPS, FedEx, Purolator, and CPC support special service *INS*.     | [optional] 
**References** | Pointer to [**ReferenceV2**](ReferenceV2.md) |  | [optional] 
**Metadata** | Pointer to [**[]ShipmentDomesticByCarrierMetadataInner**](ShipmentDomesticByCarrierMetadataInner.md) | Additional metadata that needs to be stored for this shipment can be added here.&lt;br /&gt; For now, &#39;Cost Account Name&#39; is supported. | [optional] 
**LabelSize** | **string** | Defines the label size of the Shipment, that is, the Shipping Label is available in different Doc Size. &lt;br /&gt; &#x60;Max length &#x3D; 10&#x60; | 
**LabelType** | **string** | Defines the type of the Shipment.  &lt;br /&gt; &#x60;Max length &#x3D; 14&#x60; | 
**LabelFormat** | **string** | Defines the file/format in which the label is printed.&lt;br /&gt; For ZPL2, DOC_4X6 will be supported, while for PDF, both the sizes are supported. &#x60;Max length &#x3D; 14&#x60; | 
**PrinterAliasName** | Pointer to **string** | Refers to a printer connected (directly or via network) to a computer. &#x60;Max length &#x3D; 60&#x60; | [optional] 
**DateOfShipment** | Pointer to **string** | The date when shipment is created/shipped. The format of the Date is YYYY-MM-DD. | [optional] 
**DeliveryOption** | Pointer to [**ShipmentDomesticByCarrierDeliveryOption**](ShipmentDomesticByCarrierDeliveryOption.md) |  | [optional] 

## Methods

### NewShipmentDomesticByCarrier

`func NewShipmentDomesticByCarrier(fromAddress FromAddressV2, toAddress ToAddressV2, parcelType string, rateShopBy string, labelSize string, labelType string, labelFormat string, ) *ShipmentDomesticByCarrier`

NewShipmentDomesticByCarrier instantiates a new ShipmentDomesticByCarrier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentDomesticByCarrierWithDefaults

`func NewShipmentDomesticByCarrierWithDefaults() *ShipmentDomesticByCarrier`

NewShipmentDomesticByCarrierWithDefaults instantiates a new ShipmentDomesticByCarrier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromAddress

`func (o *ShipmentDomesticByCarrier) GetFromAddress() FromAddressV2`

GetFromAddress returns the FromAddress field if non-nil, zero value otherwise.

### GetFromAddressOk

`func (o *ShipmentDomesticByCarrier) GetFromAddressOk() (*FromAddressV2, bool)`

GetFromAddressOk returns a tuple with the FromAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromAddress

`func (o *ShipmentDomesticByCarrier) SetFromAddress(v FromAddressV2)`

SetFromAddress sets FromAddress field to given value.


### GetToAddress

`func (o *ShipmentDomesticByCarrier) GetToAddress() ToAddressV2`

GetToAddress returns the ToAddress field if non-nil, zero value otherwise.

### GetToAddressOk

`func (o *ShipmentDomesticByCarrier) GetToAddressOk() (*ToAddressV2, bool)`

GetToAddressOk returns a tuple with the ToAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToAddress

`func (o *ShipmentDomesticByCarrier) SetToAddress(v ToAddressV2)`

SetToAddress sets ToAddress field to given value.


### GetParcel

`func (o *ShipmentDomesticByCarrier) GetParcel() ParcelV2`

GetParcel returns the Parcel field if non-nil, zero value otherwise.

### GetParcelOk

`func (o *ShipmentDomesticByCarrier) GetParcelOk() (*ParcelV2, bool)`

GetParcelOk returns a tuple with the Parcel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcel

`func (o *ShipmentDomesticByCarrier) SetParcel(v ParcelV2)`

SetParcel sets Parcel field to given value.

### HasParcel

`func (o *ShipmentDomesticByCarrier) HasParcel() bool`

HasParcel returns a boolean if a field has been set.

### GetParcelType

`func (o *ShipmentDomesticByCarrier) GetParcelType() string`

GetParcelType returns the ParcelType field if non-nil, zero value otherwise.

### GetParcelTypeOk

`func (o *ShipmentDomesticByCarrier) GetParcelTypeOk() (*string, bool)`

GetParcelTypeOk returns a tuple with the ParcelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParcelType

`func (o *ShipmentDomesticByCarrier) SetParcelType(v string)`

SetParcelType sets ParcelType field to given value.


### GetRateShopBy

`func (o *ShipmentDomesticByCarrier) GetRateShopBy() string`

GetRateShopBy returns the RateShopBy field if non-nil, zero value otherwise.

### GetRateShopByOk

`func (o *ShipmentDomesticByCarrier) GetRateShopByOk() (*string, bool)`

GetRateShopByOk returns a tuple with the RateShopBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateShopBy

`func (o *ShipmentDomesticByCarrier) SetRateShopBy(v string)`

SetRateShopBy sets RateShopBy field to given value.


### GetByCarrier

`func (o *ShipmentDomesticByCarrier) GetByCarrier() ByCarrierV2`

GetByCarrier returns the ByCarrier field if non-nil, zero value otherwise.

### GetByCarrierOk

`func (o *ShipmentDomesticByCarrier) GetByCarrierOk() (*ByCarrierV2, bool)`

GetByCarrierOk returns a tuple with the ByCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByCarrier

`func (o *ShipmentDomesticByCarrier) SetByCarrier(v ByCarrierV2)`

SetByCarrier sets ByCarrier field to given value.

### HasByCarrier

`func (o *ShipmentDomesticByCarrier) HasByCarrier() bool`

HasByCarrier returns a boolean if a field has been set.

### GetShipmentOptions

`func (o *ShipmentDomesticByCarrier) GetShipmentOptions() ShipmentDomesticByCarrierShipmentOptions`

GetShipmentOptions returns the ShipmentOptions field if non-nil, zero value otherwise.

### GetShipmentOptionsOk

`func (o *ShipmentDomesticByCarrier) GetShipmentOptionsOk() (*ShipmentDomesticByCarrierShipmentOptions, bool)`

GetShipmentOptionsOk returns a tuple with the ShipmentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentOptions

`func (o *ShipmentDomesticByCarrier) SetShipmentOptions(v ShipmentDomesticByCarrierShipmentOptions)`

SetShipmentOptions sets ShipmentOptions field to given value.

### HasShipmentOptions

`func (o *ShipmentDomesticByCarrier) HasShipmentOptions() bool`

HasShipmentOptions returns a boolean if a field has been set.

### GetDocTab

`func (o *ShipmentDomesticByCarrier) GetDocTab() ShipmentDomesticByCarrierDocTab`

GetDocTab returns the DocTab field if non-nil, zero value otherwise.

### GetDocTabOk

`func (o *ShipmentDomesticByCarrier) GetDocTabOk() (*ShipmentDomesticByCarrierDocTab, bool)`

GetDocTabOk returns a tuple with the DocTab field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocTab

`func (o *ShipmentDomesticByCarrier) SetDocTab(v ShipmentDomesticByCarrierDocTab)`

SetDocTab sets DocTab field to given value.

### HasDocTab

`func (o *ShipmentDomesticByCarrier) HasDocTab() bool`

HasDocTab returns a boolean if a field has been set.

### GetDeliveryConfirmation

`func (o *ShipmentDomesticByCarrier) GetDeliveryConfirmation() string`

GetDeliveryConfirmation returns the DeliveryConfirmation field if non-nil, zero value otherwise.

### GetDeliveryConfirmationOk

`func (o *ShipmentDomesticByCarrier) GetDeliveryConfirmationOk() (*string, bool)`

GetDeliveryConfirmationOk returns a tuple with the DeliveryConfirmation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryConfirmation

`func (o *ShipmentDomesticByCarrier) SetDeliveryConfirmation(v string)`

SetDeliveryConfirmation sets DeliveryConfirmation field to given value.

### HasDeliveryConfirmation

`func (o *ShipmentDomesticByCarrier) HasDeliveryConfirmation() bool`

HasDeliveryConfirmation returns a boolean if a field has been set.

### GetHandling

`func (o *ShipmentDomesticByCarrier) GetHandling() string`

GetHandling returns the Handling field if non-nil, zero value otherwise.

### GetHandlingOk

`func (o *ShipmentDomesticByCarrier) GetHandlingOk() (*string, bool)`

GetHandlingOk returns a tuple with the Handling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandling

`func (o *ShipmentDomesticByCarrier) SetHandling(v string)`

SetHandling sets Handling field to given value.

### HasHandling

`func (o *ShipmentDomesticByCarrier) HasHandling() bool`

HasHandling returns a boolean if a field has been set.

### GetInsurance

`func (o *ShipmentDomesticByCarrier) GetInsurance() string`

GetInsurance returns the Insurance field if non-nil, zero value otherwise.

### GetInsuranceOk

`func (o *ShipmentDomesticByCarrier) GetInsuranceOk() (*string, bool)`

GetInsuranceOk returns a tuple with the Insurance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsurance

`func (o *ShipmentDomesticByCarrier) SetInsurance(v string)`

SetInsurance sets Insurance field to given value.

### HasInsurance

`func (o *ShipmentDomesticByCarrier) HasInsurance() bool`

HasInsurance returns a boolean if a field has been set.

### GetReferences

`func (o *ShipmentDomesticByCarrier) GetReferences() ReferenceV2`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShipmentDomesticByCarrier) GetReferencesOk() (*ReferenceV2, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShipmentDomesticByCarrier) SetReferences(v ReferenceV2)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShipmentDomesticByCarrier) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetMetadata

`func (o *ShipmentDomesticByCarrier) GetMetadata() []ShipmentDomesticByCarrierMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ShipmentDomesticByCarrier) GetMetadataOk() (*[]ShipmentDomesticByCarrierMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ShipmentDomesticByCarrier) SetMetadata(v []ShipmentDomesticByCarrierMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ShipmentDomesticByCarrier) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetLabelSize

`func (o *ShipmentDomesticByCarrier) GetLabelSize() string`

GetLabelSize returns the LabelSize field if non-nil, zero value otherwise.

### GetLabelSizeOk

`func (o *ShipmentDomesticByCarrier) GetLabelSizeOk() (*string, bool)`

GetLabelSizeOk returns a tuple with the LabelSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelSize

`func (o *ShipmentDomesticByCarrier) SetLabelSize(v string)`

SetLabelSize sets LabelSize field to given value.


### GetLabelType

`func (o *ShipmentDomesticByCarrier) GetLabelType() string`

GetLabelType returns the LabelType field if non-nil, zero value otherwise.

### GetLabelTypeOk

`func (o *ShipmentDomesticByCarrier) GetLabelTypeOk() (*string, bool)`

GetLabelTypeOk returns a tuple with the LabelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelType

`func (o *ShipmentDomesticByCarrier) SetLabelType(v string)`

SetLabelType sets LabelType field to given value.


### GetLabelFormat

`func (o *ShipmentDomesticByCarrier) GetLabelFormat() string`

GetLabelFormat returns the LabelFormat field if non-nil, zero value otherwise.

### GetLabelFormatOk

`func (o *ShipmentDomesticByCarrier) GetLabelFormatOk() (*string, bool)`

GetLabelFormatOk returns a tuple with the LabelFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelFormat

`func (o *ShipmentDomesticByCarrier) SetLabelFormat(v string)`

SetLabelFormat sets LabelFormat field to given value.


### GetPrinterAliasName

`func (o *ShipmentDomesticByCarrier) GetPrinterAliasName() string`

GetPrinterAliasName returns the PrinterAliasName field if non-nil, zero value otherwise.

### GetPrinterAliasNameOk

`func (o *ShipmentDomesticByCarrier) GetPrinterAliasNameOk() (*string, bool)`

GetPrinterAliasNameOk returns a tuple with the PrinterAliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterAliasName

`func (o *ShipmentDomesticByCarrier) SetPrinterAliasName(v string)`

SetPrinterAliasName sets PrinterAliasName field to given value.

### HasPrinterAliasName

`func (o *ShipmentDomesticByCarrier) HasPrinterAliasName() bool`

HasPrinterAliasName returns a boolean if a field has been set.

### GetDateOfShipment

`func (o *ShipmentDomesticByCarrier) GetDateOfShipment() string`

GetDateOfShipment returns the DateOfShipment field if non-nil, zero value otherwise.

### GetDateOfShipmentOk

`func (o *ShipmentDomesticByCarrier) GetDateOfShipmentOk() (*string, bool)`

GetDateOfShipmentOk returns a tuple with the DateOfShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfShipment

`func (o *ShipmentDomesticByCarrier) SetDateOfShipment(v string)`

SetDateOfShipment sets DateOfShipment field to given value.

### HasDateOfShipment

`func (o *ShipmentDomesticByCarrier) HasDateOfShipment() bool`

HasDateOfShipment returns a boolean if a field has been set.

### GetDeliveryOption

`func (o *ShipmentDomesticByCarrier) GetDeliveryOption() ShipmentDomesticByCarrierDeliveryOption`

GetDeliveryOption returns the DeliveryOption field if non-nil, zero value otherwise.

### GetDeliveryOptionOk

`func (o *ShipmentDomesticByCarrier) GetDeliveryOptionOk() (*ShipmentDomesticByCarrierDeliveryOption, bool)`

GetDeliveryOptionOk returns a tuple with the DeliveryOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryOption

`func (o *ShipmentDomesticByCarrier) SetDeliveryOption(v ShipmentDomesticByCarrierDeliveryOption)`

SetDeliveryOption sets DeliveryOption field to given value.

### HasDeliveryOption

`func (o *ShipmentDomesticByCarrier) HasDeliveryOption() bool`

HasDeliveryOption returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


