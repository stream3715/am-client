# Silence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matchers** | [**[]Matcher**](Matcher.md) |  | 
**StartsAt** | **time.Time** |  | 
**EndsAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Comment** | **string** |  | 

## Methods

### NewSilence

`func NewSilence(matchers []Matcher, startsAt time.Time, endsAt time.Time, createdBy string, comment string, ) *Silence`

NewSilence instantiates a new Silence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSilenceWithDefaults

`func NewSilenceWithDefaults() *Silence`

NewSilenceWithDefaults instantiates a new Silence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchers

`func (o *Silence) GetMatchers() []Matcher`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *Silence) GetMatchersOk() (*[]Matcher, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *Silence) SetMatchers(v []Matcher)`

SetMatchers sets Matchers field to given value.


### GetStartsAt

`func (o *Silence) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *Silence) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *Silence) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *Silence) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *Silence) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *Silence) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetCreatedBy

`func (o *Silence) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Silence) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Silence) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetComment

`func (o *Silence) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Silence) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Silence) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


