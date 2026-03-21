# DhcpSettingsVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DhcpNextServer** | Pointer to **string** |  | [optional] 
**Dhcpns** | Pointer to **string** |  | [optional] 
**Enable** | Pointer to **bool** |  | [optional] 
**Gateway** | Pointer to **string** |  | [optional] 
**HostIP** | Pointer to **string** |  | [optional] 
**IpRangeEnd** | Pointer to **int64** |  | [optional] 
**IpRangePool** | Pointer to [**[]DhcpRangeVO**](DhcpRangeVO.md) |  | [optional] 
**IpRangeStart** | Pointer to **int64** |  | [optional] 
**IpaddrEnd** | Pointer to **string** |  | [optional] 
**IpaddrStart** | Pointer to **string** |  | [optional] 
**Leasetime** | Pointer to **int32** |  | [optional] 
**Option138** | Pointer to **string** |  | [optional] 
**Option60** | Pointer to **string** |  | [optional] 
**Option66** | Pointer to **string** |  | [optional] 
**Options** | Pointer to [**[]DhcpOptionVO**](DhcpOptionVO.md) |  | [optional] 
**PriDns** | Pointer to **string** |  | [optional] 
**SndDns** | Pointer to **string** |  | [optional] 

## Methods

### NewDhcpSettingsVO

`func NewDhcpSettingsVO() *DhcpSettingsVO`

NewDhcpSettingsVO instantiates a new DhcpSettingsVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSettingsVOWithDefaults

`func NewDhcpSettingsVOWithDefaults() *DhcpSettingsVO`

NewDhcpSettingsVOWithDefaults instantiates a new DhcpSettingsVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDhcpNextServer

`func (o *DhcpSettingsVO) GetDhcpNextServer() string`

GetDhcpNextServer returns the DhcpNextServer field if non-nil, zero value otherwise.

### GetDhcpNextServerOk

`func (o *DhcpSettingsVO) GetDhcpNextServerOk() (*string, bool)`

GetDhcpNextServerOk returns a tuple with the DhcpNextServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpNextServer

`func (o *DhcpSettingsVO) SetDhcpNextServer(v string)`

SetDhcpNextServer sets DhcpNextServer field to given value.

### HasDhcpNextServer

`func (o *DhcpSettingsVO) HasDhcpNextServer() bool`

HasDhcpNextServer returns a boolean if a field has been set.

### GetDhcpns

`func (o *DhcpSettingsVO) GetDhcpns() string`

GetDhcpns returns the Dhcpns field if non-nil, zero value otherwise.

### GetDhcpnsOk

`func (o *DhcpSettingsVO) GetDhcpnsOk() (*string, bool)`

GetDhcpnsOk returns a tuple with the Dhcpns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpns

`func (o *DhcpSettingsVO) SetDhcpns(v string)`

SetDhcpns sets Dhcpns field to given value.

### HasDhcpns

`func (o *DhcpSettingsVO) HasDhcpns() bool`

HasDhcpns returns a boolean if a field has been set.

### GetEnable

