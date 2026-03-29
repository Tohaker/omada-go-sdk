# OswStpMstpVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]OswStpInstance**](OswStpInstance.md) | Instances | [optional] 
**Region** | Pointer to [**OswStpRegion**](OswStpRegion.md) |  | [optional] 

## Methods

### NewOswStpMstpVO

`func NewOswStpMstpVO() *OswStpMstpVO`

NewOswStpMstpVO instantiates a new OswStpMstpVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStpMstpVOWithDefaults

`func NewOswStpMstpVOWithDefaults() *OswStpMstpVO`

NewOswStpMstpVOWithDefaults instantiates a new OswStpMstpVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *OswStpMstpVO) GetInstances() []OswStpInstance`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *OswStpMstpVO) GetInstancesOk() (*[]OswStpInstance, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *OswStpMstpVO) SetInstances(v []OswStpInstance)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *OswStpMstpVO) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### GetRegion

`func (o *OswStpMstpVO) GetRegion() OswStpRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OswStpMstpVO) GetRegionOk() (*OswStpRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OswStpMstpVO) SetRegion(v OswStpRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *OswStpMstpVO) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


