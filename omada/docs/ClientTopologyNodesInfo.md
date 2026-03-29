# ClientTopologyNodesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to **bool** |  | [optional] 
**ClientNode** | Pointer to [**ClientNodeInfo**](ClientNodeInfo.md) |  | [optional] 
**DeviceNode** | Pointer to [**DeviceNodeInfo**](DeviceNodeInfo.md) |  | [optional] 
**NodeType** | Pointer to **int32** | Node type, 0: device; 1: client. | [optional] 

## Methods

### NewClientTopologyNodesInfo

`func NewClientTopologyNodesInfo() *ClientTopologyNodesInfo`

NewClientTopologyNodesInfo instantiates a new ClientTopologyNodesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientTopologyNodesInfoWithDefaults

`func NewClientTopologyNodesInfoWithDefaults() *ClientTopologyNodesInfo`

NewClientTopologyNodesInfoWithDefaults instantiates a new ClientTopologyNodesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *ClientTopologyNodesInfo) GetClient() bool`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *ClientTopologyNodesInfo) GetClientOk() (*bool, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *ClientTopologyNodesInfo) SetClient(v bool)`

SetClient sets Client field to given value.

### HasClient

`func (o *ClientTopologyNodesInfo) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetClientNode

`func (o *ClientTopologyNodesInfo) GetClientNode() ClientNodeInfo`

GetClientNode returns the ClientNode field if non-nil, zero value otherwise.

### GetClientNodeOk

`func (o *ClientTopologyNodesInfo) GetClientNodeOk() (*ClientNodeInfo, bool)`

GetClientNodeOk returns a tuple with the ClientNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNode

`func (o *ClientTopologyNodesInfo) SetClientNode(v ClientNodeInfo)`

SetClientNode sets ClientNode field to given value.

### HasClientNode

`func (o *ClientTopologyNodesInfo) HasClientNode() bool`

HasClientNode returns a boolean if a field has been set.

### GetDeviceNode

`func (o *ClientTopologyNodesInfo) GetDeviceNode() DeviceNodeInfo`

GetDeviceNode returns the DeviceNode field if non-nil, zero value otherwise.

### GetDeviceNodeOk

`func (o *ClientTopologyNodesInfo) GetDeviceNodeOk() (*DeviceNodeInfo, bool)`

GetDeviceNodeOk returns a tuple with the DeviceNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceNode

`func (o *ClientTopologyNodesInfo) SetDeviceNode(v DeviceNodeInfo)`

SetDeviceNode sets DeviceNode field to given value.

### HasDeviceNode

`func (o *ClientTopologyNodesInfo) HasDeviceNode() bool`

HasDeviceNode returns a boolean if a field has been set.

### GetNodeType

`func (o *ClientTopologyNodesInfo) GetNodeType() int32`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ClientTopologyNodesInfo) GetNodeTypeOk() (*int32, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ClientTopologyNodesInfo) SetNodeType(v int32)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ClientTopologyNodesInfo) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


