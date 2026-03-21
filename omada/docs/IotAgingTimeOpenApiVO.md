# IotAgingTimeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgingTime** | **int32** | The system automatically removes a device&#39;s registry entry if no data reports are received within a predefined aging period.&lt;br/&gt;When format &#x3D; 0, The parameter aging time should be within the range of 30-86400.&lt;br/&gt;When format &#x3D; 1, The parameter aging time should be within the range of 1-1440.&lt;br/&gt;When format &#x3D; 2, The parameter aging time should be within the range of 1-24.&lt;br/&gt; | 
**Format** | **int32** | The parameter [format] should be a value as follows: [0:second 1:minute; 2:hour] | 

## Methods

### NewIotAgingTimeOpenApiVO

`func NewIotAgingTimeOpenApiVO(agingTime int32, format int32, ) *IotAgingTimeOpenApiVO`

NewIotAgingTimeOpenApiVO instantiates a new IotAgingTimeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIotAgingTimeOpenApiVOWithDefaults

`func NewIotAgingTimeOpenApiVOWithDefaults() *IotAgingTimeOpenApiVO`

NewIotAgingTimeOpenApiVOWithDefaults instantiates a new IotAgingTimeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgingTime

`func (o *IotAgingTimeOpenApiVO) GetAgingTime() int32`

GetAgingTime returns the AgingTime field if non-nil, zero value otherwise.

### GetAgingTimeOk

`func (o *IotAgingTimeOpenApiVO) GetAgingTimeOk() (*int32, bool)`

GetAgingTimeOk returns a tuple with the AgingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgingTime

`func (o *IotAgingTimeOpenApiVO) SetAgingTime(v int32)`

SetAgingTime sets AgingTime field to given value.


### GetFormat

`func (o *IotAgingTimeOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *IotAgingTimeOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *IotAgingTimeOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


