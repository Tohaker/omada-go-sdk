# OswStackMemberStatVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device mac | [optional] 
**PortMap** | Pointer to [**map[string]PortStatVO**](PortStatVO.md) | Port total traffic map | [optional] 
**StatList** | Pointer to [**[]OswStatDTO**](OswStatDTO.md) | Detailed traffic information of ports | [optional] 
**Unit** | Pointer to **int32** | Unit ID | [optional] 

## Methods

### NewOswStackMemberStatVO

`func NewOswStackMemberStatVO() *OswStackMemberStatVO`

NewOswStackMemberStatVO instantiates a new OswStackMemberStatVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberStatVOWithDefaults

`func NewOswStackMemberStatVOWithDefaults() *OswStackMemberStatVO`

NewOswStackMemberStatVOWithDefaults instantiates a new OswStackMemberStatVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswStackMemberStatVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackMemberStatVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackMemberStatVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStackMemberStatVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPortMap

`func (o *OswStackMemberStatVO) GetPortMap() map[string]PortStatVO`

GetPortMap returns the PortMap field if non-nil, zero value otherwise.

### GetPortMapOk

`func (o *OswStackMemberStatVO) GetPortMapOk() (*map[string]PortStatVO, bool)`

GetPortMapOk returns a tuple with the PortMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMap

`func (o *OswStackMemberStatVO) SetPortMap(v map[string]PortStatVO)`

SetPortMap sets PortMap field to given value.

### HasPortMap

`func (o *OswStackMemberStatVO) HasPortMap() bool`

HasPortMap returns a boolean if a field has been set.

### GetStatList

`func (o *OswStackMemberStatVO) GetStatList() []OswStatDTO`

GetStatList returns the StatList field if non-nil, zero value otherwise.

### GetStatListOk

`func (o *OswStackMemberStatVO) GetStatListOk() (*[]OswStatDTO, bool)`

GetStatListOk returns a tuple with the StatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatList

`func (o *OswStackMemberStatVO) SetStatList(v []OswStatDTO)`

SetStatList sets StatList field to given value.

### HasStatList

`func (o *OswStackMemberStatVO) HasStatList() bool`

HasStatList returns a boolean if a field has been set.

### GetUnit

`func (o *OswStackMemberStatVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackMemberStatVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackMemberStatVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswStackMemberStatVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


