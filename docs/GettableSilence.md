# GettableSilence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matchers** | [**[]Matcher**](Matcher.md) |  | 
**StartsAt** | **time.Time** |  | 
**EndsAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Comment** | **string** |  | 
**Id** | **string** |  | 
**Status** | [**SilenceStatus**](SilenceStatus.md) |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewGettableSilence

`func NewGettableSilence(matchers []Matcher, startsAt time.Time, endsAt time.Time, createdBy string, comment string, id string, status SilenceStatus, updatedAt time.Time, ) *GettableSilence`

NewGettableSilence instantiates a new GettableSilence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGettableSilenceWithDefaults

`func NewGettableSilenceWithDefaults() *GettableSilence`

NewGettableSilenceWithDefaults instantiates a new GettableSilence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchers

`func (o *GettableSilence) GetMatchers() []Matcher`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *GettableSilence) GetMatchersOk() (*[]Matcher, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *GettableSilence) SetMatchers(v []Matcher)`

SetMatchers sets Matchers field to given value.


### GetStartsAt

`func (o *GettableSilence) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *GettableSilence) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *GettableSilence) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *GettableSilence) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *GettableSilence) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *GettableSilence) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetCreatedBy

`func (o *GettableSilence) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GettableSilence) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GettableSilence) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetComment

`func (o *GettableSilence) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *GettableSilence) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *GettableSilence) SetComment(v string)`

SetComment sets Comment field to given value.


### GetId

`func (o *GettableSilence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GettableSilence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GettableSilence) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *GettableSilence) GetStatus() SilenceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GettableSilence) GetStatusOk() (*SilenceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GettableSilence) SetStatus(v SilenceStatus)`

SetStatus sets Status field to given value.


### GetUpdatedAt

`func (o *GettableSilence) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GettableSilence) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GettableSilence) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


