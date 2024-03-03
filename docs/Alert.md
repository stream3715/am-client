# Alert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **map[string]string** |  | 
**GeneratorURL** | Pointer to **string** |  | [optional] 

## Methods

### NewAlert

`func NewAlert(labels map[string]string, ) *Alert`

NewAlert instantiates a new Alert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertWithDefaults

`func NewAlertWithDefaults() *Alert`

NewAlertWithDefaults instantiates a new Alert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *Alert) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *Alert) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *Alert) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetGeneratorURL

`func (o *Alert) GetGeneratorURL() string`

GetGeneratorURL returns the GeneratorURL field if non-nil, zero value otherwise.

### GetGeneratorURLOk

`func (o *Alert) GetGeneratorURLOk() (*string, bool)`

GetGeneratorURLOk returns a tuple with the GeneratorURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorURL

`func (o *Alert) SetGeneratorURL(v string)`

SetGeneratorURL sets GeneratorURL field to given value.

### HasGeneratorURL

`func (o *Alert) HasGeneratorURL() bool`

HasGeneratorURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


