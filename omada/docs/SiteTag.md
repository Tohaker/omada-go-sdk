# SiteTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Tag name should contain 1 to 128 ASCII characters. | [optional] 
**TagId** | Pointer to **string** | Tag ID | [optional] 

## Methods

### NewSiteTag

`func NewSiteTag() *SiteTag`

NewSiteTag instantiates a new SiteTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSiteTagWithDefaults

`func NewSiteTagWithDefaults() *SiteTag`

NewSiteTagWithDefaults instantiates a new SiteTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SiteTag) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SiteTag) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SiteTag) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SiteTag) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTagId

`func (o *SiteTag) GetTagId() string`

GetTagId returns the TagId field if non-nil, zero value otherwise.

### GetTagIdOk

`func (o *SiteTag) GetTagIdOk() (*string, bool)`

GetTagIdOk returns a tuple with the TagId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagId

`func (o *SiteTag) SetTagId(v string)`

SetTagId sets TagId field to given value.

### HasTagId

`func (o *SiteTag) HasTagId() bool`

HasTagId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


