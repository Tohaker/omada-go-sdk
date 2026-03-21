# UpdateRateLimitProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int64** | Download limit(Unit: Kbps), this field is required when parameter [downLimitEnable] is true. | [optional] 
**DownLimitEnable** | **bool** | Whether to enable download limit | 
**Name** | **string** | Rate limit profile name should contain 1 to 64 characters. | 
**UpLimit** | Pointer to **int64** | Upload limit(Unit: Kbps), this field is required when parameter [upLimitEnable] is true. | [optional] 
**UpLimitEnable** | **bool** | Whether to enable upload limit | 

## Methods

### NewUpdateRateLimitProfileOpenApiVO

`func NewUpdateRateLimitProfileOpenApiVO(downLimitEnable bool, name string, upLimitEnable bool, ) *UpdateRateLimitProfileOpenApiVO`

NewUpdateRateLimitProfileOpenApiVO instantiates a new UpdateRateLimitProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRateLimitProfileOpenApiVOWithDefaults

`func NewUpdateRateLimitProfileOpenApiVOWithDefaults() *UpdateRateLimitProfileOpenApiVO`

NewUpdateRateLimitProfileOpenApiVOWithDefaults instantiates a new UpdateRateLimitProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *UpdateRateLimitProfileOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *UpdateRateLimitProfileOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *UpdateRateLimitProfileOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *UpdateRateLimitProfileOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *UpdateRateLimitProfileOpenApiVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *UpdateRateLimitProfileOpenApiVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *UpdateRateLimitProfileOpenApiVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.


### GetName

`func (o *UpdateRateLimitProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRateLimitProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRateLimitProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetUpLimit

`func (o *UpdateRateLimitProfileOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *UpdateRateLimitProfileOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *UpdateRateLimitProfileOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *UpdateRateLimitProfileOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *UpdateRateLimitProfileOpenApiVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *UpdateRateLimitProfileOpenApiVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *UpdateRateLimitProfileOpenApiVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


