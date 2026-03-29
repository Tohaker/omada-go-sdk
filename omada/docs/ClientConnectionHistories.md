# ClientConnectionHistories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connections** | Pointer to [**[]ClientConnectionInfo**](ClientConnectionInfo.md) | Client connection histories | [optional] 
**Roamings** | Pointer to [**[]ClientRoamingInfo**](ClientRoamingInfo.md) | Client roaming histories | [optional] 

## Methods

### NewClientConnectionHistories

`func NewClientConnectionHistories() *ClientConnectionHistories`

NewClientConnectionHistories instantiates a new ClientConnectionHistories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientConnectionHistoriesWithDefaults

`func NewClientConnectionHistoriesWithDefaults() *ClientConnectionHistories`

NewClientConnectionHistoriesWithDefaults instantiates a new ClientConnectionHistories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnections

`func (o *ClientConnectionHistories) GetConnections() []ClientConnectionInfo`

GetConnections returns the Connections field if non-nil, zero value otherwise.

### GetConnectionsOk

`func (o *ClientConnectionHistories) GetConnectionsOk() (*[]ClientConnectionInfo, bool)`

GetConnectionsOk returns a tuple with the Connections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnections

`func (o *ClientConnectionHistories) SetConnections(v []ClientConnectionInfo)`

SetConnections sets Connections field to given value.

### HasConnections

`func (o *ClientConnectionHistories) HasConnections() bool`

HasConnections returns a boolean if a field has been set.

### GetRoamings

`func (o *ClientConnectionHistories) GetRoamings() []ClientRoamingInfo`

GetRoamings returns the Roamings field if non-nil, zero value otherwise.

### GetRoamingsOk

`func (o *ClientConnectionHistories) GetRoamingsOk() (*[]ClientRoamingInfo, bool)`

GetRoamingsOk returns a tuple with the Roamings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoamings

`func (o *ClientConnectionHistories) SetRoamings(v []ClientRoamingInfo)`

SetRoamings sets Roamings field to given value.

### HasRoamings

`func (o *ClientConnectionHistories) HasRoamings() bool`

HasRoamings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


