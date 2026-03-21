# RemoteLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Whether to enable the feature | 
**Host** | Pointer to **string** | The IP address of the remote log server | [optional] 
**Port** | Pointer to **int32** | Port of the remote log server, port should be within the range of 1-65535 | [optional] 

## Methods

### NewRemoteLog

`func NewRemoteLog(enable bool, ) *RemoteLog`

NewRemoteLog instantiates a new RemoteLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteLogWithDefaults

`func NewRemoteLogWithDefaults() *RemoteLog`

NewRemoteLogWithDefaults instantiates a new RemoteLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RemoteLog) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RemoteLog) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RemoteLog) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetHost

`func (o *RemoteLog) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *RemoteLog) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *RemoteLog) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *RemoteLog) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *RemoteLog) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RemoteLog) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RemoteLog) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *RemoteLog) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


