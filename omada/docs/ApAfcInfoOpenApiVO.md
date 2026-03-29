# ApAfcInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationTimeSec** | Pointer to **int64** | The expiration timestamp of the current AFC information of the AP | [optional] 
**LastResponse** | Pointer to **bool** | The status of the last AFC information obtained by the AP | [optional] 
**LastResponseTimeSec** | Pointer to **int64** | The timestamp of the last AFC information obtained by the AP | [optional] 
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


