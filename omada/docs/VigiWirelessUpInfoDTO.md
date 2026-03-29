# VigiWirelessUpInfoDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel** | Pointer to **int32** | Wireless Vigi Channel | [optional] 
**MultiLink** | Pointer to [**[]MultiLinkEntryDTO**](MultiLinkEntryDTO.md) | Vigi MultiLink In MLO Mode | [optional] 
**Radio** | Pointer to **int32** | Wireless Vigi Radio | [optional] 
**Ssid** | Pointer to **string** | Wireless Vigi Ssid | [optional] 
**Support5g2** | Pointer to **bool** | Whether The Vigi Supports 5g2 Or Not | [optional] 

## Methods

### NewVigiWirelessUpInfoDTO

`func NewVigiWirelessUpInfoDTO() *VigiWirelessUpInfoDTO`

NewVigiWirelessUpInfoDTO instantiates a new VigiWirelessUpInfoDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVigiWirelessUpInfoDTOWithDefaults

`func NewVigiWirelessUpInfoDTOWithDefaults() *VigiWirelessUpInfoDTO`

NewVigiWirelessUpInfoDTOWithDefaults instantiates a new VigiWirelessUpInfoDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel

`func (o *VigiWirelessUpInfoDTO) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VigiWirelessUpInfoDTO) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VigiWirelessUpInfoDTO) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VigiWirelessUpInfoDTO) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetMultiLink

`func (o *VigiWirelessUpInfoDTO) GetMultiLink() []MultiLinkEntryDTO`

GetMultiLink returns the MultiLink field if non-nil, zero value otherwise.

### GetMultiLinkOk

`func (o *VigiWirelessUpInfoDTO) GetMultiLinkOk() (*[]MultiLinkEntryDTO, bool)`

GetMultiLinkOk returns a tuple with the MultiLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiLink

`func (o *VigiWirelessUpInfoDTO) SetMultiLink(v []MultiLinkEntryDTO)`

SetMultiLink sets MultiLink field to given value.

### HasMultiLink

`func (o *VigiWirelessUpInfoDTO) HasMultiLink() bool`

HasMultiLink returns a boolean if a field has been set.

### GetRadio

`func (o *VigiWirelessUpInfoDTO) GetRadio() int32`

GetRadio returns the Radio field if non-nil, zero value otherwise.

### GetRadioOk

`func (o *VigiWirelessUpInfoDTO) GetRadioOk() (*int32, bool)`

GetRadioOk returns a tuple with the Radio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio

`func (o *VigiWirelessUpInfoDTO) SetRadio(v int32)`

SetRadio sets Radio field to given value.

### HasRadio

`func (o *VigiWirelessUpInfoDTO) HasRadio() bool`

HasRadio returns a boolean if a field has been set.

### GetSsid

`func (o *VigiWirelessUpInfoDTO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *VigiWirelessUpInfoDTO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *VigiWirelessUpInfoDTO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *VigiWirelessUpInfoDTO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetSupport5g2

`func (o *VigiWirelessUpInfoDTO) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *VigiWirelessUpInfoDTO) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *VigiWirelessUpInfoDTO) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *VigiWirelessUpInfoDTO) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


