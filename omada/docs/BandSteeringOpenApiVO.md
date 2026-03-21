# BandSteeringOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionThreshold** | Pointer to **int32** | The range of connect threshold is from 2 to 256, with a default of 30. | [optional] 
**DifferenceThreshold** | Pointer to **int32** | The range of difference threshold is from 1 to 20, with a default of 4. | [optional] 
**Enable** | **bool** | Band steer enable. | 
**MaxFailures** | Pointer to **int32** | The range of maxFailures threshold is from 1 to 20, with a default of 5. | [optional] 

## Methods

### NewBandSteeringOpenApiVO

`func NewBandSteeringOpenApiVO(enable bool, ) *BandSteeringOpenApiVO`

NewBandSteeringOpenApiVO instantiates a new BandSteeringOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandSteeringOpenApiVOWithDefaults

`func NewBandSteeringOpenApiVOWithDefaults() *BandSteeringOpenApiVO`

NewBandSteeringOpenApiVOWithDefaults instantiates a new BandSteeringOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionThreshold

`func (o *BandSteeringOpenApiVO) GetConnectionThreshold() int32`

GetConnectionThreshold returns the ConnectionThreshold field if non-nil, zero value otherwise.

### GetConnectionThresholdOk

`func (o *BandSteeringOpenApiVO) GetConnectionThresholdOk() (*int32, bool)`

GetConnectionThresholdOk returns a tuple with the ConnectionThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionThreshold

`func (o *BandSteeringOpenApiVO) SetConnectionThreshold(v int32)`

SetConnectionThreshold sets ConnectionThreshold field to given value.

### HasConnectionThreshold

`func (o *BandSteeringOpenApiVO) HasConnectionThreshold() bool`

HasConnectionThreshold returns a boolean if a field has been set.

### GetDifferenceThreshold

`func (o *BandSteeringOpenApiVO) GetDifferenceThreshold() int32`

GetDifferenceThreshold returns the DifferenceThreshold field if non-nil, zero value otherwise.

### GetDifferenceThresholdOk

`func (o *BandSteeringOpenApiVO) GetDifferenceThresholdOk() (*int32, bool)`

GetDifferenceThresholdOk returns a tuple with the DifferenceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDifferenceThreshold

`func (o *BandSteeringOpenApiVO) SetDifferenceThreshold(v int32)`

SetDifferenceThreshold sets DifferenceThreshold field to given value.

### HasDifferenceThreshold

`func (o *BandSteeringOpenApiVO) HasDifferenceThreshold() bool`

HasDifferenceThreshold returns a boolean if a field has been set.

### GetEnable

`func (o *BandSteeringOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *BandSteeringOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *BandSteeringOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetMaxFailures

`func (o *BandSteeringOpenApiVO) GetMaxFailures() int32`

GetMaxFailures returns the MaxFailures field if non-nil, zero value otherwise.

### GetMaxFailuresOk

`func (o *BandSteeringOpenApiVO) GetMaxFailuresOk() (*int32, bool)`

GetMaxFailuresOk returns a tuple with the MaxFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxFailures

`func (o *BandSteeringOpenApiVO) SetMaxFailures(v int32)`

SetMaxFailures sets MaxFailures field to given value.

### HasMaxFailures

`func (o *BandSteeringOpenApiVO) HasMaxFailures() bool`

HasMaxFailures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


