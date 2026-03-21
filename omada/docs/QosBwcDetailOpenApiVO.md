# QosBwcDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassRatio** | Pointer to **[]int32** | The ratio of class type, value&#39;s format is [class1 ratio, class2 ratio, class3 ratio, others ratio]. | [optional] 
**Direction** | Pointer to **int32** | The direction value selected in the Direction configuration should be a value as follows: 0: Inbound; 1: Outbound; 2: Both. | [optional] 
**Id** | Pointer to **string** | The ID of Bandwidth Control rule. | [optional] 
**InBandwidth** | Pointer to **int32** | The Inbound Bandwidth of Bandwidth Control rule should be within the range of 100-1000000. | [optional] 
**OutBandwidth** | Pointer to **int32** | The Outbound Bandwidth of Bandwidth Control rule should be within the range of 100-1000000. | [optional] 
**OutPrioritization** | Pointer to **bool** | The Outbound TCP ACK Prioritize of Bandwidth Control rule, valid values are true or false. | [optional] 
**Status** | Pointer to **bool** | The status of Bandwidth Control rule, valid values are true or false. | [optional] 
**UdpBandwidthCtrl** | Pointer to **bool** | The UDP Bandwidth Control of Bandwidth Control rule, valid values are true or false. | [optional] 
**UdpRatio** | Pointer to **int32** | The Limited Bandwidth Ratio of Bandwidth Control rule. It should be within the range of 0-100 when parameter [udpBandwidthCtrl] is true. | [optional] 
**Wan** | Pointer to **string** | The ID of WAN port to which this Bandwidth Control applies. | [optional] 

## Methods

### NewQosBwcDetailOpenApiVO

`func NewQosBwcDetailOpenApiVO() *QosBwcDetailOpenApiVO`

NewQosBwcDetailOpenApiVO instantiates a new QosBwcDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQosBwcDetailOpenApiVOWithDefaults

`func NewQosBwcDetailOpenApiVOWithDefaults() *QosBwcDetailOpenApiVO`

NewQosBwcDetailOpenApiVOWithDefaults instantiates a new QosBwcDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassRatio

`func (o *QosBwcDetailOpenApiVO) GetClassRatio() []int32`

GetClassRatio returns the ClassRatio field if non-nil, zero value otherwise.

### GetClassRatioOk

`func (o *QosBwcDetailOpenApiVO) GetClassRatioOk() (*[]int32, bool)`

GetClassRatioOk returns a tuple with the ClassRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassRatio

`func (o *QosBwcDetailOpenApiVO) SetClassRatio(v []int32)`

SetClassRatio sets ClassRatio field to given value.

### HasClassRatio

`func (o *QosBwcDetailOpenApiVO) HasClassRatio() bool`

HasClassRatio returns a boolean if a field has been set.

### GetDirection

`func (o *QosBwcDetailOpenApiVO) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *QosBwcDetailOpenApiVO) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *QosBwcDetailOpenApiVO) SetDirection(v int32)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *QosBwcDetailOpenApiVO) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *QosBwcDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QosBwcDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QosBwcDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *QosBwcDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInBandwidth

`func (o *QosBwcDetailOpenApiVO) GetInBandwidth() int32`

GetInBandwidth returns the InBandwidth field if non-nil, zero value otherwise.

### GetInBandwidthOk

`func (o *QosBwcDetailOpenApiVO) GetInBandwidthOk() (*int32, bool)`

GetInBandwidthOk returns a tuple with the InBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInBandwidth

`func (o *QosBwcDetailOpenApiVO) SetInBandwidth(v int32)`

SetInBandwidth sets InBandwidth field to given value.

### HasInBandwidth

`func (o *QosBwcDetailOpenApiVO) HasInBandwidth() bool`

HasInBandwidth returns a boolean if a field has been set.

### GetOutBandwidth

`func (o *QosBwcDetailOpenApiVO) GetOutBandwidth() int32`

GetOutBandwidth returns the OutBandwidth field if non-nil, zero value otherwise.

### GetOutBandwidthOk

`func (o *QosBwcDetailOpenApiVO) GetOutBandwidthOk() (*int32, bool)`

GetOutBandwidthOk returns a tuple with the OutBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutBandwidth

`func (o *QosBwcDetailOpenApiVO) SetOutBandwidth(v int32)`

SetOutBandwidth sets OutBandwidth field to given value.

### HasOutBandwidth

`func (o *QosBwcDetailOpenApiVO) HasOutBandwidth() bool`

HasOutBandwidth returns a boolean if a field has been set.

### GetOutPrioritization

`func (o *QosBwcDetailOpenApiVO) GetOutPrioritization() bool`

GetOutPrioritization returns the OutPrioritization field if non-nil, zero value otherwise.

### GetOutPrioritizationOk

`func (o *QosBwcDetailOpenApiVO) GetOutPrioritizationOk() (*bool, bool)`

GetOutPrioritizationOk returns a tuple with the OutPrioritization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutPrioritization

`func (o *QosBwcDetailOpenApiVO) SetOutPrioritization(v bool)`

SetOutPrioritization sets OutPrioritization field to given value.

### HasOutPrioritization

`func (o *QosBwcDetailOpenApiVO) HasOutPrioritization() bool`

HasOutPrioritization returns a boolean if a field has been set.

### GetStatus

`func (o *QosBwcDetailOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *QosBwcDetailOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *QosBwcDetailOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *QosBwcDetailOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUdpBandwidthCtrl

`func (o *QosBwcDetailOpenApiVO) GetUdpBandwidthCtrl() bool`

GetUdpBandwidthCtrl returns the UdpBandwidthCtrl field if non-nil, zero value otherwise.

### GetUdpBandwidthCtrlOk

`func (o *QosBwcDetailOpenApiVO) GetUdpBandwidthCtrlOk() (*bool, bool)`

GetUdpBandwidthCtrlOk returns a tuple with the UdpBandwidthCtrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpBandwidthCtrl

`func (o *QosBwcDetailOpenApiVO) SetUdpBandwidthCtrl(v bool)`

SetUdpBandwidthCtrl sets UdpBandwidthCtrl field to given value.

### HasUdpBandwidthCtrl

`func (o *QosBwcDetailOpenApiVO) HasUdpBandwidthCtrl() bool`

HasUdpBandwidthCtrl returns a boolean if a field has been set.

### GetUdpRatio

`func (o *QosBwcDetailOpenApiVO) GetUdpRatio() int32`

GetUdpRatio returns the UdpRatio field if non-nil, zero value otherwise.

### GetUdpRatioOk

`func (o *QosBwcDetailOpenApiVO) GetUdpRatioOk() (*int32, bool)`

GetUdpRatioOk returns a tuple with the UdpRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUdpRatio

`func (o *QosBwcDetailOpenApiVO) SetUdpRatio(v int32)`

SetUdpRatio sets UdpRatio field to given value.

### HasUdpRatio

`func (o *QosBwcDetailOpenApiVO) HasUdpRatio() bool`

HasUdpRatio returns a boolean if a field has been set.

### GetWan

`func (o *QosBwcDetailOpenApiVO) GetWan() string`

GetWan returns the Wan field if non-nil, zero value otherwise.

### GetWanOk

`func (o *QosBwcDetailOpenApiVO) GetWanOk() (*string, bool)`

GetWanOk returns a tuple with the Wan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWan

`func (o *QosBwcDetailOpenApiVO) SetWan(v string)`

SetWan sets Wan field to given value.

### HasWan

`func (o *QosBwcDetailOpenApiVO) HasWan() bool`

HasWan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


