# OswCableTestLogOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CancelPortList** | Pointer to [**[]OswBriefPortInfoOpenApiVO**](OswBriefPortInfoOpenApiVO.md) | Ports that were not detected due to cancel. Used for type 2 | [optional] 
**CompletePortList** | Pointer to [**[]OswBriefPortInfoOpenApiVO**](OswBriefPortInfoOpenApiVO.md) | Ports that have been tested completely. Used for type 1,2,3 | [optional] 
**TestPortList** | Pointer to [**[]OswBriefPortInfoOpenApiVO**](OswBriefPortInfoOpenApiVO.md) | All test port. Used for type 0 | [optional] 
**TimeStamp** | Pointer to **int64** | Timestamp of the log | [optional] 
**TimeoutPortList** | Pointer to [**[]OswBriefPortInfoOpenApiVO**](OswBriefPortInfoOpenApiVO.md) | Timeout ports. Used for type 3 | [optional] 
**Type** | Pointer to **int32** | Log type. It should be a value as follows: 0:start test. 1: test done. 2:test interrupts 3:time out | [optional] 

## Methods

### NewOswCableTestLogOpenApiVO

`func NewOswCableTestLogOpenApiVO() *OswCableTestLogOpenApiVO`

NewOswCableTestLogOpenApiVO instantiates a new OswCableTestLogOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswCableTestLogOpenApiVOWithDefaults

`func NewOswCableTestLogOpenApiVOWithDefaults() *OswCableTestLogOpenApiVO`

NewOswCableTestLogOpenApiVOWithDefaults instantiates a new OswCableTestLogOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCancelPortList

`func (o *OswCableTestLogOpenApiVO) GetCancelPortList() []OswBriefPortInfoOpenApiVO`

GetCancelPortList returns the CancelPortList field if non-nil, zero value otherwise.

### GetCancelPortListOk

`func (o *OswCableTestLogOpenApiVO) GetCancelPortListOk() (*[]OswBriefPortInfoOpenApiVO, bool)`

GetCancelPortListOk returns a tuple with the CancelPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelPortList

`func (o *OswCableTestLogOpenApiVO) SetCancelPortList(v []OswBriefPortInfoOpenApiVO)`

SetCancelPortList sets CancelPortList field to given value.

### HasCancelPortList

`func (o *OswCableTestLogOpenApiVO) HasCancelPortList() bool`

HasCancelPortList returns a boolean if a field has been set.

### GetCompletePortList

`func (o *OswCableTestLogOpenApiVO) GetCompletePortList() []OswBriefPortInfoOpenApiVO`

GetCompletePortList returns the CompletePortList field if non-nil, zero value otherwise.

### GetCompletePortListOk

`func (o *OswCableTestLogOpenApiVO) GetCompletePortListOk() (*[]OswBriefPortInfoOpenApiVO, bool)`

GetCompletePortListOk returns a tuple with the CompletePortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletePortList

`func (o *OswCableTestLogOpenApiVO) SetCompletePortList(v []OswBriefPortInfoOpenApiVO)`

SetCompletePortList sets CompletePortList field to given value.

### HasCompletePortList

`func (o *OswCableTestLogOpenApiVO) HasCompletePortList() bool`

HasCompletePortList returns a boolean if a field has been set.

### GetTestPortList

`func (o *OswCableTestLogOpenApiVO) GetTestPortList() []OswBriefPortInfoOpenApiVO`

GetTestPortList returns the TestPortList field if non-nil, zero value otherwise.

### GetTestPortListOk

`func (o *OswCableTestLogOpenApiVO) GetTestPortListOk() (*[]OswBriefPortInfoOpenApiVO, bool)`

GetTestPortListOk returns a tuple with the TestPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestPortList

`func (o *OswCableTestLogOpenApiVO) SetTestPortList(v []OswBriefPortInfoOpenApiVO)`

SetTestPortList sets TestPortList field to given value.

### HasTestPortList

`func (o *OswCableTestLogOpenApiVO) HasTestPortList() bool`

HasTestPortList returns a boolean if a field has been set.

### GetTimeStamp

`func (o *OswCableTestLogOpenApiVO) GetTimeStamp() int64`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *OswCableTestLogOpenApiVO) GetTimeStampOk() (*int64, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *OswCableTestLogOpenApiVO) SetTimeStamp(v int64)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *OswCableTestLogOpenApiVO) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.

### GetTimeoutPortList

`func (o *OswCableTestLogOpenApiVO) GetTimeoutPortList() []OswBriefPortInfoOpenApiVO`

GetTimeoutPortList returns the TimeoutPortList field if non-nil, zero value otherwise.

### GetTimeoutPortListOk

`func (o *OswCableTestLogOpenApiVO) GetTimeoutPortListOk() (*[]OswBriefPortInfoOpenApiVO, bool)`

GetTimeoutPortListOk returns a tuple with the TimeoutPortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutPortList

`func (o *OswCableTestLogOpenApiVO) SetTimeoutPortList(v []OswBriefPortInfoOpenApiVO)`

SetTimeoutPortList sets TimeoutPortList field to given value.

### HasTimeoutPortList

`func (o *OswCableTestLogOpenApiVO) HasTimeoutPortList() bool`

HasTimeoutPortList returns a boolean if a field has been set.

### GetType

`func (o *OswCableTestLogOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswCableTestLogOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswCableTestLogOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswCableTestLogOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


