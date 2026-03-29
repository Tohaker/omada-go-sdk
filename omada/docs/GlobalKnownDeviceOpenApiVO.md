# GlobalKnownDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUtil** | Pointer to **int32** | Device cpuUtil | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type. 0 means basic, 1 means pro. | [optional] 
**Ip** | Pointer to **string** | Device IP | [optional] 
**Ipv6** | Pointer to **[]string** | Device IPv6 list | [optional] 
**LastSeen** | Pointer to **int64** | Device lastSeen | [optional] 
**LicenseStatus** | Pointer to **int32** | Device license status(Only for cloud base) should be a value as follows: 0: unActive; 1: Unbind; 2: Expired; 3: active | [optional] 
**Mac** | Pointer to **string** | Device MAC | [optional] 
**MemUtil** | Pointer to **int32** | Device memUtil | [optional] 
**Model** | Pointer to **string** | Device model name | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**NeedActive** | Pointer to **bool** | Device license status(Only for cloud base).If the value is true, the device is ready to be activated;If the value is false, the device cannot be activated or has already been activated. | [optional] 
**SiteName** | Pointer to **string** | The name of the site where the device is located | [optional] 
**Sn** | Pointer to **string** | Device serial number | [optional] 
**Status** | Pointer to **int32** | Device status should be a value as follows: 0: Disconnected; 1: Connected; 2: Pending; 3: Heartbeat Missed; 4: Isolated | [optional] 
**Subtype** | Pointer to **string** | Switch subtype should be a value as follows: smart: Non-Agile Series Switch; es: Agile Series Switch. | [optional] 
**TagName** | Pointer to **string** | Device tag name | [optional] 
**Type** | Pointer to **string** | Device type | [optional] 
**Uptime** | Pointer to **string** | Device uptime | [optional] 

## Methods

### NewGlobalKnownDeviceOpenApiVO

`func NewGlobalKnownDeviceOpenApiVO() *GlobalKnownDeviceOpenApiVO`

NewGlobalKnownDeviceOpenApiVO instantiates a new GlobalKnownDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalKnownDeviceOpenApiVOWithDefaults

`func NewGlobalKnownDeviceOpenApiVOWithDefaults() *GlobalKnownDeviceOpenApiVO`

NewGlobalKnownDeviceOpenApiVOWithDefaults instantiates a new GlobalKnownDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUtil

`func (o *GlobalKnownDeviceOpenApiVO) GetCpuUtil() int32`

GetCpuUtil returns the CpuUtil field if non-nil, zero value otherwise.

### GetCpuUtilOk

`func (o *GlobalKnownDeviceOpenApiVO) GetCpuUtilOk() (*int32, bool)`

GetCpuUtilOk returns a tuple with the CpuUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUtil

`func (o *GlobalKnownDeviceOpenApiVO) SetCpuUtil(v int32)`

SetCpuUtil sets CpuUtil field to given value.

### HasCpuUtil

`func (o *GlobalKnownDeviceOpenApiVO) HasCpuUtil() bool`

HasCpuUtil returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *GlobalKnownDeviceOpenApiVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *GlobalKnownDeviceOpenApiVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *GlobalKnownDeviceOpenApiVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *GlobalKnownDeviceOpenApiVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetIp

