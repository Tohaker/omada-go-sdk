# LocateResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocateEnable** | Pointer to **bool** | Indicates whether the locate function is enabled | [optional] 
**Mac** | Pointer to **string** | Mac | [optional] 
**RestTime** | Pointer to **int64** | The remaining time that the device locate switch is valid | [optional] 

## Methods

### NewLocateResultVO

`func NewLocateResultVO() *LocateResultVO`

NewLocateResultVO instantiates a new LocateResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocateResultVOWithDefaults

`func NewLocateResultVOWithDefaults() *LocateResultVO`

NewLocateResultVOWithDefaults instantiates a new LocateResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocateEnable

`func (o *LocateResultVO) GetLocateEnable() bool`

GetLocateEnable returns the LocateEnable field if non-nil, zero value otherwise.

### GetLocateEnableOk

`func (o *LocateResultVO) GetLocateEnableOk() (*bool, bool)`

GetLocateEnableOk returns a tuple with the LocateEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocateEnable

`func (o *LocateResultVO) SetLocateEnable(v bool)`

SetLocateEnable sets LocateEnable field to given value.

### HasLocateEnable

`func (o *LocateResultVO) HasLocateEnable() bool`

HasLocateEnable returns a boolean if a field has been set.

### GetMac

`func (o *LocateResultVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *LocateResultVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *LocateResultVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *LocateResultVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetRestTime

`func (o *LocateResultVO) GetRestTime() int64`

GetRestTime returns the RestTime field if non-nil, zero value otherwise.

### GetRestTimeOk

`func (o *LocateResultVO) GetRestTimeOk() (*int64, bool)`

GetRestTimeOk returns a tuple with the RestTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestTime

`func (o *LocateResultVO) SetRestTime(v int64)`

SetRestTime sets RestTime field to given value.

### HasRestTime

`func (o *LocateResultVO) HasRestTime() bool`

HasRestTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


