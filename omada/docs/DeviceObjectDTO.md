# DeviceObjectDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApTempThresholds** | Pointer to **[]int32** | AP temperature thresholds. | [optional] 
**DeviceSeriesType** | Pointer to **int32** | Device series type. 0: Advanced, 1: Pro. | [optional] 
**FilledTempThreshold** | Pointer to **bool** | Whether temperature threshold fields have been filled. | [optional] 
**Ip** | Pointer to **string** | Ip address,such as 192.168.0.105. | [optional] 
**LagPortsMap** | Pointer to **map[string][]int32** | Switch LAG ports map. Key is LAG ID, value is member port set. | [optional] 
**Model** | Pointer to **string** | Device model, such as EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device, for example:3.0. | [optional] 
**Name** | Pointer to **string** | Default uses the MAC address as the name. | [optional] 
**OsgTempThreshold** | Pointer to **int32** | Gateway temperature threshold. | [optional] 
**OswCpuTempThreshold** | Pointer to **int32** | Switch CPU temperature threshold. | [optional] 
**OswMacTempThreshold** | Pointer to **int32** | Switch MAC chip temperature threshold. | [optional] 
**OswPhyTempThreshold** | Pointer to **int32** | Switch PHY temperature threshold. | [optional] 
**OswPseTempThreshold** | Pointer to **int32** | Switch PSE temperature threshold. | [optional] 
**PortNum** | Pointer to **int32** | Number of ports.  | [optional] 
**Radio2gEnable** | Pointer to **bool** | Whether 2.4 GHz radio is enabled. | [optional] 
**Radio5g2Enable** | Pointer to **bool** | Whether 5 GHz-2 radio is enabled. | [optional] 
**Radio5gEnable** | Pointer to **bool** | Whether 5 GHz radio is enabled. | [optional] 
**Radio6gEnable** | Pointer to **bool** | Whether 6 GHz radio is enabled. | [optional] 
**Support5g** | Pointer to **bool** | Whether the device supports 5 GHz radio. | [optional] 
**Support5g2** | Pointer to **bool** | Whether the device supports 5 GHZ-2 radio. | [optional] 
**Support6g** | Pointer to **bool** | Whether the device supports 6 GHz radio. | [optional] 
**SupportAnomaly** | Pointer to **bool** | Whether the device firmware support intelligent anomaly detection. | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt. | [optional] 

## Methods

### NewDeviceObjectDTO

`func NewDeviceObjectDTO() *DeviceObjectDTO`

NewDeviceObjectDTO instantiates a new DeviceObjectDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceObjectDTOWithDefaults

`func NewDeviceObjectDTOWithDefaults() *DeviceObjectDTO`

NewDeviceObjectDTOWithDefaults instantiates a new DeviceObjectDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApTempThresholds

`func (o *DeviceObjectDTO) GetApTempThresholds() []int32`

GetApTempThresholds returns the ApTempThresholds field if non-nil, zero value otherwise.

### GetApTempThresholdsOk

`func (o *DeviceObjectDTO) GetApTempThresholdsOk() (*[]int32, bool)`

GetApTempThresholdsOk returns a tuple with the ApTempThresholds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApTempThresholds

`func (o *DeviceObjectDTO) SetApTempThresholds(v []int32)`

SetApTempThresholds sets ApTempThresholds field to given value.

### HasApTempThresholds

`func (o *DeviceObjectDTO) HasApTempThresholds() bool`

HasApTempThresholds returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *DeviceObjectDTO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *DeviceObjectDTO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *DeviceObjectDTO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *DeviceObjectDTO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetFilledTempThreshold

`func (o *DeviceObjectDTO) GetFilledTempThreshold() bool`

GetFilledTempThreshold returns the FilledTempThreshold field if non-nil, zero value otherwise.

### GetFilledTempThresholdOk

`func (o *DeviceObjectDTO) GetFilledTempThresholdOk() (*bool, bool)`

