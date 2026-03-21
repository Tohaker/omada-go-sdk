# OuiBasedVlanTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Switch Rule state. | 
**Mode** | **int32** | Switch Rule type should be a value as follows: 0: All device port | 
**Name** | **string** | Switch Rule name should contain 1 to 128 characters. | 
**RuleCombine** | [**[]VlanOuiModeOpenApiVO**](VlanOuiModeOpenApiVO.md) | Basic vlan-oui-priority configuration of oui based rule. Cannot be empty. | 

## Methods

### NewOuiBasedVlanTemplateOpenApiVO

`func NewOuiBasedVlanTemplateOpenApiVO(enable bool, mode int32, name string, ruleCombine []VlanOuiModeOpenApiVO, ) *OuiBasedVlanTemplateOpenApiVO`

NewOuiBasedVlanTemplateOpenApiVO instantiates a new OuiBasedVlanTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanTemplateOpenApiVOWithDefaults

`func NewOuiBasedVlanTemplateOpenApiVOWithDefaults() *OuiBasedVlanTemplateOpenApiVO`

NewOuiBasedVlanTemplateOpenApiVOWithDefaults instantiates a new OuiBasedVlanTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OuiBasedVlanTemplateOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanTemplateOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanTemplateOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMode

`func (o *OuiBasedVlanTemplateOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OuiBasedVlanTemplateOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OuiBasedVlanTemplateOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetName

`func (o *OuiBasedVlanTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRuleCombine

`func (o *OuiBasedVlanTemplateOpenApiVO) GetRuleCombine() []VlanOuiModeOpenApiVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanTemplateOpenApiVO) GetRuleCombineOk() (*[]VlanOuiModeOpenApiVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanTemplateOpenApiVO) SetRuleCombine(v []VlanOuiModeOpenApiVO)`

SetRuleCombine sets RuleCombine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


