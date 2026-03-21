# WireguardPeerDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowAddress** | Pointer to **[]string** | IP address list of WireGuard peer allowed. | [optional] 
**Comment** | Pointer to **string** | The comment of WireGuard peer should contain 0 to 128 characters. | [optional] 
**EndPoint** | Pointer to **string** | The end point of WireGuard peer. | [optional] 
**EndPointPort** | Pointer to **int32** | The end point port of WireGuard peer should be within the range of 1-65535. | [optional] 
**ExistDomain** | Pointer to **bool** | Whether a domain name has been configured in Endpoint. | [optional] 
**Id** | Pointer to **string** | The ID of WireGuard peer. | [optional] 
**InterfaceId** | Pointer to **string** | The ID of WireGuard VPN to which this WireGuard peer binds. The ID can be obtained from &#39;Get all wireguard&#39;s id and name info&#39; interface. | [optional] 
**InterfaceName** | Pointer to **string** | The name of WireGuard VPN to which this WireGuard peer binds. | [optional] 
**KeepAlive** | Pointer to **int32** | The keepalive second of WireGuard peer should be within the range of 0-65535. | [optional] 
**Name** | Pointer to **string** | The name of WireGuard peer should contain 1 to 64 characters. | [optional] 
**PresharedKey** | Pointer to **string** | The presharedKey of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | Pointer to **string** | The public key of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**Status** | Pointer to **bool** | The status of WireGuard peer. | [optional] 

## Methods

### NewWireguardPeerDetailOpenApiVO

`func NewWireguardPeerDetailOpenApiVO() *WireguardPeerDetailOpenApiVO`

NewWireguardPeerDetailOpenApiVO instantiates a new WireguardPeerDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWireguardPeerDetailOpenApiVOWithDefaults

`func NewWireguardPeerDetailOpenApiVOWithDefaults() *WireguardPeerDetailOpenApiVO`

NewWireguardPeerDetailOpenApiVOWithDefaults instantiates a new WireguardPeerDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowAddress

`func (o *WireguardPeerDetailOpenApiVO) GetAllowAddress() []string`

GetAllowAddress returns the AllowAddress field if non-nil, zero value otherwise.

### GetAllowAddressOk

`func (o *WireguardPeerDetailOpenApiVO) GetAllowAddressOk() (*[]string, bool)`

GetAllowAddressOk returns a tuple with the AllowAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowAddress

`func (o *WireguardPeerDetailOpenApiVO) SetAllowAddress(v []string)`

SetAllowAddress sets AllowAddress field to given value.

### HasAllowAddress

`func (o *WireguardPeerDetailOpenApiVO) HasAllowAddress() bool`

HasAllowAddress returns a boolean if a field has been set.

### GetComment

`func (o *WireguardPeerDetailOpenApiVO) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *WireguardPeerDetailOpenApiVO) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *WireguardPeerDetailOpenApiVO) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *WireguardPeerDetailOpenApiVO) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetEndPoint

`func (o *WireguardPeerDetailOpenApiVO) GetEndPoint() string`

GetEndPoint returns the EndPoint field if non-nil, zero value otherwise.

### GetEndPointOk

`func (o *WireguardPeerDetailOpenApiVO) GetEndPointOk() (*string, bool)`

GetEndPointOk returns a tuple with the EndPoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPoint

`func (o *WireguardPeerDetailOpenApiVO) SetEndPoint(v string)`

SetEndPoint sets EndPoint field to given value.

### HasEndPoint

`func (o *WireguardPeerDetailOpenApiVO) HasEndPoint() bool`

HasEndPoint returns a boolean if a field has been set.

### GetEndPointPort

`func (o *WireguardPeerDetailOpenApiVO) GetEndPointPort() int32`

GetEndPointPort returns the EndPointPort field if non-nil, zero value otherwise.

### GetEndPointPortOk

`func (o *WireguardPeerDetailOpenApiVO) GetEndPointPortOk() (*int32, bool)`

GetEndPointPortOk returns a tuple with the EndPointPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPointPort

`func (o *WireguardPeerDetailOpenApiVO) SetEndPointPort(v int32)`

SetEndPointPort sets EndPointPort field to given value.

### HasEndPointPort

`func (o *WireguardPeerDetailOpenApiVO) HasEndPointPort() bool`

HasEndPointPort returns a boolean if a field has been set.

### GetExistDomain

`func (o *WireguardPeerDetailOpenApiVO) GetExistDomain() bool`

GetExistDomain returns the ExistDomain field if non-nil, zero value otherwise.

### GetExistDomainOk

`func (o *WireguardPeerDetailOpenApiVO) GetExistDomainOk() (*bool, bool)`

GetExistDomainOk returns a tuple with the ExistDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExistDomain

`func (o *WireguardPeerDetailOpenApiVO) SetExistDomain(v bool)`

SetExistDomain sets ExistDomain field to given value.

### HasExistDomain

`func (o *WireguardPeerDetailOpenApiVO) HasExistDomain() bool`

HasExistDomain returns a boolean if a field has been set.

### GetId

`func (o *WireguardPeerDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WireguardPeerDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WireguardPeerDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WireguardPeerDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *WireguardPeerDetailOpenApiVO) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *WireguardPeerDetailOpenApiVO) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *WireguardPeerDetailOpenApiVO) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *WireguardPeerDetailOpenApiVO) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *WireguardPeerDetailOpenApiVO) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *WireguardPeerDetailOpenApiVO) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *WireguardPeerDetailOpenApiVO) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *WireguardPeerDetailOpenApiVO) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetKeepAlive

`func (o *WireguardPeerDetailOpenApiVO) GetKeepAlive() int32`

GetKeepAlive returns the KeepAlive field if non-nil, zero value otherwise.

### GetKeepAliveOk

`func (o *WireguardPeerDetailOpenApiVO) GetKeepAliveOk() (*int32, bool)`

GetKeepAliveOk returns a tuple with the KeepAlive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepAlive

`func (o *WireguardPeerDetailOpenApiVO) SetKeepAlive(v int32)`

SetKeepAlive sets KeepAlive field to given value.

### HasKeepAlive

`func (o *WireguardPeerDetailOpenApiVO) HasKeepAlive() bool`

HasKeepAlive returns a boolean if a field has been set.

### GetName

`func (o *WireguardPeerDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WireguardPeerDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WireguardPeerDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WireguardPeerDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPresharedKey

`func (o *WireguardPeerDetailOpenApiVO) GetPresharedKey() string`

GetPresharedKey returns the PresharedKey field if non-nil, zero value otherwise.

### GetPresharedKeyOk

`func (o *WireguardPeerDetailOpenApiVO) GetPresharedKeyOk() (*string, bool)`

GetPresharedKeyOk returns a tuple with the PresharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresharedKey

`func (o *WireguardPeerDetailOpenApiVO) SetPresharedKey(v string)`

SetPresharedKey sets PresharedKey field to given value.

### HasPresharedKey

`func (o *WireguardPeerDetailOpenApiVO) HasPresharedKey() bool`

HasPresharedKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *WireguardPeerDetailOpenApiVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *WireguardPeerDetailOpenApiVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *WireguardPeerDetailOpenApiVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *WireguardPeerDetailOpenApiVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetStatus

`func (o *WireguardPeerDetailOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WireguardPeerDetailOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WireguardPeerDetailOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WireguardPeerDetailOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


