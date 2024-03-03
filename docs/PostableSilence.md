# PostableSilence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Matchers** | [**[]Matcher**](Matcher.md) |  | 
**StartsAt** | **time.Time** |  | 
**EndsAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**Comment** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewPostableSilence

`func NewPostableSilence(matchers []Matcher, startsAt time.Time, endsAt time.Time, createdBy string, comment string, ) *PostableSilence`

NewPostableSilence instantiates a new PostableSilence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostableSilenceWithDefaults

`func NewPostableSilenceWithDefaults() *PostableSilence`

NewPostableSilenceWithDefaults instantiates a new PostableSilence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchers

`func (o *PostableSilence) GetMatchers() []Matcher`

GetMatchers returns the Matchers field if non-nil, zero value otherwise.

### GetMatchersOk

`func (o *PostableSilence) GetMatchersOk() (*[]Matcher, bool)`

GetMatchersOk returns a tuple with the Matchers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchers

`func (o *PostableSilence) SetMatchers(v []Matcher)`

SetMatchers sets Matchers field to given value.


### GetStartsAt

`func (o *PostableSilence) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PostableSilence) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PostableSilence) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *PostableSilence) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *PostableSilence) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *PostableSilence) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.


### GetCreatedBy

`func (o *PostableSilence) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PostableSilence) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PostableSilence) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetComment

`func (o *PostableSilence) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostableSilence) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostableSilence) SetComment(v string)`

SetComment sets Comment field to given value.


### GetId

`func (o *PostableSilence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostableSilence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostableSilence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostableSilence) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


