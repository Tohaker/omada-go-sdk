# ClientIpSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | Pointer to **string** | Client IP. | [optional] 
**NetId** | Pointer to **string** | LAN network ID. | [optional] 
**ServerMac** | Pointer to **string** | DHCP Server Mac. | [optional] 
**ServerStackId** | Pointer to **string** | DHCP Sever Stack ID. | [optional] 
**ServerType** | Pointer to **string** | DHCP Sever Type. | [optional] 
**UseFixedAddr** | **bool** | Whether to use the specified IP. | 

## Methods

### NewClientIpSetting

`func NewClientIpSetting(useFixedAddr bool, ) *ClientIpSetting`

NewClientIpSetting instantiates a new ClientIpSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientIpSettingWithDefaults

`func NewClientIpSettingWithDefaults() *ClientIpSetting`

NewClientIpSettingWithDefaults instantiates a new ClientIpSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *ClientIpSetting) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientIpSetting) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientIpSetting) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientIpSetting) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetNetId

`func (o *ClientIpSetting) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *ClientIpSetting) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *ClientIpSetting) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *ClientIpSetting) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetServerMac

`func (o *ClientIpSetting) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *ClientIpSetting) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *ClientIpSetting) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *ClientIpSetting) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *ClientIpSetting) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *ClientIpSetting) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *ClientIpSetting) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *ClientIpSetting) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *ClientIpSetting) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ClientIpSetting) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ClientIpSetting) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ClientIpSetting) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUseFixedAddr

`func (o *ClientIpSetting) GetUseFixedAddr() bool`

GetUseFixedAddr returns the UseFixedAddr field if non-nil, zero value otherwise.

### GetUseFixedAddrOk

`func (o *ClientIpSetting) GetUseFixedAddrOk() (*bool, bool)`

GetUseFixedAddrOk returns a tuple with the UseFixedAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedAddr

`func (o *ClientIpSetting) SetUseFixedAddr(v bool)`

SetUseFixedAddr sets UseFixedAddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


