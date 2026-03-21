# AuthClientOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminName** | Pointer to **string** | Admin name (authType &#x3D; ADMIN) | [optional] 
**AuthType** | Pointer to **int32** | Auth type should be a value as follows: 0: No Auth; 1: Simple Password; 2: Exrternal Radius; 3: Voucher; 4: External Portal Server; 5: Local User; 6: SMS; 7: Facebook; 8: Hotspot Radius; 9: Mac Auth (with fail over); 10: Admin auth; 12: Form auth | [optional] 
**Download** | Pointer to **int64** | Total Download (Byte) | [optional] 
**Duration** | Pointer to **int64** | Total Duration (s) | [optional] 
**End** | Pointer to **int64** | End timestamp | [optional] 
**FbVer** | Pointer to **int32** | Version of Facebook Wi-Fi (authType &#x3D; Facebook or Facebook V2) | [optional] 
**FormName** | Pointer to **string** | Form name (authType &#x3D; form auth) | [optional] 
**Id** | Pointer to **string** | AuthRecord ID | [optional] 
**Ip** | Pointer to **string** | IP address at the time of authentication | [optional] 
**LocalUserName** | Pointer to **string** | LocalUser Name (authType &#x3D; Local User) | [optional] 
**Mac** | Pointer to **string** | Client MAC | [optional] 
**Name** | Pointer to **string** | Client name | [optional] 
**NetworkName** | Pointer to **string** | Network name | [optional] 
**RadiusUsername** | Pointer to **string** | Radius userName (authType &#x3D; External RADIUS or Hotspot RADIUS) | [optional] 
**Ssid** | Pointer to **string** | SSID name | [optional] 
**Start** | Pointer to **int64** | Start timestamp | [optional] 
**Upload** | Pointer to **int64** | Total Upload (Byte) | [optional] 
**Valid** | Pointer to **bool** | Is the client valid | [optional] 
**VoucherCode** | Pointer to **string** | Voucher code (authType &#x3D; Voucher) | [optional] 
**Wireless** | Pointer to **bool** | Is it wireless type | [optional] 

## Methods

### NewAuthClientOpenApiVO

`func NewAuthClientOpenApiVO() *AuthClientOpenApiVO`

NewAuthClientOpenApiVO instantiates a new AuthClientOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthClientOpenApiVOWithDefaults

`func NewAuthClientOpenApiVOWithDefaults() *AuthClientOpenApiVO`

NewAuthClientOpenApiVOWithDefaults instantiates a new AuthClientOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminName

`func (o *AuthClientOpenApiVO) GetAdminName() string`

GetAdminName returns the AdminName field if non-nil, zero value otherwise.

### GetAdminNameOk

`func (o *AuthClientOpenApiVO) GetAdminNameOk() (*string, bool)`

GetAdminNameOk returns a tuple with the AdminName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminName

`func (o *AuthClientOpenApiVO) SetAdminName(v string)`

SetAdminName sets AdminName field to given value.

### HasAdminName

`func (o *AuthClientOpenApiVO) HasAdminName() bool`

HasAdminName returns a boolean if a field has been set.

### GetAuthType

`func (o *AuthClientOpenApiVO) GetAuthType() int32`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AuthClientOpenApiVO) GetAuthTypeOk() (*int32, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AuthClientOpenApiVO) SetAuthType(v int32)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AuthClientOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetDownload

`func (o *AuthClientOpenApiVO) GetDownload() int64`

GetDownload returns the Download field if non-nil, zero value otherwise.

### GetDownloadOk

`func (o *AuthClientOpenApiVO) GetDownloadOk() (*int64, bool)`

GetDownloadOk returns a tuple with the Download field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownload

`func (o *AuthClientOpenApiVO) SetDownload(v int64)`

SetDownload sets Download field to given value.

### HasDownload

`func (o *AuthClientOpenApiVO) HasDownload() bool`

HasDownload returns a boolean if a field has been set.

### GetDuration

`func (o *AuthClientOpenApiVO) GetDuration() int64`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *AuthClientOpenApiVO) GetDurationOk() (*int64, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *AuthClientOpenApiVO) SetDuration(v int64)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *AuthClientOpenApiVO) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEnd

`func (o *AuthClientOpenApiVO) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *AuthClientOpenApiVO) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *AuthClientOpenApiVO) SetEnd(v int64)`

SetEnd sets End field to given value.

### HasEnd

`func (o *AuthClientOpenApiVO) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetFbVer

`func (o *AuthClientOpenApiVO) GetFbVer() int32`

GetFbVer returns the FbVer field if non-nil, zero value otherwise.

### GetFbVerOk

`func (o *AuthClientOpenApiVO) GetFbVerOk() (*int32, bool)`

GetFbVerOk returns a tuple with the FbVer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFbVer

`func (o *AuthClientOpenApiVO) SetFbVer(v int32)`

SetFbVer sets FbVer field to given value.

### HasFbVer

`func (o *AuthClientOpenApiVO) HasFbVer() bool`

HasFbVer returns a boolean if a field has been set.

### GetFormName

`func (o *AuthClientOpenApiVO) GetFormName() string`

GetFormName returns the FormName field if non-nil, zero value otherwise.

### GetFormNameOk

`func (o *AuthClientOpenApiVO) GetFormNameOk() (*string, bool)`

GetFormNameOk returns a tuple with the FormName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormName

`func (o *AuthClientOpenApiVO) SetFormName(v string)`

SetFormName sets FormName field to given value.

### HasFormName

`func (o *AuthClientOpenApiVO) HasFormName() bool`

HasFormName returns a boolean if a field has been set.

### GetId

`func (o *AuthClientOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AuthClientOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AuthClientOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AuthClientOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *AuthClientOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *AuthClientOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *AuthClientOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *AuthClientOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLocalUserName

`func (o *AuthClientOpenApiVO) GetLocalUserName() string`

GetLocalUserName returns the LocalUserName field if non-nil, zero value otherwise.

### GetLocalUserNameOk

`func (o *AuthClientOpenApiVO) GetLocalUserNameOk() (*string, bool)`

GetLocalUserNameOk returns a tuple with the LocalUserName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalUserName

`func (o *AuthClientOpenApiVO) SetLocalUserName(v string)`

SetLocalUserName sets LocalUserName field to given value.

### HasLocalUserName

`func (o *AuthClientOpenApiVO) HasLocalUserName() bool`

HasLocalUserName returns a boolean if a field has been set.

### GetMac

`func (o *AuthClientOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *AuthClientOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *AuthClientOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *AuthClientOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *AuthClientOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthClientOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthClientOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthClientOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetworkName

`func (o *AuthClientOpenApiVO) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *AuthClientOpenApiVO) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *AuthClientOpenApiVO) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *AuthClientOpenApiVO) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetRadiusUsername

