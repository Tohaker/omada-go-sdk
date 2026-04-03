# ClientBatchSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpSetting** | Pointer to [**ClientBatchIPSetting**](ClientBatchIPSetting.md) |  | [optional] 
**LockToAp** | Pointer to [**ClientLockToApMacListSetting**](ClientLockToApMacListSetting.md) |  | [optional] 
**MacList** | Pointer to **[]string** | List of clients&#39; mac. | [optional] 
**RateLimit** | Pointer to [**ClientRateLimitSetting**](ClientRateLimitSetting.md) |  | [optional] 

## Methods

### NewClientBatchSetting

`func NewClientBatchSetting() *ClientBatchSetting`

NewClientBatchSetting instantiates a new ClientBatchSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientBatchSettingWithDefaults

`func NewClientBatchSettingWithDefaults() *ClientBatchSetting`

NewClientBatchSettingWithDefaults instantiates a new ClientBatchSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpSetting

`func (o *ClientBatchSetting) GetIpSetting() ClientBatchIPSetting`

GetIpSetting returns the IpSetting field if non-nil, zero value otherwise.

### GetIpSettingOk

`func (o *ClientBatchSetting) GetIpSettingOk() (*ClientBatchIPSetting, bool)`

GetIpSettingOk returns a tuple with the IpSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpSetting

`func (o *ClientBatchSetting) SetIpSetting(v ClientBatchIPSetting)`

SetIpSetting sets IpSetting field to given value.

### HasIpSetting

`func (o *ClientBatchSetting) HasIpSetting() bool`

HasIpSetting returns a boolean if a field has been set.

### GetLockToAp

`func (o *ClientBatchSetting) GetLockToAp() ClientLockToApMacListSetting`

GetLockToAp returns the LockToAp field if non-nil, zero value otherwise.

### GetLockToApOk

`func (o *ClientBatchSetting) GetLockToApOk() (*ClientLockToApMacListSetting, bool)`

GetLockToApOk returns a tuple with the LockToAp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockToAp

`func (o *ClientBatchSetting) SetLockToAp(v ClientLockToApMacListSetting)`

SetLockToAp sets LockToAp field to given value.

### HasLockToAp

`func (o *ClientBatchSetting) HasLockToAp() bool`

HasLockToAp returns a boolean if a field has been set.

### GetMacList

`func (o *ClientBatchSetting) GetMacList() []string`

GetMacList returns the MacList field if non-nil, zero value otherwise.

### GetMacListOk

`func (o *ClientBatchSetting) GetMacListOk() (*[]string, bool)`

GetMacListOk returns a tuple with the MacList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacList

`func (o *ClientBatchSetting) SetMacList(v []string)`

SetMacList sets MacList field to given value.

### HasMacList

`func (o *ClientBatchSetting) HasMacList() bool`

HasMacList returns a boolean if a field has been set.

### GetRateLimit

`func (o *ClientBatchSetting) GetRateLimit() ClientRateLimitSetting`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ClientBatchSetting) GetRateLimitOk() (*ClientRateLimitSetting, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ClientBatchSetting) SetRateLimit(v ClientRateLimitSetting)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ClientBatchSetting) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


