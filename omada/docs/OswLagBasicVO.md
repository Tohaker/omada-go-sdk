# OswLagBasicVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagId** | **int32** | Lag ID | 
**LagType** | Pointer to **int32** | LagType should be a value as follows: 1: Static LAG; 2: LACP; 3: LACP-Active; 4: LACP-Passive | [optional] 
**MlagEnable** | Pointer to **bool** | Indicates whether M-LAG port is enabled | [optional] 
**MlagName** | Pointer to **string** | M-LAG group Name | [optional] 
**MlagPeerSetting** | Pointer to [**OswMlagPeerSettingVO**](OswMlagPeerSettingVO.md) |  | [optional] 
**Ports** | Pointer to **[]int32** | Lag ports | [optional] 

## Methods

### NewOswLagBasicVO

`func NewOswLagBasicVO(lagId int32, ) *OswLagBasicVO`

NewOswLagBasicVO instantiates a new OswLagBasicVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswLagBasicVOWithDefaults

`func NewOswLagBasicVOWithDefaults() *OswLagBasicVO`

NewOswLagBasicVOWithDefaults instantiates a new OswLagBasicVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagId

`func (o *OswLagBasicVO) GetLagId() int32`

GetLagId returns the LagId field if non-nil, zero value otherwise.

### GetLagIdOk

`func (o *OswLagBasicVO) GetLagIdOk() (*int32, bool)`

GetLagIdOk returns a tuple with the LagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagId

`func (o *OswLagBasicVO) SetLagId(v int32)`

SetLagId sets LagId field to given value.


### GetLagType

`func (o *OswLagBasicVO) GetLagType() int32`

GetLagType returns the LagType field if non-nil, zero value otherwise.

### GetLagTypeOk

`func (o *OswLagBasicVO) GetLagTypeOk() (*int32, bool)`

GetLagTypeOk returns a tuple with the LagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagType

`func (o *OswLagBasicVO) SetLagType(v int32)`

SetLagType sets LagType field to given value.

### HasLagType

`func (o *OswLagBasicVO) HasLagType() bool`

HasLagType returns a boolean if a field has been set.

### GetMlagEnable

`func (o *OswLagBasicVO) GetMlagEnable() bool`

GetMlagEnable returns the MlagEnable field if non-nil, zero value otherwise.

### GetMlagEnableOk

`func (o *OswLagBasicVO) GetMlagEnableOk() (*bool, bool)`

GetMlagEnableOk returns a tuple with the MlagEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagEnable

`func (o *OswLagBasicVO) SetMlagEnable(v bool)`

SetMlagEnable sets MlagEnable field to given value.

### HasMlagEnable

`func (o *OswLagBasicVO) HasMlagEnable() bool`

HasMlagEnable returns a boolean if a field has been set.

### GetMlagName

`func (o *OswLagBasicVO) GetMlagName() string`

GetMlagName returns the MlagName field if non-nil, zero value otherwise.

### GetMlagNameOk

`func (o *OswLagBasicVO) GetMlagNameOk() (*string, bool)`

GetMlagNameOk returns a tuple with the MlagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagName

`func (o *OswLagBasicVO) SetMlagName(v string)`

SetMlagName sets MlagName field to given value.

### HasMlagName

`func (o *OswLagBasicVO) HasMlagName() bool`

HasMlagName returns a boolean if a field has been set.

### GetMlagPeerSetting

`func (o *OswLagBasicVO) GetMlagPeerSetting() OswMlagPeerSettingVO`

GetMlagPeerSetting returns the MlagPeerSetting field if non-nil, zero value otherwise.

### GetMlagPeerSettingOk

`func (o *OswLagBasicVO) GetMlagPeerSettingOk() (*OswMlagPeerSettingVO, bool)`

GetMlagPeerSettingOk returns a tuple with the MlagPeerSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerSetting

`func (o *OswLagBasicVO) SetMlagPeerSetting(v OswMlagPeerSettingVO)`

SetMlagPeerSetting sets MlagPeerSetting field to given value.

### HasMlagPeerSetting

`func (o *OswLagBasicVO) HasMlagPeerSetting() bool`

HasMlagPeerSetting returns a boolean if a field has been set.

### GetPorts

`func (o *OswLagBasicVO) GetPorts() []int32`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswLagBasicVO) GetPortsOk() (*[]int32, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswLagBasicVO) SetPorts(v []int32)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswLagBasicVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


