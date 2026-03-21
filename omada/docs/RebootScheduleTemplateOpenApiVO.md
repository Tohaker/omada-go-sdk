# RebootScheduleTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Reboot Schedule name should contain 1 to 128 characters. | 
**Status** | **bool** | Reboot Schedule status. | 
**Time** | [**RebootScheduleTimeOpenApiVO**](RebootScheduleTimeOpenApiVO.md) |  | 

## Methods

### NewRebootScheduleTemplateOpenApiVO

`func NewRebootScheduleTemplateOpenApiVO(name string, status bool, time RebootScheduleTimeOpenApiVO, ) *RebootScheduleTemplateOpenApiVO`

NewRebootScheduleTemplateOpenApiVO instantiates a new RebootScheduleTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootScheduleTemplateOpenApiVOWithDefaults

`func NewRebootScheduleTemplateOpenApiVOWithDefaults() *RebootScheduleTemplateOpenApiVO`

NewRebootScheduleTemplateOpenApiVOWithDefaults instantiates a new RebootScheduleTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RebootScheduleTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebootScheduleTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebootScheduleTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *RebootScheduleTemplateOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RebootScheduleTemplateOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RebootScheduleTemplateOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTime

`func (o *RebootScheduleTemplateOpenApiVO) GetTime() RebootScheduleTimeOpenApiVO`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RebootScheduleTemplateOpenApiVO) GetTimeOk() (*RebootScheduleTimeOpenApiVO, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RebootScheduleTemplateOpenApiVO) SetTime(v RebootScheduleTimeOpenApiVO)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


