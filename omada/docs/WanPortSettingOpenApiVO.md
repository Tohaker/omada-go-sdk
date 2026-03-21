# WanPortSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DslSetting** | Pointer to [**WanPortDslSettingOpenApiVO**](WanPortDslSettingOpenApiVO.md) |  | [optional] 
**PortDescription** | Pointer to **string** | Port description should contain 1 to 32 characters. | [optional] 
**PortId** | **string** | Port ID | 
**PortName** | Pointer to **string** | Port name | [optional] 
**WanPortIpv4Setting** | [**WanPortIpv4SettingOpenApiVO**](WanPortIpv4SettingOpenApiVO.md) |  | 
**WanPortIpv6Setting** | [**WanPortIpv6SettingOpenApiVO**](WanPortIpv6SettingOpenApiVO.md) |  | 
**WanPortMacSetting** | [**WanPortMacSettingOpenApiVO**](WanPortMacSettingOpenApiVO.md) |  | 

## Methods

### NewWanPortSettingOpenApiVO

`func NewWanPortSettingOpenApiVO(portId string, wanPortIpv4Setting WanPortIpv4SettingOpenApiVO, wanPortIpv6Setting WanPortIpv6SettingOpenApiVO, wanPortMacSetting WanPortMacSettingOpenApiVO, ) *WanPortSettingOpenApiVO`

NewWanPortSettingOpenApiVO instantiates a new WanPortSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortSettingOpenApiVOWithDefaults

`func NewWanPortSettingOpenApiVOWithDefaults() *WanPortSettingOpenApiVO`

NewWanPortSettingOpenApiVOWithDefaults instantiates a new WanPortSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDslSetting

`func (o *WanPortSettingOpenApiVO) GetDslSetting() WanPortDslSettingOpenApiVO`

GetDslSetting returns the DslSetting field if non-nil, zero value otherwise.

### GetDslSettingOk

`func (o *WanPortSettingOpenApiVO) GetDslSettingOk() (*WanPortDslSettingOpenApiVO, bool)`

GetDslSettingOk returns a tuple with the DslSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDslSetting

`func (o *WanPortSettingOpenApiVO) SetDslSetting(v WanPortDslSettingOpenApiVO)`

SetDslSetting sets DslSetting field to given value.

### HasDslSetting

`func (o *WanPortSettingOpenApiVO) HasDslSetting() bool`

HasDslSetting returns a boolean if a field has been set.

### GetPortDescription

`func (o *WanPortSettingOpenApiVO) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *WanPortSettingOpenApiVO) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *WanPortSettingOpenApiVO) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *WanPortSettingOpenApiVO) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetPortId

`func (o *WanPortSettingOpenApiVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *WanPortSettingOpenApiVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *WanPortSettingOpenApiVO) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetPortName

`func (o *WanPortSettingOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *WanPortSettingOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *WanPortSettingOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *WanPortSettingOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetWanPortIpv4Setting

`func (o *WanPortSettingOpenApiVO) GetWanPortIpv4Setting() WanPortIpv4SettingOpenApiVO`

GetWanPortIpv4Setting returns the WanPortIpv4Setting field if non-nil, zero value otherwise.

### GetWanPortIpv4SettingOk

`func (o *WanPortSettingOpenApiVO) GetWanPortIpv4SettingOk() (*WanPortIpv4SettingOpenApiVO, bool)`

GetWanPortIpv4SettingOk returns a tuple with the WanPortIpv4Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv4Setting

`func (o *WanPortSettingOpenApiVO) SetWanPortIpv4Setting(v WanPortIpv4SettingOpenApiVO)`

SetWanPortIpv4Setting sets WanPortIpv4Setting field to given value.


### GetWanPortIpv6Setting

`func (o *WanPortSettingOpenApiVO) GetWanPortIpv6Setting() WanPortIpv6SettingOpenApiVO`

GetWanPortIpv6Setting returns the WanPortIpv6Setting field if non-nil, zero value otherwise.

### GetWanPortIpv6SettingOk

`func (o *WanPortSettingOpenApiVO) GetWanPortIpv6SettingOk() (*WanPortIpv6SettingOpenApiVO, bool)`

GetWanPortIpv6SettingOk returns a tuple with the WanPortIpv6Setting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIpv6Setting

`func (o *WanPortSettingOpenApiVO) SetWanPortIpv6Setting(v WanPortIpv6SettingOpenApiVO)`

SetWanPortIpv6Setting sets WanPortIpv6Setting field to given value.


### GetWanPortMacSetting

`func (o *WanPortSettingOpenApiVO) GetWanPortMacSetting() WanPortMacSettingOpenApiVO`

GetWanPortMacSetting returns the WanPortMacSetting field if non-nil, zero value otherwise.

### GetWanPortMacSettingOk

`func (o *WanPortSettingOpenApiVO) GetWanPortMacSettingOk() (*WanPortMacSettingOpenApiVO, bool)`

GetWanPortMacSettingOk returns a tuple with the WanPortMacSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortMacSetting

`func (o *WanPortSettingOpenApiVO) SetWanPortMacSetting(v WanPortMacSettingOpenApiVO)`

SetWanPortMacSetting sets WanPortMacSetting field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


