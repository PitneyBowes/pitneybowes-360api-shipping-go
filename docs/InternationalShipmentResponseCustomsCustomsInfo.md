# InternationalShipmentResponseCustomsCustomsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EELPFC** | Pointer to **string** | A number provided by the Automated Export System (AES). This is required if the item is valued at more than $2,500 USD per Schedule B export codes. | [optional] 
**CertificateNumber** | Pointer to **string** | The certificate number associated with the commodity. | [optional] 
**Comments** | Pointer to **string** | Free-form comments regarding the exported shipment. | [optional] 
**CurrencyCode** | Pointer to **string** | The currency used for declared value. Use three uppercase letters, per ISO 4217. | [optional] 
**FromCustomsReference** | Pointer to **string** | Free-form reference information provided by the requestor of the shipment. Depending on the carrier this information may or may not be rendered on the customs documents. | [optional] 
**ImporterCustomsReference** | Pointer to **string** | A reference number used by the importer, such as a VAT number, PO number, or insured number. | [optional] 
**InvoiceNumber** | Pointer to **string** | The commercial invoice number assigned by the exporter. It is rquired for DHL Express. | [optional] 
**LicenseNumber** | Pointer to **string** | The export license number associated with the commodity. | [optional] 
**ReasonForExport** | Pointer to **string** | The reason the commodity is being exported. | [optional] 
**SdrValue** | Pointer to **float32** | When an international parcel is insured, the insured value must be expressed in Special Drawing Rights values. | [optional] 
**TermsOfSale** | Pointer to **string** | Three-digit codes describing Terms of sale required for DHL Express shipments. Terms of sale is some sort of agreement between a buyer and seller which serves the purpose of creating uniform expectations between them. Possible values are- CFR (description- Cost and Freight), CIF (description- Cost, Insurance and Freight), CIP (description- Carriage and Insurance Paid To), DAF (description- Delivered at Frontier), DAP (description- Delivered At Place), DAT (description- Delivered At Terminal), DDP (description- Delivered Duty Paid), DDU (description- Delivered Duty Unpaid), DEQ (description- Delivery ex Quay), DES (description- Delivered ex Ship), DPU (description- Delivered At Place Unloaded), EXW (description- Ex Works), FAS (description- FAS Free Alongside Ship), FCA (description- Free Carrier), FOB (description- FOB Free On Board) | [optional] 
**RecipientTaxType** | Pointer to **string** | The tax type to choose for recipient. | [optional] 
**RecipientTaxID** | Pointer to **string** | The respective taxID for the taxType chosen | [optional] 
**SenderTaxType** | Pointer to **string** | The tax type to choose for recipient. | [optional] 
**SenderTaxID** | Pointer to **string** | The respective taxID for the taxType chosen | [optional] 

## Methods

### NewInternationalShipmentResponseCustomsCustomsInfo

`func NewInternationalShipmentResponseCustomsCustomsInfo() *InternationalShipmentResponseCustomsCustomsInfo`

NewInternationalShipmentResponseCustomsCustomsInfo instantiates a new InternationalShipmentResponseCustomsCustomsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternationalShipmentResponseCustomsCustomsInfoWithDefaults

`func NewInternationalShipmentResponseCustomsCustomsInfoWithDefaults() *InternationalShipmentResponseCustomsCustomsInfo`

NewInternationalShipmentResponseCustomsCustomsInfoWithDefaults instantiates a new InternationalShipmentResponseCustomsCustomsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEELPFC

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetEELPFC() string`

GetEELPFC returns the EELPFC field if non-nil, zero value otherwise.

### GetEELPFCOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetEELPFCOk() (*string, bool)`

GetEELPFCOk returns a tuple with the EELPFC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEELPFC

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetEELPFC(v string)`

SetEELPFC sets EELPFC field to given value.

### HasEELPFC

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasEELPFC() bool`

HasEELPFC returns a boolean if a field has been set.

### GetCertificateNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetCertificateNumber() string`

GetCertificateNumber returns the CertificateNumber field if non-nil, zero value otherwise.

### GetCertificateNumberOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetCertificateNumberOk() (*string, bool)`

GetCertificateNumberOk returns a tuple with the CertificateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificateNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetCertificateNumber(v string)`

SetCertificateNumber sets CertificateNumber field to given value.

### HasCertificateNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasCertificateNumber() bool`

HasCertificateNumber returns a boolean if a field has been set.

### GetComments

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetComments() string`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetCommentsOk() (*string, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetComments(v string)`

SetComments sets Comments field to given value.

### HasComments

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetCurrencyCode

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetCurrencyCode() string`

GetCurrencyCode returns the CurrencyCode field if non-nil, zero value otherwise.

### GetCurrencyCodeOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetCurrencyCodeOk() (*string, bool)`

GetCurrencyCodeOk returns a tuple with the CurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyCode

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetCurrencyCode(v string)`

SetCurrencyCode sets CurrencyCode field to given value.

### HasCurrencyCode

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasCurrencyCode() bool`

HasCurrencyCode returns a boolean if a field has been set.

### GetFromCustomsReference

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetFromCustomsReference() string`

GetFromCustomsReference returns the FromCustomsReference field if non-nil, zero value otherwise.

### GetFromCustomsReferenceOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetFromCustomsReferenceOk() (*string, bool)`

GetFromCustomsReferenceOk returns a tuple with the FromCustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCustomsReference

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetFromCustomsReference(v string)`

SetFromCustomsReference sets FromCustomsReference field to given value.

### HasFromCustomsReference

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasFromCustomsReference() bool`

