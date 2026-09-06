# Dot1xBasicInfoEapOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | EAP 802.1x function enable status | 
**GuestVlanSetting** | Pointer to [**Dot1xGuestVlanSettingOpenApiVO**](Dot1xGuestVlanSettingOpenApiVO.md) |  | [optional] 
**MacFormat** | Pointer to **int32** | Format of the MAC address. MacFormat should be a value as follows: 0: aabbccddeeff, 1: aa-bb-cc-dd-ee-ff, 2: aa:bb:cc:dd:ee:ff, 3: AABBCCDDEEFF, 4: AA-BB-CC-DD-EE-FF, 5: AA:BB:CC:DD:EE:FF | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new RADIUS profile&#39; (&#39;Create a new RADIUS profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get RADIUS profile list&#39; (&#39;Get RADIUS profile template list&#39;) interface | [optional] 

## Methods

### NewDot1xBasicInfoEapOpenApiVO

`func NewDot1xBasicInfoEapOpenApiVO(enable bool, ) *Dot1xBasicInfoEapOpenApiVO`

NewDot1xBasicInfoEapOpenApiVO instantiates a new Dot1xBasicInfoEapOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xBasicInfoEapOpenApiVOWithDefaults

`func NewDot1xBasicInfoEapOpenApiVOWithDefaults() *Dot1xBasicInfoEapOpenApiVO`

NewDot1xBasicInfoEapOpenApiVOWithDefaults instantiates a new Dot1xBasicInfoEapOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *Dot1xBasicInfoEapOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Dot1xBasicInfoEapOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Dot1xBasicInfoEapOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetGuestVlanSetting

`func (o *Dot1xBasicInfoEapOpenApiVO) GetGuestVlanSetting() Dot1xGuestVlanSettingOpenApiVO`

GetGuestVlanSetting returns the GuestVlanSetting field if non-nil, zero value otherwise.

### GetGuestVlanSettingOk

`func (o *Dot1xBasicInfoEapOpenApiVO) GetGuestVlanSettingOk() (*Dot1xGuestVlanSettingOpenApiVO, bool)`

GetGuestVlanSettingOk returns a tuple with the GuestVlanSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestVlanSetting

`func (o *Dot1xBasicInfoEapOpenApiVO) SetGuestVlanSetting(v Dot1xGuestVlanSettingOpenApiVO)`

SetGuestVlanSetting sets GuestVlanSetting field to given value.

### HasGuestVlanSetting

`func (o *Dot1xBasicInfoEapOpenApiVO) HasGuestVlanSetting() bool`

HasGuestVlanSetting returns a boolean if a field has been set.

### GetMacFormat

`func (o *Dot1xBasicInfoEapOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *Dot1xBasicInfoEapOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *Dot1xBasicInfoEapOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *Dot1xBasicInfoEapOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *Dot1xBasicInfoEapOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *Dot1xBasicInfoEapOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *Dot1xBasicInfoEapOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *Dot1xBasicInfoEapOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


