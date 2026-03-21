# OswPortStackSettingVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LagSetting** | Pointer to [**OswStackLagVO**](OswStackLagVO.md) |  | [optional] 
**MirroredPorts** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Mirrored Standard Ports | [optional] 
**StackId** | Pointer to **string** | Stack ID | [optional] 

## Methods

### NewOswPortStackSettingVO

`func NewOswPortStackSettingVO() *OswPortStackSettingVO`

NewOswPortStackSettingVO instantiates a new OswPortStackSettingVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortStackSettingVOWithDefaults

`func NewOswPortStackSettingVOWithDefaults() *OswPortStackSettingVO`

NewOswPortStackSettingVOWithDefaults instantiates a new OswPortStackSettingVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLagSetting

`func (o *OswPortStackSettingVO) GetLagSetting() OswStackLagVO`

GetLagSetting returns the LagSetting field if non-nil, zero value otherwise.

### GetLagSettingOk

`func (o *OswPortStackSettingVO) GetLagSettingOk() (*OswStackLagVO, bool)`

GetLagSettingOk returns a tuple with the LagSetting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLagSetting

`func (o *OswPortStackSettingVO) SetLagSetting(v OswStackLagVO)`

SetLagSetting sets LagSetting field to given value.

### HasLagSetting

`func (o *OswPortStackSettingVO) HasLagSetting() bool`

HasLagSetting returns a boolean if a field has been set.

### GetMirroredPorts

`func (o *OswPortStackSettingVO) GetMirroredPorts() []OswStandPortVO`

GetMirroredPorts returns the MirroredPorts field if non-nil, zero value otherwise.

### GetMirroredPortsOk

`func (o *OswPortStackSettingVO) GetMirroredPortsOk() (*[]OswStandPortVO, bool)`

GetMirroredPortsOk returns a tuple with the MirroredPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMirroredPorts

`func (o *OswPortStackSettingVO) SetMirroredPorts(v []OswStandPortVO)`

SetMirroredPorts sets MirroredPorts field to given value.

### HasMirroredPorts

`func (o *OswPortStackSettingVO) HasMirroredPorts() bool`

HasMirroredPorts returns a boolean if a field has been set.

### GetStackId

`func (o *OswPortStackSettingVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *OswPortStackSettingVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *OswPortStackSettingVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *OswPortStackSettingVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


