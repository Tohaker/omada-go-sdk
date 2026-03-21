# BeaconControlOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BeaconInterval2g** | Pointer to **int32** | 2.4GHz beacon interval, parameter beaconInterval2g should be within the range of 40-500) | [optional] 
**BeaconInterval5g** | Pointer to **int32** | 5GHz beacon interval, parameter beaconInterval5g should be within the range of 40-500) | [optional] 
**BeaconInterval6g** | Pointer to **int32** | 6GHz beacon interval, parameter beaconInterval6g should be within the range of 40-500) | [optional] 
**BeaconIntvMode2g** | Pointer to **int32** | Configuration for Beacon Interval Mode in the 2G Band:0: Auto 1: Custom | [optional] 
**BeaconIntvMode5g** | Pointer to **int32** | Configuration for Beacon Interval Mode in the 5G Band:0: Auto 1: Custom | [optional] 
**BeaconIntvMode6g** | Pointer to **int32** | 6g band Beacon Interval mode: 0-auto, 1-custom | [optional] 
**DtimPeriod2g** | **int32** | 2.4GHz DTIM period, parameter dtimPeriod2g should be within the range of 1-255 | 
**DtimPeriod5g** | **int32** | 5GHz DTIM period, parameter dtimPeriod5g should be within the range of 1-255 | 
**DtimPeriod6g** | **int32** | 6GHz DTIM period, parameter dtimPeriod6g should be within the range of 1-255 | 
**RtsThreshold2g** | **int32** | 2.4GHz RTS threshold, parameter rtsThreshold2g should be within the range of 1-2347 | 
**RtsThreshold5g** | **int32** | 5GHz RTS threshold, parameter rtsThreshold5g should be within the range of 1-2347 | 
**RtsThreshold6g** | **int32** | 6GHz RTS threshold, parameter rtsThreshold6g should be within the range of 1-2347 | 

## Methods

### NewBeaconControlOpenApiVO

`func NewBeaconControlOpenApiVO(dtimPeriod2g int32, dtimPeriod5g int32, dtimPeriod6g int32, rtsThreshold2g int32, rtsThreshold5g int32, rtsThreshold6g int32, ) *BeaconControlOpenApiVO`

NewBeaconControlOpenApiVO instantiates a new BeaconControlOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBeaconControlOpenApiVOWithDefaults

`func NewBeaconControlOpenApiVOWithDefaults() *BeaconControlOpenApiVO`

NewBeaconControlOpenApiVOWithDefaults instantiates a new BeaconControlOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBeaconInterval2g

`func (o *BeaconControlOpenApiVO) GetBeaconInterval2g() int32`

GetBeaconInterval2g returns the BeaconInterval2g field if non-nil, zero value otherwise.

### GetBeaconInterval2gOk

`func (o *BeaconControlOpenApiVO) GetBeaconInterval2gOk() (*int32, bool)`

GetBeaconInterval2gOk returns a tuple with the BeaconInterval2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconInterval2g

`func (o *BeaconControlOpenApiVO) SetBeaconInterval2g(v int32)`

SetBeaconInterval2g sets BeaconInterval2g field to given value.

### HasBeaconInterval2g

`func (o *BeaconControlOpenApiVO) HasBeaconInterval2g() bool`

HasBeaconInterval2g returns a boolean if a field has been set.

### GetBeaconInterval5g

`func (o *BeaconControlOpenApiVO) GetBeaconInterval5g() int32`

GetBeaconInterval5g returns the BeaconInterval5g field if non-nil, zero value otherwise.

### GetBeaconInterval5gOk

`func (o *BeaconControlOpenApiVO) GetBeaconInterval5gOk() (*int32, bool)`

GetBeaconInterval5gOk returns a tuple with the BeaconInterval5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconInterval5g

`func (o *BeaconControlOpenApiVO) SetBeaconInterval5g(v int32)`

SetBeaconInterval5g sets BeaconInterval5g field to given value.

### HasBeaconInterval5g

`func (o *BeaconControlOpenApiVO) HasBeaconInterval5g() bool`

HasBeaconInterval5g returns a boolean if a field has been set.

### GetBeaconInterval6g

`func (o *BeaconControlOpenApiVO) GetBeaconInterval6g() int32`

GetBeaconInterval6g returns the BeaconInterval6g field if non-nil, zero value otherwise.

### GetBeaconInterval6gOk

`func (o *BeaconControlOpenApiVO) GetBeaconInterval6gOk() (*int32, bool)`

GetBeaconInterval6gOk returns a tuple with the BeaconInterval6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconInterval6g

`func (o *BeaconControlOpenApiVO) SetBeaconInterval6g(v int32)`

SetBeaconInterval6g sets BeaconInterval6g field to given value.

### HasBeaconInterval6g

`func (o *BeaconControlOpenApiVO) HasBeaconInterval6g() bool`

HasBeaconInterval6g returns a boolean if a field has been set.

### GetBeaconIntvMode2g

`func (o *BeaconControlOpenApiVO) GetBeaconIntvMode2g() int32`

GetBeaconIntvMode2g returns the BeaconIntvMode2g field if non-nil, zero value otherwise.

### GetBeaconIntvMode2gOk

`func (o *BeaconControlOpenApiVO) GetBeaconIntvMode2gOk() (*int32, bool)`

GetBeaconIntvMode2gOk returns a tuple with the BeaconIntvMode2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconIntvMode2g

`func (o *BeaconControlOpenApiVO) SetBeaconIntvMode2g(v int32)`

SetBeaconIntvMode2g sets BeaconIntvMode2g field to given value.

### HasBeaconIntvMode2g

