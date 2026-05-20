# IotAgingTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingTimeIotAgingTime** | **int32** | The system automatically removes a device&#39;s registry entry if no data reports are received within a predefined aging period.&lt;br/&gt;When format &#x3D; 0, The parameter aging time should be within the range of 30-86400.&lt;br/&gt;When format &#x3D; 1, The parameter aging time should be within the range of 1-1440.&lt;br/&gt;When format &#x3D; 2, The parameter aging time should be within the range of 1-24.&lt;br/&gt; | 
**FormatIotAgingTime** | **int32** | The parameter [format] should be a value as follows: [0:second 1:minute; 2:hour] | 

## Methods

### NewIotAgingTimeOpenApiVO

`func NewIotAgingTimeOpenApiVO(agingTimeIotAgingTime int32, formatIotAgingTime int32, ) *IotAgingTimeOpenApiVO`

NewIotAgingTimeOpenApiVO instantiates a new IotAgingTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotAgingTimeOpenApiVOWithDefaults

`func NewIotAgingTimeOpenApiVOWithDefaults() *IotAgingTimeOpenApiVO`

NewIotAgingTimeOpenApiVOWithDefaults instantiates a new IotAgingTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingTimeIotAgingTime

`func (o *IotAgingTimeOpenApiVO) GetAgingTimeIotAgingTime() int32`

GetAgingTimeIotAgingTime returns the AgingTimeIotAgingTime field if non-nil, zero value otherwise.

### GetAgingTimeIotAgingTimeOk

`func (o *IotAgingTimeOpenApiVO) GetAgingTimeIotAgingTimeOk() (*int32, bool)`

GetAgingTimeIotAgingTimeOk returns a tuple with the AgingTimeIotAgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingTimeIotAgingTime

`func (o *IotAgingTimeOpenApiVO) SetAgingTimeIotAgingTime(v int32)`

SetAgingTimeIotAgingTime sets AgingTimeIotAgingTime field to given value.


### GetFormatIotAgingTime

`func (o *IotAgingTimeOpenApiVO) GetFormatIotAgingTime() int32`

GetFormatIotAgingTime returns the FormatIotAgingTime field if non-nil, zero value otherwise.

### GetFormatIotAgingTimeOk

`func (o *IotAgingTimeOpenApiVO) GetFormatIotAgingTimeOk() (*int32, bool)`

GetFormatIotAgingTimeOk returns a tuple with the FormatIotAgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormatIotAgingTime

`func (o *IotAgingTimeOpenApiVO) SetFormatIotAgingTime(v int32)`

SetFormatIotAgingTime sets FormatIotAgingTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


