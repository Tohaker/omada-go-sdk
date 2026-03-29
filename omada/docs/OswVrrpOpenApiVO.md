# OswVrrpOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiseTimer** | **int32** | AdvertiseTimer should be within the range of 1-255, the default setting value is 100 and the unit is millimeter second. | 
**Authentication** | **int32** | Authentication type should be a value as follows: 0: NONE, 1: Simple, 2: MD5 | 
**DelayTime** | **int32** | DelayTime should be within the range of 0-255, default setting is 0 and the unit is second | 
**DeviceList** | Pointer to [**[]OswVrrpDeviceOpenApiVO**](OswVrrpDeviceOpenApiVO.md) | Up to 8 entries are allowed for the deviceList, and it cannot be empty. | [optional] 
**Id** | Pointer to **string** | VRRP ID | [optional] 
**IpV6Global** | Pointer to **[]string** | Up to 15 entries are allowed for the ipV6Global. | [optional] 
**IpV6LinkLocal** | Pointer to **string** | Virtual IpV6LinkLocal Address | [optional] 
**Ipv4MasterDevice** | Pointer to **string** | IPV4 Master Device | [optional] 
**Ipv6MasterDevice** | Pointer to **string** | IPV6 Master Device | [optional] 
**Key** | Pointer to **string** | Key of VRRP | [optional] 
**Name** | **string** | VRRP name, it should be visible ASCII, and should contain 1 to 64 characters. | 
**PreemptMode** | **bool** | Whether the preemption mode is enabled or not, it defaults to enabled. | 
**Status** | Pointer to **int32** | The current status of VRRP, it should be a value as follows: 0: NORMAL, 1: FAILED, 2: ABNORMAL | [optional] 
**VirtualIpv4s** | Pointer to **[]string** | Up to 32 entries are allowed for the virtualIpv4s. | [optional] 
**VrId** | **int32** | Virtual Router ID, its value should be within the range of 1-255. | 

## Methods

### NewOswVrrpOpenApiVO

`func NewOswVrrpOpenApiVO(advertiseTimer int32, authentication int32, delayTime int32, name string, preemptMode bool, vrId int32, ) *OswVrrpOpenApiVO`

NewOswVrrpOpenApiVO instantiates a new OswVrrpOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrrpOpenApiVOWithDefaults

`func NewOswVrrpOpenApiVOWithDefaults() *OswVrrpOpenApiVO`

NewOswVrrpOpenApiVOWithDefaults instantiates a new OswVrrpOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiseTimer

`func (o *OswVrrpOpenApiVO) GetAdvertiseTimer() int32`

GetAdvertiseTimer returns the AdvertiseTimer field if non-nil, zero value otherwise.

### GetAdvertiseTimerOk

`func (o *OswVrrpOpenApiVO) GetAdvertiseTimerOk() (*int32, bool)`

GetAdvertiseTimerOk returns a tuple with the AdvertiseTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseTimer

`func (o *OswVrrpOpenApiVO) SetAdvertiseTimer(v int32)`

SetAdvertiseTimer sets AdvertiseTimer field to given value.


### GetAuthentication

