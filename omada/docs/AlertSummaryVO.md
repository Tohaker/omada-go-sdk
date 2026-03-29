# AlertSummaryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientAlerts** | Pointer to **int64** | Number of client alerts | [optional] 
**CriticalAlerts** | Pointer to **int64** | Number of critical alerts | [optional] 
**DeviceAlerts** | Pointer to **int64** | Number of device alerts | [optional] 
**ErrorAlerts** | Pointer to **int64** | Number of error alerts | [optional] 
**InfoAlerts** | Pointer to **int64** | Number of info alerts | [optional] 
**OperationAlerts** | Pointer to **int64** | Number of operation alerts | [optional] 
**SystemAlerts** | Pointer to **int64** | Number of system alerts | [optional] 
**TotalGradeAlerts** | Pointer to **int64** | Number of alerts: sum by alert level | [optional] 
**TotalTypeAlerts** | Pointer to **int64** | Number of alerts: sum by alert type | [optional] 
**WarningAlerts** | Pointer to **int64** | Number of warning alerts | [optional] 

## Methods

### NewAlertSummaryVO

`func NewAlertSummaryVO() *AlertSummaryVO`

NewAlertSummaryVO instantiates a new AlertSummaryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlertSummaryVOWithDefaults

`func NewAlertSummaryVOWithDefaults() *AlertSummaryVO`

NewAlertSummaryVOWithDefaults instantiates a new AlertSummaryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientAlerts

`func (o *AlertSummaryVO) GetClientAlerts() int64`

GetClientAlerts returns the ClientAlerts field if non-nil, zero value otherwise.

### GetClientAlertsOk

`func (o *AlertSummaryVO) GetClientAlertsOk() (*int64, bool)`

GetClientAlertsOk returns a tuple with the ClientAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAlerts

`func (o *AlertSummaryVO) SetClientAlerts(v int64)`

SetClientAlerts sets ClientAlerts field to given value.

### HasClientAlerts

`func (o *AlertSummaryVO) HasClientAlerts() bool`

HasClientAlerts returns a boolean if a field has been set.

### GetCriticalAlerts

`func (o *AlertSummaryVO) GetCriticalAlerts() int64`

GetCriticalAlerts returns the CriticalAlerts field if non-nil, zero value otherwise.

### GetCriticalAlertsOk

`func (o *AlertSummaryVO) GetCriticalAlertsOk() (*int64, bool)`

GetCriticalAlertsOk returns a tuple with the CriticalAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriticalAlerts

`func (o *AlertSummaryVO) SetCriticalAlerts(v int64)`

SetCriticalAlerts sets CriticalAlerts field to given value.

### HasCriticalAlerts

`func (o *AlertSummaryVO) HasCriticalAlerts() bool`

HasCriticalAlerts returns a boolean if a field has been set.

### GetDeviceAlerts

`func (o *AlertSummaryVO) GetDeviceAlerts() int64`

GetDeviceAlerts returns the DeviceAlerts field if non-nil, zero value otherwise.

### GetDeviceAlertsOk

`func (o *AlertSummaryVO) GetDeviceAlertsOk() (*int64, bool)`

GetDeviceAlertsOk returns a tuple with the DeviceAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAlerts

`func (o *AlertSummaryVO) SetDeviceAlerts(v int64)`

SetDeviceAlerts sets DeviceAlerts field to given value.

### HasDeviceAlerts

`func (o *AlertSummaryVO) HasDeviceAlerts() bool`

HasDeviceAlerts returns a boolean if a field has been set.

### GetErrorAlerts

`func (o *AlertSummaryVO) GetErrorAlerts() int64`

GetErrorAlerts returns the ErrorAlerts field if non-nil, zero value otherwise.

### GetErrorAlertsOk

`func (o *AlertSummaryVO) GetErrorAlertsOk() (*int64, bool)`

GetErrorAlertsOk returns a tuple with the ErrorAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorAlerts

`func (o *AlertSummaryVO) SetErrorAlerts(v int64)`

SetErrorAlerts sets ErrorAlerts field to given value.

### HasErrorAlerts

`func (o *AlertSummaryVO) HasErrorAlerts() bool`

HasErrorAlerts returns a boolean if a field has been set.

### GetInfoAlerts

