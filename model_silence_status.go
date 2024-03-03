/*
Alertmanager API

API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager)

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SilenceStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SilenceStatus{}

// SilenceStatus struct for SilenceStatus
type SilenceStatus struct {
	State string `json:"state"`
}

type _SilenceStatus SilenceStatus

// NewSilenceStatus instantiates a new SilenceStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSilenceStatus(state string) *SilenceStatus {
	this := SilenceStatus{}
	this.State = state
	return &this
}

// NewSilenceStatusWithDefaults instantiates a new SilenceStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSilenceStatusWithDefaults() *SilenceStatus {
	this := SilenceStatus{}
	return &this
}

// GetState returns the State field value
func (o *SilenceStatus) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *SilenceStatus) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *SilenceStatus) SetState(v string) {
	o.State = v
}

func (o SilenceStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SilenceStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["state"] = o.State
	return toSerialize, nil
}

func (o *SilenceStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"state",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSilenceStatus := _SilenceStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSilenceStatus)

	if err != nil {
		return err
	}

	*o = SilenceStatus(varSilenceStatus)

	return err
}

type NullableSilenceStatus struct {
	value *SilenceStatus
	isSet bool
}

func (v NullableSilenceStatus) Get() *SilenceStatus {
	return v.value
}

func (v *NullableSilenceStatus) Set(val *SilenceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSilenceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSilenceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSilenceStatus(val *SilenceStatus) *NullableSilenceStatus {
	return &NullableSilenceStatus{value: val, isSet: true}
}

func (v NullableSilenceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSilenceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

