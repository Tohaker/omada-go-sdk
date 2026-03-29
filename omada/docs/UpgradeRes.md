# UpgradeRes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Finished** | Pointer to **bool** | Whether the task is complete | [optional] 
**UnfinishedCount** | Pointer to **int32** | Total number of devices that have not been upgraded | [optional] 

## Methods

### NewUpgradeRes

`func NewUpgradeRes() *UpgradeRes`

NewUpgradeRes instantiates a new UpgradeRes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeResWithDefaults

`func NewUpgradeResWithDefaults() *UpgradeRes`

NewUpgradeResWithDefaults instantiates a new UpgradeRes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFinished

`func (o *UpgradeRes) GetFinished() bool`

GetFinished returns the Finished field if non-nil, zero value otherwise.

### GetFinishedOk

`func (o *UpgradeRes) GetFinishedOk() (*bool, bool)`

GetFinishedOk returns a tuple with the Finished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinished

`func (o *UpgradeRes) SetFinished(v bool)`

SetFinished sets Finished field to given value.

### HasFinished

`func (o *UpgradeRes) HasFinished() bool`

HasFinished returns a boolean if a field has been set.

### GetUnfinishedCount

`func (o *UpgradeRes) GetUnfinishedCount() int32`

GetUnfinishedCount returns the UnfinishedCount field if non-nil, zero value otherwise.

### GetUnfinishedCountOk

`func (o *UpgradeRes) GetUnfinishedCountOk() (*int32, bool)`

GetUnfinishedCountOk returns a tuple with the UnfinishedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnfinishedCount

`func (o *UpgradeRes) SetUnfinishedCount(v int32)`

SetUnfinishedCount sets UnfinishedCount field to given value.

### HasUnfinishedCount

`func (o *UpgradeRes) HasUnfinishedCount() bool`

HasUnfinishedCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


