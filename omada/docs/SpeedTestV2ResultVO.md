# SpeedTestV2ResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortSpeedResults** | Pointer to [**[]SpeedTestV2ResultItemVO**](SpeedTestV2ResultItemVO.md) |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewSpeedTestV2ResultVO

`func NewSpeedTestV2ResultVO() *SpeedTestV2ResultVO`

NewSpeedTestV2ResultVO instantiates a new SpeedTestV2ResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSpeedTestV2ResultVOWithDefaults

`func NewSpeedTestV2ResultVOWithDefaults() *SpeedTestV2ResultVO`

NewSpeedTestV2ResultVOWithDefaults instantiates a new SpeedTestV2ResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortSpeedResults

`func (o *SpeedTestV2ResultVO) GetPortSpeedResults() []SpeedTestV2ResultItemVO`

GetPortSpeedResults returns the PortSpeedResults field if non-nil, zero value otherwise.

### GetPortSpeedResultsOk

`func (o *SpeedTestV2ResultVO) GetPortSpeedResultsOk() (*[]SpeedTestV2ResultItemVO, bool)`

GetPortSpeedResultsOk returns a tuple with the PortSpeedResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSpeedResults

`func (o *SpeedTestV2ResultVO) SetPortSpeedResults(v []SpeedTestV2ResultItemVO)`

SetPortSpeedResults sets PortSpeedResults field to given value.

### HasPortSpeedResults

`func (o *SpeedTestV2ResultVO) HasPortSpeedResults() bool`

HasPortSpeedResults returns a boolean if a field has been set.

### GetStatus

`func (o *SpeedTestV2ResultVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SpeedTestV2ResultVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SpeedTestV2ResultVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SpeedTestV2ResultVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


