# VirtualWanConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Virtual WAN name should contain 1 to 128 characters. Only letters, digits and underscores are allowed. | 
**PhysicalWanId** | **string** | Physical WAN ID. Physical WAN ID can be obtained from &#39;Get internet basic info&#39; interface. Only DSL WAN supports configuring virtual WAN.  | 
**VirtualWanDsl** | Pointer to [**VirtualWanDslOpenApiVO**](VirtualWanDslOpenApiVO.md) |  | [optional] 
**VirtualWanIpv4Setting** | [**VirtualWanIpv4SettingConfigOpenApiVO**](VirtualWanIpv4SettingConfigOpenApiVO.md) |  | 
**WanPortMacSetting** | Pointer to [**VirtualWanMacSettingOpenApiVO**](VirtualWanMacSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewVirtualWanConfigOpenApiVO

`func NewVirtualWanConfigOpenApiVO(name string, physicalWanId string, virtualWanIpv4Setting VirtualWanIpv4SettingConfigOpenApiVO, ) *VirtualWanConfigOpenApiVO`

NewVirtualWanConfigOpenApiVO instantiates a new VirtualWanConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanConfigOpenApiVOWithDefaults

`func NewVirtualWanConfigOpenApiVOWithDefaults() *VirtualWanConfigOpenApiVO`

NewVirtualWanConfigOpenApiVOWithDefaults instantiates a new VirtualWanConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VirtualWanConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualWanConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualWanConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPhysicalWanId

`func (o *VirtualWanConfigOpenApiVO) GetPhysicalWanId() string`

GetPhysicalWanId returns the PhysicalWanId field if non-nil, zero value otherwise.

### GetPhysicalWanIdOk

`func (o *VirtualWanConfigOpenApiVO) GetPhysicalWanIdOk() (*string, bool)`

GetPhysicalWanIdOk returns a tuple with the PhysicalWanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhysicalWanId

`func (o *VirtualWanConfigOpenApiVO) SetPhysicalWanId(v string)`

SetPhysicalWanId sets PhysicalWanId field to given value.


### GetVirtualWanDsl

`func (o *VirtualWanConfigOpenApiVO) GetVirtualWanDsl() VirtualWanDslOpenApiVO`

GetVirtualWanDsl returns the VirtualWanDsl field if non-nil, zero value otherwise.

### GetVirtualWanDslOk

`func (o *VirtualWanConfigOpenApiVO) GetVirtualWanDslOk() (*VirtualWanDslOpenApiVO, bool)`

GetVirtualWanDslOk returns a tuple with the VirtualWanDsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanDsl

`func (o *VirtualWanConfigOpenApiVO) SetVirtualWanDsl(v VirtualWanDslOpenApiVO)`

SetVirtualWanDsl sets VirtualWanDsl field to given value.

### HasVirtualWanDsl

`func (o *VirtualWanConfigOpenApiVO) HasVirtualWanDsl() bool`

HasVirtualWanDsl returns a boolean if a field has been set.

### GetVirtualWanIpv4Setting

`func (o *VirtualWanConfigOpenApiVO) GetVirtualWanIpv4Setting() VirtualWanIpv4SettingConfigOpenApiVO`

GetVirtualWanIpv4Setting returns the VirtualWanIpv4Setting field if non-nil, zero value otherwise.

### GetVirtualWanIpv4SettingOk

`func (o *VirtualWanConfigOpenApiVO) GetVirtualWanIpv4SettingOk() (*VirtualWanIpv4SettingConfigOpenApiVO, bool)`

GetVirtualWanIpv4SettingOk returns a tuple with the VirtualWanIpv4Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualWanIpv4Setting

`func (o *VirtualWanConfigOpenApiVO) SetVirtualWanIpv4Setting(v VirtualWanIpv4SettingConfigOpenApiVO)`

SetVirtualWanIpv4Setting sets VirtualWanIpv4Setting field to given value.


### GetWanPortMacSetting

`func (o *VirtualWanConfigOpenApiVO) GetWanPortMacSetting() VirtualWanMacSettingOpenApiVO`

GetWanPortMacSetting returns the WanPortMacSetting field if non-nil, zero value otherwise.

### GetWanPortMacSettingOk

`func (o *VirtualWanConfigOpenApiVO) GetWanPortMacSettingOk() (*VirtualWanMacSettingOpenApiVO, bool)`

GetWanPortMacSettingOk returns a tuple with the WanPortMacSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortMacSetting

`func (o *VirtualWanConfigOpenApiVO) SetWanPortMacSetting(v VirtualWanMacSettingOpenApiVO)`

SetWanPortMacSetting sets WanPortMacSetting field to given value.

### HasWanPortMacSetting

`func (o *VirtualWanConfigOpenApiVO) HasWanPortMacSetting() bool`

HasWanPortMacSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


