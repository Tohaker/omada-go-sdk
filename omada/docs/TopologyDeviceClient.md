# TopologyDeviceClient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceType** | Pointer to **string** | Type of Client | [optional] 
**GroupType** | Pointer to **int32** | Topology group type of client, 0：clientGroup  1：cameraGroup | [optional] 
**Ip** | Pointer to **string** | Device or Client Ip. | [optional] 
**IsClient** | Pointer to **bool** | Whether this entry is a client | [optional] 
**Mac** | Pointer to **string** | Device or Client MAC address, like AA-BB-CC-DD-EE-FF. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Device model version. | [optional] 
**Name** | Pointer to **string** | Device or Client Name. | [optional] 
**OmadaDeviceType** | Pointer to **string** | Type of omada device when it&#39;s type is other in topology | [optional] 
**ShowModel** | Pointer to **string** | Device model. | [optional] 
**StackGroup** | Pointer to **bool** | Whether this node is stack or not | [optional] 
**Type** | Pointer to **string** | Type of Device | [optional] 
**UplinkMac** | Pointer to **string** | Uplink Device MAC address of client. | [optional] 

## Methods

### NewTopologyDeviceClient

`func NewTopologyDeviceClient() *TopologyDeviceClient`

NewTopologyDeviceClient instantiates a new TopologyDeviceClient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyDeviceClientWithDefaults

`func NewTopologyDeviceClientWithDefaults() *TopologyDeviceClient`

NewTopologyDeviceClientWithDefaults instantiates a new TopologyDeviceClient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceType

`func (o *TopologyDeviceClient) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *TopologyDeviceClient) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *TopologyDeviceClient) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *TopologyDeviceClient) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetGroupType

`func (o *TopologyDeviceClient) GetGroupType() int32`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *TopologyDeviceClient) GetGroupTypeOk() (*int32, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *TopologyDeviceClient) SetGroupType(v int32)`

SetGroupType sets GroupType field to given value.

### HasGroupType

`func (o *TopologyDeviceClient) HasGroupType() bool`

HasGroupType returns a boolean if a field has been set.

### GetIp

`func (o *TopologyDeviceClient) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *TopologyDeviceClient) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *TopologyDeviceClient) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *TopologyDeviceClient) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIsClient

`func (o *TopologyDeviceClient) GetIsClient() bool`

GetIsClient returns the IsClient field if non-nil, zero value otherwise.

### GetIsClientOk

`func (o *TopologyDeviceClient) GetIsClientOk() (*bool, bool)`

GetIsClientOk returns a tuple with the IsClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClient

`func (o *TopologyDeviceClient) SetIsClient(v bool)`

SetIsClient sets IsClient field to given value.

### HasIsClient

`func (o *TopologyDeviceClient) HasIsClient() bool`

HasIsClient returns a boolean if a field has been set.

### GetMac

`func (o *TopologyDeviceClient) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyDeviceClient) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyDeviceClient) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyDeviceClient) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopologyDeviceClient) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopologyDeviceClient) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopologyDeviceClient) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopologyDeviceClient) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopologyDeviceClient) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopologyDeviceClient) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopologyDeviceClient) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopologyDeviceClient) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *TopologyDeviceClient) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyDeviceClient) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyDeviceClient) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyDeviceClient) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadaDeviceType

`func (o *TopologyDeviceClient) GetOmadaDeviceType() string`

GetOmadaDeviceType returns the OmadaDeviceType field if non-nil, zero value otherwise.

### GetOmadaDeviceTypeOk

`func (o *TopologyDeviceClient) GetOmadaDeviceTypeOk() (*string, bool)`

GetOmadaDeviceTypeOk returns a tuple with the OmadaDeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadaDeviceType

`func (o *TopologyDeviceClient) SetOmadaDeviceType(v string)`

SetOmadaDeviceType sets OmadaDeviceType field to given value.

### HasOmadaDeviceType

`func (o *TopologyDeviceClient) HasOmadaDeviceType() bool`

HasOmadaDeviceType returns a boolean if a field has been set.

### GetShowModel

`func (o *TopologyDeviceClient) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *TopologyDeviceClient) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *TopologyDeviceClient) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *TopologyDeviceClient) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStackGroup

`func (o *TopologyDeviceClient) GetStackGroup() bool`

GetStackGroup returns the StackGroup field if non-nil, zero value otherwise.

### GetStackGroupOk

`func (o *TopologyDeviceClient) GetStackGroupOk() (*bool, bool)`

GetStackGroupOk returns a tuple with the StackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackGroup

`func (o *TopologyDeviceClient) SetStackGroup(v bool)`

SetStackGroup sets StackGroup field to given value.

### HasStackGroup

`func (o *TopologyDeviceClient) HasStackGroup() bool`

HasStackGroup returns a boolean if a field has been set.

### GetType

`func (o *TopologyDeviceClient) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyDeviceClient) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyDeviceClient) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyDeviceClient) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUplinkMac

`func (o *TopologyDeviceClient) GetUplinkMac() string`

GetUplinkMac returns the UplinkMac field if non-nil, zero value otherwise.

### GetUplinkMacOk

`func (o *TopologyDeviceClient) GetUplinkMacOk() (*string, bool)`

GetUplinkMacOk returns a tuple with the UplinkMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkMac

`func (o *TopologyDeviceClient) SetUplinkMac(v string)`

SetUplinkMac sets UplinkMac field to given value.

### HasUplinkMac

`func (o *TopologyDeviceClient) HasUplinkMac() bool`

HasUplinkMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


