# OswStackMemberCableTestVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | MAC | [optional] 
**Model** | Pointer to **string** | Device Model. | [optional] 
**ModelVersion** | Pointer to **string** | Device Model Version. | [optional] 
**PortList** | Pointer to [**[]OswCableTestPortVO**](OswCableTestPortVO.md) | Port list | [optional] 
**Unit** | Pointer to **int32** | Unit | [optional] 

## Methods

### NewOswStackMemberCableTestVO

`func NewOswStackMemberCableTestVO() *OswStackMemberCableTestVO`

NewOswStackMemberCableTestVO instantiates a new OswStackMemberCableTestVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberCableTestVOWithDefaults

`func NewOswStackMemberCableTestVOWithDefaults() *OswStackMemberCableTestVO`

NewOswStackMemberCableTestVOWithDefaults instantiates a new OswStackMemberCableTestVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswStackMemberCableTestVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackMemberCableTestVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackMemberCableTestVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStackMemberCableTestVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *OswStackMemberCableTestVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackMemberCableTestVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackMemberCableTestVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackMemberCableTestVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackMemberCableTestVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackMemberCableTestVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackMemberCableTestVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackMemberCableTestVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetPortList

`func (o *OswStackMemberCableTestVO) GetPortList() []OswCableTestPortVO`

GetPortList returns the PortList field if non-nil, zero value otherwise.

### GetPortListOk

`func (o *OswStackMemberCableTestVO) GetPortListOk() (*[]OswCableTestPortVO, bool)`

GetPortListOk returns a tuple with the PortList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortList

`func (o *OswStackMemberCableTestVO) SetPortList(v []OswCableTestPortVO)`

SetPortList sets PortList field to given value.

### HasPortList

`func (o *OswStackMemberCableTestVO) HasPortList() bool`

HasPortList returns a boolean if a field has been set.

### GetUnit

`func (o *OswStackMemberCableTestVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackMemberCableTestVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackMemberCableTestVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswStackMemberCableTestVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


