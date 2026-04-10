# ServerWireGuardClientsConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowedAddress** | Pointer to **[]string** | IP/MASK address list of server WireGuard peer allowed. | [optional] 
**AllowedAddressStatus** | Pointer to **bool** | When parameter [allowedAddressStatus] is false, server WireGuard peer allowed address is client interface IP. | [optional] 
**InterfaceIp** | **string** | IP address. | 
**Name** | **string** | Client name. | 
**PreSharedKey** | Pointer to **string** | The preSharedKey of WireGuard peer must have 44 character of base64 and end with &#39;&#x3D;&#39;. | [optional] 
**PrivateKey** | **string** | The private key of WireGuard VPN must have 44 character of base64 and end with &#39;&#x3D;&#39;. | 

## Methods

### NewServerWireGuardClientsConfigVO

`func NewServerWireGuardClientsConfigVO(interfaceIp string, name string, privateKey string, ) *ServerWireGuardClientsConfigVO`

NewServerWireGuardClientsConfigVO instantiates a new ServerWireGuardClientsConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerWireGuardClientsConfigVOWithDefaults

`func NewServerWireGuardClientsConfigVOWithDefaults() *ServerWireGuardClientsConfigVO`

NewServerWireGuardClientsConfigVOWithDefaults instantiates a new ServerWireGuardClientsConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowedAddress

`func (o *ServerWireGuardClientsConfigVO) GetAllowedAddress() []string`

GetAllowedAddress returns the AllowedAddress field if non-nil, zero value otherwise.

### GetAllowedAddressOk

`func (o *ServerWireGuardClientsConfigVO) GetAllowedAddressOk() (*[]string, bool)`

GetAllowedAddressOk returns a tuple with the AllowedAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddress

`func (o *ServerWireGuardClientsConfigVO) SetAllowedAddress(v []string)`

SetAllowedAddress sets AllowedAddress field to given value.

### HasAllowedAddress

`func (o *ServerWireGuardClientsConfigVO) HasAllowedAddress() bool`

HasAllowedAddress returns a boolean if a field has been set.

### GetAllowedAddressStatus

`func (o *ServerWireGuardClientsConfigVO) GetAllowedAddressStatus() bool`

GetAllowedAddressStatus returns the AllowedAddressStatus field if non-nil, zero value otherwise.

### GetAllowedAddressStatusOk

`func (o *ServerWireGuardClientsConfigVO) GetAllowedAddressStatusOk() (*bool, bool)`

GetAllowedAddressStatusOk returns a tuple with the AllowedAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedAddressStatus

`func (o *ServerWireGuardClientsConfigVO) SetAllowedAddressStatus(v bool)`

SetAllowedAddressStatus sets AllowedAddressStatus field to given value.

### HasAllowedAddressStatus

`func (o *ServerWireGuardClientsConfigVO) HasAllowedAddressStatus() bool`

HasAllowedAddressStatus returns a boolean if a field has been set.

### GetInterfaceIp

`func (o *ServerWireGuardClientsConfigVO) GetInterfaceIp() string`

GetInterfaceIp returns the InterfaceIp field if non-nil, zero value otherwise.

### GetInterfaceIpOk

`func (o *ServerWireGuardClientsConfigVO) GetInterfaceIpOk() (*string, bool)`

GetInterfaceIpOk returns a tuple with the InterfaceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIp

`func (o *ServerWireGuardClientsConfigVO) SetInterfaceIp(v string)`

SetInterfaceIp sets InterfaceIp field to given value.


### GetName

`func (o *ServerWireGuardClientsConfigVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServerWireGuardClientsConfigVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServerWireGuardClientsConfigVO) SetName(v string)`

SetName sets Name field to given value.


### GetPreSharedKey

`func (o *ServerWireGuardClientsConfigVO) GetPreSharedKey() string`

GetPreSharedKey returns the PreSharedKey field if non-nil, zero value otherwise.

### GetPreSharedKeyOk

`func (o *ServerWireGuardClientsConfigVO) GetPreSharedKeyOk() (*string, bool)`

GetPreSharedKeyOk returns a tuple with the PreSharedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreSharedKey

`func (o *ServerWireGuardClientsConfigVO) SetPreSharedKey(v string)`

SetPreSharedKey sets PreSharedKey field to given value.

### HasPreSharedKey

`func (o *ServerWireGuardClientsConfigVO) HasPreSharedKey() bool`

HasPreSharedKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *ServerWireGuardClientsConfigVO) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *ServerWireGuardClientsConfigVO) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *ServerWireGuardClientsConfigVO) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


