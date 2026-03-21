# OperationResponseTrafficDistribution

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Result** | Pointer to [**TrafficDistribution**](TrafficDistribution.md) |  | [optional] 

## Methods

### NewOperationResponseTrafficDistribution

`func NewOperationResponseTrafficDistribution() *OperationResponseTrafficDistribution`

NewOperationResponseTrafficDistribution instantiates a new OperationResponseTrafficDistribution object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationResponseTrafficDistributionWithDefaults

`func NewOperationResponseTrafficDistributionWithDefaults() *OperationResponseTrafficDistribution`

NewOperationResponseTrafficDistributionWithDefaults instantiates a new OperationResponseTrafficDistribution object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *OperationResponseTrafficDistribution) GetErrorCode() int32`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *OperationResponseTrafficDistribution) GetErrorCodeOk() (*int32, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *OperationResponseTrafficDistribution) SetErrorCode(v int32)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *OperationResponseTrafficDistribution) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetMsg

`func (o *OperationResponseTrafficDistribution) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *OperationResponseTrafficDistribution) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *OperationResponseTrafficDistribution) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *OperationResponseTrafficDistribution) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetResult

`func (o *OperationResponseTrafficDistribution) GetResult() TrafficDistribution`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *OperationResponseTrafficDistribution) GetResultOk() (*TrafficDistribution, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *OperationResponseTrafficDistribution) SetResult(v TrafficDistribution)`

SetResult sets Result field to given value.

### HasResult

`func (o *OperationResponseTrafficDistribution) HasResult() bool`

HasResult returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


