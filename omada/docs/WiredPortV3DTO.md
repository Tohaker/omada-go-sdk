# WiredPortV3DTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | Pointer to **string** | Lag Id | [optional] 
**LagPorts** | Pointer to **[]int32** | Lag Ports | [optional] 
**MlagId** | Pointer to **string** | Mlag Id | [optional] 
**MlagPorts** | Pointer to **[]int32** | Mlag Ports | [optional] 
**MultiSwitchNum** | Pointer to **int32** | Multi Switch Num | [optional] 
**MultiSwitchRole** | Pointer to **int32** | Multi Switch Role | [optional] 
**Name** | Pointer to **string** | Name of the port | [optional] 
**PoePower** | Pointer to **int64** | Poe Power | [optional] 
**PoePowerDecimal** | Pointer to **float64** | Poe Power In Decimal | [optional] 
**Port** | Pointer to **string** | Port Id | [optional] 
**StandardLagPorts** | Pointer to **[]string** | Standard Lag Ports of Stack | [optional] 
**StandardMlagPorts** | Pointer to **[]string** | Standard Mlag Ports of Stack | [optional] 
**StandardPort** | Pointer to **string** | Standard Port of Stack | [optional] 
**UplinkName** | Pointer to **string** | Uplink Name | [optional] 

## Methods

### NewWiredPortV3DTO

`func NewWiredPortV3DTO() *WiredPortV3DTO`

NewWiredPortV3DTO instantiates a new WiredPortV3DTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWiredPortV3DTOWithDefaults

`func NewWiredPortV3DTOWithDefaults() *WiredPortV3DTO`

NewWiredPortV3DTOWithDefaults instantiates a new WiredPortV3DTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *WiredPortV3DTO) GetLagId() string`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *WiredPortV3DTO) GetLagIdOk() (*string, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *WiredPortV3DTO) SetLagId(v string)`

SetLagId sets LagId field to given value.

### HasLagId

`func (o *WiredPortV3DTO) HasLagId() bool`

HasLagId returns a boolean if a field has been set.

### GetLagPorts

`func (o *WiredPortV3DTO) GetLagPorts() []int32`

GetLagPorts returns the LagPorts field if non-nil, zero value otherwise.

### GetLagPortsOk

`func (o *WiredPortV3DTO) GetLagPortsOk() (*[]int32, bool)`

GetLagPortsOk returns a tuple with the LagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPorts

`func (o *WiredPortV3DTO) SetLagPorts(v []int32)`

SetLagPorts sets LagPorts field to given value.

### HasLagPorts

`func (o *WiredPortV3DTO) HasLagPorts() bool`

HasLagPorts returns a boolean if a field has been set.

### GetMlagId

`func (o *WiredPortV3DTO) GetMlagId() string`

GetMlagId returns the MlagId field if non-nil, zero value otherwise.

### GetMlagIdOk

`func (o *WiredPortV3DTO) GetMlagIdOk() (*string, bool)`

GetMlagIdOk returns a tuple with the MlagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagId

`func (o *WiredPortV3DTO) SetMlagId(v string)`

SetMlagId sets MlagId field to given value.

### HasMlagId

`func (o *WiredPortV3DTO) HasMlagId() bool`

HasMlagId returns a boolean if a field has been set.

### GetMlagPorts

`func (o *WiredPortV3DTO) GetMlagPorts() []int32`

GetMlagPorts returns the MlagPorts field if non-nil, zero value otherwise.

### GetMlagPortsOk

`func (o *WiredPortV3DTO) GetMlagPortsOk() (*[]int32, bool)`

GetMlagPortsOk returns a tuple with the MlagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPorts

`func (o *WiredPortV3DTO) SetMlagPorts(v []int32)`

SetMlagPorts sets MlagPorts field to given value.

### HasMlagPorts

`func (o *WiredPortV3DTO) HasMlagPorts() bool`

HasMlagPorts returns a boolean if a field has been set.

### GetMultiSwitchNum

`func (o *WiredPortV3DTO) GetMultiSwitchNum() int32`

GetMultiSwitchNum returns the MultiSwitchNum field if non-nil, zero value otherwise.

### GetMultiSwitchNumOk

`func (o *WiredPortV3DTO) GetMultiSwitchNumOk() (*int32, bool)`

GetMultiSwitchNumOk returns a tuple with the MultiSwitchNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchNum

`func (o *WiredPortV3DTO) SetMultiSwitchNum(v int32)`

SetMultiSwitchNum sets MultiSwitchNum field to given value.

### HasMultiSwitchNum

`func (o *WiredPortV3DTO) HasMultiSwitchNum() bool`

HasMultiSwitchNum returns a boolean if a field has been set.

### GetMultiSwitchRole

