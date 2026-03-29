# TopologyOpenApiNodeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel2g** | Pointer to **int32** | channel2g | [optional] 
**Channel5g** | Pointer to **int32** | channel5g | [optional] 
**Channel5g2** | Pointer to **int32** | channel5g2 | [optional] 
**Channel6g** | Pointer to **int32** | channel6g | [optional] 
**ClientCount** | Pointer to **int32** | Client Count | [optional] 
**ClientHealth** | Pointer to [**ClientAggHealthDTO**](ClientAggHealthDTO.md) |  | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device Series Type | [optional] 
**HealthScore** | Pointer to **int32** | Health Score | [optional] 
**LanPorts** | Pointer to [**[]LanPort**](LanPort.md) | Lan Port List | [optional] 
**Mac** | Pointer to **string** | Device Mac | [optional] 
**Model** | Pointer to **string** | Device Model | [optional] 
**ModelVersion** | Pointer to **string** | Device ModelVersion | [optional] 
**Name** | Pointer to **string** | Device Name | [optional] 
**RdMode2g** | Pointer to **string** | rdMode2g | [optional] 
**RdMode5g** | Pointer to **string** | rdMode5g | [optional] 
**RdMode5g2** | Pointer to **string** | rdMode5g2 | [optional] 
**RdMode6g** | Pointer to **string** | rdMode6g | [optional] 
**StackGroup** | Pointer to **bool** | Stack Group | [optional] 
**StackId** | Pointer to **string** | Stack Id | [optional] 
**Support5g2** | Pointer to **bool** | support5g2 | [optional] 
**Type** | Pointer to **string** | Device Type | [optional] 
**UpperPort** | Pointer to [**OswUpInfo**](OswUpInfo.md) |  | [optional] 
**WanPorts** | Pointer to [**[]WanPort**](WanPort.md) | Wan Port List | [optional] 
**WiredUpLink** | Pointer to [**WireUpLink**](WireUpLink.md) |  | [optional] 
**WirelessUpLink** | Pointer to [**WirelessUpLink**](WirelessUpLink.md) |  | [optional] 

## Methods

### NewTopologyOpenApiNodeVO

`func NewTopologyOpenApiNodeVO() *TopologyOpenApiNodeVO`

NewTopologyOpenApiNodeVO instantiates a new TopologyOpenApiNodeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopologyOpenApiNodeVOWithDefaults

`func NewTopologyOpenApiNodeVOWithDefaults() *TopologyOpenApiNodeVO`

NewTopologyOpenApiNodeVOWithDefaults instantiates a new TopologyOpenApiNodeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel2g

`func (o *TopologyOpenApiNodeVO) GetChannel2g() int32`

GetChannel2g returns the Channel2g field if non-nil, zero value otherwise.

### GetChannel2gOk

`func (o *TopologyOpenApiNodeVO) GetChannel2gOk() (*int32, bool)`

GetChannel2gOk returns a tuple with the Channel2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel2g

`func (o *TopologyOpenApiNodeVO) SetChannel2g(v int32)`

SetChannel2g sets Channel2g field to given value.

### HasChannel2g

`func (o *TopologyOpenApiNodeVO) HasChannel2g() bool`

HasChannel2g returns a boolean if a field has been set.

### GetChannel5g

`func (o *TopologyOpenApiNodeVO) GetChannel5g() int32`

GetChannel5g returns the Channel5g field if non-nil, zero value otherwise.

### GetChannel5gOk

`func (o *TopologyOpenApiNodeVO) GetChannel5gOk() (*int32, bool)`

GetChannel5gOk returns a tuple with the Channel5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g

`func (o *TopologyOpenApiNodeVO) SetChannel5g(v int32)`

SetChannel5g sets Channel5g field to given value.

### HasChannel5g

`func (o *TopologyOpenApiNodeVO) HasChannel5g() bool`

HasChannel5g returns a boolean if a field has been set.

### GetChannel5g2

