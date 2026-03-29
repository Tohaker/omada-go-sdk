# PrivilegeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Whether having all site privilege. | [optional] 
**Sites** | Pointer to **[]string** | The IDs of site that can be accessed by this user. | [optional] 

## Methods

### NewPrivilegeOpenApiVO

`func NewPrivilegeOpenApiVO() *PrivilegeOpenApiVO`

NewPrivilegeOpenApiVO instantiates a new PrivilegeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeOpenApiVOWithDefaults

`func NewPrivilegeOpenApiVOWithDefaults() *PrivilegeOpenApiVO`

NewPrivilegeOpenApiVOWithDefaults instantiates a new PrivilegeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *PrivilegeOpenApiVO) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *PrivilegeOpenApiVO) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *PrivilegeOpenApiVO) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *PrivilegeOpenApiVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetSites

`func (o *PrivilegeOpenApiVO) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PrivilegeOpenApiVO) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PrivilegeOpenApiVO) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *PrivilegeOpenApiVO) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


