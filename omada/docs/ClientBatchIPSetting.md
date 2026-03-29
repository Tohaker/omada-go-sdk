# ClientBatchIPSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IpList** | Pointer to [**[]ClientMACIPSetting**](ClientMACIPSetting.md) | List of ip and client mac. | [optional] 
**NetId** | Pointer to **string** | Network ID. | [optional] 
**ServerMac** | Pointer to **string** | DHCP Server Mac. | [optional] 
**ServerStackId** | Pointer to **string** | DHCP Server Stack ID. | [optional] 
**ServerType** | Pointer to **string** | DHCP Server Type. | [optional] 
**UseFixedAddr** | **bool** | Use fixed ip address. | 

## Methods

### NewClientBatchIPSetting

`func NewClientBatchIPSetting(useFixedAddr bool, ) *ClientBatchIPSetting`

NewClientBatchIPSetting instantiates a new ClientBatchIPSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientBatchIPSettingWithDefaults

`func NewClientBatchIPSettingWithDefaults() *ClientBatchIPSetting`

NewClientBatchIPSettingWithDefaults instantiates a new ClientBatchIPSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIpList

`func (o *ClientBatchIPSetting) GetIpList() []ClientMACIPSetting`

GetIpList returns the IpList field if non-nil, zero value otherwise.

### GetIpListOk

`func (o *ClientBatchIPSetting) GetIpListOk() (*[]ClientMACIPSetting, bool)`

GetIpListOk returns a tuple with the IpList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpList

`func (o *ClientBatchIPSetting) SetIpList(v []ClientMACIPSetting)`

SetIpList sets IpList field to given value.

### HasIpList

`func (o *ClientBatchIPSetting) HasIpList() bool`

HasIpList returns a boolean if a field has been set.

### GetNetId

`func (o *ClientBatchIPSetting) GetNetId() string`

GetNetId returns the NetId field if non-nil, zero value otherwise.

### GetNetIdOk

`func (o *ClientBatchIPSetting) GetNetIdOk() (*string, bool)`

GetNetIdOk returns a tuple with the NetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetId

`func (o *ClientBatchIPSetting) SetNetId(v string)`

SetNetId sets NetId field to given value.

### HasNetId

`func (o *ClientBatchIPSetting) HasNetId() bool`

HasNetId returns a boolean if a field has been set.

### GetServerMac

`func (o *ClientBatchIPSetting) GetServerMac() string`

GetServerMac returns the ServerMac field if non-nil, zero value otherwise.

### GetServerMacOk

`func (o *ClientBatchIPSetting) GetServerMacOk() (*string, bool)`

GetServerMacOk returns a tuple with the ServerMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMac

`func (o *ClientBatchIPSetting) SetServerMac(v string)`

SetServerMac sets ServerMac field to given value.

### HasServerMac

`func (o *ClientBatchIPSetting) HasServerMac() bool`

HasServerMac returns a boolean if a field has been set.

### GetServerStackId

`func (o *ClientBatchIPSetting) GetServerStackId() string`

GetServerStackId returns the ServerStackId field if non-nil, zero value otherwise.

### GetServerStackIdOk

`func (o *ClientBatchIPSetting) GetServerStackIdOk() (*string, bool)`

GetServerStackIdOk returns a tuple with the ServerStackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStackId

`func (o *ClientBatchIPSetting) SetServerStackId(v string)`

SetServerStackId sets ServerStackId field to given value.

### HasServerStackId

`func (o *ClientBatchIPSetting) HasServerStackId() bool`

HasServerStackId returns a boolean if a field has been set.

### GetServerType

`func (o *ClientBatchIPSetting) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ClientBatchIPSetting) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ClientBatchIPSetting) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ClientBatchIPSetting) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetUseFixedAddr

`func (o *ClientBatchIPSetting) GetUseFixedAddr() bool`

GetUseFixedAddr returns the UseFixedAddr field if non-nil, zero value otherwise.

### GetUseFixedAddrOk

`func (o *ClientBatchIPSetting) GetUseFixedAddrOk() (*bool, bool)`

GetUseFixedAddrOk returns a tuple with the UseFixedAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseFixedAddr

`func (o *ClientBatchIPSetting) SetUseFixedAddr(v bool)`

SetUseFixedAddr sets UseFixedAddr field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


