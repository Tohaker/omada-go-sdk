# ClientQueryFiltersOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthStatus** | Pointer to **[]int32** | Filter by authStatus, 0: CONNECTED, 1: PENDING, 2: AUTHED, 3: AUTH_FREE | [optional] 
**AuthType** | Pointer to **string** | Filter by authentication type | [optional] 
**Device** | Pointer to **[]string** | Filter by the name of connected device. | [optional] 
**DeviceMac** | Pointer to **[]string** | Filter by the mac of connected device. | [optional] 
**DeviceType** | Pointer to **[]string** | Filter by the type of client. | [optional] 
**Guest** | Pointer to **[]bool** | Filter by guest, effective only when wireless is true. | [optional] 
**Health** | Pointer to **int32** | Filter by health score, 0:No Data, 1 : poor 2 : fair 3 : good | [optional] 
**IpcNvr** | Pointer to **bool** | Filter by ipcNvr, This filter will not affect clientStat result. true : filter by client type IPC or NVR, false or null: NOT filter by client type. | [optional] 
**Model** | Pointer to **string** | Filter by the model. | [optional] 
**Network** | Pointer to **[]string** | Filter by lan network name. | [optional] 
**RadioId** | Pointer to **[]int32** | Filter by radio Id. 0 : 2.4G, 1 : 5G-1, 2 : 5G-2, 3 : 6G | [optional] 
**SourceTime** | Pointer to **int64** | Filter by history health data. | [optional] 
**Ssid** | Pointer to **[]string** | Filter by ssid. | [optional] 
**TimeEnd** | Pointer to **int64** | Filter by lastSeen. | [optional] 
**TimeStart** | Pointer to **int64** | Filter by lastSeen. | [optional] 
**Vendor** | Pointer to **[]string** | Filter by vendor. | [optional] 
**Vid** | Pointer to **string** | Filter by VLAN id, vid should be a string like 1,100-200. | [optional] 
**Wireless** | Pointer to **bool** | Filter by wireless, This filter will not affect clientStat result. | [optional] 

## Methods

### NewClientQueryFiltersOpenApiVO

`func NewClientQueryFiltersOpenApiVO() *ClientQueryFiltersOpenApiVO`

NewClientQueryFiltersOpenApiVO instantiates a new ClientQueryFiltersOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientQueryFiltersOpenApiVOWithDefaults

`func NewClientQueryFiltersOpenApiVOWithDefaults() *ClientQueryFiltersOpenApiVO`

NewClientQueryFiltersOpenApiVOWithDefaults instantiates a new ClientQueryFiltersOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthStatus

`func (o *ClientQueryFiltersOpenApiVO) GetAuthStatus() []int32`

GetAuthStatus returns the AuthStatus field if non-nil, zero value otherwise.

### GetAuthStatusOk

`func (o *ClientQueryFiltersOpenApiVO) GetAuthStatusOk() (*[]int32, bool)`

GetAuthStatusOk returns a tuple with the AuthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthStatus

`func (o *ClientQueryFiltersOpenApiVO) SetAuthStatus(v []int32)`

SetAuthStatus sets AuthStatus field to given value.

### HasAuthStatus

`func (o *ClientQueryFiltersOpenApiVO) HasAuthStatus() bool`

HasAuthStatus returns a boolean if a field has been set.

### GetAuthType

`func (o *ClientQueryFiltersOpenApiVO) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ClientQueryFiltersOpenApiVO) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ClientQueryFiltersOpenApiVO) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ClientQueryFiltersOpenApiVO) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetDevice

`func (o *ClientQueryFiltersOpenApiVO) GetDevice() []string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ClientQueryFiltersOpenApiVO) GetDeviceOk() (*[]string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ClientQueryFiltersOpenApiVO) SetDevice(v []string)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *ClientQueryFiltersOpenApiVO) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetDeviceMac

`func (o *ClientQueryFiltersOpenApiVO) GetDeviceMac() []string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *ClientQueryFiltersOpenApiVO) GetDeviceMacOk() (*[]string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *ClientQueryFiltersOpenApiVO) SetDeviceMac(v []string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *ClientQueryFiltersOpenApiVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetDeviceType

`func (o *ClientQueryFiltersOpenApiVO) GetDeviceType() []string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *ClientQueryFiltersOpenApiVO) GetDeviceTypeOk() (*[]string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *ClientQueryFiltersOpenApiVO) SetDeviceType(v []string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *ClientQueryFiltersOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetGuest

`func (o *ClientQueryFiltersOpenApiVO) GetGuest() []bool`

GetGuest returns the Guest field if non-nil, zero value otherwise.

### GetGuestOk

`func (o *ClientQueryFiltersOpenApiVO) GetGuestOk() (*[]bool, bool)`

GetGuestOk returns a tuple with the Guest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuest

`func (o *ClientQueryFiltersOpenApiVO) SetGuest(v []bool)`

SetGuest sets Guest field to given value.

### HasGuest

`func (o *ClientQueryFiltersOpenApiVO) HasGuest() bool`

HasGuest returns a boolean if a field has been set.

### GetHealth

`func (o *ClientQueryFiltersOpenApiVO) GetHealth() int32`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *ClientQueryFiltersOpenApiVO) GetHealthOk() (*int32, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *ClientQueryFiltersOpenApiVO) SetHealth(v int32)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *ClientQueryFiltersOpenApiVO) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetIpcNvr

`func (o *ClientQueryFiltersOpenApiVO) GetIpcNvr() bool`

GetIpcNvr returns the IpcNvr field if non-nil, zero value otherwise.

### GetIpcNvrOk

