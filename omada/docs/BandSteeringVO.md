# BandSteeringVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionThreshold** | Pointer to **int32** |  | [optional] 
**DifferenceThreshold** | Pointer to **int32** |  | [optional] 
**Enable** | **bool** |  | 
**MaxFailures** | Pointer to **int32** |  | [optional] 

## Methods

### NewBandSteeringVO

`func NewBandSteeringVO(enable bool, ) *BandSteeringVO`

NewBandSteeringVO instantiates a new BandSteeringVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandSteeringVOWithDefaults

`func NewBandSteeringVOWithDefaults() *BandSteeringVO`

NewBandSteeringVOWithDefaults instantiates a new BandSteeringVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionThreshold

`func (o *BandSteeringVO) GetConnectionThreshold() int32`

GetConnectionThreshold returns the ConnectionThreshold field if non-nil, zero value otherwise.

### GetConnectionThresholdOk

`func (o *BandSteeringVO) GetConnectionThresholdOk() (*int32, bool)`

GetConnectionThresholdOk returns a tuple with the ConnectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionThreshold

`func (o *BandSteeringVO) SetConnectionThreshold(v int32)`

SetConnectionThreshold sets ConnectionThreshold field to given value.

### HasConnectionThreshold

`func (o *BandSteeringVO) HasConnectionThreshold() bool`

HasConnectionThreshold returns a boolean if a field has been set.

### GetDifferenceThreshold

`func (o *BandSteeringVO) GetDifferenceThreshold() int32`

GetDifferenceThreshold returns the DifferenceThreshold field if non-nil, zero value otherwise.

### GetDifferenceThresholdOk

`func (o *BandSteeringVO) GetDifferenceThresholdOk() (*int32, bool)`

GetDifferenceThresholdOk returns a tuple with the DifferenceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferenceThreshold

`func (o *BandSteeringVO) SetDifferenceThreshold(v int32)`

SetDifferenceThreshold sets DifferenceThreshold field to given value.

### HasDifferenceThreshold

`func (o *BandSteeringVO) HasDifferenceThreshold() bool`

HasDifferenceThreshold returns a boolean if a field has been set.

### GetEnable

`func (o *BandSteeringVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BandSteeringVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BandSteeringVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMaxFailures

`func (o *BandSteeringVO) GetMaxFailures() int32`

GetMaxFailures returns the MaxFailures field if non-nil, zero value otherwise.

### GetMaxFailuresOk

`func (o *BandSteeringVO) GetMaxFailuresOk() (*int32, bool)`

GetMaxFailuresOk returns a tuple with the MaxFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailures

`func (o *BandSteeringVO) SetMaxFailures(v int32)`

SetMaxFailures sets MaxFailures field to given value.

### HasMaxFailures

`func (o *BandSteeringVO) HasMaxFailures() bool`

HasMaxFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


