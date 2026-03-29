# RemoteLogSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Whether to enable the feature | 
**Host** | Pointer to **string** | The IP address of the remote log server | [optional] 
**MoreClientLog** | Pointer to **bool** | Whether it contains client log | [optional] 
**Port** | Pointer to **int32** | Port of the remote log server, port should be within the range of 1-65535 | [optional] 
**Resource** | Pointer to **int32** | Data Source. Resource should be a value as follows: 0: new created; 1: from template; 2: override | [optional] 

## Methods

### NewRemoteLogSettingVO

`func NewRemoteLogSettingVO(enable bool, ) *RemoteLogSettingVO`

NewRemoteLogSettingVO instantiates a new RemoteLogSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteLogSettingVOWithDefaults

`func NewRemoteLogSettingVOWithDefaults() *RemoteLogSettingVO`

NewRemoteLogSettingVOWithDefaults instantiates a new RemoteLogSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RemoteLogSettingVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RemoteLogSettingVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RemoteLogSettingVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetHost

`func (o *RemoteLogSettingVO) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RemoteLogSettingVO) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RemoteLogSettingVO) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RemoteLogSettingVO) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMoreClientLog

`func (o *RemoteLogSettingVO) GetMoreClientLog() bool`

GetMoreClientLog returns the MoreClientLog field if non-nil, zero value otherwise.

### GetMoreClientLogOk

`func (o *RemoteLogSettingVO) GetMoreClientLogOk() (*bool, bool)`

GetMoreClientLogOk returns a tuple with the MoreClientLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreClientLog

`func (o *RemoteLogSettingVO) SetMoreClientLog(v bool)`

SetMoreClientLog sets MoreClientLog field to given value.

### HasMoreClientLog

`func (o *RemoteLogSettingVO) HasMoreClientLog() bool`

HasMoreClientLog returns a boolean if a field has been set.

### GetPort

`func (o *RemoteLogSettingVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemoteLogSettingVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemoteLogSettingVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RemoteLogSettingVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetResource

`func (o *RemoteLogSettingVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RemoteLogSettingVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RemoteLogSettingVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RemoteLogSettingVO) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


