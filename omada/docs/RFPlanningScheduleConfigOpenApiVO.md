# RFPlanningScheduleConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Occurrence** | Pointer to [**Occurrence**](Occurrence.md) |  | [optional] 
**ScheduleEnable** | **bool** | Whether by WLAN Optimization schedule. The optimization schedule function is temporarily offline. | 

## Methods

### NewRFPlanningScheduleConfigOpenApiVO

`func NewRFPlanningScheduleConfigOpenApiVO(scheduleEnable bool, ) *RFPlanningScheduleConfigOpenApiVO`

NewRFPlanningScheduleConfigOpenApiVO instantiates a new RFPlanningScheduleConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFPlanningScheduleConfigOpenApiVOWithDefaults

`func NewRFPlanningScheduleConfigOpenApiVOWithDefaults() *RFPlanningScheduleConfigOpenApiVO`

NewRFPlanningScheduleConfigOpenApiVOWithDefaults instantiates a new RFPlanningScheduleConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOccurrence

`func (o *RFPlanningScheduleConfigOpenApiVO) GetOccurrence() Occurrence`

GetOccurrence returns the Occurrence field if non-nil, zero value otherwise.

### GetOccurrenceOk

`func (o *RFPlanningScheduleConfigOpenApiVO) GetOccurrenceOk() (*Occurrence, bool)`

GetOccurrenceOk returns a tuple with the Occurrence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOccurrence

`func (o *RFPlanningScheduleConfigOpenApiVO) SetOccurrence(v Occurrence)`

SetOccurrence sets Occurrence field to given value.

### HasOccurrence

`func (o *RFPlanningScheduleConfigOpenApiVO) HasOccurrence() bool`

HasOccurrence returns a boolean if a field has been set.

### GetScheduleEnable

`func (o *RFPlanningScheduleConfigOpenApiVO) GetScheduleEnable() bool`

GetScheduleEnable returns the ScheduleEnable field if non-nil, zero value otherwise.

### GetScheduleEnableOk

`func (o *RFPlanningScheduleConfigOpenApiVO) GetScheduleEnableOk() (*bool, bool)`

GetScheduleEnableOk returns a tuple with the ScheduleEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleEnable

`func (o *RFPlanningScheduleConfigOpenApiVO) SetScheduleEnable(v bool)`

SetScheduleEnable sets ScheduleEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


