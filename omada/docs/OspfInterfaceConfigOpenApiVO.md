# OspfInterfaceConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **int32** | Authentication Type, it should be a value as follows: 0: None, 1: Simple, 2: MD5. | 
**Cost** | **int32** | The link cost. OSPF uses this value in computing shortest paths. It should be within the range of 1–65535. | 
**DeadInterval** | Pointer to **int32** | The dead interval for the specified interface in seconds. This specifies how long a router will wait to see a neighbor router&#39;s Hello packets before declaring that the router is down. This parameter must be the same for all routers attached to a network. It should be within the range of 1-65535 seconds and the default is 40. | [optional] 
**DeviceName** | **string** | Device Name | 
**HelloInterval** | **int32** | The hello interval for the specified interface in seconds. This parameter must be the same for all routers attached to a network. It should be within the range of 1-65535 seconds and the default is 10 seconds. | 
**Mac** | **string** | Device Mac | 
**Md5Key** | Pointer to **string** | Displays the key used for md5 authentication, its value should be within the range of 1-16. | [optional] 
**Md5KeyId** | Pointer to **int32** | Displays the key ID used for md5 authentication, its value should be within the range of 1-255. | [optional] 
**NetworkType** | **int32** | Network Type, it should be a value as follows: 0: Broadcast, 1: Non-Broadcast, 2: Point-to-Multipoint, 3: Point-to-Point. The default network type for Ethernet interfaces is broadcast. | 
**SimpleKey** | Pointer to **string** | Displays the key used for simple authentication, its value should be within the range of 1-8. | [optional] 
**VlanId** | **int32** | Vlan ID should be within the range of 1–4094. | 
**VlanInterfaceId** | **string** | Vlan Interface ID | 
**VlanInterfaceName** | **string** | Vlan Interface Name | 

## Methods

### NewOspfInterfaceConfigOpenApiVO

`func NewOspfInterfaceConfigOpenApiVO(authenticationType int32, cost int32, deviceName string, helloInterval int32, mac string, networkType int32, vlanId int32, vlanInterfaceId string, vlanInterfaceName string, ) *OspfInterfaceConfigOpenApiVO`

NewOspfInterfaceConfigOpenApiVO instantiates a new OspfInterfaceConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfInterfaceConfigOpenApiVOWithDefaults

`func NewOspfInterfaceConfigOpenApiVOWithDefaults() *OspfInterfaceConfigOpenApiVO`

NewOspfInterfaceConfigOpenApiVOWithDefaults instantiates a new OspfInterfaceConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *OspfInterfaceConfigOpenApiVO) GetAuthenticationType() int32`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *OspfInterfaceConfigOpenApiVO) GetAuthenticationTypeOk() (*int32, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *OspfInterfaceConfigOpenApiVO) SetAuthenticationType(v int32)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetCost

`func (o *OspfInterfaceConfigOpenApiVO) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OspfInterfaceConfigOpenApiVO) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OspfInterfaceConfigOpenApiVO) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetDeadInterval

`func (o *OspfInterfaceConfigOpenApiVO) GetDeadInterval() int32`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *OspfInterfaceConfigOpenApiVO) GetDeadIntervalOk() (*int32, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *OspfInterfaceConfigOpenApiVO) SetDeadInterval(v int32)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *OspfInterfaceConfigOpenApiVO) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetDeviceName

`func (o *OspfInterfaceConfigOpenApiVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *OspfInterfaceConfigOpenApiVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *OspfInterfaceConfigOpenApiVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetHelloInterval

`func (o *OspfInterfaceConfigOpenApiVO) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *OspfInterfaceConfigOpenApiVO) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *OspfInterfaceConfigOpenApiVO) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.


### GetMac

`func (o *OspfInterfaceConfigOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OspfInterfaceConfigOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OspfInterfaceConfigOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMd5Key

`func (o *OspfInterfaceConfigOpenApiVO) GetMd5Key() string`

GetMd5Key returns the Md5Key field if non-nil, zero value otherwise.

### GetMd5KeyOk

`func (o *OspfInterfaceConfigOpenApiVO) GetMd5KeyOk() (*string, bool)`

GetMd5KeyOk returns a tuple with the Md5Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Key

`func (o *OspfInterfaceConfigOpenApiVO) SetMd5Key(v string)`

SetMd5Key sets Md5Key field to given value.

### HasMd5Key

`func (o *OspfInterfaceConfigOpenApiVO) HasMd5Key() bool`

HasMd5Key returns a boolean if a field has been set.

### GetMd5KeyId

`func (o *OspfInterfaceConfigOpenApiVO) GetMd5KeyId() int32`

GetMd5KeyId returns the Md5KeyId field if non-nil, zero value otherwise.

### GetMd5KeyIdOk

`func (o *OspfInterfaceConfigOpenApiVO) GetMd5KeyIdOk() (*int32, bool)`

GetMd5KeyIdOk returns a tuple with the Md5KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5KeyId

`func (o *OspfInterfaceConfigOpenApiVO) SetMd5KeyId(v int32)`

SetMd5KeyId sets Md5KeyId field to given value.

### HasMd5KeyId

`func (o *OspfInterfaceConfigOpenApiVO) HasMd5KeyId() bool`

HasMd5KeyId returns a boolean if a field has been set.

### GetNetworkType

`func (o *OspfInterfaceConfigOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *OspfInterfaceConfigOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *OspfInterfaceConfigOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.


### GetSimpleKey

`func (o *OspfInterfaceConfigOpenApiVO) GetSimpleKey() string`

GetSimpleKey returns the SimpleKey field if non-nil, zero value otherwise.

### GetSimpleKeyOk

`func (o *OspfInterfaceConfigOpenApiVO) GetSimpleKeyOk() (*string, bool)`

GetSimpleKeyOk returns a tuple with the SimpleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleKey

`func (o *OspfInterfaceConfigOpenApiVO) SetSimpleKey(v string)`

SetSimpleKey sets SimpleKey field to given value.

### HasSimpleKey

`func (o *OspfInterfaceConfigOpenApiVO) HasSimpleKey() bool`

HasSimpleKey returns a boolean if a field has been set.

### GetVlanId

`func (o *OspfInterfaceConfigOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *OspfInterfaceConfigOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *OspfInterfaceConfigOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetVlanInterfaceId

`func (o *OspfInterfaceConfigOpenApiVO) GetVlanInterfaceId() string`

GetVlanInterfaceId returns the VlanInterfaceId field if non-nil, zero value otherwise.

### GetVlanInterfaceIdOk

`func (o *OspfInterfaceConfigOpenApiVO) GetVlanInterfaceIdOk() (*string, bool)`

GetVlanInterfaceIdOk returns a tuple with the VlanInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterfaceId

`func (o *OspfInterfaceConfigOpenApiVO) SetVlanInterfaceId(v string)`

SetVlanInterfaceId sets VlanInterfaceId field to given value.


### GetVlanInterfaceName

`func (o *OspfInterfaceConfigOpenApiVO) GetVlanInterfaceName() string`

GetVlanInterfaceName returns the VlanInterfaceName field if non-nil, zero value otherwise.

### GetVlanInterfaceNameOk

`func (o *OspfInterfaceConfigOpenApiVO) GetVlanInterfaceNameOk() (*string, bool)`

GetVlanInterfaceNameOk returns a tuple with the VlanInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterfaceName

`func (o *OspfInterfaceConfigOpenApiVO) SetVlanInterfaceName(v string)`

SetVlanInterfaceName sets VlanInterfaceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


