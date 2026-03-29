# ClientLockToApSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aps** | Pointer to [**[]ApBriefInfoVO**](ApBriefInfoVO.md) | AP name and MAC info list. | [optional] 
**Enable** | Pointer to **bool** | Lock to AP enable. | [optional] 

## Methods

### NewClientLockToApSetting

`func NewClientLockToApSetting() *ClientLockToApSetting`

NewClientLockToApSetting instantiates a new ClientLockToApSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientLockToApSettingWithDefaults

`func NewClientLockToApSettingWithDefaults() *ClientLockToApSetting`

NewClientLockToApSettingWithDefaults instantiates a new ClientLockToApSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAps

`func (o *ClientLockToApSetting) GetAps() []ApBriefInfoVO`

GetAps returns the Aps field if non-nil, zero value otherwise.

### GetApsOk

`func (o *ClientLockToApSetting) GetApsOk() (*[]ApBriefInfoVO, bool)`

GetApsOk returns a tuple with the Aps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAps

`func (o *ClientLockToApSetting) SetAps(v []ApBriefInfoVO)`

SetAps sets Aps field to given value.

### HasAps

`func (o *ClientLockToApSetting) HasAps() bool`

HasAps returns a boolean if a field has been set.

### GetEnable

`func (o *ClientLockToApSetting) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *ClientLockToApSetting) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *ClientLockToApSetting) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *ClientLockToApSetting) HasEnable() bool`

HasEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


