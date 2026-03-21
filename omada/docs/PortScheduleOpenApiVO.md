# PortScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Port Schedule Name should contain 1 to 128 characters. | 
**PortsMap** | **map[string][]int32** | Key:MAC(\&quot;String\&quot;), Value:Set of Ports(\&quot;Integer\&quot;) | 
**Status** | **bool** | Port Schedule Status. | 
**TurnOnTime** | **string** | Time Range ID, cannot be empty. | 

## Methods

### NewPortScheduleOpenApiVO

`func NewPortScheduleOpenApiVO(name string, portsMap map[string][]int32, status bool, turnOnTime string, ) *PortScheduleOpenApiVO`

NewPortScheduleOpenApiVO instantiates a new PortScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortScheduleOpenApiVOWithDefaults

`func NewPortScheduleOpenApiVOWithDefaults() *PortScheduleOpenApiVO`

NewPortScheduleOpenApiVOWithDefaults instantiates a new PortScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PortScheduleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortScheduleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortScheduleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPortsMap

`func (o *PortScheduleOpenApiVO) GetPortsMap() map[string][]int32`

GetPortsMap returns the PortsMap field if non-nil, zero value otherwise.

### GetPortsMapOk

`func (o *PortScheduleOpenApiVO) GetPortsMapOk() (*map[string][]int32, bool)`

GetPortsMapOk returns a tuple with the PortsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsMap

`func (o *PortScheduleOpenApiVO) SetPortsMap(v map[string][]int32)`

SetPortsMap sets PortsMap field to given value.


### GetStatus

`func (o *PortScheduleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortScheduleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortScheduleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTurnOnTime

`func (o *PortScheduleOpenApiVO) GetTurnOnTime() string`

GetTurnOnTime returns the TurnOnTime field if non-nil, zero value otherwise.

### GetTurnOnTimeOk

`func (o *PortScheduleOpenApiVO) GetTurnOnTimeOk() (*string, bool)`

GetTurnOnTimeOk returns a tuple with the TurnOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnTime

`func (o *PortScheduleOpenApiVO) SetTurnOnTime(v string)`

SetTurnOnTime sets TurnOnTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


