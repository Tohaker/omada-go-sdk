# ClientHistoryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Associated** | Pointer to **int64** |  | [optional] 
**AssociationTime** | Pointer to **int32** |  | [optional] 
**AuthInfo** | Pointer to [**[]AuthInfoVO**](AuthInfoVO.md) |  | [optional] 
**DeviceName** | Pointer to **string** |  | [optional] 
**Download** | Pointer to **int64** |  | [optional] 
**Duration** | Pointer to **int64** |  | [optional] 
**Guest** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ipv6List** | Pointer to **[]string** |  | [optional] 
**LastSeen** | Pointer to **int64** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **int32** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**ReasonType** | Pointer to **int32** |  | [optional] 
**Ssid** | Pointer to **string** |  | [optional] 
**Upload** | Pointer to **int64** |  | [optional] 
**Vid** | Pointer to **int32** |  | [optional] 

## Methods

### NewClientHistoryVO

`func NewClientHistoryVO() *ClientHistoryVO`

NewClientHistoryVO instantiates a new ClientHistoryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientHistoryVOWithDefaults

`func NewClientHistoryVOWithDefaults() *ClientHistoryVO`

NewClientHistoryVOWithDefaults instantiates a new ClientHistoryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociated

`func (o *ClientHistoryVO) GetAssociated() int64`

GetAssociated returns the Associated field if non-nil, zero value otherwise.

### GetAssociatedOk

`func (o *ClientHistoryVO) GetAssociatedOk() (*int64, bool)`

GetAssociatedOk returns a tuple with the Associated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociated

`func (o *ClientHistoryVO) SetAssociated(v int64)`

SetAssociated sets Associated field to given value.

### HasAssociated

`func (o *ClientHistoryVO) HasAssociated() bool`

HasAssociated returns a boolean if a field has been set.

### GetAssociationTime

`func (o *ClientHistoryVO) GetAssociationTime() int32`

GetAssociationTime returns the AssociationTime field if non-nil, zero value otherwise.

### GetAssociationTimeOk

`func (o *ClientHistoryVO) GetAssociationTimeOk() (*int32, bool)`

GetAssociationTimeOk returns a tuple with the AssociationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationTime

`func (o *ClientHistoryVO) SetAssociationTime(v int32)`

SetAssociationTime sets AssociationTime field to given value.

### HasAssociationTime

`func (o *ClientHistoryVO) HasAssociationTime() bool`

HasAssociationTime returns a boolean if a field has been set.

### GetAuthInfo

`func (o *ClientHistoryVO) GetAuthInfo() []AuthInfoVO`

GetAuthInfo returns the AuthInfo field if non-nil, zero value otherwise.

### GetAuthInfoOk

`func (o *ClientHistoryVO) GetAuthInfoOk() (*[]AuthInfoVO, bool)`

GetAuthInfoOk returns a tuple with the AuthInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthInfo

`func (o *ClientHistoryVO) SetAuthInfo(v []AuthInfoVO)`

SetAuthInfo sets AuthInfo field to given value.

### HasAuthInfo

`func (o *ClientHistoryVO) HasAuthInfo() bool`

HasAuthInfo returns a boolean if a field has been set.

### GetDeviceName

`func (o *ClientHistoryVO) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *ClientHistoryVO) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *ClientHistoryVO) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *ClientHistoryVO) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDownload

`func (o *ClientHistoryVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *ClientHistoryVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *ClientHistoryVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *ClientHistoryVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDuration

`func (o *ClientHistoryVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *ClientHistoryVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *ClientHistoryVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *ClientHistoryVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetGuest

`func (o *ClientHistoryVO) GetGuest() bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientHistoryVO) GetGuestOk() (*bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientHistoryVO) SetGuest(v bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientHistoryVO) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetId

`func (o *ClientHistoryVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClientHistoryVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClientHistoryVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ClientHistoryVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *ClientHistoryVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ClientHistoryVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ClientHistoryVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ClientHistoryVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *ClientHistoryVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *ClientHistoryVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *ClientHistoryVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *ClientHistoryVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLastSeen

`func (o *ClientHistoryVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *ClientHistoryVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *ClientHistoryVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *ClientHistoryVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetMac

`func (o *ClientHistoryVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ClientHistoryVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ClientHistoryVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ClientHistoryVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *ClientHistoryVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClientHistoryVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClientHistoryVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClientHistoryVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *ClientHistoryVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ClientHistoryVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ClientHistoryVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ClientHistoryVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReason

`func (o *ClientHistoryVO) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ClientHistoryVO) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ClientHistoryVO) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ClientHistoryVO) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetReasonType

`func (o *ClientHistoryVO) GetReasonType() int32`

GetReasonType returns the ReasonType field if non-nil, zero value otherwise.

### GetReasonTypeOk

`func (o *ClientHistoryVO) GetReasonTypeOk() (*int32, bool)`

GetReasonTypeOk returns a tuple with the ReasonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonType

`func (o *ClientHistoryVO) SetReasonType(v int32)`

SetReasonType sets ReasonType field to given value.

### HasReasonType

`func (o *ClientHistoryVO) HasReasonType() bool`

HasReasonType returns a boolean if a field has been set.

### GetSsid

`func (o *ClientHistoryVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientHistoryVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientHistoryVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientHistoryVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetUpload

`func (o *ClientHistoryVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *ClientHistoryVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *ClientHistoryVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *ClientHistoryVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetVid

`func (o *ClientHistoryVO) GetVid() int32`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *ClientHistoryVO) GetVidOk() (*int32, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *ClientHistoryVO) SetVid(v int32)`

SetVid sets Vid field to given value.

### HasVid

`func (o *ClientHistoryVO) HasVid() bool`

HasVid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


