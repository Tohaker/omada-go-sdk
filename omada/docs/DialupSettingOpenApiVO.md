# DialupSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apn** | Pointer to **string** | APN profile ID, only for apnMode manual | [optional] 
**ApnMode** | **int32** | Apn mode should be a value as follows: 0: auto; 1: manual. | 
**Isp** | Pointer to **string** | Only for networkSearch mode manual and the ISP from the scan result must be available. | [optional] 
**IspNum** | Pointer to **int32** | Only for networkSearch mode manual and the ISP from the scan result must be available. | [optional] 
**NetworkSearch** | **int32** | Network search mode should be a value as follows: 0: auto; 1: manual. | 

## Methods

### NewDialupSettingOpenApiVO

`func NewDialupSettingOpenApiVO(apnMode int32, networkSearch int32, ) *DialupSettingOpenApiVO`

NewDialupSettingOpenApiVO instantiates a new DialupSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDialupSettingOpenApiVOWithDefaults

`func NewDialupSettingOpenApiVOWithDefaults() *DialupSettingOpenApiVO`

NewDialupSettingOpenApiVOWithDefaults instantiates a new DialupSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApn

`func (o *DialupSettingOpenApiVO) GetApn() string`

GetApn returns the Apn field if non-nil, zero value otherwise.

### GetApnOk

`func (o *DialupSettingOpenApiVO) GetApnOk() (*string, bool)`

GetApnOk returns a tuple with the Apn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApn

`func (o *DialupSettingOpenApiVO) SetApn(v string)`

SetApn sets Apn field to given value.

### HasApn

`func (o *DialupSettingOpenApiVO) HasApn() bool`

HasApn returns a boolean if a field has been set.

### GetApnMode

`func (o *DialupSettingOpenApiVO) GetApnMode() int32`

GetApnMode returns the ApnMode field if non-nil, zero value otherwise.

### GetApnModeOk

`func (o *DialupSettingOpenApiVO) GetApnModeOk() (*int32, bool)`

GetApnModeOk returns a tuple with the ApnMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApnMode

`func (o *DialupSettingOpenApiVO) SetApnMode(v int32)`

SetApnMode sets ApnMode field to given value.


### GetIsp

`func (o *DialupSettingOpenApiVO) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *DialupSettingOpenApiVO) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *DialupSettingOpenApiVO) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *DialupSettingOpenApiVO) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetIspNum

`func (o *DialupSettingOpenApiVO) GetIspNum() int32`

GetIspNum returns the IspNum field if non-nil, zero value otherwise.

### GetIspNumOk

`func (o *DialupSettingOpenApiVO) GetIspNumOk() (*int32, bool)`

GetIspNumOk returns a tuple with the IspNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspNum

`func (o *DialupSettingOpenApiVO) SetIspNum(v int32)`

SetIspNum sets IspNum field to given value.

### HasIspNum

`func (o *DialupSettingOpenApiVO) HasIspNum() bool`

HasIspNum returns a boolean if a field has been set.

### GetNetworkSearch

`func (o *DialupSettingOpenApiVO) GetNetworkSearch() int32`

GetNetworkSearch returns the NetworkSearch field if non-nil, zero value otherwise.

### GetNetworkSearchOk

`func (o *DialupSettingOpenApiVO) GetNetworkSearchOk() (*int32, bool)`

GetNetworkSearchOk returns a tuple with the NetworkSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSearch

`func (o *DialupSettingOpenApiVO) SetNetworkSearch(v int32)`

SetNetworkSearch sets NetworkSearch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


