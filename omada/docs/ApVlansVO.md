# ApVlansVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ipaddr** | Pointer to **string** | IP address. | [optional] 
**LocalVlanId** | Pointer to **int32** | Vlan ID | [optional] 
**LocalVlanNetworkId** | Pointer to **string** | Network ID | [optional] 
**Name** | Pointer to **string** | Network Name | [optional] 
**NativePort** | Pointer to **[]string** | Native Port List | [optional] 
**TagPort** | Pointer to **[]string** | Tag Port List | [optional] 
**UntagPort** | Pointer to **[]string** | Untag Port List | [optional] 

## Methods

### NewApVlansVO

`func NewApVlansVO() *ApVlansVO`

NewApVlansVO instantiates a new ApVlansVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApVlansVOWithDefaults

`func NewApVlansVOWithDefaults() *ApVlansVO`

NewApVlansVOWithDefaults instantiates a new ApVlansVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpaddr

`func (o *ApVlansVO) GetIpaddr() string`

GetIpaddr returns the Ipaddr field if non-nil, zero value otherwise.

### GetIpaddrOk

`func (o *ApVlansVO) GetIpaddrOk() (*string, bool)`

GetIpaddrOk returns a tuple with the Ipaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddr

`func (o *ApVlansVO) SetIpaddr(v string)`

SetIpaddr sets Ipaddr field to given value.

### HasIpaddr

`func (o *ApVlansVO) HasIpaddr() bool`

HasIpaddr returns a boolean if a field has been set.

### GetLocalVlanId

`func (o *ApVlansVO) GetLocalVlanId() int32`

GetLocalVlanId returns the LocalVlanId field if non-nil, zero value otherwise.

### GetLocalVlanIdOk

`func (o *ApVlansVO) GetLocalVlanIdOk() (*int32, bool)`

GetLocalVlanIdOk returns a tuple with the LocalVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanId

`func (o *ApVlansVO) SetLocalVlanId(v int32)`

SetLocalVlanId sets LocalVlanId field to given value.

### HasLocalVlanId

`func (o *ApVlansVO) HasLocalVlanId() bool`

HasLocalVlanId returns a boolean if a field has been set.

### GetLocalVlanNetworkId

`func (o *ApVlansVO) GetLocalVlanNetworkId() string`

GetLocalVlanNetworkId returns the LocalVlanNetworkId field if non-nil, zero value otherwise.

### GetLocalVlanNetworkIdOk

`func (o *ApVlansVO) GetLocalVlanNetworkIdOk() (*string, bool)`

GetLocalVlanNetworkIdOk returns a tuple with the LocalVlanNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVlanNetworkId

`func (o *ApVlansVO) SetLocalVlanNetworkId(v string)`

SetLocalVlanNetworkId sets LocalVlanNetworkId field to given value.

### HasLocalVlanNetworkId

`func (o *ApVlansVO) HasLocalVlanNetworkId() bool`

HasLocalVlanNetworkId returns a boolean if a field has been set.

### GetName

`func (o *ApVlansVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApVlansVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApVlansVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApVlansVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNativePort

`func (o *ApVlansVO) GetNativePort() []string`

GetNativePort returns the NativePort field if non-nil, zero value otherwise.

### GetNativePortOk

`func (o *ApVlansVO) GetNativePortOk() (*[]string, bool)`

GetNativePortOk returns a tuple with the NativePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativePort

`func (o *ApVlansVO) SetNativePort(v []string)`

SetNativePort sets NativePort field to given value.

### HasNativePort

`func (o *ApVlansVO) HasNativePort() bool`

HasNativePort returns a boolean if a field has been set.

### GetTagPort

`func (o *ApVlansVO) GetTagPort() []string`

GetTagPort returns the TagPort field if non-nil, zero value otherwise.

### GetTagPortOk

`func (o *ApVlansVO) GetTagPortOk() (*[]string, bool)`

GetTagPortOk returns a tuple with the TagPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPort

`func (o *ApVlansVO) SetTagPort(v []string)`

SetTagPort sets TagPort field to given value.

### HasTagPort

`func (o *ApVlansVO) HasTagPort() bool`

HasTagPort returns a boolean if a field has been set.

### GetUntagPort

`func (o *ApVlansVO) GetUntagPort() []string`

GetUntagPort returns the UntagPort field if non-nil, zero value otherwise.

### GetUntagPortOk

`func (o *ApVlansVO) GetUntagPortOk() (*[]string, bool)`

GetUntagPortOk returns a tuple with the UntagPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntagPort

`func (o *ApVlansVO) SetUntagPort(v []string)`

SetUntagPort sets UntagPort field to given value.

### HasUntagPort

`func (o *ApVlansVO) HasUntagPort() bool`

HasUntagPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


