# \MultipieceAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MultipieceRates**](MultipieceAPI.md#MultipieceRates) | **Post** /api/v1/multipiece/rates | Multipiece Rateshop and Rates
[**MultipieceShipment**](MultipieceAPI.md#MultipieceShipment) | **Post** /api/v1/multipiece/shipments | Multipiece Shipment
[**MultipieceShipmentCancel**](MultipieceAPI.md#MultipieceShipmentCancel) | **Put** /api/v1/multipiece/shipments/{shipmentId}/cancel | Cancel Multipiece Shipment
[**MultipieceShipmentReprint**](MultipieceAPI.md#MultipieceShipmentReprint) | **Get** /api/v1/multipiece/shipments/{shipmentId}/reprint | Reprint Multipiece Shipment



## MultipieceRates

> MultipieceRates200Response MultipieceRates(ctx).MultipieceRatesRequest(multipieceRatesRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Multipiece Rateshop and Rates



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
	multipieceRatesRequest := openapiclient.multipieceRates_request{MultipieceRateShopRequest: openapiclient.NewMultipieceRateShopRequest()} // MultipieceRatesRequest | 
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceAPI.MultipieceRates(context.Background()).MultipieceRatesRequest(multipieceRatesRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceAPI.MultipieceRates``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceRates`: MultipieceRates200Response
	fmt.Fprintf(os.Stdout, "Response from `MultipieceAPI.MultipieceRates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceRatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipieceRatesRequest** | [**MultipieceRatesRequest**](MultipieceRatesRequest.md) |  | 
 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**MultipieceRates200Response**](MultipieceRates200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultipieceShipment

> MultipieceShipment200Response MultipieceShipment(ctx).MultipieceShipmentRequest(multipieceShipmentRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Multipiece Shipment



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
	multipieceShipmentRequest := openapiclient.multipieceShipment_request{MultipieceDomesticShipmentRequest: openapiclient.NewMultipieceDomesticShipmentRequest()} // MultipieceShipmentRequest | 
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceAPI.MultipieceShipment(context.Background()).MultipieceShipmentRequest(multipieceShipmentRequest).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceAPI.MultipieceShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceShipment`: MultipieceShipment200Response
	fmt.Fprintf(os.Stdout, "Response from `MultipieceAPI.MultipieceShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipieceShipmentRequest** | [**MultipieceShipmentRequest**](MultipieceShipmentRequest.md) |  | 
 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**MultipieceShipment200Response**](MultipieceShipment200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultipieceShipmentCancel

> CancelShipment MultipieceShipmentCancel(ctx, shipmentId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Cancel Multipiece Shipment



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
	shipmentId := "shipmentId_example" // string | It specifies the shipmentId of onward shipment against which return label has to be created.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceAPI.MultipieceShipmentCancel(context.Background(), shipmentId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceAPI.MultipieceShipmentCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceShipmentCancel`: CancelShipment
	fmt.Fprintf(os.Stdout, "Response from `MultipieceAPI.MultipieceShipmentCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | It specifies the shipmentId of onward shipment against which return label has to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceShipmentCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**CancelShipment**](CancelShipment.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MultipieceShipmentReprint

> MultipieceDomesticShipmentResponse MultipieceShipmentReprint(ctx, shipmentId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()

Reprint Multipiece Shipment



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
	shipmentId := "shipmentId_example" // string | It specifies the shipmentId of onward shipment against which return label has to be created.
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | This is the Develover Partner ID. When the developer is the only partner, this field is not required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MultipieceAPI.MultipieceShipmentReprint(context.Background(), shipmentId).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MultipieceAPI.MultipieceShipmentReprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MultipieceShipmentReprint`: MultipieceDomesticShipmentResponse
	fmt.Fprintf(os.Stdout, "Response from `MultipieceAPI.MultipieceShipmentReprint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**shipmentId** | **string** | It specifies the shipmentId of onward shipment against which return label has to be created. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMultipieceShipmentReprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xPBDeveloperPartnerID** | **string** | This is the Develover Partner ID. When the developer is the only partner, this field is not required. | 

### Return type

[**MultipieceDomesticShipmentResponse**](MultipieceDomesticShipmentResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

