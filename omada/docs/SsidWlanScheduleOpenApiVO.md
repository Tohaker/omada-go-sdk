# SsidWlanScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **int32** | 0 means radio off, indicating the Wi-Fi function is off during the selected period; 1 means radio on, indicating the Wi-Fi function is on during the selected period. | [optional] 
**ScheduleId** | Pointer to **string** | This field represents Time Range Profile ID. Time Range Profile can be created using Create time range profile interface, and Time Range Profile ID can be obtained from Get time range profile list interface. | [optional] 
**WlanScheduleEnable** | Pointer to **bool** | SSID WLAN schedule global config status. True: enable, false: disable. | [optional] 

## Methods

### NewSsidWlanScheduleOpenApiVO

`func NewSsidWlanScheduleOpenApiVO() *SsidWlanScheduleOpenApiVO`

NewSsidWlanScheduleOpenApiVO instantiates a new SsidWlanScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidWlanScheduleOpenApiVOWithDefaults

`func NewSsidWlanScheduleOpenApiVOWithDefaults() *SsidWlanScheduleOpenApiVO`

NewSsidWlanScheduleOpenApiVOWithDefaults instantiates a new SsidWlanScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *SsidWlanScheduleOpenApiVO) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SsidWlanScheduleOpenApiVO) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SsidWlanScheduleOpenApiVO) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *SsidWlanScheduleOpenApiVO) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetScheduleId

`func (o *SsidWlanScheduleOpenApiVO) GetScheduleId() string`

GetScheduleId returns the ScheduleId field if non-nil, zero value otherwise.

### GetScheduleIdOk

`func (o *SsidWlanScheduleOpenApiVO) GetScheduleIdOk() (*string, bool)`

GetScheduleIdOk returns a tuple with the ScheduleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleId

`func (o *SsidWlanScheduleOpenApiVO) SetScheduleId(v string)`

SetScheduleId sets ScheduleId field to given value.

### HasScheduleId

`func (o *SsidWlanScheduleOpenApiVO) HasScheduleId() bool`

HasScheduleId returns a boolean if a field has been set.

### GetWlanScheduleEnable

`func (o *SsidWlanScheduleOpenApiVO) GetWlanScheduleEnable() bool`

GetWlanScheduleEnable returns the WlanScheduleEnable field if non-nil, zero value otherwise.

### GetWlanScheduleEnableOk

`func (o *SsidWlanScheduleOpenApiVO) GetWlanScheduleEnableOk() (*bool, bool)`

GetWlanScheduleEnableOk returns a tuple with the WlanScheduleEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanScheduleEnable

`func (o *SsidWlanScheduleOpenApiVO) SetWlanScheduleEnable(v bool)`

SetWlanScheduleEnable sets WlanScheduleEnable field to given value.

### HasWlanScheduleEnable

`func (o *SsidWlanScheduleOpenApiVO) HasWlanScheduleEnable() bool`

HasWlanScheduleEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


