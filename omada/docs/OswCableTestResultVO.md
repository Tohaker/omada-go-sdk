# OswCableTestResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Length** | Pointer to **[]int32** | Cable length List. | [optional] 
**Port** | Pointer to **int32** | Port | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 
**State** | Pointer to **[]int32** | Cable state List. Each item should be a value as follows: 0:OK  1:OPEN  2:short 3:crosstalk 4:unknown-error 5:not-support 6:openshort | [optional] 
**TimeStamp** | Pointer to **int64** | Timestamp of the test result | [optional] 

## Methods

### NewOswCableTestResultVO

`func NewOswCableTestResultVO() *OswCableTestResultVO`

NewOswCableTestResultVO instantiates a new OswCableTestResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswCableTestResultVOWithDefaults

`func NewOswCableTestResultVOWithDefaults() *OswCableTestResultVO`

NewOswCableTestResultVOWithDefaults instantiates a new OswCableTestResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLength

`func (o *OswCableTestResultVO) GetLength() []int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *OswCableTestResultVO) GetLengthOk() (*[]int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *OswCableTestResultVO) SetLength(v []int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *OswCableTestResultVO) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetPort

`func (o *OswCableTestResultVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswCableTestResultVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswCableTestResultVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswCableTestResultVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswCableTestResultVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswCableTestResultVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswCableTestResultVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswCableTestResultVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetState

`func (o *OswCableTestResultVO) GetState() []int32`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *OswCableTestResultVO) GetStateOk() (*[]int32, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *OswCableTestResultVO) SetState(v []int32)`

SetState sets State field to given value.

### HasState

`func (o *OswCableTestResultVO) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTimeStamp

`func (o *OswCableTestResultVO) GetTimeStamp() int64`

GetTimeStamp returns the TimeStamp field if non-nil, zero value otherwise.

### GetTimeStampOk

`func (o *OswCableTestResultVO) GetTimeStampOk() (*int64, bool)`

GetTimeStampOk returns a tuple with the TimeStamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamp

`func (o *OswCableTestResultVO) SetTimeStamp(v int64)`

SetTimeStamp sets TimeStamp field to given value.

### HasTimeStamp

`func (o *OswCableTestResultVO) HasTimeStamp() bool`

HasTimeStamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


