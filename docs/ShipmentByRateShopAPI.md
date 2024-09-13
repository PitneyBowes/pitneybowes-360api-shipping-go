# \ShipmentByRateShopAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelShipmentByIdV2**](ShipmentByRateShopAPI.md#CancelShipmentByIdV2) | **Post** /api/v2/shipments/cancel | Cancel Shipment
[**CreateShipmentV2**](ShipmentByRateShopAPI.md#CreateShipmentV2) | **Post** /api/v2/shipments | Create Shipment
[**ReprintShipmentByIdV2**](ShipmentByRateShopAPI.md#ReprintShipmentByIdV2) | **Post** /api/v2/shipments/reprint | Reprint Shipment



## CancelShipmentByIdV2

> CancelShipmentV2 CancelShipmentByIdV2(ctx).ShipmentCancelV2(shipmentCancelV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Cancel Shipment



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
	shipmentCancelV2 := *openapiclient.NewShipmentCancelV2("PUROLATOR2200626353009030") // ShipmentCancelV2 | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique Transaction ID provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentByRateShopAPI.CancelShipmentByIdV2(context.Background()).ShipmentCancelV2(shipmentCancelV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentByRateShopAPI.CancelShipmentByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CancelShipmentByIdV2`: CancelShipmentV2
	fmt.Fprintf(os.Stdout, "Response from `ShipmentByRateShopAPI.CancelShipmentByIdV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCancelShipmentByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentCancelV2** | [**ShipmentCancelV2**](ShipmentCancelV2.md) |  | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique Transaction ID provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**CancelShipmentV2**](CancelShipmentV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateShipmentV2

> DomesticShipmentResponseV2 CreateShipmentV2(ctx).CreateShipmentV2Request(createShipmentV2Request).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Create Shipment



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
	createShipmentV2Request := openapiclient.createShipmentV2_request{ShipmentDomesticByCarrier: openapiclient.NewShipmentDomesticByCarrier(*openapiclient.NewFromAddressV2("Paul Wright", "638 Manitoba Ave", "Winnipeg", "MB", "R2W 2H2", "CA", "203-555-1213"), *openapiclient.NewToAddressV2("Paul Wright", "638 Manitoba Ave", "Winnipeg", "MB", "R2W 2H2", "CA", "203-555-1213"), "PKG", "carrier", "DOC_4X6", "SHIPPING_LABEL", "ZPL2")} // CreateShipmentV2Request | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentByRateShopAPI.CreateShipmentV2(context.Background()).CreateShipmentV2Request(createShipmentV2Request).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentByRateShopAPI.CreateShipmentV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateShipmentV2`: DomesticShipmentResponseV2
	fmt.Fprintf(os.Stdout, "Response from `ShipmentByRateShopAPI.CreateShipmentV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShipmentV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShipmentV2Request** | [**CreateShipmentV2Request**](CreateShipmentV2Request.md) |  | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**DomesticShipmentResponseV2**](DomesticShipmentResponseV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReprintShipmentByIdV2

> ReprintShipmentV2 ReprintShipmentByIdV2(ctx).ShipmentReprintV2(shipmentReprintV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()

Reprint Shipment



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
	shipmentReprintV2 := *openapiclient.NewShipmentReprintV2("PUROLATOR2200626353009030") // ShipmentReprintV2 | 
	xPBDeveloperPartnerId := "xPBDeveloperPartnerId_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	xPBLocationId := "xPBLocationId_example" // string | This is the Location ID assigned as per the Developer's and Partner's parsed locations, to which all transactions will be billed. <br /> Partner's location will be used for billing if it is configured, however, in case Partner's location is not given, then the Developer's location will be taken. Developer's location will be the default value. <br /> Additionally, Developers and Partners can use carriers belong to this location only. (optional)
	xPBTransactionId := "xPBTransactionId_example" // string | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client's transaction and the system. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ShipmentByRateShopAPI.ReprintShipmentByIdV2(context.Background()).ShipmentReprintV2(shipmentReprintV2).XPBDeveloperPartnerId(xPBDeveloperPartnerId).XPBLocationId(xPBLocationId).XPBTransactionId(xPBTransactionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ShipmentByRateShopAPI.ReprintShipmentByIdV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReprintShipmentByIdV2`: ReprintShipmentV2
	fmt.Fprintf(os.Stdout, "Response from `ShipmentByRateShopAPI.ReprintShipmentByIdV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReprintShipmentByIdV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **shipmentReprintV2** | [**ShipmentReprintV2**](ShipmentReprintV2.md) |  | 
 **xPBDeveloperPartnerId** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **xPBLocationId** | **string** | This is the Location ID assigned as per the Developer&#39;s and Partner&#39;s parsed locations, to which all transactions will be billed. &lt;br /&gt; Partner&#39;s location will be used for billing if it is configured, however, in case Partner&#39;s location is not given, then the Developer&#39;s location will be taken. Developer&#39;s location will be the default value. &lt;br /&gt; Additionally, Developers and Partners can use carriers belong to this location only. | 
 **xPBTransactionId** | **string** | A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system. | 

### Return type

[**ReprintShipmentV2**](ReprintShipmentV2.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

