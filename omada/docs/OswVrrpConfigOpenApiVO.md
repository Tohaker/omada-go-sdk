# OswVrrpConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdvertiseTimer** | **int32** | AdvertiseTimer should be within the range of 1-255, the default setting value is 100 and the unit is millimeter second. | 
**Authentication** | **int32** | Authentication type should be a value as follows: 0: NONE, 1: Simple, 2: MD5 | 
**DelayTime** | **int32** | DelayTime should be within the range of 0-255, default setting is 0 and the unit is second | 
**DeviceList** | Pointer to [**[]OswVrrpDeviceConfigOpenApiVO**](OswVrrpDeviceConfigOpenApiVO.md) | Up to 8 entries are allowed for the deviceList, and it cannot be empty. | [optional] 
**IpV6Global** | Pointer to **[]string** | Up to 15 entries are allowed for the ipV6Global. | [optional] 
**IpV6LinkLocal** | Pointer to **string** | Virtual IpV6LinkLocal Address | [optional] 
**Key** | Pointer to **string** | Key of VRRP | [optional] 
**Name** | **string** | VRRP name, it should be visible ASCII, and should contain 1 to 64 characters. | 
**PreemptMode** | **bool** | Whether the preemption mode is enabled or not, it defaults to enabled. | 
**VirtualIpv4s** | Pointer to **[]string** | Up to 32 entries are allowed for the virtualIpv4s. | [optional] 
**VrId** | **int32** | Virtual Router ID, its value should be within the range of 1-255. | 

## Methods

### NewOswVrrpConfigOpenApiVO

`func NewOswVrrpConfigOpenApiVO(advertiseTimer int32, authentication int32, delayTime int32, name string, preemptMode bool, vrId int32, ) *OswVrrpConfigOpenApiVO`

NewOswVrrpConfigOpenApiVO instantiates a new OswVrrpConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswVrrpConfigOpenApiVOWithDefaults

`func NewOswVrrpConfigOpenApiVOWithDefaults() *OswVrrpConfigOpenApiVO`

NewOswVrrpConfigOpenApiVOWithDefaults instantiates a new OswVrrpConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdvertiseTimer

`func (o *OswVrrpConfigOpenApiVO) GetAdvertiseTimer() int32`

GetAdvertiseTimer returns the AdvertiseTimer field if non-nil, zero value otherwise.

### GetAdvertiseTimerOk

`func (o *OswVrrpConfigOpenApiVO) GetAdvertiseTimerOk() (*int32, bool)`

GetAdvertiseTimerOk returns a tuple with the AdvertiseTimer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvertiseTimer

`func (o *OswVrrpConfigOpenApiVO) SetAdvertiseTimer(v int32)`

SetAdvertiseTimer sets AdvertiseTimer field to given value.


### GetAuthentication

