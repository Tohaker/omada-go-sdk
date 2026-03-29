# OswQosRuleVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindType** | Pointer to **int32** |  | [optional] 
**DPort** | Pointer to **int32** |  | [optional] 
**DeviceList** | Pointer to [**[]OswQosRuleDeviceVO**](OswQosRuleDeviceVO.md) |  | [optional] 
**Dscp** | Pointer to **int32** |  | [optional] 
**DscpEnable** | Pointer to **bool** |  | [optional] 
**DscpRe** | Pointer to **int32** |  | [optional] 
**DscpReEnable** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IpVersion** | **[]int32** |  | 
**LanNetworkEntries** | Pointer to [**[]LanNetworkEntryVO**](LanNetworkEntryVO.md) |  | [optional] 
**MacLagIdsMap** | Pointer to **map[string][]int32** |  | [optional] 
**MacPortIdsMap** | Pointer to **map[string][]int32** |  | [optional] 
**MacStdPortIdsMap** | Pointer to **map[string][]string** |  | [optional] 
**Name** | **string** |  | 
**Protocol** | Pointer to **int32** |  | [optional] 
**Queue** | **int32** |  | 
**SPort** | Pointer to **int32** |  | [optional] 
**Status** | **bool** |  | 
**Type** | **int32** |  | 

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

### GetDscpEnable

`func (o *OswQosRuleVO) GetDscpEnable() bool`

GetDscpEnable returns the DscpEnable field if non-nil, zero value otherwise.

### GetDscpEnableOk

`func (o *OswQosRuleVO) GetDscpEnableOk() (*bool, bool)`

GetDscpEnableOk returns a tuple with the DscpEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscpEnable

`func (o *OswQosRuleVO) SetDscpEnable(v bool)`

SetDscpEnable sets DscpEnable field to given value.

### HasDscpEnable

`func (o *OswQosRuleVO) HasDscpEnable() bool`

HasDscpEnable returns a boolean if a field has been set.

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

### GetMacLagIdsMap

`func (o *OswQosRuleVO) GetMacLagIdsMap() map[string][]int32`

GetMacLagIdsMap returns the MacLagIdsMap field if non-nil, zero value otherwise.

### GetMacLagIdsMapOk

`func (o *OswQosRuleVO) GetMacLagIdsMapOk() (*map[string][]int32, bool)`

GetMacLagIdsMapOk returns a tuple with the MacLagIdsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacLagIdsMap

`func (o *OswQosRuleVO) SetMacLagIdsMap(v map[string][]int32)`

SetMacLagIdsMap sets MacLagIdsMap field to given value.

### HasMacLagIdsMap

`func (o *OswQosRuleVO) HasMacLagIdsMap() bool`

HasMacLagIdsMap returns a boolean if a field has been set.

### GetMacPortIdsMap

`func (o *OswQosRuleVO) GetMacPortIdsMap() map[string][]int32`

GetMacPortIdsMap returns the MacPortIdsMap field if non-nil, zero value otherwise.

### GetMacPortIdsMapOk

`func (o *OswQosRuleVO) GetMacPortIdsMapOk() (*map[string][]int32, bool)`

GetMacPortIdsMapOk returns a tuple with the MacPortIdsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacPortIdsMap

`func (o *OswQosRuleVO) SetMacPortIdsMap(v map[string][]int32)`

SetMacPortIdsMap sets MacPortIdsMap field to given value.

### HasMacPortIdsMap

`func (o *OswQosRuleVO) HasMacPortIdsMap() bool`

HasMacPortIdsMap returns a boolean if a field has been set.

### GetMacStdPortIdsMap

`func (o *OswQosRuleVO) GetMacStdPortIdsMap() map[string][]string`

GetMacStdPortIdsMap returns the MacStdPortIdsMap field if non-nil, zero value otherwise.

### GetMacStdPortIdsMapOk

`func (o *OswQosRuleVO) GetMacStdPortIdsMapOk() (*map[string][]string, bool)`

GetMacStdPortIdsMapOk returns a tuple with the MacStdPortIdsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacStdPortIdsMap

`func (o *OswQosRuleVO) SetMacStdPortIdsMap(v map[string][]string)`

SetMacStdPortIdsMap sets MacStdPortIdsMap field to given value.

### HasMacStdPortIdsMap

`func (o *OswQosRuleVO) HasMacStdPortIdsMap() bool`

HasMacStdPortIdsMap returns a boolean if a field has been set.

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


