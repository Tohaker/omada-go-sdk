# StatisticInfoDetailQueryDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePortIndex** | **int32** | ID of service port.ServicePortIndex should be within the range of 1 to 8100 | 

## Methods

### NewStatisticInfoDetailQueryDTO

`func NewStatisticInfoDetailQueryDTO(servicePortIndex int32, ) *StatisticInfoDetailQueryDTO`

NewStatisticInfoDetailQueryDTO instantiates a new StatisticInfoDetailQueryDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticInfoDetailQueryDTOWithDefaults

`func NewStatisticInfoDetailQueryDTOWithDefaults() *StatisticInfoDetailQueryDTO`

NewStatisticInfoDetailQueryDTOWithDefaults instantiates a new StatisticInfoDetailQueryDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePortIndex

`func (o *StatisticInfoDetailQueryDTO) GetServicePortIndex() int32`

GetServicePortIndex returns the ServicePortIndex field if non-nil, zero value otherwise.

### GetServicePortIndexOk

`func (o *StatisticInfoDetailQueryDTO) GetServicePortIndexOk() (*int32, bool)`

GetServicePortIndexOk returns a tuple with the ServicePortIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePortIndex

`func (o *StatisticInfoDetailQueryDTO) SetServicePortIndex(v int32)`

SetServicePortIndex sets ServicePortIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


