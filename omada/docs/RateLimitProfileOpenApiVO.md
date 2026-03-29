# RateLimitProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultProfile** | Pointer to **bool** | Whether it is default profile. | [optional] 
**DownLimit** | Pointer to **int64** | Download limit(Unit: Kbps), this field is required when parameter [downLimitEnable] is true. | [optional] 
**DownLimitEnable** | Pointer to **bool** | Whether to enable download limit | [optional] 
**Name** | Pointer to **string** | Rate limit profile name | [optional] 
**ProfileId** | Pointer to **string** | Rate limit profile ID | [optional] 
**UpLimit** | Pointer to **int64** | Upload limit(Unit: Kbps), this field is required when parameter [upLimitEnable] is true. | [optional] 
**UpLimitEnable** | Pointer to **bool** | Whether to enable upload limit | [optional] 

## Methods

### NewRateLimitProfileOpenApiVO

`func NewRateLimitProfileOpenApiVO() *RateLimitProfileOpenApiVO`

NewRateLimitProfileOpenApiVO instantiates a new RateLimitProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitProfileOpenApiVOWithDefaults

`func NewRateLimitProfileOpenApiVOWithDefaults() *RateLimitProfileOpenApiVO`

NewRateLimitProfileOpenApiVOWithDefaults instantiates a new RateLimitProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultProfile

`func (o *RateLimitProfileOpenApiVO) GetDefaultProfile() bool`

GetDefaultProfile returns the DefaultProfile field if non-nil, zero value otherwise.

### GetDefaultProfileOk

`func (o *RateLimitProfileOpenApiVO) GetDefaultProfileOk() (*bool, bool)`

GetDefaultProfileOk returns a tuple with the DefaultProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfile

`func (o *RateLimitProfileOpenApiVO) SetDefaultProfile(v bool)`

SetDefaultProfile sets DefaultProfile field to given value.

### HasDefaultProfile

`func (o *RateLimitProfileOpenApiVO) HasDefaultProfile() bool`

HasDefaultProfile returns a boolean if a field has been set.

### GetDownLimit

`func (o *RateLimitProfileOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *RateLimitProfileOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *RateLimitProfileOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *RateLimitProfileOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *RateLimitProfileOpenApiVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *RateLimitProfileOpenApiVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *RateLimitProfileOpenApiVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.

### HasDownLimitEnable

`func (o *RateLimitProfileOpenApiVO) HasDownLimitEnable() bool`

HasDownLimitEnable returns a boolean if a field has been set.

### GetName

`func (o *RateLimitProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateLimitProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateLimitProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RateLimitProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfileId

`func (o *RateLimitProfileOpenApiVO) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *RateLimitProfileOpenApiVO) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *RateLimitProfileOpenApiVO) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *RateLimitProfileOpenApiVO) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetUpLimit

`func (o *RateLimitProfileOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *RateLimitProfileOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *RateLimitProfileOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *RateLimitProfileOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *RateLimitProfileOpenApiVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *RateLimitProfileOpenApiVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *RateLimitProfileOpenApiVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.

### HasUpLimitEnable

`func (o *RateLimitProfileOpenApiVO) HasUpLimitEnable() bool`

HasUpLimitEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


