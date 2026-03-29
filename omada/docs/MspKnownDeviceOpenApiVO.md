# MspKnownDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | Device cpuUtil | [optional] 
**CustomerName** | Pointer to **string** | The name of the customer where the device is located | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type. 0 means basic, 1 means pro. | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Ipv6** | Pointer to **[]string** | Device IPv6 list | [optional] 
**LastSeen** | Pointer to **int64** | Device lastSeen, unit: ms | [optional] 
**LicenseStatus** | Pointer to **int32** | Device license status(Only for pro site) should be a value as follows: 0:unActive; 1:Unbind; 2:Expired; 3:active | [optional] 
**Mac** | Pointer to **string** | Device MAC | [optional] 
**MemUtil** | Pointer to **int32** | Device memUtil | [optional] 
**Model** | Pointer to **string** | Device model name | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**NeedActive** | Pointer to **bool** | Device license status(Only for pro site).If the value is true, the device is ready to be activated;If the value is false, the device cannot be activated or has already been activated. | [optional] 
**SiteName** | Pointer to **string** | The name of the site where the device is located | [optional] 
**Sn** | Pointer to **string** | Device serial number | [optional] 
**Status** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 
**Subtype** | Pointer to **string** | Switch subtype should be a value as follows: smart: Non-Agile Series Switch; es: Agile Series Switch. | [optional] 
**TagName** | Pointer to **string** | Device tag name | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 
**Uptime** | Pointer to **string** | Device uptime | [optional] 

## Methods

### NewMspKnownDeviceOpenApiVO

`func NewMspKnownDeviceOpenApiVO() *MspKnownDeviceOpenApiVO`

NewMspKnownDeviceOpenApiVO instantiates a new MspKnownDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMspKnownDeviceOpenApiVOWithDefaults

`func NewMspKnownDeviceOpenApiVOWithDefaults() *MspKnownDeviceOpenApiVO`

NewMspKnownDeviceOpenApiVOWithDefaults instantiates a new MspKnownDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *MspKnownDeviceOpenApiVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *MspKnownDeviceOpenApiVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *MspKnownDeviceOpenApiVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *MspKnownDeviceOpenApiVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetCustomerName

`func (o *MspKnownDeviceOpenApiVO) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *MspKnownDeviceOpenApiVO) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *MspKnownDeviceOpenApiVO) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *MspKnownDeviceOpenApiVO) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *MspKnownDeviceOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *MspKnownDeviceOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *MspKnownDeviceOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *MspKnownDeviceOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetIp

`func (o *MspKnownDeviceOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *MspKnownDeviceOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *MspKnownDeviceOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *MspKnownDeviceOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *MspKnownDeviceOpenApiVO) GetIpv6() []string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *MspKnownDeviceOpenApiVO) GetIpv6Ok() (*[]string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *MspKnownDeviceOpenApiVO) SetIpv6(v []string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *MspKnownDeviceOpenApiVO) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLastSeen

`func (o *MspKnownDeviceOpenApiVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *MspKnownDeviceOpenApiVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *MspKnownDeviceOpenApiVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *MspKnownDeviceOpenApiVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *MspKnownDeviceOpenApiVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *MspKnownDeviceOpenApiVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *MspKnownDeviceOpenApiVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *MspKnownDeviceOpenApiVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetMac

`func (o *MspKnownDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *MspKnownDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *MspKnownDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *MspKnownDeviceOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *MspKnownDeviceOpenApiVO) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *MspKnownDeviceOpenApiVO) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *MspKnownDeviceOpenApiVO) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *MspKnownDeviceOpenApiVO) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetModel

`func (o *MspKnownDeviceOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MspKnownDeviceOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MspKnownDeviceOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *MspKnownDeviceOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *MspKnownDeviceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MspKnownDeviceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MspKnownDeviceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MspKnownDeviceOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedActive

`func (o *MspKnownDeviceOpenApiVO) GetNeedActive() bool`

GetNeedActive returns the NeedActive field if non-nil, zero value otherwise.

### GetNeedActiveOk

`func (o *MspKnownDeviceOpenApiVO) GetNeedActiveOk() (*bool, bool)`

GetNeedActiveOk returns a tuple with the NeedActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedActive

`func (o *MspKnownDeviceOpenApiVO) SetNeedActive(v bool)`

SetNeedActive sets NeedActive field to given value.

### HasNeedActive

`func (o *MspKnownDeviceOpenApiVO) HasNeedActive() bool`

HasNeedActive returns a boolean if a field has been set.

### GetSiteName

`func (o *MspKnownDeviceOpenApiVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *MspKnownDeviceOpenApiVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *MspKnownDeviceOpenApiVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *MspKnownDeviceOpenApiVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSn

`func (o *MspKnownDeviceOpenApiVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *MspKnownDeviceOpenApiVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *MspKnownDeviceOpenApiVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *MspKnownDeviceOpenApiVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *MspKnownDeviceOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MspKnownDeviceOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MspKnownDeviceOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MspKnownDeviceOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *MspKnownDeviceOpenApiVO) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *MspKnownDeviceOpenApiVO) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *MspKnownDeviceOpenApiVO) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *MspKnownDeviceOpenApiVO) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTagName

`func (o *MspKnownDeviceOpenApiVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *MspKnownDeviceOpenApiVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *MspKnownDeviceOpenApiVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *MspKnownDeviceOpenApiVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetType

`func (o *MspKnownDeviceOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MspKnownDeviceOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MspKnownDeviceOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MspKnownDeviceOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *MspKnownDeviceOpenApiVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *MspKnownDeviceOpenApiVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *MspKnownDeviceOpenApiVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *MspKnownDeviceOpenApiVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


