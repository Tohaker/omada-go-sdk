# OuiBasedVlanApQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Ap Rule state. | [optional] 
**Id** | Pointer to **string** | Rule ID | [optional] 
**Name** | Pointer to **string** | Ap Rule name should contain 1 to 128 characters. | [optional] 
**RuleCombine** | Pointer to [**[]VlanOuiModeQueryOpenApiVO**](VlanOuiModeQueryOpenApiVO.md) | Basic vlan-oui-priority configuration of oui based rule. | [optional] 
**SsidList** | Pointer to [**[]SsidOuiModeOpenApiVO**](SsidOuiModeOpenApiVO.md) | Configured ssid list. | [optional] 

## Methods

### NewOuiBasedVlanApQueryOpenApiVO

`func NewOuiBasedVlanApQueryOpenApiVO() *OuiBasedVlanApQueryOpenApiVO`

NewOuiBasedVlanApQueryOpenApiVO instantiates a new OuiBasedVlanApQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanApQueryOpenApiVOWithDefaults

`func NewOuiBasedVlanApQueryOpenApiVOWithDefaults() *OuiBasedVlanApQueryOpenApiVO`

NewOuiBasedVlanApQueryOpenApiVOWithDefaults instantiates a new OuiBasedVlanApQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OuiBasedVlanApQueryOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanApQueryOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanApQueryOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OuiBasedVlanApQueryOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetId

`func (o *OuiBasedVlanApQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OuiBasedVlanApQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OuiBasedVlanApQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OuiBasedVlanApQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OuiBasedVlanApQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanApQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanApQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OuiBasedVlanApQueryOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRuleCombine

`func (o *OuiBasedVlanApQueryOpenApiVO) GetRuleCombine() []VlanOuiModeQueryOpenApiVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanApQueryOpenApiVO) GetRuleCombineOk() (*[]VlanOuiModeQueryOpenApiVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanApQueryOpenApiVO) SetRuleCombine(v []VlanOuiModeQueryOpenApiVO)`

SetRuleCombine sets RuleCombine field to given value.

### HasRuleCombine

`func (o *OuiBasedVlanApQueryOpenApiVO) HasRuleCombine() bool`

HasRuleCombine returns a boolean if a field has been set.

### GetSsidList

`func (o *OuiBasedVlanApQueryOpenApiVO) GetSsidList() []SsidOuiModeOpenApiVO`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *OuiBasedVlanApQueryOpenApiVO) GetSsidListOk() (*[]SsidOuiModeOpenApiVO, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *OuiBasedVlanApQueryOpenApiVO) SetSsidList(v []SsidOuiModeOpenApiVO)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *OuiBasedVlanApQueryOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


