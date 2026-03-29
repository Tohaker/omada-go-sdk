# LteWanPortSettingConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandMode** | Pointer to **int32** | First SIM card&#39;s bandMode: 0:auto, 1:manual, only for networkMode 4g Only and 4G Preferred | [optional] 
**BandMode2** | Pointer to **int32** | Second SIM card&#39;s bandMode: 0:auto, 1:manual, only for networkMode 4g Only and 4G Preferred | [optional] 
**Bands** | Pointer to **[]string** | First SIM card&#39;s bands. Only for band mode manual | [optional] 
**Bands2** | Pointer to **[]string** | Second SIM card&#39;s bands. Only for band mode manual | [optional] 
**Bands5g** | Pointer to **[]string** | First SIM card&#39;s. Only for band mode manual | [optional] 
**Bands5g2** | Pointer to **[]string** | Second SIM card&#39;s. Only for band mode manual | [optional] 
**DataRoaming** | **bool** | First SIM card&#39;s data roaming | 
**DataRoaming2** | Pointer to **bool** | Second SIM card&#39;s data roaming | [optional] 
**DialupSetting** | [**DialupSettingOpenApiVO**](DialupSettingOpenApiVO.md) |  | 
**DialupSetting2** | Pointer to [**DialupSettingOpenApiVO**](DialupSettingOpenApiVO.md) |  | [optional] 
**FailoverTimeout** | Pointer to **int32** | First SIM card&#39;s failoverTimeout. Set the dial-up timeout (100 to 3552 seconds). If the connection is not successfully established within the specified time, the gateway will use the other SIM card to connect to the internet. | [optional] 
**FailoverTimeout2** | Pointer to **int32** | Second SIM card&#39;s failoverTimeout. Set the dial-up timeout (100 to 3552 seconds). If the connection is not successfully established within the specified time, the gateway will use the other SIM card to connect to the internet. | [optional] 
**MobileData** | **bool** | mobile data | 
**NetworkMode** | **int32** | First SIM card&#39;s networkMode: 1:3G Only, 2:4G Only, 3:4G Preferred4: 5G-NSA/4G, 5: 5G-SA, 6: 5G/4G/3G | 
**NetworkMode2** | Pointer to **int32** | Second SIM card&#39;s networkMode: 1:3G Only, 2:4G Only, 3:4G Preferred4: 5G-NSA/4G, 5: 5G-SA, 6: 5G/4G/3G | [optional] 
**PortDescription** | Pointer to **string** | Port description should contain 1 to 32 characters. | [optional] 
**PortId** | **string** | Port ID | 
**SimPriority** | Pointer to **int32** | Set which SIM card is used first. SIM Priority takes effect only when the device is powered on and the priority is changed. If only one SIM card is inserted, this card is used by default. | [optional] 

## Methods

### NewLteWanPortSettingConfigOpenApiVO

`func NewLteWanPortSettingConfigOpenApiVO(dataRoaming bool, dialupSetting DialupSettingOpenApiVO, mobileData bool, networkMode int32, portId string, ) *LteWanPortSettingConfigOpenApiVO`

NewLteWanPortSettingConfigOpenApiVO instantiates a new LteWanPortSettingConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLteWanPortSettingConfigOpenApiVOWithDefaults

`func NewLteWanPortSettingConfigOpenApiVOWithDefaults() *LteWanPortSettingConfigOpenApiVO`

NewLteWanPortSettingConfigOpenApiVOWithDefaults instantiates a new LteWanPortSettingConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandMode

`func (o *LteWanPortSettingConfigOpenApiVO) GetBandMode() int32`

GetBandMode returns the BandMode field if non-nil, zero value otherwise.

### GetBandModeOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetBandModeOk() (*int32, bool)`

GetBandModeOk returns a tuple with the BandMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandMode

`func (o *LteWanPortSettingConfigOpenApiVO) SetBandMode(v int32)`

SetBandMode sets BandMode field to given value.

### HasBandMode

`func (o *LteWanPortSettingConfigOpenApiVO) HasBandMode() bool`

HasBandMode returns a boolean if a field has been set.

### GetBandMode2

`func (o *LteWanPortSettingConfigOpenApiVO) GetBandMode2() int32`

GetBandMode2 returns the BandMode2 field if non-nil, zero value otherwise.

### GetBandMode2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetBandMode2Ok() (*int32, bool)`

GetBandMode2Ok returns a tuple with the BandMode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandMode2

