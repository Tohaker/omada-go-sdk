# DnsCacheOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Whether to enable DNS Cache. Options: &#39;true&#39; or &#39;false&#39;. | [optional] 
**Ttl** | Pointer to **int32** | TTL should be within the range of 1 or 86400. | [optional] 

## Methods

### NewDnsCacheOpenApiVO

`func NewDnsCacheOpenApiVO() *DnsCacheOpenApiVO`

NewDnsCacheOpenApiVO instantiates a new DnsCacheOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCacheOpenApiVOWithDefaults

`func NewDnsCacheOpenApiVOWithDefaults() *DnsCacheOpenApiVO`

NewDnsCacheOpenApiVOWithDefaults instantiates a new DnsCacheOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *DnsCacheOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DnsCacheOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DnsCacheOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DnsCacheOpenApiVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetTtl

`func (o *DnsCacheOpenApiVO) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsCacheOpenApiVO) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsCacheOpenApiVO) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsCacheOpenApiVO) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


