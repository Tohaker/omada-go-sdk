# GlobalDeviceStatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | Pointer to **int32** | device activated | [optional] 
**AllAps** | Pointer to **int32** | eap health stat | [optional] 
**AllDevices** | Pointer to **int32** | the number of all devices | [optional] 
**AllOlts** | Pointer to **int32** | eap health stat | [optional] 
**AllSwitchsAndGateway** | Pointer to **int32** | the number of switchs and gateway | [optional] 
**Config** | Pointer to **int32** | device config | [optional] 
**EapHealth** | Pointer to [**HealthStatOpenApiVO**](HealthStatOpenApiVO.md) |  | [optional] 
**Expired** | Pointer to **int32** | device expired | [optional] 
**Fair** | Pointer to **int32** | device health stat is fair | [optional] 
**GatewayHealth** | Pointer to [**HealthStatOpenApiVO**](HealthStatOpenApiVO.md) |  | [optional] 
**Good** | Pointer to **int32** | device health stat is good | [optional] 
**Mesh** | Pointer to **int32** | device mesh | [optional] 
**NoData** | Pointer to **int32** | no health stat data | [optional] 
**Performance** | Pointer to **int32** | device performance | [optional] 
**Poor** | Pointer to **int32** | device health stat is poor | [optional] 
**SwitchHealth** | Pointer to [**HealthStatOpenApiVO**](HealthStatOpenApiVO.md) |  | [optional] 
**UnActivated** | Pointer to **int32** | device unActivated | [optional] 

## Methods

### NewGlobalDeviceStatOpenApiVO

`func NewGlobalDeviceStatOpenApiVO() *GlobalDeviceStatOpenApiVO`

NewGlobalDeviceStatOpenApiVO instantiates a new GlobalDeviceStatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalDeviceStatOpenApiVOWithDefaults

`func NewGlobalDeviceStatOpenApiVOWithDefaults() *GlobalDeviceStatOpenApiVO`

