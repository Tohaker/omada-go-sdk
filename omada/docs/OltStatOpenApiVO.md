# OltStatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortMap** | Pointer to [**map[string]PortStatVO**](PortStatVO.md) | Port total traffic map | [optional] 
**StatList** | Pointer to [**[]OltStatDetailOpenApiVO**](OltStatDetailOpenApiVO.md) | Detailed traffic information of ports | [optional] 

## Methods

### NewOltStatOpenApiVO

`func NewOltStatOpenApiVO() *OltStatOpenApiVO`

NewOltStatOpenApiVO instantiates a new OltStatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltStatOpenApiVOWithDefaults

`func NewOltStatOpenApiVOWithDefaults() *OltStatOpenApiVO`

NewOltStatOpenApiVOWithDefaults instantiates a new OltStatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortMap

`func (o *OltStatOpenApiVO) GetPortMap() map[string]PortStatVO`

GetPortMap returns the PortMap field if non-nil, zero value otherwise.

### GetPortMapOk

`func (o *OltStatOpenApiVO) GetPortMapOk() (*map[string]PortStatVO, bool)`

GetPortMapOk returns a tuple with the PortMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortMap

`func (o *OltStatOpenApiVO) SetPortMap(v map[string]PortStatVO)`

SetPortMap sets PortMap field to given value.

### HasPortMap

`func (o *OltStatOpenApiVO) HasPortMap() bool`

HasPortMap returns a boolean if a field has been set.

### GetStatList

`func (o *OltStatOpenApiVO) GetStatList() []OltStatDetailOpenApiVO`

GetStatList returns the StatList field if non-nil, zero value otherwise.

### GetStatListOk

`func (o *OltStatOpenApiVO) GetStatListOk() (*[]OltStatDetailOpenApiVO, bool)`

GetStatListOk returns a tuple with the StatList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatList

`func (o *OltStatOpenApiVO) SetStatList(v []OltStatDetailOpenApiVO)`

SetStatList sets StatList field to given value.

### HasStatList

`func (o *OltStatOpenApiVO) HasStatList() bool`

HasStatList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


