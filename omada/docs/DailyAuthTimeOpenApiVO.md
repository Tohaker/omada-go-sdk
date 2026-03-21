# DailyAuthTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | **int32** | Validity period should be a value as follows: 0: Custom; 1: 30 Minutes; 2: 1 Hour; 3: 2 Hours; 4: 4 Hours; 5: 8 Hours; | 
**CustomTimeout** | Pointer to **int32** | Custom timeout should be within the range of 1 ~ 1,440 minutes(when parameter[customTimeoutUnit] value is 1), or 1 ~ 24 hours(when parameter[customTimeoutUnit] value is 2). | [optional] 
**CustomTimeoutUnit** | Pointer to **int32** | Custom timeout unit should be a value as follows: 1: min; 2: hour. | [optional] 

## Methods

### NewDailyAuthTimeOpenApiVO

`func NewDailyAuthTimeOpenApiVO(authTimeout int32, ) *DailyAuthTimeOpenApiVO`

NewDailyAuthTimeOpenApiVO instantiates a new DailyAuthTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDailyAuthTimeOpenApiVOWithDefaults

`func NewDailyAuthTimeOpenApiVOWithDefaults() *DailyAuthTimeOpenApiVO`

NewDailyAuthTimeOpenApiVOWithDefaults instantiates a new DailyAuthTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *DailyAuthTimeOpenApiVO) GetAuthTimeout() int32`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *DailyAuthTimeOpenApiVO) GetAuthTimeoutOk() (*int32, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *DailyAuthTimeOpenApiVO) SetAuthTimeout(v int32)`

SetAuthTimeout sets AuthTimeout field to given value.


### GetCustomTimeout

`func (o *DailyAuthTimeOpenApiVO) GetCustomTimeout() int32`

GetCustomTimeout returns the CustomTimeout field if non-nil, zero value otherwise.

### GetCustomTimeoutOk

`func (o *DailyAuthTimeOpenApiVO) GetCustomTimeoutOk() (*int32, bool)`

GetCustomTimeoutOk returns a tuple with the CustomTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeout

`func (o *DailyAuthTimeOpenApiVO) SetCustomTimeout(v int32)`

SetCustomTimeout sets CustomTimeout field to given value.

### HasCustomTimeout

`func (o *DailyAuthTimeOpenApiVO) HasCustomTimeout() bool`

HasCustomTimeout returns a boolean if a field has been set.

### GetCustomTimeoutUnit

`func (o *DailyAuthTimeOpenApiVO) GetCustomTimeoutUnit() int32`

GetCustomTimeoutUnit returns the CustomTimeoutUnit field if non-nil, zero value otherwise.

### GetCustomTimeoutUnitOk

`func (o *DailyAuthTimeOpenApiVO) GetCustomTimeoutUnitOk() (*int32, bool)`

GetCustomTimeoutUnitOk returns a tuple with the CustomTimeoutUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeoutUnit

`func (o *DailyAuthTimeOpenApiVO) SetCustomTimeoutUnit(v int32)`

SetCustomTimeoutUnit sets CustomTimeoutUnit field to given value.

### HasCustomTimeoutUnit

`func (o *DailyAuthTimeOpenApiVO) HasCustomTimeoutUnit() bool`

HasCustomTimeoutUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