`func (o *GlobalKnownDeviceOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *GlobalKnownDeviceOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *GlobalKnownDeviceOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *GlobalKnownDeviceOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetIpv6

`func (o *GlobalKnownDeviceOpenApiVO) GetIpv6() []string`

GetIpv6 returns the Ipv6 field if non-nil, zero value otherwise.

### GetIpv6Ok

`func (o *GlobalKnownDeviceOpenApiVO) GetIpv6Ok() (*[]string, bool)`

GetIpv6Ok returns a tuple with the Ipv6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6

`func (o *GlobalKnownDeviceOpenApiVO) SetIpv6(v []string)`

SetIpv6 sets Ipv6 field to given value.

### HasIpv6

`func (o *GlobalKnownDeviceOpenApiVO) HasIpv6() bool`

HasIpv6 returns a boolean if a field has been set.

### GetLastSeen

`func (o *GlobalKnownDeviceOpenApiVO) GetLastSeen() int64`

GetLastSeen returns the LastSeen field if non-nil, zero value otherwise.

### GetLastSeenOk

`func (o *GlobalKnownDeviceOpenApiVO) GetLastSeenOk() (*int64, bool)`

GetLastSeenOk returns a tuple with the LastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeen

`func (o *GlobalKnownDeviceOpenApiVO) SetLastSeen(v int64)`

SetLastSeen sets LastSeen field to given value.

### HasLastSeen

`func (o *GlobalKnownDeviceOpenApiVO) HasLastSeen() bool`

HasLastSeen returns a boolean if a field has been set.

### GetLicenseStatus

`func (o *GlobalKnownDeviceOpenApiVO) GetLicenseStatus() int32`

GetLicenseStatus returns the LicenseStatus field if non-nil, zero value otherwise.

### GetLicenseStatusOk

`func (o *GlobalKnownDeviceOpenApiVO) GetLicenseStatusOk() (*int32, bool)`

GetLicenseStatusOk returns a tuple with the LicenseStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicenseStatus

`func (o *GlobalKnownDeviceOpenApiVO) SetLicenseStatus(v int32)`

SetLicenseStatus sets LicenseStatus field to given value.

### HasLicenseStatus

`func (o *GlobalKnownDeviceOpenApiVO) HasLicenseStatus() bool`

HasLicenseStatus returns a boolean if a field has been set.

### GetMac

`func (o *GlobalKnownDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *GlobalKnownDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *GlobalKnownDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *GlobalKnownDeviceOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMemUtil

`func (o *GlobalKnownDeviceOpenApiVO) GetMemUtil() int32`

GetMemUtil returns the MemUtil field if non-nil, zero value otherwise.

### GetMemUtilOk

`func (o *GlobalKnownDeviceOpenApiVO) GetMemUtilOk() (*int32, bool)`

GetMemUtilOk returns a tuple with the MemUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemUtil

`func (o *GlobalKnownDeviceOpenApiVO) SetMemUtil(v int32)`

SetMemUtil sets MemUtil field to given value.

### HasMemUtil

`func (o *GlobalKnownDeviceOpenApiVO) HasMemUtil() bool`

HasMemUtil returns a boolean if a field has been set.

### GetModel

`func (o *GlobalKnownDeviceOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GlobalKnownDeviceOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GlobalKnownDeviceOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GlobalKnownDeviceOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *GlobalKnownDeviceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GlobalKnownDeviceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GlobalKnownDeviceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GlobalKnownDeviceOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedActive

`func (o *GlobalKnownDeviceOpenApiVO) GetNeedActive() bool`

GetNeedActive returns the NeedActive field if non-nil, zero value otherwise.

### GetNeedActiveOk

`func (o *GlobalKnownDeviceOpenApiVO) GetNeedActiveOk() (*bool, bool)`

GetNeedActiveOk returns a tuple with the NeedActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedActive

`func (o *GlobalKnownDeviceOpenApiVO) SetNeedActive(v bool)`

SetNeedActive sets NeedActive field to given value.

### HasNeedActive

`func (o *GlobalKnownDeviceOpenApiVO) HasNeedActive() bool`

HasNeedActive returns a boolean if a field has been set.

### GetSiteName

`func (o *GlobalKnownDeviceOpenApiVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *GlobalKnownDeviceOpenApiVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *GlobalKnownDeviceOpenApiVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *GlobalKnownDeviceOpenApiVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSn

`func (o *GlobalKnownDeviceOpenApiVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *GlobalKnownDeviceOpenApiVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *GlobalKnownDeviceOpenApiVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *GlobalKnownDeviceOpenApiVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *GlobalKnownDeviceOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GlobalKnownDeviceOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GlobalKnownDeviceOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GlobalKnownDeviceOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubtype

`func (o *GlobalKnownDeviceOpenApiVO) GetSubtype() string`

GetSubtype returns the Subtype field if non-nil, zero value otherwise.

### GetSubtypeOk

`func (o *GlobalKnownDeviceOpenApiVO) GetSubtypeOk() (*string, bool)`

GetSubtypeOk returns a tuple with the Subtype field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtype

`func (o *GlobalKnownDeviceOpenApiVO) SetSubtype(v string)`

SetSubtype sets Subtype field to given value.

### HasSubtype

`func (o *GlobalKnownDeviceOpenApiVO) HasSubtype() bool`

HasSubtype returns a boolean if a field has been set.

### GetTagName

`func (o *GlobalKnownDeviceOpenApiVO) GetTagName() string`

GetTagName returns the TagName field if non-nil, zero value otherwise.

### GetTagNameOk

`func (o *GlobalKnownDeviceOpenApiVO) GetTagNameOk() (*string, bool)`

GetTagNameOk returns a tuple with the TagName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagName

`func (o *GlobalKnownDeviceOpenApiVO) SetTagName(v string)`

SetTagName sets TagName field to given value.

### HasTagName

`func (o *GlobalKnownDeviceOpenApiVO) HasTagName() bool`

HasTagName returns a boolean if a field has been set.

### GetType

`func (o *GlobalKnownDeviceOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GlobalKnownDeviceOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GlobalKnownDeviceOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GlobalKnownDeviceOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUptime

`func (o *GlobalKnownDeviceOpenApiVO) GetUptime() string`

GetUptime returns the Uptime field if non-nil, zero value otherwise.

### GetUptimeOk

`func (o *GlobalKnownDeviceOpenApiVO) GetUptimeOk() (*string, bool)`

GetUptimeOk returns a tuple with the Uptime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUptime

`func (o *GlobalKnownDeviceOpenApiVO) SetUptime(v string)`

SetUptime sets Uptime field to given value.

### HasUptime

`func (o *GlobalKnownDeviceOpenApiVO) HasUptime() bool`

HasUptime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


