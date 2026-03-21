# DisasterModeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | Format should be a value as follows: 0: not support; 1: disaster mode enable; 2: disaster mode disable | [optional] 

## Methods

### NewDisasterModeOpenApiVO

`func NewDisasterModeOpenApiVO() *DisasterModeOpenApiVO`

NewDisasterModeOpenApiVO instantiates a new DisasterModeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisasterModeOpenApiVOWithDefaults

`func NewDisasterModeOpenApiVOWithDefaults() *DisasterModeOpenApiVO`

NewDisasterModeOpenApiVOWithDefaults instantiates a new DisasterModeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *DisasterModeOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisasterModeOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisasterModeOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DisasterModeOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


