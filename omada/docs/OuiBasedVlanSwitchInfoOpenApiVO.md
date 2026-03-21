# OuiBasedVlanSwitchInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfiguredInOuiId** | Pointer to **string** | Indicate ID of \&quot;oldFirmwareDevice\&quot; is configured in specific oui based vlan. | [optional] 
**Lags** | Pointer to [**[]LagInfoOpenApiVO**](LagInfoOpenApiVO.md) | Switch lag info list. | [optional] 
**Mac** | Pointer to **string** | Device mac. | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**OldFirmwareDevice** | Pointer to **bool** | If switch supports only one oui combine, this key word is \&quot;true\&quot; | [optional] 
**PortNum** | Pointer to **int32** | Switch port number. Valid port range is from 1 to \&quot;portNum\&quot; | [optional] 

## Methods

### NewOuiBasedVlanSwitchInfoOpenApiVO

`func NewOuiBasedVlanSwitchInfoOpenApiVO() *OuiBasedVlanSwitchInfoOpenApiVO`

NewOuiBasedVlanSwitchInfoOpenApiVO instantiates a new OuiBasedVlanSwitchInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOuiBasedVlanSwitchInfoOpenApiVOWithDefaults

`func NewOuiBasedVlanSwitchInfoOpenApiVOWithDefaults() *OuiBasedVlanSwitchInfoOpenApiVO`

NewOuiBasedVlanSwitchInfoOpenApiVOWithDefaults instantiates a new OuiBasedVlanSwitchInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguredInOuiId

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetConfiguredInOuiId() string`

GetConfiguredInOuiId returns the ConfiguredInOuiId field if non-nil, zero value otherwise.

### GetConfiguredInOuiIdOk

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetConfiguredInOuiIdOk() (*string, bool)`

GetConfiguredInOuiIdOk returns a tuple with the ConfiguredInOuiId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguredInOuiId

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) SetConfiguredInOuiId(v string)`

SetConfiguredInOuiId sets ConfiguredInOuiId field to given value.

### HasConfiguredInOuiId

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) HasConfiguredInOuiId() bool`

HasConfiguredInOuiId returns a boolean if a field has been set.

### GetLags

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetLags() []LagInfoOpenApiVO`

GetLags returns the Lags field if non-nil, zero value otherwise.

### GetLagsOk

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetLagsOk() (*[]LagInfoOpenApiVO, bool)`

GetLagsOk returns a tuple with the Lags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLags

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) SetLags(v []LagInfoOpenApiVO)`

SetLags sets Lags field to given value.

### HasLags

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) HasLags() bool`

HasLags returns a boolean if a field has been set.

### GetMac

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOldFirmwareDevice

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetOldFirmwareDevice() bool`

GetOldFirmwareDevice returns the OldFirmwareDevice field if non-nil, zero value otherwise.

### GetOldFirmwareDeviceOk

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetOldFirmwareDeviceOk() (*bool, bool)`

GetOldFirmwareDeviceOk returns a tuple with the OldFirmwareDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldFirmwareDevice

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) SetOldFirmwareDevice(v bool)`

SetOldFirmwareDevice sets OldFirmwareDevice field to given value.

### HasOldFirmwareDevice

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) HasOldFirmwareDevice() bool`

HasOldFirmwareDevice returns a boolean if a field has been set.

### GetPortNum

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *OuiBasedVlanSwitchInfoOpenApiVO) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


