# OswQosRuleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindType** | Pointer to **int32** | Switch port bind type(0: all switch port, 1: custom switch port). | [optional] 
**DPort** | Pointer to **int32** | The destination port(0-65535). When the type is \&quot;Custom\&quot; and the protocol is TCP or UDP, it can be issued. | [optional] 
**DeviceList** | Pointer to [**[]OswQosRuleDeviceVO**](OswQosRuleDeviceVO.md) | List of switch devices to which QoS rule are bound, only for bindType 1 | [optional] 
**Dscp** | Pointer to **int32** | The dscp value(0-63). When the type is \&quot;Custom\&quot;, it can be issued. | [optional] 
**DscpRe** | Pointer to **int32** | The remarked dscp value(0-63). If it is set to \&quot;Auto\&quot;, the actual issued value is 64. | [optional] 
**DscpReEnable** | Pointer to **bool** | Whether to enable DSCP remark | [optional] 
**Id** | Pointer to **string** | The Qos rule id. This parameter is not required when creating or modifying Qos rule. | [optional] 
**Index** | Pointer to **int32** | The index of Qos rule. This parameter is not required when creating or modifying Qos rule. | [optional] 
**IpVersion** | **[]int32** | The selected ipVersion list of Qos rule, IPv4:[0], IPv6:[1], IPv4&amp;IPv6:[0,1] | 
**LanNetworkEntries** | Pointer to [**[]LanNetworkEntryVO**](LanNetworkEntryVO.md) | It can be issued when the type is either \&quot;Network\&quot; or \&quot;Custom\&quot;. For the \&quot;Network\&quot; type, multiple options are available; for the \&quot;Custom\&quot; type, a single option is required. | [optional] 
**Name** | **string** | The name of Qos rule. | 
**Protocol** | Pointer to **int32** | Network Protocol(Reference to chapter 5.5.1 ACL Protocol Template at Home page). When the type is \&quot;Custom\&quot;, it can be issued. | [optional] 
**Queue** | **int32** | The queue of Qos rule(0-7). | 
**SPort** | Pointer to **int32** | The source port(0-65535). When the type is \&quot;Custom\&quot; and the protocol is TCP or UDP, it can be issued. | [optional] 
**Status** | **bool** | The status of Qos rule, true: enable, false:disable | 
**Type** | **int32** | The type of Qos rule, 0:Network, 1:Port, 2:Custom | 

## Methods

### NewOswQosRuleVO

`func NewOswQosRuleVO(ipVersion []int32, name string, queue int32, status bool, type_ int32, ) *OswQosRuleVO`

NewOswQosRuleVO instantiates a new OswQosRuleVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswQosRuleVOWithDefaults

`func NewOswQosRuleVOWithDefaults() *OswQosRuleVO`

NewOswQosRuleVOWithDefaults instantiates a new OswQosRuleVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindType

`func (o *OswQosRuleVO) GetBindType() int32`

GetBindType returns the BindType field if non-nil, zero value otherwise.

### GetBindTypeOk

`func (o *OswQosRuleVO) GetBindTypeOk() (*int32, bool)`

GetBindTypeOk returns a tuple with the BindType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindType

`func (o *OswQosRuleVO) SetBindType(v int32)`

SetBindType sets BindType field to given value.

### HasBindType

`func (o *OswQosRuleVO) HasBindType() bool`

HasBindType returns a boolean if a field has been set.

### GetDPort

`func (o *OswQosRuleVO) GetDPort() int32`

GetDPort returns the DPort field if non-nil, zero value otherwise.

### GetDPortOk

`func (o *OswQosRuleVO) GetDPortOk() (*int32, bool)`

GetDPortOk returns a tuple with the DPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDPort

`func (o *OswQosRuleVO) SetDPort(v int32)`

SetDPort sets DPort field to given value.

### HasDPort

`func (o *OswQosRuleVO) HasDPort() bool`

HasDPort returns a boolean if a field has been set.

### GetDeviceList

`func (o *OswQosRuleVO) GetDeviceList() []OswQosRuleDeviceVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *OswQosRuleVO) GetDeviceListOk() (*[]OswQosRuleDeviceVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *OswQosRuleVO) SetDeviceList(v []OswQosRuleDeviceVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *OswQosRuleVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetDscp

`func (o *OswQosRuleVO) GetDscp() int32`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *OswQosRuleVO) GetDscpOk() (*int32, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *OswQosRuleVO) SetDscp(v int32)`

