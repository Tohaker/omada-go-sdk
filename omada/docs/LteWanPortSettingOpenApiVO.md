# LteWanPortSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandMode** | Pointer to **int32** | First SIM card&#39;s bandMode: 0:auto, 1:manual, only for networkMode 4g Only and 4G Preferred | [optional] 
**BandMode2** | Pointer to **int32** | Second SIM card&#39;s bandMode: 0:auto, 1:manual, only for networkMode 4g Only and 4G Preferred | [optional] 
**Bands** | Pointer to **[]string** | First SIM card&#39;s. Only for band mode manual | [optional] 
**Bands2** | Pointer to **[]string** | Second SIM card&#39;s bands. Only for band mode manual | [optional] 
**Bands5g** | Pointer to **[]string** | First SIM card&#39;s. Only for band mode manual | [optional] 
**Bands5g2** | Pointer to **[]string** | Second SIM card&#39;s. Only for band mode manual | [optional] 
**CardStatus** | Pointer to **int32** | cardStatus: 0:no_sim, 1:unknown, 2:pin_lock, 3:pin_verified, 4:ready, 5:puk_lock, 6:blocked | [optional] 
**DataRoaming** | **bool** | First SIM card&#39;s data roaming | 
**DataRoaming2** | Pointer to **bool** | Second SIM card&#39;s data roaming | [optional] 
**DialupSetting** | [**DialupSettingOpenApiVO**](DialupSettingOpenApiVO.md) |  | 
**DialupSetting2** | Pointer to [**DialupSettingOpenApiVO**](DialupSettingOpenApiVO.md) |  | [optional] 
**FailoverTimeout** | Pointer to **int32** | First SIM card&#39;s failoverTimeout. Set the dial-up timeout (100 to 3552 seconds). If the connection is not successfully established within the specified time, the gateway will use the other SIM card to connect to the internet. | [optional] 
**FailoverTimeout2** | Pointer to **int32** | Second SIM card&#39;s failoverTimeout. Set the dial-up timeout (100 to 3552 seconds). If the connection is not successfully established within the specified time, the gateway will use the other SIM card to connect to the internet. | [optional] 
**MobileData** | **bool** | mobile data | 
**NetworkMode** | **int32** | First SIM card&#39;s networkMode: 1:3G Only, 2:4G Only, 3:4G Preferred, 4: 5G-NSA/4G, 5: 5G-SA, 6: 5G/4G/3G | 
**NetworkMode2** | Pointer to **int32** | Second SIM card&#39;s networkMode: 1:3G Only, 2:4G Only, 3:4G Preferred4: 5G-NSA/4G, 5: 5G-SA, 6: 5G/4G/3G | [optional] 
**PortDescription** | Pointer to **string** | Port description | [optional] 
**PortName** | Pointer to **string** | Port name | [optional] 
**PortUuId** | **string** | Port UUID | 
**SimCardUsed** | Pointer to **int32** | The SIM card being used | [optional] 
**SimPriority** | Pointer to **int32** | Which SIM card is used first. | [optional] 
**Status** | Pointer to **int32** | 0: unknown, 1: disconnected, 2:disconnecting, 3:connected, 4:connecting | [optional] 
**SupportNetworkMode** | Pointer to **[]int32** | The network mode that the device can select. | [optional] 

## Methods

### NewLteWanPortSettingOpenApiVO

`func NewLteWanPortSettingOpenApiVO(dataRoaming bool, dialupSetting DialupSettingOpenApiVO, mobileData bool, networkMode int32, portUuId string, ) *LteWanPortSettingOpenApiVO`

NewLteWanPortSettingOpenApiVO instantiates a new LteWanPortSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLteWanPortSettingOpenApiVOWithDefaults

`func NewLteWanPortSettingOpenApiVOWithDefaults() *LteWanPortSettingOpenApiVO`

NewLteWanPortSettingOpenApiVOWithDefaults instantiates a new LteWanPortSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandMode

`func (o *LteWanPortSettingOpenApiVO) GetBandMode() int32`

GetBandMode returns the BandMode field if non-nil, zero value otherwise.

### GetBandModeOk

`func (o *LteWanPortSettingOpenApiVO) GetBandModeOk() (*int32, bool)`

GetBandModeOk returns a tuple with the BandMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandMode

`func (o *LteWanPortSettingOpenApiVO) SetBandMode(v int32)`

SetBandMode sets BandMode field to given value.

### HasBandMode

`func (o *LteWanPortSettingOpenApiVO) HasBandMode() bool`

