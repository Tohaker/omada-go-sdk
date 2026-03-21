# RebootScheduleTemplateQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Reboot Schedule ID. | [optional] 
**Name** | **string** | Reboot Schedule name should contain 1 to 128 characters. | 
**NextExecute** | Pointer to **int64** | Reboot Schedule next execution timestamp(ms). | [optional] 
**Status** | **bool** | Reboot Schedule status. | 
**Time** | [**RebootScheduleTimeOpenApiVO**](RebootScheduleTimeOpenApiVO.md) |  | 

## Methods

### NewRebootScheduleTemplateQueryOpenApiVO

`func NewRebootScheduleTemplateQueryOpenApiVO(name string, status bool, time RebootScheduleTimeOpenApiVO, ) *RebootScheduleTemplateQueryOpenApiVO`

NewRebootScheduleTemplateQueryOpenApiVO instantiates a new RebootScheduleTemplateQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootScheduleTemplateQueryOpenApiVOWithDefaults

`func NewRebootScheduleTemplateQueryOpenApiVOWithDefaults() *RebootScheduleTemplateQueryOpenApiVO`

NewRebootScheduleTemplateQueryOpenApiVOWithDefaults instantiates a new RebootScheduleTemplateQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RebootScheduleTemplateQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RebootScheduleTemplateQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RebootScheduleTemplateQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecute

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetNextExecute() int64`

GetNextExecute returns the NextExecute field if non-nil, zero value otherwise.

### GetNextExecuteOk

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetNextExecuteOk() (*int64, bool)`

GetNextExecuteOk returns a tuple with the NextExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecute

`func (o *RebootScheduleTemplateQueryOpenApiVO) SetNextExecute(v int64)`

SetNextExecute sets NextExecute field to given value.

### HasNextExecute

`func (o *RebootScheduleTemplateQueryOpenApiVO) HasNextExecute() bool`

HasNextExecute returns a boolean if a field has been set.

### GetStatus

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RebootScheduleTemplateQueryOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTime

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetTime() RebootScheduleTimeOpenApiVO`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *RebootScheduleTemplateQueryOpenApiVO) GetTimeOk() (*RebootScheduleTimeOpenApiVO, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *RebootScheduleTemplateQueryOpenApiVO) SetTime(v RebootScheduleTimeOpenApiVO)`

SetTime sets Time field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


