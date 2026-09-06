# OswQosRuleResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailDevices** | Pointer to [**[]QosRuleDeviceInfo**](QosRuleDeviceInfo.md) | List of device information for which QoS rules failed to be issued. | [optional] 
**SuccessDevices** | Pointer to [**[]QosRuleDeviceInfo**](QosRuleDeviceInfo.md) | List of device information for which QoS rules have been successfully issued. | [optional] 

## Methods

### NewOswQosRuleResultVO

`func NewOswQosRuleResultVO() *OswQosRuleResultVO`

NewOswQosRuleResultVO instantiates a new OswQosRuleResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswQosRuleResultVOWithDefaults

`func NewOswQosRuleResultVOWithDefaults() *OswQosRuleResultVO`

NewOswQosRuleResultVOWithDefaults instantiates a new OswQosRuleResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailDevices

`func (o *OswQosRuleResultVO) GetFailDevices() []QosRuleDeviceInfo`

GetFailDevices returns the FailDevices field if non-nil, zero value otherwise.

### GetFailDevicesOk

`func (o *OswQosRuleResultVO) GetFailDevicesOk() (*[]QosRuleDeviceInfo, bool)`

GetFailDevicesOk returns a tuple with the FailDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailDevices

`func (o *OswQosRuleResultVO) SetFailDevices(v []QosRuleDeviceInfo)`

SetFailDevices sets FailDevices field to given value.

### HasFailDevices

`func (o *OswQosRuleResultVO) HasFailDevices() bool`

HasFailDevices returns a boolean if a field has been set.

### GetSuccessDevices

`func (o *OswQosRuleResultVO) GetSuccessDevices() []QosRuleDeviceInfo`

GetSuccessDevices returns the SuccessDevices field if non-nil, zero value otherwise.

### GetSuccessDevicesOk

`func (o *OswQosRuleResultVO) GetSuccessDevicesOk() (*[]QosRuleDeviceInfo, bool)`

GetSuccessDevicesOk returns a tuple with the SuccessDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessDevices

`func (o *OswQosRuleResultVO) SetSuccessDevices(v []QosRuleDeviceInfo)`

SetSuccessDevices sets SuccessDevices field to given value.

### HasSuccessDevices

`func (o *OswQosRuleResultVO) HasSuccessDevices() bool`

HasSuccessDevices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


