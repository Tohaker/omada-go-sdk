# GatewayCustomACLUpdateEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Adds** | Pointer to [**[]GatewayCustomACLAddEntity**](GatewayCustomACLAddEntity.md) | Added Custom ACLs. | [optional] 
**Deletes** | Pointer to **[]string** | Deleted Custom ACLs. | [optional] 
**Modifies** | Pointer to [**[]GatewayCustomACLModifyEntity**](GatewayCustomACLModifyEntity.md) | Modified Custom ACLs. | [optional] 

## Methods

### NewGatewayCustomACLUpdateEntity

`func NewGatewayCustomACLUpdateEntity() *GatewayCustomACLUpdateEntity`

NewGatewayCustomACLUpdateEntity instantiates a new GatewayCustomACLUpdateEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCustomACLUpdateEntityWithDefaults

`func NewGatewayCustomACLUpdateEntityWithDefaults() *GatewayCustomACLUpdateEntity`

NewGatewayCustomACLUpdateEntityWithDefaults instantiates a new GatewayCustomACLUpdateEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdds

`func (o *GatewayCustomACLUpdateEntity) GetAdds() []GatewayCustomACLAddEntity`

GetAdds returns the Adds field if non-nil, zero value otherwise.

### GetAddsOk

`func (o *GatewayCustomACLUpdateEntity) GetAddsOk() (*[]GatewayCustomACLAddEntity, bool)`

GetAddsOk returns a tuple with the Adds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdds

`func (o *GatewayCustomACLUpdateEntity) SetAdds(v []GatewayCustomACLAddEntity)`

SetAdds sets Adds field to given value.

### HasAdds

`func (o *GatewayCustomACLUpdateEntity) HasAdds() bool`

HasAdds returns a boolean if a field has been set.

### GetDeletes

`func (o *GatewayCustomACLUpdateEntity) GetDeletes() []string`

GetDeletes returns the Deletes field if non-nil, zero value otherwise.

### GetDeletesOk

`func (o *GatewayCustomACLUpdateEntity) GetDeletesOk() (*[]string, bool)`

GetDeletesOk returns a tuple with the Deletes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletes

`func (o *GatewayCustomACLUpdateEntity) SetDeletes(v []string)`

SetDeletes sets Deletes field to given value.

### HasDeletes

`func (o *GatewayCustomACLUpdateEntity) HasDeletes() bool`

HasDeletes returns a boolean if a field has been set.

### GetModifies

`func (o *GatewayCustomACLUpdateEntity) GetModifies() []GatewayCustomACLModifyEntity`

GetModifies returns the Modifies field if non-nil, zero value otherwise.

### GetModifiesOk

`func (o *GatewayCustomACLUpdateEntity) GetModifiesOk() (*[]GatewayCustomACLModifyEntity, bool)`

GetModifiesOk returns a tuple with the Modifies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifies

`func (o *GatewayCustomACLUpdateEntity) SetModifies(v []GatewayCustomACLModifyEntity)`

SetModifies sets Modifies field to given value.

### HasModifies

`func (o *GatewayCustomACLUpdateEntity) HasModifies() bool`

HasModifies returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


