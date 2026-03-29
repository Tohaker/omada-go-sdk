# OspfInterfaceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthenticationType** | **int32** | Authentication Type, it should be a value as follows: 0: None, 1: Simple, 2: MD5. | 
**Cost** | **int32** | The link cost. OSPF uses this value in computing shortest paths. cost should be within the range of 1 to 65535. | 
**DeadInterval** | Pointer to **int32** | The dead interval for the specified interface in seconds. This specifies how long a router will wait to see a neighbor router&#39;s Hello packets before declaring that the router is down. This parameter must be the same for all routers attached to a network. It should be within the range of 1-65535 seconds and the default is 40. | [optional] 
**DeviceName** | **string** | Device Name | 
**HelloInterval** | **int32** | The hello interval for the specified interface in seconds. This parameter must be the same for all routers attached to a network. It should be within the range of 1-65535 seconds and the default is 10 seconds. | 
**Id** | Pointer to **string** | OSPF Interface ID | [optional] 
**IsStack** | Pointer to **bool** | Indicates whether the device is a stack member device. | [optional] 
**Mac** | **string** | Device Mac | 
**Md5Key** | Pointer to **string** | Displays the key used for md5 authentication, its value should be within the range of 1-16. | [optional] 
**Md5KeyId** | Pointer to **int32** | Displays the key ID used for md5 authentication, its value should be within the range of 1-255. | [optional] 
**NetworkType** | **int32** | Network Type, it should be a value as follows: 0: Broadcast, 1: Non-Broadcast, 2: Point-to-Multipoint, 3: Point-to-Point. The default network type for Ethernet interfaces is broadcast. | 
**SimpleKey** | Pointer to **string** | Displays the key used for simple authentication, its value should be within the range of 1-8. | [optional] 
**StackId** | Pointer to **string** | Stack ID, used for backend verification, there is no need to transmit values, and it has value when the mac is master. | [optional] 
**VlanId** | **int32** | Vlan ID should be within the range of 1–4094. | 
**VlanInterfaceId** | **string** | Vlan Interface ID | 
**VlanInterfaceName** | **string** | Vlan Interface Name | 

## Methods

### NewOspfInterfaceOpenApiVO

`func NewOspfInterfaceOpenApiVO(authenticationType int32, cost int32, deviceName string, helloInterval int32, mac string, networkType int32, vlanId int32, vlanInterfaceId string, vlanInterfaceName string, ) *OspfInterfaceOpenApiVO`

NewOspfInterfaceOpenApiVO instantiates a new OspfInterfaceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfInterfaceOpenApiVOWithDefaults

`func NewOspfInterfaceOpenApiVOWithDefaults() *OspfInterfaceOpenApiVO`

NewOspfInterfaceOpenApiVOWithDefaults instantiates a new OspfInterfaceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticationType

`func (o *OspfInterfaceOpenApiVO) GetAuthenticationType() int32`

GetAuthenticationType returns the AuthenticationType field if non-nil, zero value otherwise.

### GetAuthenticationTypeOk

`func (o *OspfInterfaceOpenApiVO) GetAuthenticationTypeOk() (*int32, bool)`

GetAuthenticationTypeOk returns a tuple with the AuthenticationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationType

`func (o *OspfInterfaceOpenApiVO) SetAuthenticationType(v int32)`

SetAuthenticationType sets AuthenticationType field to given value.


### GetCost

`func (o *OspfInterfaceOpenApiVO) GetCost() int32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *OspfInterfaceOpenApiVO) GetCostOk() (*int32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *OspfInterfaceOpenApiVO) SetCost(v int32)`

SetCost sets Cost field to given value.


### GetDeadInterval

`func (o *OspfInterfaceOpenApiVO) GetDeadInterval() int32`

GetDeadInterval returns the DeadInterval field if non-nil, zero value otherwise.

### GetDeadIntervalOk

`func (o *OspfInterfaceOpenApiVO) GetDeadIntervalOk() (*int32, bool)`

GetDeadIntervalOk returns a tuple with the DeadInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeadInterval

`func (o *OspfInterfaceOpenApiVO) SetDeadInterval(v int32)`

SetDeadInterval sets DeadInterval field to given value.

### HasDeadInterval

`func (o *OspfInterfaceOpenApiVO) HasDeadInterval() bool`

HasDeadInterval returns a boolean if a field has been set.

### GetDeviceName

`func (o *OspfInterfaceOpenApiVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *OspfInterfaceOpenApiVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *OspfInterfaceOpenApiVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.


### GetHelloInterval

`func (o *OspfInterfaceOpenApiVO) GetHelloInterval() int32`

GetHelloInterval returns the HelloInterval field if non-nil, zero value otherwise.

### GetHelloIntervalOk

`func (o *OspfInterfaceOpenApiVO) GetHelloIntervalOk() (*int32, bool)`

GetHelloIntervalOk returns a tuple with the HelloInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelloInterval

`func (o *OspfInterfaceOpenApiVO) SetHelloInterval(v int32)`

SetHelloInterval sets HelloInterval field to given value.


### GetId

`func (o *OspfInterfaceOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OspfInterfaceOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OspfInterfaceOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OspfInterfaceOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsStack

`func (o *OspfInterfaceOpenApiVO) GetIsStack() bool`

GetIsStack returns the IsStack field if non-nil, zero value otherwise.

### GetIsStackOk

`func (o *OspfInterfaceOpenApiVO) GetIsStackOk() (*bool, bool)`

GetIsStackOk returns a tuple with the IsStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStack

`func (o *OspfInterfaceOpenApiVO) SetIsStack(v bool)`

SetIsStack sets IsStack field to given value.

### HasIsStack

`func (o *OspfInterfaceOpenApiVO) HasIsStack() bool`

HasIsStack returns a boolean if a field has been set.

### GetMac

`func (o *OspfInterfaceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OspfInterfaceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OspfInterfaceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetMd5Key

`func (o *OspfInterfaceOpenApiVO) GetMd5Key() string`

GetMd5Key returns the Md5Key field if non-nil, zero value otherwise.

### GetMd5KeyOk

`func (o *OspfInterfaceOpenApiVO) GetMd5KeyOk() (*string, bool)`

GetMd5KeyOk returns a tuple with the Md5Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5Key

`func (o *OspfInterfaceOpenApiVO) SetMd5Key(v string)`

SetMd5Key sets Md5Key field to given value.

### HasMd5Key

`func (o *OspfInterfaceOpenApiVO) HasMd5Key() bool`

HasMd5Key returns a boolean if a field has been set.

### GetMd5KeyId

`func (o *OspfInterfaceOpenApiVO) GetMd5KeyId() int32`

GetMd5KeyId returns the Md5KeyId field if non-nil, zero value otherwise.

### GetMd5KeyIdOk

`func (o *OspfInterfaceOpenApiVO) GetMd5KeyIdOk() (*int32, bool)`

GetMd5KeyIdOk returns a tuple with the Md5KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMd5KeyId

`func (o *OspfInterfaceOpenApiVO) SetMd5KeyId(v int32)`

SetMd5KeyId sets Md5KeyId field to given value.

### HasMd5KeyId

`func (o *OspfInterfaceOpenApiVO) HasMd5KeyId() bool`

HasMd5KeyId returns a boolean if a field has been set.

### GetNetworkType

`func (o *OspfInterfaceOpenApiVO) GetNetworkType() int32`

GetNetworkType returns the NetworkType field if non-nil, zero value otherwise.

### GetNetworkTypeOk

`func (o *OspfInterfaceOpenApiVO) GetNetworkTypeOk() (*int32, bool)`

GetNetworkTypeOk returns a tuple with the NetworkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkType

`func (o *OspfInterfaceOpenApiVO) SetNetworkType(v int32)`

SetNetworkType sets NetworkType field to given value.


### GetSimpleKey

`func (o *OspfInterfaceOpenApiVO) GetSimpleKey() string`

GetSimpleKey returns the SimpleKey field if non-nil, zero value otherwise.

### GetSimpleKeyOk

`func (o *OspfInterfaceOpenApiVO) GetSimpleKeyOk() (*string, bool)`

GetSimpleKeyOk returns a tuple with the SimpleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimpleKey

`func (o *OspfInterfaceOpenApiVO) SetSimpleKey(v string)`

SetSimpleKey sets SimpleKey field to given value.

### HasSimpleKey

`func (o *OspfInterfaceOpenApiVO) HasSimpleKey() bool`

HasSimpleKey returns a boolean if a field has been set.

### GetStackId

`func (o *OspfInterfaceOpenApiVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OspfInterfaceOpenApiVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OspfInterfaceOpenApiVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OspfInterfaceOpenApiVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetVlanId

`func (o *OspfInterfaceOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *OspfInterfaceOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *OspfInterfaceOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.


### GetVlanInterfaceId

`func (o *OspfInterfaceOpenApiVO) GetVlanInterfaceId() string`

GetVlanInterfaceId returns the VlanInterfaceId field if non-nil, zero value otherwise.

### GetVlanInterfaceIdOk

`func (o *OspfInterfaceOpenApiVO) GetVlanInterfaceIdOk() (*string, bool)`

GetVlanInterfaceIdOk returns a tuple with the VlanInterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterfaceId

`func (o *OspfInterfaceOpenApiVO) SetVlanInterfaceId(v string)`

SetVlanInterfaceId sets VlanInterfaceId field to given value.


### GetVlanInterfaceName

`func (o *OspfInterfaceOpenApiVO) GetVlanInterfaceName() string`

GetVlanInterfaceName returns the VlanInterfaceName field if non-nil, zero value otherwise.

### GetVlanInterfaceNameOk

`func (o *OspfInterfaceOpenApiVO) GetVlanInterfaceNameOk() (*string, bool)`

GetVlanInterfaceNameOk returns a tuple with the VlanInterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterfaceName

`func (o *OspfInterfaceOpenApiVO) SetVlanInterfaceName(v string)`

SetVlanInterfaceName sets VlanInterfaceName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


