# MlagAdoptOswVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** |  | [optional] 
**AddedInAdvanced** | Pointer to **bool** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**DeviceType** | Pointer to **int32** |  | [optional] 
**DueTime** | Pointer to **int64** |  | [optional] 
**DueTimeLeft** | Pointer to **int64** |  | [optional] 
**Ip** | Pointer to **string** |  | [optional] 
**Ipv6List** | Pointer to **[]string** |  | [optional] 
**LicenseStatus** | Pointer to **int32** |  | [optional] 
**Mac** | Pointer to **string** |  | [optional] 
**MlagGroupId** | Pointer to **int32** |  | [optional] 
**MlagVersion** | Pointer to **string** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**ModelVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Ports** | Pointer to [**[]OswMlagPortVO**](OswMlagPortVO.md) |  | [optional] 
**PublicIp** | Pointer to **string** |  | [optional] 
**ShowModel** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**SupportMlagIpv6** | Pointer to **bool** |  | [optional] 
**SupportPeerLinkPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewMlagAdoptOswVO

`func NewMlagAdoptOswVO() *MlagAdoptOswVO`

NewMlagAdoptOswVO instantiates a new MlagAdoptOswVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlagAdoptOswVOWithDefaults

`func NewMlagAdoptOswVOWithDefaults() *MlagAdoptOswVO`

NewMlagAdoptOswVOWithDefaults instantiates a new MlagAdoptOswVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *MlagAdoptOswVO) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *MlagAdoptOswVO) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *MlagAdoptOswVO) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *MlagAdoptOswVO) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAddedInAdvanced

`func (o *MlagAdoptOswVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *MlagAdoptOswVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *MlagAdoptOswVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *MlagAdoptOswVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetCategory

`func (o *MlagAdoptOswVO) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *MlagAdoptOswVO) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *MlagAdoptOswVO) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *MlagAdoptOswVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDeviceType

`func (o *MlagAdoptOswVO) GetDeviceType() int32`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *MlagAdoptOswVO) GetDeviceTypeOk() (*int32, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *MlagAdoptOswVO) SetDeviceType(v int32)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *MlagAdoptOswVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetDueTime

`func (o *MlagAdoptOswVO) GetDueTime() int64`

GetDueTime returns the DueTime field if non-nil, zero value otherwise.

### GetDueTimeOk

`func (o *MlagAdoptOswVO) GetDueTimeOk() (*int64, bool)`

GetDueTimeOk returns a tuple with the DueTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTime

`func (o *MlagAdoptOswVO) SetDueTime(v int64)`

SetDueTime sets DueTime field to given value.

### HasDueTime

`func (o *MlagAdoptOswVO) HasDueTime() bool`

HasDueTime returns a boolean if a field has been set.

### GetDueTimeLeft

`func (o *MlagAdoptOswVO) GetDueTimeLeft() int64`

GetDueTimeLeft returns the DueTimeLeft field if non-nil, zero value otherwise.

### GetDueTimeLeftOk

`func (o *MlagAdoptOswVO) GetDueTimeLeftOk() (*int64, bool)`

GetDueTimeLeftOk returns a tuple with the DueTimeLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueTimeLeft

`func (o *MlagAdoptOswVO) SetDueTimeLeft(v int64)`

SetDueTimeLeft sets DueTimeLeft field to given value.

### HasDueTimeLeft

`func (o *MlagAdoptOswVO) HasDueTimeLeft() bool`

HasDueTimeLeft returns a boolean if a field has been set.

### GetIp

