# CreateSitePrivilegeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | Whether user has all site permission, including new created site. | [optional] 
**ServiceType** | Pointer to **int32** | Service type should be a value as follows: 1: Omada System. | [optional] 
**Sites** | Pointer to **[]string** | User site privilege list | [optional] 

## Methods

### NewCreateSitePrivilegeVO

`func NewCreateSitePrivilegeVO() *CreateSitePrivilegeVO`

NewCreateSitePrivilegeVO instantiates a new CreateSitePrivilegeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSitePrivilegeVOWithDefaults

`func NewCreateSitePrivilegeVOWithDefaults() *CreateSitePrivilegeVO`

NewCreateSitePrivilegeVOWithDefaults instantiates a new CreateSitePrivilegeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *CreateSitePrivilegeVO) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *CreateSitePrivilegeVO) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *CreateSitePrivilegeVO) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *CreateSitePrivilegeVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetServiceType

`func (o *CreateSitePrivilegeVO) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateSitePrivilegeVO) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateSitePrivilegeVO) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *CreateSitePrivilegeVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSites

`func (o *CreateSitePrivilegeVO) GetSites() []string`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *CreateSitePrivilegeVO) GetSitesOk() (*[]string, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *CreateSitePrivilegeVO) SetSites(v []string)`

SetSites sets Sites field to given value.

### HasSites

`func (o *CreateSitePrivilegeVO) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


