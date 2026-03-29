# BatchSelectMacsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Macs** | Pointer to **[]string** | When selectType is set to all, the macs do not need to be passed and all entries are processed, when selectType is set to include, the mac entries contained in the macs are processed, when selectType is set to exclude, the mac entries that are not contained in the macs are processed | [optional] 
**SelectType** | **string** | SelectType all, include or exclude | 

## Methods

### NewBatchSelectMacsVO

`func NewBatchSelectMacsVO(selectType string, ) *BatchSelectMacsVO`

NewBatchSelectMacsVO instantiates a new BatchSelectMacsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchSelectMacsVOWithDefaults

`func NewBatchSelectMacsVOWithDefaults() *BatchSelectMacsVO`

NewBatchSelectMacsVOWithDefaults instantiates a new BatchSelectMacsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMacs

`func (o *BatchSelectMacsVO) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *BatchSelectMacsVO) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *BatchSelectMacsVO) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *BatchSelectMacsVO) HasMacs() bool`

HasMacs returns a boolean if a field has been set.

### GetSelectType

`func (o *BatchSelectMacsVO) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *BatchSelectMacsVO) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *BatchSelectMacsVO) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


