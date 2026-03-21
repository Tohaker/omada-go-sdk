# RFPlanningHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterIndex** | Pointer to **int32** | Index after WLAN Optimization, between 0 and 100. | [optional] 
**ApNum** | Pointer to **int32** | The number of EAPs. | [optional] 
**BandDeployEnable** | Pointer to **bool** | Whether to enable band deployment. | [optional] 
**BeforeIndex** | Pointer to **int32** | Index before WLAN Optimization, between 0 and 100. | [optional] 
**ChannelDeployEnable** | Pointer to **bool** | Whether to enable channel deployment. | [optional] 
**ChannelWidthDeployEnable** | Pointer to **bool** | Whether to enable channel width deployment. | [optional] 
**CurrentConfig** | Pointer to **int32** | 0: This entry does not correspond to the current configuration. 1: The recommended values are the current configuration. 2: The previous/origin values are the current configuration. | [optional] 
**FailedMacs2g** | Pointer to **string** |  | [optional] 
**FailedMacs5g** | Pointer to **string** |  | [optional] 
**FailedMacs6g** | Pointer to **string** |  | [optional] 
**IsAppliedSuccess** | Pointer to **bool** | True: Succeeded to apply settings recommended by the algorithm. False or null: Failed to apply settings recommended by the algorithm.  | [optional] 
**IsScanSuccess** | Pointer to **bool** | True: Succeeded to obtain device scan results. False: Failed to obtain device scan results. | [optional] 
**Length** | Pointer to **int32** | Parameter [length] means the duration of the WLAN Optimization in seconds. | [optional] 
**Mode** | Pointer to **int32** | 0: by WLAN Optimization schedule. 1: by one-click WLAN Optimization. | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**MsgType** | Pointer to **int32** |  | [optional] 
**PlanningHistroyId** | Pointer to **string** | Planning histroy ID. | [optional] 
**PowerAdjustEnable** | Pointer to **bool** | Whether to enable power adjustment. | [optional] 
**Time** | Pointer to **int64** | Timestamp corresponding to the start of optimization. | [optional] 

## Methods

### NewRFPlanningHistory

`func NewRFPlanningHistory() *RFPlanningHistory`

NewRFPlanningHistory instantiates a new RFPlanningHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFPlanningHistoryWithDefaults

`func NewRFPlanningHistoryWithDefaults() *RFPlanningHistory`

NewRFPlanningHistoryWithDefaults instantiates a new RFPlanningHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterIndex

`func (o *RFPlanningHistory) GetAfterIndex() int32`

GetAfterIndex returns the AfterIndex field if non-nil, zero value otherwise.

### GetAfterIndexOk

`func (o *RFPlanningHistory) GetAfterIndexOk() (*int32, bool)`

GetAfterIndexOk returns a tuple with the AfterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterIndex

`func (o *RFPlanningHistory) SetAfterIndex(v int32)`

SetAfterIndex sets AfterIndex field to given value.

### HasAfterIndex

`func (o *RFPlanningHistory) HasAfterIndex() bool`

HasAfterIndex returns a boolean if a field has been set.

### GetApNum

`func (o *RFPlanningHistory) GetApNum() int32`

GetApNum returns the ApNum field if non-nil, zero value otherwise.

### GetApNumOk

`func (o *RFPlanningHistory) GetApNumOk() (*int32, bool)`

GetApNumOk returns a tuple with the ApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApNum

`func (o *RFPlanningHistory) SetApNum(v int32)`

SetApNum sets ApNum field to given value.

### HasApNum

`func (o *RFPlanningHistory) HasApNum() bool`

HasApNum returns a boolean if a field has been set.

### GetBandDeployEnable

`func (o *RFPlanningHistory) GetBandDeployEnable() bool`

GetBandDeployEnable returns the BandDeployEnable field if non-nil, zero value otherwise.

### GetBandDeployEnableOk

`func (o *RFPlanningHistory) GetBandDeployEnableOk() (*bool, bool)`

GetBandDeployEnableOk returns a tuple with the BandDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandDeployEnable

`func (o *RFPlanningHistory) SetBandDeployEnable(v bool)`

SetBandDeployEnable sets BandDeployEnable field to given value.

### HasBandDeployEnable

`func (o *RFPlanningHistory) HasBandDeployEnable() bool`

HasBandDeployEnable returns a boolean if a field has been set.

### GetBeforeIndex

`func (o *RFPlanningHistory) GetBeforeIndex() int32`

GetBeforeIndex returns the BeforeIndex field if non-nil, zero value otherwise.

### GetBeforeIndexOk

