# WanSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UsbPortsConfig** | Pointer to [**[]UsbLteSettingOpenApiVO**](UsbLteSettingOpenApiVO.md) | USB lte ports config | [optional] 
**WanPortsConfig** | Pointer to [**[]WanPortSettingOpenApiVO**](WanPortSettingOpenApiVO.md) | WAN ports config | [optional] 

## Methods

### NewWanSettingOpenApiVO

`func NewWanSettingOpenApiVO() *WanSettingOpenApiVO`

NewWanSettingOpenApiVO instantiates a new WanSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanSettingOpenApiVOWithDefaults

`func NewWanSettingOpenApiVOWithDefaults() *WanSettingOpenApiVO`

NewWanSettingOpenApiVOWithDefaults instantiates a new WanSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsbPortsConfig

`func (o *WanSettingOpenApiVO) GetUsbPortsConfig() []UsbLteSettingOpenApiVO`

GetUsbPortsConfig returns the UsbPortsConfig field if non-nil, zero value otherwise.

### GetUsbPortsConfigOk

`func (o *WanSettingOpenApiVO) GetUsbPortsConfigOk() (*[]UsbLteSettingOpenApiVO, bool)`

GetUsbPortsConfigOk returns a tuple with the UsbPortsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsbPortsConfig

`func (o *WanSettingOpenApiVO) SetUsbPortsConfig(v []UsbLteSettingOpenApiVO)`

SetUsbPortsConfig sets UsbPortsConfig field to given value.

### HasUsbPortsConfig

`func (o *WanSettingOpenApiVO) HasUsbPortsConfig() bool`

HasUsbPortsConfig returns a boolean if a field has been set.

### GetWanPortsConfig

`func (o *WanSettingOpenApiVO) GetWanPortsConfig() []WanPortSettingOpenApiVO`

GetWanPortsConfig returns the WanPortsConfig field if non-nil, zero value otherwise.

### GetWanPortsConfigOk

`func (o *WanSettingOpenApiVO) GetWanPortsConfigOk() (*[]WanPortSettingOpenApiVO, bool)`

GetWanPortsConfigOk returns a tuple with the WanPortsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortsConfig

`func (o *WanSettingOpenApiVO) SetWanPortsConfig(v []WanPortSettingOpenApiVO)`

SetWanPortsConfig sets WanPortsConfig field to given value.

### HasWanPortsConfig

`func (o *WanSettingOpenApiVO) HasWanPortsConfig() bool`

HasWanPortsConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


