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

// checks if the VersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionInfo{}

// VersionInfo struct for VersionInfo
type VersionInfo struct {
	Version string `json:"version"`
	Revision string `json:"revision"`
	Branch string `json:"branch"`
	BuildUser string `json:"buildUser"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
}

type _VersionInfo VersionInfo

// NewVersionInfo instantiates a new VersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfo(version string, revision string, branch string, buildUser string, buildDate string, goVersion string) *VersionInfo {
	this := VersionInfo{}
	this.Version = version
	this.Revision = revision
	this.Branch = branch
	this.BuildUser = buildUser
	this.BuildDate = buildDate
	this.GoVersion = goVersion
	return &this
}

// NewVersionInfoWithDefaults instantiates a new VersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoWithDefaults() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// GetVersion returns the Version field value
func (o *VersionInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VersionInfo) SetVersion(v string) {
	o.Version = v
}

// GetRevision returns the Revision field value
func (o *VersionInfo) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *VersionInfo) SetRevision(v string) {
	o.Revision = v
}

// GetBranch returns the Branch field value
func (o *VersionInfo) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *VersionInfo) SetBranch(v string) {
	o.Branch = v
}

// GetBuildUser returns the BuildUser field value
func (o *VersionInfo) GetBuildUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildUser
}

// GetBuildUserOk returns a tuple with the BuildUser field value
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetBuildUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildUser, true
}

// SetBuildUser sets field value
func (o *VersionInfo) SetBuildUser(v string) {
	o.BuildUser = v
}

// GetBuildDate returns the BuildDate field value
func (o *VersionInfo) GetBuildDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BuildDate
}

// GetBuildDateOk returns a tuple with the BuildDate field value
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetBuildDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BuildDate, true
}

// SetBuildDate sets field value
func (o *VersionInfo) SetBuildDate(v string) {
	o.BuildDate = v
}

// GetGoVersion returns the GoVersion field value
func (o *VersionInfo) GetGoVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GoVersion
}

// GetGoVersionOk returns a tuple with the GoVersion field value
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetGoVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GoVersion, true
}

// SetGoVersion sets field value
func (o *VersionInfo) SetGoVersion(v string) {
	o.GoVersion = v
}

func (o VersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["version"] = o.Version
	toSerialize["revision"] = o.Revision
	toSerialize["branch"] = o.Branch
	toSerialize["buildUser"] = o.BuildUser
	toSerialize["buildDate"] = o.BuildDate
	toSerialize["goVersion"] = o.GoVersion
	return toSerialize, nil
}

func (o *VersionInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"version",
		"revision",
		"branch",
		"buildUser",
		"buildDate",
		"goVersion",
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

	varVersionInfo := _VersionInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVersionInfo)

	if err != nil {
		return err
	}

	*o = VersionInfo(varVersionInfo)

	return err
}

type NullableVersionInfo struct {
	value *VersionInfo
	isSet bool
}

func (v NullableVersionInfo) Get() *VersionInfo {
	return v.value
}

func (v *NullableVersionInfo) Set(val *VersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfo(val *VersionInfo) *NullableVersionInfo {
	return &NullableVersionInfo{value: val, isSet: true}
}

func (v NullableVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


