# AlertGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Labels** | **map[string]string** |  | 
**Receiver** | [**Receiver**](Receiver.md) |  | 
**Alerts** | [**[]GettableAlert**](GettableAlert.md) |  | 

## Methods

### NewAlertGroup

`func NewAlertGroup(labels map[string]string, receiver Receiver, alerts []GettableAlert, ) *AlertGroup`

NewAlertGroup instantiates a new AlertGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertGroupWithDefaults

`func NewAlertGroupWithDefaults() *AlertGroup`

NewAlertGroupWithDefaults instantiates a new AlertGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabels

`func (o *AlertGroup) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AlertGroup) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AlertGroup) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.


### GetReceiver

`func (o *AlertGroup) GetReceiver() Receiver`

GetReceiver returns the Receiver field if non-nil, zero value otherwise.

### GetReceiverOk

`func (o *AlertGroup) GetReceiverOk() (*Receiver, bool)`

GetReceiverOk returns a tuple with the Receiver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiver

`func (o *AlertGroup) SetReceiver(v Receiver)`

SetReceiver sets Receiver field to given value.


### GetAlerts

`func (o *AlertGroup) GetAlerts() []GettableAlert`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *AlertGroup) GetAlertsOk() (*[]GettableAlert, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *AlertGroup) SetAlerts(v []GettableAlert)`

SetAlerts sets Alerts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


