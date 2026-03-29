# PrivilegeResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**All** | Pointer to **bool** | whether user has all site privilege | [optional] 
**Favorites** | Pointer to **[]string** |  | [optional] 
**LastVisited** | Pointer to **string** | user last visited site | [optional] 
**LastSiteCategory** | Pointer to **string** |  | [optional] 
**ServiceType** | Pointer to **int32** |  | [optional] 
**SiteNum** | Pointer to **int32** | user site num | [optional] 
**Sites** | Pointer to **[]map[string]map[string]interface{}** | user site privilege list | [optional] 

## Methods

### NewPrivilegeResultVO

`func NewPrivilegeResultVO() *PrivilegeResultVO`

NewPrivilegeResultVO instantiates a new PrivilegeResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivilegeResultVOWithDefaults

`func NewPrivilegeResultVOWithDefaults() *PrivilegeResultVO`

NewPrivilegeResultVOWithDefaults instantiates a new PrivilegeResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAll

`func (o *PrivilegeResultVO) GetAll() bool`

GetAll returns the All field if non-nil, zero value otherwise.

### GetAllOk

`func (o *PrivilegeResultVO) GetAllOk() (*bool, bool)`

GetAllOk returns a tuple with the All field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAll

`func (o *PrivilegeResultVO) SetAll(v bool)`

SetAll sets All field to given value.

### HasAll

`func (o *PrivilegeResultVO) HasAll() bool`

HasAll returns a boolean if a field has been set.

### GetFavorites

`func (o *PrivilegeResultVO) GetFavorites() []string`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *PrivilegeResultVO) GetFavoritesOk() (*[]string, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *PrivilegeResultVO) SetFavorites(v []string)`

SetFavorites sets Favorites field to given value.

### HasFavorites

`func (o *PrivilegeResultVO) HasFavorites() bool`

HasFavorites returns a boolean if a field has been set.

### GetLastVisited

`func (o *PrivilegeResultVO) GetLastVisited() string`

GetLastVisited returns the LastVisited field if non-nil, zero value otherwise.

### GetLastVisitedOk

`func (o *PrivilegeResultVO) GetLastVisitedOk() (*string, bool)`

GetLastVisitedOk returns a tuple with the LastVisited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVisited

`func (o *PrivilegeResultVO) SetLastVisited(v string)`

SetLastVisited sets LastVisited field to given value.

### HasLastVisited

`func (o *PrivilegeResultVO) HasLastVisited() bool`

HasLastVisited returns a boolean if a field has been set.

### GetLastSiteCategory

`func (o *PrivilegeResultVO) GetLastSiteCategory() string`

GetLastSiteCategory returns the LastSiteCategory field if non-nil, zero value otherwise.

### GetLastSiteCategoryOk

`func (o *PrivilegeResultVO) GetLastSiteCategoryOk() (*string, bool)`

GetLastSiteCategoryOk returns a tuple with the LastSiteCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSiteCategory

`func (o *PrivilegeResultVO) SetLastSiteCategory(v string)`

SetLastSiteCategory sets LastSiteCategory field to given value.

### HasLastSiteCategory

`func (o *PrivilegeResultVO) HasLastSiteCategory() bool`

HasLastSiteCategory returns a boolean if a field has been set.

### GetServiceType

`func (o *PrivilegeResultVO) GetServiceType() int32`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *PrivilegeResultVO) GetServiceTypeOk() (*int32, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *PrivilegeResultVO) SetServiceType(v int32)`

SetServiceType sets ServiceType field to given value.

### HasServiceType

`func (o *PrivilegeResultVO) HasServiceType() bool`

HasServiceType returns a boolean if a field has been set.

### GetSiteNum

`func (o *PrivilegeResultVO) GetSiteNum() int32`

GetSiteNum returns the SiteNum field if non-nil, zero value otherwise.

### GetSiteNumOk

`func (o *PrivilegeResultVO) GetSiteNumOk() (*int32, bool)`

GetSiteNumOk returns a tuple with the SiteNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteNum

`func (o *PrivilegeResultVO) SetSiteNum(v int32)`

SetSiteNum sets SiteNum field to given value.

### HasSiteNum

`func (o *PrivilegeResultVO) HasSiteNum() bool`

HasSiteNum returns a boolean if a field has been set.

### GetSites

`func (o *PrivilegeResultVO) GetSites() []map[string]map[string]interface{}`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *PrivilegeResultVO) GetSitesOk() (*[]map[string]map[string]interface{}, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *PrivilegeResultVO) SetSites(v []map[string]map[string]interface{})`

SetSites sets Sites field to given value.

### HasSites

`func (o *PrivilegeResultVO) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


