# DialupSettingOpenApiV2VO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | APN profile ID, only for apnMode manual | [optional] 
**ApnMode** | **int32** | Apn mode should be a value as follows: 0: auto; 1: manual. | 
**BandMode** | Pointer to **int32** | SIM card&#39;s bandMode: 0:auto, 1:manual, only for networkMode 4g Only and 4G Preferred | [optional] 
**Bands** | Pointer to **[]string** | SIM card&#39;s 5g bands. Only for band mode manual | [optional] 
**DataRoaming** | **bool** | SIM card&#39;s data roaming | 
**FailoverTimeout** | Pointer to **int32** | SIM card&#39;s failoverTimeout. Set the dial-up timeout (100 to 3552 seconds). If the connection is not successfully established within the specified time, the gateway will use the other SIM card to connect to the internet. | [optional] 
**Isp** | Pointer to **string** | Only for networkSearch mode manual. | [optional] 
**IspNum** | Pointer to **int32** | Only for networkSearch mode manual. | [optional] 
**NetworkMode** | **int32** | SIM card&#39;s networkMode: 1:3G Only, 2:4G Only, 3:4G Preferred4: 5G-NSA/4G, 5: 5G-SA, 6: 5G/4G/3G | 
**NetworkSearch** | **int32** | Network search mode should be a value as follows: 0: auto; 1: manual. | 

## Methods

### NewDialupSettingOpenApiV2VO

`func NewDialupSettingOpenApiV2VO(apnMode int32, dataRoaming bool, networkMode int32, networkSearch int32, ) *DialupSettingOpenApiV2VO`

NewDialupSettingOpenApiV2VO instantiates a new DialupSettingOpenApiV2VO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDialupSettingOpenApiV2VOWithDefaults

`func NewDialupSettingOpenApiV2VOWithDefaults() *DialupSettingOpenApiV2VO`

NewDialupSettingOpenApiV2VOWithDefaults instantiates a new DialupSettingOpenApiV2VO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *DialupSettingOpenApiV2VO) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *DialupSettingOpenApiV2VO) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *DialupSettingOpenApiV2VO) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *DialupSettingOpenApiV2VO) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetApnMode

`func (o *DialupSettingOpenApiV2VO) GetApnMode() int32`

GetApnMode returns the ApnMode field if non-nil, zero value otherwise.

### GetApnModeOk

`func (o *DialupSettingOpenApiV2VO) GetApnModeOk() (*int32, bool)`

GetApnModeOk returns a tuple with the ApnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnMode

`func (o *DialupSettingOpenApiV2VO) SetApnMode(v int32)`

SetApnMode sets ApnMode field to given value.


### GetBandMode

`func (o *DialupSettingOpenApiV2VO) GetBandMode() int32`

GetBandMode returns the BandMode field if non-nil, zero value otherwise.

### GetBandModeOk

`func (o *DialupSettingOpenApiV2VO) GetBandModeOk() (*int32, bool)`

GetBandModeOk returns a tuple with the BandMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandMode

`func (o *DialupSettingOpenApiV2VO) SetBandMode(v int32)`

SetBandMode sets BandMode field to given value.

### HasBandMode

`func (o *DialupSettingOpenApiV2VO) HasBandMode() bool`

HasBandMode returns a boolean if a field has been set.

### GetBands

`func (o *DialupSettingOpenApiV2VO) GetBands() []string`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *DialupSettingOpenApiV2VO) GetBandsOk() (*[]string, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *DialupSettingOpenApiV2VO) SetBands(v []string)`

SetBands sets Bands field to given value.

### HasBands

`func (o *DialupSettingOpenApiV2VO) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetDataRoaming

`func (o *DialupSettingOpenApiV2VO) GetDataRoaming() bool`

GetDataRoaming returns the DataRoaming field if non-nil, zero value otherwise.

### GetDataRoamingOk

`func (o *DialupSettingOpenApiV2VO) GetDataRoamingOk() (*bool, bool)`

GetDataRoamingOk returns a tuple with the DataRoaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataRoaming

`func (o *DialupSettingOpenApiV2VO) SetDataRoaming(v bool)`

SetDataRoaming sets DataRoaming field to given value.


### GetFailoverTimeout

`func (o *DialupSettingOpenApiV2VO) GetFailoverTimeout() int32`

GetFailoverTimeout returns the FailoverTimeout field if non-nil, zero value otherwise.

### GetFailoverTimeoutOk

`func (o *DialupSettingOpenApiV2VO) GetFailoverTimeoutOk() (*int32, bool)`

GetFailoverTimeoutOk returns a tuple with the FailoverTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverTimeout

`func (o *DialupSettingOpenApiV2VO) SetFailoverTimeout(v int32)`

SetFailoverTimeout sets FailoverTimeout field to given value.

### HasFailoverTimeout

`func (o *DialupSettingOpenApiV2VO) HasFailoverTimeout() bool`

HasFailoverTimeout returns a boolean if a field has been set.

### GetIsp

`func (o *DialupSettingOpenApiV2VO) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *DialupSettingOpenApiV2VO) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *DialupSettingOpenApiV2VO) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *DialupSettingOpenApiV2VO) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIspNum

`func (o *DialupSettingOpenApiV2VO) GetIspNum() int32`

GetIspNum returns the IspNum field if non-nil, zero value otherwise.

### GetIspNumOk

`func (o *DialupSettingOpenApiV2VO) GetIspNumOk() (*int32, bool)`

GetIspNumOk returns a tuple with the IspNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspNum

`func (o *DialupSettingOpenApiV2VO) SetIspNum(v int32)`

SetIspNum sets IspNum field to given value.

### HasIspNum

`func (o *DialupSettingOpenApiV2VO) HasIspNum() bool`

HasIspNum returns a boolean if a field has been set.

### GetNetworkMode

`func (o *DialupSettingOpenApiV2VO) GetNetworkMode() int32`

GetNetworkMode returns the NetworkMode field if non-nil, zero value otherwise.

### GetNetworkModeOk

`func (o *DialupSettingOpenApiV2VO) GetNetworkModeOk() (*int32, bool)`

GetNetworkModeOk returns a tuple with the NetworkMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMode

`func (o *DialupSettingOpenApiV2VO) SetNetworkMode(v int32)`

SetNetworkMode sets NetworkMode field to given value.


### GetNetworkSearch

`func (o *DialupSettingOpenApiV2VO) GetNetworkSearch() int32`

GetNetworkSearch returns the NetworkSearch field if non-nil, zero value otherwise.

### GetNetworkSearchOk

`func (o *DialupSettingOpenApiV2VO) GetNetworkSearchOk() (*int32, bool)`

GetNetworkSearchOk returns a tuple with the NetworkSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSearch

`func (o *DialupSettingOpenApiV2VO) SetNetworkSearch(v int32)`

SetNetworkSearch sets NetworkSearch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


