# OswStackPortGroupVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ID** | **int32** | The ID of the stack port aggregation group | 
**GroupSpeed** | Pointer to **int32** | Stack port aggregation group link speed config | [optional] 
**Name** | Pointer to **string** | Name of the stacking port aggregation group | [optional] 
**Ports** | Pointer to [**[]OswStandPortVO**](OswStandPortVO.md) | Stack port aggregation group member port | [optional] 

## Methods

### NewOswStackPortGroupVO

`func NewOswStackPortGroupVO(iD int32, ) *OswStackPortGroupVO`

NewOswStackPortGroupVO instantiates a new OswStackPortGroupVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackPortGroupVOWithDefaults

`func NewOswStackPortGroupVOWithDefaults() *OswStackPortGroupVO`

NewOswStackPortGroupVOWithDefaults instantiates a new OswStackPortGroupVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetID

`func (o *OswStackPortGroupVO) GetID() int32`

GetID returns the ID field if non-nil, zero value otherwise.

### GetIDOk

`func (o *OswStackPortGroupVO) GetIDOk() (*int32, bool)`

GetIDOk returns a tuple with the ID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetID

`func (o *OswStackPortGroupVO) SetID(v int32)`

SetID sets ID field to given value.


### GetGroupSpeed

`func (o *OswStackPortGroupVO) GetGroupSpeed() int32`

GetGroupSpeed returns the GroupSpeed field if non-nil, zero value otherwise.

### GetGroupSpeedOk

`func (o *OswStackPortGroupVO) GetGroupSpeedOk() (*int32, bool)`

GetGroupSpeedOk returns a tuple with the GroupSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSpeed

`func (o *OswStackPortGroupVO) SetGroupSpeed(v int32)`

SetGroupSpeed sets GroupSpeed field to given value.

### HasGroupSpeed

`func (o *OswStackPortGroupVO) HasGroupSpeed() bool`

HasGroupSpeed returns a boolean if a field has been set.

### GetName

`func (o *OswStackPortGroupVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackPortGroupVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackPortGroupVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStackPortGroupVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPorts

`func (o *OswStackPortGroupVO) GetPorts() []OswStandPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStackPortGroupVO) GetPortsOk() (*[]OswStandPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStackPortGroupVO) SetPorts(v []OswStandPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStackPortGroupVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


