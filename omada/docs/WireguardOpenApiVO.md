# WireguardOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListenPort** | **int32** | The listening port for WireGuard VPN should be within the range of 1-65535. | 
**LocalIp** | **string** | The local IP address of WireGuard VPN. | 
**Mtu** | **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | 
**Name** | **string** | The name of WireGuard VPN should contain 1 to 64 characters. | 
**PrivateKey** | **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 
**Status** | **bool** | The status of WireGuard VPN. Valid value is true or false. | 

## Methods

### NewWireguardOpenApiVO

`func NewWireguardOpenApiVO(listenPort int32, localIp string, mtu int32, name string, privateKey string, status bool, ) *WireguardOpenApiVO`

NewWireguardOpenApiVO instantiates a new WireguardOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardOpenApiVOWithDefaults

`func NewWireguardOpenApiVOWithDefaults() *WireguardOpenApiVO`

NewWireguardOpenApiVOWithDefaults instantiates a new WireguardOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListenPort

`func (o *WireguardOpenApiVO) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *WireguardOpenApiVO) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *WireguardOpenApiVO) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.


### GetLocalIp

`func (o *WireguardOpenApiVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *WireguardOpenApiVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *WireguardOpenApiVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.


### GetMtu

`func (o *WireguardOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *WireguardOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *WireguardOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.


### GetName

`func (o *WireguardOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPrivateKey

`func (o *WireguardOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WireguardOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WireguardOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetStatus

`func (o *WireguardOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


