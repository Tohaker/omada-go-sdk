# PortOrPoeScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Port Schedule Name should contain 1 to 128 characters. | 
**NextExecute** | Pointer to **int64** | Show PoE Schedule next execution timestamp. | [optional] 
**NotInDst** | Pointer to **bool** | When notInDst is true, not displayed in DST. | [optional] 
**PortsMap** | **map[string][]int32** | Key:Mac(\&quot;String\&quot;), Value:Set of Ports(\&quot;Integer\&quot;) | 
**ScheduleType** | **int32** | Port Schedule Type should be 0 or 1. | 
**Status** | **bool** | Port Schedule Status. | 
**TurnOnTime** | **string** | Time Range ID, cannot be empty. | 

## Methods

### NewPortOrPoeScheduleOpenApiVO

`func NewPortOrPoeScheduleOpenApiVO(name string, portsMap map[string][]int32, scheduleType int32, status bool, turnOnTime string, ) *PortOrPoeScheduleOpenApiVO`

NewPortOrPoeScheduleOpenApiVO instantiates a new PortOrPoeScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortOrPoeScheduleOpenApiVOWithDefaults

`func NewPortOrPoeScheduleOpenApiVOWithDefaults() *PortOrPoeScheduleOpenApiVO`

NewPortOrPoeScheduleOpenApiVOWithDefaults instantiates a new PortOrPoeScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PortOrPoeScheduleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortOrPoeScheduleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortOrPoeScheduleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecute

`func (o *PortOrPoeScheduleOpenApiVO) GetNextExecute() int64`

GetNextExecute returns the NextExecute field if non-nil, zero value otherwise.

### GetNextExecuteOk

`func (o *PortOrPoeScheduleOpenApiVO) GetNextExecuteOk() (*int64, bool)`

GetNextExecuteOk returns a tuple with the NextExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecute

`func (o *PortOrPoeScheduleOpenApiVO) SetNextExecute(v int64)`

SetNextExecute sets NextExecute field to given value.

### HasNextExecute

`func (o *PortOrPoeScheduleOpenApiVO) HasNextExecute() bool`

HasNextExecute returns a boolean if a field has been set.

### GetNotInDst

`func (o *PortOrPoeScheduleOpenApiVO) GetNotInDst() bool`

GetNotInDst returns the NotInDst field if non-nil, zero value otherwise.

### GetNotInDstOk

`func (o *PortOrPoeScheduleOpenApiVO) GetNotInDstOk() (*bool, bool)`

GetNotInDstOk returns a tuple with the NotInDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInDst

`func (o *PortOrPoeScheduleOpenApiVO) SetNotInDst(v bool)`

SetNotInDst sets NotInDst field to given value.

### HasNotInDst

`func (o *PortOrPoeScheduleOpenApiVO) HasNotInDst() bool`

HasNotInDst returns a boolean if a field has been set.

### GetPortsMap

`func (o *PortOrPoeScheduleOpenApiVO) GetPortsMap() map[string][]int32`

GetPortsMap returns the PortsMap field if non-nil, zero value otherwise.

### GetPortsMapOk

`func (o *PortOrPoeScheduleOpenApiVO) GetPortsMapOk() (*map[string][]int32, bool)`

GetPortsMapOk returns a tuple with the PortsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsMap

`func (o *PortOrPoeScheduleOpenApiVO) SetPortsMap(v map[string][]int32)`

SetPortsMap sets PortsMap field to given value.


### GetScheduleType

`func (o *PortOrPoeScheduleOpenApiVO) GetScheduleType() int32`

GetScheduleType returns the ScheduleType field if non-nil, zero value otherwise.

### GetScheduleTypeOk

`func (o *PortOrPoeScheduleOpenApiVO) GetScheduleTypeOk() (*int32, bool)`

GetScheduleTypeOk returns a tuple with the ScheduleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduleType

`func (o *PortOrPoeScheduleOpenApiVO) SetScheduleType(v int32)`

SetScheduleType sets ScheduleType field to given value.


### GetStatus

`func (o *PortOrPoeScheduleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortOrPoeScheduleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortOrPoeScheduleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTurnOnTime

`func (o *PortOrPoeScheduleOpenApiVO) GetTurnOnTime() string`

GetTurnOnTime returns the TurnOnTime field if non-nil, zero value otherwise.

### GetTurnOnTimeOk

`func (o *PortOrPoeScheduleOpenApiVO) GetTurnOnTimeOk() (*string, bool)`

GetTurnOnTimeOk returns a tuple with the TurnOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnTime

`func (o *PortOrPoeScheduleOpenApiVO) SetTurnOnTime(v string)`

SetTurnOnTime sets TurnOnTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


