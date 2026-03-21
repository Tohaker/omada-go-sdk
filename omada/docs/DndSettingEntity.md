# DndSettingEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayMode** | Pointer to **int32** | The days you want to block the incoming calls. 1-daily，2-weekend，3-weekday | [optional] 
**Enable** | **bool** | Enable DND or not | 
**OmadacId** | Pointer to **string** | Omadac ID | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**TimeBegin** | Pointer to **int32** | The start time of the DND period you want to block incoming calls. It should be an integer value between 0 and 1438. | [optional] 
**TimeEnd** | Pointer to **int32** | The end time of the DND period you want to block incoming calls. It should be an integer value between 1 and 1439. | [optional] 

## Methods

### NewDndSettingEntity

`func NewDndSettingEntity(enable bool, ) *DndSettingEntity`

NewDndSettingEntity instantiates a new DndSettingEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDndSettingEntityWithDefaults

`func NewDndSettingEntityWithDefaults() *DndSettingEntity`

NewDndSettingEntityWithDefaults instantiates a new DndSettingEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayMode

`func (o *DndSettingEntity) GetDayMode() int32`

GetDayMode returns the DayMode field if non-nil, zero value otherwise.

### GetDayModeOk

`func (o *DndSettingEntity) GetDayModeOk() (*int32, bool)`

GetDayModeOk returns a tuple with the DayMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayMode

`func (o *DndSettingEntity) SetDayMode(v int32)`

SetDayMode sets DayMode field to given value.

### HasDayMode

`func (o *DndSettingEntity) HasDayMode() bool`

HasDayMode returns a boolean if a field has been set.

### GetEnable

`func (o *DndSettingEntity) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DndSettingEntity) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DndSettingEntity) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetOmadacId

`func (o *DndSettingEntity) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *DndSettingEntity) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *DndSettingEntity) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *DndSettingEntity) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSiteId

`func (o *DndSettingEntity) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *DndSettingEntity) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *DndSettingEntity) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *DndSettingEntity) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetTimeBegin

`func (o *DndSettingEntity) GetTimeBegin() int32`

GetTimeBegin returns the TimeBegin field if non-nil, zero value otherwise.

### GetTimeBeginOk

`func (o *DndSettingEntity) GetTimeBeginOk() (*int32, bool)`

GetTimeBeginOk returns a tuple with the TimeBegin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeBegin

`func (o *DndSettingEntity) SetTimeBegin(v int32)`

SetTimeBegin sets TimeBegin field to given value.

### HasTimeBegin

`func (o *DndSettingEntity) HasTimeBegin() bool`

HasTimeBegin returns a boolean if a field has been set.

### GetTimeEnd

`func (o *DndSettingEntity) GetTimeEnd() int32`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *DndSettingEntity) GetTimeEndOk() (*int32, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *DndSettingEntity) SetTimeEnd(v int32)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *DndSettingEntity) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


