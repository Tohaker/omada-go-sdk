# WireguardPeerOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAddress** | **[]string** | IP/MASK address list of WireGuard peer allowed. | 
**Comment** | Pointer to **string** | The comment of WireGuard peer should contain 0 to 128 characters. | [optional] 
**EndPoint** | Pointer to **string** | The end point IP of WireGuard peer. Get whether supports domain from interface &#39;Get wireguard peer list&#39;. | [optional] 
**EndPointPort** | Pointer to **int32** | The end point port of WireGuard peer should be within the range of 1-65535. | [optional] 
**InterfaceId** | **string** | The ID of WireGuard VPN to which this WireGuard peer binds. The ID can be obtained from &#39;Get all wireguard&#39;s id and name info&#39; interface. | 
**KeepAlive** | **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | 
**Name** | **string** | The name of WireGuard peer should contain 1 to 64 characters. | 
**PresharedKey** | Pointer to **string** | The presharedKey of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | **string** | The public key of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 
**Status** | **bool** | The status of WireGuard peer. Valid value is true or false. | 

## Methods

### NewWireguardPeerOpenApiVO

`func NewWireguardPeerOpenApiVO(allowAddress []string, interfaceId string, keepAlive int32, name string, publicKey string, status bool, ) *WireguardPeerOpenApiVO`

NewWireguardPeerOpenApiVO instantiates a new WireguardPeerOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerOpenApiVOWithDefaults

`func NewWireguardPeerOpenApiVOWithDefaults() *WireguardPeerOpenApiVO`

NewWireguardPeerOpenApiVOWithDefaults instantiates a new WireguardPeerOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAddress

`func (o *WireguardPeerOpenApiVO) GetAllowAddress() []string`

GetAllowAddress returns the AllowAddress field if non-nil, zero value otherwise.

### GetAllowAddressOk

`func (o *WireguardPeerOpenApiVO) GetAllowAddressOk() (*[]string, bool)`

GetAllowAddressOk returns a tuple with the AllowAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAddress

`func (o *WireguardPeerOpenApiVO) SetAllowAddress(v []string)`

SetAllowAddress sets AllowAddress field to given value.


### GetComment

`func (o *WireguardPeerOpenApiVO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WireguardPeerOpenApiVO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WireguardPeerOpenApiVO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WireguardPeerOpenApiVO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndPoint

`func (o *WireguardPeerOpenApiVO) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *WireguardPeerOpenApiVO) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *WireguardPeerOpenApiVO) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *WireguardPeerOpenApiVO) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.

### GetEndPointPort

`func (o *WireguardPeerOpenApiVO) GetEndPointPort() int32`

GetEndPointPort returns the EndPointPort field if non-nil, zero value otherwise.

### GetEndPointPortOk

`func (o *WireguardPeerOpenApiVO) GetEndPointPortOk() (*int32, bool)`

GetEndPointPortOk returns a tuple with the EndPointPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointPort

`func (o *WireguardPeerOpenApiVO) SetEndPointPort(v int32)`

SetEndPointPort sets EndPointPort field to given value.

### HasEndPointPort

`func (o *WireguardPeerOpenApiVO) HasEndPointPort() bool`

HasEndPointPort returns a boolean if a field has been set.

### GetInterfaceId

`func (o *WireguardPeerOpenApiVO) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *WireguardPeerOpenApiVO) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *WireguardPeerOpenApiVO) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.


### GetKeepAlive

`func (o *WireguardPeerOpenApiVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *WireguardPeerOpenApiVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *WireguardPeerOpenApiVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.


### GetName

`func (o *WireguardPeerOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardPeerOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardPeerOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPresharedKey

`func (o *WireguardPeerOpenApiVO) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *WireguardPeerOpenApiVO) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *WireguardPeerOpenApiVO) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *WireguardPeerOpenApiVO) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardPeerOpenApiVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardPeerOpenApiVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardPeerOpenApiVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetStatus

`func (o *WireguardPeerOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardPeerOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardPeerOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


