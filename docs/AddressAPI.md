# \AddressAPI

All URIs are relative to *https://api-sandbox.sendpro360.pitneybowes.com/shipping*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressSuggest**](AddressAPI.md#AddressSuggest) | **Post** /api/v1/address/suggest | Address Suggest
[**AddressValidate**](AddressAPI.md#AddressValidate) | **Post** /api/v1/address/verify | Address Validate



## AddressSuggest

> AddressSuggestResponse AddressSuggest(ctx).AddressSuggestRequest(addressSuggestRequest).Execute()

Address Suggest



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
	addressSuggestRequest := *openapiclient.NewAddressSuggestRequest(*openapiclient.NewAddressSuggestRequestAddress("27 Watervw Dr", "Shelton", "US", "06484", "CT")) // AddressSuggestRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressSuggest(context.Background()).AddressSuggestRequest(addressSuggestRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressSuggest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressSuggest`: AddressSuggestResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressSuggest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressSuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressSuggestRequest** | [**AddressSuggestRequest**](AddressSuggestRequest.md) |  | 

### Return type

[**AddressSuggestResponse**](AddressSuggestResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddressValidate

> AddressValidateResponse AddressValidate(ctx).AddressValidateRequest(addressValidateRequest).MinimalAddressValidation(minimalAddressValidation).Execute()

Address Validate



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
	addressValidateRequest := *openapiclient.NewAddressValidateRequest("27 Watervw Dr", "US", "06484") // AddressValidateRequest | 
	minimalAddressValidation := true // bool | If true, then only City, State, and PostalCode (zip) are validated. This option is specific for US domestic addresses only. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressAPI.AddressValidate(context.Background()).AddressValidateRequest(addressValidateRequest).MinimalAddressValidation(minimalAddressValidation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressAPI.AddressValidate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressValidate`: AddressValidateResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressAPI.AddressValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressValidateRequest** | [**AddressValidateRequest**](AddressValidateRequest.md) |  | 
 **minimalAddressValidation** | **bool** | If true, then only City, State, and PostalCode (zip) are validated. This option is specific for US domestic addresses only. | 

### Return type

[**AddressValidateResponse**](AddressValidateResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