GetFilledTempThresholdOk returns a tuple with the FilledTempThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledTempThreshold

`func (o *DeviceObjectDTO) SetFilledTempThreshold(v bool)`

SetFilledTempThreshold sets FilledTempThreshold field to given value.

### HasFilledTempThreshold

`func (o *DeviceObjectDTO) HasFilledTempThreshold() bool`

HasFilledTempThreshold returns a boolean if a field has been set.

### GetIp

`func (o *DeviceObjectDTO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *DeviceObjectDTO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *DeviceObjectDTO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *DeviceObjectDTO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetLagPortsMap

`func (o *DeviceObjectDTO) GetLagPortsMap() map[string][]int32`

GetLagPortsMap returns the LagPortsMap field if non-nil, zero value otherwise.

### GetLagPortsMapOk

`func (o *DeviceObjectDTO) GetLagPortsMapOk() (*map[string][]int32, bool)`

GetLagPortsMapOk returns a tuple with the LagPortsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagPortsMap

`func (o *DeviceObjectDTO) SetLagPortsMap(v map[string][]int32)`

SetLagPortsMap sets LagPortsMap field to given value.

### HasLagPortsMap

`func (o *DeviceObjectDTO) HasLagPortsMap() bool`

HasLagPortsMap returns a boolean if a field has been set.

### GetModel

`func (o *DeviceObjectDTO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DeviceObjectDTO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DeviceObjectDTO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DeviceObjectDTO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DeviceObjectDTO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DeviceObjectDTO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DeviceObjectDTO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DeviceObjectDTO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DeviceObjectDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceObjectDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceObjectDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceObjectDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsgTempThreshold

`func (o *DeviceObjectDTO) GetOsgTempThreshold() int32`

GetOsgTempThreshold returns the OsgTempThreshold field if non-nil, zero value otherwise.

### GetOsgTempThresholdOk

`func (o *DeviceObjectDTO) GetOsgTempThresholdOk() (*int32, bool)`

GetOsgTempThresholdOk returns a tuple with the OsgTempThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsgTempThreshold

`func (o *DeviceObjectDTO) SetOsgTempThreshold(v int32)`

SetOsgTempThreshold sets OsgTempThreshold field to given value.

### HasOsgTempThreshold

`func (o *DeviceObjectDTO) HasOsgTempThreshold() bool`

HasOsgTempThreshold returns a boolean if a field has been set.

### GetOswCpuTempThreshold

`func (o *DeviceObjectDTO) GetOswCpuTempThreshold() int32`

GetOswCpuTempThreshold returns the OswCpuTempThreshold field if non-nil, zero value otherwise.

### GetOswCpuTempThresholdOk

`func (o *DeviceObjectDTO) GetOswCpuTempThresholdOk() (*int32, bool)`

GetOswCpuTempThresholdOk returns a tuple with the OswCpuTempThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswCpuTempThreshold

`func (o *DeviceObjectDTO) SetOswCpuTempThreshold(v int32)`

SetOswCpuTempThreshold sets OswCpuTempThreshold field to given value.

### HasOswCpuTempThreshold

`func (o *DeviceObjectDTO) HasOswCpuTempThreshold() bool`

HasOswCpuTempThreshold returns a boolean if a field has been set.

### GetOswMacTempThreshold

`func (o *DeviceObjectDTO) GetOswMacTempThreshold() int32`

GetOswMacTempThreshold returns the OswMacTempThreshold field if non-nil, zero value otherwise.

### GetOswMacTempThresholdOk

`func (o *DeviceObjectDTO) GetOswMacTempThresholdOk() (*int32, bool)`

GetOswMacTempThresholdOk returns a tuple with the OswMacTempThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswMacTempThreshold

`func (o *DeviceObjectDTO) SetOswMacTempThreshold(v int32)`

SetOswMacTempThreshold sets OswMacTempThreshold field to given value.

### HasOswMacTempThreshold

`func (o *DeviceObjectDTO) HasOswMacTempThreshold() bool`

HasOswMacTempThreshold returns a boolean if a field has been set.

### GetOswPhyTempThreshold

`func (o *DeviceObjectDTO) GetOswPhyTempThreshold() int32`

GetOswPhyTempThreshold returns the OswPhyTempThreshold field if non-nil, zero value otherwise.

### GetOswPhyTempThresholdOk

`func (o *DeviceObjectDTO) GetOswPhyTempThresholdOk() (*int32, bool)`

GetOswPhyTempThresholdOk returns a tuple with the OswPhyTempThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswPhyTempThreshold

`func (o *DeviceObjectDTO) SetOswPhyTempThreshold(v int32)`

SetOswPhyTempThreshold sets OswPhyTempThreshold field to given value.

### HasOswPhyTempThreshold

`func (o *DeviceObjectDTO) HasOswPhyTempThreshold() bool`

HasOswPhyTempThreshold returns a boolean if a field has been set.

### GetOswPseTempThreshold

`func (o *DeviceObjectDTO) GetOswPseTempThreshold() int32`

GetOswPseTempThreshold returns the OswPseTempThreshold field if non-nil, zero value otherwise.

### GetOswPseTempThresholdOk

`func (o *DeviceObjectDTO) GetOswPseTempThresholdOk() (*int32, bool)`

GetOswPseTempThresholdOk returns a tuple with the OswPseTempThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswPseTempThreshold

`func (o *DeviceObjectDTO) SetOswPseTempThreshold(v int32)`

SetOswPseTempThreshold sets OswPseTempThreshold field to given value.

### HasOswPseTempThreshold

`func (o *DeviceObjectDTO) HasOswPseTempThreshold() bool`

HasOswPseTempThreshold returns a boolean if a field has been set.

### GetPortNum

`func (o *DeviceObjectDTO) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *DeviceObjectDTO) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *DeviceObjectDTO) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *DeviceObjectDTO) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.

### GetRadio2gEnable

`func (o *DeviceObjectDTO) GetRadio2gEnable() bool`

GetRadio2gEnable returns the Radio2gEnable field if non-nil, zero value otherwise.

### GetRadio2gEnableOk

`func (o *DeviceObjectDTO) GetRadio2gEnableOk() (*bool, bool)`

GetRadio2gEnableOk returns a tuple with the Radio2gEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio2gEnable

`func (o *DeviceObjectDTO) SetRadio2gEnable(v bool)`

SetRadio2gEnable sets Radio2gEnable field to given value.

### HasRadio2gEnable

`func (o *DeviceObjectDTO) HasRadio2gEnable() bool`

HasRadio2gEnable returns a boolean if a field has been set.

### GetRadio5g2Enable

`func (o *DeviceObjectDTO) GetRadio5g2Enable() bool`

GetRadio5g2Enable returns the Radio5g2Enable field if non-nil, zero value otherwise.

### GetRadio5g2EnableOk

`func (o *DeviceObjectDTO) GetRadio5g2EnableOk() (*bool, bool)`

GetRadio5g2EnableOk returns a tuple with the Radio5g2Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio5g2Enable

`func (o *DeviceObjectDTO) SetRadio5g2Enable(v bool)`

SetRadio5g2Enable sets Radio5g2Enable field to given value.

### HasRadio5g2Enable

`func (o *DeviceObjectDTO) HasRadio5g2Enable() bool`

HasRadio5g2Enable returns a boolean if a field has been set.

### GetRadio5gEnable

`func (o *DeviceObjectDTO) GetRadio5gEnable() bool`

GetRadio5gEnable returns the Radio5gEnable field if non-nil, zero value otherwise.

### GetRadio5gEnableOk

`func (o *DeviceObjectDTO) GetRadio5gEnableOk() (*bool, bool)`

GetRadio5gEnableOk returns a tuple with the Radio5gEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio5gEnable

`func (o *DeviceObjectDTO) SetRadio5gEnable(v bool)`

SetRadio5gEnable sets Radio5gEnable field to given value.

### HasRadio5gEnable

`func (o *DeviceObjectDTO) HasRadio5gEnable() bool`

HasRadio5gEnable returns a boolean if a field has been set.

### GetRadio6gEnable

`func (o *DeviceObjectDTO) GetRadio6gEnable() bool`

GetRadio6gEnable returns the Radio6gEnable field if non-nil, zero value otherwise.

### GetRadio6gEnableOk

`func (o *DeviceObjectDTO) GetRadio6gEnableOk() (*bool, bool)`

GetRadio6gEnableOk returns a tuple with the Radio6gEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadio6gEnable

`func (o *DeviceObjectDTO) SetRadio6gEnable(v bool)`

SetRadio6gEnable sets Radio6gEnable field to given value.

### HasRadio6gEnable

`func (o *DeviceObjectDTO) HasRadio6gEnable() bool`

HasRadio6gEnable returns a boolean if a field has been set.

### GetSupport5g

`func (o *DeviceObjectDTO) GetSupport5g() bool`

GetSupport5g returns the Support5g field if non-nil, zero value otherwise.

### GetSupport5gOk

`func (o *DeviceObjectDTO) GetSupport5gOk() (*bool, bool)`

GetSupport5gOk returns a tuple with the Support5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g

`func (o *DeviceObjectDTO) SetSupport5g(v bool)`

SetSupport5g sets Support5g field to given value.

### HasSupport5g

`func (o *DeviceObjectDTO) HasSupport5g() bool`

HasSupport5g returns a boolean if a field has been set.

### GetSupport5g2

`func (o *DeviceObjectDTO) GetSupport5g2() bool`

GetSupport5g2 returns the Support5g2 field if non-nil, zero value otherwise.

### GetSupport5g2Ok

`func (o *DeviceObjectDTO) GetSupport5g2Ok() (*bool, bool)`

GetSupport5g2Ok returns a tuple with the Support5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g2

`func (o *DeviceObjectDTO) SetSupport5g2(v bool)`

SetSupport5g2 sets Support5g2 field to given value.

### HasSupport5g2

`func (o *DeviceObjectDTO) HasSupport5g2() bool`

HasSupport5g2 returns a boolean if a field has been set.

### GetSupport6g

`func (o *DeviceObjectDTO) GetSupport6g() bool`

GetSupport6g returns the Support6g field if non-nil, zero value otherwise.

### GetSupport6gOk

`func (o *DeviceObjectDTO) GetSupport6gOk() (*bool, bool)`

GetSupport6gOk returns a tuple with the Support6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport6g

`func (o *DeviceObjectDTO) SetSupport6g(v bool)`

SetSupport6g sets Support6g field to given value.

### HasSupport6g

`func (o *DeviceObjectDTO) HasSupport6g() bool`

HasSupport6g returns a boolean if a field has been set.

### GetSupportAnomaly

`func (o *DeviceObjectDTO) GetSupportAnomaly() bool`

GetSupportAnomaly returns the SupportAnomaly field if non-nil, zero value otherwise.

### GetSupportAnomalyOk

`func (o *DeviceObjectDTO) GetSupportAnomalyOk() (*bool, bool)`

GetSupportAnomalyOk returns a tuple with the SupportAnomaly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportAnomaly

`func (o *DeviceObjectDTO) SetSupportAnomaly(v bool)`

SetSupportAnomaly sets SupportAnomaly field to given value.

### HasSupportAnomaly

`func (o *DeviceObjectDTO) HasSupportAnomaly() bool`

HasSupportAnomaly returns a boolean if a field has been set.

### GetType

`func (o *DeviceObjectDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceObjectDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceObjectDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceObjectDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


