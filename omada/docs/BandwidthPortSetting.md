# BandwidthPortSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownstreamBandwidth** | Pointer to **int32** | Downstream bandwidth should be within the range of 1–9999999. | [optional] 
**DownstreamBandwidthUnit** | Pointer to **int32** | Downstream bandwidth unit should be a value as follows: 1: Kbps, 2: Mbps. | [optional] 
**PortUuid** | **string** | Port uuid. | 
**UpstreamBandwidth** | Pointer to **int32** | Upstream bandwidth should be within the range of 1–9999999. | [optional] 
**UpstreamBandwidthUnit** | Pointer to **int32** | Upstream bandwidth unit should be a value as follows: 1: Kbps, 2: Mbps. | [optional] 

## Methods

### NewBandwidthPortSetting

`func NewBandwidthPortSetting(portUuid string, ) *BandwidthPortSetting`

NewBandwidthPortSetting instantiates a new BandwidthPortSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthPortSettingWithDefaults

`func NewBandwidthPortSettingWithDefaults() *BandwidthPortSetting`

NewBandwidthPortSettingWithDefaults instantiates a new BandwidthPortSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownstreamBandwidth

`func (o *BandwidthPortSetting) GetDownstreamBandwidth() int32`

GetDownstreamBandwidth returns the DownstreamBandwidth field if non-nil, zero value otherwise.

### GetDownstreamBandwidthOk

`func (o *BandwidthPortSetting) GetDownstreamBandwidthOk() (*int32, bool)`

GetDownstreamBandwidthOk returns a tuple with the DownstreamBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamBandwidth

`func (o *BandwidthPortSetting) SetDownstreamBandwidth(v int32)`

SetDownstreamBandwidth sets DownstreamBandwidth field to given value.

### HasDownstreamBandwidth

`func (o *BandwidthPortSetting) HasDownstreamBandwidth() bool`

HasDownstreamBandwidth returns a boolean if a field has been set.

### GetDownstreamBandwidthUnit

`func (o *BandwidthPortSetting) GetDownstreamBandwidthUnit() int32`

GetDownstreamBandwidthUnit returns the DownstreamBandwidthUnit field if non-nil, zero value otherwise.

### GetDownstreamBandwidthUnitOk

`func (o *BandwidthPortSetting) GetDownstreamBandwidthUnitOk() (*int32, bool)`

GetDownstreamBandwidthUnitOk returns a tuple with the DownstreamBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownstreamBandwidthUnit

`func (o *BandwidthPortSetting) SetDownstreamBandwidthUnit(v int32)`

SetDownstreamBandwidthUnit sets DownstreamBandwidthUnit field to given value.

### HasDownstreamBandwidthUnit

`func (o *BandwidthPortSetting) HasDownstreamBandwidthUnit() bool`

HasDownstreamBandwidthUnit returns a boolean if a field has been set.

### GetPortUuid

`func (o *BandwidthPortSetting) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *BandwidthPortSetting) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *BandwidthPortSetting) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.


### GetUpstreamBandwidth

`func (o *BandwidthPortSetting) GetUpstreamBandwidth() int32`

GetUpstreamBandwidth returns the UpstreamBandwidth field if non-nil, zero value otherwise.

### GetUpstreamBandwidthOk

`func (o *BandwidthPortSetting) GetUpstreamBandwidthOk() (*int32, bool)`

GetUpstreamBandwidthOk returns a tuple with the UpstreamBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamBandwidth

`func (o *BandwidthPortSetting) SetUpstreamBandwidth(v int32)`

SetUpstreamBandwidth sets UpstreamBandwidth field to given value.

### HasUpstreamBandwidth

`func (o *BandwidthPortSetting) HasUpstreamBandwidth() bool`

HasUpstreamBandwidth returns a boolean if a field has been set.

### GetUpstreamBandwidthUnit

`func (o *BandwidthPortSetting) GetUpstreamBandwidthUnit() int32`

GetUpstreamBandwidthUnit returns the UpstreamBandwidthUnit field if non-nil, zero value otherwise.

### GetUpstreamBandwidthUnitOk

`func (o *BandwidthPortSetting) GetUpstreamBandwidthUnitOk() (*int32, bool)`

GetUpstreamBandwidthUnitOk returns a tuple with the UpstreamBandwidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpstreamBandwidthUnit

`func (o *BandwidthPortSetting) SetUpstreamBandwidthUnit(v int32)`

SetUpstreamBandwidthUnit sets UpstreamBandwidthUnit field to given value.

### HasUpstreamBandwidthUnit

`func (o *BandwidthPortSetting) HasUpstreamBandwidthUnit() bool`

HasUpstreamBandwidthUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