`func (o *OswVrrpOpenApiVO) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *OswVrrpOpenApiVO) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *OswVrrpOpenApiVO) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.


### GetDelayTime

`func (o *OswVrrpOpenApiVO) GetDelayTime() int32`

GetDelayTime returns the DelayTime field if non-nil, zero value otherwise.

### GetDelayTimeOk

`func (o *OswVrrpOpenApiVO) GetDelayTimeOk() (*int32, bool)`

GetDelayTimeOk returns a tuple with the DelayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTime

`func (o *OswVrrpOpenApiVO) SetDelayTime(v int32)`

SetDelayTime sets DelayTime field to given value.


### GetDeviceList

`func (o *OswVrrpOpenApiVO) GetDeviceList() []OswVrrpDeviceOpenApiVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *OswVrrpOpenApiVO) GetDeviceListOk() (*[]OswVrrpDeviceOpenApiVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *OswVrrpOpenApiVO) SetDeviceList(v []OswVrrpDeviceOpenApiVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *OswVrrpOpenApiVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetId

`func (o *OswVrrpOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswVrrpOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswVrrpOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswVrrpOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIpV6Global

`func (o *OswVrrpOpenApiVO) GetIpV6Global() []string`

GetIpV6Global returns the IpV6Global field if non-nil, zero value otherwise.

### GetIpV6GlobalOk

`func (o *OswVrrpOpenApiVO) GetIpV6GlobalOk() (*[]string, bool)`

GetIpV6GlobalOk returns a tuple with the IpV6Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Global

`func (o *OswVrrpOpenApiVO) SetIpV6Global(v []string)`

SetIpV6Global sets IpV6Global field to given value.

### HasIpV6Global

`func (o *OswVrrpOpenApiVO) HasIpV6Global() bool`

HasIpV6Global returns a boolean if a field has been set.

### GetIpV6LinkLocal

`func (o *OswVrrpOpenApiVO) GetIpV6LinkLocal() string`

GetIpV6LinkLocal returns the IpV6LinkLocal field if non-nil, zero value otherwise.

### GetIpV6LinkLocalOk

`func (o *OswVrrpOpenApiVO) GetIpV6LinkLocalOk() (*string, bool)`

GetIpV6LinkLocalOk returns a tuple with the IpV6LinkLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6LinkLocal

`func (o *OswVrrpOpenApiVO) SetIpV6LinkLocal(v string)`

SetIpV6LinkLocal sets IpV6LinkLocal field to given value.

### HasIpV6LinkLocal

`func (o *OswVrrpOpenApiVO) HasIpV6LinkLocal() bool`

HasIpV6LinkLocal returns a boolean if a field has been set.

### GetIpv4MasterDevice

`func (o *OswVrrpOpenApiVO) GetIpv4MasterDevice() string`

GetIpv4MasterDevice returns the Ipv4MasterDevice field if non-nil, zero value otherwise.

### GetIpv4MasterDeviceOk

`func (o *OswVrrpOpenApiVO) GetIpv4MasterDeviceOk() (*string, bool)`

GetIpv4MasterDeviceOk returns a tuple with the Ipv4MasterDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv4MasterDevice

`func (o *OswVrrpOpenApiVO) SetIpv4MasterDevice(v string)`

SetIpv4MasterDevice sets Ipv4MasterDevice field to given value.

### HasIpv4MasterDevice

`func (o *OswVrrpOpenApiVO) HasIpv4MasterDevice() bool`

HasIpv4MasterDevice returns a boolean if a field has been set.

### GetIpv6MasterDevice

`func (o *OswVrrpOpenApiVO) GetIpv6MasterDevice() string`

GetIpv6MasterDevice returns the Ipv6MasterDevice field if non-nil, zero value otherwise.

### GetIpv6MasterDeviceOk

`func (o *OswVrrpOpenApiVO) GetIpv6MasterDeviceOk() (*string, bool)`

GetIpv6MasterDeviceOk returns a tuple with the Ipv6MasterDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpv6MasterDevice

`func (o *OswVrrpOpenApiVO) SetIpv6MasterDevice(v string)`

SetIpv6MasterDevice sets Ipv6MasterDevice field to given value.

### HasIpv6MasterDevice

`func (o *OswVrrpOpenApiVO) HasIpv6MasterDevice() bool`

HasIpv6MasterDevice returns a boolean if a field has been set.

### GetKey

`func (o *OswVrrpOpenApiVO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OswVrrpOpenApiVO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OswVrrpOpenApiVO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OswVrrpOpenApiVO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *OswVrrpOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswVrrpOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswVrrpOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPreemptMode

`func (o *OswVrrpOpenApiVO) GetPreemptMode() bool`

GetPreemptMode returns the PreemptMode field if non-nil, zero value otherwise.

### GetPreemptModeOk

`func (o *OswVrrpOpenApiVO) GetPreemptModeOk() (*bool, bool)`

GetPreemptModeOk returns a tuple with the PreemptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptMode

`func (o *OswVrrpOpenApiVO) SetPreemptMode(v bool)`

SetPreemptMode sets PreemptMode field to given value.


### GetStatus

`func (o *OswVrrpOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswVrrpOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswVrrpOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswVrrpOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVirtualIpv4s

`func (o *OswVrrpOpenApiVO) GetVirtualIpv4s() []string`

GetVirtualIpv4s returns the VirtualIpv4s field if non-nil, zero value otherwise.

### GetVirtualIpv4sOk

`func (o *OswVrrpOpenApiVO) GetVirtualIpv4sOk() (*[]string, bool)`

GetVirtualIpv4sOk returns a tuple with the VirtualIpv4s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpv4s

`func (o *OswVrrpOpenApiVO) SetVirtualIpv4s(v []string)`

SetVirtualIpv4s sets VirtualIpv4s field to given value.

### HasVirtualIpv4s

`func (o *OswVrrpOpenApiVO) HasVirtualIpv4s() bool`

HasVirtualIpv4s returns a boolean if a field has been set.

### GetVrId

`func (o *OswVrrpOpenApiVO) GetVrId() int32`

GetVrId returns the VrId field if non-nil, zero value otherwise.

### GetVrIdOk

`func (o *OswVrrpOpenApiVO) GetVrIdOk() (*int32, bool)`

GetVrIdOk returns a tuple with the VrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrId

`func (o *OswVrrpOpenApiVO) SetVrId(v int32)`

SetVrId sets VrId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


