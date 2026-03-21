# RFScanRadio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chan** | Pointer to **int32** | Channel number | [optional] 
**ChanWidth** | Pointer to **int32** | ChanWidth should be a value as follows: 2: 20MHz, 3: 40MHz, 5: 80MHz | [optional] 
**Inter** | Pointer to [**[]Interference**](Interference.md) | At most two types of interference data are reported in Inter, sorted in descending order of interference intensity | [optional] 
**Util** | Pointer to **int32** | Channel utilization should be within the range of 0–100. | [optional] 

## Methods

### NewRFScanRadio

`func NewRFScanRadio() *RFScanRadio`

NewRFScanRadio instantiates a new RFScanRadio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRFScanRadioWithDefaults

`func NewRFScanRadioWithDefaults() *RFScanRadio`

NewRFScanRadioWithDefaults instantiates a new RFScanRadio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChan

`func (o *RFScanRadio) GetChan() int32`

GetChan returns the Chan field if non-nil, zero value otherwise.

### GetChanOk

`func (o *RFScanRadio) GetChanOk() (*int32, bool)`

GetChanOk returns a tuple with the Chan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChan

`func (o *RFScanRadio) SetChan(v int32)`

SetChan sets Chan field to given value.

### HasChan

`func (o *RFScanRadio) HasChan() bool`

HasChan returns a boolean if a field has been set.

### GetChanWidth

`func (o *RFScanRadio) GetChanWidth() int32`

GetChanWidth returns the ChanWidth field if non-nil, zero value otherwise.

### GetChanWidthOk

`func (o *RFScanRadio) GetChanWidthOk() (*int32, bool)`

GetChanWidthOk returns a tuple with the ChanWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanWidth

`func (o *RFScanRadio) SetChanWidth(v int32)`

SetChanWidth sets ChanWidth field to given value.

### HasChanWidth

`func (o *RFScanRadio) HasChanWidth() bool`

HasChanWidth returns a boolean if a field has been set.

### GetInter

`func (o *RFScanRadio) GetInter() []Interference`

GetInter returns the Inter field if non-nil, zero value otherwise.

### GetInterOk

`func (o *RFScanRadio) GetInterOk() (*[]Interference, bool)`

GetInterOk returns a tuple with the Inter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInter

`func (o *RFScanRadio) SetInter(v []Interference)`

SetInter sets Inter field to given value.

### HasInter

`func (o *RFScanRadio) HasInter() bool`

HasInter returns a boolean if a field has been set.

### GetUtil

`func (o *RFScanRadio) GetUtil() int32`

GetUtil returns the Util field if non-nil, zero value otherwise.

### GetUtilOk

`func (o *RFScanRadio) GetUtilOk() (*int32, bool)`

GetUtilOk returns a tuple with the Util field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUtil

`func (o *RFScanRadio) SetUtil(v int32)`

SetUtil sets Util field to given value.

### HasUtil

`func (o *RFScanRadio) HasUtil() bool`

HasUtil returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


