# ModelLatestFwInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID | [optional] 
**LatestVersion** | Pointer to **string** | Latest version | [optional] 
**ModelTypeInfo** | Pointer to [**ModelTypeInfoOpenApiVO**](ModelTypeInfoOpenApiVO.md) |  | [optional] 
**ReleaseTime** | Pointer to **int64** | Latest version release time, timestamp ms | [optional] 

## Methods

### NewModelLatestFwInfo

`func NewModelLatestFwInfo() *ModelLatestFwInfo`

NewModelLatestFwInfo instantiates a new ModelLatestFwInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelLatestFwInfoWithDefaults

`func NewModelLatestFwInfoWithDefaults() *ModelLatestFwInfo`

NewModelLatestFwInfoWithDefaults instantiates a new ModelLatestFwInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelLatestFwInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelLatestFwInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelLatestFwInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ModelLatestFwInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatestVersion

`func (o *ModelLatestFwInfo) GetLatestVersion() string`

GetLatestVersion returns the LatestVersion field if non-nil, zero value otherwise.

### GetLatestVersionOk

`func (o *ModelLatestFwInfo) GetLatestVersionOk() (*string, bool)`

GetLatestVersionOk returns a tuple with the LatestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestVersion

`func (o *ModelLatestFwInfo) SetLatestVersion(v string)`

SetLatestVersion sets LatestVersion field to given value.

### HasLatestVersion

`func (o *ModelLatestFwInfo) HasLatestVersion() bool`

HasLatestVersion returns a boolean if a field has been set.

### GetModelTypeInfo

`func (o *ModelLatestFwInfo) GetModelTypeInfo() ModelTypeInfoOpenApiVO`

GetModelTypeInfo returns the ModelTypeInfo field if non-nil, zero value otherwise.

### GetModelTypeInfoOk

`func (o *ModelLatestFwInfo) GetModelTypeInfoOk() (*ModelTypeInfoOpenApiVO, bool)`

GetModelTypeInfoOk returns a tuple with the ModelTypeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelTypeInfo

`func (o *ModelLatestFwInfo) SetModelTypeInfo(v ModelTypeInfoOpenApiVO)`

SetModelTypeInfo sets ModelTypeInfo field to given value.

### HasModelTypeInfo

`func (o *ModelLatestFwInfo) HasModelTypeInfo() bool`

HasModelTypeInfo returns a boolean if a field has been set.

### GetReleaseTime

`func (o *ModelLatestFwInfo) GetReleaseTime() int64`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *ModelLatestFwInfo) GetReleaseTimeOk() (*int64, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *ModelLatestFwInfo) SetReleaseTime(v int64)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *ModelLatestFwInfo) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


