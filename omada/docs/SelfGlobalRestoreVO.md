# SelfGlobalRestoreVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileName** | **string** | Site backup file name. Parameter [fileName] should be 1 - 128 ASCII characters. | 

## Methods

### NewSelfGlobalRestoreVO

`func NewSelfGlobalRestoreVO(fileName string, ) *SelfGlobalRestoreVO`

NewSelfGlobalRestoreVO instantiates a new SelfGlobalRestoreVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfGlobalRestoreVOWithDefaults

`func NewSelfGlobalRestoreVOWithDefaults() *SelfGlobalRestoreVO`

NewSelfGlobalRestoreVOWithDefaults instantiates a new SelfGlobalRestoreVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileName

`func (o *SelfGlobalRestoreVO) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *SelfGlobalRestoreVO) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *SelfGlobalRestoreVO) SetFileName(v string)`

SetFileName sets FileName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


