# DiskCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Disk condition name | [optional] 
**OcStorage** | Pointer to **bool** | OC storage of disk | [optional] 
**TotalStorage** | Pointer to **float64** | Total storage of disk | [optional] 
**UsedStorage** | Pointer to **float64** | Used storage of disk | [optional] 

## Methods

### NewDiskCondition

`func NewDiskCondition() *DiskCondition`

NewDiskCondition instantiates a new DiskCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiskConditionWithDefaults

`func NewDiskConditionWithDefaults() *DiskCondition`

NewDiskConditionWithDefaults instantiates a new DiskCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DiskCondition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DiskCondition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DiskCondition) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DiskCondition) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOcStorage

`func (o *DiskCondition) GetOcStorage() bool`

GetOcStorage returns the OcStorage field if non-nil, zero value otherwise.

### GetOcStorageOk

`func (o *DiskCondition) GetOcStorageOk() (*bool, bool)`

GetOcStorageOk returns a tuple with the OcStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcStorage

`func (o *DiskCondition) SetOcStorage(v bool)`

SetOcStorage sets OcStorage field to given value.

### HasOcStorage

`func (o *DiskCondition) HasOcStorage() bool`

HasOcStorage returns a boolean if a field has been set.

### GetTotalStorage

`func (o *DiskCondition) GetTotalStorage() float64`

GetTotalStorage returns the TotalStorage field if non-nil, zero value otherwise.

### GetTotalStorageOk

`func (o *DiskCondition) GetTotalStorageOk() (*float64, bool)`

GetTotalStorageOk returns a tuple with the TotalStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalStorage

`func (o *DiskCondition) SetTotalStorage(v float64)`

SetTotalStorage sets TotalStorage field to given value.

### HasTotalStorage

`func (o *DiskCondition) HasTotalStorage() bool`

HasTotalStorage returns a boolean if a field has been set.

### GetUsedStorage

`func (o *DiskCondition) GetUsedStorage() float64`

GetUsedStorage returns the UsedStorage field if non-nil, zero value otherwise.

### GetUsedStorageOk

`func (o *DiskCondition) GetUsedStorageOk() (*float64, bool)`

GetUsedStorageOk returns a tuple with the UsedStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedStorage

`func (o *DiskCondition) SetUsedStorage(v float64)`

SetUsedStorage sets UsedStorage field to given value.

### HasUsedStorage

`func (o *DiskCondition) HasUsedStorage() bool`

HasUsedStorage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