`func (o *LteWanPortSettingConfigOpenApiVO) SetBandMode2(v int32)`

SetBandMode2 sets BandMode2 field to given value.

### HasBandMode2

`func (o *LteWanPortSettingConfigOpenApiVO) HasBandMode2() bool`

HasBandMode2 returns a boolean if a field has been set.

### GetBands

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands() []string`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetBandsOk() (*[]string, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *LteWanPortSettingConfigOpenApiVO) SetBands(v []string)`

SetBands sets Bands field to given value.

### HasBands

`func (o *LteWanPortSettingConfigOpenApiVO) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetBands2

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands2() []string`

GetBands2 returns the Bands2 field if non-nil, zero value otherwise.

### GetBands2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands2Ok() (*[]string, bool)`

GetBands2Ok returns a tuple with the Bands2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands2

`func (o *LteWanPortSettingConfigOpenApiVO) SetBands2(v []string)`

SetBands2 sets Bands2 field to given value.

### HasBands2

`func (o *LteWanPortSettingConfigOpenApiVO) HasBands2() bool`

HasBands2 returns a boolean if a field has been set.

### GetBands5g

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands5g() []string`

GetBands5g returns the Bands5g field if non-nil, zero value otherwise.

### GetBands5gOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands5gOk() (*[]string, bool)`

GetBands5gOk returns a tuple with the Bands5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands5g

`func (o *LteWanPortSettingConfigOpenApiVO) SetBands5g(v []string)`

SetBands5g sets Bands5g field to given value.

### HasBands5g

`func (o *LteWanPortSettingConfigOpenApiVO) HasBands5g() bool`

HasBands5g returns a boolean if a field has been set.

### GetBands5g2

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands5g2() []string`

GetBands5g2 returns the Bands5g2 field if non-nil, zero value otherwise.

### GetBands5g2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetBands5g2Ok() (*[]string, bool)`

GetBands5g2Ok returns a tuple with the Bands5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands5g2

`func (o *LteWanPortSettingConfigOpenApiVO) SetBands5g2(v []string)`

SetBands5g2 sets Bands5g2 field to given value.

### HasBands5g2

`func (o *LteWanPortSettingConfigOpenApiVO) HasBands5g2() bool`

HasBands5g2 returns a boolean if a field has been set.

### GetDataRoaming

`func (o *LteWanPortSettingConfigOpenApiVO) GetDataRoaming() bool`

GetDataRoaming returns the DataRoaming field if non-nil, zero value otherwise.

### GetDataRoamingOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetDataRoamingOk() (*bool, bool)`

GetDataRoamingOk returns a tuple with the DataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming

`func (o *LteWanPortSettingConfigOpenApiVO) SetDataRoaming(v bool)`

SetDataRoaming sets DataRoaming field to given value.


### GetDataRoaming2

`func (o *LteWanPortSettingConfigOpenApiVO) GetDataRoaming2() bool`

GetDataRoaming2 returns the DataRoaming2 field if non-nil, zero value otherwise.

### GetDataRoaming2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetDataRoaming2Ok() (*bool, bool)`

GetDataRoaming2Ok returns a tuple with the DataRoaming2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming2

`func (o *LteWanPortSettingConfigOpenApiVO) SetDataRoaming2(v bool)`

SetDataRoaming2 sets DataRoaming2 field to given value.

### HasDataRoaming2

`func (o *LteWanPortSettingConfigOpenApiVO) HasDataRoaming2() bool`

HasDataRoaming2 returns a boolean if a field has been set.

### GetDialupSetting

`func (o *LteWanPortSettingConfigOpenApiVO) GetDialupSetting() DialupSettingOpenApiVO`

GetDialupSetting returns the DialupSetting field if non-nil, zero value otherwise.

### GetDialupSettingOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetDialupSettingOk() (*DialupSettingOpenApiVO, bool)`

GetDialupSettingOk returns a tuple with the DialupSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialupSetting

`func (o *LteWanPortSettingConfigOpenApiVO) SetDialupSetting(v DialupSettingOpenApiVO)`

SetDialupSetting sets DialupSetting field to given value.


### GetDialupSetting2

`func (o *LteWanPortSettingConfigOpenApiVO) GetDialupSetting2() DialupSettingOpenApiVO`

GetDialupSetting2 returns the DialupSetting2 field if non-nil, zero value otherwise.

### GetDialupSetting2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetDialupSetting2Ok() (*DialupSettingOpenApiVO, bool)`

