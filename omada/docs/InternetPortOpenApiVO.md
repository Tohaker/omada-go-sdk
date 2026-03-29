# InternetPortOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LteWanSetting** | Pointer to [**LteWanSettingOpenApiVO**](LteWanSettingOpenApiVO.md) |  | [optional] 
**Type** | Pointer to **int32** | The Type of the Internet port as follows: 0: WAN; 1: USB; 2: LTE; 3: DSL. | [optional] 
**UsbLteSetting** | Pointer to [**UsbLteSettingOpenApiVO**](UsbLteSettingOpenApiVO.md) |  | [optional] 
**WanPortSetting** | Pointer to [**WanPortSettingOpenApiVO**](WanPortSettingOpenApiVO.md) |  | [optional] 

## Methods

### NewInternetPortOpenApiVO

`func NewInternetPortOpenApiVO() *InternetPortOpenApiVO`

NewInternetPortOpenApiVO instantiates a new InternetPortOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInternetPortOpenApiVOWithDefaults

`func NewInternetPortOpenApiVOWithDefaults() *InternetPortOpenApiVO`

NewInternetPortOpenApiVOWithDefaults instantiates a new InternetPortOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLteWanSetting

`func (o *InternetPortOpenApiVO) GetLteWanSetting() LteWanSettingOpenApiVO`

GetLteWanSetting returns the LteWanSetting field if non-nil, zero value otherwise.

### GetLteWanSettingOk

`func (o *InternetPortOpenApiVO) GetLteWanSettingOk() (*LteWanSettingOpenApiVO, bool)`

GetLteWanSettingOk returns a tuple with the LteWanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteWanSetting

`func (o *InternetPortOpenApiVO) SetLteWanSetting(v LteWanSettingOpenApiVO)`

SetLteWanSetting sets LteWanSetting field to given value.

### HasLteWanSetting

`func (o *InternetPortOpenApiVO) HasLteWanSetting() bool`

HasLteWanSetting returns a boolean if a field has been set.

### GetType

`func (o *InternetPortOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InternetPortOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InternetPortOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *InternetPortOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUsbLteSetting

`func (o *InternetPortOpenApiVO) GetUsbLteSetting() UsbLteSettingOpenApiVO`

GetUsbLteSetting returns the UsbLteSetting field if non-nil, zero value otherwise.

### GetUsbLteSettingOk

`func (o *InternetPortOpenApiVO) GetUsbLteSettingOk() (*UsbLteSettingOpenApiVO, bool)`

GetUsbLteSettingOk returns a tuple with the UsbLteSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbLteSetting

`func (o *InternetPortOpenApiVO) SetUsbLteSetting(v UsbLteSettingOpenApiVO)`

SetUsbLteSetting sets UsbLteSetting field to given value.

### HasUsbLteSetting

`func (o *InternetPortOpenApiVO) HasUsbLteSetting() bool`

HasUsbLteSetting returns a boolean if a field has been set.

### GetWanPortSetting

`func (o *InternetPortOpenApiVO) GetWanPortSetting() WanPortSettingOpenApiVO`

GetWanPortSetting returns the WanPortSetting field if non-nil, zero value otherwise.

### GetWanPortSettingOk

`func (o *InternetPortOpenApiVO) GetWanPortSettingOk() (*WanPortSettingOpenApiVO, bool)`

GetWanPortSettingOk returns a tuple with the WanPortSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortSetting

`func (o *InternetPortOpenApiVO) SetWanPortSetting(v WanPortSettingOpenApiVO)`

SetWanPortSetting sets WanPortSetting field to given value.

### HasWanPortSetting

`func (o *InternetPortOpenApiVO) HasWanPortSetting() bool`

HasWanPortSetting returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


