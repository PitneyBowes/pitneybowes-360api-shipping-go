/*
Shipping APIs

### Introduction  The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.   Some of the key API operations available in the Shipping API includes: ### Shipment API  | Operation      | Description | | ----------- | ----------- |  | Get Carriers    | This operation fetches all onboarded carriers. Typically, user will use this service to get list of onboarded carriers and supported properties for those carriers.  |  | Get Countries | This operation fetches list of supported destination countries for a provided carrier and origin country.  |  | Get Services | This operation fetches a list of supported services for a carrier with respect to specific origin and destination country. |  | Get ParcelTypes| This operation fetches ParcelTypes based on carrier, origin and destination country. |  | Get Special Services| This operation fetches Special Services for a given carrier, service, origin and destination country. |  | Get Carrier Accounts| This operation retrieves onboarded Carriers with their Carrier Account Ids which uniquely identify multiple accounts of same carrier.  |  | Rate Shop and Get Single Rate| This API contains 2 operations, rate shop and single rate. Rate shop will fetch rates for all carrier services based on the given addresses (From and To), weight, and dimension for given parcelType. Single rate will get rate for specific service and special service (if requested) based on the given addresses (From and To), weight, and dimension, parcelType and serviceId with or without specialServices. Single rate will be used mainly to a rate a shipment before creating shipment.  |  | Create Shipment| This operation creates a new Shipment or Shipment Label. This is for both Domestic and International. | | Get All Shipments| This operation fetches all created Shipments. |  | Get Shipment by Id| Retrieves single shipment using Shipment Id. |  | Reprint Shipment| This operation reprints Shipment by the shipmentId. It retrieves an existing shipping label to reprint. The API sends the shipmentId returned by the original Created Shipment request. Use this only if the shipping label in the Create Shipment response was spoilt or lost. |  | Cancel Shipment| This operation cancels previously created shipment. |  

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package shipping

import (
	"encoding/json"
	validator "gopkg.in/validator.v2"
	"fmt"
)

// GetRatesRequest - struct for GetRatesRequest
type GetRatesRequest struct {
	RateShop *RateShop
	SingleRate *SingleRate
}

// RateShopAsGetRatesRequest is a convenience function that returns RateShop wrapped in GetRatesRequest
func RateShopAsGetRatesRequest(v *RateShop) GetRatesRequest {
	return GetRatesRequest{
		RateShop: v,
	}
}

// SingleRateAsGetRatesRequest is a convenience function that returns SingleRate wrapped in GetRatesRequest
func SingleRateAsGetRatesRequest(v *SingleRate) GetRatesRequest {
	return GetRatesRequest{
		SingleRate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *GetRatesRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into RateShop
	err = newStrictDecoder(data).Decode(&dst.RateShop)
	if err == nil {
		jsonRateShop, _ := json.Marshal(dst.RateShop)
		if string(jsonRateShop) == "{}" { // empty struct
			dst.RateShop = nil
		} else {
			if err = validator.Validate(dst.RateShop); err != nil {
				dst.RateShop = nil
			} else {
				match++
			}
		}
	} else {
		dst.RateShop = nil
	}

	// try to unmarshal data into SingleRate
	err = newStrictDecoder(data).Decode(&dst.SingleRate)
	if err == nil {
		jsonSingleRate, _ := json.Marshal(dst.SingleRate)
		if string(jsonSingleRate) == "{}" { // empty struct
			dst.SingleRate = nil
		} else {
			if err = validator.Validate(dst.SingleRate); err != nil {
				dst.SingleRate = nil
			} else {
				match++
			}
		}
	} else {
		dst.SingleRate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.RateShop = nil
		dst.SingleRate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GetRatesRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GetRatesRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GetRatesRequest) MarshalJSON() ([]byte, error) {
	if src.RateShop != nil {
		return json.Marshal(&src.RateShop)
	}

	if src.SingleRate != nil {
		return json.Marshal(&src.SingleRate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GetRatesRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RateShop != nil {
		return obj.RateShop
	}

	if obj.SingleRate != nil {
		return obj.SingleRate
	}

	// all schemas are nil
	return nil
}

type NullableGetRatesRequest struct {
	value *GetRatesRequest
	isSet bool
}

func (v NullableGetRatesRequest) Get() *GetRatesRequest {
	return v.value
}

func (v *NullableGetRatesRequest) Set(val *GetRatesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRatesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRatesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRatesRequest(val *GetRatesRequest) *NullableGetRatesRequest {
	return &NullableGetRatesRequest{value: val, isSet: true}
}

func (v NullableGetRatesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRatesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


