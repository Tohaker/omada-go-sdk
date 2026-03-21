# SnoopConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LeaveTime** | Pointer to **int32** | The time it takes to remove the port from the group after receiving the outbound message (unit: s). | [optional] 
**MemberPortAgingTime** | Pointer to **int32** | The aging time of member ports (unit: s). | [optional] 
**ReportSuppressionEnable** | Pointer to **bool** | Indicates whether to enable message suppression. | [optional] 
**ReportSuppressionExceptDevice** | Pointer to [**ReportSuppressionConfigVO**](ReportSuppressionConfigVO.md) |  | [optional] 
**RouterPortAgingTime** | Pointer to **int32** | The aging time of the routing port (unit: s). | [optional] 

## Methods

### NewSnoopConfigVO

`func NewSnoopConfigVO() *SnoopConfigVO`

NewSnoopConfigVO instantiates a new SnoopConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnoopConfigVOWithDefaults

`func NewSnoopConfigVOWithDefaults() *SnoopConfigVO`

NewSnoopConfigVOWithDefaults instantiates a new SnoopConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLeaveTime

`func (o *SnoopConfigVO) GetLeaveTime() int32`

GetLeaveTime returns the LeaveTime field if non-nil, zero value otherwise.

### GetLeaveTimeOk

`func (o *SnoopConfigVO) GetLeaveTimeOk() (*int32, bool)`

GetLeaveTimeOk returns a tuple with the LeaveTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaveTime

`func (o *SnoopConfigVO) SetLeaveTime(v int32)`

SetLeaveTime sets LeaveTime field to given value.

### HasLeaveTime

`func (o *SnoopConfigVO) HasLeaveTime() bool`

HasLeaveTime returns a boolean if a field has been set.

### GetMemberPortAgingTime

`func (o *SnoopConfigVO) GetMemberPortAgingTime() int32`

GetMemberPortAgingTime returns the MemberPortAgingTime field if non-nil, zero value otherwise.

### GetMemberPortAgingTimeOk

`func (o *SnoopConfigVO) GetMemberPortAgingTimeOk() (*int32, bool)`

GetMemberPortAgingTimeOk returns a tuple with the MemberPortAgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberPortAgingTime

`func (o *SnoopConfigVO) SetMemberPortAgingTime(v int32)`

SetMemberPortAgingTime sets MemberPortAgingTime field to given value.

### HasMemberPortAgingTime

`func (o *SnoopConfigVO) HasMemberPortAgingTime() bool`

HasMemberPortAgingTime returns a boolean if a field has been set.

### GetReportSuppressionEnable

`func (o *SnoopConfigVO) GetReportSuppressionEnable() bool`

GetReportSuppressionEnable returns the ReportSuppressionEnable field if non-nil, zero value otherwise.

### GetReportSuppressionEnableOk

`func (o *SnoopConfigVO) GetReportSuppressionEnableOk() (*bool, bool)`

GetReportSuppressionEnableOk returns a tuple with the ReportSuppressionEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportSuppressionEnable

`func (o *SnoopConfigVO) SetReportSuppressionEnable(v bool)`

SetReportSuppressionEnable sets ReportSuppressionEnable field to given value.

### HasReportSuppressionEnable

`func (o *SnoopConfigVO) HasReportSuppressionEnable() bool`

HasReportSuppressionEnable returns a boolean if a field has been set.

### GetReportSuppressionExceptDevice

`func (o *SnoopConfigVO) GetReportSuppressionExceptDevice() ReportSuppressionConfigVO`

GetReportSuppressionExceptDevice returns the ReportSuppressionExceptDevice field if non-nil, zero value otherwise.

### GetReportSuppressionExceptDeviceOk

`func (o *SnoopConfigVO) GetReportSuppressionExceptDeviceOk() (*ReportSuppressionConfigVO, bool)`

GetReportSuppressionExceptDeviceOk returns a tuple with the ReportSuppressionExceptDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportSuppressionExceptDevice

`func (o *SnoopConfigVO) SetReportSuppressionExceptDevice(v ReportSuppressionConfigVO)`

SetReportSuppressionExceptDevice sets ReportSuppressionExceptDevice field to given value.

### HasReportSuppressionExceptDevice

`func (o *SnoopConfigVO) HasReportSuppressionExceptDevice() bool`

HasReportSuppressionExceptDevice returns a boolean if a field has been set.

### GetRouterPortAgingTime

`func (o *SnoopConfigVO) GetRouterPortAgingTime() int32`

GetRouterPortAgingTime returns the RouterPortAgingTime field if non-nil, zero value otherwise.

### GetRouterPortAgingTimeOk

`func (o *SnoopConfigVO) GetRouterPortAgingTimeOk() (*int32, bool)`

GetRouterPortAgingTimeOk returns a tuple with the RouterPortAgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterPortAgingTime

`func (o *SnoopConfigVO) SetRouterPortAgingTime(v int32)`

SetRouterPortAgingTime sets RouterPortAgingTime field to given value.

### HasRouterPortAgingTime

`func (o *SnoopConfigVO) HasRouterPortAgingTime() bool`

HasRouterPortAgingTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


