# VpnUserServerBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the VPN user. | [optional] 
**Name** | Pointer to **string** | Name of the VPN user. | [optional] 
**VpnType** | Pointer to **int32** | Server Vpn type. 0: L2TP; 1: PPTP; 2: IPSec; 3: OpenVPN; 4: WireGuard; 5: SSL VPN. | [optional] 

## Methods

### NewVpnUserServerBriefVO

`func NewVpnUserServerBriefVO() *VpnUserServerBriefVO`

NewVpnUserServerBriefVO instantiates a new VpnUserServerBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnUserServerBriefVOWithDefaults

`func NewVpnUserServerBriefVOWithDefaults() *VpnUserServerBriefVO`

NewVpnUserServerBriefVOWithDefaults instantiates a new VpnUserServerBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VpnUserServerBriefVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VpnUserServerBriefVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VpnUserServerBriefVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VpnUserServerBriefVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VpnUserServerBriefVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VpnUserServerBriefVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VpnUserServerBriefVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VpnUserServerBriefVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetVpnType

`func (o *VpnUserServerBriefVO) GetVpnType() int32`

GetVpnType returns the VpnType field if non-nil, zero value otherwise.

### GetVpnTypeOk

`func (o *VpnUserServerBriefVO) GetVpnTypeOk() (*int32, bool)`

GetVpnTypeOk returns a tuple with the VpnType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpnType

`func (o *VpnUserServerBriefVO) SetVpnType(v int32)`

SetVpnType sets VpnType field to given value.

### HasVpnType

`func (o *VpnUserServerBriefVO) HasVpnType() bool`

HasVpnType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


