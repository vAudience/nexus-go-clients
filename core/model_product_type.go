/*
Vaudience AI Core API

API for the Vaudience AI Core

API version: 0.2.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package core

import (
	"encoding/json"
	"fmt"
)

// ProductType the model 'ProductType'
type ProductType string

// List of ProductType
const (
	ProductTypePayment ProductType = "payment"
	ProductTypeStripeSubscription ProductType = "stripeSubscription"
	ProductTypeVaudSubscription ProductType = "vaudSubscription"
)

// All allowed values of ProductType enum
var AllowedProductTypeEnumValues = []ProductType{
	"payment",
	"stripeSubscription",
	"vaudSubscription",
}

func (v *ProductType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProductType(value)
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProductType", value)
}

// NewProductTypeFromValue returns a pointer to a valid ProductType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProductTypeFromValue(v string) (*ProductType, error) {
	ev := ProductType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProductType: valid values are %v", v, AllowedProductTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProductType) IsValid() bool {
	for _, existing := range AllowedProductTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProductType value
func (v ProductType) Ptr() *ProductType {
	return &v
}

type NullableProductType struct {
	value *ProductType
	isSet bool
}

func (v NullableProductType) Get() *ProductType {
	return v.value
}

func (v *NullableProductType) Set(val *ProductType) {
	v.value = val
	v.isSet = true
}

func (v NullableProductType) IsSet() bool {
	return v.isSet
}

func (v *NullableProductType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductType(val *ProductType) *NullableProductType {
	return &NullableProductType{value: val, isSet: true}
}

func (v NullableProductType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

