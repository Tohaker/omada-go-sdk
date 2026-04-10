# VpnValueAvailableVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | VPN ID. | [optional] 
**Name** | Pointer to **string** | VPN name. Parameter [name] should not be null when [type] is name. | [optional] 
**Port** | Pointer to **int32** | Parameter [port] should not be null when [type] is port. | [optional] 
**ServiceType** | Pointer to **int32** | Service type of the Open VPN should be a value as follows: 0: UDP; 1: TCP. | [optional] 
**Type** | **int32** | Parameter [type] should be a value as follows: 0, name; 1, port. | 
**Usage** | Pointer to **int32** | Parameter [usage] should not be null when [type] is name. Parameter [usage] should be a value as follows: 0, server; 1, client; 2: site-to-site. | [optional] 
**VpnType** | Pointer to **int32** | Vpn type should be a value as follows: 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | [optional] 
**Wan** | Pointer to **string** | WAN of the VPN. | [optional] 

## Methods

### NewVpnValueAvailableVO

`func NewVpnValueAvailableVO(type_ int32, ) *VpnValueAvailableVO`

NewVpnValueAvailableVO instantiates a new VpnValueAvailableVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnValueAvailableVOWithDefaults

`func NewVpnValueAvailableVOWithDefaults() *VpnValueAvailableVO`

NewVpnValueAvailableVOWithDefaults instantiates a new VpnValueAvailableVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpnValueAvailableVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnValueAvailableVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnValueAvailableVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnValueAvailableVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VpnValueAvailableVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnValueAvailableVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnValueAvailableVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnValueAvailableVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *VpnValueAvailableVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VpnValueAvailableVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VpnValueAvailableVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *VpnValueAvailableVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetServiceType

`func (o *VpnValueAvailableVO) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *VpnValueAvailableVO) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *VpnValueAvailableVO) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *VpnValueAvailableVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetType

`func (o *VpnValueAvailableVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VpnValueAvailableVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VpnValueAvailableVO) SetType(v int32)`

SetType sets Type field to given value.


### GetUsage

`func (o *VpnValueAvailableVO) GetUsage() int32`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *VpnValueAvailableVO) GetUsageOk() (*int32, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *VpnValueAvailableVO) SetUsage(v int32)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *VpnValueAvailableVO) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnValueAvailableVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnValueAvailableVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnValueAvailableVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnValueAvailableVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.

### GetWan

`func (o *VpnValueAvailableVO) GetWan() string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *VpnValueAvailableVO) GetWanOk() (*string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *VpnValueAvailableVO) SetWan(v string)`

SetWan sets Wan field to given value.

### HasWan

`func (o *VpnValueAvailableVO) HasWan() bool`

HasWan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


