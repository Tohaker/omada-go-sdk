# NewMeshSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutoFailoverEnable** | Pointer to **bool** | Whether to enable auto failover | [optional] 
**DefGatewayEnable** | Pointer to **bool** | Connectivity detection, parameter defGatewayEnable should be a value as follows: true: Auto(Recommended) ; false: Customer IP | [optional] 
**FullSector** | Pointer to **bool** | Whether to enable full-sector DFS | [optional] 
**Gateway** | Pointer to **string** | Customer IP | [optional] 
**MeshEnable** | **bool** | Whether to enable mesh | 

## Methods

### NewNewMeshSettingOpenApiVO

`func NewNewMeshSettingOpenApiVO(meshEnable bool, ) *NewMeshSettingOpenApiVO`

NewNewMeshSettingOpenApiVO instantiates a new NewMeshSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewMeshSettingOpenApiVOWithDefaults

`func NewNewMeshSettingOpenApiVOWithDefaults() *NewMeshSettingOpenApiVO`

NewNewMeshSettingOpenApiVOWithDefaults instantiates a new NewMeshSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoFailoverEnable

`func (o *NewMeshSettingOpenApiVO) GetAutoFailoverEnable() bool`

GetAutoFailoverEnable returns the AutoFailoverEnable field if non-nil, zero value otherwise.

### GetAutoFailoverEnableOk

`func (o *NewMeshSettingOpenApiVO) GetAutoFailoverEnableOk() (*bool, bool)`

GetAutoFailoverEnableOk returns a tuple with the AutoFailoverEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoFailoverEnable

`func (o *NewMeshSettingOpenApiVO) SetAutoFailoverEnable(v bool)`

SetAutoFailoverEnable sets AutoFailoverEnable field to given value.

### HasAutoFailoverEnable

`func (o *NewMeshSettingOpenApiVO) HasAutoFailoverEnable() bool`

HasAutoFailoverEnable returns a boolean if a field has been set.

### GetDefGatewayEnable

`func (o *NewMeshSettingOpenApiVO) GetDefGatewayEnable() bool`

GetDefGatewayEnable returns the DefGatewayEnable field if non-nil, zero value otherwise.

### GetDefGatewayEnableOk

`func (o *NewMeshSettingOpenApiVO) GetDefGatewayEnableOk() (*bool, bool)`

GetDefGatewayEnableOk returns a tuple with the DefGatewayEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefGatewayEnable

`func (o *NewMeshSettingOpenApiVO) SetDefGatewayEnable(v bool)`

SetDefGatewayEnable sets DefGatewayEnable field to given value.

### HasDefGatewayEnable

`func (o *NewMeshSettingOpenApiVO) HasDefGatewayEnable() bool`

HasDefGatewayEnable returns a boolean if a field has been set.

### GetFullSector

`func (o *NewMeshSettingOpenApiVO) GetFullSector() bool`

GetFullSector returns the FullSector field if non-nil, zero value otherwise.

### GetFullSectorOk

`func (o *NewMeshSettingOpenApiVO) GetFullSectorOk() (*bool, bool)`

GetFullSectorOk returns a tuple with the FullSector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullSector

`func (o *NewMeshSettingOpenApiVO) SetFullSector(v bool)`

SetFullSector sets FullSector field to given value.

### HasFullSector

`func (o *NewMeshSettingOpenApiVO) HasFullSector() bool`

HasFullSector returns a boolean if a field has been set.

### GetGateway

`func (o *NewMeshSettingOpenApiVO) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *NewMeshSettingOpenApiVO) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *NewMeshSettingOpenApiVO) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *NewMeshSettingOpenApiVO) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetMeshEnable

`func (o *NewMeshSettingOpenApiVO) GetMeshEnable() bool`

GetMeshEnable returns the MeshEnable field if non-nil, zero value otherwise.

### GetMeshEnableOk

`func (o *NewMeshSettingOpenApiVO) GetMeshEnableOk() (*bool, bool)`

GetMeshEnableOk returns a tuple with the MeshEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeshEnable

`func (o *NewMeshSettingOpenApiVO) SetMeshEnable(v bool)`

SetMeshEnable sets MeshEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


