# AlertStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**SilencedBy** | **[]string** |  | 
**InhibitedBy** | **[]string** |  | 

## Methods

### NewAlertStatus

`func NewAlertStatus(state string, silencedBy []string, inhibitedBy []string, ) *AlertStatus`

NewAlertStatus instantiates a new AlertStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertStatusWithDefaults

`func NewAlertStatusWithDefaults() *AlertStatus`

NewAlertStatusWithDefaults instantiates a new AlertStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *AlertStatus) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AlertStatus) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AlertStatus) SetState(v string)`

SetState sets State field to given value.


### GetSilencedBy

`func (o *AlertStatus) GetSilencedBy() []string`

GetSilencedBy returns the SilencedBy field if non-nil, zero value otherwise.

### GetSilencedByOk

`func (o *AlertStatus) GetSilencedByOk() (*[]string, bool)`

GetSilencedByOk returns a tuple with the SilencedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilencedBy

`func (o *AlertStatus) SetSilencedBy(v []string)`

SetSilencedBy sets SilencedBy field to given value.


### GetInhibitedBy

`func (o *AlertStatus) GetInhibitedBy() []string`

GetInhibitedBy returns the InhibitedBy field if non-nil, zero value otherwise.

### GetInhibitedByOk

`func (o *AlertStatus) GetInhibitedByOk() (*[]string, bool)`

GetInhibitedByOk returns a tuple with the InhibitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInhibitedBy

`func (o *AlertStatus) SetInhibitedBy(v []string)`

SetInhibitedBy sets InhibitedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


