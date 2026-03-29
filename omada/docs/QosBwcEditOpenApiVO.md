# QosBwcEditOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassRatio** | **[]int32** | The ratio of class type, value&#39;s format is [class1 ratio, class2 ratio, class3 ratio, others ratio]. | 
**Direction** | **int32** | The direction value selected in the Direction configuration should be a value as follows: 0: Inbound; 1: Outbound; 2: Both. | 
**InBandwidth** | **int32** | The Inbound Bandwidth of Bandwidth Control rule should be within the range of 100-1000000. | 
**OutBandwidth** | **int32** | The Outbound Bandwidth of Bandwidth Control rule should be within the range of 100-1000000. | 
**OutPrioritization** | **bool** | The Outbound TCP ACK Prioritize of Bandwidth Control rule, valid values is true or false. | 
**Status** | **bool** | The status of Bandwidth Control rule, valid values is true or false. | 
**UdpBandwidthCtrl** | **bool** | The UDP Bandwidth Control of Bandwidth Control rule, valid values is true or false. | 
**UdpRatio** | Pointer to **int32** | The Limited Bandwidth Ratio of Bandwidth Control rule. It should be within the range of 0-100 when parameter [udpBandwidthCtrl] is true. | [optional] 
**Wan** | Pointer to **string** | The wan of Bandwidth Control rule. | [optional] 

## Methods

### NewQosBwcEditOpenApiVO

`func NewQosBwcEditOpenApiVO(classRatio []int32, direction int32, inBandwidth int32, outBandwidth int32, outPrioritization bool, status bool, udpBandwidthCtrl bool, ) *QosBwcEditOpenApiVO`

NewQosBwcEditOpenApiVO instantiates a new QosBwcEditOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosBwcEditOpenApiVOWithDefaults

`func NewQosBwcEditOpenApiVOWithDefaults() *QosBwcEditOpenApiVO`

NewQosBwcEditOpenApiVOWithDefaults instantiates a new QosBwcEditOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassRatio

`func (o *QosBwcEditOpenApiVO) GetClassRatio() []int32`

GetClassRatio returns the ClassRatio field if non-nil, zero value otherwise.

### GetClassRatioOk

`func (o *QosBwcEditOpenApiVO) GetClassRatioOk() (*[]int32, bool)`

GetClassRatioOk returns a tuple with the ClassRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassRatio

`func (o *QosBwcEditOpenApiVO) SetClassRatio(v []int32)`

SetClassRatio sets ClassRatio field to given value.


### GetDirection

`func (o *QosBwcEditOpenApiVO) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QosBwcEditOpenApiVO) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QosBwcEditOpenApiVO) SetDirection(v int32)`

SetDirection sets Direction field to given value.


### GetInBandwidth

`func (o *QosBwcEditOpenApiVO) GetInBandwidth() int32`

GetInBandwidth returns the InBandwidth field if non-nil, zero value otherwise.

### GetInBandwidthOk

`func (o *QosBwcEditOpenApiVO) GetInBandwidthOk() (*int32, bool)`

GetInBandwidthOk returns a tuple with the InBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBandwidth

`func (o *QosBwcEditOpenApiVO) SetInBandwidth(v int32)`

SetInBandwidth sets InBandwidth field to given value.


### GetOutBandwidth

`func (o *QosBwcEditOpenApiVO) GetOutBandwidth() int32`

GetOutBandwidth returns the OutBandwidth field if non-nil, zero value otherwise.

### GetOutBandwidthOk

`func (o *QosBwcEditOpenApiVO) GetOutBandwidthOk() (*int32, bool)`

GetOutBandwidthOk returns a tuple with the OutBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutBandwidth

`func (o *QosBwcEditOpenApiVO) SetOutBandwidth(v int32)`

SetOutBandwidth sets OutBandwidth field to given value.


### GetOutPrioritization

`func (o *QosBwcEditOpenApiVO) GetOutPrioritization() bool`

GetOutPrioritization returns the OutPrioritization field if non-nil, zero value otherwise.

### GetOutPrioritizationOk

`func (o *QosBwcEditOpenApiVO) GetOutPrioritizationOk() (*bool, bool)`

GetOutPrioritizationOk returns a tuple with the OutPrioritization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutPrioritization

`func (o *QosBwcEditOpenApiVO) SetOutPrioritization(v bool)`

SetOutPrioritization sets OutPrioritization field to given value.


### GetStatus

`func (o *QosBwcEditOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QosBwcEditOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QosBwcEditOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetUdpBandwidthCtrl

`func (o *QosBwcEditOpenApiVO) GetUdpBandwidthCtrl() bool`

GetUdpBandwidthCtrl returns the UdpBandwidthCtrl field if non-nil, zero value otherwise.

### GetUdpBandwidthCtrlOk

`func (o *QosBwcEditOpenApiVO) GetUdpBandwidthCtrlOk() (*bool, bool)`

GetUdpBandwidthCtrlOk returns a tuple with the UdpBandwidthCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpBandwidthCtrl

`func (o *QosBwcEditOpenApiVO) SetUdpBandwidthCtrl(v bool)`

SetUdpBandwidthCtrl sets UdpBandwidthCtrl field to given value.


### GetUdpRatio

`func (o *QosBwcEditOpenApiVO) GetUdpRatio() int32`

GetUdpRatio returns the UdpRatio field if non-nil, zero value otherwise.

### GetUdpRatioOk

`func (o *QosBwcEditOpenApiVO) GetUdpRatioOk() (*int32, bool)`

GetUdpRatioOk returns a tuple with the UdpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpRatio

`func (o *QosBwcEditOpenApiVO) SetUdpRatio(v int32)`

SetUdpRatio sets UdpRatio field to given value.

### HasUdpRatio

`func (o *QosBwcEditOpenApiVO) HasUdpRatio() bool`

HasUdpRatio returns a boolean if a field has been set.

### GetWan

`func (o *QosBwcEditOpenApiVO) GetWan() string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *QosBwcEditOpenApiVO) GetWanOk() (*string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *QosBwcEditOpenApiVO) SetWan(v string)`

SetWan sets Wan field to given value.

### HasWan

`func (o *QosBwcEditOpenApiVO) HasWan() bool`

HasWan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


