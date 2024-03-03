# AlertmanagerStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | [**ClusterStatus**](ClusterStatus.md) |  | 
**VersionInfo** | [**VersionInfo**](VersionInfo.md) |  | 
**Config** | [**AlertmanagerConfig**](AlertmanagerConfig.md) |  | 
**Uptime** | **time.Time** |  | 

## Methods

### NewAlertmanagerStatus

`func NewAlertmanagerStatus(cluster ClusterStatus, versionInfo VersionInfo, config AlertmanagerConfig, uptime time.Time, ) *AlertmanagerStatus`

NewAlertmanagerStatus instantiates a new AlertmanagerStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertmanagerStatusWithDefaults

`func NewAlertmanagerStatusWithDefaults() *AlertmanagerStatus`

NewAlertmanagerStatusWithDefaults instantiates a new AlertmanagerStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *AlertmanagerStatus) GetCluster() ClusterStatus`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *AlertmanagerStatus) GetClusterOk() (*ClusterStatus, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *AlertmanagerStatus) SetCluster(v ClusterStatus)`

SetCluster sets Cluster field to given value.


### GetVersionInfo

`func (o *AlertmanagerStatus) GetVersionInfo() VersionInfo`

GetVersionInfo returns the VersionInfo field if non-nil, zero value otherwise.

### GetVersionInfoOk

`func (o *AlertmanagerStatus) GetVersionInfoOk() (*VersionInfo, bool)`

GetVersionInfoOk returns a tuple with the VersionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionInfo

`func (o *AlertmanagerStatus) SetVersionInfo(v VersionInfo)`

SetVersionInfo sets VersionInfo field to given value.


### GetConfig

`func (o *AlertmanagerStatus) GetConfig() AlertmanagerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AlertmanagerStatus) GetConfigOk() (*AlertmanagerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AlertmanagerStatus) SetConfig(v AlertmanagerConfig)`

SetConfig sets Config field to given value.


### GetUptime

`func (o *AlertmanagerStatus) GetUptime() time.Time`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *AlertmanagerStatus) GetUptimeOk() (*time.Time, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *AlertmanagerStatus) SetUptime(v time.Time)`

SetUptime sets Uptime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


