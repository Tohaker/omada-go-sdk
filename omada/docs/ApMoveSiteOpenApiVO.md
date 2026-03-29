# ApMoveSiteOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMacs** | Pointer to **[]string** | aps mac list | [optional] 
**TargetSite** | **string** | Target site ID | 

## Methods

### NewApMoveSiteOpenApiVO

`func NewApMoveSiteOpenApiVO(targetSite string, ) *ApMoveSiteOpenApiVO`

NewApMoveSiteOpenApiVO instantiates a new ApMoveSiteOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMoveSiteOpenApiVOWithDefaults

`func NewApMoveSiteOpenApiVOWithDefaults() *ApMoveSiteOpenApiVO`

NewApMoveSiteOpenApiVOWithDefaults instantiates a new ApMoveSiteOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMacs

`func (o *ApMoveSiteOpenApiVO) GetApMacs() []string`

GetApMacs returns the ApMacs field if non-nil, zero value otherwise.

### GetApMacsOk

`func (o *ApMoveSiteOpenApiVO) GetApMacsOk() (*[]string, bool)`

GetApMacsOk returns a tuple with the ApMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMacs

`func (o *ApMoveSiteOpenApiVO) SetApMacs(v []string)`

SetApMacs sets ApMacs field to given value.

### HasApMacs

`func (o *ApMoveSiteOpenApiVO) HasApMacs() bool`

HasApMacs returns a boolean if a field has been set.

### GetTargetSite

`func (o *ApMoveSiteOpenApiVO) GetTargetSite() string`

GetTargetSite returns the TargetSite field if non-nil, zero value otherwise.

### GetTargetSiteOk

`func (o *ApMoveSiteOpenApiVO) GetTargetSiteOk() (*string, bool)`

GetTargetSiteOk returns a tuple with the TargetSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSite

`func (o *ApMoveSiteOpenApiVO) SetTargetSite(v string)`

SetTargetSite sets TargetSite field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


