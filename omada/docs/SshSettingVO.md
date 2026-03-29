# SshSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layer3Access** | Pointer to **bool** |  | [optional] 
**Resource** | Pointer to **int32** |  | [optional] 
**SshEnable** | **bool** |  | 
**SshServerPort** | Pointer to **int32** |  | [optional] 

## Methods

### NewSshSettingVO

`func NewSshSettingVO(sshEnable bool, ) *SshSettingVO`

NewSshSettingVO instantiates a new SshSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSshSettingVOWithDefaults

`func NewSshSettingVOWithDefaults() *SshSettingVO`

NewSshSettingVOWithDefaults instantiates a new SshSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayer3Access

`func (o *SshSettingVO) GetLayer3Access() bool`

GetLayer3Access returns the Layer3Access field if non-nil, zero value otherwise.

### GetLayer3AccessOk

`func (o *SshSettingVO) GetLayer3AccessOk() (*bool, bool)`

GetLayer3AccessOk returns a tuple with the Layer3Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer3Access

`func (o *SshSettingVO) SetLayer3Access(v bool)`

SetLayer3Access sets Layer3Access field to given value.

### HasLayer3Access

`func (o *SshSettingVO) HasLayer3Access() bool`

HasLayer3Access returns a boolean if a field has been set.

### GetResource

`func (o *SshSettingVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SshSettingVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SshSettingVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SshSettingVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetSshEnable

`func (o *SshSettingVO) GetSshEnable() bool`

GetSshEnable returns the SshEnable field if non-nil, zero value otherwise.

### GetSshEnableOk

`func (o *SshSettingVO) GetSshEnableOk() (*bool, bool)`

GetSshEnableOk returns a tuple with the SshEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshEnable

`func (o *SshSettingVO) SetSshEnable(v bool)`

SetSshEnable sets SshEnable field to given value.


### GetSshServerPort

`func (o *SshSettingVO) GetSshServerPort() int32`

GetSshServerPort returns the SshServerPort field if non-nil, zero value otherwise.

### GetSshServerPortOk

`func (o *SshSettingVO) GetSshServerPortOk() (*int32, bool)`

GetSshServerPortOk returns a tuple with the SshServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshServerPort

`func (o *SshSettingVO) SetSshServerPort(v int32)`

SetSshServerPort sets SshServerPort field to given value.

### HasSshServerPort

`func (o *SshSettingVO) HasSshServerPort() bool`

HasSshServerPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


