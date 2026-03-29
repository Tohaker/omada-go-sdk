# AppCategoryTrafficsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to [**[]CategoryTrafficDetailVO**](CategoryTrafficDetailVO.md) |  | [optional] 
**TotalTraffic** | Pointer to **int64** | total traffic | [optional] 

## Methods

### NewAppCategoryTrafficsVO

`func NewAppCategoryTrafficsVO() *AppCategoryTrafficsVO`

NewAppCategoryTrafficsVO instantiates a new AppCategoryTrafficsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppCategoryTrafficsVOWithDefaults

`func NewAppCategoryTrafficsVOWithDefaults() *AppCategoryTrafficsVO`

NewAppCategoryTrafficsVOWithDefaults instantiates a new AppCategoryTrafficsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *AppCategoryTrafficsVO) GetCategories() []CategoryTrafficDetailVO`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *AppCategoryTrafficsVO) GetCategoriesOk() (*[]CategoryTrafficDetailVO, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *AppCategoryTrafficsVO) SetCategories(v []CategoryTrafficDetailVO)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *AppCategoryTrafficsVO) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetTotalTraffic

`func (o *AppCategoryTrafficsVO) GetTotalTraffic() int64`

GetTotalTraffic returns the TotalTraffic field if non-nil, zero value otherwise.

### GetTotalTrafficOk

`func (o *AppCategoryTrafficsVO) GetTotalTrafficOk() (*int64, bool)`

GetTotalTrafficOk returns a tuple with the TotalTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTraffic

`func (o *AppCategoryTrafficsVO) SetTotalTraffic(v int64)`

SetTotalTraffic sets TotalTraffic field to given value.

### HasTotalTraffic

`func (o *AppCategoryTrafficsVO) HasTotalTraffic() bool`

HasTotalTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