HasBandMode returns a boolean if a field has been set.

### GetBandMode2

`func (o *LteWanPortSettingOpenApiVO) GetBandMode2() int32`

GetBandMode2 returns the BandMode2 field if non-nil, zero value otherwise.

### GetBandMode2Ok

`func (o *LteWanPortSettingOpenApiVO) GetBandMode2Ok() (*int32, bool)`

GetBandMode2Ok returns a tuple with the BandMode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandMode2

`func (o *LteWanPortSettingOpenApiVO) SetBandMode2(v int32)`

SetBandMode2 sets BandMode2 field to given value.

### HasBandMode2

`func (o *LteWanPortSettingOpenApiVO) HasBandMode2() bool`

HasBandMode2 returns a boolean if a field has been set.

### GetBands

`func (o *LteWanPortSettingOpenApiVO) GetBands() []string`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *LteWanPortSettingOpenApiVO) GetBandsOk() (*[]string, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *LteWanPortSettingOpenApiVO) SetBands(v []string)`

SetBands sets Bands field to given value.

### HasBands

`func (o *LteWanPortSettingOpenApiVO) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetBands2

`func (o *LteWanPortSettingOpenApiVO) GetBands2() []string`

GetBands2 returns the Bands2 field if non-nil, zero value otherwise.

### GetBands2Ok

`func (o *LteWanPortSettingOpenApiVO) GetBands2Ok() (*[]string, bool)`

GetBands2Ok returns a tuple with the Bands2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands2

`func (o *LteWanPortSettingOpenApiVO) SetBands2(v []string)`

SetBands2 sets Bands2 field to given value.

### HasBands2

`func (o *LteWanPortSettingOpenApiVO) HasBands2() bool`

HasBands2 returns a boolean if a field has been set.

### GetBands5g

`func (o *LteWanPortSettingOpenApiVO) GetBands5g() []string`

GetBands5g returns the Bands5g field if non-nil, zero value otherwise.

### GetBands5gOk

`func (o *LteWanPortSettingOpenApiVO) GetBands5gOk() (*[]string, bool)`

GetBands5gOk returns a tuple with the Bands5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands5g

`func (o *LteWanPortSettingOpenApiVO) SetBands5g(v []string)`

SetBands5g sets Bands5g field to given value.

### HasBands5g

`func (o *LteWanPortSettingOpenApiVO) HasBands5g() bool`

HasBands5g returns a boolean if a field has been set.

### GetBands5g2

`func (o *LteWanPortSettingOpenApiVO) GetBands5g2() []string`

GetBands5g2 returns the Bands5g2 field if non-nil, zero value otherwise.

### GetBands5g2Ok

`func (o *LteWanPortSettingOpenApiVO) GetBands5g2Ok() (*[]string, bool)`

GetBands5g2Ok returns a tuple with the Bands5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands5g2

`func (o *LteWanPortSettingOpenApiVO) SetBands5g2(v []string)`

SetBands5g2 sets Bands5g2 field to given value.

### HasBands5g2

`func (o *LteWanPortSettingOpenApiVO) HasBands5g2() bool`

HasBands5g2 returns a boolean if a field has been set.

### GetCardStatus

`func (o *LteWanPortSettingOpenApiVO) GetCardStatus() int32`

GetCardStatus returns the CardStatus field if non-nil, zero value otherwise.

### GetCardStatusOk

`func (o *LteWanPortSettingOpenApiVO) GetCardStatusOk() (*int32, bool)`

GetCardStatusOk returns a tuple with the CardStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardStatus

`func (o *LteWanPortSettingOpenApiVO) SetCardStatus(v int32)`

SetCardStatus sets CardStatus field to given value.

### HasCardStatus

`func (o *LteWanPortSettingOpenApiVO) HasCardStatus() bool`

HasCardStatus returns a boolean if a field has been set.

### GetDataRoaming

`func (o *LteWanPortSettingOpenApiVO) GetDataRoaming() bool`

GetDataRoaming returns the DataRoaming field if non-nil, zero value otherwise.

### GetDataRoamingOk

`func (o *LteWanPortSettingOpenApiVO) GetDataRoamingOk() (*bool, bool)`

GetDataRoamingOk returns a tuple with the DataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming

`func (o *LteWanPortSettingOpenApiVO) SetDataRoaming(v bool)`

SetDataRoaming sets DataRoaming field to given value.


### GetDataRoaming2

`func (o *LteWanPortSettingOpenApiVO) GetDataRoaming2() bool`

