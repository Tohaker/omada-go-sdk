# LocationOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Index** | **int32** | Location index | 
**IspInfo** | Pointer to [**[]IspOpenApiVO**](IspOpenApiVO.md) | ISP info in the location, including index and name | [optional] 
**Name** | **string** | Location name | 

## Methods

### NewLocationOpenApiVO

`func NewLocationOpenApiVO(index int32, name string, ) *LocationOpenApiVO`

NewLocationOpenApiVO instantiates a new LocationOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationOpenApiVOWithDefaults

`func NewLocationOpenApiVOWithDefaults() *LocationOpenApiVO`

NewLocationOpenApiVOWithDefaults instantiates a new LocationOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIndex

`func (o *LocationOpenApiVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *LocationOpenApiVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *LocationOpenApiVO) SetIndex(v int32)`

SetIndex sets Index field to given value.


### GetIspInfo

`func (o *LocationOpenApiVO) GetIspInfo() []IspOpenApiVO`

GetIspInfo returns the IspInfo field if non-nil, zero value otherwise.

### GetIspInfoOk

`func (o *LocationOpenApiVO) GetIspInfoOk() (*[]IspOpenApiVO, bool)`

GetIspInfoOk returns a tuple with the IspInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspInfo

`func (o *LocationOpenApiVO) SetIspInfo(v []IspOpenApiVO)`

SetIspInfo sets IspInfo field to given value.

### HasIspInfo

`func (o *LocationOpenApiVO) HasIspInfo() bool`

HasIspInfo returns a boolean if a field has been set.

### GetName

`func (o *LocationOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocationOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocationOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


