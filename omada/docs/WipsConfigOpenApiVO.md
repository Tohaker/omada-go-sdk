# WipsConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeauthEn** | Pointer to **bool** | Wireless IPS deauthenticate config status; true:enable, false:disable. | [optional] 
**DynamicEn** | Pointer to **bool** | Wireless IPS dynamic block list config status; true:enable, false:disable. | [optional] 
**LockTime** | Pointer to **int32** | Wireless IPS device locking duration config status; It should be within the range of 300–36000; this field is required when parameter [dynamicEn] is true. | [optional] 
**Status** | Pointer to **bool** | Wireless IPS config status; true:enable, false:disable. | [optional] 

## Methods

### NewWipsConfigOpenApiVO

`func NewWipsConfigOpenApiVO() *WipsConfigOpenApiVO`

NewWipsConfigOpenApiVO instantiates a new WipsConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWipsConfigOpenApiVOWithDefaults

`func NewWipsConfigOpenApiVOWithDefaults() *WipsConfigOpenApiVO`

NewWipsConfigOpenApiVOWithDefaults instantiates a new WipsConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeauthEn

`func (o *WipsConfigOpenApiVO) GetDeauthEn() bool`

GetDeauthEn returns the DeauthEn field if non-nil, zero value otherwise.

### GetDeauthEnOk

`func (o *WipsConfigOpenApiVO) GetDeauthEnOk() (*bool, bool)`

GetDeauthEnOk returns a tuple with the DeauthEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeauthEn

`func (o *WipsConfigOpenApiVO) SetDeauthEn(v bool)`

SetDeauthEn sets DeauthEn field to given value.

### HasDeauthEn

`func (o *WipsConfigOpenApiVO) HasDeauthEn() bool`

HasDeauthEn returns a boolean if a field has been set.

### GetDynamicEn

`func (o *WipsConfigOpenApiVO) GetDynamicEn() bool`

GetDynamicEn returns the DynamicEn field if non-nil, zero value otherwise.

### GetDynamicEnOk

`func (o *WipsConfigOpenApiVO) GetDynamicEnOk() (*bool, bool)`

GetDynamicEnOk returns a tuple with the DynamicEn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicEn

`func (o *WipsConfigOpenApiVO) SetDynamicEn(v bool)`

SetDynamicEn sets DynamicEn field to given value.

### HasDynamicEn

`func (o *WipsConfigOpenApiVO) HasDynamicEn() bool`

HasDynamicEn returns a boolean if a field has been set.

### GetLockTime

`func (o *WipsConfigOpenApiVO) GetLockTime() int32`

GetLockTime returns the LockTime field if non-nil, zero value otherwise.

### GetLockTimeOk

`func (o *WipsConfigOpenApiVO) GetLockTimeOk() (*int32, bool)`

GetLockTimeOk returns a tuple with the LockTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockTime

`func (o *WipsConfigOpenApiVO) SetLockTime(v int32)`

SetLockTime sets LockTime field to given value.

### HasLockTime

`func (o *WipsConfigOpenApiVO) HasLockTime() bool`

HasLockTime returns a boolean if a field has been set.

### GetStatus

`func (o *WipsConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WipsConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WipsConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WipsConfigOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