`func (o *BeaconControlOpenApiVO) HasBeaconIntvMode2g() bool`

HasBeaconIntvMode2g returns a boolean if a field has been set.

### GetBeaconIntvMode5g

`func (o *BeaconControlOpenApiVO) GetBeaconIntvMode5g() int32`

GetBeaconIntvMode5g returns the BeaconIntvMode5g field if non-nil, zero value otherwise.

### GetBeaconIntvMode5gOk

`func (o *BeaconControlOpenApiVO) GetBeaconIntvMode5gOk() (*int32, bool)`

GetBeaconIntvMode5gOk returns a tuple with the BeaconIntvMode5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconIntvMode5g

`func (o *BeaconControlOpenApiVO) SetBeaconIntvMode5g(v int32)`

SetBeaconIntvMode5g sets BeaconIntvMode5g field to given value.

### HasBeaconIntvMode5g

`func (o *BeaconControlOpenApiVO) HasBeaconIntvMode5g() bool`

HasBeaconIntvMode5g returns a boolean if a field has been set.

### GetBeaconIntvMode6g

`func (o *BeaconControlOpenApiVO) GetBeaconIntvMode6g() int32`

GetBeaconIntvMode6g returns the BeaconIntvMode6g field if non-nil, zero value otherwise.

### GetBeaconIntvMode6gOk

`func (o *BeaconControlOpenApiVO) GetBeaconIntvMode6gOk() (*int32, bool)`

GetBeaconIntvMode6gOk returns a tuple with the BeaconIntvMode6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBeaconIntvMode6g

`func (o *BeaconControlOpenApiVO) SetBeaconIntvMode6g(v int32)`

SetBeaconIntvMode6g sets BeaconIntvMode6g field to given value.

### HasBeaconIntvMode6g

`func (o *BeaconControlOpenApiVO) HasBeaconIntvMode6g() bool`

HasBeaconIntvMode6g returns a boolean if a field has been set.

### GetDtimPeriod2g

`func (o *BeaconControlOpenApiVO) GetDtimPeriod2g() int32`

GetDtimPeriod2g returns the DtimPeriod2g field if non-nil, zero value otherwise.

### GetDtimPeriod2gOk

`func (o *BeaconControlOpenApiVO) GetDtimPeriod2gOk() (*int32, bool)`

GetDtimPeriod2gOk returns a tuple with the DtimPeriod2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtimPeriod2g

`func (o *BeaconControlOpenApiVO) SetDtimPeriod2g(v int32)`

SetDtimPeriod2g sets DtimPeriod2g field to given value.


### GetDtimPeriod5g

`func (o *BeaconControlOpenApiVO) GetDtimPeriod5g() int32`

GetDtimPeriod5g returns the DtimPeriod5g field if non-nil, zero value otherwise.

### GetDtimPeriod5gOk

`func (o *BeaconControlOpenApiVO) GetDtimPeriod5gOk() (*int32, bool)`

GetDtimPeriod5gOk returns a tuple with the DtimPeriod5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtimPeriod5g

`func (o *BeaconControlOpenApiVO) SetDtimPeriod5g(v int32)`

SetDtimPeriod5g sets DtimPeriod5g field to given value.


### GetDtimPeriod6g

`func (o *BeaconControlOpenApiVO) GetDtimPeriod6g() int32`

GetDtimPeriod6g returns the DtimPeriod6g field if non-nil, zero value otherwise.

### GetDtimPeriod6gOk

`func (o *BeaconControlOpenApiVO) GetDtimPeriod6gOk() (*int32, bool)`

GetDtimPeriod6gOk returns a tuple with the DtimPeriod6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtimPeriod6g

`func (o *BeaconControlOpenApiVO) SetDtimPeriod6g(v int32)`

SetDtimPeriod6g sets DtimPeriod6g field to given value.


### GetRtsThreshold2g

`func (o *BeaconControlOpenApiVO) GetRtsThreshold2g() int32`

GetRtsThreshold2g returns the RtsThreshold2g field if non-nil, zero value otherwise.

### GetRtsThreshold2gOk

`func (o *BeaconControlOpenApiVO) GetRtsThreshold2gOk() (*int32, bool)`

GetRtsThreshold2gOk returns a tuple with the RtsThreshold2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtsThreshold2g

`func (o *BeaconControlOpenApiVO) SetRtsThreshold2g(v int32)`

SetRtsThreshold2g sets RtsThreshold2g field to given value.


### GetRtsThreshold5g

`func (o *BeaconControlOpenApiVO) GetRtsThreshold5g() int32`

GetRtsThreshold5g returns the RtsThreshold5g field if non-nil, zero value otherwise.

### GetRtsThreshold5gOk

`func (o *BeaconControlOpenApiVO) GetRtsThreshold5gOk() (*int32, bool)`

GetRtsThreshold5gOk returns a tuple with the RtsThreshold5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtsThreshold5g

`func (o *BeaconControlOpenApiVO) SetRtsThreshold5g(v int32)`

SetRtsThreshold5g sets RtsThreshold5g field to given value.


### GetRtsThreshold6g

`func (o *BeaconControlOpenApiVO) GetRtsThreshold6g() int32`

GetRtsThreshold6g returns the RtsThreshold6g field if non-nil, zero value otherwise.

### GetRtsThreshold6gOk

`func (o *BeaconControlOpenApiVO) GetRtsThreshold6gOk() (*int32, bool)`

GetRtsThreshold6gOk returns a tuple with the RtsThreshold6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtsThreshold6g

`func (o *BeaconControlOpenApiVO) SetRtsThreshold6g(v int32)`

SetRtsThreshold6g sets RtsThreshold6g field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


