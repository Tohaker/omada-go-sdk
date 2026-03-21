# RateLimitProfileVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Site ID | [optional] 
**DownLimit** | Pointer to **int64** | Down limit. Unit is Kbps, integer from 1 to 10485760. | [optional] 
**DownLimitEnable** | **bool** | Down limit enable. | 
**Id** | Pointer to **string** | Rate limit profile id. | [optional] 
**IsDefault** | Pointer to **bool** | Whether it is default profile. | [optional] 
**Name** | **string** | Rate limit profile name. | 
**Resource** | Pointer to **int32** |  | [optional] 
**UpLimit** | Pointer to **int64** | Up limit. Unit is Kbps, integer from 1 to 10485760. | [optional] 
**UpLimitEnable** | **bool** | Up limit enable. | 

## Methods

### NewRateLimitProfileVO

`func NewRateLimitProfileVO(downLimitEnable bool, name string, upLimitEnable bool, ) *RateLimitProfileVO`

NewRateLimitProfileVO instantiates a new RateLimitProfileVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitProfileVOWithDefaults

`func NewRateLimitProfileVOWithDefaults() *RateLimitProfileVO`

NewRateLimitProfileVOWithDefaults instantiates a new RateLimitProfileVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *RateLimitProfileVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *RateLimitProfileVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *RateLimitProfileVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *RateLimitProfileVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetDownLimit

`func (o *RateLimitProfileVO) GetDownLimit() int64`

GetDownLimit returns the DownLimit field if non-nil, zero value otherwise.

### GetDownLimitOk

`func (o *RateLimitProfileVO) GetDownLimitOk() (*int64, bool)`

GetDownLimitOk returns a tuple with the DownLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimit

`func (o *RateLimitProfileVO) SetDownLimit(v int64)`

SetDownLimit sets DownLimit field to given value.

### HasDownLimit

`func (o *RateLimitProfileVO) HasDownLimit() bool`

HasDownLimit returns a boolean if a field has been set.

### GetDownLimitEnable

`func (o *RateLimitProfileVO) GetDownLimitEnable() bool`

GetDownLimitEnable returns the DownLimitEnable field if non-nil, zero value otherwise.

### GetDownLimitEnableOk

`func (o *RateLimitProfileVO) GetDownLimitEnableOk() (*bool, bool)`

GetDownLimitEnableOk returns a tuple with the DownLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownLimitEnable

`func (o *RateLimitProfileVO) SetDownLimitEnable(v bool)`

SetDownLimitEnable sets DownLimitEnable field to given value.


### GetId

`func (o *RateLimitProfileVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RateLimitProfileVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RateLimitProfileVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RateLimitProfileVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *RateLimitProfileVO) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *RateLimitProfileVO) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *RateLimitProfileVO) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *RateLimitProfileVO) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *RateLimitProfileVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RateLimitProfileVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RateLimitProfileVO) SetName(v string)`

SetName sets Name field to given value.


### GetResource

`func (o *RateLimitProfileVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RateLimitProfileVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RateLimitProfileVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RateLimitProfileVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetUpLimit

`func (o *RateLimitProfileVO) GetUpLimit() int64`

GetUpLimit returns the UpLimit field if non-nil, zero value otherwise.

### GetUpLimitOk

`func (o *RateLimitProfileVO) GetUpLimitOk() (*int64, bool)`

GetUpLimitOk returns a tuple with the UpLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimit

`func (o *RateLimitProfileVO) SetUpLimit(v int64)`

SetUpLimit sets UpLimit field to given value.

### HasUpLimit

`func (o *RateLimitProfileVO) HasUpLimit() bool`

HasUpLimit returns a boolean if a field has been set.

### GetUpLimitEnable

`func (o *RateLimitProfileVO) GetUpLimitEnable() bool`

GetUpLimitEnable returns the UpLimitEnable field if non-nil, zero value otherwise.

### GetUpLimitEnableOk

`func (o *RateLimitProfileVO) GetUpLimitEnableOk() (*bool, bool)`

GetUpLimitEnableOk returns a tuple with the UpLimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpLimitEnable

`func (o *RateLimitProfileVO) SetUpLimitEnable(v bool)`

SetUpLimitEnable sets UpLimitEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


