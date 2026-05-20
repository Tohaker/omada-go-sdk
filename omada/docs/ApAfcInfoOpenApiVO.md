# ApAfcInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available6g** | Pointer to **bool** | 6G radio available status | [optional] 
**ErrCode** | Pointer to **int32** | The error code of last afc status | [optional] 
**ErrDetail** | Pointer to **int32** | The error detail | [optional] 
**ErrMain** | Pointer to **int32** | The error main reason | [optional] 
**ExpirationTimeSec** | Pointer to **int64** | The expiration timestamp of the current AFC information of the AP | [optional] 
**LastResponse** | Pointer to **bool** | The status of the last AFC information obtained by the AP | [optional] 
**LastResponseTimeSec** | Pointer to **int64** | The timestamp of the last AFC information obtained by the AP | [optional] 
**Processing** | Pointer to **bool** | Whether the afc status is being retrieved | [optional] 
**Status** | Pointer to **bool** | Ap AFC working status | [optional] 

## Methods

### NewApAfcInfoOpenApiVO

`func NewApAfcInfoOpenApiVO() *ApAfcInfoOpenApiVO`

NewApAfcInfoOpenApiVO instantiates a new ApAfcInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApAfcInfoOpenApiVOWithDefaults

`func NewApAfcInfoOpenApiVOWithDefaults() *ApAfcInfoOpenApiVO`

NewApAfcInfoOpenApiVOWithDefaults instantiates a new ApAfcInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable6g

`func (o *ApAfcInfoOpenApiVO) GetAvailable6g() bool`

GetAvailable6g returns the Available6g field if non-nil, zero value otherwise.

### GetAvailable6gOk

`func (o *ApAfcInfoOpenApiVO) GetAvailable6gOk() (*bool, bool)`

GetAvailable6gOk returns a tuple with the Available6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable6g

`func (o *ApAfcInfoOpenApiVO) SetAvailable6g(v bool)`

SetAvailable6g sets Available6g field to given value.

### HasAvailable6g

`func (o *ApAfcInfoOpenApiVO) HasAvailable6g() bool`

HasAvailable6g returns a boolean if a field has been set.

### GetErrCode

`func (o *ApAfcInfoOpenApiVO) GetErrCode() int32`

GetErrCode returns the ErrCode field if non-nil, zero value otherwise.

### GetErrCodeOk

`func (o *ApAfcInfoOpenApiVO) GetErrCodeOk() (*int32, bool)`

GetErrCodeOk returns a tuple with the ErrCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrCode

`func (o *ApAfcInfoOpenApiVO) SetErrCode(v int32)`

SetErrCode sets ErrCode field to given value.

### HasErrCode

`func (o *ApAfcInfoOpenApiVO) HasErrCode() bool`

HasErrCode returns a boolean if a field has been set.

### GetErrDetail

`func (o *ApAfcInfoOpenApiVO) GetErrDetail() int32`

GetErrDetail returns the ErrDetail field if non-nil, zero value otherwise.

### GetErrDetailOk

`func (o *ApAfcInfoOpenApiVO) GetErrDetailOk() (*int32, bool)`

GetErrDetailOk returns a tuple with the ErrDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrDetail

`func (o *ApAfcInfoOpenApiVO) SetErrDetail(v int32)`

SetErrDetail sets ErrDetail field to given value.

### HasErrDetail

`func (o *ApAfcInfoOpenApiVO) HasErrDetail() bool`

HasErrDetail returns a boolean if a field has been set.

### GetErrMain

`func (o *ApAfcInfoOpenApiVO) GetErrMain() int32`

GetErrMain returns the ErrMain field if non-nil, zero value otherwise.

### GetErrMainOk

`func (o *ApAfcInfoOpenApiVO) GetErrMainOk() (*int32, bool)`

GetErrMainOk returns a tuple with the ErrMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrMain

`func (o *ApAfcInfoOpenApiVO) SetErrMain(v int32)`

SetErrMain sets ErrMain field to given value.

### HasErrMain

`func (o *ApAfcInfoOpenApiVO) HasErrMain() bool`

HasErrMain returns a boolean if a field has been set.

### GetExpirationTimeSec

`func (o *ApAfcInfoOpenApiVO) GetExpirationTimeSec() int64`

GetExpirationTimeSec returns the ExpirationTimeSec field if non-nil, zero value otherwise.

### GetExpirationTimeSecOk

`func (o *ApAfcInfoOpenApiVO) GetExpirationTimeSecOk() (*int64, bool)`

GetExpirationTimeSecOk returns a tuple with the ExpirationTimeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimeSec

`func (o *ApAfcInfoOpenApiVO) SetExpirationTimeSec(v int64)`

SetExpirationTimeSec sets ExpirationTimeSec field to given value.

### HasExpirationTimeSec

`func (o *ApAfcInfoOpenApiVO) HasExpirationTimeSec() bool`

HasExpirationTimeSec returns a boolean if a field has been set.

### GetLastResponse

`func (o *ApAfcInfoOpenApiVO) GetLastResponse() bool`

GetLastResponse returns the LastResponse field if non-nil, zero value otherwise.

### GetLastResponseOk

`func (o *ApAfcInfoOpenApiVO) GetLastResponseOk() (*bool, bool)`

GetLastResponseOk returns a tuple with the LastResponse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponse

`func (o *ApAfcInfoOpenApiVO) SetLastResponse(v bool)`

SetLastResponse sets LastResponse field to given value.

### HasLastResponse

`func (o *ApAfcInfoOpenApiVO) HasLastResponse() bool`

HasLastResponse returns a boolean if a field has been set.

### GetLastResponseTimeSec

`func (o *ApAfcInfoOpenApiVO) GetLastResponseTimeSec() int64`

GetLastResponseTimeSec returns the LastResponseTimeSec field if non-nil, zero value otherwise.

### GetLastResponseTimeSecOk

`func (o *ApAfcInfoOpenApiVO) GetLastResponseTimeSecOk() (*int64, bool)`

GetLastResponseTimeSecOk returns a tuple with the LastResponseTimeSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastResponseTimeSec

`func (o *ApAfcInfoOpenApiVO) SetLastResponseTimeSec(v int64)`

SetLastResponseTimeSec sets LastResponseTimeSec field to given value.

### HasLastResponseTimeSec

`func (o *ApAfcInfoOpenApiVO) HasLastResponseTimeSec() bool`

HasLastResponseTimeSec returns a boolean if a field has been set.

### GetProcessing

`func (o *ApAfcInfoOpenApiVO) GetProcessing() bool`

GetProcessing returns the Processing field if non-nil, zero value otherwise.

### GetProcessingOk

`func (o *ApAfcInfoOpenApiVO) GetProcessingOk() (*bool, bool)`

GetProcessingOk returns a tuple with the Processing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessing

`func (o *ApAfcInfoOpenApiVO) SetProcessing(v bool)`

SetProcessing sets Processing field to given value.

### HasProcessing

`func (o *ApAfcInfoOpenApiVO) HasProcessing() bool`

HasProcessing returns a boolean if a field has been set.

### GetStatus

`func (o *ApAfcInfoOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApAfcInfoOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApAfcInfoOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApAfcInfoOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


