# OswQosRuleDeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagList** | Pointer to **[]int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**PortList** | Pointer to **[]string** |  | [optional] 
**StackDevice** | Pointer to **bool** |  | [optional] 

## Methods

### NewOswQosRuleDeviceVO

`func NewOswQosRuleDeviceVO() *OswQosRuleDeviceVO`

NewOswQosRuleDeviceVO instantiates a new OswQosRuleDeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswQosRuleDeviceVOWithDefaults

`func NewOswQosRuleDeviceVOWithDefaults() *OswQosRuleDeviceVO`

NewOswQosRuleDeviceVOWithDefaults instantiates a new OswQosRuleDeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagList

`func (o *OswQosRuleDeviceVO) GetLagList() []int32`

GetLagList returns the LagList field if non-nil, zero value otherwise.

### GetLagListOk

`func (o *OswQosRuleDeviceVO) GetLagListOk() (*[]int32, bool)`

GetLagListOk returns a tuple with the LagList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagList

`func (o *OswQosRuleDeviceVO) SetLagList(v []int32)`

SetLagList sets LagList field to given value.

### HasLagList

`func (o *OswQosRuleDeviceVO) HasLagList() bool`

HasLagList returns a boolean if a field has been set.

### GetMac

`func (o *OswQosRuleDeviceVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswQosRuleDeviceVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswQosRuleDeviceVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswQosRuleDeviceVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPortList

`func (o *OswQosRuleDeviceVO) GetPortList() []string`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *OswQosRuleDeviceVO) GetPortListOk() (*[]string, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *OswQosRuleDeviceVO) SetPortList(v []string)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *OswQosRuleDeviceVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetStackDevice

`func (o *OswQosRuleDeviceVO) GetStackDevice() bool`

GetStackDevice returns the StackDevice field if non-nil, zero value otherwise.

### GetStackDeviceOk

`func (o *OswQosRuleDeviceVO) GetStackDeviceOk() (*bool, bool)`

GetStackDeviceOk returns a tuple with the StackDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackDevice

`func (o *OswQosRuleDeviceVO) SetStackDevice(v bool)`

SetStackDevice sets StackDevice field to given value.

### HasStackDevice

`func (o *OswQosRuleDeviceVO) HasStackDevice() bool`

HasStackDevice returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


