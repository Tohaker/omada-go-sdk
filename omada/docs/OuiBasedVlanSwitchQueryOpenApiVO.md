# OuiBasedVlanSwitchQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceList** | Pointer to [**[]DeviceOuiModeQueryOpenApiVO**](DeviceOuiModeQueryOpenApiVO.md) | When mode is 0, Show configured device info. | [optional] 
**Enable** | Pointer to **bool** | Switch Rule state. | [optional] 
**Id** | Pointer to **string** | Rule ID | [optional] 
**Mode** | Pointer to **int32** | Switch Rule type should be a value as follows: 0: All device port; 1: Custom device port | [optional] 
**Name** | Pointer to **string** | Switch Rule name should contain 1 to 128 characters. | [optional] 
**RuleCombine** | Pointer to [**[]VlanOuiModeQueryOpenApiVO**](VlanOuiModeQueryOpenApiVO.md) | Basic vlan-oui-priority configuration of oui based vlan rule. | [optional] 

## Methods

### NewOuiBasedVlanSwitchQueryOpenApiVO

`func NewOuiBasedVlanSwitchQueryOpenApiVO() *OuiBasedVlanSwitchQueryOpenApiVO`

NewOuiBasedVlanSwitchQueryOpenApiVO instantiates a new OuiBasedVlanSwitchQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanSwitchQueryOpenApiVOWithDefaults

`func NewOuiBasedVlanSwitchQueryOpenApiVOWithDefaults() *OuiBasedVlanSwitchQueryOpenApiVO`

NewOuiBasedVlanSwitchQueryOpenApiVOWithDefaults instantiates a new OuiBasedVlanSwitchQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceList

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetDeviceList() []DeviceOuiModeQueryOpenApiVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetDeviceListOk() (*[]DeviceOuiModeQueryOpenApiVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) SetDeviceList(v []DeviceOuiModeQueryOpenApiVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetEnable

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetId

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleCombine

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetRuleCombine() []VlanOuiModeQueryOpenApiVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) GetRuleCombineOk() (*[]VlanOuiModeQueryOpenApiVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) SetRuleCombine(v []VlanOuiModeQueryOpenApiVO)`

SetRuleCombine sets RuleCombine field to given value.

### HasRuleCombine

`func (o *OuiBasedVlanSwitchQueryOpenApiVO) HasRuleCombine() bool`

HasRuleCombine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


