# OuiBasedVlanTemplateSwitchQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Switch Rule state. | [optional] 
**Id** | Pointer to **string** | Rule ID | [optional] 
**Mode** | Pointer to **int32** | Switch Rule type should be a value as follows: 0: All device port; 1: Custom device port | [optional] 
**Name** | Pointer to **string** | Switch Rule name should contain 1 to 128 characters. | [optional] 
**RuleCombine** | Pointer to [**[]VlanOuiModeQueryOpenApiVO**](VlanOuiModeQueryOpenApiVO.md) | Basic vlan-oui-priority configuration of oui based vlan rule. | [optional] 

## Methods

### NewOuiBasedVlanTemplateSwitchQueryOpenApiVO

`func NewOuiBasedVlanTemplateSwitchQueryOpenApiVO() *OuiBasedVlanTemplateSwitchQueryOpenApiVO`

NewOuiBasedVlanTemplateSwitchQueryOpenApiVO instantiates a new OuiBasedVlanTemplateSwitchQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanTemplateSwitchQueryOpenApiVOWithDefaults

`func NewOuiBasedVlanTemplateSwitchQueryOpenApiVOWithDefaults() *OuiBasedVlanTemplateSwitchQueryOpenApiVO`

NewOuiBasedVlanTemplateSwitchQueryOpenApiVOWithDefaults instantiates a new OuiBasedVlanTemplateSwitchQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetId

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleCombine

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetRuleCombine() []VlanOuiModeQueryOpenApiVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) GetRuleCombineOk() (*[]VlanOuiModeQueryOpenApiVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) SetRuleCombine(v []VlanOuiModeQueryOpenApiVO)`

SetRuleCombine sets RuleCombine field to given value.

### HasRuleCombine

`func (o *OuiBasedVlanTemplateSwitchQueryOpenApiVO) HasRuleCombine() bool`

HasRuleCombine returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


