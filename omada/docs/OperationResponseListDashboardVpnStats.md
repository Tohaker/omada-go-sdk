# OperationResponseListDashboardVpnStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**[]DashboardVpnStats**](DashboardVpnStats.md) |  | [optional] 

## Methods

### NewOperationResponseListDashboardVpnStats

`func NewOperationResponseListDashboardVpnStats() *OperationResponseListDashboardVpnStats`

NewOperationResponseListDashboardVpnStats instantiates a new OperationResponseListDashboardVpnStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseListDashboardVpnStatsWithDefaults

`func NewOperationResponseListDashboardVpnStatsWithDefaults() *OperationResponseListDashboardVpnStats`

NewOperationResponseListDashboardVpnStatsWithDefaults instantiates a new OperationResponseListDashboardVpnStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseListDashboardVpnStats) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseListDashboardVpnStats) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseListDashboardVpnStats) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseListDashboardVpnStats) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseListDashboardVpnStats) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseListDashboardVpnStats) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseListDashboardVpnStats) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseListDashboardVpnStats) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseListDashboardVpnStats) GetResult() []DashboardVpnStats`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseListDashboardVpnStats) GetResultOk() (*[]DashboardVpnStats, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseListDashboardVpnStats) SetResult(v []DashboardVpnStats)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseListDashboardVpnStats) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


