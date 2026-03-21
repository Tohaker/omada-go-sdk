# IgmpConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap, gateway, switch, olt | [optional] 
**GeneralQuerySourceIp** | Pointer to **string** | The source IP of the general query message sent. | [optional] 
**LastMemberQueryCount** | Pointer to **int32** | The number of specific group query messages sent by the device. | [optional] 
**LastMemberQueryInterval** | Pointer to **int32** | The time interval for the device to send a specific group of query messages (unit: s). | [optional] 
**Mac** | Pointer to **string** | The unique identification of the device. | [optional] 
**MaximumResponseTime** | Pointer to **int32** | The maximum response time of the host to the general query message (unit: s). | [optional] 
**NetworkId** | Pointer to **string** | The primary id of the network. | [optional] 
**NetworkName** | Pointer to **string** | The name of the network.Name should contain 1 to 128 characters. | [optional] 
**OswStack** | Pointer to [**OswStackInfoVO**](OswStackInfoVO.md) |  | [optional] 
**QueryInterval** | Pointer to **int32** | The time interval for the IGMP query to send query messages (unit: s). | [optional] 
**StackId** | Pointer to **string** | The id of the stacking devices. | [optional] 
**StackName** | Pointer to **string** | The name of the stacking devices. | [optional] 
**Vlan** | Pointer to **int32** | The vlan of the network.The vlan should be within the range of 1 to 4090. | [optional] 

## Methods

### NewIgmpConfigVO

`func NewIgmpConfigVO() *IgmpConfigVO`

NewIgmpConfigVO instantiates a new IgmpConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIgmpConfigVOWithDefaults

`func NewIgmpConfigVOWithDefaults() *IgmpConfigVO`

NewIgmpConfigVOWithDefaults instantiates a new IgmpConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *IgmpConfigVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *IgmpConfigVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *IgmpConfigVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *IgmpConfigVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *IgmpConfigVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *IgmpConfigVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *IgmpConfigVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *IgmpConfigVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *IgmpConfigVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *IgmpConfigVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *IgmpConfigVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *IgmpConfigVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *IgmpConfigVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *IgmpConfigVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *IgmpConfigVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *IgmpConfigVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetGeneralQuerySourceIp

`func (o *IgmpConfigVO) GetGeneralQuerySourceIp() string`

GetGeneralQuerySourceIp returns the GeneralQuerySourceIp field if non-nil, zero value otherwise.

### GetGeneralQuerySourceIpOk

`func (o *IgmpConfigVO) GetGeneralQuerySourceIpOk() (*string, bool)`

GetGeneralQuerySourceIpOk returns a tuple with the GeneralQuerySourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralQuerySourceIp

`func (o *IgmpConfigVO) SetGeneralQuerySourceIp(v string)`

SetGeneralQuerySourceIp sets GeneralQuerySourceIp field to given value.

### HasGeneralQuerySourceIp

`func (o *IgmpConfigVO) HasGeneralQuerySourceIp() bool`

HasGeneralQuerySourceIp returns a boolean if a field has been set.

### GetLastMemberQueryCount

`func (o *IgmpConfigVO) GetLastMemberQueryCount() int32`

GetLastMemberQueryCount returns the LastMemberQueryCount field if non-nil, zero value otherwise.

### GetLastMemberQueryCountOk

`func (o *IgmpConfigVO) GetLastMemberQueryCountOk() (*int32, bool)`

GetLastMemberQueryCountOk returns a tuple with the LastMemberQueryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMemberQueryCount

`func (o *IgmpConfigVO) SetLastMemberQueryCount(v int32)`

SetLastMemberQueryCount sets LastMemberQueryCount field to given value.

### HasLastMemberQueryCount

`func (o *IgmpConfigVO) HasLastMemberQueryCount() bool`

HasLastMemberQueryCount returns a boolean if a field has been set.

### GetLastMemberQueryInterval

`func (o *IgmpConfigVO) GetLastMemberQueryInterval() int32`

GetLastMemberQueryInterval returns the LastMemberQueryInterval field if non-nil, zero value otherwise.

### GetLastMemberQueryIntervalOk

`func (o *IgmpConfigVO) GetLastMemberQueryIntervalOk() (*int32, bool)`

GetLastMemberQueryIntervalOk returns a tuple with the LastMemberQueryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMemberQueryInterval

`func (o *IgmpConfigVO) SetLastMemberQueryInterval(v int32)`

SetLastMemberQueryInterval sets LastMemberQueryInterval field to given value.

### HasLastMemberQueryInterval

`func (o *IgmpConfigVO) HasLastMemberQueryInterval() bool`

HasLastMemberQueryInterval returns a boolean if a field has been set.

### GetMac

`func (o *IgmpConfigVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *IgmpConfigVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *IgmpConfigVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *IgmpConfigVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMaximumResponseTime

`func (o *IgmpConfigVO) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *IgmpConfigVO) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *IgmpConfigVO) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *IgmpConfigVO) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### GetNetworkId

`func (o *IgmpConfigVO) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IgmpConfigVO) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IgmpConfigVO) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IgmpConfigVO) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *IgmpConfigVO) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *IgmpConfigVO) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *IgmpConfigVO) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *IgmpConfigVO) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetOswStack

`func (o *IgmpConfigVO) GetOswStack() OswStackInfoVO`

GetOswStack returns the OswStack field if non-nil, zero value otherwise.

### GetOswStackOk

`func (o *IgmpConfigVO) GetOswStackOk() (*OswStackInfoVO, bool)`

GetOswStackOk returns a tuple with the OswStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStack

`func (o *IgmpConfigVO) SetOswStack(v OswStackInfoVO)`

SetOswStack sets OswStack field to given value.

### HasOswStack

`func (o *IgmpConfigVO) HasOswStack() bool`

HasOswStack returns a boolean if a field has been set.

### GetQueryInterval

`func (o *IgmpConfigVO) GetQueryInterval() int32`

GetQueryInterval returns the QueryInterval field if non-nil, zero value otherwise.

### GetQueryIntervalOk

`func (o *IgmpConfigVO) GetQueryIntervalOk() (*int32, bool)`

GetQueryIntervalOk returns a tuple with the QueryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryInterval

`func (o *IgmpConfigVO) SetQueryInterval(v int32)`

SetQueryInterval sets QueryInterval field to given value.

### HasQueryInterval

`func (o *IgmpConfigVO) HasQueryInterval() bool`

HasQueryInterval returns a boolean if a field has been set.

### GetStackId

`func (o *IgmpConfigVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *IgmpConfigVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *IgmpConfigVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *IgmpConfigVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *IgmpConfigVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *IgmpConfigVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *IgmpConfigVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *IgmpConfigVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetVlan

`func (o *IgmpConfigVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *IgmpConfigVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *IgmpConfigVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *IgmpConfigVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


