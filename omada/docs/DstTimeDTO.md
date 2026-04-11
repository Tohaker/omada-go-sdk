# DstTimeDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Day** | Pointer to **int32** | Day of the DST config should be a value as follows: 1: Monday; 2: Tuesday; 3: Wednesday; 4: Thursday; 5: Friday; 6: Saturday; 7: Sunday. | [optional] 
**Hour** | Pointer to **int32** | Hour of the DST config should be within the range of 0–23. | [optional] 
**Minute** | Pointer to **int32** | Minute of the DST config should be within the range of 0–59. | [optional] 
**Month** | Pointer to **int32** | Month of the DST config should be a value as follows: 1: January; 2: February; 3: March; 4: April; 5: May; 6: June; 7: July; 8: August; 9: September; 10: October; 11: November; 12: December. | [optional] 
**Serial** | Pointer to **int32** | Week of the DST config should be a value as follows: 1: 1st; 2: 2nd; 3: 3rd; 4: 4th; 5: Last. | [optional] 

## Methods

### NewDstTimeDTO

`func NewDstTimeDTO() *DstTimeDTO`

NewDstTimeDTO instantiates a new DstTimeDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDstTimeDTOWithDefaults

`func NewDstTimeDTOWithDefaults() *DstTimeDTO`

NewDstTimeDTOWithDefaults instantiates a new DstTimeDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDay

`func (o *DstTimeDTO) GetDay() int32`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DstTimeDTO) GetDayOk() (*int32, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DstTimeDTO) SetDay(v int32)`

SetDay sets Day field to given value.

### HasDay

`func (o *DstTimeDTO) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetHour

`func (o *DstTimeDTO) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *DstTimeDTO) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *DstTimeDTO) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *DstTimeDTO) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetMinute

`func (o *DstTimeDTO) GetMinute() int32`

GetMinute returns the Minute field if non-nil, zero value otherwise.

### GetMinuteOk

`func (o *DstTimeDTO) GetMinuteOk() (*int32, bool)`

GetMinuteOk returns a tuple with the Minute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinute

`func (o *DstTimeDTO) SetMinute(v int32)`

SetMinute sets Minute field to given value.

### HasMinute

`func (o *DstTimeDTO) HasMinute() bool`

HasMinute returns a boolean if a field has been set.

### GetMonth

`func (o *DstTimeDTO) GetMonth() int32`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *DstTimeDTO) GetMonthOk() (*int32, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *DstTimeDTO) SetMonth(v int32)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *DstTimeDTO) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetSerial

`func (o *DstTimeDTO) GetSerial() int32`

GetSerial returns the Serial field if non-nil, zero value otherwise.

### GetSerialOk

`func (o *DstTimeDTO) GetSerialOk() (*int32, bool)`

GetSerialOk returns a tuple with the Serial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerial

`func (o *DstTimeDTO) SetSerial(v int32)`

SetSerial sets Serial field to given value.

### HasSerial

`func (o *DstTimeDTO) HasSerial() bool`

HasSerial returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


