# PostableAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **map[string]string** |  | 
**GeneratorURL** | Pointer to **string** |  | [optional] 
**StartsAt** | Pointer to **time.Time** |  | [optional] 
**EndsAt** | Pointer to **time.Time** |  | [optional] 
**Annotations** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewPostableAlert

`func NewPostableAlert(labels map[string]string, ) *PostableAlert`

NewPostableAlert instantiates a new PostableAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostableAlertWithDefaults

`func NewPostableAlertWithDefaults() *PostableAlert`

NewPostableAlertWithDefaults instantiates a new PostableAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *PostableAlert) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *PostableAlert) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *PostableAlert) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetGeneratorURL

`func (o *PostableAlert) GetGeneratorURL() string`

GetGeneratorURL returns the GeneratorURL field if non-nil, zero value otherwise.

### GetGeneratorURLOk

`func (o *PostableAlert) GetGeneratorURLOk() (*string, bool)`

GetGeneratorURLOk returns a tuple with the GeneratorURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorURL

`func (o *PostableAlert) SetGeneratorURL(v string)`

SetGeneratorURL sets GeneratorURL field to given value.

### HasGeneratorURL

`func (o *PostableAlert) HasGeneratorURL() bool`

HasGeneratorURL returns a boolean if a field has been set.

### GetStartsAt

`func (o *PostableAlert) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PostableAlert) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PostableAlert) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *PostableAlert) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetEndsAt

`func (o *PostableAlert) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *PostableAlert) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *PostableAlert) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *PostableAlert) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### GetAnnotations

`func (o *PostableAlert) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *PostableAlert) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *PostableAlert) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.

### HasAnnotations

`func (o *PostableAlert) HasAnnotations() bool`

HasAnnotations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