GetDataRoaming2 returns the DataRoaming2 field if non-nil, zero value otherwise.

### GetDataRoaming2Ok

`func (o *LteWanPortSettingOpenApiVO) GetDataRoaming2Ok() (*bool, bool)`

GetDataRoaming2Ok returns a tuple with the DataRoaming2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming2

`func (o *LteWanPortSettingOpenApiVO) SetDataRoaming2(v bool)`

SetDataRoaming2 sets DataRoaming2 field to given value.

### HasDataRoaming2

`func (o *LteWanPortSettingOpenApiVO) HasDataRoaming2() bool`

HasDataRoaming2 returns a boolean if a field has been set.

### GetDialupSetting

`func (o *LteWanPortSettingOpenApiVO) GetDialupSetting() DialupSettingOpenApiVO`

GetDialupSetting returns the DialupSetting field if non-nil, zero value otherwise.

### GetDialupSettingOk

`func (o *LteWanPortSettingOpenApiVO) GetDialupSettingOk() (*DialupSettingOpenApiVO, bool)`

GetDialupSettingOk returns a tuple with the DialupSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialupSetting

`func (o *LteWanPortSettingOpenApiVO) SetDialupSetting(v DialupSettingOpenApiVO)`

SetDialupSetting sets DialupSetting field to given value.


### GetDialupSetting2

`func (o *LteWanPortSettingOpenApiVO) GetDialupSetting2() DialupSettingOpenApiVO`

GetDialupSetting2 returns the DialupSetting2 field if non-nil, zero value otherwise.

### GetDialupSetting2Ok

`func (o *LteWanPortSettingOpenApiVO) GetDialupSetting2Ok() (*DialupSettingOpenApiVO, bool)`

GetDialupSetting2Ok returns a tuple with the DialupSetting2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialupSetting2

`func (o *LteWanPortSettingOpenApiVO) SetDialupSetting2(v DialupSettingOpenApiVO)`

SetDialupSetting2 sets DialupSetting2 field to given value.

### HasDialupSetting2

`func (o *LteWanPortSettingOpenApiVO) HasDialupSetting2() bool`

HasDialupSetting2 returns a boolean if a field has been set.

### GetFailoverTimeout

`func (o *LteWanPortSettingOpenApiVO) GetFailoverTimeout() int32`

GetFailoverTimeout returns the FailoverTimeout field if non-nil, zero value otherwise.

### GetFailoverTimeoutOk

`func (o *LteWanPortSettingOpenApiVO) GetFailoverTimeoutOk() (*int32, bool)`

GetFailoverTimeoutOk returns a tuple with the FailoverTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverTimeout

`func (o *LteWanPortSettingOpenApiVO) SetFailoverTimeout(v int32)`

SetFailoverTimeout sets FailoverTimeout field to given value.

### HasFailoverTimeout

`func (o *LteWanPortSettingOpenApiVO) HasFailoverTimeout() bool`

HasFailoverTimeout returns a boolean if a field has been set.

### GetFailoverTimeout2

`func (o *LteWanPortSettingOpenApiVO) GetFailoverTimeout2() int32`

GetFailoverTimeout2 returns the FailoverTimeout2 field if non-nil, zero value otherwise.

### GetFailoverTimeout2Ok

`func (o *LteWanPortSettingOpenApiVO) GetFailoverTimeout2Ok() (*int32, bool)`

GetFailoverTimeout2Ok returns a tuple with the FailoverTimeout2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverTimeout2

`func (o *LteWanPortSettingOpenApiVO) SetFailoverTimeout2(v int32)`

SetFailoverTimeout2 sets FailoverTimeout2 field to given value.

### HasFailoverTimeout2

`func (o *LteWanPortSettingOpenApiVO) HasFailoverTimeout2() bool`

HasFailoverTimeout2 returns a boolean if a field has been set.

### GetMobileData

`func (o *LteWanPortSettingOpenApiVO) GetMobileData() bool`

GetMobileData returns the MobileData field if non-nil, zero value otherwise.

### GetMobileDataOk

`func (o *LteWanPortSettingOpenApiVO) GetMobileDataOk() (*bool, bool)`

GetMobileDataOk returns a tuple with the MobileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileData

`func (o *LteWanPortSettingOpenApiVO) SetMobileData(v bool)`

SetMobileData sets MobileData field to given value.


### GetNetworkMode

