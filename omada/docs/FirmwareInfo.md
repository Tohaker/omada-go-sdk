# FirmwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of firmware | [optional] 
**Id** | Pointer to **string** | ID | [optional] 
**ModelTypeInfo** | Pointer to [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | [optional] 
**Name** | Pointer to **string** | File name | [optional] 
**UploadTime** | Pointer to **int64** | Uploaded timestamp (ms) | [optional] 

## Methods

### NewFirmwareInfo

`func NewFirmwareInfo() *FirmwareInfo`

NewFirmwareInfo instantiates a new FirmwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareInfoWithDefaults

`func NewFirmwareInfoWithDefaults() *FirmwareInfo`

NewFirmwareInfoWithDefaults instantiates a new FirmwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *FirmwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *FirmwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *FirmwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *FirmwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *FirmwareInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirmwareInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *FirmwareInfo) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *FirmwareInfo) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *FirmwareInfo) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.

### HasModelTypeInfo

`func (o *FirmwareInfo) HasModelTypeInfo() bool`

HasModelTypeInfo returns a boolean if a field has been set.

### GetName

`func (o *FirmwareInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FirmwareInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FirmwareInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FirmwareInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUploadTime

`func (o *FirmwareInfo) GetUploadTime() int64`

GetUploadTime returns the UploadTime field if non-nil, zero value otherwise.

### GetUploadTimeOk

`func (o *FirmwareInfo) GetUploadTimeOk() (*int64, bool)`

GetUploadTimeOk returns a tuple with the UploadTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadTime

`func (o *FirmwareInfo) SetUploadTime(v int64)`

SetUploadTime sets UploadTime field to given value.

### HasUploadTime

`func (o *FirmwareInfo) HasUploadTime() bool`

HasUploadTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


