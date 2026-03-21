# ExtendOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Period** | **int64** | Extended timestamp. Unit:ms. Period should be within the range of 60000 to 86400000000000(60s to 1000000days). | 

## Methods

### NewExtendOpenApiVO

`func NewExtendOpenApiVO(period int64, ) *ExtendOpenApiVO`

NewExtendOpenApiVO instantiates a new ExtendOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendOpenApiVOWithDefaults

`func NewExtendOpenApiVOWithDefaults() *ExtendOpenApiVO`

NewExtendOpenApiVOWithDefaults instantiates a new ExtendOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriod

`func (o *ExtendOpenApiVO) GetPeriod() int64`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *ExtendOpenApiVO) GetPeriodOk() (*int64, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *ExtendOpenApiVO) SetPeriod(v int64)`

SetPeriod sets Period field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


