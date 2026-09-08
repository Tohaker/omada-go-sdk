# IncidentVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Critical** | Pointer to **int32** | Critical incident count | [optional] 
**Error** | Pointer to **int32** | Error incident count | [optional] 
**Info** | Pointer to **int32** | Info incident count | [optional] 
**Total** | Pointer to **int32** | Total incident count | [optional] 
**Warning** | Pointer to **int32** | Warning incident count | [optional] 

## Methods

### NewIncidentVO

`func NewIncidentVO() *IncidentVO`

NewIncidentVO instantiates a new IncidentVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncidentVOWithDefaults

`func NewIncidentVOWithDefaults() *IncidentVO`

NewIncidentVOWithDefaults instantiates a new IncidentVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCritical

`func (o *IncidentVO) GetCritical() int32`

GetCritical returns the Critical field if non-nil, zero value otherwise.

### GetCriticalOk

`func (o *IncidentVO) GetCriticalOk() (*int32, bool)`

GetCriticalOk returns a tuple with the Critical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCritical

`func (o *IncidentVO) SetCritical(v int32)`

SetCritical sets Critical field to given value.

### HasCritical

`func (o *IncidentVO) HasCritical() bool`

HasCritical returns a boolean if a field has been set.

### GetError

`func (o *IncidentVO) GetError() int32`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IncidentVO) GetErrorOk() (*int32, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IncidentVO) SetError(v int32)`

SetError sets Error field to given value.

### HasError

`func (o *IncidentVO) HasError() bool`

HasError returns a boolean if a field has been set.

### GetInfo

`func (o *IncidentVO) GetInfo() int32`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *IncidentVO) GetInfoOk() (*int32, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *IncidentVO) SetInfo(v int32)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *IncidentVO) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetTotal

`func (o *IncidentVO) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IncidentVO) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IncidentVO) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IncidentVO) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWarning

`func (o *IncidentVO) GetWarning() int32`

GetWarning returns the Warning field if non-nil, zero value otherwise.

### GetWarningOk

`func (o *IncidentVO) GetWarningOk() (*int32, bool)`

GetWarningOk returns a tuple with the Warning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarning

`func (o *IncidentVO) SetWarning(v int32)`

SetWarning sets Warning field to given value.

### HasWarning

`func (o *IncidentVO) HasWarning() bool`

HasWarning returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


