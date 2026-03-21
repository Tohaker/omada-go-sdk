# WireguardDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of WireGuard VPN. | [optional] 
**ListenPort** | Pointer to **int32** | The listening port for WireGuard VPN should be within the range of 1-65535. | [optional] 
**LocalIp** | Pointer to **string** | The local IP address of WireGuard VPN. | [optional] 
**Mtu** | Pointer to **int32** | The MTU of WireGuard VPN should be within the range of 576-1440. | [optional] 
**Name** | Pointer to **string** | The name of WireGuard VPN should contain 1 to 64 characters. | [optional] 
**PrivateKey** | Pointer to **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | Pointer to **string** | The publicKey key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**Status** | Pointer to **bool** | The status of WireGuard VPN. | [optional] 

## Methods

### NewWireguardDetailOpenApiVO

`func NewWireguardDetailOpenApiVO() *WireguardDetailOpenApiVO`

NewWireguardDetailOpenApiVO instantiates a new WireguardDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardDetailOpenApiVOWithDefaults

`func NewWireguardDetailOpenApiVOWithDefaults() *WireguardDetailOpenApiVO`

NewWireguardDetailOpenApiVOWithDefaults instantiates a new WireguardDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WireguardDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireguardDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetListenPort

`func (o *WireguardDetailOpenApiVO) GetListenPort() int32`

GetListenPort returns the ListenPort field if non-nil, zero value otherwise.

### GetListenPortOk

`func (o *WireguardDetailOpenApiVO) GetListenPortOk() (*int32, bool)`

GetListenPortOk returns a tuple with the ListenPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListenPort

`func (o *WireguardDetailOpenApiVO) SetListenPort(v int32)`

SetListenPort sets ListenPort field to given value.

### HasListenPort

`func (o *WireguardDetailOpenApiVO) HasListenPort() bool`

HasListenPort returns a boolean if a field has been set.

### GetLocalIp

`func (o *WireguardDetailOpenApiVO) GetLocalIp() string`

GetLocalIp returns the LocalIp field if non-nil, zero value otherwise.

### GetLocalIpOk

`func (o *WireguardDetailOpenApiVO) GetLocalIpOk() (*string, bool)`

GetLocalIpOk returns a tuple with the LocalIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIp

`func (o *WireguardDetailOpenApiVO) SetLocalIp(v string)`

SetLocalIp sets LocalIp field to given value.

### HasLocalIp

`func (o *WireguardDetailOpenApiVO) HasLocalIp() bool`

HasLocalIp returns a boolean if a field has been set.

### GetMtu

`func (o *WireguardDetailOpenApiVO) GetMtu() int32`

GetMtu returns the Mtu field if non-nil, zero value otherwise.

### GetMtuOk

`func (o *WireguardDetailOpenApiVO) GetMtuOk() (*int32, bool)`

GetMtuOk returns a tuple with the Mtu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtu

`func (o *WireguardDetailOpenApiVO) SetMtu(v int32)`

SetMtu sets Mtu field to given value.

### HasMtu

`func (o *WireguardDetailOpenApiVO) HasMtu() bool`

HasMtu returns a boolean if a field has been set.

### GetName

`func (o *WireguardDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WireguardDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrivateKey

`func (o *WireguardDetailOpenApiVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *WireguardDetailOpenApiVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *WireguardDetailOpenApiVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *WireguardDetailOpenApiVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardDetailOpenApiVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardDetailOpenApiVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardDetailOpenApiVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardDetailOpenApiVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetStatus

`func (o *WireguardDetailOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardDetailOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardDetailOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WireguardDetailOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


