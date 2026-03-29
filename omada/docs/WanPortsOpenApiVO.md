# WanPortsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | whether to enable wan Settings Override | 
**ForceChange** | Pointer to **bool** | Force change wan ports. If set true, wan ports will be modified without check, and the gateway will reboot. If set false, consequence will be checked and return check result. | [optional] 
**PortUuids** | Pointer to **[]string** | select wan ports to open, and other wan ports will be closed, you can get port uuid list from &#39;Get internet basic info&#39;. | [optional] 
**PreOsgModel** | Pointer to **int32** | select gateway model, you can query option from &#39;Get supported gateway model list for pre-configuration&#39;. If already choose or adopted gateway,use current gateway model id. | [optional] 
**WanPortNum** | Pointer to **int32** | custom number of wan ports, only for preOsgModel at 2 | [optional] 

## Methods

### NewWanPortsOpenApiVO

`func NewWanPortsOpenApiVO(enable bool, ) *WanPortsOpenApiVO`

NewWanPortsOpenApiVO instantiates a new WanPortsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortsOpenApiVOWithDefaults

`func NewWanPortsOpenApiVOWithDefaults() *WanPortsOpenApiVO`

NewWanPortsOpenApiVOWithDefaults instantiates a new WanPortsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *WanPortsOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *WanPortsOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *WanPortsOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetForceChange

`func (o *WanPortsOpenApiVO) GetForceChange() bool`

GetForceChange returns the ForceChange field if non-nil, zero value otherwise.

### GetForceChangeOk

`func (o *WanPortsOpenApiVO) GetForceChangeOk() (*bool, bool)`

GetForceChangeOk returns a tuple with the ForceChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceChange

`func (o *WanPortsOpenApiVO) SetForceChange(v bool)`

SetForceChange sets ForceChange field to given value.

### HasForceChange

`func (o *WanPortsOpenApiVO) HasForceChange() bool`

HasForceChange returns a boolean if a field has been set.

### GetPortUuids

`func (o *WanPortsOpenApiVO) GetPortUuids() []string`

GetPortUuids returns the PortUuids field if non-nil, zero value otherwise.

### GetPortUuidsOk

`func (o *WanPortsOpenApiVO) GetPortUuidsOk() (*[]string, bool)`

GetPortUuidsOk returns a tuple with the PortUuids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuids

`func (o *WanPortsOpenApiVO) SetPortUuids(v []string)`

SetPortUuids sets PortUuids field to given value.

### HasPortUuids

`func (o *WanPortsOpenApiVO) HasPortUuids() bool`

HasPortUuids returns a boolean if a field has been set.

### GetPreOsgModel

`func (o *WanPortsOpenApiVO) GetPreOsgModel() int32`

GetPreOsgModel returns the PreOsgModel field if non-nil, zero value otherwise.

### GetPreOsgModelOk

`func (o *WanPortsOpenApiVO) GetPreOsgModelOk() (*int32, bool)`

GetPreOsgModelOk returns a tuple with the PreOsgModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreOsgModel

`func (o *WanPortsOpenApiVO) SetPreOsgModel(v int32)`

SetPreOsgModel sets PreOsgModel field to given value.

### HasPreOsgModel

`func (o *WanPortsOpenApiVO) HasPreOsgModel() bool`

HasPreOsgModel returns a boolean if a field has been set.

### GetWanPortNum

`func (o *WanPortsOpenApiVO) GetWanPortNum() int32`

GetWanPortNum returns the WanPortNum field if non-nil, zero value otherwise.

### GetWanPortNumOk

`func (o *WanPortsOpenApiVO) GetWanPortNumOk() (*int32, bool)`

GetWanPortNumOk returns a tuple with the WanPortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortNum

`func (o *WanPortsOpenApiVO) SetWanPortNum(v int32)`

SetWanPortNum sets WanPortNum field to given value.

### HasWanPortNum

`func (o *WanPortsOpenApiVO) HasWanPortNum() bool`

HasWanPortNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


