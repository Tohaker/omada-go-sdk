# OswStatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortMap** | Pointer to [**map[string]PortStatVO**](PortStatVO.md) | Port total traffic map | [optional] 
**StatList** | Pointer to [**[]OswStatDetailOpenApiVO**](OswStatDetailOpenApiVO.md) | Detailed traffic information of ports | [optional] 
**SwitchType** | Pointer to **int32** | Type of Switch should be a value as follows: 0:Non-Agile Series Switch, 1:Agile Series Switch | [optional] 

## Methods

### NewOswStatOpenApiVO

`func NewOswStatOpenApiVO() *OswStatOpenApiVO`

NewOswStatOpenApiVO instantiates a new OswStatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatOpenApiVOWithDefaults

`func NewOswStatOpenApiVOWithDefaults() *OswStatOpenApiVO`

NewOswStatOpenApiVOWithDefaults instantiates a new OswStatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortMap

`func (o *OswStatOpenApiVO) GetPortMap() map[string]PortStatVO`

GetPortMap returns the PortMap field if non-nil, zero value otherwise.

### GetPortMapOk

`func (o *OswStatOpenApiVO) GetPortMapOk() (*map[string]PortStatVO, bool)`

GetPortMapOk returns a tuple with the PortMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMap

`func (o *OswStatOpenApiVO) SetPortMap(v map[string]PortStatVO)`

SetPortMap sets PortMap field to given value.

### HasPortMap

`func (o *OswStatOpenApiVO) HasPortMap() bool`

HasPortMap returns a boolean if a field has been set.

### GetStatList

`func (o *OswStatOpenApiVO) GetStatList() []OswStatDetailOpenApiVO`

GetStatList returns the StatList field if non-nil, zero value otherwise.

### GetStatListOk

`func (o *OswStatOpenApiVO) GetStatListOk() (*[]OswStatDetailOpenApiVO, bool)`

GetStatListOk returns a tuple with the StatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatList

`func (o *OswStatOpenApiVO) SetStatList(v []OswStatDetailOpenApiVO)`

SetStatList sets StatList field to given value.

### HasStatList

`func (o *OswStatOpenApiVO) HasStatList() bool`

HasStatList returns a boolean if a field has been set.

### GetSwitchType

`func (o *OswStatOpenApiVO) GetSwitchType() int32`

GetSwitchType returns the SwitchType field if non-nil, zero value otherwise.

### GetSwitchTypeOk

`func (o *OswStatOpenApiVO) GetSwitchTypeOk() (*int32, bool)`

GetSwitchTypeOk returns a tuple with the SwitchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchType

`func (o *OswStatOpenApiVO) SetSwitchType(v int32)`

SetSwitchType sets SwitchType field to given value.

### HasSwitchType

`func (o *OswStatOpenApiVO) HasSwitchType() bool`

HasSwitchType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