`func (o *TopologyOpenApiNodeVO) GetChannel5g2() int32`

GetChannel5g2 returns the Channel5g2 field if non-nil, zero value otherwise.

### GetChannel5g2Ok

`func (o *TopologyOpenApiNodeVO) GetChannel5g2Ok() (*int32, bool)`

GetChannel5g2Ok returns a tuple with the Channel5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g2

`func (o *TopologyOpenApiNodeVO) SetChannel5g2(v int32)`

SetChannel5g2 sets Channel5g2 field to given value.

### HasChannel5g2

`func (o *TopologyOpenApiNodeVO) HasChannel5g2() bool`

HasChannel5g2 returns a boolean if a field has been set.

### GetChannel6g

`func (o *TopologyOpenApiNodeVO) GetChannel6g() int32`

GetChannel6g returns the Channel6g field if non-nil, zero value otherwise.

### GetChannel6gOk

`func (o *TopologyOpenApiNodeVO) GetChannel6gOk() (*int32, bool)`

GetChannel6gOk returns a tuple with the Channel6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel6g

`func (o *TopologyOpenApiNodeVO) SetChannel6g(v int32)`

SetChannel6g sets Channel6g field to given value.

### HasChannel6g

`func (o *TopologyOpenApiNodeVO) HasChannel6g() bool`

HasChannel6g returns a boolean if a field has been set.

### GetClientCount

`func (o *TopologyOpenApiNodeVO) GetClientCount() int32`

GetClientCount returns the ClientCount field if non-nil, zero value otherwise.

### GetClientCountOk

`func (o *TopologyOpenApiNodeVO) GetClientCountOk() (*int32, bool)`

GetClientCountOk returns a tuple with the ClientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCount

`func (o *TopologyOpenApiNodeVO) SetClientCount(v int32)`

SetClientCount sets ClientCount field to given value.

### HasClientCount

`func (o *TopologyOpenApiNodeVO) HasClientCount() bool`

HasClientCount returns a boolean if a field has been set.

### GetClientHealth

`func (o *TopologyOpenApiNodeVO) GetClientHealth() ClientAggHealthDTO`

GetClientHealth returns the ClientHealth field if non-nil, zero value otherwise.

### GetClientHealthOk

`func (o *TopologyOpenApiNodeVO) GetClientHealthOk() (*ClientAggHealthDTO, bool)`

GetClientHealthOk returns a tuple with the ClientHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealth

`func (o *TopologyOpenApiNodeVO) SetClientHealth(v ClientAggHealthDTO)`

SetClientHealth sets ClientHealth field to given value.

### HasClientHealth

`func (o *TopologyOpenApiNodeVO) HasClientHealth() bool`

HasClientHealth returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *TopologyOpenApiNodeVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *TopologyOpenApiNodeVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *TopologyOpenApiNodeVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *TopologyOpenApiNodeVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetHealthScore

`func (o *TopologyOpenApiNodeVO) GetHealthScore() int32`

GetHealthScore returns the HealthScore field if non-nil, zero value otherwise.

### GetHealthScoreOk

`func (o *TopologyOpenApiNodeVO) GetHealthScoreOk() (*int32, bool)`

GetHealthScoreOk returns a tuple with the HealthScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthScore

`func (o *TopologyOpenApiNodeVO) SetHealthScore(v int32)`

SetHealthScore sets HealthScore field to given value.

### HasHealthScore

`func (o *TopologyOpenApiNodeVO) HasHealthScore() bool`

HasHealthScore returns a boolean if a field has been set.

### GetLanPorts

`func (o *TopologyOpenApiNodeVO) GetLanPorts() []LanPort`

GetLanPorts returns the LanPorts field if non-nil, zero value otherwise.

### GetLanPortsOk

`func (o *TopologyOpenApiNodeVO) GetLanPortsOk() (*[]LanPort, bool)`

GetLanPortsOk returns a tuple with the LanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanPorts

`func (o *TopologyOpenApiNodeVO) SetLanPorts(v []LanPort)`

