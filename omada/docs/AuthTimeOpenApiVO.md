# AuthTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthTimeout** | **int32** | Validity period should be a value as follows: 0: Custom; 1: 30 Minutes; 2: 1 Hour; 3: 2 Hours; 4: 4 Hours; 5: 8 Hours; 6: 1 Day; 7: 7 Days integer. | 
**CustomTimeout** | Pointer to **int32** | Custom timeout should be within the range of 1 ~ 1,000,000 minutes(when parameter[customTimeoutUnit] value is 1), or 1 ~ 10,000 hours(when parameter[customTimeoutUnit] value is 2), or 1 ~ 1,000 days(when parameter[customTimeoutUnit] value is 3). | [optional] 
**CustomTimeoutUnit** | Pointer to **int32** | Custom timeout unit should be a value as follows: 1: min; 2: hour; 3: day. | [optional] 

## Methods

### NewAuthTimeOpenApiVO

`func NewAuthTimeOpenApiVO(authTimeout int32, ) *AuthTimeOpenApiVO`

NewAuthTimeOpenApiVO instantiates a new AuthTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthTimeOpenApiVOWithDefaults

`func NewAuthTimeOpenApiVOWithDefaults() *AuthTimeOpenApiVO`

NewAuthTimeOpenApiVOWithDefaults instantiates a new AuthTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthTimeout

`func (o *AuthTimeOpenApiVO) GetAuthTimeout() int32`

GetAuthTimeout returns the AuthTimeout field if non-nil, zero value otherwise.

### GetAuthTimeoutOk

`func (o *AuthTimeOpenApiVO) GetAuthTimeoutOk() (*int32, bool)`

GetAuthTimeoutOk returns a tuple with the AuthTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTimeout

`func (o *AuthTimeOpenApiVO) SetAuthTimeout(v int32)`

SetAuthTimeout sets AuthTimeout field to given value.


### GetCustomTimeout

`func (o *AuthTimeOpenApiVO) GetCustomTimeout() int32`

GetCustomTimeout returns the CustomTimeout field if non-nil, zero value otherwise.

### GetCustomTimeoutOk

`func (o *AuthTimeOpenApiVO) GetCustomTimeoutOk() (*int32, bool)`

GetCustomTimeoutOk returns a tuple with the CustomTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeout

`func (o *AuthTimeOpenApiVO) SetCustomTimeout(v int32)`

SetCustomTimeout sets CustomTimeout field to given value.

### HasCustomTimeout

`func (o *AuthTimeOpenApiVO) HasCustomTimeout() bool`

HasCustomTimeout returns a boolean if a field has been set.

### GetCustomTimeoutUnit

`func (o *AuthTimeOpenApiVO) GetCustomTimeoutUnit() int32`

GetCustomTimeoutUnit returns the CustomTimeoutUnit field if non-nil, zero value otherwise.

### GetCustomTimeoutUnitOk

`func (o *AuthTimeOpenApiVO) GetCustomTimeoutUnitOk() (*int32, bool)`

GetCustomTimeoutUnitOk returns a tuple with the CustomTimeoutUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTimeoutUnit

`func (o *AuthTimeOpenApiVO) SetCustomTimeoutUnit(v int32)`

SetCustomTimeoutUnit sets CustomTimeoutUnit field to given value.

### HasCustomTimeoutUnit

`func (o *AuthTimeOpenApiVO) HasCustomTimeoutUnit() bool`

HasCustomTimeoutUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


