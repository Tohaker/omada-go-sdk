# OuiBasedVlanApOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Ap Rule state. | 
**Name** | **string** | Ap Rule name should contain 1 to 128 characters. | 
**RuleCombine** | [**[]VlanOuiModeOpenApiVO**](VlanOuiModeOpenApiVO.md) | Basic vlan-oui-priority configuration of oui based vlan rule. Cannot be empty. | 
**SsidIdList** | **[]string** | Configured ssid list. | 

## Methods

### NewOuiBasedVlanApOpenApiVO

`func NewOuiBasedVlanApOpenApiVO(enable bool, name string, ruleCombine []VlanOuiModeOpenApiVO, ssidIdList []string, ) *OuiBasedVlanApOpenApiVO`

NewOuiBasedVlanApOpenApiVO instantiates a new OuiBasedVlanApOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanApOpenApiVOWithDefaults

`func NewOuiBasedVlanApOpenApiVOWithDefaults() *OuiBasedVlanApOpenApiVO`

NewOuiBasedVlanApOpenApiVOWithDefaults instantiates a new OuiBasedVlanApOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OuiBasedVlanApOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanApOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanApOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetName

`func (o *OuiBasedVlanApOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanApOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanApOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetRuleCombine

`func (o *OuiBasedVlanApOpenApiVO) GetRuleCombine() []VlanOuiModeOpenApiVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanApOpenApiVO) GetRuleCombineOk() (*[]VlanOuiModeOpenApiVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanApOpenApiVO) SetRuleCombine(v []VlanOuiModeOpenApiVO)`

SetRuleCombine sets RuleCombine field to given value.


### GetSsidIdList

`func (o *OuiBasedVlanApOpenApiVO) GetSsidIdList() []string`

GetSsidIdList returns the SsidIdList field if non-nil, zero value otherwise.

### GetSsidIdListOk

`func (o *OuiBasedVlanApOpenApiVO) GetSsidIdListOk() (*[]string, bool)`

GetSsidIdListOk returns a tuple with the SsidIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidIdList

`func (o *OuiBasedVlanApOpenApiVO) SetSsidIdList(v []string)`

SetSsidIdList sets SsidIdList field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


