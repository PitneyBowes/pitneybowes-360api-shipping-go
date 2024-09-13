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

// CreateBulkShipmentsAPIRequest - struct for CreateBulkShipmentsAPIRequest
type CreateBulkShipmentsAPIRequest struct {
	CreateBulkShipmentInternational *CreateBulkShipmentInternational
	CreateBulkShipments *CreateBulkShipments
}

// CreateBulkShipmentInternationalAsCreateBulkShipmentsAPIRequest is a convenience function that returns CreateBulkShipmentInternational wrapped in CreateBulkShipmentsAPIRequest
func CreateBulkShipmentInternationalAsCreateBulkShipmentsAPIRequest(v *CreateBulkShipmentInternational) CreateBulkShipmentsAPIRequest {
	return CreateBulkShipmentsAPIRequest{
		CreateBulkShipmentInternational: v,
	}
}

// CreateBulkShipmentsAsCreateBulkShipmentsAPIRequest is a convenience function that returns CreateBulkShipments wrapped in CreateBulkShipmentsAPIRequest
func CreateBulkShipmentsAsCreateBulkShipmentsAPIRequest(v *CreateBulkShipments) CreateBulkShipmentsAPIRequest {
	return CreateBulkShipmentsAPIRequest{
		CreateBulkShipments: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateBulkShipmentsAPIRequest) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into CreateBulkShipmentInternational
	err = newStrictDecoder(data).Decode(&dst.CreateBulkShipmentInternational)
	if err == nil {
		jsonCreateBulkShipmentInternational, _ := json.Marshal(dst.CreateBulkShipmentInternational)
		if string(jsonCreateBulkShipmentInternational) == "{}" { // empty struct
			dst.CreateBulkShipmentInternational = nil
		} else {
			if err = validator.Validate(dst.CreateBulkShipmentInternational); err != nil {
				dst.CreateBulkShipmentInternational = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateBulkShipmentInternational = nil
	}

	// try to unmarshal data into CreateBulkShipments
	err = newStrictDecoder(data).Decode(&dst.CreateBulkShipments)
	if err == nil {
		jsonCreateBulkShipments, _ := json.Marshal(dst.CreateBulkShipments)
		if string(jsonCreateBulkShipments) == "{}" { // empty struct
			dst.CreateBulkShipments = nil
		} else {
			if err = validator.Validate(dst.CreateBulkShipments); err != nil {
				dst.CreateBulkShipments = nil
			} else {
				match++
			}
		}
	} else {
		dst.CreateBulkShipments = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.CreateBulkShipmentInternational = nil
		dst.CreateBulkShipments = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateBulkShipmentsAPIRequest)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateBulkShipmentsAPIRequest)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateBulkShipmentsAPIRequest) MarshalJSON() ([]byte, error) {
	if src.CreateBulkShipmentInternational != nil {
		return json.Marshal(&src.CreateBulkShipmentInternational)
	}

	if src.CreateBulkShipments != nil {
		return json.Marshal(&src.CreateBulkShipments)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateBulkShipmentsAPIRequest) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.CreateBulkShipmentInternational != nil {
		return obj.CreateBulkShipmentInternational
	}

	if obj.CreateBulkShipments != nil {
		return obj.CreateBulkShipments
	}

	// all schemas are nil
	return nil
}

type NullableCreateBulkShipmentsAPIRequest struct {
	value *CreateBulkShipmentsAPIRequest
	isSet bool
}

func (v NullableCreateBulkShipmentsAPIRequest) Get() *CreateBulkShipmentsAPIRequest {
	return v.value
}

func (v *NullableCreateBulkShipmentsAPIRequest) Set(val *CreateBulkShipmentsAPIRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBulkShipmentsAPIRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBulkShipmentsAPIRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBulkShipmentsAPIRequest(val *CreateBulkShipmentsAPIRequest) *NullableCreateBulkShipmentsAPIRequest {
	return &NullableCreateBulkShipmentsAPIRequest{value: val, isSet: true}
}

func (v NullableCreateBulkShipmentsAPIRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBulkShipmentsAPIRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