`func (o *RFPlanningHistory) GetBeforeIndexOk() (*int32, bool)`

GetBeforeIndexOk returns a tuple with the BeforeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeIndex

`func (o *RFPlanningHistory) SetBeforeIndex(v int32)`

SetBeforeIndex sets BeforeIndex field to given value.

### HasBeforeIndex

`func (o *RFPlanningHistory) HasBeforeIndex() bool`

HasBeforeIndex returns a boolean if a field has been set.

### GetChannelDeployEnable

`func (o *RFPlanningHistory) GetChannelDeployEnable() bool`

GetChannelDeployEnable returns the ChannelDeployEnable field if non-nil, zero value otherwise.

### GetChannelDeployEnableOk

`func (o *RFPlanningHistory) GetChannelDeployEnableOk() (*bool, bool)`

GetChannelDeployEnableOk returns a tuple with the ChannelDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelDeployEnable

`func (o *RFPlanningHistory) SetChannelDeployEnable(v bool)`

SetChannelDeployEnable sets ChannelDeployEnable field to given value.

### HasChannelDeployEnable

`func (o *RFPlanningHistory) HasChannelDeployEnable() bool`

HasChannelDeployEnable returns a boolean if a field has been set.

### GetChannelWidthDeployEnable

`func (o *RFPlanningHistory) GetChannelWidthDeployEnable() bool`

GetChannelWidthDeployEnable returns the ChannelWidthDeployEnable field if non-nil, zero value otherwise.

### GetChannelWidthDeployEnableOk

`func (o *RFPlanningHistory) GetChannelWidthDeployEnableOk() (*bool, bool)`

GetChannelWidthDeployEnableOk returns a tuple with the ChannelWidthDeployEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelWidthDeployEnable

`func (o *RFPlanningHistory) SetChannelWidthDeployEnable(v bool)`

SetChannelWidthDeployEnable sets ChannelWidthDeployEnable field to given value.

### HasChannelWidthDeployEnable

`func (o *RFPlanningHistory) HasChannelWidthDeployEnable() bool`

HasChannelWidthDeployEnable returns a boolean if a field has been set.

### GetCurrentConfig

`func (o *RFPlanningHistory) GetCurrentConfig() int32`

GetCurrentConfig returns the CurrentConfig field if non-nil, zero value otherwise.

### GetCurrentConfigOk

`func (o *RFPlanningHistory) GetCurrentConfigOk() (*int32, bool)`

GetCurrentConfigOk returns a tuple with the CurrentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentConfig

`func (o *RFPlanningHistory) SetCurrentConfig(v int32)`

SetCurrentConfig sets CurrentConfig field to given value.

### HasCurrentConfig

`func (o *RFPlanningHistory) HasCurrentConfig() bool`

HasCurrentConfig returns a boolean if a field has been set.

### GetFailedMacs2g

`func (o *RFPlanningHistory) GetFailedMacs2g() string`

GetFailedMacs2g returns the FailedMacs2g field if non-nil, zero value otherwise.

### GetFailedMacs2gOk

`func (o *RFPlanningHistory) GetFailedMacs2gOk() (*string, bool)`

GetFailedMacs2gOk returns a tuple with the FailedMacs2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMacs2g

`func (o *RFPlanningHistory) SetFailedMacs2g(v string)`

SetFailedMacs2g sets FailedMacs2g field to given value.

### HasFailedMacs2g

`func (o *RFPlanningHistory) HasFailedMacs2g() bool`

HasFailedMacs2g returns a boolean if a field has been set.

### GetFailedMacs5g

`func (o *RFPlanningHistory) GetFailedMacs5g() string`

GetFailedMacs5g returns the FailedMacs5g field if non-nil, zero value otherwise.

### GetFailedMacs5gOk

`func (o *RFPlanningHistory) GetFailedMacs5gOk() (*string, bool)`

GetFailedMacs5gOk returns a tuple with the FailedMacs5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMacs5g

`func (o *RFPlanningHistory) SetFailedMacs5g(v string)`

SetFailedMacs5g sets FailedMacs5g field to given value.

### HasFailedMacs5g

`func (o *RFPlanningHistory) HasFailedMacs5g() bool`

HasFailedMacs5g returns a boolean if a field has been set.

### GetFailedMacs6g

`func (o *RFPlanningHistory) GetFailedMacs6g() string`

GetFailedMacs6g returns the FailedMacs6g field if non-nil, zero value otherwise.

### GetFailedMacs6gOk

`func (o *RFPlanningHistory) GetFailedMacs6gOk() (*string, bool)`

GetFailedMacs6gOk returns a tuple with the FailedMacs6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedMacs6g