`func (o *OswVrrpConfigOpenApiVO) GetAuthentication() int32`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *OswVrrpConfigOpenApiVO) GetAuthenticationOk() (*int32, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *OswVrrpConfigOpenApiVO) SetAuthentication(v int32)`

SetAuthentication sets Authentication field to given value.


### GetDelayTime

`func (o *OswVrrpConfigOpenApiVO) GetDelayTime() int32`

GetDelayTime returns the DelayTime field if non-nil, zero value otherwise.

### GetDelayTimeOk

`func (o *OswVrrpConfigOpenApiVO) GetDelayTimeOk() (*int32, bool)`

GetDelayTimeOk returns a tuple with the DelayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayTime

`func (o *OswVrrpConfigOpenApiVO) SetDelayTime(v int32)`

SetDelayTime sets DelayTime field to given value.


### GetDeviceList

`func (o *OswVrrpConfigOpenApiVO) GetDeviceList() []OswVrrpDeviceConfigOpenApiVO`

GetDeviceList returns the DeviceList field if non-nil, zero value otherwise.

### GetDeviceListOk

`func (o *OswVrrpConfigOpenApiVO) GetDeviceListOk() (*[]OswVrrpDeviceConfigOpenApiVO, bool)`

GetDeviceListOk returns a tuple with the DeviceList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceList

`func (o *OswVrrpConfigOpenApiVO) SetDeviceList(v []OswVrrpDeviceConfigOpenApiVO)`

SetDeviceList sets DeviceList field to given value.

### HasDeviceList

`func (o *OswVrrpConfigOpenApiVO) HasDeviceList() bool`

HasDeviceList returns a boolean if a field has been set.

### GetIpV6Global

`func (o *OswVrrpConfigOpenApiVO) GetIpV6Global() []string`

GetIpV6Global returns the IpV6Global field if non-nil, zero value otherwise.

### GetIpV6GlobalOk

`func (o *OswVrrpConfigOpenApiVO) GetIpV6GlobalOk() (*[]string, bool)`

GetIpV6GlobalOk returns a tuple with the IpV6Global field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6Global

`func (o *OswVrrpConfigOpenApiVO) SetIpV6Global(v []string)`

SetIpV6Global sets IpV6Global field to given value.

### HasIpV6Global

`func (o *OswVrrpConfigOpenApiVO) HasIpV6Global() bool`

HasIpV6Global returns a boolean if a field has been set.

### GetIpV6LinkLocal

`func (o *OswVrrpConfigOpenApiVO) GetIpV6LinkLocal() string`

GetIpV6LinkLocal returns the IpV6LinkLocal field if non-nil, zero value otherwise.

### GetIpV6LinkLocalOk

`func (o *OswVrrpConfigOpenApiVO) GetIpV6LinkLocalOk() (*string, bool)`

GetIpV6LinkLocalOk returns a tuple with the IpV6LinkLocal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpV6LinkLocal

`func (o *OswVrrpConfigOpenApiVO) SetIpV6LinkLocal(v string)`

SetIpV6LinkLocal sets IpV6LinkLocal field to given value.

### HasIpV6LinkLocal

`func (o *OswVrrpConfigOpenApiVO) HasIpV6LinkLocal() bool`

HasIpV6LinkLocal returns a boolean if a field has been set.

### GetKey

`func (o *OswVrrpConfigOpenApiVO) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *OswVrrpConfigOpenApiVO) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *OswVrrpConfigOpenApiVO) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *OswVrrpConfigOpenApiVO) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *OswVrrpConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswVrrpConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswVrrpConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPreemptMode

`func (o *OswVrrpConfigOpenApiVO) GetPreemptMode() bool`

GetPreemptMode returns the PreemptMode field if non-nil, zero value otherwise.

### GetPreemptModeOk

`func (o *OswVrrpConfigOpenApiVO) GetPreemptModeOk() (*bool, bool)`

GetPreemptModeOk returns a tuple with the PreemptMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreemptMode

`func (o *OswVrrpConfigOpenApiVO) SetPreemptMode(v bool)`

SetPreemptMode sets PreemptMode field to given value.


### GetVirtualIpv4s

`func (o *OswVrrpConfigOpenApiVO) GetVirtualIpv4s() []string`

GetVirtualIpv4s returns the VirtualIpv4s field if non-nil, zero value otherwise.

### GetVirtualIpv4sOk

`func (o *OswVrrpConfigOpenApiVO) GetVirtualIpv4sOk() (*[]string, bool)`

GetVirtualIpv4sOk returns a tuple with the VirtualIpv4s field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualIpv4s

`func (o *OswVrrpConfigOpenApiVO) SetVirtualIpv4s(v []string)`

SetVirtualIpv4s sets VirtualIpv4s field to given value.

### HasVirtualIpv4s

`func (o *OswVrrpConfigOpenApiVO) HasVirtualIpv4s() bool`

HasVirtualIpv4s returns a boolean if a field has been set.

### GetVrId

`func (o *OswVrrpConfigOpenApiVO) GetVrId() int32`

GetVrId returns the VrId field if non-nil, zero value otherwise.

### GetVrIdOk

`func (o *OswVrrpConfigOpenApiVO) GetVrIdOk() (*int32, bool)`

GetVrIdOk returns a tuple with the VrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrId

`func (o *OswVrrpConfigOpenApiVO) SetVrId(v int32)`

SetVrId sets VrId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


