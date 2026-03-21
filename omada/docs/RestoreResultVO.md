# RestoreResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **int32** | Status should be a value as follows: 0: restore finished; 1: restore prepared; 2: restore running; 3: restore failed | [optional] 

## Methods

### NewRestoreResultVO

`func NewRestoreResultVO() *RestoreResultVO`

NewRestoreResultVO instantiates a new RestoreResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestoreResultVOWithDefaults

`func NewRestoreResultVOWithDefaults() *RestoreResultVO`

NewRestoreResultVOWithDefaults instantiates a new RestoreResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *RestoreResultVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RestoreResultVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RestoreResultVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RestoreResultVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


