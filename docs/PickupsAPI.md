# \PickupsAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelPickups**](PickupsAPI.md#CancelPickups) | **Put** /api/v1/pickups/cancel | Cancel Pickups
[**CancelledPickupDocument**](PickupsAPI.md#CancelledPickupDocument) | **Post** /api/v1/pickups/document | Cancelled Pickup Document
[**GetPickupDocument**](PickupsAPI.md#GetPickupDocument) | **Get** /api/v1/pickups/{pickupId}/document | Get Pickup Document
[**GetPickups**](PickupsAPI.md#GetPickups) | **Get** /api/v1/pickups | Get Pickups
[**SchedulePickup**](PickupsAPI.md#SchedulePickup) | **Post** /api/v1/pickups | Schedule Pickup



## CancelPickups

> SchedulePickupCancelResponse CancelPickups(ctx).SchedulePickupCancelRequest(schedulePickupCancelRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Cancel Pickups



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
	schedulePickupCancelRequest := *openapiclient.NewSchedulePickupCancelRequest([]string{"UPS10191697701340901"}) // SchedulePickupCancelRequest | 
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupsAPI.CancelPickups(context.Background()).SchedulePickupCancelRequest(schedulePickupCancelRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupsAPI.CancelPickups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelPickups`: SchedulePickupCancelResponse
	fmt.Fprintf(os.Stdout, "Response from `PickupsAPI.CancelPickups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelPickupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulePickupCancelRequest** | [**SchedulePickupCancelRequest**](SchedulePickupCancelRequest.md) |  | 
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**SchedulePickupCancelResponse**](SchedulePickupCancelResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelledPickupDocument

> GetPickupCancelledDocumentResponse CancelledPickupDocument(ctx).Type_(type_).GetPickupCancelledDocumentRequest(getPickupCancelledDocumentRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Cancelled Pickup Document



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
	type_ := "cancelled" // string | Indicates type of pickup. Supported value is `cancelled`.
	getPickupCancelledDocumentRequest := *openapiclient.NewGetPickupCancelledDocumentRequest() // GetPickupCancelledDocumentRequest | 
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupsAPI.CancelledPickupDocument(context.Background()).Type_(type_).GetPickupCancelledDocumentRequest(getPickupCancelledDocumentRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupsAPI.CancelledPickupDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelledPickupDocument`: GetPickupCancelledDocumentResponse
	fmt.Fprintf(os.Stdout, "Response from `PickupsAPI.CancelledPickupDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelledPickupDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Indicates type of pickup. Supported value is &#x60;cancelled&#x60;. | 
 **getPickupCancelledDocumentRequest** | [**GetPickupCancelledDocumentRequest**](GetPickupCancelledDocumentRequest.md) |  | 
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetPickupCancelledDocumentResponse**](GetPickupCancelledDocumentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPickupDocument

> GetPickupDocument GetPickupDocument(ctx, pickupId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Get Pickup Document



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
	pickupId := "pickupId_example" // string | It specified the pickupId for which PDF receipt is needed.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupsAPI.GetPickupDocument(context.Background(), pickupId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupsAPI.GetPickupDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPickupDocument`: GetPickupDocument
	fmt.Fprintf(os.Stdout, "Response from `PickupsAPI.GetPickupDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pickupId** | **string** | It specified the pickupId for which PDF receipt is needed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPickupDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**GetPickupDocument**](GetPickupDocument.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPickups

> GetAllPickups GetPickups(ctx).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Carrier(carrier).StartDate(startDate).EndDate(endDate).Status(status).Page(page).Size(size).Execute()

Get Pickups



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
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)
	carrier := "carrier_example" // string | Indicates CarrierID. If not provided, it would show pickups for all the carriers. (optional)
	startDate := "startDate_example" // string | Indicates start date from when you want to see the history. If not provided, it will take today's date. (optional)
	endDate := "endDate_example" // string | Indicates end date till you want to see the pickups history. If not provide, it will take today's date. (optional)
	status := "status_example" // string | Indicates status of the pickup(schedule or cancel). If not provided, it will show the response for both status. (optional)
	page := float32(8.14) // float32 | Indicates page number, if not provided page would be 1. (optional)
	size := float32(8.14) // float32 | Indicates size of records, if not provided size would be 20 (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupsAPI.GetPickups(context.Background()).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Carrier(carrier).StartDate(startDate).EndDate(endDate).Status(status).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupsAPI.GetPickups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPickups`: GetAllPickups
	fmt.Fprintf(os.Stdout, "Response from `PickupsAPI.GetPickups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPickupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **carrier** | **string** | Indicates CarrierID. If not provided, it would show pickups for all the carriers. | 
 **startDate** | **string** | Indicates start date from when you want to see the history. If not provided, it will take today&#39;s date. | 
 **endDate** | **string** | Indicates end date till you want to see the pickups history. If not provide, it will take today&#39;s date. | 
 **status** | **string** | Indicates status of the pickup(schedule or cancel). If not provided, it will show the response for both status. | 
 **page** | **float32** | Indicates page number, if not provided page would be 1. | 
 **size** | **float32** | Indicates size of records, if not provided size would be 20 | 

### Return type

[**GetAllPickups**](GetAllPickups.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SchedulePickup

> SchedulePickup200Response SchedulePickup(ctx).SchedulePickupRequest(schedulePickupRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Schedule Pickup



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
	schedulePickupRequest := openapiclient.schedulePickup_request{SchedulePickupDHLEXPRequest: openapiclient.NewSchedulePickupDHLEXPRequest("NONE", "j4pqLKjQ5dn", *openapiclient.NewSchedulePickupDHLEXPRequestPickupAddress("Test", "27 Waterview Dr", "Shelton", "CT", "06484-4301", "US", "1234567890"), *openapiclient.NewSchedulePickupDHLEXPRequestPickupOptions("2023-11-06T17:05:05Z", "2023-11-06T18:30:00Z"))} // SchedulePickupRequest | 
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Developer Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PickupsAPI.SchedulePickup(context.Background()).SchedulePickupRequest(schedulePickupRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PickupsAPI.SchedulePickup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SchedulePickup`: SchedulePickup200Response
	fmt.Fprintf(os.Stdout, "Response from `PickupsAPI.SchedulePickup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSchedulePickupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schedulePickupRequest** | [**SchedulePickupRequest**](SchedulePickupRequest.md) |  | 
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**SchedulePickup200Response**](SchedulePickup200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

