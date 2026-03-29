# OuiBasedVlanSwitchOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceList** | Pointer to [**[]DeviceOuiModeOpenApiVO**](DeviceOuiModeOpenApiVO.md) | When mode is 0, should configure device info. | [optional] 
**Enable** | **bool** | Switch Rule state. | 
**Mode** | **int32** | Switch Rule type. 0:\&quot;All device port\&quot;, 1:\&quot;Custom device port\&quot; | 
**Name** | **string** | Switch Rule name should contain 1 to 128 characters. | 
**RuleCombine** | [**[]VlanOuiModeOpenApiVO**](VlanOuiModeOpenApiVO.md) | Basic vlan-oui-priority configuration of oui based rule. Cannot be empty. | 

## Methods

### NewOuiBasedVlanSwitchOpenApiVO

`func NewOuiBasedVlanSwitchOpenApiVO(enable bool, mode int32, name string, ruleCombine []VlanOuiModeOpenApiVO, ) *OuiBasedVlanSwitchOpenApiVO`

NewOuiBasedVlanSwitchOpenApiVO instantiates a new OuiBasedVlanSwitchOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanSwitchOpenApiVOWithDefaults

`func NewOuiBasedVlanSwitchOpenApiVOWithDefaults() *OuiBasedVlanSwitchOpenApiVO`

NewOuiBasedVlanSwitchOpenApiVOWithDefaults instantiates a new OuiBasedVlanSwitchOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceList

`func (o *OuiBasedVlanSwitchOpenApiVO) GetDeviceList() []DeviceOuiModeOpenApiVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *OuiBasedVlanSwitchOpenApiVO) GetDeviceListOk() (*[]DeviceOuiModeOpenApiVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *OuiBasedVlanSwitchOpenApiVO) SetDeviceList(v []DeviceOuiModeOpenApiVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *OuiBasedVlanSwitchOpenApiVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetEnable

`func (o *OuiBasedVlanSwitchOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanSwitchOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanSwitchOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMode

`func (o *OuiBasedVlanSwitchOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OuiBasedVlanSwitchOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OuiBasedVlanSwitchOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetName

`func (o *OuiBasedVlanSwitchOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanSwitchOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanSwitchOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRuleCombine

`func (o *OuiBasedVlanSwitchOpenApiVO) GetRuleCombine() []VlanOuiModeOpenApiVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanSwitchOpenApiVO) GetRuleCombineOk() (*[]VlanOuiModeOpenApiVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanSwitchOpenApiVO) SetRuleCombine(v []VlanOuiModeOpenApiVO)`

SetRuleCombine sets RuleCombine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


