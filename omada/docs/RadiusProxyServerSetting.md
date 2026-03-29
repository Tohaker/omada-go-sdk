# RadiusProxyServerSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Whether the RADIUS proxy server is enabled. | 
**Port** | **int32** | Server port. | 
**Secret** | **string** | Server secrect. | 

## Methods

### NewRadiusProxyServerSetting

`func NewRadiusProxyServerSetting(enable bool, port int32, secret string, ) *RadiusProxyServerSetting`

NewRadiusProxyServerSetting instantiates a new RadiusProxyServerSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusProxyServerSettingWithDefaults

`func NewRadiusProxyServerSettingWithDefaults() *RadiusProxyServerSetting`

NewRadiusProxyServerSettingWithDefaults instantiates a new RadiusProxyServerSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *RadiusProxyServerSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *RadiusProxyServerSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *RadiusProxyServerSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetPort

`func (o *RadiusProxyServerSetting) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *RadiusProxyServerSetting) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *RadiusProxyServerSetting) SetPort(v int32)`

SetPort sets Port field to given value.


### GetSecret

`func (o *RadiusProxyServerSetting) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *RadiusProxyServerSetting) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *RadiusProxyServerSetting) SetSecret(v string)`

SetSecret sets Secret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


