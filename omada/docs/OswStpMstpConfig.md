# OswStpMstpConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]OswStpInstanceConfig**](OswStpInstanceConfig.md) | Instances | 
**Region** | [**OswStpRegionConfig**](OswStpRegionConfig.md) |  | 

## Methods

### NewOswStpMstpConfig

`func NewOswStpMstpConfig(instances []OswStpInstanceConfig, region OswStpRegionConfig, ) *OswStpMstpConfig`

NewOswStpMstpConfig instantiates a new OswStpMstpConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStpMstpConfigWithDefaults

`func NewOswStpMstpConfigWithDefaults() *OswStpMstpConfig`

NewOswStpMstpConfigWithDefaults instantiates a new OswStpMstpConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *OswStpMstpConfig) GetInstances() []OswStpInstanceConfig`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *OswStpMstpConfig) GetInstancesOk() (*[]OswStpInstanceConfig, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *OswStpMstpConfig) SetInstances(v []OswStpInstanceConfig)`

SetInstances sets Instances field to given value.


### GetRegion

`func (o *OswStpMstpConfig) GetRegion() OswStpRegionConfig`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OswStpMstpConfig) GetRegionOk() (*OswStpRegionConfig, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OswStpMstpConfig) SetRegion(v OswStpRegionConfig)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


