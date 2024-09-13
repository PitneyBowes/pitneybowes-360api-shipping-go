# \BatchAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkImportAPI**](BatchAPI.md#BulkImportAPI) | **Post** /api/v1/shipments/importUrl | Bulk Import Shipments
[**BulkImportAPIERR**](BatchAPI.md#BulkImportAPIERR) | **Post** /api/v1/err/shipments/importUrl | Bulk Import Shipments ERR
[**CreateBulkShipmentsAPI**](BatchAPI.md#CreateBulkShipmentsAPI) | **Post** /api/v1/bulkShipments | Create Bulk Shipments
[**CreateBulkShipmentsAPIERR**](BatchAPI.md#CreateBulkShipmentsAPIERR) | **Post** /api/v1/err/bulkShipments | Create Bulk Shipments ERR
[**GetBatchStatusAPI**](BatchAPI.md#GetBatchStatusAPI) | **Get** /api/v1/shipments/batch/{batchId}/status | Get Batch Status
[**GetShipmentDetailsForBatchAPI**](BatchAPI.md#GetShipmentDetailsForBatchAPI) | **Get** /api/v1/shipments/batch/{batchId}/shipments | Get Batch Shipment Details
[**ProcessBatchAPI**](BatchAPI.md#ProcessBatchAPI) | **Post** /api/v1/shipments/batch/{batchId}/process | Process Batch
[**VoidShippingLabel**](BatchAPI.md#VoidShippingLabel) | **Post** /api/v1/shipments/batch/{batchId}/void | Void Batch Shipping Labels



## BulkImportAPI

> ShipmentBatch BulkImportAPI(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Bulk Import Shipments



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
	body := *openapiclient.NewCreateBatchRequest("DOC_8X11", "SHIPPING_LABEL", "CarrierAccountId_example", "ServiceId_example", "ParcelType_example") // CreateBatchRequest |  This is the Request body to bulk import shipments.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BulkImportAPI(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BulkImportAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkImportAPI`: ShipmentBatch
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BulkImportAPI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkImportAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBatchRequest**](CreateBatchRequest.md) |  This is the Request body to bulk import shipments. | 
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**ShipmentBatch**](ShipmentBatch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkImportAPIERR

> ShipmentBatchResponseERR BulkImportAPIERR(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Bulk Import Shipments ERR



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
	body := *openapiclient.NewCreateBatchRequestERR("DOC_8X11", "SHIPPING_LABEL", "CarrierAccountId_example", "ServiceId_example", "ParcelType_example") // CreateBatchRequestERR |  This is the request body to import ERR Bulk shipments.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The PB-Developer-Partner-ID is assigned by PB to uniquely identify a developer's strategic business partners. If the developer is the sole business partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.BulkImportAPIERR(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.BulkImportAPIERR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkImportAPIERR`: ShipmentBatchResponseERR
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.BulkImportAPIERR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkImportAPIERRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBatchRequestERR**](CreateBatchRequestERR.md) |  This is the request body to import ERR Bulk shipments. | 
 **xPBDeveloperPartnerID** | **string** | The PB-Developer-Partner-ID is assigned by PB to uniquely identify a developer&#39;s strategic business partners. If the developer is the sole business partner, this field is not required. | 

### Return type

[**ShipmentBatchResponseERR**](ShipmentBatchResponseERR.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBulkShipmentsAPI

> BulkShipmentResponse CreateBulkShipmentsAPI(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Create Bulk Shipments



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
	body := openapiclient.CreateBulkShipmentsAPI_request{CreateBulkShipmentInternational: openapiclient.NewCreateBulkShipmentInternational("DOC_8X11", "SHIPPING_LABEL", "name", "abcd123", "PKG", "PMI", []openapiclient.ShipmentInternational{*openapiclient.NewShipmentInternational(*openapiclient.NewShipmentInternationalFromAddress("27 Watervw Dr", "US", "06905", "CT"), *openapiclient.NewShipmentInternationalParcel(float32(1), float32(2), float32(1), "IN", "OZ", float32(1)), "asas2223", "PKG", "PMI", *openapiclient.NewShipmentInternationalToAddress("70 Hanlan RD", "CA", "L4L3P6", "ON"), *openapiclient.NewShipmentInternationalCustoms(*openapiclient.NewShipmentInternationalCustomsCustomsInfo("ReasonForExport_example", float32(123), "USD")))})} // CreateBulkShipmentsAPIRequest | This is the Request body to create bulk shipments.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.CreateBulkShipmentsAPI(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.CreateBulkShipmentsAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkShipmentsAPI`: BulkShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.CreateBulkShipmentsAPI`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkShipmentsAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBulkShipmentsAPIRequest**](CreateBulkShipmentsAPIRequest.md) | This is the Request body to create bulk shipments. | 
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**BulkShipmentResponse**](BulkShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBulkShipmentsAPIERR

> BulkShipmentResponseERR CreateBulkShipmentsAPIERR(ctx).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Create Bulk Shipments ERR



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
	body := openapiclient.CreateBulkShipmentsAPIERR_request{CreateBulkShipmentsERRCoversheet: openapiclient.NewCreateBulkShipmentsERRCoversheet("ERR-Coversheet07", "10", "COVERSHEET", "qnpB1VkRMmw", "LTR", "FCM", []openapiclient.ShipmentERRCoversheet{*openapiclient.NewShipmentERRCoversheet(*openapiclient.NewFromAddress("Anupama", "27 Watervw Dr", "Stamford", "CT", "06905", "US"), *openapiclient.NewToAddress("John Smith", "98 Watervw Dr", "Shelton", "CT", "06905", "US"), *openapiclient.NewParcel("OZ"))})} // CreateBulkShipmentsAPIERRRequest | This is the request body to create Bulk Shipment for ERR.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The Developer-Partner-ID is assigned by PB to uniquely identify a developer's strategic business partners. If the developer is the sole business partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.CreateBulkShipmentsAPIERR(context.Background()).Body(body).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.CreateBulkShipmentsAPIERR``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBulkShipmentsAPIERR`: BulkShipmentResponseERR
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.CreateBulkShipmentsAPIERR`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkShipmentsAPIERRRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateBulkShipmentsAPIERRRequest**](CreateBulkShipmentsAPIERRRequest.md) | This is the request body to create Bulk Shipment for ERR. | 
 **xPBDeveloperPartnerID** | **string** | The Developer-Partner-ID is assigned by PB to uniquely identify a developer&#39;s strategic business partners. If the developer is the sole business partner, this field is not required. | 

### Return type

[**BulkShipmentResponseERR**](BulkShipmentResponseERR.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBatchStatusAPI

> GetStatusDetailedResponse GetBatchStatusAPI(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Get Batch Status



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
	batchId := "batchId_example" // string | This is a system-generated unique identifier assigned to the Batch while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The Developer-Partner- ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetBatchStatusAPI(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetBatchStatusAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBatchStatusAPI`: GetStatusDetailedResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetBatchStatusAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | This is a system-generated unique identifier assigned to the Batch while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBatchStatusAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | The Developer-Partner- ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field is not required. | 

### Return type

[**GetStatusDetailedResponse**](GetStatusDetailedResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetShipmentDetailsForBatchAPI

> GetShipmentsForBatch GetShipmentDetailsForBatchAPI(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Page(page).Size(size).Status(status).Step(step).Execute()

Get Batch Shipment Details



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
	batchId := "batchId_example" // string | This is a system-generated unique identifier assigned to the Batch while it is processed
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The Developer-Partner- ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field is not required. (optional)
	page := int32(56) // int32 | It returns detailed information for shipments status and it can cover in one or more pages. The default value for page number is 1. (optional)
	size := int32(56) // int32 | Indicates the number of records per page. The default value for records is 20. (optional)
	status := "status_example" // string | The status of the shipment. Values can be Failed or Success. (optional)
	step := "step_example" // string | Indicates various stages of the batch processing. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.GetShipmentDetailsForBatchAPI(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Page(page).Size(size).Status(status).Step(step).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.GetShipmentDetailsForBatchAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetShipmentDetailsForBatchAPI`: GetShipmentsForBatch
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.GetShipmentDetailsForBatchAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | This is a system-generated unique identifier assigned to the Batch while it is processed | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShipmentDetailsForBatchAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | The Developer-Partner- ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field is not required. | 
 **page** | **int32** | It returns detailed information for shipments status and it can cover in one or more pages. The default value for page number is 1. | 
 **size** | **int32** | Indicates the number of records per page. The default value for records is 20. | 
 **status** | **string** | The status of the shipment. Values can be Failed or Success. | 
 **step** | **string** | Indicates various stages of the batch processing. | 

### Return type

[**GetShipmentsForBatch**](GetShipmentsForBatch.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessBatchAPI

> ProcessShipmentResponse ProcessBatchAPI(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Process Batch



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
	batchId := "batchId_example" // string | This is a system-generated unique identifier assigned to the Batch while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The Developer-Partner-ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.ProcessBatchAPI(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.ProcessBatchAPI``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ProcessBatchAPI`: ProcessShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.ProcessBatchAPI`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | This is a system-generated unique identifier assigned to the Batch while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiProcessBatchAPIRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | The Developer-Partner-ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field is not required. | 

### Return type

[**ProcessShipmentResponse**](ProcessShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VoidShippingLabel

> VoidBatchResponse VoidShippingLabel(ctx, batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).VoidBatchRequest(voidBatchRequest).Execute()

Void Batch Shipping Labels



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
	batchId := "batchId_example" // string | This is a system-generated unique identifier assigned to the Batch while it is processed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The Developer-Partner-ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field is not required. (optional)
	voidBatchRequest := *openapiclient.NewVoidBatchRequest("batchTest") // VoidBatchRequest |  This is the request body for cancelling the selected shipments or entire Batch of shipments*. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BatchAPI.VoidShippingLabel(context.Background(), batchId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).VoidBatchRequest(voidBatchRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BatchAPI.VoidShippingLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VoidShippingLabel`: VoidBatchResponse
	fmt.Fprintf(os.Stdout, "Response from `BatchAPI.VoidShippingLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**batchId** | **string** | This is a system-generated unique identifier assigned to the Batch while it is processed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiVoidShippingLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | The Developer-Partner-ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field is not required. | 
 **voidBatchRequest** | [**VoidBatchRequest**](VoidBatchRequest.md) |  This is the request body for cancelling the selected shipments or entire Batch of shipments*. | 

### Return type

[**VoidBatchResponse**](VoidBatchResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