`func (o *ClientQueryFiltersOpenApiVO) GetIpcNvrOk() (*bool, bool)`

GetIpcNvrOk returns a tuple with the IpcNvr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpcNvr

`func (o *ClientQueryFiltersOpenApiVO) SetIpcNvr(v bool)`

SetIpcNvr sets IpcNvr field to given value.

### HasIpcNvr

`func (o *ClientQueryFiltersOpenApiVO) HasIpcNvr() bool`

HasIpcNvr returns a boolean if a field has been set.

### GetModel

`func (o *ClientQueryFiltersOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ClientQueryFiltersOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ClientQueryFiltersOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ClientQueryFiltersOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetNetwork

`func (o *ClientQueryFiltersOpenApiVO) GetNetwork() []string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClientQueryFiltersOpenApiVO) GetNetworkOk() (*[]string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClientQueryFiltersOpenApiVO) SetNetwork(v []string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClientQueryFiltersOpenApiVO) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetRadioId

`func (o *ClientQueryFiltersOpenApiVO) GetRadioId() []int32`

GetRadioId returns the RadioId field if non-nil, zero value otherwise.

### GetRadioIdOk

`func (o *ClientQueryFiltersOpenApiVO) GetRadioIdOk() (*[]int32, bool)`

GetRadioIdOk returns a tuple with the RadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadioId

`func (o *ClientQueryFiltersOpenApiVO) SetRadioId(v []int32)`

SetRadioId sets RadioId field to given value.

### HasRadioId

`func (o *ClientQueryFiltersOpenApiVO) HasRadioId() bool`

HasRadioId returns a boolean if a field has been set.

### GetSourceTime

`func (o *ClientQueryFiltersOpenApiVO) GetSourceTime() int64`

GetSourceTime returns the SourceTime field if non-nil, zero value otherwise.

### GetSourceTimeOk

`func (o *ClientQueryFiltersOpenApiVO) GetSourceTimeOk() (*int64, bool)`

GetSourceTimeOk returns a tuple with the SourceTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTime

`func (o *ClientQueryFiltersOpenApiVO) SetSourceTime(v int64)`

SetSourceTime sets SourceTime field to given value.

### HasSourceTime

`func (o *ClientQueryFiltersOpenApiVO) HasSourceTime() bool`

HasSourceTime returns a boolean if a field has been set.

### GetSsid

`func (o *ClientQueryFiltersOpenApiVO) GetSsid() []string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *ClientQueryFiltersOpenApiVO) GetSsidOk() (*[]string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *ClientQueryFiltersOpenApiVO) SetSsid(v []string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *ClientQueryFiltersOpenApiVO) HasSsid() bool`

HasSsid returns a boolean if a field has been set.

### GetTimeEnd

`func (o *ClientQueryFiltersOpenApiVO) GetTimeEnd() int64`

GetTimeEnd returns the TimeEnd field if non-nil, zero value otherwise.

### GetTimeEndOk

`func (o *ClientQueryFiltersOpenApiVO) GetTimeEndOk() (*int64, bool)`

GetTimeEndOk returns a tuple with the TimeEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeEnd

`func (o *ClientQueryFiltersOpenApiVO) SetTimeEnd(v int64)`

SetTimeEnd sets TimeEnd field to given value.

### HasTimeEnd

`func (o *ClientQueryFiltersOpenApiVO) HasTimeEnd() bool`

HasTimeEnd returns a boolean if a field has been set.

### GetTimeStart

`func (o *ClientQueryFiltersOpenApiVO) GetTimeStart() int64`

GetTimeStart returns the TimeStart field if non-nil, zero value otherwise.

### GetTimeStartOk

`func (o *ClientQueryFiltersOpenApiVO) GetTimeStartOk() (*int64, bool)`

GetTimeStartOk returns a tuple with the TimeStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStart

`func (o *ClientQueryFiltersOpenApiVO) SetTimeStart(v int64)`

SetTimeStart sets TimeStart field to given value.

### HasTimeStart

`func (o *ClientQueryFiltersOpenApiVO) HasTimeStart() bool`

HasTimeStart returns a boolean if a field has been set.

### GetVendor

`func (o *ClientQueryFiltersOpenApiVO) GetVendor() []string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *ClientQueryFiltersOpenApiVO) GetVendorOk() (*[]string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *ClientQueryFiltersOpenApiVO) SetVendor(v []string)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *ClientQueryFiltersOpenApiVO) HasVendor() bool`

HasVendor returns a boolean if a field has been set.

### GetVid

`func (o *ClientQueryFiltersOpenApiVO) GetVid() string`

GetVid returns the Vid field if non-nil, zero value otherwise.

### GetVidOk

`func (o *ClientQueryFiltersOpenApiVO) GetVidOk() (*string, bool)`

GetVidOk returns a tuple with the Vid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVid

`func (o *ClientQueryFiltersOpenApiVO) SetVid(v string)`

SetVid sets Vid field to given value.

### HasVid

`func (o *ClientQueryFiltersOpenApiVO) HasVid() bool`

HasVid returns a boolean if a field has been set.

### GetWireless

`func (o *ClientQueryFiltersOpenApiVO) GetWireless() bool`

GetWireless returns the Wireless field if non-nil, zero value otherwise.

### GetWirelessOk

`func (o *ClientQueryFiltersOpenApiVO) GetWirelessOk() (*bool, bool)`

GetWirelessOk returns a tuple with the Wireless field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireless

`func (o *ClientQueryFiltersOpenApiVO) SetWireless(v bool)`

SetWireless sets Wireless field to given value.

### HasWireless

`func (o *ClientQueryFiltersOpenApiVO) HasWireless() bool`

HasWireless returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


