# WiredPortDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **int32** | Lag Id | [optional] 
**LagPorts** | Pointer to **[]int32** | Lag Ports | [optional] 
**MultiSwitchNum** | Pointer to **int32** | Multi Switch Num | [optional] 
**MultiSwitchRole** | Pointer to **int32** | Multi Switch Role | [optional] 
**Name** | Pointer to **string** | Name of the port | [optional] 
**Port** | Pointer to **int32** | Port Id | [optional] 
**StandardLagPorts** | Pointer to **[]string** | Standard Lag Ports of Stack | [optional] 
**StandardPort** | Pointer to **string** | Standard Port of Stack | [optional] 
**UplinkName** | Pointer to **string** | Uplink Name | [optional] 

## Methods

### NewWiredPortDTO

`func NewWiredPortDTO() *WiredPortDTO`

NewWiredPortDTO instantiates a new WiredPortDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredPortDTOWithDefaults

`func NewWiredPortDTOWithDefaults() *WiredPortDTO`

NewWiredPortDTOWithDefaults instantiates a new WiredPortDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *WiredPortDTO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *WiredPortDTO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *WiredPortDTO) SetLagId(v int32)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *WiredPortDTO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagPorts

`func (o *WiredPortDTO) GetLagPorts() []int32`

GetLagPorts returns the LagPorts field if non-nil, zero value otherwise.

### GetLagPortsOk

`func (o *WiredPortDTO) GetLagPortsOk() (*[]int32, bool)`

GetLagPortsOk returns a tuple with the LagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPorts

`func (o *WiredPortDTO) SetLagPorts(v []int32)`

SetLagPorts sets LagPorts field to given value.

### HasLagPorts

`func (o *WiredPortDTO) HasLagPorts() bool`

HasLagPorts returns a boolean if a field has been set.

### GetMultiSwitchNum

`func (o *WiredPortDTO) GetMultiSwitchNum() int32`

GetMultiSwitchNum returns the MultiSwitchNum field if non-nil, zero value otherwise.

### GetMultiSwitchNumOk

`func (o *WiredPortDTO) GetMultiSwitchNumOk() (*int32, bool)`

GetMultiSwitchNumOk returns a tuple with the MultiSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchNum

`func (o *WiredPortDTO) SetMultiSwitchNum(v int32)`

SetMultiSwitchNum sets MultiSwitchNum field to given value.

### HasMultiSwitchNum

`func (o *WiredPortDTO) HasMultiSwitchNum() bool`

HasMultiSwitchNum returns a boolean if a field has been set.

### GetMultiSwitchRole

`func (o *WiredPortDTO) GetMultiSwitchRole() int32`

GetMultiSwitchRole returns the MultiSwitchRole field if non-nil, zero value otherwise.

### GetMultiSwitchRoleOk

`func (o *WiredPortDTO) GetMultiSwitchRoleOk() (*int32, bool)`

GetMultiSwitchRoleOk returns a tuple with the MultiSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchRole

`func (o *WiredPortDTO) SetMultiSwitchRole(v int32)`

SetMultiSwitchRole sets MultiSwitchRole field to given value.

### HasMultiSwitchRole

`func (o *WiredPortDTO) HasMultiSwitchRole() bool`

HasMultiSwitchRole returns a boolean if a field has been set.

### GetName

`func (o *WiredPortDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WiredPortDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WiredPortDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WiredPortDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *WiredPortDTO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WiredPortDTO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WiredPortDTO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *WiredPortDTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardLagPorts

`func (o *WiredPortDTO) GetStandardLagPorts() []string`

GetStandardLagPorts returns the StandardLagPorts field if non-nil, zero value otherwise.

### GetStandardLagPortsOk

`func (o *WiredPortDTO) GetStandardLagPortsOk() (*[]string, bool)`

GetStandardLagPortsOk returns a tuple with the StandardLagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLagPorts

`func (o *WiredPortDTO) SetStandardLagPorts(v []string)`

SetStandardLagPorts sets StandardLagPorts field to given value.

### HasStandardLagPorts

`func (o *WiredPortDTO) HasStandardLagPorts() bool`

HasStandardLagPorts returns a boolean if a field has been set.

### GetStandardPort

`func (o *WiredPortDTO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *WiredPortDTO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *WiredPortDTO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *WiredPortDTO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetUplinkName

`func (o *WiredPortDTO) GetUplinkName() string`

GetUplinkName returns the UplinkName field if non-nil, zero value otherwise.

### GetUplinkNameOk

`func (o *WiredPortDTO) GetUplinkNameOk() (*string, bool)`

GetUplinkNameOk returns a tuple with the UplinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkName

`func (o *WiredPortDTO) SetUplinkName(v string)`

SetUplinkName sets UplinkName field to given value.

### HasUplinkName

`func (o *WiredPortDTO) HasUplinkName() bool`

HasUplinkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


