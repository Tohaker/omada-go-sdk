# TopologyOpenApiEdgeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedType** | Pointer to **int32** | Blocked Type | [optional] 
**BlockedVlans** | Pointer to **string** | Blocked Vlans | [optional] 
**DownLinkMac** | Pointer to **string** | DownLink Mac | [optional] 
**Port** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 
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


