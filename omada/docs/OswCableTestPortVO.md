# OswCableTestPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DownlinkDevices** | Pointer to [**[]DeviceBriefVO**](DeviceBriefVO.md) | Downlink Devices | [optional] 
**EditEnable** | Pointer to **bool** | It indicates whether the port can run cable test. | [optional] 
**IsDisabled** | Pointer to **bool** | It indicates whether the port is disabled. | [optional] 
**IsStackPort** | Pointer to **bool** | It indicates whether the port is stack port | [optional] 
**IsUpperPort** | Pointer to **bool** | It indicates whether the port is upper port. | [optional] 
**Name** | Pointer to **string** | Port name. | [optional] 
**Port** | Pointer to **int32** | Port Id | [optional] 
**Reasons** | Pointer to **[]int32** | Only valid when editEnable is false.It indicates why the port can not run cable test. Each item should be a value as follows: -1:uplink port. -2:disabled port. -3:SFP port. -4:console port. -5:USB port. -5:Stack port | [optional] 
**StandardPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 
**Type** | Pointer to **int32** | Port type, it should be a value as follows: 1: Copper, 2: Combo, 3: SFP | [optional] 

## Methods

### NewOswCableTestPortVO

`func NewOswCableTestPortVO() *OswCableTestPortVO`

NewOswCableTestPortVO instantiates a new OswCableTestPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswCableTestPortVOWithDefaults

`func NewOswCableTestPortVOWithDefaults() *OswCableTestPortVO`

NewOswCableTestPortVOWithDefaults instantiates a new OswCableTestPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDownlinkDevices

`func (o *OswCableTestPortVO) GetDownlinkDevices() []DeviceBriefVO`

GetDownlinkDevices returns the DownlinkDevices field if non-nil, zero value otherwise.

### GetDownlinkDevicesOk

`func (o *OswCableTestPortVO) GetDownlinkDevicesOk() (*[]DeviceBriefVO, bool)`

GetDownlinkDevicesOk returns a tuple with the DownlinkDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlinkDevices

`func (o *OswCableTestPortVO) SetDownlinkDevices(v []DeviceBriefVO)`

SetDownlinkDevices sets DownlinkDevices field to given value.

### HasDownlinkDevices

`func (o *OswCableTestPortVO) HasDownlinkDevices() bool`

HasDownlinkDevices returns a boolean if a field has been set.

### GetEditEnable

`func (o *OswCableTestPortVO) GetEditEnable() bool`

GetEditEnable returns the EditEnable field if non-nil, zero value otherwise.

### GetEditEnableOk

`func (o *OswCableTestPortVO) GetEditEnableOk() (*bool, bool)`

GetEditEnableOk returns a tuple with the EditEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditEnable

`func (o *OswCableTestPortVO) SetEditEnable(v bool)`

SetEditEnable sets EditEnable field to given value.

### HasEditEnable

`func (o *OswCableTestPortVO) HasEditEnable() bool`

HasEditEnable returns a boolean if a field has been set.

### GetIsDisabled

`func (o *OswCableTestPortVO) GetIsDisabled() bool`

GetIsDisabled returns the IsDisabled field if non-nil, zero value otherwise.

### GetIsDisabledOk

`func (o *OswCableTestPortVO) GetIsDisabledOk() (*bool, bool)`

GetIsDisabledOk returns a tuple with the IsDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisabled

`func (o *OswCableTestPortVO) SetIsDisabled(v bool)`

SetIsDisabled sets IsDisabled field to given value.

### HasIsDisabled

`func (o *OswCableTestPortVO) HasIsDisabled() bool`

HasIsDisabled returns a boolean if a field has been set.

### GetIsStackPort

`func (o *OswCableTestPortVO) GetIsStackPort() bool`

GetIsStackPort returns the IsStackPort field if non-nil, zero value otherwise.

### GetIsStackPortOk

`func (o *OswCableTestPortVO) GetIsStackPortOk() (*bool, bool)`

GetIsStackPortOk returns a tuple with the IsStackPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStackPort

`func (o *OswCableTestPortVO) SetIsStackPort(v bool)`

SetIsStackPort sets IsStackPort field to given value.

### HasIsStackPort

`func (o *OswCableTestPortVO) HasIsStackPort() bool`

HasIsStackPort returns a boolean if a field has been set.

### GetIsUpperPort

`func (o *OswCableTestPortVO) GetIsUpperPort() bool`

GetIsUpperPort returns the IsUpperPort field if non-nil, zero value otherwise.

### GetIsUpperPortOk

`func (o *OswCableTestPortVO) GetIsUpperPortOk() (*bool, bool)`

GetIsUpperPortOk returns a tuple with the IsUpperPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUpperPort

`func (o *OswCableTestPortVO) SetIsUpperPort(v bool)`

SetIsUpperPort sets IsUpperPort field to given value.

### HasIsUpperPort

`func (o *OswCableTestPortVO) HasIsUpperPort() bool`

HasIsUpperPort returns a boolean if a field has been set.

### GetName

`func (o *OswCableTestPortVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswCableTestPortVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswCableTestPortVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswCableTestPortVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPort

`func (o *OswCableTestPortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswCableTestPortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswCableTestPortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswCableTestPortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetReasons

`func (o *OswCableTestPortVO) GetReasons() []int32`

GetReasons returns the Reasons field if non-nil, zero value otherwise.

### GetReasonsOk

`func (o *OswCableTestPortVO) GetReasonsOk() (*[]int32, bool)`

GetReasonsOk returns a tuple with the Reasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasons

`func (o *OswCableTestPortVO) SetReasons(v []int32)`

SetReasons sets Reasons field to given value.

### HasReasons

`func (o *OswCableTestPortVO) HasReasons() bool`

HasReasons returns a boolean if a field has been set.

### GetStandardPort

`func (o *OswCableTestPortVO) GetStandardPort() OswStandPortVO`

GetStandardPort returns the StandardPort field if non-nil, zero value otherwise.

### GetStandardPortOk

`func (o *OswCableTestPortVO) GetStandardPortOk() (*OswStandPortVO, bool)`

GetStandardPortOk returns a tuple with the StandardPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardPort

`func (o *OswCableTestPortVO) SetStandardPort(v OswStandPortVO)`

SetStandardPort sets StandardPort field to given value.

### HasStandardPort

`func (o *OswCableTestPortVO) HasStandardPort() bool`

HasStandardPort returns a boolean if a field has been set.

### GetType

`func (o *OswCableTestPortVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswCableTestPortVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswCableTestPortVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswCableTestPortVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


