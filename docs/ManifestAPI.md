# \ManifestAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateManifest**](ManifestAPI.md#CreateManifest) | **Post** /api/v1/manifests | Create Manifest
[**ReprintManifest**](ManifestAPI.md#ReprintManifest) | **Post** /api/v1/manifests/reprint | Reprint manifest



## CreateManifest

> CreateManifest200Response CreateManifest(ctx).XPBDeveloperPartnerID(xPBDeveloperPartnerID).CompactResponse(compactResponse).CreateManifestRequest(createManifestRequest).Execute()

Create Manifest



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
	compactResponse := false // bool | This header defines if the response required is detailed or compact. When value is set to true, it will only return manifest details in response.  (optional)
	createManifestRequest := *openapiclient.NewCreateManifestRequest("abcd12", *openapiclient.NewCreateManifestRequestFromAddress("27 Watervw Dr", "US", "06484")) // CreateManifestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManifestAPI.CreateManifest(context.Background()).XPBDeveloperPartnerID(xPBDeveloperPartnerID).CompactResponse(compactResponse).CreateManifestRequest(createManifestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.CreateManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateManifest`: CreateManifest200Response
	fmt.Fprintf(os.Stdout, "Response from `ManifestAPI.CreateManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **compactResponse** | **bool** | This header defines if the response required is detailed or compact. When value is set to true, it will only return manifest details in response.  | 
 **createManifestRequest** | [**CreateManifestRequest**](CreateManifestRequest.md) |  | 

### Return type

[**CreateManifest200Response**](CreateManifest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReprintManifest

> CreateManifest200Response ReprintManifest(ctx).XPBDeveloperPartnerID(xPBDeveloperPartnerID).CompactResponse(compactResponse).ReprintManifestRequest(reprintManifestRequest).Execute()

Reprint manifest



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
	compactResponse := false // bool | This header defines if the response required is detailed or compact. When value is set to true, it will only return manifest details in response. (optional)
	reprintManifestRequest := *openapiclient.NewReprintManifestRequest("JORx6bVG8mr", "1234567890") // ReprintManifestRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ManifestAPI.ReprintManifest(context.Background()).XPBDeveloperPartnerID(xPBDeveloperPartnerID).CompactResponse(compactResponse).ReprintManifestRequest(reprintManifestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ManifestAPI.ReprintManifest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReprintManifest`: CreateManifest200Response
	fmt.Fprintf(os.Stdout, "Response from `ManifestAPI.ReprintManifest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReprintManifestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerID** | **string** | This is the Developer Partner ID. When the developer is the only partner, this field is not required. | 
 **compactResponse** | **bool** | This header defines if the response required is detailed or compact. When value is set to true, it will only return manifest details in response. | 
 **reprintManifestRequest** | [**ReprintManifestRequest**](ReprintManifestRequest.md) |  | 

### Return type

[**CreateManifest200Response**](CreateManifest200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

