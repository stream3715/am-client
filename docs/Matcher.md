# Matcher

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Value** | **string** |  | 
**IsRegex** | **bool** |  | 
**IsEqual** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewMatcher

`func NewMatcher(name string, value string, isRegex bool, ) *Matcher`

NewMatcher instantiates a new Matcher object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatcherWithDefaults

`func NewMatcherWithDefaults() *Matcher`

NewMatcherWithDefaults instantiates a new Matcher object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Matcher) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Matcher) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Matcher) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *Matcher) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Matcher) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Matcher) SetValue(v string)`

SetValue sets Value field to given value.


### GetIsRegex

`func (o *Matcher) GetIsRegex() bool`

GetIsRegex returns the IsRegex field if non-nil, zero value otherwise.

### GetIsRegexOk

`func (o *Matcher) GetIsRegexOk() (*bool, bool)`

GetIsRegexOk returns a tuple with the IsRegex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegex

`func (o *Matcher) SetIsRegex(v bool)`

SetIsRegex sets IsRegex field to given value.


### GetIsEqual

`func (o *Matcher) GetIsEqual() bool`

GetIsEqual returns the IsEqual field if non-nil, zero value otherwise.

### GetIsEqualOk

`func (o *Matcher) GetIsEqualOk() (*bool, bool)`

GetIsEqualOk returns a tuple with the IsEqual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEqual

`func (o *Matcher) SetIsEqual(v bool)`

SetIsEqual sets IsEqual field to given value.

### HasIsEqual

`func (o *Matcher) HasIsEqual() bool`

HasIsEqual returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


