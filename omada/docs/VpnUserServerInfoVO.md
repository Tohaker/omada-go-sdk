# VpnUserServerInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the VPN. | [optional] 
**Name** | Pointer to **string** | VPN name. | [optional] 
**VpnType** | Pointer to **int32** | Server Vpn type. 0: L2TP; 1: PPTP; 3: OpenVPN; 5: SSL VPN. | [optional] 

## Methods

### NewVpnUserServerInfoVO

`func NewVpnUserServerInfoVO() *VpnUserServerInfoVO`

NewVpnUserServerInfoVO instantiates a new VpnUserServerInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserServerInfoVOWithDefaults

`func NewVpnUserServerInfoVOWithDefaults() *VpnUserServerInfoVO`

NewVpnUserServerInfoVOWithDefaults instantiates a new VpnUserServerInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpnUserServerInfoVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnUserServerInfoVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnUserServerInfoVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnUserServerInfoVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VpnUserServerInfoVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnUserServerInfoVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnUserServerInfoVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnUserServerInfoVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnUserServerInfoVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnUserServerInfoVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnUserServerInfoVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnUserServerInfoVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


