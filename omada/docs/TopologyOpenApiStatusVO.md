# TopologyOpenApiStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client2g** | Pointer to **int32** | 2g Client Count | [optional] 
**Client5g** | Pointer to **int32** | 5g Client Count | [optional] 
**Client6g** | Pointer to **int32** | 6g Client Count | [optional] 
**Connected** | Pointer to **int32** | Connected Devices Count | [optional] 
**DirectlyShowClientsDeviceMacs** | Pointer to [**[]ClientsQueryMacAndFilterType**](ClientsQueryMacAndFilterType.md) | Devices Macs That Show Clients Directly | [optional] 
**Disconnected** | Pointer to **int32** | Disconnected Devices Count | [optional] 
**ExistUMVmsDev** | Pointer to **bool** | Whether Unmanaged Vms Devices Exist | [optional] 
**StpLoops** | Pointer to [**map[string][]TopologyOpenApiEdgeVO**](TopologyOpenApiEdgeVO.md) | Stp Loops | [optional] 
**WiredClient** | Pointer to **int32** | Wired Client Count | [optional] 

## Methods

### NewTopologyOpenApiStatusVO

`func NewTopologyOpenApiStatusVO() *TopologyOpenApiStatusVO`

NewTopologyOpenApiStatusVO instantiates a new TopologyOpenApiStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyOpenApiStatusVOWithDefaults

`func NewTopologyOpenApiStatusVOWithDefaults() *TopologyOpenApiStatusVO`

NewTopologyOpenApiStatusVOWithDefaults instantiates a new TopologyOpenApiStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient2g

`func (o *TopologyOpenApiStatusVO) GetClient2g() int32`

GetClient2g returns the Client2g field if non-nil, zero value otherwise.

### GetClient2gOk

`func (o *TopologyOpenApiStatusVO) GetClient2gOk() (*int32, bool)`

GetClient2gOk returns a tuple with the Client2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient2g

`func (o *TopologyOpenApiStatusVO) SetClient2g(v int32)`

SetClient2g sets Client2g field to given value.

### HasClient2g

`func (o *TopologyOpenApiStatusVO) HasClient2g() bool`

HasClient2g returns a boolean if a field has been set.

### GetClient5g

`func (o *TopologyOpenApiStatusVO) GetClient5g() int32`

GetClient5g returns the Client5g field if non-nil, zero value otherwise.

### GetClient5gOk

`func (o *TopologyOpenApiStatusVO) GetClient5gOk() (*int32, bool)`

GetClient5gOk returns a tuple with the Client5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient5g

`func (o *TopologyOpenApiStatusVO) SetClient5g(v int32)`

SetClient5g sets Client5g field to given value.

### HasClient5g

`func (o *TopologyOpenApiStatusVO) HasClient5g() bool`

HasClient5g returns a boolean if a field has been set.

### GetClient6g

`func (o *TopologyOpenApiStatusVO) GetClient6g() int32`

GetClient6g returns the Client6g field if non-nil, zero value otherwise.

### GetClient6gOk

`func (o *TopologyOpenApiStatusVO) GetClient6gOk() (*int32, bool)`

GetClient6gOk returns a tuple with the Client6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient6g

`func (o *TopologyOpenApiStatusVO) SetClient6g(v int32)`

SetClient6g sets Client6g field to given value.

### HasClient6g

`func (o *TopologyOpenApiStatusVO) HasClient6g() bool`

HasClient6g returns a boolean if a field has been set.

### GetConnected

`func (o *TopologyOpenApiStatusVO) GetConnected() int32`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *TopologyOpenApiStatusVO) GetConnectedOk() (*int32, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *TopologyOpenApiStatusVO) SetConnected(v int32)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *TopologyOpenApiStatusVO) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetDirectlyShowClientsDeviceMacs

`func (o *TopologyOpenApiStatusVO) GetDirectlyShowClientsDeviceMacs() []ClientsQueryMacAndFilterType`

GetDirectlyShowClientsDeviceMacs returns the DirectlyShowClientsDeviceMacs field if non-nil, zero value otherwise.

### GetDirectlyShowClientsDeviceMacsOk

`func (o *TopologyOpenApiStatusVO) GetDirectlyShowClientsDeviceMacsOk() (*[]ClientsQueryMacAndFilterType, bool)`

GetDirectlyShowClientsDeviceMacsOk returns a tuple with the DirectlyShowClientsDeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectlyShowClientsDeviceMacs

`func (o *TopologyOpenApiStatusVO) SetDirectlyShowClientsDeviceMacs(v []ClientsQueryMacAndFilterType)`

SetDirectlyShowClientsDeviceMacs sets DirectlyShowClientsDeviceMacs field to given value.

### HasDirectlyShowClientsDeviceMacs

`func (o *TopologyOpenApiStatusVO) HasDirectlyShowClientsDeviceMacs() bool`

HasDirectlyShowClientsDeviceMacs returns a boolean if a field has been set.

### GetDisconnected

`func (o *TopologyOpenApiStatusVO) GetDisconnected() int32`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *TopologyOpenApiStatusVO) GetDisconnectedOk() (*int32, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *TopologyOpenApiStatusVO) SetDisconnected(v int32)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *TopologyOpenApiStatusVO) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetExistUMVmsDev

`func (o *TopologyOpenApiStatusVO) GetExistUMVmsDev() bool`

GetExistUMVmsDev returns the ExistUMVmsDev field if non-nil, zero value otherwise.

### GetExistUMVmsDevOk

`func (o *TopologyOpenApiStatusVO) GetExistUMVmsDevOk() (*bool, bool)`

GetExistUMVmsDevOk returns a tuple with the ExistUMVmsDev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistUMVmsDev

`func (o *TopologyOpenApiStatusVO) SetExistUMVmsDev(v bool)`

SetExistUMVmsDev sets ExistUMVmsDev field to given value.

### HasExistUMVmsDev

`func (o *TopologyOpenApiStatusVO) HasExistUMVmsDev() bool`

HasExistUMVmsDev returns a boolean if a field has been set.

### GetStpLoops

`func (o *TopologyOpenApiStatusVO) GetStpLoops() map[string][]TopologyOpenApiEdgeVO`

GetStpLoops returns the StpLoops field if non-nil, zero value otherwise.

### GetStpLoopsOk

`func (o *TopologyOpenApiStatusVO) GetStpLoopsOk() (*map[string][]TopologyOpenApiEdgeVO, bool)`

GetStpLoopsOk returns a tuple with the StpLoops field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStpLoops

`func (o *TopologyOpenApiStatusVO) SetStpLoops(v map[string][]TopologyOpenApiEdgeVO)`

SetStpLoops sets StpLoops field to given value.

### HasStpLoops

`func (o *TopologyOpenApiStatusVO) HasStpLoops() bool`

HasStpLoops returns a boolean if a field has been set.

### GetWiredClient

`func (o *TopologyOpenApiStatusVO) GetWiredClient() int32`

GetWiredClient returns the WiredClient field if non-nil, zero value otherwise.

### GetWiredClientOk

`func (o *TopologyOpenApiStatusVO) GetWiredClientOk() (*int32, bool)`

GetWiredClientOk returns a tuple with the WiredClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredClient

`func (o *TopologyOpenApiStatusVO) SetWiredClient(v int32)`

SetWiredClient sets WiredClient field to given value.

### HasWiredClient

`func (o *TopologyOpenApiStatusVO) HasWiredClient() bool`

HasWiredClient returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