`func (o *LteWanPortSettingOpenApiVO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *LteWanPortSettingOpenApiVO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *LteWanPortSettingOpenApiVO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.


### GetNetworkMode2

`func (o *LteWanPortSettingOpenApiVO) GetNetworkMode2() int32`

GetNetworkMode2 returns the NetworkMode2 field if non-nil, zero value otherwise.

### GetNetworkMode2Ok

`func (o *LteWanPortSettingOpenApiVO) GetNetworkMode2Ok() (*int32, bool)`

GetNetworkMode2Ok returns a tuple with the NetworkMode2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode2

`func (o *LteWanPortSettingOpenApiVO) SetNetworkMode2(v int32)`

SetNetworkMode2 sets NetworkMode2 field to given value.

### HasNetworkMode2

`func (o *LteWanPortSettingOpenApiVO) HasNetworkMode2() bool`

HasNetworkMode2 returns a boolean if a field has been set.

### GetPortDescription

`func (o *LteWanPortSettingOpenApiVO) GetPortDescription() string`

GetPortDescription returns the PortDescription field if non-nil, zero value otherwise.

### GetPortDescriptionOk

`func (o *LteWanPortSettingOpenApiVO) GetPortDescriptionOk() (*string, bool)`

GetPortDescriptionOk returns a tuple with the PortDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDescription

`func (o *LteWanPortSettingOpenApiVO) SetPortDescription(v string)`

SetPortDescription sets PortDescription field to given value.

### HasPortDescription

`func (o *LteWanPortSettingOpenApiVO) HasPortDescription() bool`

HasPortDescription returns a boolean if a field has been set.

### GetPortName

`func (o *LteWanPortSettingOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *LteWanPortSettingOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *LteWanPortSettingOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *LteWanPortSettingOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortUuId

`func (o *LteWanPortSettingOpenApiVO) GetPortUuId() string`

GetPortUuId returns the PortUuId field if non-nil, zero value otherwise.

### GetPortUuIdOk

`func (o *LteWanPortSettingOpenApiVO) GetPortUuIdOk() (*string, bool)`

GetPortUuIdOk returns a tuple with the PortUuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuId

`func (o *LteWanPortSettingOpenApiVO) SetPortUuId(v string)`

SetPortUuId sets PortUuId field to given value.


### GetSimCardUsed

`func (o *LteWanPortSettingOpenApiVO) GetSimCardUsed() int32`

GetSimCardUsed returns the SimCardUsed field if non-nil, zero value otherwise.

### GetSimCardUsedOk

`func (o *LteWanPortSettingOpenApiVO) GetSimCardUsedOk() (*int32, bool)`

GetSimCardUsedOk returns a tuple with the SimCardUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardUsed

`func (o *LteWanPortSettingOpenApiVO) SetSimCardUsed(v int32)`

SetSimCardUsed sets SimCardUsed field to given value.

### HasSimCardUsed

`func (o *LteWanPortSettingOpenApiVO) HasSimCardUsed() bool`

HasSimCardUsed returns a boolean if a field has been set.

### GetSimPriority

`func (o *LteWanPortSettingOpenApiVO) GetSimPriority() int32`

GetSimPriority returns the SimPriority field if non-nil, zero value otherwise.

### GetSimPriorityOk

`func (o *LteWanPortSettingOpenApiVO) GetSimPriorityOk() (*int32, bool)`

GetSimPriorityOk returns a tuple with the SimPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimPriority

`func (o *LteWanPortSettingOpenApiVO) SetSimPriority(v int32)`

SetSimPriority sets SimPriority field to given value.

### HasSimPriority

`func (o *LteWanPortSettingOpenApiVO) HasSimPriority() bool`

HasSimPriority returns a boolean if a field has been set.

### GetStatus

`func (o *LteWanPortSettingOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LteWanPortSettingOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LteWanPortSettingOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LteWanPortSettingOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportNetworkMode

`func (o *LteWanPortSettingOpenApiVO) GetSupportNetworkMode() []int32`

GetSupportNetworkMode returns the SupportNetworkMode field if non-nil, zero value otherwise.

### GetSupportNetworkModeOk

`func (o *LteWanPortSettingOpenApiVO) GetSupportNetworkModeOk() (*[]int32, bool)`

GetSupportNetworkModeOk returns a tuple with the SupportNetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportNetworkMode

`func (o *LteWanPortSettingOpenApiVO) SetSupportNetworkMode(v []int32)`

SetSupportNetworkMode sets SupportNetworkMode field to given value.

### HasSupportNetworkMode

`func (o *LteWanPortSettingOpenApiVO) HasSupportNetworkMode() bool`

HasSupportNetworkMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


