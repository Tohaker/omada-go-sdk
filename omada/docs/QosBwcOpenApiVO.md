# QosBwcOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassRatio** | **[]int32** | The ratio of class type, value&#39;s format is [class1 ratio, class2 ratio, class3 ratio, others ratio], and the total sum should be 100. | 
**Direction** | **int32** | The direction value selected in the Direction configuration should be a value as follows: 0: Inbound; 1: Outbound; 2: Both. | 
**InBandwidth** | **int32** | The Inbound Bandwidth of Bandwidth Control rule should be within the range of 100-1000000. | 
**OutBandwidth** | **int32** | The Outbound Bandwidth of Bandwidth Control rule should be within the range of 100-1000000. | 
**OutPrioritization** | **bool** | The Outbound TCP ACK Prioritize of Bandwidth Control rule, valid values are true or false. | 
**Status** | **bool** | The status of Bandwidth Control rule, valid values are true or false. | 
**UdpBandwidthCtrl** | **bool** | The UDP Bandwidth Control of Bandwidth Control rule, valid values are true or false. | 
**UdpRatio** | Pointer to **int32** | The Limited Bandwidth Ratio of Bandwidth Control rule. It should be within the range of 0-100 when parameter [udpBandwidthCtrl] is true. | [optional] 
**Wan** | **string** | The ID of WAN port to which this Bandwidth Control applies. Each WAN port can only be configured with one rule. WAN ID can be obtained from &#39;Get WAN ports info for Gateway QoS&#39; interface. | 

## Methods

### NewQosBwcOpenApiVO

`func NewQosBwcOpenApiVO(classRatio []int32, direction int32, inBandwidth int32, outBandwidth int32, outPrioritization bool, status bool, udpBandwidthCtrl bool, wan string, ) *QosBwcOpenApiVO`

NewQosBwcOpenApiVO instantiates a new QosBwcOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosBwcOpenApiVOWithDefaults

`func NewQosBwcOpenApiVOWithDefaults() *QosBwcOpenApiVO`

NewQosBwcOpenApiVOWithDefaults instantiates a new QosBwcOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassRatio

`func (o *QosBwcOpenApiVO) GetClassRatio() []int32`

GetClassRatio returns the ClassRatio field if non-nil, zero value otherwise.

### GetClassRatioOk

`func (o *QosBwcOpenApiVO) GetClassRatioOk() (*[]int32, bool)`

GetClassRatioOk returns a tuple with the ClassRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassRatio

`func (o *QosBwcOpenApiVO) SetClassRatio(v []int32)`

SetClassRatio sets ClassRatio field to given value.


### GetDirection

`func (o *QosBwcOpenApiVO) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QosBwcOpenApiVO) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QosBwcOpenApiVO) SetDirection(v int32)`

SetDirection sets Direction field to given value.


### GetInBandwidth

`func (o *QosBwcOpenApiVO) GetInBandwidth() int32`

GetInBandwidth returns the InBandwidth field if non-nil, zero value otherwise.

### GetInBandwidthOk

`func (o *QosBwcOpenApiVO) GetInBandwidthOk() (*int32, bool)`

GetInBandwidthOk returns a tuple with the InBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBandwidth

`func (o *QosBwcOpenApiVO) SetInBandwidth(v int32)`

SetInBandwidth sets InBandwidth field to given value.


### GetOutBandwidth

`func (o *QosBwcOpenApiVO) GetOutBandwidth() int32`

GetOutBandwidth returns the OutBandwidth field if non-nil, zero value otherwise.

### GetOutBandwidthOk

`func (o *QosBwcOpenApiVO) GetOutBandwidthOk() (*int32, bool)`

GetOutBandwidthOk returns a tuple with the OutBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutBandwidth

`func (o *QosBwcOpenApiVO) SetOutBandwidth(v int32)`

SetOutBandwidth sets OutBandwidth field to given value.


### GetOutPrioritization

`func (o *QosBwcOpenApiVO) GetOutPrioritization() bool`

GetOutPrioritization returns the OutPrioritization field if non-nil, zero value otherwise.

### GetOutPrioritizationOk

`func (o *QosBwcOpenApiVO) GetOutPrioritizationOk() (*bool, bool)`

GetOutPrioritizationOk returns a tuple with the OutPrioritization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutPrioritization

`func (o *QosBwcOpenApiVO) SetOutPrioritization(v bool)`

SetOutPrioritization sets OutPrioritization field to given value.


### GetStatus

`func (o *QosBwcOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QosBwcOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QosBwcOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetUdpBandwidthCtrl

`func (o *QosBwcOpenApiVO) GetUdpBandwidthCtrl() bool`

GetUdpBandwidthCtrl returns the UdpBandwidthCtrl field if non-nil, zero value otherwise.

### GetUdpBandwidthCtrlOk

`func (o *QosBwcOpenApiVO) GetUdpBandwidthCtrlOk() (*bool, bool)`

GetUdpBandwidthCtrlOk returns a tuple with the UdpBandwidthCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpBandwidthCtrl

`func (o *QosBwcOpenApiVO) SetUdpBandwidthCtrl(v bool)`

SetUdpBandwidthCtrl sets UdpBandwidthCtrl field to given value.


### GetUdpRatio

`func (o *QosBwcOpenApiVO) GetUdpRatio() int32`

GetUdpRatio returns the UdpRatio field if non-nil, zero value otherwise.

### GetUdpRatioOk

`func (o *QosBwcOpenApiVO) GetUdpRatioOk() (*int32, bool)`

GetUdpRatioOk returns a tuple with the UdpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpRatio

`func (o *QosBwcOpenApiVO) SetUdpRatio(v int32)`

SetUdpRatio sets UdpRatio field to given value.

### HasUdpRatio

`func (o *QosBwcOpenApiVO) HasUdpRatio() bool`

HasUdpRatio returns a boolean if a field has been set.

### GetWan

`func (o *QosBwcOpenApiVO) GetWan() string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *QosBwcOpenApiVO) GetWanOk() (*string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *QosBwcOpenApiVO) SetWan(v string)`

SetWan sets Wan field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