HasFromCustomsReference returns a boolean if a field has been set.

### GetImporterCustomsReference

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetImporterCustomsReference() string`

GetImporterCustomsReference returns the ImporterCustomsReference field if non-nil, zero value otherwise.

### GetImporterCustomsReferenceOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetImporterCustomsReferenceOk() (*string, bool)`

GetImporterCustomsReferenceOk returns a tuple with the ImporterCustomsReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImporterCustomsReference

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetImporterCustomsReference(v string)`

SetImporterCustomsReference sets ImporterCustomsReference field to given value.

### HasImporterCustomsReference

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasImporterCustomsReference() bool`

HasImporterCustomsReference returns a boolean if a field has been set.

### GetInvoiceNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetInvoiceNumber() string`

GetInvoiceNumber returns the InvoiceNumber field if non-nil, zero value otherwise.

### GetInvoiceNumberOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetInvoiceNumberOk() (*string, bool)`

GetInvoiceNumberOk returns a tuple with the InvoiceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetInvoiceNumber(v string)`

SetInvoiceNumber sets InvoiceNumber field to given value.

### HasInvoiceNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasInvoiceNumber() bool`

HasInvoiceNumber returns a boolean if a field has been set.

### GetLicenseNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetLicenseNumber() string`

GetLicenseNumber returns the LicenseNumber field if non-nil, zero value otherwise.

### GetLicenseNumberOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetLicenseNumberOk() (*string, bool)`

GetLicenseNumberOk returns a tuple with the LicenseNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetLicenseNumber(v string)`

SetLicenseNumber sets LicenseNumber field to given value.

### HasLicenseNumber

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasLicenseNumber() bool`

HasLicenseNumber returns a boolean if a field has been set.

### GetReasonForExport

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetReasonForExport() string`

GetReasonForExport returns the ReasonForExport field if non-nil, zero value otherwise.

### GetReasonForExportOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetReasonForExportOk() (*string, bool)`

GetReasonForExportOk returns a tuple with the ReasonForExport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonForExport

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetReasonForExport(v string)`

SetReasonForExport sets ReasonForExport field to given value.

### HasReasonForExport

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasReasonForExport() bool`

HasReasonForExport returns a boolean if a field has been set.

### GetSdrValue

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetSdrValue() float32`

GetSdrValue returns the SdrValue field if non-nil, zero value otherwise.

### GetSdrValueOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetSdrValueOk() (*float32, bool)`

GetSdrValueOk returns a tuple with the SdrValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdrValue

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetSdrValue(v float32)`

SetSdrValue sets SdrValue field to given value.

### HasSdrValue

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasSdrValue() bool`

HasSdrValue returns a boolean if a field has been set.

### GetTermsOfSale

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetTermsOfSale() string`

GetTermsOfSale returns the TermsOfSale field if non-nil, zero value otherwise.

### GetTermsOfSaleOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetTermsOfSaleOk() (*string, bool)`

GetTermsOfSaleOk returns a tuple with the TermsOfSale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfSale

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetTermsOfSale(v string)`

SetTermsOfSale sets TermsOfSale field to given value.

### HasTermsOfSale

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasTermsOfSale() bool`

HasTermsOfSale returns a boolean if a field has been set.

### GetRecipientTaxType

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetRecipientTaxType() string`

GetRecipientTaxType returns the RecipientTaxType field if non-nil, zero value otherwise.

### GetRecipientTaxTypeOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetRecipientTaxTypeOk() (*string, bool)`

GetRecipientTaxTypeOk returns a tuple with the RecipientTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientTaxType

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetRecipientTaxType(v string)`

SetRecipientTaxType sets RecipientTaxType field to given value.

### HasRecipientTaxType

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasRecipientTaxType() bool`

HasRecipientTaxType returns a boolean if a field has been set.

### GetRecipientTaxID

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetRecipientTaxID() string`

GetRecipientTaxID returns the RecipientTaxID field if non-nil, zero value otherwise.

### GetRecipientTaxIDOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetRecipientTaxIDOk() (*string, bool)`

GetRecipientTaxIDOk returns a tuple with the RecipientTaxID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientTaxID

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetRecipientTaxID(v string)`

SetRecipientTaxID sets RecipientTaxID field to given value.

### HasRecipientTaxID

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasRecipientTaxID() bool`

HasRecipientTaxID returns a boolean if a field has been set.

### GetSenderTaxType

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetSenderTaxType() string`

GetSenderTaxType returns the SenderTaxType field if non-nil, zero value otherwise.

### GetSenderTaxTypeOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetSenderTaxTypeOk() (*string, bool)`

GetSenderTaxTypeOk returns a tuple with the SenderTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderTaxType

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetSenderTaxType(v string)`

SetSenderTaxType sets SenderTaxType field to given value.

### HasSenderTaxType

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasSenderTaxType() bool`

HasSenderTaxType returns a boolean if a field has been set.

### GetSenderTaxID

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetSenderTaxID() string`

GetSenderTaxID returns the SenderTaxID field if non-nil, zero value otherwise.

### GetSenderTaxIDOk

`func (o *InternationalShipmentResponseCustomsCustomsInfo) GetSenderTaxIDOk() (*string, bool)`

GetSenderTaxIDOk returns a tuple with the SenderTaxID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderTaxID

`func (o *InternationalShipmentResponseCustomsCustomsInfo) SetSenderTaxID(v string)`

SetSenderTaxID sets SenderTaxID field to given value.

### HasSenderTaxID

`func (o *InternationalShipmentResponseCustomsCustomsInfo) HasSenderTaxID() bool`

HasSenderTaxID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


