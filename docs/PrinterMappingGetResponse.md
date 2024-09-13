# PrinterMappingGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | The printer connected (directly or via network) to a computer. | [optional] 
**SerialNumber** | Pointer to **string** | A Device Serial number. | [optional] 
**PrinterName** | Pointer to **string** | The Printer name which is used for mapping. | [optional] 
**SubId** | Pointer to **string** | The subscription ID used for mapping. | [optional] 
**InsertTimeStamp** | Pointer to **time.Time** | It records saved Timestamp. | [optional] 
**UpdateTimeStamp** | Pointer to **time.Time** | It records updated Timestamp. | [optional] 

## Methods

### NewPrinterMappingGetResponse

`func NewPrinterMappingGetResponse() *PrinterMappingGetResponse`

NewPrinterMappingGetResponse instantiates a new PrinterMappingGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrinterMappingGetResponseWithDefaults

`func NewPrinterMappingGetResponseWithDefaults() *PrinterMappingGetResponse`

NewPrinterMappingGetResponseWithDefaults instantiates a new PrinterMappingGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *PrinterMappingGetResponse) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *PrinterMappingGetResponse) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *PrinterMappingGetResponse) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *PrinterMappingGetResponse) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetSerialNumber

`func (o *PrinterMappingGetResponse) GetSerialNumber() string`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *PrinterMappingGetResponse) GetSerialNumberOk() (*string, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *PrinterMappingGetResponse) SetSerialNumber(v string)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *PrinterMappingGetResponse) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetPrinterName

`func (o *PrinterMappingGetResponse) GetPrinterName() string`

GetPrinterName returns the PrinterName field if non-nil, zero value otherwise.

### GetPrinterNameOk

`func (o *PrinterMappingGetResponse) GetPrinterNameOk() (*string, bool)`

GetPrinterNameOk returns a tuple with the PrinterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterName

`func (o *PrinterMappingGetResponse) SetPrinterName(v string)`

SetPrinterName sets PrinterName field to given value.

### HasPrinterName

`func (o *PrinterMappingGetResponse) HasPrinterName() bool`

HasPrinterName returns a boolean if a field has been set.

### GetSubId

`func (o *PrinterMappingGetResponse) GetSubId() string`

GetSubId returns the SubId field if non-nil, zero value otherwise.

### GetSubIdOk

`func (o *PrinterMappingGetResponse) GetSubIdOk() (*string, bool)`

GetSubIdOk returns a tuple with the SubId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubId

`func (o *PrinterMappingGetResponse) SetSubId(v string)`

SetSubId sets SubId field to given value.

### HasSubId

`func (o *PrinterMappingGetResponse) HasSubId() bool`

HasSubId returns a boolean if a field has been set.

### GetInsertTimeStamp

`func (o *PrinterMappingGetResponse) GetInsertTimeStamp() time.Time`

GetInsertTimeStamp returns the InsertTimeStamp field if non-nil, zero value otherwise.

### GetInsertTimeStampOk

`func (o *PrinterMappingGetResponse) GetInsertTimeStampOk() (*time.Time, bool)`

GetInsertTimeStampOk returns a tuple with the InsertTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsertTimeStamp

`func (o *PrinterMappingGetResponse) SetInsertTimeStamp(v time.Time)`

SetInsertTimeStamp sets InsertTimeStamp field to given value.

### HasInsertTimeStamp

`func (o *PrinterMappingGetResponse) HasInsertTimeStamp() bool`

HasInsertTimeStamp returns a boolean if a field has been set.

### GetUpdateTimeStamp

`func (o *PrinterMappingGetResponse) GetUpdateTimeStamp() time.Time`

GetUpdateTimeStamp returns the UpdateTimeStamp field if non-nil, zero value otherwise.

### GetUpdateTimeStampOk

`func (o *PrinterMappingGetResponse) GetUpdateTimeStampOk() (*time.Time, bool)`

GetUpdateTimeStampOk returns a tuple with the UpdateTimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTimeStamp

`func (o *PrinterMappingGetResponse) SetUpdateTimeStamp(v time.Time)`

SetUpdateTimeStamp sets UpdateTimeStamp field to given value.

### HasUpdateTimeStamp

`func (o *PrinterMappingGetResponse) HasUpdateTimeStamp() bool`

HasUpdateTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


