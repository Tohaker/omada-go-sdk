# VirtualWanInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DslSetting** | Pointer to [**VirtualWanDslOpenApiVO**](VirtualWanDslOpenApiVO.md) |  | [optional] 
**Id** | Pointer to **string** | Virtual WAN ID. | [optional] 
**Name** | Pointer to **string** | Virtual WAN name.  | [optional] 
**PhysicalWanId** | Pointer to **string** | Physical WAN ID. | [optional] 
**PhysicalWanName** | Pointer to **string** | Physical WAN name. | [optional] 
**PhysicalWanPortId** | Pointer to **int32** | Physical WAN port ID. | [optional] 
**Status** | Pointer to **bool** | Virtual WAN status. | [optional] 
**WanPortIpv4Setting** | Pointer to [**VirtualWanIpv4SettingInfoOpenApiVO**](VirtualWanIpv4SettingInfoOpenApiVO.md) |  | [optional] 

## Methods

### NewVirtualWanInfoOpenApiVO

`func NewVirtualWanInfoOpenApiVO() *VirtualWanInfoOpenApiVO`

NewVirtualWanInfoOpenApiVO instantiates a new VirtualWanInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanInfoOpenApiVOWithDefaults

`func NewVirtualWanInfoOpenApiVOWithDefaults() *VirtualWanInfoOpenApiVO`

NewVirtualWanInfoOpenApiVOWithDefaults instantiates a new VirtualWanInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDslSetting

`func (o *VirtualWanInfoOpenApiVO) GetDslSetting() VirtualWanDslOpenApiVO`

GetDslSetting returns the DslSetting field if non-nil, zero value otherwise.

### GetDslSettingOk

`func (o *VirtualWanInfoOpenApiVO) GetDslSettingOk() (*VirtualWanDslOpenApiVO, bool)`

GetDslSettingOk returns a tuple with the DslSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslSetting

`func (o *VirtualWanInfoOpenApiVO) SetDslSetting(v VirtualWanDslOpenApiVO)`

SetDslSetting sets DslSetting field to given value.

### HasDslSetting

`func (o *VirtualWanInfoOpenApiVO) HasDslSetting() bool`

HasDslSetting returns a boolean if a field has been set.

### GetId

`func (o *VirtualWanInfoOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualWanInfoOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualWanInfoOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualWanInfoOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VirtualWanInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualWanInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualWanInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualWanInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPhysicalWanId

`func (o *VirtualWanInfoOpenApiVO) GetPhysicalWanId() string`

GetPhysicalWanId returns the PhysicalWanId field if non-nil, zero value otherwise.

### GetPhysicalWanIdOk

`func (o *VirtualWanInfoOpenApiVO) GetPhysicalWanIdOk() (*string, bool)`

GetPhysicalWanIdOk returns a tuple with the PhysicalWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalWanId

`func (o *VirtualWanInfoOpenApiVO) SetPhysicalWanId(v string)`

SetPhysicalWanId sets PhysicalWanId field to given value.

### HasPhysicalWanId

`func (o *VirtualWanInfoOpenApiVO) HasPhysicalWanId() bool`

HasPhysicalWanId returns a boolean if a field has been set.

### GetPhysicalWanName

`func (o *VirtualWanInfoOpenApiVO) GetPhysicalWanName() string`

GetPhysicalWanName returns the PhysicalWanName field if non-nil, zero value otherwise.

### GetPhysicalWanNameOk

`func (o *VirtualWanInfoOpenApiVO) GetPhysicalWanNameOk() (*string, bool)`

GetPhysicalWanNameOk returns a tuple with the PhysicalWanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalWanName

`func (o *VirtualWanInfoOpenApiVO) SetPhysicalWanName(v string)`

SetPhysicalWanName sets PhysicalWanName field to given value.

### HasPhysicalWanName

`func (o *VirtualWanInfoOpenApiVO) HasPhysicalWanName() bool`

HasPhysicalWanName returns a boolean if a field has been set.

### GetPhysicalWanPortId

`func (o *VirtualWanInfoOpenApiVO) GetPhysicalWanPortId() int32`

GetPhysicalWanPortId returns the PhysicalWanPortId field if non-nil, zero value otherwise.

### GetPhysicalWanPortIdOk

`func (o *VirtualWanInfoOpenApiVO) GetPhysicalWanPortIdOk() (*int32, bool)`

GetPhysicalWanPortIdOk returns a tuple with the PhysicalWanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalWanPortId

`func (o *VirtualWanInfoOpenApiVO) SetPhysicalWanPortId(v int32)`

SetPhysicalWanPortId sets PhysicalWanPortId field to given value.

### HasPhysicalWanPortId

`func (o *VirtualWanInfoOpenApiVO) HasPhysicalWanPortId() bool`

HasPhysicalWanPortId returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualWanInfoOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualWanInfoOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualWanInfoOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualWanInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetWanPortIpv4Setting

`func (o *VirtualWanInfoOpenApiVO) GetWanPortIpv4Setting() VirtualWanIpv4SettingInfoOpenApiVO`

GetWanPortIpv4Setting returns the WanPortIpv4Setting field if non-nil, zero value otherwise.

### GetWanPortIpv4SettingOk

`func (o *VirtualWanInfoOpenApiVO) GetWanPortIpv4SettingOk() (*VirtualWanIpv4SettingInfoOpenApiVO, bool)`

GetWanPortIpv4SettingOk returns a tuple with the WanPortIpv4Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv4Setting

`func (o *VirtualWanInfoOpenApiVO) SetWanPortIpv4Setting(v VirtualWanIpv4SettingInfoOpenApiVO)`

SetWanPortIpv4Setting sets WanPortIpv4Setting field to given value.

### HasWanPortIpv4Setting

`func (o *VirtualWanInfoOpenApiVO) HasWanPortIpv4Setting() bool`

HasWanPortIpv4Setting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