SetDscp sets Dscp field to given value.

### HasDscp

`func (o *OswQosRuleVO) HasDscp() bool`

HasDscp returns a boolean if a field has been set.

### GetDscpRe

`func (o *OswQosRuleVO) GetDscpRe() int32`

GetDscpRe returns the DscpRe field if non-nil, zero value otherwise.

### GetDscpReOk

`func (o *OswQosRuleVO) GetDscpReOk() (*int32, bool)`

GetDscpReOk returns a tuple with the DscpRe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpRe

`func (o *OswQosRuleVO) SetDscpRe(v int32)`

SetDscpRe sets DscpRe field to given value.

### HasDscpRe

`func (o *OswQosRuleVO) HasDscpRe() bool`

HasDscpRe returns a boolean if a field has been set.

### GetDscpReEnable

`func (o *OswQosRuleVO) GetDscpReEnable() bool`

GetDscpReEnable returns the DscpReEnable field if non-nil, zero value otherwise.

### GetDscpReEnableOk

`func (o *OswQosRuleVO) GetDscpReEnableOk() (*bool, bool)`

GetDscpReEnableOk returns a tuple with the DscpReEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpReEnable

`func (o *OswQosRuleVO) SetDscpReEnable(v bool)`

SetDscpReEnable sets DscpReEnable field to given value.

### HasDscpReEnable

`func (o *OswQosRuleVO) HasDscpReEnable() bool`

HasDscpReEnable returns a boolean if a field has been set.

### GetId

`func (o *OswQosRuleVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswQosRuleVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswQosRuleVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswQosRuleVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIndex

`func (o *OswQosRuleVO) GetIndex() int32`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *OswQosRuleVO) GetIndexOk() (*int32, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *OswQosRuleVO) SetIndex(v int32)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *OswQosRuleVO) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetIpVersion

`func (o *OswQosRuleVO) GetIpVersion() []int32`

GetIpVersion returns the IpVersion field if non-nil, zero value otherwise.

### GetIpVersionOk

`func (o *OswQosRuleVO) GetIpVersionOk() (*[]int32, bool)`

GetIpVersionOk returns a tuple with the IpVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpVersion

`func (o *OswQosRuleVO) SetIpVersion(v []int32)`

SetIpVersion sets IpVersion field to given value.


### GetLanNetworkEntries

`func (o *OswQosRuleVO) GetLanNetworkEntries() []LanNetworkEntryVO`

GetLanNetworkEntries returns the LanNetworkEntries field if non-nil, zero value otherwise.

### GetLanNetworkEntriesOk

`func (o *OswQosRuleVO) GetLanNetworkEntriesOk() (*[]LanNetworkEntryVO, bool)`

GetLanNetworkEntriesOk returns a tuple with the LanNetworkEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanNetworkEntries

`func (o *OswQosRuleVO) SetLanNetworkEntries(v []LanNetworkEntryVO)`

SetLanNetworkEntries sets LanNetworkEntries field to given value.

### HasLanNetworkEntries

`func (o *OswQosRuleVO) HasLanNetworkEntries() bool`

HasLanNetworkEntries returns a boolean if a field has been set.

### GetName

`func (o *OswQosRuleVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswQosRuleVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswQosRuleVO) SetName(v string)`

SetName sets Name field to given value.


### GetProtocol

`func (o *OswQosRuleVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *OswQosRuleVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *OswQosRuleVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *OswQosRuleVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetQueue

`func (o *OswQosRuleVO) GetQueue() int32`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *OswQosRuleVO) GetQueueOk() (*int32, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *OswQosRuleVO) SetQueue(v int32)`

SetQueue sets Queue field to given value.


### GetSPort

`func (o *OswQosRuleVO) GetSPort() int32`

GetSPort returns the SPort field if non-nil, zero value otherwise.

### GetSPortOk

`func (o *OswQosRuleVO) GetSPortOk() (*int32, bool)`

GetSPortOk returns a tuple with the SPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSPort

`func (o *OswQosRuleVO) SetSPort(v int32)`

SetSPort sets SPort field to given value.

### HasSPort

`func (o *OswQosRuleVO) HasSPort() bool`

HasSPort returns a boolean if a field has been set.

### GetStatus

`func (o *OswQosRuleVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswQosRuleVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswQosRuleVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetType

`func (o *OswQosRuleVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswQosRuleVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswQosRuleVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


