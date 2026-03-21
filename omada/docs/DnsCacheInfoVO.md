# DnsCacheInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Domain** | Pointer to **string** |  | [optional] 
**IpList** | Pointer to **[]string** |  | [optional] 
**Ttl** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewDnsCacheInfoVO

`func NewDnsCacheInfoVO() *DnsCacheInfoVO`

NewDnsCacheInfoVO instantiates a new DnsCacheInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsCacheInfoVOWithDefaults

`func NewDnsCacheInfoVOWithDefaults() *DnsCacheInfoVO`

NewDnsCacheInfoVOWithDefaults instantiates a new DnsCacheInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomain

`func (o *DnsCacheInfoVO) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *DnsCacheInfoVO) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *DnsCacheInfoVO) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *DnsCacheInfoVO) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetIpList

`func (o *DnsCacheInfoVO) GetIpList() []string`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *DnsCacheInfoVO) GetIpListOk() (*[]string, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *DnsCacheInfoVO) SetIpList(v []string)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *DnsCacheInfoVO) HasIpList() bool`

HasIpList returns a boolean if a field has been set.

### GetTtl

`func (o *DnsCacheInfoVO) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *DnsCacheInfoVO) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *DnsCacheInfoVO) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *DnsCacheInfoVO) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetType

`func (o *DnsCacheInfoVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DnsCacheInfoVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DnsCacheInfoVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *DnsCacheInfoVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


