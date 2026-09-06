# SelectLocalUsersOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]string** | IDs of localUsers to be deleted in batch operation | [optional] 
**Type** | **int32** | Select type. It should be a value as follows: 0: Represents selecting all localUsers, this selection does not pass parameter [ids]. 1: Parameter [ids] includes the IDs of the localUsers to be selected. 2: Parameter [ids] includes the IDs of the localUsers not to be selected | 

## Methods

### NewSelectLocalUsersOpenApiVO

`func NewSelectLocalUsersOpenApiVO(type_ int32, ) *SelectLocalUsersOpenApiVO`

NewSelectLocalUsersOpenApiVO instantiates a new SelectLocalUsersOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectLocalUsersOpenApiVOWithDefaults

`func NewSelectLocalUsersOpenApiVOWithDefaults() *SelectLocalUsersOpenApiVO`

NewSelectLocalUsersOpenApiVOWithDefaults instantiates a new SelectLocalUsersOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *SelectLocalUsersOpenApiVO) GetIds() []string`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *SelectLocalUsersOpenApiVO) GetIdsOk() (*[]string, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *SelectLocalUsersOpenApiVO) SetIds(v []string)`

SetIds sets Ids field to given value.

### HasIds

`func (o *SelectLocalUsersOpenApiVO) HasIds() bool`

HasIds returns a boolean if a field has been set.

### GetType

`func (o *SelectLocalUsersOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SelectLocalUsersOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SelectLocalUsersOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


