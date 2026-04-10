# VpnDefaultValueReqVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **int32** | Parameter [type] should be a value as follows: 0, name; 1, port; 2, tunnelIp. | 
**Usage** | Pointer to **int32** | Parameter [usage] should not be null when parameter [type] is name. Parameter [usage] should be a value as follows: 0, server; 1, client; 2, site-to-site. | [optional] 
**VpnType** | Pointer to **int32** | Parameter [vpnType] should not be null when parameter [type] is port. Parameter [vpnType] should be a value as follows: 3, Open VPN; 4, Wire Guard; 5, SSL VPN. | [optional] 

## Methods

### NewVpnDefaultValueReqVO

`func NewVpnDefaultValueReqVO(type_ int32, ) *VpnDefaultValueReqVO`

NewVpnDefaultValueReqVO instantiates a new VpnDefaultValueReqVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnDefaultValueReqVOWithDefaults

`func NewVpnDefaultValueReqVOWithDefaults() *VpnDefaultValueReqVO`

NewVpnDefaultValueReqVOWithDefaults instantiates a new VpnDefaultValueReqVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VpnDefaultValueReqVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VpnDefaultValueReqVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VpnDefaultValueReqVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUsage

`func (o *VpnDefaultValueReqVO) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *VpnDefaultValueReqVO) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *VpnDefaultValueReqVO) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *VpnDefaultValueReqVO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnDefaultValueReqVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnDefaultValueReqVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnDefaultValueReqVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnDefaultValueReqVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


