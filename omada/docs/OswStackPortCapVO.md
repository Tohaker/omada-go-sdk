# OswStackPortCapVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupSpeedCap** | Pointer to **[]int32** | Stack port aggregation group link speed capability | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 

## Methods

### NewOswStackPortCapVO

`func NewOswStackPortCapVO() *OswStackPortCapVO`

NewOswStackPortCapVO instantiates a new OswStackPortCapVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackPortCapVOWithDefaults

`func NewOswStackPortCapVOWithDefaults() *OswStackPortCapVO`

NewOswStackPortCapVOWithDefaults instantiates a new OswStackPortCapVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupSpeedCap

`func (o *OswStackPortCapVO) GetGroupSpeedCap() []int32`

GetGroupSpeedCap returns the GroupSpeedCap field if non-nil, zero value otherwise.

### GetGroupSpeedCapOk

`func (o *OswStackPortCapVO) GetGroupSpeedCapOk() (*[]int32, bool)`

GetGroupSpeedCapOk returns a tuple with the GroupSpeedCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSpeedCap

`func (o *OswStackPortCapVO) SetGroupSpeedCap(v []int32)`

SetGroupSpeedCap sets GroupSpeedCap field to given value.

### HasGroupSpeedCap

`func (o *OswStackPortCapVO) HasGroupSpeedCap() bool`

HasGroupSpeedCap returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswStackPortCapVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswStackPortCapVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswStackPortCapVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswStackPortCapVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


