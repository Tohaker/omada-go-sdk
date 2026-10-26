# OuiBasedVlanRuleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceList** | Pointer to [**[]DeviceOuiModeVO**](DeviceOuiModeVO.md) |  | [optional] 
**Enable** | **bool** |  | 
**ExcludeDeviceList** | Pointer to [**[]OuiBasedVlanDeviceInfoVO**](OuiBasedVlanDeviceInfoVO.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Resource** | Pointer to **int32** |  | [optional] 
**RuleCombine** | [**[]VlanOuiModeVO**](VlanOuiModeVO.md) |  | 
**RuleInterface** | Pointer to **[]int32** |  | [optional] 
**SsidIdList** | Pointer to **[]string** |  | [optional] 
**SsidList** | Pointer to [**[]SsidOuiModeVO**](SsidOuiModeVO.md) |  | [optional] 
**Type** | **int32** |  | 

## Methods

### NewOuiBasedVlanRuleVO

`func NewOuiBasedVlanRuleVO(enable bool, name string, ruleCombine []VlanOuiModeVO, type_ int32, ) *OuiBasedVlanRuleVO`

NewOuiBasedVlanRuleVO instantiates a new OuiBasedVlanRuleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanRuleVOWithDefaults

`func NewOuiBasedVlanRuleVOWithDefaults() *OuiBasedVlanRuleVO`

NewOuiBasedVlanRuleVOWithDefaults instantiates a new OuiBasedVlanRuleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceList

`func (o *OuiBasedVlanRuleVO) GetDeviceList() []DeviceOuiModeVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *OuiBasedVlanRuleVO) GetDeviceListOk() (*[]DeviceOuiModeVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *OuiBasedVlanRuleVO) SetDeviceList(v []DeviceOuiModeVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *OuiBasedVlanRuleVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetEnable

`func (o *OuiBasedVlanRuleVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OuiBasedVlanRuleVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OuiBasedVlanRuleVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetExcludeDeviceList

`func (o *OuiBasedVlanRuleVO) GetExcludeDeviceList() []OuiBasedVlanDeviceInfoVO`

GetExcludeDeviceList returns the ExcludeDeviceList field if non-nil, zero value otherwise.

### GetExcludeDeviceListOk

`func (o *OuiBasedVlanRuleVO) GetExcludeDeviceListOk() (*[]OuiBasedVlanDeviceInfoVO, bool)`

GetExcludeDeviceListOk returns a tuple with the ExcludeDeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeDeviceList

`func (o *OuiBasedVlanRuleVO) SetExcludeDeviceList(v []OuiBasedVlanDeviceInfoVO)`

SetExcludeDeviceList sets ExcludeDeviceList field to given value.

### HasExcludeDeviceList

`func (o *OuiBasedVlanRuleVO) HasExcludeDeviceList() bool`

HasExcludeDeviceList returns a boolean if a field has been set.

### GetId

`func (o *OuiBasedVlanRuleVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OuiBasedVlanRuleVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OuiBasedVlanRuleVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OuiBasedVlanRuleVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMode

`func (o *OuiBasedVlanRuleVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OuiBasedVlanRuleVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OuiBasedVlanRuleVO) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *OuiBasedVlanRuleVO) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetName

`func (o *OuiBasedVlanRuleVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanRuleVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanRuleVO) SetName(v string)`

SetName sets Name field to given value.


### GetResource

`func (o *OuiBasedVlanRuleVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *OuiBasedVlanRuleVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *OuiBasedVlanRuleVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *OuiBasedVlanRuleVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRuleCombine

`func (o *OuiBasedVlanRuleVO) GetRuleCombine() []VlanOuiModeVO`

GetRuleCombine returns the RuleCombine field if non-nil, zero value otherwise.

### GetRuleCombineOk

`func (o *OuiBasedVlanRuleVO) GetRuleCombineOk() (*[]VlanOuiModeVO, bool)`

GetRuleCombineOk returns a tuple with the RuleCombine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleCombine

`func (o *OuiBasedVlanRuleVO) SetRuleCombine(v []VlanOuiModeVO)`

SetRuleCombine sets RuleCombine field to given value.


### GetRuleInterface

`func (o *OuiBasedVlanRuleVO) GetRuleInterface() []int32`

GetRuleInterface returns the RuleInterface field if non-nil, zero value otherwise.

### GetRuleInterfaceOk

`func (o *OuiBasedVlanRuleVO) GetRuleInterfaceOk() (*[]int32, bool)`

GetRuleInterfaceOk returns a tuple with the RuleInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleInterface

`func (o *OuiBasedVlanRuleVO) SetRuleInterface(v []int32)`

SetRuleInterface sets RuleInterface field to given value.

### HasRuleInterface

`func (o *OuiBasedVlanRuleVO) HasRuleInterface() bool`

HasRuleInterface returns a boolean if a field has been set.

### GetSsidIdList

`func (o *OuiBasedVlanRuleVO) GetSsidIdList() []string`

GetSsidIdList returns the SsidIdList field if non-nil, zero value otherwise.

### GetSsidIdListOk

`func (o *OuiBasedVlanRuleVO) GetSsidIdListOk() (*[]string, bool)`

GetSsidIdListOk returns a tuple with the SsidIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidIdList

`func (o *OuiBasedVlanRuleVO) SetSsidIdList(v []string)`

SetSsidIdList sets SsidIdList field to given value.

### HasSsidIdList

`func (o *OuiBasedVlanRuleVO) HasSsidIdList() bool`

HasSsidIdList returns a boolean if a field has been set.

### GetSsidList

`func (o *OuiBasedVlanRuleVO) GetSsidList() []SsidOuiModeVO`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *OuiBasedVlanRuleVO) GetSsidListOk() (*[]SsidOuiModeVO, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *OuiBasedVlanRuleVO) SetSsidList(v []SsidOuiModeVO)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *OuiBasedVlanRuleVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.

### GetType

`func (o *OuiBasedVlanRuleVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OuiBasedVlanRuleVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OuiBasedVlanRuleVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


