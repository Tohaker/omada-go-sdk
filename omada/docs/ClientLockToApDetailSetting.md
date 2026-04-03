# ClientLockToApDetailSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aps** | Pointer to [**[]ApBriefInfoVO**](ApBriefInfoVO.md) | AP name and MAC info list. | [optional] 
**Enable** | Pointer to **bool** | Lock to AP enable. | [optional] 

## Methods

### NewClientLockToApDetailSetting

`func NewClientLockToApDetailSetting() *ClientLockToApDetailSetting`

NewClientLockToApDetailSetting instantiates a new ClientLockToApDetailSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLockToApDetailSettingWithDefaults

`func NewClientLockToApDetailSettingWithDefaults() *ClientLockToApDetailSetting`

NewClientLockToApDetailSettingWithDefaults instantiates a new ClientLockToApDetailSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAps

`func (o *ClientLockToApDetailSetting) GetAps() []ApBriefInfoVO`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *ClientLockToApDetailSetting) GetApsOk() (*[]ApBriefInfoVO, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *ClientLockToApDetailSetting) SetAps(v []ApBriefInfoVO)`

SetAps sets Aps field to given value.

### HasAps

`func (o *ClientLockToApDetailSetting) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetEnable

`func (o *ClientLockToApDetailSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ClientLockToApDetailSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ClientLockToApDetailSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ClientLockToApDetailSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


