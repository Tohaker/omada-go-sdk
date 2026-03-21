# DBAProfileModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assure** | Pointer to **int32** | The value for guaranteed bandwidth is required when the Type is set to ASSURE, ASSURE_MAX, or FIX_ASSURE_MAX. Assure should be within the range of 128 to 1200000, with the unit in kbit/s. | [optional] 
**DbaId** | **int32** | DBA ID should be within the range of 1 to 512 and should not be null | 
**Fix** | Pointer to **int32** | The value for fixed bandwidth is required when the Type is set to FIX or FIX_ASSURE_MAX. Fix should be within the range of 128 to 720000, with the unit in kbit/s. | [optional] 
**Max** | Pointer to **int32** | The value for maximum bandwidth is required when the Type is set to MAX, ASSURE_MAX, or FIX_ASSURE_MAX. Max should be within the range of 128 to 1244160, with the unit in kbit/s. | [optional] 
**Name** | Pointer to **string** | Name of DBA profile should contain 1 to 32 characters including digits, upper and lower letters, and the following six characters: -@_:/. . | [optional] 
**Type** | Pointer to **string** | DBA bandwidth allocation methods.Type should be a value as follows: FIX, ASSURE, MAX, ASSURE_MAX, and FIX_ASSURE_MAX. The default value is FIX. | [optional] 

## Methods

### NewDBAProfileModifyDTO

`func NewDBAProfileModifyDTO(dbaId int32, ) *DBAProfileModifyDTO`

NewDBAProfileModifyDTO instantiates a new DBAProfileModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDBAProfileModifyDTOWithDefaults

`func NewDBAProfileModifyDTOWithDefaults() *DBAProfileModifyDTO`

NewDBAProfileModifyDTOWithDefaults instantiates a new DBAProfileModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssure

`func (o *DBAProfileModifyDTO) GetAssure() int32`

GetAssure returns the Assure field if non-nil, zero value otherwise.

### GetAssureOk

`func (o *DBAProfileModifyDTO) GetAssureOk() (*int32, bool)`

GetAssureOk returns a tuple with the Assure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssure

`func (o *DBAProfileModifyDTO) SetAssure(v int32)`

SetAssure sets Assure field to given value.

### HasAssure

`func (o *DBAProfileModifyDTO) HasAssure() bool`

HasAssure returns a boolean if a field has been set.

### GetDbaId

`func (o *DBAProfileModifyDTO) GetDbaId() int32`

GetDbaId returns the DbaId field if non-nil, zero value otherwise.

### GetDbaIdOk

`func (o *DBAProfileModifyDTO) GetDbaIdOk() (*int32, bool)`

GetDbaIdOk returns a tuple with the DbaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDbaId

`func (o *DBAProfileModifyDTO) SetDbaId(v int32)`

SetDbaId sets DbaId field to given value.


### GetFix

`func (o *DBAProfileModifyDTO) GetFix() int32`

GetFix returns the Fix field if non-nil, zero value otherwise.

### GetFixOk

`func (o *DBAProfileModifyDTO) GetFixOk() (*int32, bool)`

GetFixOk returns a tuple with the Fix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFix

`func (o *DBAProfileModifyDTO) SetFix(v int32)`

SetFix sets Fix field to given value.

### HasFix

`func (o *DBAProfileModifyDTO) HasFix() bool`

HasFix returns a boolean if a field has been set.

### GetMax

`func (o *DBAProfileModifyDTO) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *DBAProfileModifyDTO) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *DBAProfileModifyDTO) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *DBAProfileModifyDTO) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetName

`func (o *DBAProfileModifyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DBAProfileModifyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DBAProfileModifyDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DBAProfileModifyDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *DBAProfileModifyDTO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DBAProfileModifyDTO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DBAProfileModifyDTO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DBAProfileModifyDTO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


