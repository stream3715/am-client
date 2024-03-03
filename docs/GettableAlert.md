# GettableAlert

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **map[string]string** |  | 
**GeneratorURL** | Pointer to **string** |  | [optional] 
**Annotations** | **map[string]string** |  | 
**Receivers** | [**[]Receiver**](Receiver.md) |  | 
**Fingerprint** | **string** |  | 
**StartsAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**EndsAt** | **time.Time** |  | 
**Status** | [**AlertStatus**](AlertStatus.md) |  | 

## Methods

### NewGettableAlert

`func NewGettableAlert(labels map[string]string, annotations map[string]string, receivers []Receiver, fingerprint string, startsAt time.Time, updatedAt time.Time, endsAt time.Time, status AlertStatus, ) *GettableAlert`

NewGettableAlert instantiates a new GettableAlert object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGettableAlertWithDefaults

`func NewGettableAlertWithDefaults() *GettableAlert`

NewGettableAlertWithDefaults instantiates a new GettableAlert object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *GettableAlert) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GettableAlert) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GettableAlert) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetGeneratorURL

`func (o *GettableAlert) GetGeneratorURL() string`

GetGeneratorURL returns the GeneratorURL field if non-nil, zero value otherwise.

### GetGeneratorURLOk

`func (o *GettableAlert) GetGeneratorURLOk() (*string, bool)`

GetGeneratorURLOk returns a tuple with the GeneratorURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratorURL

`func (o *GettableAlert) SetGeneratorURL(v string)`

SetGeneratorURL sets GeneratorURL field to given value.

### HasGeneratorURL

`func (o *GettableAlert) HasGeneratorURL() bool`

HasGeneratorURL returns a boolean if a field has been set.

### GetAnnotations

`func (o *GettableAlert) GetAnnotations() map[string]string`

GetAnnotations returns the Annotations field if non-nil, zero value otherwise.

### GetAnnotationsOk

`func (o *GettableAlert) GetAnnotationsOk() (*map[string]string, bool)`

GetAnnotationsOk returns a tuple with the Annotations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnotations

`func (o *GettableAlert) SetAnnotations(v map[string]string)`

SetAnnotations sets Annotations field to given value.


### GetReceivers

`func (o *GettableAlert) GetReceivers() []Receiver`

GetReceivers returns the Receivers field if non-nil, zero value otherwise.

### GetReceiversOk

`func (o *GettableAlert) GetReceiversOk() (*[]Receiver, bool)`

GetReceiversOk returns a tuple with the Receivers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivers

`func (o *GettableAlert) SetReceivers(v []Receiver)`

SetReceivers sets Receivers field to given value.


### GetFingerprint

`func (o *GettableAlert) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GettableAlert) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GettableAlert) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.


### GetStartsAt

`func (o *GettableAlert) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GettableAlert) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GettableAlert) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetUpdatedAt

`func (o *GettableAlert) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GettableAlert) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GettableAlert) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEndsAt

`func (o *GettableAlert) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *GettableAlert) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *GettableAlert) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetStatus

`func (o *GettableAlert) GetStatus() AlertStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GettableAlert) GetStatusOk() (*AlertStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GettableAlert) SetStatus(v AlertStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


