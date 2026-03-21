# GlobalPortalDomainOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoRefresh** | Pointer to **bool** | Whether Portal URL enable Auto Refresh | [optional] 
**Url** | Pointer to **string** | Portal URL, valid when parameter [autoRefresh] is false | [optional] 

## Methods

### NewGlobalPortalDomainOpenApiVO

`func NewGlobalPortalDomainOpenApiVO() *GlobalPortalDomainOpenApiVO`

NewGlobalPortalDomainOpenApiVO instantiates a new GlobalPortalDomainOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalPortalDomainOpenApiVOWithDefaults

`func NewGlobalPortalDomainOpenApiVOWithDefaults() *GlobalPortalDomainOpenApiVO`

NewGlobalPortalDomainOpenApiVOWithDefaults instantiates a new GlobalPortalDomainOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoRefresh

`func (o *GlobalPortalDomainOpenApiVO) GetAutoRefresh() bool`

GetAutoRefresh returns the AutoRefresh field if non-nil, zero value otherwise.

### GetAutoRefreshOk

`func (o *GlobalPortalDomainOpenApiVO) GetAutoRefreshOk() (*bool, bool)`

GetAutoRefreshOk returns a tuple with the AutoRefresh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoRefresh

`func (o *GlobalPortalDomainOpenApiVO) SetAutoRefresh(v bool)`

SetAutoRefresh sets AutoRefresh field to given value.

### HasAutoRefresh

`func (o *GlobalPortalDomainOpenApiVO) HasAutoRefresh() bool`

HasAutoRefresh returns a boolean if a field has been set.

### GetUrl

`func (o *GlobalPortalDomainOpenApiVO) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GlobalPortalDomainOpenApiVO) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GlobalPortalDomainOpenApiVO) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GlobalPortalDomainOpenApiVO) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


