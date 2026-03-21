# MldConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceModel** | Pointer to **string** | Model of device,for example:EAP225 | [optional] 
**DeviceModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**DeviceName** | Pointer to **string** | The name of the device. | [optional] 
**DeviceType** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 
**GeneralQuerySourceIp** | Pointer to **string** | The source IP of the general query message sent. | [optional] 
**LastMemberQueryCount** | Pointer to **int32** | The number of specific group query messages sent by the device. | [optional] 
**LastMemberQueryInterval** | Pointer to **int32** | The time interval for the device to send a specific group of query messages (unit: s). | [optional] 
**Mac** | Pointer to **string** | The unique identification of the device. | [optional] 
**MaximumResponseTime** | Pointer to **int32** | The maximum response time of the host to the general query message (unit: s). | [optional] 
**NetworkId** | Pointer to **string** | The primary id of the network. | [optional] 
**NetworkName** | Pointer to **string** | The name of the network.Name should contain 1 to 128 characters. | [optional] 
**OswStack** | Pointer to [**OswStackInfoVO**](OswStackInfoVO.md) |  | [optional] 
**QueryInterval** | Pointer to **int32** | The time interval for the MLD query to send query messages (unit: s). | [optional] 
**StackId** | Pointer to **string** | The id of the stacking devices. | [optional] 
**StackName** | Pointer to **string** | The name of the stacking devices. | [optional] 
**Vlan** | Pointer to **int32** | The vlan of the network.The vlan should be within the range of 1 to 4090. | [optional] 

## Methods

### NewMldConfigVO

`func NewMldConfigVO() *MldConfigVO`

NewMldConfigVO instantiates a new MldConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMldConfigVOWithDefaults

`func NewMldConfigVOWithDefaults() *MldConfigVO`

NewMldConfigVOWithDefaults instantiates a new MldConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceModel

`func (o *MldConfigVO) GetDeviceModel() string`

GetDeviceModel returns the DeviceModel field if non-nil, zero value otherwise.

### GetDeviceModelOk

`func (o *MldConfigVO) GetDeviceModelOk() (*string, bool)`

GetDeviceModelOk returns a tuple with the DeviceModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModel

`func (o *MldConfigVO) SetDeviceModel(v string)`

SetDeviceModel sets DeviceModel field to given value.

### HasDeviceModel

`func (o *MldConfigVO) HasDeviceModel() bool`

HasDeviceModel returns a boolean if a field has been set.

### GetDeviceModelVersion

`func (o *MldConfigVO) GetDeviceModelVersion() string`

GetDeviceModelVersion returns the DeviceModelVersion field if non-nil, zero value otherwise.

### GetDeviceModelVersionOk

`func (o *MldConfigVO) GetDeviceModelVersionOk() (*string, bool)`

GetDeviceModelVersionOk returns a tuple with the DeviceModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceModelVersion

`func (o *MldConfigVO) SetDeviceModelVersion(v string)`

SetDeviceModelVersion sets DeviceModelVersion field to given value.

### HasDeviceModelVersion

`func (o *MldConfigVO) HasDeviceModelVersion() bool`

HasDeviceModelVersion returns a boolean if a field has been set.

### GetDeviceName