NewGlobalDeviceStatOpenApiVOWithDefaults instantiates a new GlobalDeviceStatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *GlobalDeviceStatOpenApiVO) GetActivated() int32`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *GlobalDeviceStatOpenApiVO) GetActivatedOk() (*int32, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *GlobalDeviceStatOpenApiVO) SetActivated(v int32)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *GlobalDeviceStatOpenApiVO) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetAllAps

`func (o *GlobalDeviceStatOpenApiVO) GetAllAps() int32`

GetAllAps returns the AllAps field if non-nil, zero value otherwise.

### GetAllApsOk

`func (o *GlobalDeviceStatOpenApiVO) GetAllApsOk() (*int32, bool)`

GetAllApsOk returns a tuple with the AllAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllAps

`func (o *GlobalDeviceStatOpenApiVO) SetAllAps(v int32)`

SetAllAps sets AllAps field to given value.

### HasAllAps

`func (o *GlobalDeviceStatOpenApiVO) HasAllAps() bool`

HasAllAps returns a boolean if a field has been set.

### GetAllDevices

`func (o *GlobalDeviceStatOpenApiVO) GetAllDevices() int32`

GetAllDevices returns the AllDevices field if non-nil, zero value otherwise.

### GetAllDevicesOk

`func (o *GlobalDeviceStatOpenApiVO) GetAllDevicesOk() (*int32, bool)`

GetAllDevicesOk returns a tuple with the AllDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllDevices

`func (o *GlobalDeviceStatOpenApiVO) SetAllDevices(v int32)`

SetAllDevices sets AllDevices field to given value.

### HasAllDevices

`func (o *GlobalDeviceStatOpenApiVO) HasAllDevices() bool`

HasAllDevices returns a boolean if a field has been set.

### GetAllOlts

`func (o *GlobalDeviceStatOpenApiVO) GetAllOlts() int32`

GetAllOlts returns the AllOlts field if non-nil, zero value otherwise.

### GetAllOltsOk

`func (o *GlobalDeviceStatOpenApiVO) GetAllOltsOk() (*int32, bool)`

GetAllOltsOk returns a tuple with the AllOlts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOlts

`func (o *GlobalDeviceStatOpenApiVO) SetAllOlts(v int32)`

SetAllOlts sets AllOlts field to given value.

### HasAllOlts

`func (o *GlobalDeviceStatOpenApiVO) HasAllOlts() bool`

HasAllOlts returns a boolean if a field has been set.

### GetAllSwitchsAndGateway

`func (o *GlobalDeviceStatOpenApiVO) GetAllSwitchsAndGateway() int32`

GetAllSwitchsAndGateway returns the AllSwitchsAndGateway field if non-nil, zero value otherwise.

### GetAllSwitchsAndGatewayOk

`func (o *GlobalDeviceStatOpenApiVO) GetAllSwitchsAndGatewayOk() (*int32, bool)`

GetAllSwitchsAndGatewayOk returns a tuple with the AllSwitchsAndGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllSwitchsAndGateway

`func (o *GlobalDeviceStatOpenApiVO) SetAllSwitchsAndGateway(v int32)`

SetAllSwitchsAndGateway sets AllSwitchsAndGateway field to given value.

### HasAllSwitchsAndGateway

`func (o *GlobalDeviceStatOpenApiVO) HasAllSwitchsAndGateway() bool`

HasAllSwitchsAndGateway returns a boolean if a field has been set.

### GetConfig

`func (o *GlobalDeviceStatOpenApiVO) GetConfig() int32`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GlobalDeviceStatOpenApiVO) GetConfigOk() (*int32, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GlobalDeviceStatOpenApiVO) SetConfig(v int32)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GlobalDeviceStatOpenApiVO) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetEapHealth

`func (o *GlobalDeviceStatOpenApiVO) GetEapHealth() HealthStatOpenApiVO`

GetEapHealth returns the EapHealth field if non-nil, zero value otherwise.

### GetEapHealthOk

`func (o *GlobalDeviceStatOpenApiVO) GetEapHealthOk() (*HealthStatOpenApiVO, bool)`

GetEapHealthOk returns a tuple with the EapHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEapHealth

`func (o *GlobalDeviceStatOpenApiVO) SetEapHealth(v HealthStatOpenApiVO)`

SetEapHealth sets EapHealth field to given value.

### HasEapHealth

`func (o *GlobalDeviceStatOpenApiVO) HasEapHealth() bool`

HasEapHealth returns a boolean if a field has been set.

### GetExpired

`func (o *GlobalDeviceStatOpenApiVO) GetExpired() int32`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *GlobalDeviceStatOpenApiVO) GetExpiredOk() (*int32, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *GlobalDeviceStatOpenApiVO) SetExpired(v int32)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *GlobalDeviceStatOpenApiVO) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetFair

`func (o *GlobalDeviceStatOpenApiVO) GetFair() int32`

GetFair returns the Fair field if non-nil, zero value otherwise.

### GetFairOk

`func (o *GlobalDeviceStatOpenApiVO) GetFairOk() (*int32, bool)`

GetFairOk returns a tuple with the Fair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFair

`func (o *GlobalDeviceStatOpenApiVO) SetFair(v int32)`

SetFair sets Fair field to given value.

### HasFair

`func (o *GlobalDeviceStatOpenApiVO) HasFair() bool`

HasFair returns a boolean if a field has been set.

### GetGatewayHealth

`func (o *GlobalDeviceStatOpenApiVO) GetGatewayHealth() HealthStatOpenApiVO`

GetGatewayHealth returns the GatewayHealth field if non-nil, zero value otherwise.

### GetGatewayHealthOk

`func (o *GlobalDeviceStatOpenApiVO) GetGatewayHealthOk() (*HealthStatOpenApiVO, bool)`

GetGatewayHealthOk returns a tuple with the GatewayHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewayHealth

`func (o *GlobalDeviceStatOpenApiVO) SetGatewayHealth(v HealthStatOpenApiVO)`

SetGatewayHealth sets GatewayHealth field to given value.

### HasGatewayHealth

`func (o *GlobalDeviceStatOpenApiVO) HasGatewayHealth() bool`

HasGatewayHealth returns a boolean if a field has been set.

### GetGood

`func (o *GlobalDeviceStatOpenApiVO) GetGood() int32`

GetGood returns the Good field if non-nil, zero value otherwise.

### GetGoodOk

`func (o *GlobalDeviceStatOpenApiVO) GetGoodOk() (*int32, bool)`

GetGoodOk returns a tuple with the Good field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGood

`func (o *GlobalDeviceStatOpenApiVO) SetGood(v int32)`

SetGood sets Good field to given value.

### HasGood

`func (o *GlobalDeviceStatOpenApiVO) HasGood() bool`

HasGood returns a boolean if a field has been set.

### GetMesh

`func (o *GlobalDeviceStatOpenApiVO) GetMesh() int32`

GetMesh returns the Mesh field if non-nil, zero value otherwise.

### GetMeshOk

`func (o *GlobalDeviceStatOpenApiVO) GetMeshOk() (*int32, bool)`

GetMeshOk returns a tuple with the Mesh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMesh

`func (o *GlobalDeviceStatOpenApiVO) SetMesh(v int32)`

SetMesh sets Mesh field to given value.

### HasMesh

`func (o *GlobalDeviceStatOpenApiVO) HasMesh() bool`

HasMesh returns a boolean if a field has been set.

### GetNoData

`func (o *GlobalDeviceStatOpenApiVO) GetNoData() int32`

GetNoData returns the NoData field if non-nil, zero value otherwise.

### GetNoDataOk

`func (o *GlobalDeviceStatOpenApiVO) GetNoDataOk() (*int32, bool)`

GetNoDataOk returns a tuple with the NoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoData

`func (o *GlobalDeviceStatOpenApiVO) SetNoData(v int32)`

SetNoData sets NoData field to given value.

### HasNoData

`func (o *GlobalDeviceStatOpenApiVO) HasNoData() bool`

HasNoData returns a boolean if a field has been set.

### GetPerformance

`func (o *GlobalDeviceStatOpenApiVO) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *GlobalDeviceStatOpenApiVO) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *GlobalDeviceStatOpenApiVO) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *GlobalDeviceStatOpenApiVO) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.

