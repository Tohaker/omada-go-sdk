# RemoteLogSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Whether to enable the feature | 
**Host** | Pointer to **string** | The IP address of the remote log server | [optional] 
**MoreClientLog** | Pointer to **bool** | Whether it contains client log | [optional] 
**Port** | Pointer to **int32** | Port of the remote log server, port should be within the range of 1-65535 | [optional] 

## Methods

### NewRemoteLogSettingOpenApiVO

`func NewRemoteLogSettingOpenApiVO(enable bool, ) *RemoteLogSettingOpenApiVO`

NewRemoteLogSettingOpenApiVO instantiates a new RemoteLogSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteLogSettingOpenApiVOWithDefaults

`func NewRemoteLogSettingOpenApiVOWithDefaults() *RemoteLogSettingOpenApiVO`

NewRemoteLogSettingOpenApiVOWithDefaults instantiates a new RemoteLogSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RemoteLogSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RemoteLogSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RemoteLogSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetHost

`func (o *RemoteLogSettingOpenApiVO) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RemoteLogSettingOpenApiVO) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RemoteLogSettingOpenApiVO) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RemoteLogSettingOpenApiVO) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetMoreClientLog

`func (o *RemoteLogSettingOpenApiVO) GetMoreClientLog() bool`

GetMoreClientLog returns the MoreClientLog field if non-nil, zero value otherwise.

### GetMoreClientLogOk

`func (o *RemoteLogSettingOpenApiVO) GetMoreClientLogOk() (*bool, bool)`

GetMoreClientLogOk returns a tuple with the MoreClientLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoreClientLog

`func (o *RemoteLogSettingOpenApiVO) SetMoreClientLog(v bool)`

SetMoreClientLog sets MoreClientLog field to given value.

### HasMoreClientLog

`func (o *RemoteLogSettingOpenApiVO) HasMoreClientLog() bool`

HasMoreClientLog returns a boolean if a field has been set.

### GetPort

`func (o *RemoteLogSettingOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemoteLogSettingOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemoteLogSettingOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RemoteLogSettingOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


