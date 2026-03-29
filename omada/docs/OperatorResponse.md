# OperatorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Operator ID | [optional] 
**Name** | **string** | Operator name should contain 1 to 128 ASCII characters. | 
**Note** | Pointer to **string** | Operator note should contain 1 to 256 ASCII characters. | [optional] 
**OperatorRoleType** | Pointer to **int32** | Operator role type should be a value as follows: 0: Administrator; 1: Viewer. | [optional] 
**Password** | **string** | Operator password should contain 1 to 128 ASCII characters. | 
**SelectedSites** | **[]string** | Selected site ID list should contain at least one site for each operator. | 
**Sites** | Pointer to [**[]SiteInfoOpenApiVO**](SiteInfoOpenApiVO.md) | Site ID List | [optional] 

## Methods

### NewOperatorResponse

`func NewOperatorResponse(name string, password string, selectedSites []string, ) *OperatorResponse`

NewOperatorResponse instantiates a new OperatorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperatorResponseWithDefaults

`func NewOperatorResponseWithDefaults() *OperatorResponse`

NewOperatorResponseWithDefaults instantiates a new OperatorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OperatorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OperatorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OperatorResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OperatorResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OperatorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OperatorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OperatorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *OperatorResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *OperatorResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *OperatorResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *OperatorResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOperatorRoleType

`func (o *OperatorResponse) GetOperatorRoleType() int32`

GetOperatorRoleType returns the OperatorRoleType field if non-nil, zero value otherwise.

### GetOperatorRoleTypeOk

`func (o *OperatorResponse) GetOperatorRoleTypeOk() (*int32, bool)`

GetOperatorRoleTypeOk returns a tuple with the OperatorRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorRoleType

`func (o *OperatorResponse) SetOperatorRoleType(v int32)`

SetOperatorRoleType sets OperatorRoleType field to given value.

### HasOperatorRoleType

`func (o *OperatorResponse) HasOperatorRoleType() bool`

HasOperatorRoleType returns a boolean if a field has been set.

### GetPassword

`func (o *OperatorResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *OperatorResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *OperatorResponse) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSelectedSites

`func (o *OperatorResponse) GetSelectedSites() []string`

GetSelectedSites returns the SelectedSites field if non-nil, zero value otherwise.

### GetSelectedSitesOk

`func (o *OperatorResponse) GetSelectedSitesOk() (*[]string, bool)`

GetSelectedSitesOk returns a tuple with the SelectedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSites

`func (o *OperatorResponse) SetSelectedSites(v []string)`

SetSelectedSites sets SelectedSites field to given value.


### GetSites

`func (o *OperatorResponse) GetSites() []SiteInfoOpenApiVO`

GetSites returns the Sites field if non-nil, zero value otherwise.

### GetSitesOk

`func (o *OperatorResponse) GetSitesOk() (*[]SiteInfoOpenApiVO, bool)`

GetSitesOk returns a tuple with the Sites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSites

`func (o *OperatorResponse) SetSites(v []SiteInfoOpenApiVO)`

SetSites sets Sites field to given value.

### HasSites

`func (o *OperatorResponse) HasSites() bool`

HasSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