`func (o *DhcpSettingsVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpSettingsVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpSettingsVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DhcpSettingsVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetGateway

`func (o *DhcpSettingsVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *DhcpSettingsVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *DhcpSettingsVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *DhcpSettingsVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetHostIP

`func (o *DhcpSettingsVO) GetHostIP() string`

GetHostIP returns the HostIP field if non-nil, zero value otherwise.

### GetHostIPOk

`func (o *DhcpSettingsVO) GetHostIPOk() (*string, bool)`

GetHostIPOk returns a tuple with the HostIP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostIP

`func (o *DhcpSettingsVO) SetHostIP(v string)`

SetHostIP sets HostIP field to given value.

### HasHostIP

`func (o *DhcpSettingsVO) HasHostIP() bool`

HasHostIP returns a boolean if a field has been set.

### GetIpRangeEnd

`func (o *DhcpSettingsVO) GetIpRangeEnd() int64`

GetIpRangeEnd returns the IpRangeEnd field if non-nil, zero value otherwise.

### GetIpRangeEndOk

`func (o *DhcpSettingsVO) GetIpRangeEndOk() (*int64, bool)`

GetIpRangeEndOk returns a tuple with the IpRangeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeEnd

`func (o *DhcpSettingsVO) SetIpRangeEnd(v int64)`

SetIpRangeEnd sets IpRangeEnd field to given value.

### HasIpRangeEnd

`func (o *DhcpSettingsVO) HasIpRangeEnd() bool`

HasIpRangeEnd returns a boolean if a field has been set.

### GetIpRangePool

`func (o *DhcpSettingsVO) GetIpRangePool() []DhcpRangeVO`

GetIpRangePool returns the IpRangePool field if non-nil, zero value otherwise.

### GetIpRangePoolOk

`func (o *DhcpSettingsVO) GetIpRangePoolOk() (*[]DhcpRangeVO, bool)`

GetIpRangePoolOk returns a tuple with the IpRangePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangePool

`func (o *DhcpSettingsVO) SetIpRangePool(v []DhcpRangeVO)`

SetIpRangePool sets IpRangePool field to given value.

### HasIpRangePool

`func (o *DhcpSettingsVO) HasIpRangePool() bool`

HasIpRangePool returns a boolean if a field has been set.

### GetIpRangeStart

`func (o *DhcpSettingsVO) GetIpRangeStart() int64`

GetIpRangeStart returns the IpRangeStart field if non-nil, zero value otherwise.

### GetIpRangeStartOk

`func (o *DhcpSettingsVO) GetIpRangeStartOk() (*int64, bool)`

GetIpRangeStartOk returns a tuple with the IpRangeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRangeStart

`func (o *DhcpSettingsVO) SetIpRangeStart(v int64)`

SetIpRangeStart sets IpRangeStart field to given value.

### HasIpRangeStart

`func (o *DhcpSettingsVO) HasIpRangeStart() bool`

HasIpRangeStart returns a boolean if a field has been set.

### GetIpaddrEnd

`func (o *DhcpSettingsVO) GetIpaddrEnd() string`

GetIpaddrEnd returns the IpaddrEnd field if non-nil, zero value otherwise.

### GetIpaddrEndOk

`func (o *DhcpSettingsVO) GetIpaddrEndOk() (*string, bool)`

GetIpaddrEndOk returns a tuple with the IpaddrEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrEnd

`func (o *DhcpSettingsVO) SetIpaddrEnd(v string)`

SetIpaddrEnd sets IpaddrEnd field to given value.

### HasIpaddrEnd

`func (o *DhcpSettingsVO) HasIpaddrEnd() bool`

HasIpaddrEnd returns a boolean if a field has been set.

### GetIpaddrStart

`func (o *DhcpSettingsVO) GetIpaddrStart() string`

GetIpaddrStart returns the IpaddrStart field if non-nil, zero value otherwise.

### GetIpaddrStartOk

`func (o *DhcpSettingsVO) GetIpaddrStartOk() (*string, bool)`

GetIpaddrStartOk returns a tuple with the IpaddrStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpaddrStart

`func (o *DhcpSettingsVO) SetIpaddrStart(v string)`

SetIpaddrStart sets IpaddrStart field to given value.

### HasIpaddrStart

`func (o *DhcpSettingsVO) HasIpaddrStart() bool`

HasIpaddrStart returns a boolean if a field has been set.

### GetLeasetime

`func (o *DhcpSettingsVO) GetLeasetime() int32`

GetLeasetime returns the Leasetime field if non-nil, zero value otherwise.

### GetLeasetimeOk

`func (o *DhcpSettingsVO) GetLeasetimeOk() (*int32, bool)`

GetLeasetimeOk returns a tuple with the Leasetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasetime

`func (o *DhcpSettingsVO) SetLeasetime(v int32)`

SetLeasetime sets Leasetime field to given value.

### HasLeasetime

`func (o *DhcpSettingsVO) HasLeasetime() bool`

HasLeasetime returns a boolean if a field has been set.

### GetOption138

`func (o *DhcpSettingsVO) GetOption138() string`

GetOption138 returns the Option138 field if non-nil, zero value otherwise.

### GetOption138Ok

`func (o *DhcpSettingsVO) GetOption138Ok() (*string, bool)`

GetOption138Ok returns a tuple with the Option138 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption138

`func (o *DhcpSettingsVO) SetOption138(v string)`

SetOption138 sets Option138 field to given value.

### HasOption138

`func (o *DhcpSettingsVO) HasOption138() bool`

HasOption138 returns a boolean if a field has been set.

### GetOption60

`func (o *DhcpSettingsVO) GetOption60() string`

GetOption60 returns the Option60 field if non-nil, zero value otherwise.

### GetOption60Ok

`func (o *DhcpSettingsVO) GetOption60Ok() (*string, bool)`

GetOption60Ok returns a tuple with the Option60 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption60

`func (o *DhcpSettingsVO) SetOption60(v string)`

SetOption60 sets Option60 field to given value.

### HasOption60

`func (o *DhcpSettingsVO) HasOption60() bool`

HasOption60 returns a boolean if a field has been set.

### GetOption66

`func (o *DhcpSettingsVO) GetOption66() string`

GetOption66 returns the Option66 field if non-nil, zero value otherwise.

### GetOption66Ok

`func (o *DhcpSettingsVO) GetOption66Ok() (*string, bool)`

GetOption66Ok returns a tuple with the Option66 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption66

`func (o *DhcpSettingsVO) SetOption66(v string)`

SetOption66 sets Option66 field to given value.

### HasOption66

`func (o *DhcpSettingsVO) HasOption66() bool`

HasOption66 returns a boolean if a field has been set.

### GetOptions

`func (o *DhcpSettingsVO) GetOptions() []DhcpOptionVO`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *DhcpSettingsVO) GetOptionsOk() (*[]DhcpOptionVO, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *DhcpSettingsVO) SetOptions(v []DhcpOptionVO)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *DhcpSettingsVO) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPriDns

`func (o *DhcpSettingsVO) GetPriDns() string`

GetPriDns returns the PriDns field if non-nil, zero value otherwise.

### GetPriDnsOk

`func (o *DhcpSettingsVO) GetPriDnsOk() (*string, bool)`

GetPriDnsOk returns a tuple with the PriDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriDns

`func (o *DhcpSettingsVO) SetPriDns(v string)`

SetPriDns sets PriDns field to given value.

### HasPriDns

`func (o *DhcpSettingsVO) HasPriDns() bool`

HasPriDns returns a boolean if a field has been set.

### GetSndDns

`func (o *DhcpSettingsVO) GetSndDns() string`

GetSndDns returns the SndDns field if non-nil, zero value otherwise.

### GetSndDnsOk

`func (o *DhcpSettingsVO) GetSndDnsOk() (*string, bool)`

GetSndDnsOk returns a tuple with the SndDns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSndDns

`func (o *DhcpSettingsVO) SetSndDns(v string)`

SetSndDns sets SndDns field to given value.

### HasSndDns

`func (o *DhcpSettingsVO) HasSndDns() bool`

HasSndDns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


