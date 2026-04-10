# ApManagementSsidConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Broadcast** | **bool** | SSID broadcast config status. True: enable, false: disable. | 
**EntSetting** | Pointer to [**ApMgtSsidEnterpriseSettingOpenApiVO**](ApMgtSsidEnterpriseSettingOpenApiVO.md) |  | [optional] 
**Name** | **string** | SSID name. It should contain 1 to 32 UTF-8 characters. | 
**PskSetting** | Pointer to [**ApMgtSsidPskSettingOpenApiVO**](ApMgtSsidPskSettingOpenApiVO.md) |  | [optional] 
**Security** | **int32** | SSID security mode; Security should be a value as follows: 0: None; 2: WPA-Enterprise; 3: WPA-Personal.  | 
**Status** | **bool** | SSID config status. True: enable, false: disable. | 
**VlanSetting** | Pointer to [**ApMgtSsidVlanSettingOpenApiVO**](ApMgtSsidVlanSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewApManagementSsidConfig

`func NewApManagementSsidConfig(broadcast bool, name string, security int32, status bool, ) *ApManagementSsidConfig`

NewApManagementSsidConfig instantiates a new ApManagementSsidConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApManagementSsidConfigWithDefaults

`func NewApManagementSsidConfigWithDefaults() *ApManagementSsidConfig`

NewApManagementSsidConfigWithDefaults instantiates a new ApManagementSsidConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroadcast

`func (o *ApManagementSsidConfig) GetBroadcast() bool`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *ApManagementSsidConfig) GetBroadcastOk() (*bool, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *ApManagementSsidConfig) SetBroadcast(v bool)`

SetBroadcast sets Broadcast field to given value.


### GetEntSetting

`func (o *ApManagementSsidConfig) GetEntSetting() ApMgtSsidEnterpriseSettingOpenApiVO`

GetEntSetting returns the EntSetting field if non-nil, zero value otherwise.

### GetEntSettingOk

`func (o *ApManagementSsidConfig) GetEntSettingOk() (*ApMgtSsidEnterpriseSettingOpenApiVO, bool)`

GetEntSettingOk returns a tuple with the EntSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntSetting

`func (o *ApManagementSsidConfig) SetEntSetting(v ApMgtSsidEnterpriseSettingOpenApiVO)`

SetEntSetting sets EntSetting field to given value.

### HasEntSetting

`func (o *ApManagementSsidConfig) HasEntSetting() bool`

HasEntSetting returns a boolean if a field has been set.

### GetName

`func (o *ApManagementSsidConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApManagementSsidConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApManagementSsidConfig) SetName(v string)`

SetName sets Name field to given value.


### GetPskSetting

`func (o *ApManagementSsidConfig) GetPskSetting() ApMgtSsidPskSettingOpenApiVO`

GetPskSetting returns the PskSetting field if non-nil, zero value otherwise.

### GetPskSettingOk

`func (o *ApManagementSsidConfig) GetPskSettingOk() (*ApMgtSsidPskSettingOpenApiVO, bool)`

GetPskSettingOk returns a tuple with the PskSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPskSetting

`func (o *ApManagementSsidConfig) SetPskSetting(v ApMgtSsidPskSettingOpenApiVO)`

SetPskSetting sets PskSetting field to given value.

### HasPskSetting

`func (o *ApManagementSsidConfig) HasPskSetting() bool`

HasPskSetting returns a boolean if a field has been set.

### GetSecurity

`func (o *ApManagementSsidConfig) GetSecurity() int32`

GetSecurity returns the Security field if non-nil, zero value otherwise.

### GetSecurityOk

`func (o *ApManagementSsidConfig) GetSecurityOk() (*int32, bool)`

GetSecurityOk returns a tuple with the Security field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurity

`func (o *ApManagementSsidConfig) SetSecurity(v int32)`

SetSecurity sets Security field to given value.


### GetStatus

`func (o *ApManagementSsidConfig) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApManagementSsidConfig) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApManagementSsidConfig) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetVlanSetting

`func (o *ApManagementSsidConfig) GetVlanSetting() ApMgtSsidVlanSettingOpenApiVO`

GetVlanSetting returns the VlanSetting field if non-nil, zero value otherwise.

### GetVlanSettingOk

`func (o *ApManagementSsidConfig) GetVlanSettingOk() (*ApMgtSsidVlanSettingOpenApiVO, bool)`

GetVlanSettingOk returns a tuple with the VlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanSetting

`func (o *ApManagementSsidConfig) SetVlanSetting(v ApMgtSsidVlanSettingOpenApiVO)`

SetVlanSetting sets VlanSetting field to given value.

### HasVlanSetting

`func (o *ApManagementSsidConfig) HasVlanSetting() bool`

HasVlanSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


