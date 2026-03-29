# DnsCacheQueryOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Start page number. Start from 1. This subsection is deprecated. | [optional] 
**CurrentPageSize** | Pointer to **int32** | Number of entries per page. It should be within the range of 1–100. This subsection is deprecated. | [optional] 
**Type** | Pointer to **int32** | Type of DNS follows: 0: IPV4; 1: IPV6. This subsection is deprecated. | [optional] 

## Methods

### NewDnsCacheQueryOpenApiV2VO

`func NewDnsCacheQueryOpenApiV2VO() *DnsCacheQueryOpenApiV2VO`

NewDnsCacheQueryOpenApiV2VO instantiates a new DnsCacheQueryOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCacheQueryOpenApiV2VOWithDefaults

`func NewDnsCacheQueryOpenApiV2VOWithDefaults() *DnsCacheQueryOpenApiV2VO`

NewDnsCacheQueryOpenApiV2VOWithDefaults instantiates a new DnsCacheQueryOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *DnsCacheQueryOpenApiV2VO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *DnsCacheQueryOpenApiV2VO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *DnsCacheQueryOpenApiV2VO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *DnsCacheQueryOpenApiV2VO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentPageSize

`func (o *DnsCacheQueryOpenApiV2VO) GetCurrentPageSize() int32`

GetCurrentPageSize returns the CurrentPageSize field if non-nil, zero value otherwise.

### GetCurrentPageSizeOk

`func (o *DnsCacheQueryOpenApiV2VO) GetCurrentPageSizeOk() (*int32, bool)`

GetCurrentPageSizeOk returns a tuple with the CurrentPageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageSize

`func (o *DnsCacheQueryOpenApiV2VO) SetCurrentPageSize(v int32)`

SetCurrentPageSize sets CurrentPageSize field to given value.

### HasCurrentPageSize

`func (o *DnsCacheQueryOpenApiV2VO) HasCurrentPageSize() bool`

HasCurrentPageSize returns a boolean if a field has been set.

### GetType

`func (o *DnsCacheQueryOpenApiV2VO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsCacheQueryOpenApiV2VO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsCacheQueryOpenApiV2VO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsCacheQueryOpenApiV2VO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


