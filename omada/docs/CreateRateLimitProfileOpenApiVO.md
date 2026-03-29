# CreateRateLimitProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownLimit** | Pointer to **int64** | Download limit(Unit: Kbps), this field is required when parameter [downLimitEnable] is true; It should be within the range of 1–10485760. | [optional] 
**DownLimitEnable** | **bool** | Whether to enable download limit | 
**Name** | **string** | Rate limit profile name should contain 1 to 64 characters. | 
**UpLimit** | Pointer to **int64** | Upload limit(Unit: Kbps), this field is required when parameter [upLimitEnable] is true; It should be within the range of 1–10485760. | [optional] 
**UpLimitEnable** | **bool** | Whether to enable upload limit | 

## Methods

### NewCreateRateLimitProfileOpenApiVO

`func NewCreateRateLimitProfileOpenApiVO(downLimitEnable bool, name string, upLimitEnable bool, ) *CreateRateLimitProfileOpenApiVO`

NewCreateRateLimitProfileOpenApiVO instantiates a new CreateRateLimitProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRateLimitProfileOpenApiVOWithDefaults

`func NewCreateRateLimitProfileOpenApiVOWithDefaults() *CreateRateLimitProfileOpenApiVO`

NewCreateRateLimitProfileOpenApiVOWithDefaults instantiates a new CreateRateLimitProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownLimit

`func (o *CreateRateLimitProfileOpenApiVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *CreateRateLimitProfileOpenApiVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *CreateRateLimitProfileOpenApiVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *CreateRateLimitProfileOpenApiVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *CreateRateLimitProfileOpenApiVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *CreateRateLimitProfileOpenApiVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *CreateRateLimitProfileOpenApiVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.


### GetName

`func (o *CreateRateLimitProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRateLimitProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRateLimitProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetUpLimit

`func (o *CreateRateLimitProfileOpenApiVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *CreateRateLimitProfileOpenApiVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *CreateRateLimitProfileOpenApiVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *CreateRateLimitProfileOpenApiVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *CreateRateLimitProfileOpenApiVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *CreateRateLimitProfileOpenApiVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *CreateRateLimitProfileOpenApiVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


