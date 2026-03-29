# SoftwareDetailConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Software0Active** | Pointer to **string** | Status of ONU software 0.Software0Active should be a value as follows:ACTIVE,INACTIVE | [optional] 
**Software0Commited** | Pointer to **string** | Commit status of ONU software 0.Software0Commited should be a value as follows:COMMITED,UNCOMMITED | [optional] 
**Software0Valid** | Pointer to **string** | Validity of ONU software 0.Software0Valid should be a value as follows:VALID,INVALID | [optional] 
**Software0Version** | Pointer to **string** | Software0 version should contain 1 to 14 ASCII characters. | [optional] 
**Software1Active** | Pointer to **string** | Status of ONU software 1.software1Active should be a value as follows:ACTIVE,INACTIVE | [optional] 
**Software1Commited** | Pointer to **string** | Commit status of ONU software 1.Software1Commited should be a value as follows:COMMITED,UNCOMMITED | [optional] 
**Software1Valid** | Pointer to **string** | Validity of ONU software 1.software1Valid should be a value as follows:VALID,INVALID | [optional] 
**Software1Version** | Pointer to **string** | Software1 version should contain 1 to 14 ASCII characters. | [optional] 

## Methods

### NewSoftwareDetailConfigDTO

`func NewSoftwareDetailConfigDTO() *SoftwareDetailConfigDTO`

NewSoftwareDetailConfigDTO instantiates a new SoftwareDetailConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoftwareDetailConfigDTOWithDefaults

`func NewSoftwareDetailConfigDTOWithDefaults() *SoftwareDetailConfigDTO`

NewSoftwareDetailConfigDTOWithDefaults instantiates a new SoftwareDetailConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSoftware0Active

`func (o *SoftwareDetailConfigDTO) GetSoftware0Active() string`

GetSoftware0Active returns the Software0Active field if non-nil, zero value otherwise.

### GetSoftware0ActiveOk

`func (o *SoftwareDetailConfigDTO) GetSoftware0ActiveOk() (*string, bool)`

GetSoftware0ActiveOk returns a tuple with the Software0Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware0Active

`func (o *SoftwareDetailConfigDTO) SetSoftware0Active(v string)`

SetSoftware0Active sets Software0Active field to given value.

### HasSoftware0Active

`func (o *SoftwareDetailConfigDTO) HasSoftware0Active() bool`

HasSoftware0Active returns a boolean if a field has been set.

### GetSoftware0Commited

`func (o *SoftwareDetailConfigDTO) GetSoftware0Commited() string`

GetSoftware0Commited returns the Software0Commited field if non-nil, zero value otherwise.

### GetSoftware0CommitedOk

`func (o *SoftwareDetailConfigDTO) GetSoftware0CommitedOk() (*string, bool)`

GetSoftware0CommitedOk returns a tuple with the Software0Commited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware0Commited

`func (o *SoftwareDetailConfigDTO) SetSoftware0Commited(v string)`

SetSoftware0Commited sets Software0Commited field to given value.

### HasSoftware0Commited

`func (o *SoftwareDetailConfigDTO) HasSoftware0Commited() bool`

HasSoftware0Commited returns a boolean if a field has been set.

### GetSoftware0Valid

`func (o *SoftwareDetailConfigDTO) GetSoftware0Valid() string`

GetSoftware0Valid returns the Software0Valid field if non-nil, zero value otherwise.

### GetSoftware0ValidOk

`func (o *SoftwareDetailConfigDTO) GetSoftware0ValidOk() (*string, bool)`

GetSoftware0ValidOk returns a tuple with the Software0Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware0Valid

`func (o *SoftwareDetailConfigDTO) SetSoftware0Valid(v string)`

SetSoftware0Valid sets Software0Valid field to given value.

### HasSoftware0Valid

`func (o *SoftwareDetailConfigDTO) HasSoftware0Valid() bool`

HasSoftware0Valid returns a boolean if a field has been set.

### GetSoftware0Version

`func (o *SoftwareDetailConfigDTO) GetSoftware0Version() string`

GetSoftware0Version returns the Software0Version field if non-nil, zero value otherwise.

### GetSoftware0VersionOk

`func (o *SoftwareDetailConfigDTO) GetSoftware0VersionOk() (*string, bool)`

GetSoftware0VersionOk returns a tuple with the Software0Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware0Version

`func (o *SoftwareDetailConfigDTO) SetSoftware0Version(v string)`

SetSoftware0Version sets Software0Version field to given value.

### HasSoftware0Version

`func (o *SoftwareDetailConfigDTO) HasSoftware0Version() bool`

HasSoftware0Version returns a boolean if a field has been set.

### GetSoftware1Active

`func (o *SoftwareDetailConfigDTO) GetSoftware1Active() string`

GetSoftware1Active returns the Software1Active field if non-nil, zero value otherwise.

### GetSoftware1ActiveOk

`func (o *SoftwareDetailConfigDTO) GetSoftware1ActiveOk() (*string, bool)`

GetSoftware1ActiveOk returns a tuple with the Software1Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware1Active

`func (o *SoftwareDetailConfigDTO) SetSoftware1Active(v string)`

SetSoftware1Active sets Software1Active field to given value.

### HasSoftware1Active

`func (o *SoftwareDetailConfigDTO) HasSoftware1Active() bool`

HasSoftware1Active returns a boolean if a field has been set.

### GetSoftware1Commited

`func (o *SoftwareDetailConfigDTO) GetSoftware1Commited() string`

GetSoftware1Commited returns the Software1Commited field if non-nil, zero value otherwise.

### GetSoftware1CommitedOk

`func (o *SoftwareDetailConfigDTO) GetSoftware1CommitedOk() (*string, bool)`

GetSoftware1CommitedOk returns a tuple with the Software1Commited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware1Commited

`func (o *SoftwareDetailConfigDTO) SetSoftware1Commited(v string)`

SetSoftware1Commited sets Software1Commited field to given value.

### HasSoftware1Commited

`func (o *SoftwareDetailConfigDTO) HasSoftware1Commited() bool`

HasSoftware1Commited returns a boolean if a field has been set.

### GetSoftware1Valid

`func (o *SoftwareDetailConfigDTO) GetSoftware1Valid() string`

GetSoftware1Valid returns the Software1Valid field if non-nil, zero value otherwise.

### GetSoftware1ValidOk

`func (o *SoftwareDetailConfigDTO) GetSoftware1ValidOk() (*string, bool)`

GetSoftware1ValidOk returns a tuple with the Software1Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware1Valid

`func (o *SoftwareDetailConfigDTO) SetSoftware1Valid(v string)`

SetSoftware1Valid sets Software1Valid field to given value.

### HasSoftware1Valid

`func (o *SoftwareDetailConfigDTO) HasSoftware1Valid() bool`

HasSoftware1Valid returns a boolean if a field has been set.

### GetSoftware1Version

`func (o *SoftwareDetailConfigDTO) GetSoftware1Version() string`

GetSoftware1Version returns the Software1Version field if non-nil, zero value otherwise.

### GetSoftware1VersionOk

`func (o *SoftwareDetailConfigDTO) GetSoftware1VersionOk() (*string, bool)`

GetSoftware1VersionOk returns a tuple with the Software1Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware1Version

`func (o *SoftwareDetailConfigDTO) SetSoftware1Version(v string)`

SetSoftware1Version sets Software1Version field to given value.

### HasSoftware1Version

`func (o *SoftwareDetailConfigDTO) HasSoftware1Version() bool`

HasSoftware1Version returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


