# ExportMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndTime** | **int64** | The timeStamp of the message end time. | 
**Format** | **int32** | Message parameter [format] should be 0 or 1. 0:csv; 1:xlsx. | 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM cardm parameter [simCard] should be 1 or 2. 1: SIM1; 2: SIM2. | [optional] 
**StartTime** | **int64** | The timeStamp of the message start time. | 
**Type** | **int32** | Message parameter [type] should be 0 or 1. 0:inbox; 1:outbox. | 

## Methods

### NewExportMessage

`func NewExportMessage(endTime int64, format int32, startTime int64, type_ int32, ) *ExportMessage`

NewExportMessage instantiates a new ExportMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportMessageWithDefaults

`func NewExportMessageWithDefaults() *ExportMessage`

NewExportMessageWithDefaults instantiates a new ExportMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndTime

`func (o *ExportMessage) GetEndTime() int64`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *ExportMessage) GetEndTimeOk() (*int64, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *ExportMessage) SetEndTime(v int64)`

SetEndTime sets EndTime field to given value.


### GetFormat

`func (o *ExportMessage) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportMessage) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportMessage) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetSimCard

`func (o *ExportMessage) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *ExportMessage) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *ExportMessage) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *ExportMessage) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetStartTime

`func (o *ExportMessage) GetStartTime() int64`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *ExportMessage) GetStartTimeOk() (*int64, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *ExportMessage) SetStartTime(v int64)`

SetStartTime sets StartTime field to given value.


### GetType

`func (o *ExportMessage) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExportMessage) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExportMessage) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


