# \DefaultsAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDefaults**](DefaultsAPI.md#CreateDefaults) | **Post** /api/v1/defaults | Create Defaults
[**DeleteDefaultsById**](DefaultsAPI.md#DeleteDefaultsById) | **Delete** /api/v1/defaults/{defaultID} | Delete Defaults by ID
[**GetAllDefaults**](DefaultsAPI.md#GetAllDefaults) | **Get** /api/v1/defaults | Get All Defaults
[**GetDefaultsById**](DefaultsAPI.md#GetDefaultsById) | **Get** /api/v1/defaults/{defaultID} | Get Defaults By ID
[**PutDefaultsById**](DefaultsAPI.md#PutDefaultsById) | **Put** /api/v1/defaults/{defaultID} | Update Defaults



## CreateDefaults

> CreateDefaultsResponse CreateDefaults(ctx).CreateDefaults(createDefaults).Execute()

Create Defaults



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
	createDefaults := *openapiclient.NewCreateDefaults() // CreateDefaults | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultsAPI.CreateDefaults(context.Background()).CreateDefaults(createDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultsAPI.CreateDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDefaults`: CreateDefaultsResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultsAPI.CreateDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDefaults** | [**CreateDefaults**](CreateDefaults.md) |  | 

### Return type

[**CreateDefaultsResponse**](CreateDefaultsResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDefaultsById

> DeleteDefaultsById(ctx, defaultID).Execute()

Delete Defaults by ID



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
	defaultID := "defaultID_example" // string | This is unique identifier assigned to Defaults while its creation using CreateDefaults API.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultsAPI.DeleteDefaultsById(context.Background(), defaultID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultsAPI.DeleteDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultID** | **string** | This is unique identifier assigned to Defaults while its creation using CreateDefaults API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetAllDefaults

> AllDefaults GetAllDefaults(ctx).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Page(page).Size(size).Execute()

Get All Defaults



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
	xPBDeveloperPartnerID := "xPBDeveloperPartnerID_example" // string | The Developer Partner ID is assigned by PB to uniquely identify a Developer's strategic business partners. If the developer is the sole business partner, this field isn't required. (optional)
	page := "page_example" // string | The page of the Defaults search result list. (optional)
	size := "size_example" // string | The size/count of the searched result list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultsAPI.GetAllDefaults(context.Background()).XPBDeveloperPartnerID(xPBDeveloperPartnerID).Page(page).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultsAPI.GetAllDefaults``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDefaults`: AllDefaults
	fmt.Fprintf(os.Stdout, "Response from `DefaultsAPI.GetAllDefaults`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDefaultsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xPBDeveloperPartnerID** | **string** | The Developer Partner ID is assigned by PB to uniquely identify a Developer&#39;s strategic business partners. If the developer is the sole business partner, this field isn&#39;t required. | 
 **page** | **string** | The page of the Defaults search result list. | 
 **size** | **string** | The size/count of the searched result list. | 

### Return type

[**AllDefaults**](AllDefaults.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultsById

> DefaultResponse GetDefaultsById(ctx, defaultID).Execute()

Get Defaults By ID



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
	defaultID := "defaultID_example" // string | This is unique identifier assigned to Defaults while its creation using CreateDefaults API.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultsAPI.GetDefaultsById(context.Background(), defaultID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultsAPI.GetDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultsById`: DefaultResponse
	fmt.Fprintf(os.Stdout, "Response from `DefaultsAPI.GetDefaultsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultID** | **string** | This is unique identifier assigned to Defaults while its creation using CreateDefaults API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DefaultResponse**](DefaultResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutDefaultsById

> PutDefaultsById(ctx, defaultID).CreateDefaults(createDefaults).Execute()

Update Defaults



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
	defaultID := "defaultID_example" // string | This is unique identifier assigned to Defaults while its creation using CreateDefaults API.
	createDefaults := *openapiclient.NewCreateDefaults() // CreateDefaults | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DefaultsAPI.PutDefaultsById(context.Background(), defaultID).CreateDefaults(createDefaults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultsAPI.PutDefaultsById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**defaultID** | **string** | This is unique identifier assigned to Defaults while its creation using CreateDefaults API. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutDefaultsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDefaults** | [**CreateDefaults**](CreateDefaults.md) |  | 

### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