### GetPoor

`func (o *GlobalDeviceStatOpenApiVO) GetPoor() int32`

GetPoor returns the Poor field if non-nil, zero value otherwise.

### GetPoorOk

`func (o *GlobalDeviceStatOpenApiVO) GetPoorOk() (*int32, bool)`

GetPoorOk returns a tuple with the Poor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoor

`func (o *GlobalDeviceStatOpenApiVO) SetPoor(v int32)`

SetPoor sets Poor field to given value.

### HasPoor

`func (o *GlobalDeviceStatOpenApiVO) HasPoor() bool`

HasPoor returns a boolean if a field has been set.

### GetSwitchHealth

`func (o *GlobalDeviceStatOpenApiVO) GetSwitchHealth() HealthStatOpenApiVO`

GetSwitchHealth returns the SwitchHealth field if non-nil, zero value otherwise.

### GetSwitchHealthOk

`func (o *GlobalDeviceStatOpenApiVO) GetSwitchHealthOk() (*HealthStatOpenApiVO, bool)`

GetSwitchHealthOk returns a tuple with the SwitchHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchHealth

`func (o *GlobalDeviceStatOpenApiVO) SetSwitchHealth(v HealthStatOpenApiVO)`

SetSwitchHealth sets SwitchHealth field to given value.

### HasSwitchHealth

`func (o *GlobalDeviceStatOpenApiVO) HasSwitchHealth() bool`

HasSwitchHealth returns a boolean if a field has been set.

### GetUnActivated

`func (o *GlobalDeviceStatOpenApiVO) GetUnActivated() int32`

GetUnActivated returns the UnActivated field if non-nil, zero value otherwise.

### GetUnActivatedOk

`func (o *GlobalDeviceStatOpenApiVO) GetUnActivatedOk() (*int32, bool)`

GetUnActivatedOk returns a tuple with the UnActivated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnActivated

`func (o *GlobalDeviceStatOpenApiVO) SetUnActivated(v int32)`

SetUnActivated sets UnActivated field to given value.

### HasUnActivated

`func (o *GlobalDeviceStatOpenApiVO) HasUnActivated() bool`

HasUnActivated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


