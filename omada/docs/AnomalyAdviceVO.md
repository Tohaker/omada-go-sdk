# AnomalyAdviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentParams** | Pointer to **map[string]string** | Content parameter map for rendering the advice content template. | [optional] 
**Devices** | Pointer to [**map[string]DeviceObjectDTO**](DeviceObjectDTO.md) | Device objects map used by this advice. Key is MAC address, value is device info. | [optional] 
**Helpful** | Pointer to **int32** | Whether the advice is helpful. 0: Not set, 1: Helpful, 2: Not helpful. | [optional] 

## Methods

### NewAnomalyAdviceVO

`func NewAnomalyAdviceVO() *AnomalyAdviceVO`

NewAnomalyAdviceVO instantiates a new AnomalyAdviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyAdviceVOWithDefaults

`func NewAnomalyAdviceVOWithDefaults() *AnomalyAdviceVO`

NewAnomalyAdviceVOWithDefaults instantiates a new AnomalyAdviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentParams

`func (o *AnomalyAdviceVO) GetContentParams() map[string]string`

GetContentParams returns the ContentParams field if non-nil, zero value otherwise.

### GetContentParamsOk

`func (o *AnomalyAdviceVO) GetContentParamsOk() (*map[string]string, bool)`

GetContentParamsOk returns a tuple with the ContentParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentParams

`func (o *AnomalyAdviceVO) SetContentParams(v map[string]string)`

SetContentParams sets ContentParams field to given value.

### HasContentParams

`func (o *AnomalyAdviceVO) HasContentParams() bool`

HasContentParams returns a boolean if a field has been set.

### GetDevices

`func (o *AnomalyAdviceVO) GetDevices() map[string]DeviceObjectDTO`

GetDevices returns the Devices field if non-nil, zero value otherwise.

### GetDevicesOk

`func (o *AnomalyAdviceVO) GetDevicesOk() (*map[string]DeviceObjectDTO, bool)`

GetDevicesOk returns a tuple with the Devices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevices

`func (o *AnomalyAdviceVO) SetDevices(v map[string]DeviceObjectDTO)`

SetDevices sets Devices field to given value.

### HasDevices

`func (o *AnomalyAdviceVO) HasDevices() bool`

HasDevices returns a boolean if a field has been set.

### GetHelpful

`func (o *AnomalyAdviceVO) GetHelpful() int32`

GetHelpful returns the Helpful field if non-nil, zero value otherwise.

### GetHelpfulOk

`func (o *AnomalyAdviceVO) GetHelpfulOk() (*int32, bool)`

GetHelpfulOk returns a tuple with the Helpful field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHelpful

`func (o *AnomalyAdviceVO) SetHelpful(v int32)`

SetHelpful sets Helpful field to given value.

### HasHelpful

`func (o *AnomalyAdviceVO) HasHelpful() bool`

HasHelpful returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


