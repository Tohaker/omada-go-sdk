# OswMlagPeerSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device mac | [optional] 
**MlagPeerPorts** | Pointer to **[]int32** | M-LAG group peer device LAG ports | [optional] 
**MlagPeerStandardPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | M-LAG group peer device LAG standard ports | [optional] 

## Methods

### NewOswMlagPeerSettingVO

`func NewOswMlagPeerSettingVO() *OswMlagPeerSettingVO`

NewOswMlagPeerSettingVO instantiates a new OswMlagPeerSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswMlagPeerSettingVOWithDefaults

`func NewOswMlagPeerSettingVOWithDefaults() *OswMlagPeerSettingVO`

NewOswMlagPeerSettingVOWithDefaults instantiates a new OswMlagPeerSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswMlagPeerSettingVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswMlagPeerSettingVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswMlagPeerSettingVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswMlagPeerSettingVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMlagPeerPorts

`func (o *OswMlagPeerSettingVO) GetMlagPeerPorts() []int32`

GetMlagPeerPorts returns the MlagPeerPorts field if non-nil, zero value otherwise.

### GetMlagPeerPortsOk

`func (o *OswMlagPeerSettingVO) GetMlagPeerPortsOk() (*[]int32, bool)`

GetMlagPeerPortsOk returns a tuple with the MlagPeerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerPorts

`func (o *OswMlagPeerSettingVO) SetMlagPeerPorts(v []int32)`

SetMlagPeerPorts sets MlagPeerPorts field to given value.

### HasMlagPeerPorts

`func (o *OswMlagPeerSettingVO) HasMlagPeerPorts() bool`

HasMlagPeerPorts returns a boolean if a field has been set.

### GetMlagPeerStandardPorts

`func (o *OswMlagPeerSettingVO) GetMlagPeerStandardPorts() []OswStandPortVO`

GetMlagPeerStandardPorts returns the MlagPeerStandardPorts field if non-nil, zero value otherwise.

### GetMlagPeerStandardPortsOk

`func (o *OswMlagPeerSettingVO) GetMlagPeerStandardPortsOk() (*[]OswStandPortVO, bool)`

GetMlagPeerStandardPortsOk returns a tuple with the MlagPeerStandardPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagPeerStandardPorts

`func (o *OswMlagPeerSettingVO) SetMlagPeerStandardPorts(v []OswStandPortVO)`

SetMlagPeerStandardPorts sets MlagPeerStandardPorts field to given value.

### HasMlagPeerStandardPorts

`func (o *OswMlagPeerSettingVO) HasMlagPeerStandardPorts() bool`

HasMlagPeerStandardPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


