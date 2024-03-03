# ClusterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Status** | **string** |  | 
**Peers** | Pointer to [**[]PeerStatus**](PeerStatus.md) |  | [optional] 

## Methods

### NewClusterStatus

`func NewClusterStatus(status string, ) *ClusterStatus`

NewClusterStatus instantiates a new ClusterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterStatusWithDefaults

`func NewClusterStatusWithDefaults() *ClusterStatus`

NewClusterStatusWithDefaults instantiates a new ClusterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterStatus) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterStatus) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterStatus) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterStatus) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *ClusterStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusterStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusterStatus) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPeers

`func (o *ClusterStatus) GetPeers() []PeerStatus`

GetPeers returns the Peers field if non-nil, zero value otherwise.

### GetPeersOk

`func (o *ClusterStatus) GetPeersOk() (*[]PeerStatus, bool)`

GetPeersOk returns a tuple with the Peers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeers

`func (o *ClusterStatus) SetPeers(v []PeerStatus)`

SetPeers sets Peers field to given value.

### HasPeers

`func (o *ClusterStatus) HasPeers() bool`

HasPeers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


