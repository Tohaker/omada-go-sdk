# CliConfigTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CliConfig** | **string** | CLI configuration content | 
**CliType** | Pointer to **int32** | CLI configuration type, it should be a value as follows: 0：site CLI. | [optional] 
**Description** | Pointer to **string** | CLI configuration description, it should be within the range of 0 - 256 characters. | [optional] 
**DeviceType** | Pointer to **string** | Device type for CLI configuration application, it should be a value as follows: switch. | [optional] 
**Id** | Pointer to **string** | CLI configuration ID | [optional] 
**Name** | Pointer to **string** | CLI configuration name | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Status** | Pointer to **int32** | CLI configuration status, it should be a value as follows: 0: active, 1: inactive | [optional] 

## Methods

### NewCliConfigTemplateOpenApiVO

`func NewCliConfigTemplateOpenApiVO(cliConfig string, ) *CliConfigTemplateOpenApiVO`

NewCliConfigTemplateOpenApiVO instantiates a new CliConfigTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliConfigTemplateOpenApiVOWithDefaults

`func NewCliConfigTemplateOpenApiVOWithDefaults() *CliConfigTemplateOpenApiVO`

NewCliConfigTemplateOpenApiVOWithDefaults instantiates a new CliConfigTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCliConfig

`func (o *CliConfigTemplateOpenApiVO) GetCliConfig() string`

GetCliConfig returns the CliConfig field if non-nil, zero value otherwise.

### GetCliConfigOk

`func (o *CliConfigTemplateOpenApiVO) GetCliConfigOk() (*string, bool)`

GetCliConfigOk returns a tuple with the CliConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliConfig

`func (o *CliConfigTemplateOpenApiVO) SetCliConfig(v string)`

SetCliConfig sets CliConfig field to given value.


### GetCliType

`func (o *CliConfigTemplateOpenApiVO) GetCliType() int32`

GetCliType returns the CliType field if non-nil, zero value otherwise.

### GetCliTypeOk

`func (o *CliConfigTemplateOpenApiVO) GetCliTypeOk() (*int32, bool)`

GetCliTypeOk returns a tuple with the CliType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCliType

`func (o *CliConfigTemplateOpenApiVO) SetCliType(v int32)`

SetCliType sets CliType field to given value.

### HasCliType

`func (o *CliConfigTemplateOpenApiVO) HasCliType() bool`

HasCliType returns a boolean if a field has been set.

### GetDescription

`func (o *CliConfigTemplateOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CliConfigTemplateOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CliConfigTemplateOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CliConfigTemplateOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDeviceType

`func (o *CliConfigTemplateOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *CliConfigTemplateOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *CliConfigTemplateOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *CliConfigTemplateOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetId

`func (o *CliConfigTemplateOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CliConfigTemplateOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CliConfigTemplateOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CliConfigTemplateOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CliConfigTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CliConfigTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CliConfigTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CliConfigTemplateOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOmadacId

`func (o *CliConfigTemplateOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *CliConfigTemplateOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *CliConfigTemplateOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *CliConfigTemplateOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetSiteId

`func (o *CliConfigTemplateOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *CliConfigTemplateOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *CliConfigTemplateOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *CliConfigTemplateOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *CliConfigTemplateOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CliConfigTemplateOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CliConfigTemplateOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CliConfigTemplateOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