SetLanPorts sets LanPorts field to given value.

### HasLanPorts

`func (o *TopologyOpenApiNodeVO) HasLanPorts() bool`

HasLanPorts returns a boolean if a field has been set.

### GetMac

`func (o *TopologyOpenApiNodeVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *TopologyOpenApiNodeVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *TopologyOpenApiNodeVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *TopologyOpenApiNodeVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *TopologyOpenApiNodeVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *TopologyOpenApiNodeVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *TopologyOpenApiNodeVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *TopologyOpenApiNodeVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *TopologyOpenApiNodeVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *TopologyOpenApiNodeVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *TopologyOpenApiNodeVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *TopologyOpenApiNodeVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *TopologyOpenApiNodeVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TopologyOpenApiNodeVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TopologyOpenApiNodeVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TopologyOpenApiNodeVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRdMode2g

`func (o *TopologyOpenApiNodeVO) GetRdMode2g() string`

GetRdMode2g returns the RdMode2g field if non-nil, zero value otherwise.

### GetRdMode2gOk

`func (o *TopologyOpenApiNodeVO) GetRdMode2gOk() (*string, bool)`

GetRdMode2gOk returns a tuple with the RdMode2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode2g

`func (o *TopologyOpenApiNodeVO) SetRdMode2g(v string)`

SetRdMode2g sets RdMode2g field to given value.

### HasRdMode2g

`func (o *TopologyOpenApiNodeVO) HasRdMode2g() bool`

HasRdMode2g returns a boolean if a field has been set.

### GetRdMode5g

`func (o *TopologyOpenApiNodeVO) GetRdMode5g() string`

GetRdMode5g returns the RdMode5g field if non-nil, zero value otherwise.

### GetRdMode5gOk

`func (o *TopologyOpenApiNodeVO) GetRdMode5gOk() (*string, bool)`

GetRdMode5gOk returns a tuple with the RdMode5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode5g

`func (o *TopologyOpenApiNodeVO) SetRdMode5g(v string)`

SetRdMode5g sets RdMode5g field to given value.

### HasRdMode5g

`func (o *TopologyOpenApiNodeVO) HasRdMode5g() bool`

HasRdMode5g returns a boolean if a field has been set.

### GetRdMode5g2

`func (o *TopologyOpenApiNodeVO) GetRdMode5g2() string`

GetRdMode5g2 returns the RdMode5g2 field if non-nil, zero value otherwise.

### GetRdMode5g2Ok

`func (o *TopologyOpenApiNodeVO) GetRdMode5g2Ok() (*string, bool)`

GetRdMode5g2Ok returns a tuple with the RdMode5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode5g2

`func (o *TopologyOpenApiNodeVO) SetRdMode5g2(v string)`

SetRdMode5g2 sets RdMode5g2 field to given value.

### HasRdMode5g2

`func (o *TopologyOpenApiNodeVO) HasRdMode5g2() bool`

HasRdMode5g2 returns a boolean if a field has been set.

### GetRdMode6g

`func (o *TopologyOpenApiNodeVO) GetRdMode6g() string`

GetRdMode6g returns the RdMode6g field if non-nil, zero value otherwise.

### GetRdMode6gOk

`func (o *TopologyOpenApiNodeVO) GetRdMode6gOk() (*string, bool)`

GetRdMode6gOk returns a tuple with the RdMode6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRdMode6g

`func (o *TopologyOpenApiNodeVO) SetRdMode6g(v string)`

SetRdMode6g sets RdMode6g field to given value.

### HasRdMode6g

`func (o *TopologyOpenApiNodeVO) HasRdMode6g() bool`

HasRdMode6g returns a boolean if a field has been set.

### GetStackGroup

`func (o *TopologyOpenApiNodeVO) GetStackGroup() bool`

GetStackGroup returns the StackGroup field if non-nil, zero value otherwise.

### GetStackGroupOk

`func (o *TopologyOpenApiNodeVO) GetStackGroupOk() (*bool, bool)`

GetStackGroupOk returns a tuple with the StackGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackGroup

`func (o *TopologyOpenApiNodeVO) SetStackGroup(v bool)`

SetStackGroup sets StackGroup field to given value.

### HasStackGroup

`func (o *TopologyOpenApiNodeVO) HasStackGroup() bool`

HasStackGroup returns a boolean if a field has been set.

### GetStackId

`func (o *TopologyOpenApiNodeVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *TopologyOpenApiNodeVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *TopologyOpenApiNodeVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *TopologyOpenApiNodeVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetSupport5g2

`func (o *TopologyOpenApiNodeVO) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *TopologyOpenApiNodeVO) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *TopologyOpenApiNodeVO) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *TopologyOpenApiNodeVO) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetType

`func (o *TopologyOpenApiNodeVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TopologyOpenApiNodeVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TopologyOpenApiNodeVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TopologyOpenApiNodeVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpperPort

`func (o *TopologyOpenApiNodeVO) GetUpperPort() OswUpInfo`

GetUpperPort returns the UpperPort field if non-nil, zero value otherwise.

### GetUpperPortOk

`func (o *TopologyOpenApiNodeVO) GetUpperPortOk() (*OswUpInfo, bool)`

GetUpperPortOk returns a tuple with the UpperPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpperPort

`func (o *TopologyOpenApiNodeVO) SetUpperPort(v OswUpInfo)`

SetUpperPort sets UpperPort field to given value.

### HasUpperPort

`func (o *TopologyOpenApiNodeVO) HasUpperPort() bool`

HasUpperPort returns a boolean if a field has been set.

### GetWanPorts

`func (o *TopologyOpenApiNodeVO) GetWanPorts() []WanPort`

GetWanPorts returns the WanPorts field if non-nil, zero value otherwise.

### GetWanPortsOk

`func (o *TopologyOpenApiNodeVO) GetWanPortsOk() (*[]WanPort, bool)`

GetWanPortsOk returns a tuple with the WanPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPorts

`func (o *TopologyOpenApiNodeVO) SetWanPorts(v []WanPort)`

SetWanPorts sets WanPorts field to given value.

### HasWanPorts

`func (o *TopologyOpenApiNodeVO) HasWanPorts() bool`

HasWanPorts returns a boolean if a field has been set.

### GetWiredUpLink

`func (o *TopologyOpenApiNodeVO) GetWiredUpLink() WireUpLink`

GetWiredUpLink returns the WiredUpLink field if non-nil, zero value otherwise.

### GetWiredUpLinkOk

`func (o *TopologyOpenApiNodeVO) GetWiredUpLinkOk() (*WireUpLink, bool)`

GetWiredUpLinkOk returns a tuple with the WiredUpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWiredUpLink

`func (o *TopologyOpenApiNodeVO) SetWiredUpLink(v WireUpLink)`

SetWiredUpLink sets WiredUpLink field to given value.

### HasWiredUpLink

`func (o *TopologyOpenApiNodeVO) HasWiredUpLink() bool`

HasWiredUpLink returns a boolean if a field has been set.

### GetWirelessUpLink

`func (o *TopologyOpenApiNodeVO) GetWirelessUpLink() WirelessUpLink`

GetWirelessUpLink returns the WirelessUpLink field if non-nil, zero value otherwise.

### GetWirelessUpLinkOk

`func (o *TopologyOpenApiNodeVO) GetWirelessUpLinkOk() (*WirelessUpLink, bool)`

GetWirelessUpLinkOk returns a tuple with the WirelessUpLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUpLink

`func (o *TopologyOpenApiNodeVO) SetWirelessUpLink(v WirelessUpLink)`

SetWirelessUpLink sets WirelessUpLink field to given value.

### HasWirelessUpLink

`func (o *TopologyOpenApiNodeVO) HasWirelessUpLink() bool`

HasWirelessUpLink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


