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

// SchedulePickupRequest - struct for SchedulePickupRequest
type SchedulePickupRequest struct {
	SchedulePickupDHLEXPRequest *SchedulePickupDHLEXPRequest
	SchedulePickupFedexRequest *SchedulePickupFedexRequest
	SchedulePickupUPSRequest *SchedulePickupUPSRequest
	SchedulePickupUSPSRequest *SchedulePickupUSPSRequest
}

// SchedulePickupDHLEXPRequestAsSchedulePickupRequest is a convenience function that returns SchedulePickupDHLEXPRequest wrapped in SchedulePickupRequest
func SchedulePickupDHLEXPRequestAsSchedulePickupRequest(v *SchedulePickupDHLEXPRequest) SchedulePickupRequest {
	return SchedulePickupRequest{
		SchedulePickupDHLEXPRequest: v,
	}
}

// SchedulePickupFedexRequestAsSchedulePickupRequest is a convenience function that returns SchedulePickupFedexRequest wrapped in SchedulePickupRequest
func SchedulePickupFedexRequestAsSchedulePickupRequest(v *SchedulePickupFedexRequest) SchedulePickupRequest {
	return SchedulePickupRequest{
		SchedulePickupFedexRequest: v,
	}
}

// SchedulePickupUPSRequestAsSchedulePickupRequest is a convenience function that returns SchedulePickupUPSRequest wrapped in SchedulePickupRequest
func SchedulePickupUPSRequestAsSchedulePickupRequest(v *SchedulePickupUPSRequest) SchedulePickupRequest {
	return SchedulePickupRequest{
		SchedulePickupUPSRequest: v,
	}
}

// SchedulePickupUSPSRequestAsSchedulePickupRequest is a convenience function that returns SchedulePickupUSPSRequest wrapped in SchedulePickupRequest
func SchedulePickupUSPSRequestAsSchedulePickupRequest(v *SchedulePickupUSPSRequest) SchedulePickupRequest {
	return SchedulePickupRequest{
		SchedulePickupUSPSRequest: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SchedulePickupRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into SchedulePickupDHLEXPRequest
	err = newStrictDecoder(data).Decode(&dst.SchedulePickupDHLEXPRequest)
	if err == nil {
		jsonSchedulePickupDHLEXPRequest, _ := json.Marshal(dst.SchedulePickupDHLEXPRequest)
		if string(jsonSchedulePickupDHLEXPRequest) == "{}" { // empty struct
			dst.SchedulePickupDHLEXPRequest = nil
		} else {
			if err = validator.Validate(dst.SchedulePickupDHLEXPRequest); err != nil {
				dst.SchedulePickupDHLEXPRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SchedulePickupDHLEXPRequest = nil
	}

	// try to unmarshal data into SchedulePickupFedexRequest
	err = newStrictDecoder(data).Decode(&dst.SchedulePickupFedexRequest)
	if err == nil {
		jsonSchedulePickupFedexRequest, _ := json.Marshal(dst.SchedulePickupFedexRequest)
		if string(jsonSchedulePickupFedexRequest) == "{}" { // empty struct
			dst.SchedulePickupFedexRequest = nil
		} else {
			if err = validator.Validate(dst.SchedulePickupFedexRequest); err != nil {
				dst.SchedulePickupFedexRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SchedulePickupFedexRequest = nil
	}

	// try to unmarshal data into SchedulePickupUPSRequest
	err = newStrictDecoder(data).Decode(&dst.SchedulePickupUPSRequest)
	if err == nil {
		jsonSchedulePickupUPSRequest, _ := json.Marshal(dst.SchedulePickupUPSRequest)
		if string(jsonSchedulePickupUPSRequest) == "{}" { // empty struct
			dst.SchedulePickupUPSRequest = nil
		} else {
			if err = validator.Validate(dst.SchedulePickupUPSRequest); err != nil {
				dst.SchedulePickupUPSRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SchedulePickupUPSRequest = nil
	}

	// try to unmarshal data into SchedulePickupUSPSRequest
	err = newStrictDecoder(data).Decode(&dst.SchedulePickupUSPSRequest)
	if err == nil {
		jsonSchedulePickupUSPSRequest, _ := json.Marshal(dst.SchedulePickupUSPSRequest)
		if string(jsonSchedulePickupUSPSRequest) == "{}" { // empty struct
			dst.SchedulePickupUSPSRequest = nil
		} else {
			if err = validator.Validate(dst.SchedulePickupUSPSRequest); err != nil {
				dst.SchedulePickupUSPSRequest = nil
			} else {
				match++
			}
		}
	} else {
		dst.SchedulePickupUSPSRequest = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.SchedulePickupDHLEXPRequest = nil
		dst.SchedulePickupFedexRequest = nil
		dst.SchedulePickupUPSRequest = nil
		dst.SchedulePickupUSPSRequest = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SchedulePickupRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SchedulePickupRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SchedulePickupRequest) MarshalJSON() ([]byte, error) {
	if src.SchedulePickupDHLEXPRequest != nil {
		return json.Marshal(&src.SchedulePickupDHLEXPRequest)
	}

	if src.SchedulePickupFedexRequest != nil {
		return json.Marshal(&src.SchedulePickupFedexRequest)
	}

	if src.SchedulePickupUPSRequest != nil {
		return json.Marshal(&src.SchedulePickupUPSRequest)
	}

	if src.SchedulePickupUSPSRequest != nil {
		return json.Marshal(&src.SchedulePickupUSPSRequest)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SchedulePickupRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.SchedulePickupDHLEXPRequest != nil {
		return obj.SchedulePickupDHLEXPRequest
	}

	if obj.SchedulePickupFedexRequest != nil {
		return obj.SchedulePickupFedexRequest
	}

	if obj.SchedulePickupUPSRequest != nil {
		return obj.SchedulePickupUPSRequest
	}

	if obj.SchedulePickupUSPSRequest != nil {
		return obj.SchedulePickupUSPSRequest
	}

	// all schemas are nil
	return nil
}

type NullableSchedulePickupRequest struct {
	value *SchedulePickupRequest
	isSet bool
}

func (v NullableSchedulePickupRequest) Get() *SchedulePickupRequest {
	return v.value
}

func (v *NullableSchedulePickupRequest) Set(val *SchedulePickupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSchedulePickupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSchedulePickupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchedulePickupRequest(val *SchedulePickupRequest) *NullableSchedulePickupRequest {
	return &NullableSchedulePickupRequest{value: val, isSet: true}
}

func (v NullableSchedulePickupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchedulePickupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


