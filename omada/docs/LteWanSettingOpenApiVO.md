# LteWanSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LteWanPortsConfig** | Pointer to [**[]LteWanPortSettingOpenApiVO**](LteWanPortSettingOpenApiVO.md) | LTE WAN ports config | [optional] 
**PortDesc** | Pointer to **string** | LTE WAN port description | [optional] 
**PortName** | Pointer to **string** | LTE WAN port name | [optional] 
**PortUuid** | **string** | LTE WAN port uuid | 

## Methods

### NewLteWanSettingOpenApiVO

`func NewLteWanSettingOpenApiVO(portUuid string, ) *LteWanSettingOpenApiVO`

NewLteWanSettingOpenApiVO instantiates a new LteWanSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLteWanSettingOpenApiVOWithDefaults

`func NewLteWanSettingOpenApiVOWithDefaults() *LteWanSettingOpenApiVO`

NewLteWanSettingOpenApiVOWithDefaults instantiates a new LteWanSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLteWanPortsConfig

`func (o *LteWanSettingOpenApiVO) GetLteWanPortsConfig() []LteWanPortSettingOpenApiVO`

GetLteWanPortsConfig returns the LteWanPortsConfig field if non-nil, zero value otherwise.

### GetLteWanPortsConfigOk

`func (o *LteWanSettingOpenApiVO) GetLteWanPortsConfigOk() (*[]LteWanPortSettingOpenApiVO, bool)`

GetLteWanPortsConfigOk returns a tuple with the LteWanPortsConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLteWanPortsConfig

`func (o *LteWanSettingOpenApiVO) SetLteWanPortsConfig(v []LteWanPortSettingOpenApiVO)`

SetLteWanPortsConfig sets LteWanPortsConfig field to given value.

### HasLteWanPortsConfig

`func (o *LteWanSettingOpenApiVO) HasLteWanPortsConfig() bool`

HasLteWanPortsConfig returns a boolean if a field has been set.

### GetPortDesc

`func (o *LteWanSettingOpenApiVO) GetPortDesc() string`

GetPortDesc returns the PortDesc field if non-nil, zero value otherwise.

### GetPortDescOk

`func (o *LteWanSettingOpenApiVO) GetPortDescOk() (*string, bool)`

GetPortDescOk returns a tuple with the PortDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortDesc

`func (o *LteWanSettingOpenApiVO) SetPortDesc(v string)`

SetPortDesc sets PortDesc field to given value.

### HasPortDesc

`func (o *LteWanSettingOpenApiVO) HasPortDesc() bool`

HasPortDesc returns a boolean if a field has been set.

### GetPortName

`func (o *LteWanSettingOpenApiVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *LteWanSettingOpenApiVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *LteWanSettingOpenApiVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *LteWanSettingOpenApiVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortUuid

`func (o *LteWanSettingOpenApiVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *LteWanSettingOpenApiVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *LteWanSettingOpenApiVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


