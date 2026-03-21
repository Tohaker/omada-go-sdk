# RFPlanningResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AfterIndex** | Pointer to **int32** | Index after WLAN Optimization, between 0 and 100. | [optional] 
**BeforeIndex** | Pointer to **int32** | Index before WLAN Optimization, between 0 and 100. | [optional] 
**FailType** | Pointer to **int32** | Type of failure. 0: Failure before deployment. 1: Failure after deployment. | [optional] 
**PlanningHistroyId** | Pointer to **string** | Planning histroy ID. | [optional] 
**Status** | **int32** | 0: Not in scanning status and a WLAN Optimization has been just successfully executed. 1: Not in scanning status and there&#39;s no WLAN Optimization result. 2: In RF planning status. 3: In RF planning canceling status. | 

## Methods

### NewRFPlanningResult

`func NewRFPlanningResult(status int32, ) *RFPlanningResult`

NewRFPlanningResult instantiates a new RFPlanningResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFPlanningResultWithDefaults

`func NewRFPlanningResultWithDefaults() *RFPlanningResult`

NewRFPlanningResultWithDefaults instantiates a new RFPlanningResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAfterIndex

`func (o *RFPlanningResult) GetAfterIndex() int32`

GetAfterIndex returns the AfterIndex field if non-nil, zero value otherwise.

### GetAfterIndexOk

`func (o *RFPlanningResult) GetAfterIndexOk() (*int32, bool)`

GetAfterIndexOk returns a tuple with the AfterIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAfterIndex

`func (o *RFPlanningResult) SetAfterIndex(v int32)`

SetAfterIndex sets AfterIndex field to given value.

### HasAfterIndex

`func (o *RFPlanningResult) HasAfterIndex() bool`

HasAfterIndex returns a boolean if a field has been set.

### GetBeforeIndex

`func (o *RFPlanningResult) GetBeforeIndex() int32`

GetBeforeIndex returns the BeforeIndex field if non-nil, zero value otherwise.

### GetBeforeIndexOk

`func (o *RFPlanningResult) GetBeforeIndexOk() (*int32, bool)`

GetBeforeIndexOk returns a tuple with the BeforeIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeforeIndex

`func (o *RFPlanningResult) SetBeforeIndex(v int32)`

SetBeforeIndex sets BeforeIndex field to given value.

### HasBeforeIndex

`func (o *RFPlanningResult) HasBeforeIndex() bool`

HasBeforeIndex returns a boolean if a field has been set.

### GetFailType

`func (o *RFPlanningResult) GetFailType() int32`

GetFailType returns the FailType field if non-nil, zero value otherwise.

### GetFailTypeOk

`func (o *RFPlanningResult) GetFailTypeOk() (*int32, bool)`

GetFailTypeOk returns a tuple with the FailType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailType

`func (o *RFPlanningResult) SetFailType(v int32)`

SetFailType sets FailType field to given value.

### HasFailType

`func (o *RFPlanningResult) HasFailType() bool`

HasFailType returns a boolean if a field has been set.

### GetPlanningHistroyId

`func (o *RFPlanningResult) GetPlanningHistroyId() string`

GetPlanningHistroyId returns the PlanningHistroyId field if non-nil, zero value otherwise.

### GetPlanningHistroyIdOk

`func (o *RFPlanningResult) GetPlanningHistroyIdOk() (*string, bool)`

GetPlanningHistroyIdOk returns a tuple with the PlanningHistroyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanningHistroyId

`func (o *RFPlanningResult) SetPlanningHistroyId(v string)`

SetPlanningHistroyId sets PlanningHistroyId field to given value.

### HasPlanningHistroyId

`func (o *RFPlanningResult) HasPlanningHistroyId() bool`

HasPlanningHistroyId returns a boolean if a field has been set.

### GetStatus

`func (o *RFPlanningResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RFPlanningResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RFPlanningResult) SetStatus(v int32)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


