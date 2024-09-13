# JobStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** | Job ID  can be used to check the status of the print request. | [optional] 
**Status** | Pointer to **string** | Refers to the status of Print Job. | [optional] 
**PrintStatusTransaction** | Pointer to [**[]JobStatusPrintStatusTransactionInner**](JobStatusPrintStatusTransactionInner.md) |  | [optional] 

## Methods

### NewJobStatus

`func NewJobStatus() *JobStatus`

NewJobStatus instantiates a new JobStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobStatusWithDefaults

`func NewJobStatusWithDefaults() *JobStatus`

NewJobStatusWithDefaults instantiates a new JobStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *JobStatus) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *JobStatus) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *JobStatus) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *JobStatus) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetStatus

`func (o *JobStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *JobStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPrintStatusTransaction

`func (o *JobStatus) GetPrintStatusTransaction() []JobStatusPrintStatusTransactionInner`

GetPrintStatusTransaction returns the PrintStatusTransaction field if non-nil, zero value otherwise.

### GetPrintStatusTransactionOk

`func (o *JobStatus) GetPrintStatusTransactionOk() (*[]JobStatusPrintStatusTransactionInner, bool)`

GetPrintStatusTransactionOk returns a tuple with the PrintStatusTransaction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintStatusTransaction

`func (o *JobStatus) SetPrintStatusTransaction(v []JobStatusPrintStatusTransactionInner)`

SetPrintStatusTransaction sets PrintStatusTransaction field to given value.

### HasPrintStatusTransaction

`func (o *JobStatus) HasPrintStatusTransaction() bool`

HasPrintStatusTransaction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