`func (o *AlertSummaryVO) GetInfoAlerts() int64`

GetInfoAlerts returns the InfoAlerts field if non-nil, zero value otherwise.

### GetInfoAlertsOk

`func (o *AlertSummaryVO) GetInfoAlertsOk() (*int64, bool)`

GetInfoAlertsOk returns a tuple with the InfoAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfoAlerts

`func (o *AlertSummaryVO) SetInfoAlerts(v int64)`

SetInfoAlerts sets InfoAlerts field to given value.

### HasInfoAlerts

`func (o *AlertSummaryVO) HasInfoAlerts() bool`

HasInfoAlerts returns a boolean if a field has been set.

### GetOperationAlerts

`func (o *AlertSummaryVO) GetOperationAlerts() int64`

GetOperationAlerts returns the OperationAlerts field if non-nil, zero value otherwise.

### GetOperationAlertsOk

`func (o *AlertSummaryVO) GetOperationAlertsOk() (*int64, bool)`

GetOperationAlertsOk returns a tuple with the OperationAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationAlerts

`func (o *AlertSummaryVO) SetOperationAlerts(v int64)`

SetOperationAlerts sets OperationAlerts field to given value.

### HasOperationAlerts

`func (o *AlertSummaryVO) HasOperationAlerts() bool`

HasOperationAlerts returns a boolean if a field has been set.

### GetSystemAlerts

`func (o *AlertSummaryVO) GetSystemAlerts() int64`

GetSystemAlerts returns the SystemAlerts field if non-nil, zero value otherwise.

### GetSystemAlertsOk

`func (o *AlertSummaryVO) GetSystemAlertsOk() (*int64, bool)`

GetSystemAlertsOk returns a tuple with the SystemAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemAlerts

`func (o *AlertSummaryVO) SetSystemAlerts(v int64)`

SetSystemAlerts sets SystemAlerts field to given value.

### HasSystemAlerts

`func (o *AlertSummaryVO) HasSystemAlerts() bool`

HasSystemAlerts returns a boolean if a field has been set.

### GetTotalGradeAlerts

`func (o *AlertSummaryVO) GetTotalGradeAlerts() int64`

GetTotalGradeAlerts returns the TotalGradeAlerts field if non-nil, zero value otherwise.

### GetTotalGradeAlertsOk

`func (o *AlertSummaryVO) GetTotalGradeAlertsOk() (*int64, bool)`

GetTotalGradeAlertsOk returns a tuple with the TotalGradeAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGradeAlerts

`func (o *AlertSummaryVO) SetTotalGradeAlerts(v int64)`

SetTotalGradeAlerts sets TotalGradeAlerts field to given value.

### HasTotalGradeAlerts

`func (o *AlertSummaryVO) HasTotalGradeAlerts() bool`

HasTotalGradeAlerts returns a boolean if a field has been set.

### GetTotalTypeAlerts

`func (o *AlertSummaryVO) GetTotalTypeAlerts() int64`

GetTotalTypeAlerts returns the TotalTypeAlerts field if non-nil, zero value otherwise.

### GetTotalTypeAlertsOk

`func (o *AlertSummaryVO) GetTotalTypeAlertsOk() (*int64, bool)`

GetTotalTypeAlertsOk returns a tuple with the TotalTypeAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTypeAlerts

`func (o *AlertSummaryVO) SetTotalTypeAlerts(v int64)`

SetTotalTypeAlerts sets TotalTypeAlerts field to given value.

### HasTotalTypeAlerts

`func (o *AlertSummaryVO) HasTotalTypeAlerts() bool`

HasTotalTypeAlerts returns a boolean if a field has been set.

### GetWarningAlerts

`func (o *AlertSummaryVO) GetWarningAlerts() int64`

GetWarningAlerts returns the WarningAlerts field if non-nil, zero value otherwise.

### GetWarningAlertsOk

`func (o *AlertSummaryVO) GetWarningAlertsOk() (*int64, bool)`

GetWarningAlertsOk returns a tuple with the WarningAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarningAlerts

`func (o *AlertSummaryVO) SetWarningAlerts(v int64)`

SetWarningAlerts sets WarningAlerts field to given value.

### HasWarningAlerts

`func (o *AlertSummaryVO) HasWarningAlerts() bool`

HasWarningAlerts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


