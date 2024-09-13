# \PrintAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePrinterMapping**](PrintAPI.md#DeletePrinterMapping) | **Delete** /api/v1/printer/mapping | Delete Printer mapping
[**GetPrinterMapping**](PrintAPI.md#GetPrinterMapping) | **Get** /api/v1/printer/mapping | Get Printer mapping
[**JobStatus**](PrintAPI.md#JobStatus) | **Get** /api/v1/jobs/{jobId} | Job status
[**PrintDocument**](PrintAPI.md#PrintDocument) | **Post** /api/v1/document/print | Print Document
[**PrinterMapping**](PrintAPI.md#PrinterMapping) | **Post** /api/v1/printer/mapping | Printer mapping



## DeletePrinterMapping

> DeletePrinterMapping(ctx).Alias(alias).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Delete Printer mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PitneyBowes/pitneybowes-360api-shipping-go"
)

func main() {
	alias := "alias_example" // string | Refers to a printer connected (directly or via network) to a computer.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PrintAPI.DeletePrinterMapping(context.Background()).Alias(alias).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintAPI.DeletePrinterMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrinterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alias** | **string** | Refers to a printer connected (directly or via network) to a computer. | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrinterMapping

> PrinterMappingGetResponse GetPrinterMapping(ctx).Alias(alias).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Get Printer mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PitneyBowes/pitneybowes-360api-shipping-go"
)

func main() {
	alias := "alias_example" // string | Refers to a printer connected (directly or via network) to a computer.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintAPI.GetPrinterMapping(context.Background()).Alias(alias).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintAPI.GetPrinterMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrinterMapping`: PrinterMappingGetResponse
	fmt.Fprintf(os.Stdout, "Response from `PrintAPI.GetPrinterMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrinterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **alias** | **string** | Refers to a printer connected (directly or via network) to a computer. | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**PrinterMappingGetResponse**](PrinterMappingGetResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## JobStatus

> JobStatus JobStatus(ctx, jobId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Job status



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PitneyBowes/pitneybowes-360api-shipping-go"
)

func main() {
	jobId := "jobId_example" // string | The jobId, a unique identifier assigned for the job.
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintAPI.JobStatus(context.Background(), jobId).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintAPI.JobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `JobStatus`: JobStatus
	fmt.Fprintf(os.Stdout, "Response from `PrintAPI.JobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | The jobId, a unique identifier assigned for the job. | 

### Other Parameters

Other parameters are passed through a pointer to a apiJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**JobStatus**](JobStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrintDocument

> PrintDocumentResponse PrintDocument(ctx).PrintDocumentRequest(printDocumentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Print Document



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PitneyBowes/pitneybowes-360api-shipping-go"
)

func main() {
	printDocumentRequest := *openapiclient.NewPrintDocumentRequest("Pitney Bowes Printer", "<<base64string>>", "BASE64", "ZPL2", "ENV10") // PrintDocumentRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintAPI.PrintDocument(context.Background()).PrintDocumentRequest(printDocumentRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintAPI.PrintDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrintDocument`: PrintDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `PrintAPI.PrintDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrintDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **printDocumentRequest** | [**PrintDocumentRequest**](PrintDocumentRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**PrintDocumentResponse**](PrintDocumentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PrinterMapping

> PrinterMappingResponse PrinterMapping(ctx).PrinterMappingRequest(printerMappingRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Printer mapping



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/PitneyBowes/pitneybowes-360api-shipping-go"
)

func main() {
	printerMappingRequest := *openapiclient.NewPrinterMappingRequest("rohit2", "AP-DP1LZH3-0x919e698164-dev", "Brother QL-1110NWB") // PrinterMappingRequest | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrintAPI.PrinterMapping(context.Background()).PrinterMappingRequest(printerMappingRequest).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrintAPI.PrinterMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PrinterMapping`: PrinterMappingResponse
	fmt.Fprintf(os.Stdout, "Response from `PrintAPI.PrinterMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrinterMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **printerMappingRequest** | [**PrinterMappingRequest**](PrinterMappingRequest.md) |  | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**PrinterMappingResponse**](PrinterMappingResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