`func (o *MldConfigVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *MldConfigVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *MldConfigVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *MldConfigVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceType

`func (o *MldConfigVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MldConfigVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MldConfigVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *MldConfigVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetGeneralQuerySourceIp

`func (o *MldConfigVO) GetGeneralQuerySourceIp() string`

GetGeneralQuerySourceIp returns the GeneralQuerySourceIp field if non-nil, zero value otherwise.

### GetGeneralQuerySourceIpOk

`func (o *MldConfigVO) GetGeneralQuerySourceIpOk() (*string, bool)`

GetGeneralQuerySourceIpOk returns a tuple with the GeneralQuerySourceIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneralQuerySourceIp

`func (o *MldConfigVO) SetGeneralQuerySourceIp(v string)`

SetGeneralQuerySourceIp sets GeneralQuerySourceIp field to given value.

### HasGeneralQuerySourceIp

`func (o *MldConfigVO) HasGeneralQuerySourceIp() bool`

HasGeneralQuerySourceIp returns a boolean if a field has been set.

### GetLastMemberQueryCount

`func (o *MldConfigVO) GetLastMemberQueryCount() int32`

GetLastMemberQueryCount returns the LastMemberQueryCount field if non-nil, zero value otherwise.

### GetLastMemberQueryCountOk

`func (o *MldConfigVO) GetLastMemberQueryCountOk() (*int32, bool)`

GetLastMemberQueryCountOk returns a tuple with the LastMemberQueryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMemberQueryCount

`func (o *MldConfigVO) SetLastMemberQueryCount(v int32)`

SetLastMemberQueryCount sets LastMemberQueryCount field to given value.

### HasLastMemberQueryCount

`func (o *MldConfigVO) HasLastMemberQueryCount() bool`

HasLastMemberQueryCount returns a boolean if a field has been set.

### GetLastMemberQueryInterval

`func (o *MldConfigVO) GetLastMemberQueryInterval() int32`

GetLastMemberQueryInterval returns the LastMemberQueryInterval field if non-nil, zero value otherwise.

### GetLastMemberQueryIntervalOk

`func (o *MldConfigVO) GetLastMemberQueryIntervalOk() (*int32, bool)`

GetLastMemberQueryIntervalOk returns a tuple with the LastMemberQueryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMemberQueryInterval

`func (o *MldConfigVO) SetLastMemberQueryInterval(v int32)`

SetLastMemberQueryInterval sets LastMemberQueryInterval field to given value.

### HasLastMemberQueryInterval

`func (o *MldConfigVO) HasLastMemberQueryInterval() bool`

HasLastMemberQueryInterval returns a boolean if a field has been set.

### GetMac

`func (o *MldConfigVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MldConfigVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MldConfigVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MldConfigVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMaximumResponseTime

`func (o *MldConfigVO) GetMaximumResponseTime() int32`

GetMaximumResponseTime returns the MaximumResponseTime field if non-nil, zero value otherwise.

### GetMaximumResponseTimeOk

`func (o *MldConfigVO) GetMaximumResponseTimeOk() (*int32, bool)`

GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumResponseTime

`func (o *MldConfigVO) SetMaximumResponseTime(v int32)`

SetMaximumResponseTime sets MaximumResponseTime field to given value.

### HasMaximumResponseTime

`func (o *MldConfigVO) HasMaximumResponseTime() bool`

HasMaximumResponseTime returns a boolean if a field has been set.

### GetNetworkId

`func (o *MldConfigVO) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *MldConfigVO) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *MldConfigVO) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *MldConfigVO) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *MldConfigVO) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *MldConfigVO) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *MldConfigVO) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *MldConfigVO) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetOswStack

`func (o *MldConfigVO) GetOswStack() OswStackInfoVO`

GetOswStack returns the OswStack field if non-nil, zero value otherwise.

### GetOswStackOk

`func (o *MldConfigVO) GetOswStackOk() (*OswStackInfoVO, bool)`

GetOswStackOk returns a tuple with the OswStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStack

`func (o *MldConfigVO) SetOswStack(v OswStackInfoVO)`

SetOswStack sets OswStack field to given value.

### HasOswStack

`func (o *MldConfigVO) HasOswStack() bool`

HasOswStack returns a boolean if a field has been set.

### GetQueryInterval

`func (o *MldConfigVO) GetQueryInterval() int32`

GetQueryInterval returns the QueryInterval field if non-nil, zero value otherwise.

### GetQueryIntervalOk

`func (o *MldConfigVO) GetQueryIntervalOk() (*int32, bool)`

GetQueryIntervalOk returns a tuple with the QueryInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryInterval

`func (o *MldConfigVO) SetQueryInterval(v int32)`

SetQueryInterval sets QueryInterval field to given value.

### HasQueryInterval

`func (o *MldConfigVO) HasQueryInterval() bool`

HasQueryInterval returns a boolean if a field has been set.

### GetStackId

`func (o *MldConfigVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *MldConfigVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *MldConfigVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *MldConfigVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *MldConfigVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *MldConfigVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *MldConfigVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *MldConfigVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetVlan

`func (o *MldConfigVO) GetVlan() int32`

GetVlan returns the Vlan field if non-nil, zero value otherwise.

### GetVlanOk

`func (o *MldConfigVO) GetVlanOk() (*int32, bool)`

GetVlanOk returns a tuple with the Vlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlan

`func (o *MldConfigVO) SetVlan(v int32)`

SetVlan sets Vlan field to given value.

### HasVlan

`func (o *MldConfigVO) HasVlan() bool`

HasVlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


