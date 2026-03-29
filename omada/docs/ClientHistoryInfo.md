# ClientHistoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationTime** | Pointer to **int32** | (Wireless) The time (ms) it takes for the client to connect to SSID. | [optional] 
**AuthInfo** | Pointer to [**[]AuthInfoOpenApiVO**](AuthInfoOpenApiVO.md) | Client portal authentication information | [optional] 
**DeviceName** | Pointer to **string** | Device name. | [optional] 
**Download** | Pointer to **int64** | Downstream traffic (Byte). | [optional] 
**Duration** | Pointer to **int64** | Up time (unit: s). | [optional] 
**FirstSeen** | Pointer to **int64** | The timestamp (ms) when the client connected. | [optional] 
**Guest** | Pointer to **bool** | (Wireless) Whether it is Guest (used to display the wireless Guest client icon). | [optional] 
**Id** | Pointer to **string** | Client History ID. | [optional] 
**Ip** | Pointer to **string** | IP Address. | [optional] 
**Ipv6List** | Pointer to **[]string** | IPv6 Address. | [optional] 
**LastSeen** | Pointer to **int64** | Last found time, timestamp (ms). | [optional] 
**Mac** | Pointer to **string** | Client MAC Address. | [optional] 
**Name** | Pointer to **string** | Client Name. | [optional] 
**Port** | Pointer to **int32** | (Wired) Port ID. | [optional] 
**Ssid** | Pointer to **string** | (Wireless)  SSID name. | [optional] 
**Upload** | Pointer to **int64** | Upstream traffic (Byte). | [optional] 

## Methods

### NewClientHistoryInfo

`func NewClientHistoryInfo() *ClientHistoryInfo`

NewClientHistoryInfo instantiates a new ClientHistoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientHistoryInfoWithDefaults

`func NewClientHistoryInfoWithDefaults() *ClientHistoryInfo`

NewClientHistoryInfoWithDefaults instantiates a new ClientHistoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationTime

`func (o *ClientHistoryInfo) GetAssociationTime() int32`

GetAssociationTime returns the AssociationTime field if non-nil, zero value otherwise.

### GetAssociationTimeOk

`func (o *ClientHistoryInfo) GetAssociationTimeOk() (*int32, bool)`

GetAssociationTimeOk returns a tuple with the AssociationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTime

`func (o *ClientHistoryInfo) SetAssociationTime(v int32)`

SetAssociationTime sets AssociationTime field to given value.

### HasAssociationTime

`func (o *ClientHistoryInfo) HasAssociationTime() bool`

HasAssociationTime returns a boolean if a field has been set.

### GetAuthInfo

`func (o *ClientHistoryInfo) GetAuthInfo() []AuthInfoOpenApiVO`

GetAuthInfo returns the AuthInfo field if non-nil, zero value otherwise.

### GetAuthInfoOk

`func (o *ClientHistoryInfo) GetAuthInfoOk() (*[]AuthInfoOpenApiVO, bool)`

GetAuthInfoOk returns a tuple with the AuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInfo

`func (o *ClientHistoryInfo) SetAuthInfo(v []AuthInfoOpenApiVO)`

SetAuthInfo sets AuthInfo field to given value.

### HasAuthInfo

`func (o *ClientHistoryInfo) HasAuthInfo() bool`

HasAuthInfo returns a boolean if a field has been set.

### GetDeviceName

`func (o *ClientHistoryInfo) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ClientHistoryInfo) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ClientHistoryInfo) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ClientHistoryInfo) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDownload

`func (o *ClientHistoryInfo) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ClientHistoryInfo) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ClientHistoryInfo) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ClientHistoryInfo) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDuration

`func (o *ClientHistoryInfo) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClientHistoryInfo) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClientHistoryInfo) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClientHistoryInfo) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFirstSeen

`func (o *ClientHistoryInfo) GetFirstSeen() int64`

GetFirstSeen returns the FirstSeen field if non-nil, zero value otherwise.

### GetFirstSeenOk

`func (o *ClientHistoryInfo) GetFirstSeenOk() (*int64, bool)`

GetFirstSeenOk returns a tuple with the FirstSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSeen

`func (o *ClientHistoryInfo) SetFirstSeen(v int64)`

SetFirstSeen sets FirstSeen field to given value.

### HasFirstSeen

`func (o *ClientHistoryInfo) HasFirstSeen() bool`

HasFirstSeen returns a boolean if a field has been set.

### GetGuest

`func (o *ClientHistoryInfo) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientHistoryInfo) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientHistoryInfo) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientHistoryInfo) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetId

`func (o *ClientHistoryInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientHistoryInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientHistoryInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientHistoryInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *ClientHistoryInfo) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientHistoryInfo) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientHistoryInfo) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientHistoryInfo) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *ClientHistoryInfo) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *ClientHistoryInfo) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *ClientHistoryInfo) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *ClientHistoryInfo) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLastSeen

`func (o *ClientHistoryInfo) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClientHistoryInfo) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClientHistoryInfo) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ClientHistoryInfo) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *ClientHistoryInfo) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientHistoryInfo) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientHistoryInfo) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientHistoryInfo) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *ClientHistoryInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientHistoryInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientHistoryInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientHistoryInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ClientHistoryInfo) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClientHistoryInfo) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClientHistoryInfo) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ClientHistoryInfo) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetSsid

`func (o *ClientHistoryInfo) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientHistoryInfo) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientHistoryInfo) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientHistoryInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetUpload

`func (o *ClientHistoryInfo) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ClientHistoryInfo) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ClientHistoryInfo) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ClientHistoryInfo) HasUpload() bool`

HasUpload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


