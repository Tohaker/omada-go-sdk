# Dot1xEapOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eaps** | Pointer to [**[]Dot1xEapSettingOpenApiVO**](Dot1xEapSettingOpenApiVO.md) | Enabled eap ports | [optional] 
**Enable** | **bool** | EAP 802.1x function enable status | 
**MacFormat** | Pointer to **int32** | Format of the MAC address. MacFormat should be a value as follows: 0: aabbccddeeff, 1: aa-bb-cc-dd-ee-ff, 2: aa:bb:cc:dd:ee:ff, 3: AABBCCDDEEFF, 4: AA-BB-CC-DD-EE-FF, 5: AA:BB:CC:DD:EE:FF | [optional] 
**RadiusProfileId** | Pointer to **string** | This field represents radius profile ID. Radius profile can be created using &#39;Create a new Radius profile&#39; (&#39;Create a new Radius profile template&#39;) interface, and radius profile ID can be obtained from &#39;Get Radius profile list&#39; (&#39;Get Radius profile template list&#39;) interface | [optional] 

## Methods

### NewDot1xEapOpenApiVO

`func NewDot1xEapOpenApiVO(enable bool, ) *Dot1xEapOpenApiVO`

NewDot1xEapOpenApiVO instantiates a new Dot1xEapOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xEapOpenApiVOWithDefaults

`func NewDot1xEapOpenApiVOWithDefaults() *Dot1xEapOpenApiVO`

NewDot1xEapOpenApiVOWithDefaults instantiates a new Dot1xEapOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEaps

`func (o *Dot1xEapOpenApiVO) GetEaps() []Dot1xEapSettingOpenApiVO`

GetEaps returns the Eaps field if non-nil, zero value otherwise.

### GetEapsOk

`func (o *Dot1xEapOpenApiVO) GetEapsOk() (*[]Dot1xEapSettingOpenApiVO, bool)`

GetEapsOk returns a tuple with the Eaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEaps

`func (o *Dot1xEapOpenApiVO) SetEaps(v []Dot1xEapSettingOpenApiVO)`

SetEaps sets Eaps field to given value.

### HasEaps

`func (o *Dot1xEapOpenApiVO) HasEaps() bool`

HasEaps returns a boolean if a field has been set.

### GetEnable

`func (o *Dot1xEapOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *Dot1xEapOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *Dot1xEapOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMacFormat

`func (o *Dot1xEapOpenApiVO) GetMacFormat() int32`

GetMacFormat returns the MacFormat field if non-nil, zero value otherwise.

### GetMacFormatOk

`func (o *Dot1xEapOpenApiVO) GetMacFormatOk() (*int32, bool)`

GetMacFormatOk returns a tuple with the MacFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacFormat

`func (o *Dot1xEapOpenApiVO) SetMacFormat(v int32)`

SetMacFormat sets MacFormat field to given value.

### HasMacFormat

`func (o *Dot1xEapOpenApiVO) HasMacFormat() bool`

HasMacFormat returns a boolean if a field has been set.

### GetRadiusProfileId

`func (o *Dot1xEapOpenApiVO) GetRadiusProfileId() string`

GetRadiusProfileId returns the RadiusProfileId field if non-nil, zero value otherwise.

### GetRadiusProfileIdOk

`func (o *Dot1xEapOpenApiVO) GetRadiusProfileIdOk() (*string, bool)`

GetRadiusProfileIdOk returns a tuple with the RadiusProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusProfileId

`func (o *Dot1xEapOpenApiVO) SetRadiusProfileId(v string)`

SetRadiusProfileId sets RadiusProfileId field to given value.

### HasRadiusProfileId

`func (o *Dot1xEapOpenApiVO) HasRadiusProfileId() bool`

HasRadiusProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


