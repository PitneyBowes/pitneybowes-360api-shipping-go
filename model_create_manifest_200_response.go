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

// CreateManifest200Response - struct for CreateManifest200Response
type CreateManifest200Response struct {
	ManifestCompactResponse *ManifestCompactResponse
	ManifestDetailedResponse *ManifestDetailedResponse
}

// ManifestCompactResponseAsCreateManifest200Response is a convenience function that returns ManifestCompactResponse wrapped in CreateManifest200Response
func ManifestCompactResponseAsCreateManifest200Response(v *ManifestCompactResponse) CreateManifest200Response {
	return CreateManifest200Response{
		ManifestCompactResponse: v,
	}
}

// ManifestDetailedResponseAsCreateManifest200Response is a convenience function that returns ManifestDetailedResponse wrapped in CreateManifest200Response
func ManifestDetailedResponseAsCreateManifest200Response(v *ManifestDetailedResponse) CreateManifest200Response {
	return CreateManifest200Response{
		ManifestDetailedResponse: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateManifest200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ManifestCompactResponse
	err = newStrictDecoder(data).Decode(&dst.ManifestCompactResponse)
	if err == nil {
		jsonManifestCompactResponse, _ := json.Marshal(dst.ManifestCompactResponse)
		if string(jsonManifestCompactResponse) == "{}" { // empty struct
			dst.ManifestCompactResponse = nil
		} else {
			if err = validator.Validate(dst.ManifestCompactResponse); err != nil {
				dst.ManifestCompactResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ManifestCompactResponse = nil
	}

	// try to unmarshal data into ManifestDetailedResponse
	err = newStrictDecoder(data).Decode(&dst.ManifestDetailedResponse)
	if err == nil {
		jsonManifestDetailedResponse, _ := json.Marshal(dst.ManifestDetailedResponse)
		if string(jsonManifestDetailedResponse) == "{}" { // empty struct
			dst.ManifestDetailedResponse = nil
		} else {
			if err = validator.Validate(dst.ManifestDetailedResponse); err != nil {
				dst.ManifestDetailedResponse = nil
			} else {
				match++
			}
		}
	} else {
		dst.ManifestDetailedResponse = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ManifestCompactResponse = nil
		dst.ManifestDetailedResponse = nil

		return fmt.Errorf("data matches more than one schema in oneOf(CreateManifest200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(CreateManifest200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateManifest200Response) MarshalJSON() ([]byte, error) {
	if src.ManifestCompactResponse != nil {
		return json.Marshal(&src.ManifestCompactResponse)
	}

	if src.ManifestDetailedResponse != nil {
		return json.Marshal(&src.ManifestDetailedResponse)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateManifest200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ManifestCompactResponse != nil {
		return obj.ManifestCompactResponse
	}

	if obj.ManifestDetailedResponse != nil {
		return obj.ManifestDetailedResponse
	}

	// all schemas are nil
	return nil
}

type NullableCreateManifest200Response struct {
	value *CreateManifest200Response
	isSet bool
}

func (v NullableCreateManifest200Response) Get() *CreateManifest200Response {
	return v.value
}

func (v *NullableCreateManifest200Response) Set(val *CreateManifest200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateManifest200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateManifest200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateManifest200Response(val *CreateManifest200Response) *NullableCreateManifest200Response {
	return &NullableCreateManifest200Response{value: val, isSet: true}
}

func (v NullableCreateManifest200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateManifest200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