`func (o *WiredPortV3DTO) GetMultiSwitchRole() int32`

GetMultiSwitchRole returns the MultiSwitchRole field if non-nil, zero value otherwise.

### GetMultiSwitchRoleOk

`func (o *WiredPortV3DTO) GetMultiSwitchRoleOk() (*int32, bool)`

GetMultiSwitchRoleOk returns a tuple with the MultiSwitchRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSwitchRole

`func (o *WiredPortV3DTO) SetMultiSwitchRole(v int32)`

SetMultiSwitchRole sets MultiSwitchRole field to given value.

### HasMultiSwitchRole

`func (o *WiredPortV3DTO) HasMultiSwitchRole() bool`

HasMultiSwitchRole returns a boolean if a field has been set.

### GetName

`func (o *WiredPortV3DTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WiredPortV3DTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WiredPortV3DTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WiredPortV3DTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPoePower

`func (o *WiredPortV3DTO) GetPoePower() int64`

GetPoePower returns the PoePower field if non-nil, zero value otherwise.

### GetPoePowerOk

`func (o *WiredPortV3DTO) GetPoePowerOk() (*int64, bool)`

GetPoePowerOk returns a tuple with the PoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePower

`func (o *WiredPortV3DTO) SetPoePower(v int64)`

SetPoePower sets PoePower field to given value.

### HasPoePower

`func (o *WiredPortV3DTO) HasPoePower() bool`

HasPoePower returns a boolean if a field has been set.

### GetPoePowerDecimal

`func (o *WiredPortV3DTO) GetPoePowerDecimal() float64`

GetPoePowerDecimal returns the PoePowerDecimal field if non-nil, zero value otherwise.

### GetPoePowerDecimalOk

`func (o *WiredPortV3DTO) GetPoePowerDecimalOk() (*float64, bool)`

GetPoePowerDecimalOk returns a tuple with the PoePowerDecimal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePowerDecimal

`func (o *WiredPortV3DTO) SetPoePowerDecimal(v float64)`

SetPoePowerDecimal sets PoePowerDecimal field to given value.

### HasPoePowerDecimal

`func (o *WiredPortV3DTO) HasPoePowerDecimal() bool`

HasPoePowerDecimal returns a boolean if a field has been set.

### GetPort

`func (o *WiredPortV3DTO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *WiredPortV3DTO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *WiredPortV3DTO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *WiredPortV3DTO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardLagPorts

`func (o *WiredPortV3DTO) GetStandardLagPorts() []string`

GetStandardLagPorts returns the StandardLagPorts field if non-nil, zero value otherwise.

### GetStandardLagPortsOk

`func (o *WiredPortV3DTO) GetStandardLagPortsOk() (*[]string, bool)`

GetStandardLagPortsOk returns a tuple with the StandardLagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardLagPorts

`func (o *WiredPortV3DTO) SetStandardLagPorts(v []string)`

SetStandardLagPorts sets StandardLagPorts field to given value.

### HasStandardLagPorts

`func (o *WiredPortV3DTO) HasStandardLagPorts() bool`

HasStandardLagPorts returns a boolean if a field has been set.

### GetStandardMlagPorts

`func (o *WiredPortV3DTO) GetStandardMlagPorts() []string`

GetStandardMlagPorts returns the StandardMlagPorts field if non-nil, zero value otherwise.

### GetStandardMlagPortsOk

`func (o *WiredPortV3DTO) GetStandardMlagPortsOk() (*[]string, bool)`

GetStandardMlagPortsOk returns a tuple with the StandardMlagPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardMlagPorts

`func (o *WiredPortV3DTO) SetStandardMlagPorts(v []string)`

SetStandardMlagPorts sets StandardMlagPorts field to given value.

### HasStandardMlagPorts

`func (o *WiredPortV3DTO) HasStandardMlagPorts() bool`

HasStandardMlagPorts returns a boolean if a field has been set.

### GetStandardPort

`func (o *WiredPortV3DTO) GetStandardPort() string`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *WiredPortV3DTO) GetStandardPortOk() (*string, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *WiredPortV3DTO) SetStandardPort(v string)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *WiredPortV3DTO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetUplinkName

`func (o *WiredPortV3DTO) GetUplinkName() string`

GetUplinkName returns the UplinkName field if non-nil, zero value otherwise.

### GetUplinkNameOk

`func (o *WiredPortV3DTO) GetUplinkNameOk() (*string, bool)`

GetUplinkNameOk returns a tuple with the UplinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplinkName

`func (o *WiredPortV3DTO) SetUplinkName(v string)`

SetUplinkName sets UplinkName field to given value.

### HasUplinkName

`func (o *WiredPortV3DTO) HasUplinkName() bool`

HasUplinkName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


