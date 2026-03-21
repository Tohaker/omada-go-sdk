# RebootResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac | [optional] 
**Status** | Pointer to **int32** | Status | [optional] 
**StatusCategory** | Pointer to **int32** | Status Category | [optional] 

## Methods

### NewRebootResultVO

`func NewRebootResultVO() *RebootResultVO`

NewRebootResultVO instantiates a new RebootResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRebootResultVOWithDefaults

`func NewRebootResultVOWithDefaults() *RebootResultVO`

NewRebootResultVOWithDefaults instantiates a new RebootResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *RebootResultVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *RebootResultVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *RebootResultVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *RebootResultVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *RebootResultVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RebootResultVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RebootResultVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RebootResultVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *RebootResultVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *RebootResultVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *RebootResultVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *RebootResultVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


