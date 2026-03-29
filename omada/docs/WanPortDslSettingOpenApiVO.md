# WanPortDslSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncapMode** | Pointer to **int32** | Encap mode assigned by your ISP. 0: LLC, 1: VC-MUX, 2: VC/MUX, 3: 1483 Bridged IP LLC, 4: 1483 Routed IP LLC | [optional] 
**Isp** | **int32** | Your ISP (Internet Service Provider) from the drop-down list,Select Other to customize the settings. | 
**Location** | **string** | Your country/region. | 
**MerEnable** | Pointer to **bool** | If your ISP requires MER for network connection(e.g., Sky VDSL).MER switch configuration, only configurable when in VDSL mode and DHCP dial-up mode. | [optional] 
**MerPassword** | Pointer to **string** | MER password configuration, only configurable when in VDSL mode and DHCP dial-up mode. | [optional] 
**MerUsername** | Pointer to **string** | MER username configuration, only configurable when in VDSL mode and DHCP dial-up mode. | [optional] 
**ModulationType** | **int32** | The modulation type used for your DSL connection. 0: VDSL, 1: ADSL. | 
**PortDesc** | Pointer to **string** | Wan port description. | [optional] 
**PortName** | Pointer to **string** | Wan port name. | [optional] 
**PortUuid** | **string** | Wan port UUID. | 
**Vci** | Pointer to **int32** | The VCI(1~65535) assigned by your ISP to specify the virtual channel endpoints in an ATM network. | [optional] 
**Vpi** | Pointer to **int32** | The VPI(0~255) assigned by your ISP to specify the virtual path between endpoints in an ATM network. | [optional] 

## Methods

### NewWanPortDslSettingOpenApiVO

`func NewWanPortDslSettingOpenApiVO(isp int32, location string, modulationType int32, portUuid string, ) *WanPortDslSettingOpenApiVO`

NewWanPortDslSettingOpenApiVO instantiates a new WanPortDslSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWanPortDslSettingOpenApiVOWithDefaults

`func NewWanPortDslSettingOpenApiVOWithDefaults() *WanPortDslSettingOpenApiVO`

NewWanPortDslSettingOpenApiVOWithDefaults instantiates a new WanPortDslSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncapMode

`func (o *WanPortDslSettingOpenApiVO) GetEncapMode() int32`

GetEncapMode returns the EncapMode field if non-nil, zero value otherwise.

### GetEncapModeOk

`func (o *WanPortDslSettingOpenApiVO) GetEncapModeOk() (*int32, bool)`

GetEncapModeOk returns a tuple with the EncapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapMode

`func (o *WanPortDslSettingOpenApiVO) SetEncapMode(v int32)`

SetEncapMode sets EncapMode field to given value.

### HasEncapMode

`func (o *WanPortDslSettingOpenApiVO) HasEncapMode() bool`

HasEncapMode returns a boolean if a field has been set.

### GetIsp

`func (o *WanPortDslSettingOpenApiVO) GetIsp() int32`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *WanPortDslSettingOpenApiVO) GetIspOk() (*int32, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *WanPortDslSettingOpenApiVO) SetIsp(v int32)`

SetIsp sets Isp field to given value.


### GetLocation

`func (o *WanPortDslSettingOpenApiVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *WanPortDslSettingOpenApiVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *WanPortDslSettingOpenApiVO) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetMerEnable

`func (o *WanPortDslSettingOpenApiVO) GetMerEnable() bool`

GetMerEnable returns the MerEnable field if non-nil, zero value otherwise.

### GetMerEnableOk

`func (o *WanPortDslSettingOpenApiVO) GetMerEnableOk() (*bool, bool)`

GetMerEnableOk returns a tuple with the MerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerEnable

`func (o *WanPortDslSettingOpenApiVO) SetMerEnable(v bool)`

SetMerEnable sets MerEnable field to given value.

### HasMerEnable

`func (o *WanPortDslSettingOpenApiVO) HasMerEnable() bool`

HasMerEnable returns a boolean if a field has been set.

### GetMerPassword

