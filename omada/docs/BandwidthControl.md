# BandwidthControl

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BandwidthPortSettings** | Pointer to [**[]BandwidthPortSetting**](BandwidthPortSetting.md) | Bandwidth port settings of the bandwidth control. | [optional] 
**Enable** | **bool** | Enable the bandwidth control. | 
**ThresholdControlEnable** | Pointer to **bool** | Enable the threshold control of the bandwidth control. | [optional] 
**ThresholdValue** | Pointer to **int32** | Threshold value should be within the range of 1–100. Threshold value must be entered when threshold control is enable. | [optional] 

## Methods

### NewBandwidthControl

`func NewBandwidthControl(enable bool, ) *BandwidthControl`

NewBandwidthControl instantiates a new BandwidthControl object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandwidthControlWithDefaults

`func NewBandwidthControlWithDefaults() *BandwidthControl`

NewBandwidthControlWithDefaults instantiates a new BandwidthControl object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBandwidthPortSettings

`func (o *BandwidthControl) GetBandwidthPortSettings() []BandwidthPortSetting`

GetBandwidthPortSettings returns the BandwidthPortSettings field if non-nil, zero value otherwise.

### GetBandwidthPortSettingsOk

`func (o *BandwidthControl) GetBandwidthPortSettingsOk() (*[]BandwidthPortSetting, bool)`

GetBandwidthPortSettingsOk returns a tuple with the BandwidthPortSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthPortSettings

`func (o *BandwidthControl) SetBandwidthPortSettings(v []BandwidthPortSetting)`

SetBandwidthPortSettings sets BandwidthPortSettings field to given value.

### HasBandwidthPortSettings

`func (o *BandwidthControl) HasBandwidthPortSettings() bool`

HasBandwidthPortSettings returns a boolean if a field has been set.

### GetEnable

`func (o *BandwidthControl) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BandwidthControl) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BandwidthControl) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetThresholdControlEnable

`func (o *BandwidthControl) GetThresholdControlEnable() bool`

GetThresholdControlEnable returns the ThresholdControlEnable field if non-nil, zero value otherwise.

### GetThresholdControlEnableOk

`func (o *BandwidthControl) GetThresholdControlEnableOk() (*bool, bool)`

GetThresholdControlEnableOk returns a tuple with the ThresholdControlEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdControlEnable

`func (o *BandwidthControl) SetThresholdControlEnable(v bool)`

SetThresholdControlEnable sets ThresholdControlEnable field to given value.

### HasThresholdControlEnable

`func (o *BandwidthControl) HasThresholdControlEnable() bool`

HasThresholdControlEnable returns a boolean if a field has been set.

### GetThresholdValue

`func (o *BandwidthControl) GetThresholdValue() int32`

GetThresholdValue returns the ThresholdValue field if non-nil, zero value otherwise.

### GetThresholdValueOk

`func (o *BandwidthControl) GetThresholdValueOk() (*int32, bool)`

GetThresholdValueOk returns a tuple with the ThresholdValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdValue

`func (o *BandwidthControl) SetThresholdValue(v int32)`

SetThresholdValue sets ThresholdValue field to given value.

### HasThresholdValue

`func (o *BandwidthControl) HasThresholdValue() bool`

HasThresholdValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


