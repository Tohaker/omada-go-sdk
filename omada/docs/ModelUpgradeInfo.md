# ModelUpgradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ModelList** | **[]string** | The selected model list. | 
**SelectType** | **string** | SelectType should be a value as follows: all: select all; include: include selected modelList; exclude: all but exclude selected modelList. | 

## Methods

### NewModelUpgradeInfo

`func NewModelUpgradeInfo(modelList []string, selectType string, ) *ModelUpgradeInfo`

NewModelUpgradeInfo instantiates a new ModelUpgradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelUpgradeInfoWithDefaults

`func NewModelUpgradeInfoWithDefaults() *ModelUpgradeInfo`

NewModelUpgradeInfoWithDefaults instantiates a new ModelUpgradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModelList

`func (o *ModelUpgradeInfo) GetModelList() []string`

GetModelList returns the ModelList field if non-nil, zero value otherwise.

### GetModelListOk

`func (o *ModelUpgradeInfo) GetModelListOk() (*[]string, bool)`

GetModelListOk returns a tuple with the ModelList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelList

`func (o *ModelUpgradeInfo) SetModelList(v []string)`

SetModelList sets ModelList field to given value.


### GetSelectType

`func (o *ModelUpgradeInfo) GetSelectType() string`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *ModelUpgradeInfo) GetSelectTypeOk() (*string, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *ModelUpgradeInfo) SetSelectType(v string)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


