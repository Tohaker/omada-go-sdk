# VirtualWanDslOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncapMode** | Pointer to **int32** | Encap mode assigned by your ISP. 0: LLC, 1: VC-MUX, 2: VC/MUX, 3: 1483 Bridged IP LLC, 4: 1483 Routed IP LLC. | [optional] 
**Isp** | **int32** | Your ISP (Internet Service Provider) from the drop-down list,Select Other to customize the settings. | 
**Location** | **string** | Your country/region. | 
**MerEnable** | Pointer to **bool** | If your ISP requires MER for network connection(e.g., Sky VDSL).MER switch configuration, only configurable when in VDSL mode and DHCP dial-up mode. | [optional] 
**MerPassword** | Pointer to **string** | MER password configuration, only configurable when in VDSL mode and DHCP dial-up mode. | [optional] 
**MerUsername** | Pointer to **string** | MER username configuration, only configurable when in VDSL mode and DHCP dial-up mode. | [optional] 
**ModulationType** | **int32** | The modulation type used for your DSL connection. 0: VDSL, 1: ADSL. | 
**Vci** | Pointer to **int32** | The VCI(1~65535) assigned by your ISP to specify the virtual channel endpoints in an ATM network. | [optional] 
**Vpi** | Pointer to **int32** | The VPI(0~255) assigned by your ISP to specify the virtual path between endpoints in an ATM network. | [optional] 

## Methods

### NewVirtualWanDslOpenApiVO

`func NewVirtualWanDslOpenApiVO(isp int32, location string, modulationType int32, ) *VirtualWanDslOpenApiVO`

NewVirtualWanDslOpenApiVO instantiates a new VirtualWanDslOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualWanDslOpenApiVOWithDefaults

`func NewVirtualWanDslOpenApiVOWithDefaults() *VirtualWanDslOpenApiVO`

NewVirtualWanDslOpenApiVOWithDefaults instantiates a new VirtualWanDslOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncapMode

`func (o *VirtualWanDslOpenApiVO) GetEncapMode() int32`

GetEncapMode returns the EncapMode field if non-nil, zero value otherwise.

### GetEncapModeOk

`func (o *VirtualWanDslOpenApiVO) GetEncapModeOk() (*int32, bool)`

GetEncapModeOk returns a tuple with the EncapMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapMode

`func (o *VirtualWanDslOpenApiVO) SetEncapMode(v int32)`

SetEncapMode sets EncapMode field to given value.

### HasEncapMode

`func (o *VirtualWanDslOpenApiVO) HasEncapMode() bool`

HasEncapMode returns a boolean if a field has been set.

### GetIsp

`func (o *VirtualWanDslOpenApiVO) GetIsp() int32`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *VirtualWanDslOpenApiVO) GetIspOk() (*int32, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *VirtualWanDslOpenApiVO) SetIsp(v int32)`

SetIsp sets Isp field to given value.


### GetLocation

`func (o *VirtualWanDslOpenApiVO) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VirtualWanDslOpenApiVO) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VirtualWanDslOpenApiVO) SetLocation(v string)`

SetLocation sets Location field to given value.


### GetMerEnable

`func (o *VirtualWanDslOpenApiVO) GetMerEnable() bool`

GetMerEnable returns the MerEnable field if non-nil, zero value otherwise.

### GetMerEnableOk

`func (o *VirtualWanDslOpenApiVO) GetMerEnableOk() (*bool, bool)`

GetMerEnableOk returns a tuple with the MerEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerEnable

`func (o *VirtualWanDslOpenApiVO) SetMerEnable(v bool)`

SetMerEnable sets MerEnable field to given value.

### HasMerEnable

`func (o *VirtualWanDslOpenApiVO) HasMerEnable() bool`

HasMerEnable returns a boolean if a field has been set.

### GetMerPassword

`func (o *VirtualWanDslOpenApiVO) GetMerPassword() string`

GetMerPassword returns the MerPassword field if non-nil, zero value otherwise.

### GetMerPasswordOk

`func (o *VirtualWanDslOpenApiVO) GetMerPasswordOk() (*string, bool)`

GetMerPasswordOk returns a tuple with the MerPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerPassword

`func (o *VirtualWanDslOpenApiVO) SetMerPassword(v string)`

SetMerPassword sets MerPassword field to given value.

### HasMerPassword

`func (o *VirtualWanDslOpenApiVO) HasMerPassword() bool`

HasMerPassword returns a boolean if a field has been set.

### GetMerUsername

`func (o *VirtualWanDslOpenApiVO) GetMerUsername() string`

GetMerUsername returns the MerUsername field if non-nil, zero value otherwise.

### GetMerUsernameOk

`func (o *VirtualWanDslOpenApiVO) GetMerUsernameOk() (*string, bool)`

GetMerUsernameOk returns a tuple with the MerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerUsername

`func (o *VirtualWanDslOpenApiVO) SetMerUsername(v string)`

SetMerUsername sets MerUsername field to given value.

### HasMerUsername

`func (o *VirtualWanDslOpenApiVO) HasMerUsername() bool`

HasMerUsername returns a boolean if a field has been set.

### GetModulationType

`func (o *VirtualWanDslOpenApiVO) GetModulationType() int32`

GetModulationType returns the ModulationType field if non-nil, zero value otherwise.

### GetModulationTypeOk

`func (o *VirtualWanDslOpenApiVO) GetModulationTypeOk() (*int32, bool)`

GetModulationTypeOk returns a tuple with the ModulationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModulationType

`func (o *VirtualWanDslOpenApiVO) SetModulationType(v int32)`

SetModulationType sets ModulationType field to given value.


### GetVci

`func (o *VirtualWanDslOpenApiVO) GetVci() int32`

GetVci returns the Vci field if non-nil, zero value otherwise.

### GetVciOk

`func (o *VirtualWanDslOpenApiVO) GetVciOk() (*int32, bool)`

GetVciOk returns a tuple with the Vci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVci

`func (o *VirtualWanDslOpenApiVO) SetVci(v int32)`

SetVci sets Vci field to given value.

### HasVci

`func (o *VirtualWanDslOpenApiVO) HasVci() bool`

HasVci returns a boolean if a field has been set.

### GetVpi

`func (o *VirtualWanDslOpenApiVO) GetVpi() int32`

GetVpi returns the Vpi field if non-nil, zero value otherwise.

### GetVpiOk

`func (o *VirtualWanDslOpenApiVO) GetVpiOk() (*int32, bool)`

GetVpiOk returns a tuple with the Vpi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpi

`func (o *VirtualWanDslOpenApiVO) SetVpi(v int32)`

SetVpi sets Vpi field to given value.

### HasVpi

`func (o *VirtualWanDslOpenApiVO) HasVpi() bool`

HasVpi returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


