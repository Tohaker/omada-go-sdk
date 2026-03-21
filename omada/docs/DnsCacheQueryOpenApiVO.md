# DnsCacheQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | Pointer to **int32** | Queried page | [optional] 
**PageSize** | Pointer to **int32** | Queried page size | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 0: ipv4; 1: ipv6. Default type is ipv4. | [optional] 

## Methods

### NewDnsCacheQueryOpenApiVO

`func NewDnsCacheQueryOpenApiVO() *DnsCacheQueryOpenApiVO`

NewDnsCacheQueryOpenApiVO instantiates a new DnsCacheQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCacheQueryOpenApiVOWithDefaults

`func NewDnsCacheQueryOpenApiVOWithDefaults() *DnsCacheQueryOpenApiVO`

NewDnsCacheQueryOpenApiVOWithDefaults instantiates a new DnsCacheQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *DnsCacheQueryOpenApiVO) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *DnsCacheQueryOpenApiVO) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *DnsCacheQueryOpenApiVO) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *DnsCacheQueryOpenApiVO) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetPageSize

`func (o *DnsCacheQueryOpenApiVO) GetPageSize() int32`

GetPageSize returns the PageSize field if non-nil, zero value otherwise.

### GetPageSizeOk

`func (o *DnsCacheQueryOpenApiVO) GetPageSizeOk() (*int32, bool)`

GetPageSizeOk returns a tuple with the PageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageSize

`func (o *DnsCacheQueryOpenApiVO) SetPageSize(v int32)`

SetPageSize sets PageSize field to given value.

### HasPageSize

`func (o *DnsCacheQueryOpenApiVO) HasPageSize() bool`

HasPageSize returns a boolean if a field has been set.

### GetType

`func (o *DnsCacheQueryOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsCacheQueryOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsCacheQueryOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsCacheQueryOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