`func (o *AuthClientOpenApiVO) GetRadiusUsername() string`

GetRadiusUsername returns the RadiusUsername field if non-nil, zero value otherwise.

### GetRadiusUsernameOk

`func (o *AuthClientOpenApiVO) GetRadiusUsernameOk() (*string, bool)`

GetRadiusUsernameOk returns a tuple with the RadiusUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusUsername

`func (o *AuthClientOpenApiVO) SetRadiusUsername(v string)`

SetRadiusUsername sets RadiusUsername field to given value.

### HasRadiusUsername

`func (o *AuthClientOpenApiVO) HasRadiusUsername() bool`

HasRadiusUsername returns a boolean if a field has been set.

### GetSsid

`func (o *AuthClientOpenApiVO) GetSsid() string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *AuthClientOpenApiVO) GetSsidOk() (*string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *AuthClientOpenApiVO) SetSsid(v string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *AuthClientOpenApiVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetStart

`func (o *AuthClientOpenApiVO) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *AuthClientOpenApiVO) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *AuthClientOpenApiVO) SetStart(v int64)`

SetStart sets Start field to given value.

### HasStart

`func (o *AuthClientOpenApiVO) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetUpload

`func (o *AuthClientOpenApiVO) GetUpload() int64`

GetUpload returns the Upload field if non-nil, zero value otherwise.

### GetUploadOk

`func (o *AuthClientOpenApiVO) GetUploadOk() (*int64, bool)`

GetUploadOk returns a tuple with the Upload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpload

`func (o *AuthClientOpenApiVO) SetUpload(v int64)`

SetUpload sets Upload field to given value.

### HasUpload

`func (o *AuthClientOpenApiVO) HasUpload() bool`

HasUpload returns a boolean if a field has been set.

### GetValid

`func (o *AuthClientOpenApiVO) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AuthClientOpenApiVO) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AuthClientOpenApiVO) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AuthClientOpenApiVO) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetVoucherCode

`func (o *AuthClientOpenApiVO) GetVoucherCode() string`

GetVoucherCode returns the VoucherCode field if non-nil, zero value otherwise.

### GetVoucherCodeOk

`func (o *AuthClientOpenApiVO) GetVoucherCodeOk() (*string, bool)`

GetVoucherCodeOk returns a tuple with the VoucherCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoucherCode

`func (o *AuthClientOpenApiVO) SetVoucherCode(v string)`

SetVoucherCode sets VoucherCode field to given value.

### HasVoucherCode

`func (o *AuthClientOpenApiVO) HasVoucherCode() bool`

HasVoucherCode returns a boolean if a field has been set.

### GetWireless

`func (o *AuthClientOpenApiVO) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *AuthClientOpenApiVO) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *AuthClientOpenApiVO) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *AuthClientOpenApiVO) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


