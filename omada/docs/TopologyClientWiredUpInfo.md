# TopologyClientWiredUpInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duplex** | Pointer to **int32** | Duplex, it should be a value as follows: 1:Half Duplex. | [optional] 
**LinkSpeed** | Pointer to **int32** | Link Speed, it should be a value as follows: 1:10Mbps, 2:100Mbps, 3:1000Mbps, 4:2.5Gbps, 5:10Gbps, 6:5Gbps, 7:25Gbps. | [optional] 
**Port** | Pointer to [**TopologyClientUplinkPort**](TopologyClientUplinkPort.md) |  | [optional] 
**UplinkPort** | Pointer to [**WiredPortV3DTO**](WiredPortV3DTO.md) |  | [optional] 

## Methods

### NewTopologyClientWiredUpInfo

`func NewTopologyClientWiredUpInfo() *TopologyClientWiredUpInfo`

NewTopologyClientWiredUpInfo instantiates a new TopologyClientWiredUpInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyClientWiredUpInfoWithDefaults

`func NewTopologyClientWiredUpInfoWithDefaults() *TopologyClientWiredUpInfo`

NewTopologyClientWiredUpInfoWithDefaults instantiates a new TopologyClientWiredUpInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuplex

`func (o *TopologyClientWiredUpInfo) GetDuplex() int32`

GetDuplex returns the Duplex field if non-nil, zero value otherwise.

### GetDuplexOk

`func (o *TopologyClientWiredUpInfo) GetDuplexOk() (*int32, bool)`

GetDuplexOk returns a tuple with the Duplex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplex

`func (o *TopologyClientWiredUpInfo) SetDuplex(v int32)`

SetDuplex sets Duplex field to given value.

### HasDuplex

`func (o *TopologyClientWiredUpInfo) HasDuplex() bool`

HasDuplex returns a boolean if a field has been set.

### GetLinkSpeed

`func (o *TopologyClientWiredUpInfo) GetLinkSpeed() int32`

GetLinkSpeed returns the LinkSpeed field if non-nil, zero value otherwise.

### GetLinkSpeedOk

`func (o *TopologyClientWiredUpInfo) GetLinkSpeedOk() (*int32, bool)`

GetLinkSpeedOk returns a tuple with the LinkSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkSpeed

`func (o *TopologyClientWiredUpInfo) SetLinkSpeed(v int32)`

SetLinkSpeed sets LinkSpeed field to given value.

### HasLinkSpeed

`func (o *TopologyClientWiredUpInfo) HasLinkSpeed() bool`

HasLinkSpeed returns a boolean if a field has been set.

### GetPort

`func (o *TopologyClientWiredUpInfo) GetPort() TopologyClientUplinkPort`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *TopologyClientWiredUpInfo) GetPortOk() (*TopologyClientUplinkPort, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *TopologyClientWiredUpInfo) SetPort(v TopologyClientUplinkPort)`

SetPort sets Port field to given value.

### HasPort

`func (o *TopologyClientWiredUpInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUplinkPort

`func (o *TopologyClientWiredUpInfo) GetUplinkPort() WiredPortV3DTO`

GetUplinkPort returns the UplinkPort field if non-nil, zero value otherwise.

### GetUplinkPortOk

`func (o *TopologyClientWiredUpInfo) GetUplinkPortOk() (*WiredPortV3DTO, bool)`

GetUplinkPortOk returns a tuple with the UplinkPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkPort

`func (o *TopologyClientWiredUpInfo) SetUplinkPort(v WiredPortV3DTO)`

SetUplinkPort sets UplinkPort field to given value.

### HasUplinkPort

`func (o *TopologyClientWiredUpInfo) HasUplinkPort() bool`

HasUplinkPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


