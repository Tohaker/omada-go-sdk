# OsgGeneralConfigOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LedSetting** | **int32** | Led setting should be a value as follows: 0:off; 1:on; 2:Use Site Settings | 
**Location** | Pointer to [**DeviceLocationDetailOpenApiVO**](DeviceLocationDetailOpenApiVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name should contain 1 to 128 characters. This subsection is deprecated. | [optional] 
**RememberDevice** | Pointer to **int32** |  | [optional] 
**Snmp** | Pointer to [**OsgSnmpOpenApiVO**](OsgSnmpOpenApiVO.md) |  | [optional] 
**TagIds** | Pointer to **[]string** | Tag IDs. This subsection is deprecated. | [optional] 

## Methods

### NewOsgGeneralConfigOpenApiV2VO

`func NewOsgGeneralConfigOpenApiV2VO(ledSetting int32, ) *OsgGeneralConfigOpenApiV2VO`

NewOsgGeneralConfigOpenApiV2VO instantiates a new OsgGeneralConfigOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgGeneralConfigOpenApiV2VOWithDefaults

`func NewOsgGeneralConfigOpenApiV2VOWithDefaults() *OsgGeneralConfigOpenApiV2VO`

NewOsgGeneralConfigOpenApiV2VOWithDefaults instantiates a new OsgGeneralConfigOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLedSetting

`func (o *OsgGeneralConfigOpenApiV2VO) GetLedSetting() int32`

GetLedSetting returns the LedSetting field if non-nil, zero value otherwise.

### GetLedSettingOk

`func (o *OsgGeneralConfigOpenApiV2VO) GetLedSettingOk() (*int32, bool)`

GetLedSettingOk returns a tuple with the LedSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLedSetting

`func (o *OsgGeneralConfigOpenApiV2VO) SetLedSetting(v int32)`

SetLedSetting sets LedSetting field to given value.


### GetLocation

`func (o *OsgGeneralConfigOpenApiV2VO) GetLocation() DeviceLocationDetailOpenApiVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OsgGeneralConfigOpenApiV2VO) GetLocationOk() (*DeviceLocationDetailOpenApiVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OsgGeneralConfigOpenApiV2VO) SetLocation(v DeviceLocationDetailOpenApiVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OsgGeneralConfigOpenApiV2VO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *OsgGeneralConfigOpenApiV2VO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OsgGeneralConfigOpenApiV2VO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OsgGeneralConfigOpenApiV2VO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OsgGeneralConfigOpenApiV2VO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRememberDevice

`func (o *OsgGeneralConfigOpenApiV2VO) GetRememberDevice() int32`

GetRememberDevice returns the RememberDevice field if non-nil, zero value otherwise.

### GetRememberDeviceOk

`func (o *OsgGeneralConfigOpenApiV2VO) GetRememberDeviceOk() (*int32, bool)`

GetRememberDeviceOk returns a tuple with the RememberDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRememberDevice

`func (o *OsgGeneralConfigOpenApiV2VO) SetRememberDevice(v int32)`

SetRememberDevice sets RememberDevice field to given value.

### HasRememberDevice

`func (o *OsgGeneralConfigOpenApiV2VO) HasRememberDevice() bool`

HasRememberDevice returns a boolean if a field has been set.

### GetSnmp

`func (o *OsgGeneralConfigOpenApiV2VO) GetSnmp() OsgSnmpOpenApiVO`

GetSnmp returns the Snmp field if non-nil, zero value otherwise.

### GetSnmpOk

`func (o *OsgGeneralConfigOpenApiV2VO) GetSnmpOk() (*OsgSnmpOpenApiVO, bool)`

GetSnmpOk returns a tuple with the Snmp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnmp

`func (o *OsgGeneralConfigOpenApiV2VO) SetSnmp(v OsgSnmpOpenApiVO)`

SetSnmp sets Snmp field to given value.

### HasSnmp

`func (o *OsgGeneralConfigOpenApiV2VO) HasSnmp() bool`

HasSnmp returns a boolean if a field has been set.

### GetTagIds

`func (o *OsgGeneralConfigOpenApiV2VO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OsgGeneralConfigOpenApiV2VO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OsgGeneralConfigOpenApiV2VO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OsgGeneralConfigOpenApiV2VO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


