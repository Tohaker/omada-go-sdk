# AlertVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CriticalAlerts** | Pointer to **int64** | Total number of critical alerts | [optional] 
**ErrorAlerts** | Pointer to **int64** | Total number of error alerts | [optional] 
**InfoAlerts** | Pointer to **int64** | Total number of info alerts | [optional] 
**TotalAlerts** | Pointer to **int64** | Total number of alerts | [optional] 
**WarningAlerts** | Pointer to **int64** | Total number of warning alerts | [optional] 

## Methods

### NewAlertVO

`func NewAlertVO() *AlertVO`

NewAlertVO instantiates a new AlertVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertVOWithDefaults

`func NewAlertVOWithDefaults() *AlertVO`

NewAlertVOWithDefaults instantiates a new AlertVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCriticalAlerts

`func (o *AlertVO) GetCriticalAlerts() int64`

GetCriticalAlerts returns the CriticalAlerts field if non-nil, zero value otherwise.

### GetCriticalAlertsOk

`func (o *AlertVO) GetCriticalAlertsOk() (*int64, bool)`

GetCriticalAlertsOk returns a tuple with the CriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAlerts

`func (o *AlertVO) SetCriticalAlerts(v int64)`

SetCriticalAlerts sets CriticalAlerts field to given value.

### HasCriticalAlerts

`func (o *AlertVO) HasCriticalAlerts() bool`

HasCriticalAlerts returns a boolean if a field has been set.

### GetErrorAlerts

`func (o *AlertVO) GetErrorAlerts() int64`

GetErrorAlerts returns the ErrorAlerts field if non-nil, zero value otherwise.

### GetErrorAlertsOk

`func (o *AlertVO) GetErrorAlertsOk() (*int64, bool)`

GetErrorAlertsOk returns a tuple with the ErrorAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorAlerts

`func (o *AlertVO) SetErrorAlerts(v int64)`

SetErrorAlerts sets ErrorAlerts field to given value.

### HasErrorAlerts

`func (o *AlertVO) HasErrorAlerts() bool`

HasErrorAlerts returns a boolean if a field has been set.

### GetInfoAlerts

`func (o *AlertVO) GetInfoAlerts() int64`

GetInfoAlerts returns the InfoAlerts field if non-nil, zero value otherwise.

### GetInfoAlertsOk

`func (o *AlertVO) GetInfoAlertsOk() (*int64, bool)`

GetInfoAlertsOk returns a tuple with the InfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoAlerts

`func (o *AlertVO) SetInfoAlerts(v int64)`

SetInfoAlerts sets InfoAlerts field to given value.

### HasInfoAlerts

`func (o *AlertVO) HasInfoAlerts() bool`

HasInfoAlerts returns a boolean if a field has been set.

### GetTotalAlerts

`func (o *AlertVO) GetTotalAlerts() int64`

GetTotalAlerts returns the TotalAlerts field if non-nil, zero value otherwise.

### GetTotalAlertsOk

`func (o *AlertVO) GetTotalAlertsOk() (*int64, bool)`

GetTotalAlertsOk returns a tuple with the TotalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAlerts

`func (o *AlertVO) SetTotalAlerts(v int64)`

SetTotalAlerts sets TotalAlerts field to given value.

### HasTotalAlerts

`func (o *AlertVO) HasTotalAlerts() bool`

HasTotalAlerts returns a boolean if a field has been set.

### GetWarningAlerts

`func (o *AlertVO) GetWarningAlerts() int64`

GetWarningAlerts returns the WarningAlerts field if non-nil, zero value otherwise.

### GetWarningAlertsOk

`func (o *AlertVO) GetWarningAlertsOk() (*int64, bool)`

GetWarningAlertsOk returns a tuple with the WarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningAlerts

`func (o *AlertVO) SetWarningAlerts(v int64)`

SetWarningAlerts sets WarningAlerts field to given value.

### HasWarningAlerts

`func (o *AlertVO) HasWarningAlerts() bool`

HasWarningAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


