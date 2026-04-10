# ServerWireGuardClientsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAddress** | Pointer to **[]string** | IP/MASK address list of server WireGuard peer allowed. | [optional] 
**AllowedAddressStatus** | Pointer to **bool** | When parameter [allowedAddressStatus] is false, server WireGuard peer allowed address is client interface IP. | [optional] 
**Id** | Pointer to **string** | ID of the client. | [optional] 
**InterfaceIp** | Pointer to **string** | IP address. | [optional] 
**Name** | Pointer to **string** | Client name. | [optional] 
**PreSharedKey** | Pointer to **string** | The preSharedKey of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PrivateKey** | Pointer to **string** | The privateKey key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PublicKey** | Pointer to **string** | The public key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 

## Methods

### NewServerWireGuardClientsVO

`func NewServerWireGuardClientsVO() *ServerWireGuardClientsVO`

NewServerWireGuardClientsVO instantiates a new ServerWireGuardClientsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWireGuardClientsVOWithDefaults

`func NewServerWireGuardClientsVOWithDefaults() *ServerWireGuardClientsVO`

NewServerWireGuardClientsVOWithDefaults instantiates a new ServerWireGuardClientsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAddress

`func (o *ServerWireGuardClientsVO) GetAllowedAddress() []string`

GetAllowedAddress returns the AllowedAddress field if non-nil, zero value otherwise.

### GetAllowedAddressOk

`func (o *ServerWireGuardClientsVO) GetAllowedAddressOk() (*[]string, bool)`

GetAllowedAddressOk returns a tuple with the AllowedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddress

`func (o *ServerWireGuardClientsVO) SetAllowedAddress(v []string)`

SetAllowedAddress sets AllowedAddress field to given value.

### HasAllowedAddress

`func (o *ServerWireGuardClientsVO) HasAllowedAddress() bool`

HasAllowedAddress returns a boolean if a field has been set.

### GetAllowedAddressStatus

`func (o *ServerWireGuardClientsVO) GetAllowedAddressStatus() bool`

GetAllowedAddressStatus returns the AllowedAddressStatus field if non-nil, zero value otherwise.

### GetAllowedAddressStatusOk

`func (o *ServerWireGuardClientsVO) GetAllowedAddressStatusOk() (*bool, bool)`

GetAllowedAddressStatusOk returns a tuple with the AllowedAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddressStatus

`func (o *ServerWireGuardClientsVO) SetAllowedAddressStatus(v bool)`

SetAllowedAddressStatus sets AllowedAddressStatus field to given value.

### HasAllowedAddressStatus

`func (o *ServerWireGuardClientsVO) HasAllowedAddressStatus() bool`

HasAllowedAddressStatus returns a boolean if a field has been set.

### GetId

`func (o *ServerWireGuardClientsVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServerWireGuardClientsVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServerWireGuardClientsVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServerWireGuardClientsVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *ServerWireGuardClientsVO) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *ServerWireGuardClientsVO) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *ServerWireGuardClientsVO) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.

### HasInterfaceIp

`func (o *ServerWireGuardClientsVO) HasInterfaceIp() bool`

HasInterfaceIp returns a boolean if a field has been set.

### GetName

`func (o *ServerWireGuardClientsVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerWireGuardClientsVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerWireGuardClientsVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServerWireGuardClientsVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPreSharedKey

`func (o *ServerWireGuardClientsVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *ServerWireGuardClientsVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *ServerWireGuardClientsVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *ServerWireGuardClientsVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *ServerWireGuardClientsVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ServerWireGuardClientsVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ServerWireGuardClientsVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *ServerWireGuardClientsVO) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPublicKey

`func (o *ServerWireGuardClientsVO) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ServerWireGuardClientsVO) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ServerWireGuardClientsVO) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ServerWireGuardClientsVO) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


