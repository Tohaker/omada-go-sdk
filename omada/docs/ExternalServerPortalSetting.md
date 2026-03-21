# ExternalServerPortalSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HostType** | **int32** | Host type, should be a value as follows: 1: IP; 2: URL | 
**ServerIp** | Pointer to **string** | Server IP, required when [hostType] is 1, pattern as \&quot;xx.xx.xx.xx\&quot;. | [optional] 
**ServerPort** | Pointer to **int32** | Server port, required when [hostType] is 1, from 1 to 65535. | [optional] 
**ServerUrl** | Pointer to **string** | Server url, required when [hostType] is 2. | [optional] 
**ServerUrlScheme** | Pointer to **string** | Server url scheme, required when [hostType] is 2, value is http or https. | [optional] 

## Methods

### NewExternalServerPortalSetting

`func NewExternalServerPortalSetting(hostType int32, ) *ExternalServerPortalSetting`

NewExternalServerPortalSetting instantiates a new ExternalServerPortalSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalServerPortalSettingWithDefaults

`func NewExternalServerPortalSettingWithDefaults() *ExternalServerPortalSetting`

NewExternalServerPortalSettingWithDefaults instantiates a new ExternalServerPortalSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHostType

`func (o *ExternalServerPortalSetting) GetHostType() int32`

GetHostType returns the HostType field if non-nil, zero value otherwise.

### GetHostTypeOk

`func (o *ExternalServerPortalSetting) GetHostTypeOk() (*int32, bool)`

GetHostTypeOk returns a tuple with the HostType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostType

`func (o *ExternalServerPortalSetting) SetHostType(v int32)`

SetHostType sets HostType field to given value.


### GetServerIp

`func (o *ExternalServerPortalSetting) GetServerIp() string`

GetServerIp returns the ServerIp field if non-nil, zero value otherwise.

### GetServerIpOk

`func (o *ExternalServerPortalSetting) GetServerIpOk() (*string, bool)`

GetServerIpOk returns a tuple with the ServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIp

`func (o *ExternalServerPortalSetting) SetServerIp(v string)`

SetServerIp sets ServerIp field to given value.

### HasServerIp

`func (o *ExternalServerPortalSetting) HasServerIp() bool`

HasServerIp returns a boolean if a field has been set.

### GetServerPort

`func (o *ExternalServerPortalSetting) GetServerPort() int32`

GetServerPort returns the ServerPort field if non-nil, zero value otherwise.

### GetServerPortOk

`func (o *ExternalServerPortalSetting) GetServerPortOk() (*int32, bool)`

GetServerPortOk returns a tuple with the ServerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPort

`func (o *ExternalServerPortalSetting) SetServerPort(v int32)`

SetServerPort sets ServerPort field to given value.

### HasServerPort

`func (o *ExternalServerPortalSetting) HasServerPort() bool`

HasServerPort returns a boolean if a field has been set.

### GetServerUrl

`func (o *ExternalServerPortalSetting) GetServerUrl() string`

GetServerUrl returns the ServerUrl field if non-nil, zero value otherwise.

### GetServerUrlOk

`func (o *ExternalServerPortalSetting) GetServerUrlOk() (*string, bool)`

GetServerUrlOk returns a tuple with the ServerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrl

`func (o *ExternalServerPortalSetting) SetServerUrl(v string)`

SetServerUrl sets ServerUrl field to given value.

### HasServerUrl

`func (o *ExternalServerPortalSetting) HasServerUrl() bool`

HasServerUrl returns a boolean if a field has been set.

### GetServerUrlScheme

`func (o *ExternalServerPortalSetting) GetServerUrlScheme() string`

GetServerUrlScheme returns the ServerUrlScheme field if non-nil, zero value otherwise.

### GetServerUrlSchemeOk

`func (o *ExternalServerPortalSetting) GetServerUrlSchemeOk() (*string, bool)`

GetServerUrlSchemeOk returns a tuple with the ServerUrlScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUrlScheme

`func (o *ExternalServerPortalSetting) SetServerUrlScheme(v string)`

SetServerUrlScheme sets ServerUrlScheme field to given value.

### HasServerUrlScheme

`func (o *ExternalServerPortalSetting) HasServerUrlScheme() bool`

HasServerUrlScheme returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


