# WlanGroupOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | site Id | [optional] 
**Clone** | Pointer to **bool** | Whether it is cloned | [optional] 
**CloneWlanId** | Pointer to **string** | cloneWlanId | [optional] 
**Name** | Pointer to **string** | WLAN group name should contain 1 to 128 characters. | [optional] 
**Primary** | Pointer to **bool** | Whether it is the default WLAN group | [optional] 
**Resource** | Pointer to **int32** | resource. 0 is new created, 1 is from template, 2 is override template. | [optional] 
**WlanId** | Pointer to **string** | WLAN group ID | [optional] 

## Methods

### NewWlanGroupOpenApiVO

`func NewWlanGroupOpenApiVO() *WlanGroupOpenApiVO`

NewWlanGroupOpenApiVO instantiates a new WlanGroupOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWlanGroupOpenApiVOWithDefaults

`func NewWlanGroupOpenApiVOWithDefaults() *WlanGroupOpenApiVO`

NewWlanGroupOpenApiVOWithDefaults instantiates a new WlanGroupOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *WlanGroupOpenApiVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *WlanGroupOpenApiVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *WlanGroupOpenApiVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *WlanGroupOpenApiVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetClone

`func (o *WlanGroupOpenApiVO) GetClone() bool`

GetClone returns the Clone field if non-nil, zero value otherwise.

### GetCloneOk

`func (o *WlanGroupOpenApiVO) GetCloneOk() (*bool, bool)`

GetCloneOk returns a tuple with the Clone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClone

`func (o *WlanGroupOpenApiVO) SetClone(v bool)`

SetClone sets Clone field to given value.

### HasClone

`func (o *WlanGroupOpenApiVO) HasClone() bool`

HasClone returns a boolean if a field has been set.

### GetCloneWlanId

`func (o *WlanGroupOpenApiVO) GetCloneWlanId() string`

GetCloneWlanId returns the CloneWlanId field if non-nil, zero value otherwise.

### GetCloneWlanIdOk

`func (o *WlanGroupOpenApiVO) GetCloneWlanIdOk() (*string, bool)`

GetCloneWlanIdOk returns a tuple with the CloneWlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneWlanId

`func (o *WlanGroupOpenApiVO) SetCloneWlanId(v string)`

SetCloneWlanId sets CloneWlanId field to given value.

### HasCloneWlanId

`func (o *WlanGroupOpenApiVO) HasCloneWlanId() bool`

HasCloneWlanId returns a boolean if a field has been set.

### GetName

`func (o *WlanGroupOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WlanGroupOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WlanGroupOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WlanGroupOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPrimary

`func (o *WlanGroupOpenApiVO) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *WlanGroupOpenApiVO) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *WlanGroupOpenApiVO) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *WlanGroupOpenApiVO) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.

### GetResource

`func (o *WlanGroupOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WlanGroupOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WlanGroupOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *WlanGroupOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetWlanId

`func (o *WlanGroupOpenApiVO) GetWlanId() string`

GetWlanId returns the WlanId field if non-nil, zero value otherwise.

### GetWlanIdOk

`func (o *WlanGroupOpenApiVO) GetWlanIdOk() (*string, bool)`

GetWlanIdOk returns a tuple with the WlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanId

`func (o *WlanGroupOpenApiVO) SetWlanId(v string)`

SetWlanId sets WlanId field to given value.

### HasWlanId

`func (o *WlanGroupOpenApiVO) HasWlanId() bool`

HasWlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


