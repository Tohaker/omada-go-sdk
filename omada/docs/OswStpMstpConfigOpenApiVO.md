# OswStpMstpConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | [**[]OswStpInstanceConfigOpenApiVO**](OswStpInstanceConfigOpenApiVO.md) | Instances | 
**Region** | [**OswStpRegionConfigOpenApiVO**](OswStpRegionConfigOpenApiVO.md) |  | 

## Methods

### NewOswStpMstpConfigOpenApiVO

`func NewOswStpMstpConfigOpenApiVO(instances []OswStpInstanceConfigOpenApiVO, region OswStpRegionConfigOpenApiVO, ) *OswStpMstpConfigOpenApiVO`

NewOswStpMstpConfigOpenApiVO instantiates a new OswStpMstpConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStpMstpConfigOpenApiVOWithDefaults

`func NewOswStpMstpConfigOpenApiVOWithDefaults() *OswStpMstpConfigOpenApiVO`

NewOswStpMstpConfigOpenApiVOWithDefaults instantiates a new OswStpMstpConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstances

`func (o *OswStpMstpConfigOpenApiVO) GetInstances() []OswStpInstanceConfigOpenApiVO`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *OswStpMstpConfigOpenApiVO) GetInstancesOk() (*[]OswStpInstanceConfigOpenApiVO, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *OswStpMstpConfigOpenApiVO) SetInstances(v []OswStpInstanceConfigOpenApiVO)`

SetInstances sets Instances field to given value.


### GetRegion

`func (o *OswStpMstpConfigOpenApiVO) GetRegion() OswStpRegionConfigOpenApiVO`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *OswStpMstpConfigOpenApiVO) GetRegionOk() (*OswStpRegionConfigOpenApiVO, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *OswStpMstpConfigOpenApiVO) SetRegion(v OswStpRegionConfigOpenApiVO)`

SetRegion sets Region field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