`func (o *RFPlanningHistory) SetFailedMacs6g(v string)`

SetFailedMacs6g sets FailedMacs6g field to given value.

### HasFailedMacs6g

`func (o *RFPlanningHistory) HasFailedMacs6g() bool`

HasFailedMacs6g returns a boolean if a field has been set.

### GetIsAppliedSuccess

`func (o *RFPlanningHistory) GetIsAppliedSuccess() bool`

GetIsAppliedSuccess returns the IsAppliedSuccess field if non-nil, zero value otherwise.

### GetIsAppliedSuccessOk

`func (o *RFPlanningHistory) GetIsAppliedSuccessOk() (*bool, bool)`

GetIsAppliedSuccessOk returns a tuple with the IsAppliedSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAppliedSuccess

`func (o *RFPlanningHistory) SetIsAppliedSuccess(v bool)`

SetIsAppliedSuccess sets IsAppliedSuccess field to given value.

### HasIsAppliedSuccess

`func (o *RFPlanningHistory) HasIsAppliedSuccess() bool`

HasIsAppliedSuccess returns a boolean if a field has been set.

### GetIsScanSuccess

`func (o *RFPlanningHistory) GetIsScanSuccess() bool`

GetIsScanSuccess returns the IsScanSuccess field if non-nil, zero value otherwise.

### GetIsScanSuccessOk

`func (o *RFPlanningHistory) GetIsScanSuccessOk() (*bool, bool)`

GetIsScanSuccessOk returns a tuple with the IsScanSuccess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsScanSuccess

`func (o *RFPlanningHistory) SetIsScanSuccess(v bool)`

SetIsScanSuccess sets IsScanSuccess field to given value.

### HasIsScanSuccess

`func (o *RFPlanningHistory) HasIsScanSuccess() bool`

HasIsScanSuccess returns a boolean if a field has been set.

### GetLength

`func (o *RFPlanningHistory) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *RFPlanningHistory) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *RFPlanningHistory) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *RFPlanningHistory) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMode

`func (o *RFPlanningHistory) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RFPlanningHistory) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RFPlanningHistory) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RFPlanningHistory) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetMsg

`func (o *RFPlanningHistory) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *RFPlanningHistory) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *RFPlanningHistory) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *RFPlanningHistory) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetMsgType

`func (o *RFPlanningHistory) GetMsgType() int32`

GetMsgType returns the MsgType field if non-nil, zero value otherwise.

### GetMsgTypeOk

`func (o *RFPlanningHistory) GetMsgTypeOk() (*int32, bool)`

GetMsgTypeOk returns a tuple with the MsgType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsgType

`func (o *RFPlanningHistory) SetMsgType(v int32)`

SetMsgType sets MsgType field to given value.

### HasMsgType

`func (o *RFPlanningHistory) HasMsgType() bool`

HasMsgType returns a boolean if a field has been set.

### GetPlanningHistroyId

`func (o *RFPlanningHistory) GetPlanningHistroyId() string`

GetPlanningHistroyId returns the PlanningHistroyId field if non-nil, zero value otherwise.

### GetPlanningHistroyIdOk

`func (o *RFPlanningHistory) GetPlanningHistroyIdOk() (*string, bool)`

GetPlanningHistroyIdOk returns a tuple with the PlanningHistroyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningHistroyId

`func (o *RFPlanningHistory) SetPlanningHistroyId(v string)`

SetPlanningHistroyId sets PlanningHistroyId field to given value.

### HasPlanningHistroyId

`func (o *RFPlanningHistory) HasPlanningHistroyId() bool`

HasPlanningHistroyId returns a boolean if a field has been set.

### GetPowerAdjustEnable

`func (o *RFPlanningHistory) GetPowerAdjustEnable() bool`

GetPowerAdjustEnable returns the PowerAdjustEnable field if non-nil, zero value otherwise.

### GetPowerAdjustEnableOk

`func (o *RFPlanningHistory) GetPowerAdjustEnableOk() (*bool, bool)`

GetPowerAdjustEnableOk returns a tuple with the PowerAdjustEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerAdjustEnable

`func (o *RFPlanningHistory) SetPowerAdjustEnable(v bool)`

SetPowerAdjustEnable sets PowerAdjustEnable field to given value.

### HasPowerAdjustEnable

`func (o *RFPlanningHistory) HasPowerAdjustEnable() bool`

HasPowerAdjustEnable returns a boolean if a field has been set.

### GetTime

`func (o *RFPlanningHistory) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RFPlanningHistory) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RFPlanningHistory) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *RFPlanningHistory) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