`func (o *MlagAdoptOswVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MlagAdoptOswVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MlagAdoptOswVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MlagAdoptOswVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6List

`func (o *MlagAdoptOswVO) GetIpv6List() []string`

GetIpv6List returns the Ipv6List field if non-nil, zero value otherwise.

### GetIpv6ListOk

`func (o *MlagAdoptOswVO) GetIpv6ListOk() (*[]string, bool)`

GetIpv6ListOk returns a tuple with the Ipv6List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6List

`func (o *MlagAdoptOswVO) SetIpv6List(v []string)`

SetIpv6List sets Ipv6List field to given value.

### HasIpv6List

`func (o *MlagAdoptOswVO) HasIpv6List() bool`

HasIpv6List returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *MlagAdoptOswVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *MlagAdoptOswVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *MlagAdoptOswVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *MlagAdoptOswVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetMac

`func (o *MlagAdoptOswVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MlagAdoptOswVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MlagAdoptOswVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MlagAdoptOswVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMlagGroupId

`func (o *MlagAdoptOswVO) GetMlagGroupId() int32`

GetMlagGroupId returns the MlagGroupId field if non-nil, zero value otherwise.

### GetMlagGroupIdOk

`func (o *MlagAdoptOswVO) GetMlagGroupIdOk() (*int32, bool)`

GetMlagGroupIdOk returns a tuple with the MlagGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagGroupId

`func (o *MlagAdoptOswVO) SetMlagGroupId(v int32)`

SetMlagGroupId sets MlagGroupId field to given value.

### HasMlagGroupId

`func (o *MlagAdoptOswVO) HasMlagGroupId() bool`

HasMlagGroupId returns a boolean if a field has been set.

### GetMlagVersion

`func (o *MlagAdoptOswVO) GetMlagVersion() string`

GetMlagVersion returns the MlagVersion field if non-nil, zero value otherwise.

### GetMlagVersionOk

`func (o *MlagAdoptOswVO) GetMlagVersionOk() (*string, bool)`

GetMlagVersionOk returns a tuple with the MlagVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagVersion

`func (o *MlagAdoptOswVO) SetMlagVersion(v string)`

SetMlagVersion sets MlagVersion field to given value.

### HasMlagVersion

`func (o *MlagAdoptOswVO) HasMlagVersion() bool`

HasMlagVersion returns a boolean if a field has been set.

### GetModel

`func (o *MlagAdoptOswVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MlagAdoptOswVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MlagAdoptOswVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MlagAdoptOswVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *MlagAdoptOswVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *MlagAdoptOswVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *MlagAdoptOswVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *MlagAdoptOswVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *MlagAdoptOswVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MlagAdoptOswVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MlagAdoptOswVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MlagAdoptOswVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *MlagAdoptOswVO) GetPorts() []OswMlagPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *MlagAdoptOswVO) GetPortsOk() (*[]OswMlagPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *MlagAdoptOswVO) SetPorts(v []OswMlagPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *MlagAdoptOswVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetPublicIp

`func (o *MlagAdoptOswVO) GetPublicIp() string`

GetPublicIp returns the PublicIp field if non-nil, zero value otherwise.

### GetPublicIpOk

`func (o *MlagAdoptOswVO) GetPublicIpOk() (*string, bool)`

GetPublicIpOk returns a tuple with the PublicIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIp

`func (o *MlagAdoptOswVO) SetPublicIp(v string)`

SetPublicIp sets PublicIp field to given value.

### HasPublicIp

`func (o *MlagAdoptOswVO) HasPublicIp() bool`

HasPublicIp returns a boolean if a field has been set.

### GetShowModel

`func (o *MlagAdoptOswVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *MlagAdoptOswVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *MlagAdoptOswVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *MlagAdoptOswVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStatus

`func (o *MlagAdoptOswVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MlagAdoptOswVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MlagAdoptOswVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MlagAdoptOswVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportMlagIpv6

`func (o *MlagAdoptOswVO) GetSupportMlagIpv6() bool`

GetSupportMlagIpv6 returns the SupportMlagIpv6 field if non-nil, zero value otherwise.

### GetSupportMlagIpv6Ok

`func (o *MlagAdoptOswVO) GetSupportMlagIpv6Ok() (*bool, bool)`

GetSupportMlagIpv6Ok returns a tuple with the SupportMlagIpv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportMlagIpv6

`func (o *MlagAdoptOswVO) SetSupportMlagIpv6(v bool)`

SetSupportMlagIpv6 sets SupportMlagIpv6 field to given value.

### HasSupportMlagIpv6

`func (o *MlagAdoptOswVO) HasSupportMlagIpv6() bool`

HasSupportMlagIpv6 returns a boolean if a field has been set.

### GetSupportPeerLinkPorts

`func (o *MlagAdoptOswVO) GetSupportPeerLinkPorts() []OswStandPortVO`

GetSupportPeerLinkPorts returns the SupportPeerLinkPorts field if non-nil, zero value otherwise.

### GetSupportPeerLinkPortsOk

`func (o *MlagAdoptOswVO) GetSupportPeerLinkPortsOk() (*[]OswStandPortVO, bool)`

GetSupportPeerLinkPortsOk returns a tuple with the SupportPeerLinkPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPeerLinkPorts

`func (o *MlagAdoptOswVO) SetSupportPeerLinkPorts(v []OswStandPortVO)`

SetSupportPeerLinkPorts sets SupportPeerLinkPorts field to given value.

### HasSupportPeerLinkPorts

`func (o *MlagAdoptOswVO) HasSupportPeerLinkPorts() bool`

HasSupportPeerLinkPorts returns a boolean if a field has been set.

### GetVersion

`func (o *MlagAdoptOswVO) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MlagAdoptOswVO) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MlagAdoptOswVO) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *MlagAdoptOswVO) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


