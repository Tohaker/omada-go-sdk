# UpdateWipsConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeauthEn** | **bool** | Wireless IPS deauthenticate config status; true:enable, false:disable. | 
**DynamicEn** | **bool** | Wireless IPS dynamic block list config status; true:enable, false:disable. | 
**LockTime** | Pointer to **int32** | Wireless IPS device locking duration config status; It should be within the range of 300–36000; this field is required when parameter [dynamicEn] is true. | [optional] 
**Status** | **bool** | Wireless IPS config status; true:enable, false:disable. | 

## Methods

### NewUpdateWipsConfigOpenApiVO

`func NewUpdateWipsConfigOpenApiVO(deauthEn bool, dynamicEn bool, status bool, ) *UpdateWipsConfigOpenApiVO`

NewUpdateWipsConfigOpenApiVO instantiates a new UpdateWipsConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWipsConfigOpenApiVOWithDefaults

`func NewUpdateWipsConfigOpenApiVOWithDefaults() *UpdateWipsConfigOpenApiVO`

NewUpdateWipsConfigOpenApiVOWithDefaults instantiates a new UpdateWipsConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeauthEn

`func (o *UpdateWipsConfigOpenApiVO) GetDeauthEn() bool`

GetDeauthEn returns the DeauthEn field if non-nil, zero value otherwise.

### GetDeauthEnOk

`func (o *UpdateWipsConfigOpenApiVO) GetDeauthEnOk() (*bool, bool)`

GetDeauthEnOk returns a tuple with the DeauthEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeauthEn

`func (o *UpdateWipsConfigOpenApiVO) SetDeauthEn(v bool)`

SetDeauthEn sets DeauthEn field to given value.


### GetDynamicEn

`func (o *UpdateWipsConfigOpenApiVO) GetDynamicEn() bool`

GetDynamicEn returns the DynamicEn field if non-nil, zero value otherwise.

### GetDynamicEnOk

`func (o *UpdateWipsConfigOpenApiVO) GetDynamicEnOk() (*bool, bool)`

GetDynamicEnOk returns a tuple with the DynamicEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicEn

`func (o *UpdateWipsConfigOpenApiVO) SetDynamicEn(v bool)`

SetDynamicEn sets DynamicEn field to given value.


### GetLockTime

`func (o *UpdateWipsConfigOpenApiVO) GetLockTime() int32`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *UpdateWipsConfigOpenApiVO) GetLockTimeOk() (*int32, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *UpdateWipsConfigOpenApiVO) SetLockTime(v int32)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *UpdateWipsConfigOpenApiVO) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateWipsConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateWipsConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateWipsConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


