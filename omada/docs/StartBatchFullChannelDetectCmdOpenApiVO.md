# StartBatchFullChannelDetectCmdOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EnableChannelUtil** | Pointer to **bool** | Whether to enable channel load detect. | [optional] 
**EnableWifiInterference** | Pointer to **bool** | Whether to enable wifi interference detect. | [optional] 
**MacList** | Pointer to **[]string** | Select the Aps to full channel detect;. | [optional] 
**SelectType** | **string** | Select type of macs. include: include selected aps, exclude: all but exclude selected aps, all: include all aps(Parameter [macList] need input &#39;[]&#39;). | 

## Methods

### NewStartBatchFullChannelDetectCmdOpenApiVO

`func NewStartBatchFullChannelDetectCmdOpenApiVO(selectType string, ) *StartBatchFullChannelDetectCmdOpenApiVO`

NewStartBatchFullChannelDetectCmdOpenApiVO instantiates a new StartBatchFullChannelDetectCmdOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStartBatchFullChannelDetectCmdOpenApiVOWithDefaults

`func NewStartBatchFullChannelDetectCmdOpenApiVOWithDefaults() *StartBatchFullChannelDetectCmdOpenApiVO`

NewStartBatchFullChannelDetectCmdOpenApiVOWithDefaults instantiates a new StartBatchFullChannelDetectCmdOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnableChannelUtil

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetEnableChannelUtil() bool`

GetEnableChannelUtil returns the EnableChannelUtil field if non-nil, zero value otherwise.

### GetEnableChannelUtilOk

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetEnableChannelUtilOk() (*bool, bool)`

GetEnableChannelUtilOk returns a tuple with the EnableChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableChannelUtil

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) SetEnableChannelUtil(v bool)`

SetEnableChannelUtil sets EnableChannelUtil field to given value.

### HasEnableChannelUtil

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) HasEnableChannelUtil() bool`

HasEnableChannelUtil returns a boolean if a field has been set.

### GetEnableWifiInterference

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetEnableWifiInterference() bool`

GetEnableWifiInterference returns the EnableWifiInterference field if non-nil, zero value otherwise.

### GetEnableWifiInterferenceOk

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetEnableWifiInterferenceOk() (*bool, bool)`

GetEnableWifiInterferenceOk returns a tuple with the EnableWifiInterference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableWifiInterference

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) SetEnableWifiInterference(v bool)`

SetEnableWifiInterference sets EnableWifiInterference field to given value.

### HasEnableWifiInterference

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) HasEnableWifiInterference() bool`

HasEnableWifiInterference returns a boolean if a field has been set.

### GetMacList

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) SetMacList(v []string)`

SetMacList sets MacList field to given value.

### HasMacList

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) HasMacList() bool`

HasMacList returns a boolean if a field has been set.

### GetSelectType

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *StartBatchFullChannelDetectCmdOpenApiVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


