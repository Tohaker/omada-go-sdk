# GatewayGeneralConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedSetting** | **int32** | Led setting should be a value as follows: 0:off; 1:on; 2:Use Site Settings | 
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name should contain 1 to 128 characters. | [optional] 
**RememberDevice** | Pointer to **int32** |  | [optional] 
**Snmp** | Pointer to [**OsgSnmpOpenApiVO**](OsgSnmpOpenApiVO.md) |  | [optional] 
**TagIds** | Pointer to **[]string** | Tag IDs | [optional] 

## Methods

### NewGatewayGeneralConfig

`func NewGatewayGeneralConfig(ledSetting int32, ) *GatewayGeneralConfig`

NewGatewayGeneralConfig instantiates a new GatewayGeneralConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayGeneralConfigWithDefaults

`func NewGatewayGeneralConfigWithDefaults() *GatewayGeneralConfig`

NewGatewayGeneralConfigWithDefaults instantiates a new GatewayGeneralConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedSetting

`func (o *GatewayGeneralConfig) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *GatewayGeneralConfig) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *GatewayGeneralConfig) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.


### GetLocation

`func (o *GatewayGeneralConfig) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GatewayGeneralConfig) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GatewayGeneralConfig) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GatewayGeneralConfig) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *GatewayGeneralConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayGeneralConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayGeneralConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayGeneralConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRememberDevice

`func (o *GatewayGeneralConfig) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *GatewayGeneralConfig) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *GatewayGeneralConfig) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *GatewayGeneralConfig) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetSnmp

`func (o *GatewayGeneralConfig) GetSnmp() OsgSnmpOpenApiVO`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *GatewayGeneralConfig) GetSnmpOk() (*OsgSnmpOpenApiVO, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *GatewayGeneralConfig) SetSnmp(v OsgSnmpOpenApiVO)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *GatewayGeneralConfig) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetTagIds

`func (o *GatewayGeneralConfig) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *GatewayGeneralConfig) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *GatewayGeneralConfig) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *GatewayGeneralConfig) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


