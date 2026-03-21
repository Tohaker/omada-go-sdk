# CustomerRemoteLogTipOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Whether to enable the feature | 
**Host** | Pointer to **string** | The IP address of the remote log server | [optional] 
**Port** | Pointer to **int32** | Port of the remote log server, port should be within the range of 1-65535 | [optional] 

## Methods

### NewCustomerRemoteLogTipOpenApiVO

`func NewCustomerRemoteLogTipOpenApiVO(enable bool, ) *CustomerRemoteLogTipOpenApiVO`

NewCustomerRemoteLogTipOpenApiVO instantiates a new CustomerRemoteLogTipOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerRemoteLogTipOpenApiVOWithDefaults

`func NewCustomerRemoteLogTipOpenApiVOWithDefaults() *CustomerRemoteLogTipOpenApiVO`

NewCustomerRemoteLogTipOpenApiVOWithDefaults instantiates a new CustomerRemoteLogTipOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *CustomerRemoteLogTipOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *CustomerRemoteLogTipOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *CustomerRemoteLogTipOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetHost

`func (o *CustomerRemoteLogTipOpenApiVO) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *CustomerRemoteLogTipOpenApiVO) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *CustomerRemoteLogTipOpenApiVO) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *CustomerRemoteLogTipOpenApiVO) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *CustomerRemoteLogTipOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *CustomerRemoteLogTipOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *CustomerRemoteLogTipOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *CustomerRemoteLogTipOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