GetDialupSetting2Ok returns a tuple with the DialupSetting2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialupSetting2

`func (o *LteWanPortSettingConfigOpenApiVO) SetDialupSetting2(v DialupSettingOpenApiVO)`

SetDialupSetting2 sets DialupSetting2 field to given value.

### HasDialupSetting2

`func (o *LteWanPortSettingConfigOpenApiVO) HasDialupSetting2() bool`

HasDialupSetting2 returns a boolean if a field has been set.

### GetFailoverTimeout

`func (o *LteWanPortSettingConfigOpenApiVO) GetFailoverTimeout() int32`

GetFailoverTimeout returns the FailoverTimeout field if non-nil, zero value otherwise.

### GetFailoverTimeoutOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetFailoverTimeoutOk() (*int32, bool)`

GetFailoverTimeoutOk returns a tuple with the FailoverTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverTimeout

`func (o *LteWanPortSettingConfigOpenApiVO) SetFailoverTimeout(v int32)`

SetFailoverTimeout sets FailoverTimeout field to given value.

### HasFailoverTimeout

`func (o *LteWanPortSettingConfigOpenApiVO) HasFailoverTimeout() bool`

HasFailoverTimeout returns a boolean if a field has been set.

### GetFailoverTimeout2

`func (o *LteWanPortSettingConfigOpenApiVO) GetFailoverTimeout2() int32`

GetFailoverTimeout2 returns the FailoverTimeout2 field if non-nil, zero value otherwise.

### GetFailoverTimeout2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetFailoverTimeout2Ok() (*int32, bool)`

GetFailoverTimeout2Ok returns a tuple with the FailoverTimeout2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverTimeout2

`func (o *LteWanPortSettingConfigOpenApiVO) SetFailoverTimeout2(v int32)`

SetFailoverTimeout2 sets FailoverTimeout2 field to given value.

### HasFailoverTimeout2

`func (o *LteWanPortSettingConfigOpenApiVO) HasFailoverTimeout2() bool`

HasFailoverTimeout2 returns a boolean if a field has been set.

### GetMobileData

`func (o *LteWanPortSettingConfigOpenApiVO) GetMobileData() bool`

GetMobileData returns the MobileData field if non-nil, zero value otherwise.

### GetMobileDataOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetMobileDataOk() (*bool, bool)`

GetMobileDataOk returns a tuple with the MobileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileData

`func (o *LteWanPortSettingConfigOpenApiVO) SetMobileData(v bool)`

SetMobileData sets MobileData field to given value.


### GetNetworkMode

`func (o *LteWanPortSettingConfigOpenApiVO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *LteWanPortSettingConfigOpenApiVO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.


### GetNetworkMode2

`func (o *LteWanPortSettingConfigOpenApiVO) GetNetworkMode2() int32`

GetNetworkMode2 returns the NetworkMode2 field if non-nil, zero value otherwise.

### GetNetworkMode2Ok

`func (o *LteWanPortSettingConfigOpenApiVO) GetNetworkMode2Ok() (*int32, bool)`

GetNetworkMode2Ok returns a tuple with the NetworkMode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode2

`func (o *LteWanPortSettingConfigOpenApiVO) SetNetworkMode2(v int32)`

SetNetworkMode2 sets NetworkMode2 field to given value.

### HasNetworkMode2

`func (o *LteWanPortSettingConfigOpenApiVO) HasNetworkMode2() bool`

HasNetworkMode2 returns a boolean if a field has been set.

### GetPortDescription

`func (o *LteWanPortSettingConfigOpenApiVO) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *LteWanPortSettingConfigOpenApiVO) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *LteWanPortSettingConfigOpenApiVO) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetPortId

`func (o *LteWanPortSettingConfigOpenApiVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *LteWanPortSettingConfigOpenApiVO) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetSimPriority

`func (o *LteWanPortSettingConfigOpenApiVO) GetSimPriority() int32`

GetSimPriority returns the SimPriority field if non-nil, zero value otherwise.

### GetSimPriorityOk

`func (o *LteWanPortSettingConfigOpenApiVO) GetSimPriorityOk() (*int32, bool)`

GetSimPriorityOk returns a tuple with the SimPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimPriority

`func (o *LteWanPortSettingConfigOpenApiVO) SetSimPriority(v int32)`

SetSimPriority sets SimPriority field to given value.

### HasSimPriority

`func (o *LteWanPortSettingConfigOpenApiVO) HasSimPriority() bool`

HasSimPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


