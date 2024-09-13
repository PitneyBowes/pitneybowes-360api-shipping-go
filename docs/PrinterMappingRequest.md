# PrinterMappingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | **string** | Name of the printer connected (directly or via network) to a computer. | 
**SerialNumber** | **string** | A Device Serial number. | 
**PrinterName** | **string** | The Printer name which is used for mapping. | 

## Methods

### NewPrinterMappingRequest

`func NewPrinterMappingRequest(alias string, serialNumber string, printerName string, ) *PrinterMappingRequest`

NewPrinterMappingRequest instantiates a new PrinterMappingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterMappingRequestWithDefaults

`func NewPrinterMappingRequestWithDefaults() *PrinterMappingRequest`

NewPrinterMappingRequestWithDefaults instantiates a new PrinterMappingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *PrinterMappingRequest) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PrinterMappingRequest) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PrinterMappingRequest) SetAlias(v string)`

SetAlias sets Alias field to given value.


### GetSerialNumber

`func (o *PrinterMappingRequest) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PrinterMappingRequest) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PrinterMappingRequest) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.


### GetPrinterName

`func (o *PrinterMappingRequest) GetPrinterName() string`

GetPrinterName returns the PrinterName field if non-nil, zero value otherwise.

### GetPrinterNameOk

`func (o *PrinterMappingRequest) GetPrinterNameOk() (*string, bool)`

GetPrinterNameOk returns a tuple with the PrinterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterName

`func (o *PrinterMappingRequest) SetPrinterName(v string)`

SetPrinterName sets PrinterName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


