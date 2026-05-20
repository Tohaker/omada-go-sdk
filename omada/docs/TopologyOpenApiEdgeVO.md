# TopologyOpenApiEdgeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to **bool** | Blocked Or Not | [optional] 
**BlockedType** | Pointer to **int32** | Blocked Type | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**DownLinkMac** | Pointer to **string** | DownLink Mac | [optional] 
**Port** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 
**RemainBlockNum** | Pointer to **int32** | Remain Block Num | [optional] 
**RemainBlockedPortList** | Pointer to [**[]TopologyOpenApiEdgeVO**](TopologyOpenApiEdgeVO.md) | Remain Blocked PortList | [optional] 
**UpLinkMac** | Pointer to **string** | UpLink Mac | [optional] 
**UpLinkPort** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 

## Methods

### NewTopologyOpenApiEdgeVO

`func NewTopologyOpenApiEdgeVO() *TopologyOpenApiEdgeVO`

NewTopologyOpenApiEdgeVO instantiates a new TopologyOpenApiEdgeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyOpenApiEdgeVOWithDefaults

`func NewTopologyOpenApiEdgeVOWithDefaults() *TopologyOpenApiEdgeVO`

NewTopologyOpenApiEdgeVOWithDefaults instantiates a new TopologyOpenApiEdgeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *TopologyOpenApiEdgeVO) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *TopologyOpenApiEdgeVO) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *TopologyOpenApiEdgeVO) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *TopologyOpenApiEdgeVO) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetBlockedType

`func (o *TopologyOpenApiEdgeVO) GetBlockedType() int32`

GetBlockedType returns the BlockedType field if non-nil, zero value otherwise.

### GetBlockedTypeOk

`func (o *TopologyOpenApiEdgeVO) GetBlockedTypeOk() (*int32, bool)`

GetBlockedTypeOk returns a tuple with the BlockedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedType

`func (o *TopologyOpenApiEdgeVO) SetBlockedType(v int32)`

SetBlockedType sets BlockedType field to given value.

### HasBlockedType

`func (o *TopologyOpenApiEdgeVO) HasBlockedType() bool`

HasBlockedType returns a boolean if a field has been set.

### GetBlockedVlans

`func (o *TopologyOpenApiEdgeVO) GetBlockedVlans() string`

GetBlockedVlans returns the BlockedVlans field if non-nil, zero value otherwise.

### GetBlockedVlansOk

`func (o *TopologyOpenApiEdgeVO) GetBlockedVlansOk() (*string, bool)`

GetBlockedVlansOk returns a tuple with the BlockedVlans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedVlans

`func (o *TopologyOpenApiEdgeVO) SetBlockedVlans(v string)`

SetBlockedVlans sets BlockedVlans field to given value.

### HasBlockedVlans

`func (o *TopologyOpenApiEdgeVO) HasBlockedVlans() bool`

HasBlockedVlans returns a boolean if a field has been set.

### GetDownLinkMac

`func (o *TopologyOpenApiEdgeVO) GetDownLinkMac() string`

GetDownLinkMac returns the DownLinkMac field if non-nil, zero value otherwise.

### GetDownLinkMacOk

`func (o *TopologyOpenApiEdgeVO) GetDownLinkMacOk() (*string, bool)`

GetDownLinkMacOk returns a tuple with the DownLinkMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLinkMac

`func (o *TopologyOpenApiEdgeVO) SetDownLinkMac(v string)`

SetDownLinkMac sets DownLinkMac field to given value.

### HasDownLinkMac

`func (o *TopologyOpenApiEdgeVO) HasDownLinkMac() bool`

HasDownLinkMac returns a boolean if a field has been set.

### GetPort

`func (o *TopologyOpenApiEdgeVO) GetPort() WiredPortV3DTO`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TopologyOpenApiEdgeVO) GetPortOk() (*WiredPortV3DTO, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TopologyOpenApiEdgeVO) SetPort(v WiredPortV3DTO)`

SetPort sets Port field to given value.

### HasPort

`func (o *TopologyOpenApiEdgeVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetRemainBlockNum

`func (o *TopologyOpenApiEdgeVO) GetRemainBlockNum() int32`

GetRemainBlockNum returns the RemainBlockNum field if non-nil, zero value otherwise.

### GetRemainBlockNumOk

`func (o *TopologyOpenApiEdgeVO) GetRemainBlockNumOk() (*int32, bool)`

GetRemainBlockNumOk returns a tuple with the RemainBlockNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainBlockNum

`func (o *TopologyOpenApiEdgeVO) SetRemainBlockNum(v int32)`

SetRemainBlockNum sets RemainBlockNum field to given value.

### HasRemainBlockNum

`func (o *TopologyOpenApiEdgeVO) HasRemainBlockNum() bool`

HasRemainBlockNum returns a boolean if a field has been set.

### GetRemainBlockedPortList

`func (o *TopologyOpenApiEdgeVO) GetRemainBlockedPortList() []TopologyOpenApiEdgeVO`

GetRemainBlockedPortList returns the RemainBlockedPortList field if non-nil, zero value otherwise.

### GetRemainBlockedPortListOk

`func (o *TopologyOpenApiEdgeVO) GetRemainBlockedPortListOk() (*[]TopologyOpenApiEdgeVO, bool)`

GetRemainBlockedPortListOk returns a tuple with the RemainBlockedPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainBlockedPortList

`func (o *TopologyOpenApiEdgeVO) SetRemainBlockedPortList(v []TopologyOpenApiEdgeVO)`

SetRemainBlockedPortList sets RemainBlockedPortList field to given value.

### HasRemainBlockedPortList

`func (o *TopologyOpenApiEdgeVO) HasRemainBlockedPortList() bool`

HasRemainBlockedPortList returns a boolean if a field has been set.

### GetUpLinkMac

`func (o *TopologyOpenApiEdgeVO) GetUpLinkMac() string`

GetUpLinkMac returns the UpLinkMac field if non-nil, zero value otherwise.

### GetUpLinkMacOk

`func (o *TopologyOpenApiEdgeVO) GetUpLinkMacOk() (*string, bool)`

GetUpLinkMacOk returns a tuple with the UpLinkMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkMac

`func (o *TopologyOpenApiEdgeVO) SetUpLinkMac(v string)`

SetUpLinkMac sets UpLinkMac field to given value.

### HasUpLinkMac

`func (o *TopologyOpenApiEdgeVO) HasUpLinkMac() bool`

HasUpLinkMac returns a boolean if a field has been set.

### GetUpLinkPort

`func (o *TopologyOpenApiEdgeVO) GetUpLinkPort() WiredPortV3DTO`

GetUpLinkPort returns the UpLinkPort field if non-nil, zero value otherwise.

### GetUpLinkPortOk

`func (o *TopologyOpenApiEdgeVO) GetUpLinkPortOk() (*WiredPortV3DTO, bool)`

GetUpLinkPortOk returns a tuple with the UpLinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLinkPort

`func (o *TopologyOpenApiEdgeVO) SetUpLinkPort(v WiredPortV3DTO)`

SetUpLinkPort sets UpLinkPort field to given value.

### HasUpLinkPort

`func (o *TopologyOpenApiEdgeVO) HasUpLinkPort() bool`

HasUpLinkPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


