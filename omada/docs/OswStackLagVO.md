# OswStackLagVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | **int32** | Lag ID | 
**LagType** | Pointer to **int32** | LagType should be a value as follows: 1: Static LAG; 2: LACP; 3: LACP-Active; 4: LACP-Passive | [optional] 
**MlagEnable** | Pointer to **bool** | Indicates whether M-LAG port is enabled | [optional] 
**MlagName** | Pointer to **string** | M-LAG group Name | [optional] 
**MlagPeerSetting** | Pointer to [**OswMlagPeerSettingVO**](OswMlagPeerSettingVO.md) |  | [optional] 
**Ports** | Pointer to **[]int32** | Lag ports | [optional] 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**StandPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Standard LAG Ports | [optional] 

## Methods

### NewOswStackLagVO

`func NewOswStackLagVO(lagId int32, ) *OswStackLagVO`

NewOswStackLagVO instantiates a new OswStackLagVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackLagVOWithDefaults

`func NewOswStackLagVOWithDefaults() *OswStackLagVO`

NewOswStackLagVOWithDefaults instantiates a new OswStackLagVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *OswStackLagVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswStackLagVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswStackLagVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.


### GetLagType

`func (o *OswStackLagVO) GetLagType() int32`

GetLagType returns the LagType field if non-nil, zero value otherwise.

### GetLagTypeOk

`func (o *OswStackLagVO) GetLagTypeOk() (*int32, bool)`

GetLagTypeOk returns a tuple with the LagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagType

`func (o *OswStackLagVO) SetLagType(v int32)`

SetLagType sets LagType field to given value.

### HasLagType

`func (o *OswStackLagVO) HasLagType() bool`

HasLagType returns a boolean if a field has been set.

### GetMlagEnable

`func (o *OswStackLagVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *OswStackLagVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *OswStackLagVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *OswStackLagVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.

### GetMlagName

`func (o *OswStackLagVO) GetMlagName() string`

GetMlagName returns the MlagName field if non-nil, zero value otherwise.

### GetMlagNameOk

`func (o *OswStackLagVO) GetMlagNameOk() (*string, bool)`

GetMlagNameOk returns a tuple with the MlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagName

`func (o *OswStackLagVO) SetMlagName(v string)`

SetMlagName sets MlagName field to given value.

### HasMlagName

`func (o *OswStackLagVO) HasMlagName() bool`

HasMlagName returns a boolean if a field has been set.

### GetMlagPeerSetting

`func (o *OswStackLagVO) GetMlagPeerSetting() OswMlagPeerSettingVO`

GetMlagPeerSetting returns the MlagPeerSetting field if non-nil, zero value otherwise.

### GetMlagPeerSettingOk

`func (o *OswStackLagVO) GetMlagPeerSettingOk() (*OswMlagPeerSettingVO, bool)`

GetMlagPeerSettingOk returns a tuple with the MlagPeerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerSetting

`func (o *OswStackLagVO) SetMlagPeerSetting(v OswMlagPeerSettingVO)`

SetMlagPeerSetting sets MlagPeerSetting field to given value.

### HasMlagPeerSetting

`func (o *OswStackLagVO) HasMlagPeerSetting() bool`

HasMlagPeerSetting returns a boolean if a field has been set.

### GetPorts

`func (o *OswStackLagVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStackLagVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStackLagVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStackLagVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetStackId

`func (o *OswStackLagVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswStackLagVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswStackLagVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswStackLagVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStandPorts

`func (o *OswStackLagVO) GetStandPorts() []OswStandPortVO`

GetStandPorts returns the StandPorts field if non-nil, zero value otherwise.

### GetStandPortsOk

`func (o *OswStackLagVO) GetStandPortsOk() (*[]OswStandPortVO, bool)`

GetStandPortsOk returns a tuple with the StandPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandPorts

`func (o *OswStackLagVO) SetStandPorts(v []OswStandPortVO)`

SetStandPorts sets StandPorts field to given value.

### HasStandPorts

`func (o *OswStackLagVO) HasStandPorts() bool`

HasStandPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


