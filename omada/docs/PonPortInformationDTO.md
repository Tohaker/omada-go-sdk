# PonPortInformationDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualBandwidth** | Pointer to **int32** | Actual usable bandwidth of the PON port: 0 to 1310720 kbps. | [optional] 
**MaxBandwidth** | Pointer to **int32** | Maximum usable bandwidth of the PON port: 0 to 1310720 kbps. | [optional] 
**OnuNum** | Pointer to **int32** | The number of successfully registered ONU devices on this port, onuNum should be within the range of 0 to 128. | [optional] 
**OpticalBias** | Pointer to **string** | PON port optical current, in mA. | [optional] 
**OpticalPower** | Pointer to **string** | PON port power, in dBm. | [optional] 
**OpticalVcc** | Pointer to **string** | PON port optical voltage, in V. | [optional] 
**PortId** | Pointer to **string** | OLT port id | [optional] 
**RemainBandwidth** | Pointer to **int32** | Remaining usable bandwidth of the PON port: 0 to 1310720 kbps. | [optional] 
**Status** | Pointer to **string** | Whether to enable the laser for the PON port.Status should be a value as follows:DISABLE,ENABLE. | [optional] 

## Methods

### NewPonPortInformationDTO

`func NewPonPortInformationDTO() *PonPortInformationDTO`

NewPonPortInformationDTO instantiates a new PonPortInformationDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPonPortInformationDTOWithDefaults

`func NewPonPortInformationDTOWithDefaults() *PonPortInformationDTO`

NewPonPortInformationDTOWithDefaults instantiates a new PonPortInformationDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualBandwidth

`func (o *PonPortInformationDTO) GetActualBandwidth() int32`

GetActualBandwidth returns the ActualBandwidth field if non-nil, zero value otherwise.

### GetActualBandwidthOk

`func (o *PonPortInformationDTO) GetActualBandwidthOk() (*int32, bool)`

GetActualBandwidthOk returns a tuple with the ActualBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualBandwidth

`func (o *PonPortInformationDTO) SetActualBandwidth(v int32)`

SetActualBandwidth sets ActualBandwidth field to given value.

### HasActualBandwidth

`func (o *PonPortInformationDTO) HasActualBandwidth() bool`

HasActualBandwidth returns a boolean if a field has been set.

### GetMaxBandwidth

`func (o *PonPortInformationDTO) GetMaxBandwidth() int32`

GetMaxBandwidth returns the MaxBandwidth field if non-nil, zero value otherwise.

### GetMaxBandwidthOk

`func (o *PonPortInformationDTO) GetMaxBandwidthOk() (*int32, bool)`

GetMaxBandwidthOk returns a tuple with the MaxBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBandwidth

`func (o *PonPortInformationDTO) SetMaxBandwidth(v int32)`

SetMaxBandwidth sets MaxBandwidth field to given value.

### HasMaxBandwidth

`func (o *PonPortInformationDTO) HasMaxBandwidth() bool`

HasMaxBandwidth returns a boolean if a field has been set.

### GetOnuNum

`func (o *PonPortInformationDTO) GetOnuNum() int32`

GetOnuNum returns the OnuNum field if non-nil, zero value otherwise.

### GetOnuNumOk

`func (o *PonPortInformationDTO) GetOnuNumOk() (*int32, bool)`

GetOnuNumOk returns a tuple with the OnuNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnuNum

`func (o *PonPortInformationDTO) SetOnuNum(v int32)`

SetOnuNum sets OnuNum field to given value.

### HasOnuNum

`func (o *PonPortInformationDTO) HasOnuNum() bool`

HasOnuNum returns a boolean if a field has been set.

### GetOpticalBias

`func (o *PonPortInformationDTO) GetOpticalBias() string`

GetOpticalBias returns the OpticalBias field if non-nil, zero value otherwise.

### GetOpticalBiasOk

`func (o *PonPortInformationDTO) GetOpticalBiasOk() (*string, bool)`

GetOpticalBiasOk returns a tuple with the OpticalBias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalBias

`func (o *PonPortInformationDTO) SetOpticalBias(v string)`

SetOpticalBias sets OpticalBias field to given value.

### HasOpticalBias

`func (o *PonPortInformationDTO) HasOpticalBias() bool`

HasOpticalBias returns a boolean if a field has been set.

### GetOpticalPower

`func (o *PonPortInformationDTO) GetOpticalPower() string`

GetOpticalPower returns the OpticalPower field if non-nil, zero value otherwise.

### GetOpticalPowerOk

`func (o *PonPortInformationDTO) GetOpticalPowerOk() (*string, bool)`

GetOpticalPowerOk returns a tuple with the OpticalPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalPower

`func (o *PonPortInformationDTO) SetOpticalPower(v string)`

SetOpticalPower sets OpticalPower field to given value.

### HasOpticalPower

`func (o *PonPortInformationDTO) HasOpticalPower() bool`

HasOpticalPower returns a boolean if a field has been set.

### GetOpticalVcc

`func (o *PonPortInformationDTO) GetOpticalVcc() string`

GetOpticalVcc returns the OpticalVcc field if non-nil, zero value otherwise.

### GetOpticalVccOk

`func (o *PonPortInformationDTO) GetOpticalVccOk() (*string, bool)`

GetOpticalVccOk returns a tuple with the OpticalVcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpticalVcc

`func (o *PonPortInformationDTO) SetOpticalVcc(v string)`

SetOpticalVcc sets OpticalVcc field to given value.

### HasOpticalVcc

`func (o *PonPortInformationDTO) HasOpticalVcc() bool`

HasOpticalVcc returns a boolean if a field has been set.

### GetPortId

`func (o *PonPortInformationDTO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *PonPortInformationDTO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *PonPortInformationDTO) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *PonPortInformationDTO) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetRemainBandwidth

`func (o *PonPortInformationDTO) GetRemainBandwidth() int32`

GetRemainBandwidth returns the RemainBandwidth field if non-nil, zero value otherwise.

### GetRemainBandwidthOk

`func (o *PonPortInformationDTO) GetRemainBandwidthOk() (*int32, bool)`

GetRemainBandwidthOk returns a tuple with the RemainBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainBandwidth

`func (o *PonPortInformationDTO) SetRemainBandwidth(v int32)`

SetRemainBandwidth sets RemainBandwidth field to given value.

### HasRemainBandwidth

`func (o *PonPortInformationDTO) HasRemainBandwidth() bool`

HasRemainBandwidth returns a boolean if a field has been set.

### GetStatus

`func (o *PonPortInformationDTO) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PonPortInformationDTO) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PonPortInformationDTO) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PonPortInformationDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