`func (o *WanPortDslSettingOpenApiVO) GetMerPassword() string`

GetMerPassword returns the MerPassword field if non-nil, zero value otherwise.

### GetMerPasswordOk

`func (o *WanPortDslSettingOpenApiVO) GetMerPasswordOk() (*string, bool)`

GetMerPasswordOk returns a tuple with the MerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerPassword

`func (o *WanPortDslSettingOpenApiVO) SetMerPassword(v string)`

SetMerPassword sets MerPassword field to given value.

### HasMerPassword

`func (o *WanPortDslSettingOpenApiVO) HasMerPassword() bool`

HasMerPassword returns a boolean if a field has been set.

### GetMerUsername

`func (o *WanPortDslSettingOpenApiVO) GetMerUsername() string`

GetMerUsername returns the MerUsername field if non-nil, zero value otherwise.

### GetMerUsernameOk

`func (o *WanPortDslSettingOpenApiVO) GetMerUsernameOk() (*string, bool)`

GetMerUsernameOk returns a tuple with the MerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerUsername

`func (o *WanPortDslSettingOpenApiVO) SetMerUsername(v string)`

SetMerUsername sets MerUsername field to given value.

### HasMerUsername

`func (o *WanPortDslSettingOpenApiVO) HasMerUsername() bool`

HasMerUsername returns a boolean if a field has been set.

### GetModulationType

`func (o *WanPortDslSettingOpenApiVO) GetModulationType() int32`

GetModulationType returns the ModulationType field if non-nil, zero value otherwise.

### GetModulationTypeOk

`func (o *WanPortDslSettingOpenApiVO) GetModulationTypeOk() (*int32, bool)`

GetModulationTypeOk returns a tuple with the ModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulationType

`func (o *WanPortDslSettingOpenApiVO) SetModulationType(v int32)`

SetModulationType sets ModulationType field to given value.


### GetPortDesc

`func (o *WanPortDslSettingOpenApiVO) GetPortDesc() string`

GetPortDesc returns the PortDesc field if non-nil, zero value otherwise.

### GetPortDescOk

`func (o *WanPortDslSettingOpenApiVO) GetPortDescOk() (*string, bool)`

GetPortDescOk returns a tuple with the PortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDesc

`func (o *WanPortDslSettingOpenApiVO) SetPortDesc(v string)`

SetPortDesc sets PortDesc field to given value.

### HasPortDesc

`func (o *WanPortDslSettingOpenApiVO) HasPortDesc() bool`

HasPortDesc returns a boolean if a field has been set.

### GetPortName

`func (o *WanPortDslSettingOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *WanPortDslSettingOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *WanPortDslSettingOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *WanPortDslSettingOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortUuid

`func (o *WanPortDslSettingOpenApiVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *WanPortDslSettingOpenApiVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *WanPortDslSettingOpenApiVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.


### GetVci

`func (o *WanPortDslSettingOpenApiVO) GetVci() int32`

GetVci returns the Vci field if non-nil, zero value otherwise.

### GetVciOk

`func (o *WanPortDslSettingOpenApiVO) GetVciOk() (*int32, bool)`

GetVciOk returns a tuple with the Vci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVci

`func (o *WanPortDslSettingOpenApiVO) SetVci(v int32)`

SetVci sets Vci field to given value.

### HasVci

`func (o *WanPortDslSettingOpenApiVO) HasVci() bool`

HasVci returns a boolean if a field has been set.

### GetVpi

`func (o *WanPortDslSettingOpenApiVO) GetVpi() int32`

GetVpi returns the Vpi field if non-nil, zero value otherwise.

### GetVpiOk

`func (o *WanPortDslSettingOpenApiVO) GetVpiOk() (*int32, bool)`

GetVpiOk returns a tuple with the Vpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpi

`func (o *WanPortDslSettingOpenApiVO) SetVpi(v int32)`

SetVpi sets Vpi field to given value.

### HasVpi

`func (o *WanPortDslSettingOpenApiVO) HasVpi() bool`

HasVpi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


