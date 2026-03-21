# SSHSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layer3Access** | Pointer to **bool** | Whether to enable layer 3 accessibility | [optional] 
**SshEnable** | **bool** | Whether to enable SSH | 
**SshServerPort** | **int32** | SSH server port should be 22 or within the range of 1025-65535. | 

## Methods

### NewSSHSetting

`func NewSSHSetting(sshEnable bool, sshServerPort int32, ) *SSHSetting`

NewSSHSetting instantiates a new SSHSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSHSettingWithDefaults

`func NewSSHSettingWithDefaults() *SSHSetting`

NewSSHSettingWithDefaults instantiates a new SSHSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayer3Access

`func (o *SSHSetting) GetLayer3Access() bool`

GetLayer3Access returns the Layer3Access field if non-nil, zero value otherwise.

### GetLayer3AccessOk

`func (o *SSHSetting) GetLayer3AccessOk() (*bool, bool)`

GetLayer3AccessOk returns a tuple with the Layer3Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayer3Access

`func (o *SSHSetting) SetLayer3Access(v bool)`

SetLayer3Access sets Layer3Access field to given value.

### HasLayer3Access

`func (o *SSHSetting) HasLayer3Access() bool`

HasLayer3Access returns a boolean if a field has been set.

### GetSshEnable

`func (o *SSHSetting) GetSshEnable() bool`

GetSshEnable returns the SshEnable field if non-nil, zero value otherwise.

### GetSshEnableOk

`func (o *SSHSetting) GetSshEnableOk() (*bool, bool)`

GetSshEnableOk returns a tuple with the SshEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshEnable

`func (o *SSHSetting) SetSshEnable(v bool)`

SetSshEnable sets SshEnable field to given value.


### GetSshServerPort

`func (o *SSHSetting) GetSshServerPort() int32`

GetSshServerPort returns the SshServerPort field if non-nil, zero value otherwise.

### GetSshServerPortOk

`func (o *SSHSetting) GetSshServerPortOk() (*int32, bool)`

GetSshServerPortOk returns a tuple with the SshServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshServerPort

`func (o *SSHSetting) SetSshServerPort(v int32)`

SetSshServerPort sets SshServerPort field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


