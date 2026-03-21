# BatchOspfInterfaceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceIdList** | **[]string** | List of Interface ID | 
**SearchKey** | Pointer to **string** | SearchKey | [optional] 
**SelectType** | **int32** | Select Type, it should be a value as follows: 1: SELECT ALL, 2: SELECT INCLUDE, 3: SELECT EXCLUDE. | 

## Methods

### NewBatchOspfInterfaceOpenApiVO

`func NewBatchOspfInterfaceOpenApiVO(interfaceIdList []string, selectType int32, ) *BatchOspfInterfaceOpenApiVO`

NewBatchOspfInterfaceOpenApiVO instantiates a new BatchOspfInterfaceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBatchOspfInterfaceOpenApiVOWithDefaults

`func NewBatchOspfInterfaceOpenApiVOWithDefaults() *BatchOspfInterfaceOpenApiVO`

NewBatchOspfInterfaceOpenApiVOWithDefaults instantiates a new BatchOspfInterfaceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceIdList

`func (o *BatchOspfInterfaceOpenApiVO) GetInterfaceIdList() []string`

GetInterfaceIdList returns the InterfaceIdList field if non-nil, zero value otherwise.

### GetInterfaceIdListOk

`func (o *BatchOspfInterfaceOpenApiVO) GetInterfaceIdListOk() (*[]string, bool)`

GetInterfaceIdListOk returns a tuple with the InterfaceIdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceIdList

`func (o *BatchOspfInterfaceOpenApiVO) SetInterfaceIdList(v []string)`

SetInterfaceIdList sets InterfaceIdList field to given value.


### GetSearchKey

`func (o *BatchOspfInterfaceOpenApiVO) GetSearchKey() string`

GetSearchKey returns the SearchKey field if non-nil, zero value otherwise.

### GetSearchKeyOk

`func (o *BatchOspfInterfaceOpenApiVO) GetSearchKeyOk() (*string, bool)`

GetSearchKeyOk returns a tuple with the SearchKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchKey

`func (o *BatchOspfInterfaceOpenApiVO) SetSearchKey(v string)`

SetSearchKey sets SearchKey field to given value.

### HasSearchKey

`func (o *BatchOspfInterfaceOpenApiVO) HasSearchKey() bool`

HasSearchKey returns a boolean if a field has been set.

### GetSelectType

`func (o *BatchOspfInterfaceOpenApiVO) GetSelectType() int32`

GetSelectType returns the SelectType field if non-nil, zero value otherwise.

### GetSelectTypeOk

`func (o *BatchOspfInterfaceOpenApiVO) GetSelectTypeOk() (*int32, bool)`

GetSelectTypeOk returns a tuple with the SelectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectType

`func (o *BatchOspfInterfaceOpenApiVO) SetSelectType(v int32)`

SetSelectType sets SelectType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


