# ApSpeedTestResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownRate** | Pointer to **float64** | Downlink rate | [optional] 
**ErrorCode** | Pointer to **int32** | This field will only be present if status is an error status. | [optional] 
**ErrorMsg** | Pointer to **string** | This field will only be present if status is an error status. | [optional] 
**Role** | Pointer to **int32** | Uplink or downlink AP for the current AP. The parameter [role] should be a value as follows:[0:Uplink; 1:Downlink]. | [optional] 
**Status** | Pointer to **int32** | The parameter [status] should be a value as follows:[-2:Device error; -1:System error; 0:No result; 1:Normal results]. | [optional] 
**Time** | Pointer to **int64** | Time in milliseconds. | [optional] 
**UpRate** | Pointer to **float64** | Uplink rate | [optional] 

## Methods

### NewApSpeedTestResultOpenApiVO

`func NewApSpeedTestResultOpenApiVO() *ApSpeedTestResultOpenApiVO`

NewApSpeedTestResultOpenApiVO instantiates a new ApSpeedTestResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApSpeedTestResultOpenApiVOWithDefaults

`func NewApSpeedTestResultOpenApiVOWithDefaults() *ApSpeedTestResultOpenApiVO`

NewApSpeedTestResultOpenApiVOWithDefaults instantiates a new ApSpeedTestResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownRate

`func (o *ApSpeedTestResultOpenApiVO) GetDownRate() float64`

GetDownRate returns the DownRate field if non-nil, zero value otherwise.

### GetDownRateOk

`func (o *ApSpeedTestResultOpenApiVO) GetDownRateOk() (*float64, bool)`

GetDownRateOk returns a tuple with the DownRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownRate

`func (o *ApSpeedTestResultOpenApiVO) SetDownRate(v float64)`

SetDownRate sets DownRate field to given value.

### HasDownRate

`func (o *ApSpeedTestResultOpenApiVO) HasDownRate() bool`

HasDownRate returns a boolean if a field has been set.

### GetErrorCode

`func (o *ApSpeedTestResultOpenApiVO) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ApSpeedTestResultOpenApiVO) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ApSpeedTestResultOpenApiVO) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ApSpeedTestResultOpenApiVO) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorMsg

`func (o *ApSpeedTestResultOpenApiVO) GetErrorMsg() string`

GetErrorMsg returns the ErrorMsg field if non-nil, zero value otherwise.

### GetErrorMsgOk

`func (o *ApSpeedTestResultOpenApiVO) GetErrorMsgOk() (*string, bool)`

GetErrorMsgOk returns a tuple with the ErrorMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMsg

`func (o *ApSpeedTestResultOpenApiVO) SetErrorMsg(v string)`

SetErrorMsg sets ErrorMsg field to given value.

### HasErrorMsg

`func (o *ApSpeedTestResultOpenApiVO) HasErrorMsg() bool`

HasErrorMsg returns a boolean if a field has been set.

### GetRole

`func (o *ApSpeedTestResultOpenApiVO) GetRole() int32`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ApSpeedTestResultOpenApiVO) GetRoleOk() (*int32, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ApSpeedTestResultOpenApiVO) SetRole(v int32)`

SetRole sets Role field to given value.

### HasRole

`func (o *ApSpeedTestResultOpenApiVO) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetStatus

`func (o *ApSpeedTestResultOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApSpeedTestResultOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApSpeedTestResultOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApSpeedTestResultOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *ApSpeedTestResultOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApSpeedTestResultOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApSpeedTestResultOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ApSpeedTestResultOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUpRate

`func (o *ApSpeedTestResultOpenApiVO) GetUpRate() float64`

GetUpRate returns the UpRate field if non-nil, zero value otherwise.

### GetUpRateOk

`func (o *ApSpeedTestResultOpenApiVO) GetUpRateOk() (*float64, bool)`

GetUpRateOk returns a tuple with the UpRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpRate

`func (o *ApSpeedTestResultOpenApiVO) SetUpRate(v float64)`

SetUpRate sets UpRate field to given value.

### HasUpRate

`func (o *ApSpeedTestResultOpenApiVO) HasUpRate() bool`

HasUpRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


