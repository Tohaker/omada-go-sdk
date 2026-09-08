# ClientAccessTimeInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationTime** | Pointer to **int32** | Association time | [optional] 
**AuthTime** | Pointer to **int32** | Authorization time | [optional] 
**ClientModel** | Pointer to **string** | Client model for icon display | [optional] 
**ClientName** | Pointer to **string** | Client name | [optional] 
**ClientType** | Pointer to **string** | Client type for icon display (e.g. Mobile, Laptop, IPC) | [optional] 
**DeviceType** | Pointer to **string** | Client-connected device type | [optional] 
**DhcpTime** | Pointer to **int32** | DHCP time | [optional] 
**DnsTime** | Pointer to **int32** | DNS time | [optional] 
**Incidents** | Pointer to **int32** | Incident counts | [optional] 
**Ip** | Pointer to **string** | ip | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**TotalTime** | Pointer to **int32** | Total access time, only including associationTime now | [optional] 

## Methods

### NewClientAccessTimeInfoVO

`func NewClientAccessTimeInfoVO() *ClientAccessTimeInfoVO`

NewClientAccessTimeInfoVO instantiates a new ClientAccessTimeInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientAccessTimeInfoVOWithDefaults

`func NewClientAccessTimeInfoVOWithDefaults() *ClientAccessTimeInfoVO`

NewClientAccessTimeInfoVOWithDefaults instantiates a new ClientAccessTimeInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationTime

`func (o *ClientAccessTimeInfoVO) GetAssociationTime() int32`

GetAssociationTime returns the AssociationTime field if non-nil, zero value otherwise.

### GetAssociationTimeOk

`func (o *ClientAccessTimeInfoVO) GetAssociationTimeOk() (*int32, bool)`

GetAssociationTimeOk returns a tuple with the AssociationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTime

`func (o *ClientAccessTimeInfoVO) SetAssociationTime(v int32)`

SetAssociationTime sets AssociationTime field to given value.

### HasAssociationTime

`func (o *ClientAccessTimeInfoVO) HasAssociationTime() bool`

HasAssociationTime returns a boolean if a field has been set.

### GetAuthTime

`func (o *ClientAccessTimeInfoVO) GetAuthTime() int32`

GetAuthTime returns the AuthTime field if non-nil, zero value otherwise.

### GetAuthTimeOk

`func (o *ClientAccessTimeInfoVO) GetAuthTimeOk() (*int32, bool)`

GetAuthTimeOk returns a tuple with the AuthTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthTime

`func (o *ClientAccessTimeInfoVO) SetAuthTime(v int32)`

SetAuthTime sets AuthTime field to given value.

### HasAuthTime

`func (o *ClientAccessTimeInfoVO) HasAuthTime() bool`

HasAuthTime returns a boolean if a field has been set.

### GetClientModel

`func (o *ClientAccessTimeInfoVO) GetClientModel() string`

GetClientModel returns the ClientModel field if non-nil, zero value otherwise.

### GetClientModelOk

`func (o *ClientAccessTimeInfoVO) GetClientModelOk() (*string, bool)`

GetClientModelOk returns a tuple with the ClientModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientModel

`func (o *ClientAccessTimeInfoVO) SetClientModel(v string)`

SetClientModel sets ClientModel field to given value.

### HasClientModel

`func (o *ClientAccessTimeInfoVO) HasClientModel() bool`

HasClientModel returns a boolean if a field has been set.

### GetClientName

`func (o *ClientAccessTimeInfoVO) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientAccessTimeInfoVO) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientAccessTimeInfoVO) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientAccessTimeInfoVO) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### GetClientType

`func (o *ClientAccessTimeInfoVO) GetClientType() string`

GetClientType returns the ClientType field if non-nil, zero value otherwise.

### GetClientTypeOk

`func (o *ClientAccessTimeInfoVO) GetClientTypeOk() (*string, bool)`

GetClientTypeOk returns a tuple with the ClientType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientType

`func (o *ClientAccessTimeInfoVO) SetClientType(v string)`

SetClientType sets ClientType field to given value.

### HasClientType

`func (o *ClientAccessTimeInfoVO) HasClientType() bool`

HasClientType returns a boolean if a field has been set.

### GetDeviceType

`func (o *ClientAccessTimeInfoVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClientAccessTimeInfoVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClientAccessTimeInfoVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClientAccessTimeInfoVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDhcpTime

`func (o *ClientAccessTimeInfoVO) GetDhcpTime() int32`

GetDhcpTime returns the DhcpTime field if non-nil, zero value otherwise.

### GetDhcpTimeOk

`func (o *ClientAccessTimeInfoVO) GetDhcpTimeOk() (*int32, bool)`

GetDhcpTimeOk returns a tuple with the DhcpTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpTime

`func (o *ClientAccessTimeInfoVO) SetDhcpTime(v int32)`

SetDhcpTime sets DhcpTime field to given value.

### HasDhcpTime

`func (o *ClientAccessTimeInfoVO) HasDhcpTime() bool`

HasDhcpTime returns a boolean if a field has been set.

### GetDnsTime

`func (o *ClientAccessTimeInfoVO) GetDnsTime() int32`

GetDnsTime returns the DnsTime field if non-nil, zero value otherwise.

### GetDnsTimeOk

`func (o *ClientAccessTimeInfoVO) GetDnsTimeOk() (*int32, bool)`

GetDnsTimeOk returns a tuple with the DnsTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnsTime

`func (o *ClientAccessTimeInfoVO) SetDnsTime(v int32)`

SetDnsTime sets DnsTime field to given value.

### HasDnsTime

`func (o *ClientAccessTimeInfoVO) HasDnsTime() bool`

HasDnsTime returns a boolean if a field has been set.

### GetIncidents

`func (o *ClientAccessTimeInfoVO) GetIncidents() int32`

GetIncidents returns the Incidents field if non-nil, zero value otherwise.

### GetIncidentsOk

`func (o *ClientAccessTimeInfoVO) GetIncidentsOk() (*int32, bool)`

GetIncidentsOk returns a tuple with the Incidents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncidents

`func (o *ClientAccessTimeInfoVO) SetIncidents(v int32)`

SetIncidents sets Incidents field to given value.

### HasIncidents

`func (o *ClientAccessTimeInfoVO) HasIncidents() bool`

HasIncidents returns a boolean if a field has been set.

### GetIp

`func (o *ClientAccessTimeInfoVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientAccessTimeInfoVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientAccessTimeInfoVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientAccessTimeInfoVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ClientAccessTimeInfoVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientAccessTimeInfoVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientAccessTimeInfoVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientAccessTimeInfoVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetTotalTime

`func (o *ClientAccessTimeInfoVO) GetTotalTime() int32`

GetTotalTime returns the TotalTime field if non-nil, zero value otherwise.

### GetTotalTimeOk

`func (o *ClientAccessTimeInfoVO) GetTotalTimeOk() (*int32, bool)`

GetTotalTimeOk returns a tuple with the TotalTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTime

`func (o *ClientAccessTimeInfoVO) SetTotalTime(v int32)`

SetTotalTime sets TotalTime field to given value.

### HasTotalTime

`func (o *ClientAccessTimeInfoVO) HasTotalTime() bool`

HasTotalTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


