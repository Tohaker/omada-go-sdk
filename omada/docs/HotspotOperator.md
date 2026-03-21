# HotspotOperator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSite** | Pointer to **string** | Last Site ID | [optional] 
**Name** | **string** | Operator name should contain 1 to 128 ASCII characters. | 
**Note** | Pointer to **string** | Operator note should contain 1 to 256 ASCII characters. | [optional] 
**OperatorRoleType** | Pointer to **int32** | Operator role type should be a value as follows: 0: Administrator; 1: Viewer. | [optional] 
**Password** | **string** | Operator password should contain 1 to 128 ASCII characters. | 
**SelectedSites** | **[]string** | Selected site ID list should contain at least one site for each operator. | 

## Methods

### NewHotspotOperator

`func NewHotspotOperator(name string, password string, selectedSites []string, ) *HotspotOperator`

NewHotspotOperator instantiates a new HotspotOperator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotOperatorWithDefaults

`func NewHotspotOperatorWithDefaults() *HotspotOperator`

NewHotspotOperatorWithDefaults instantiates a new HotspotOperator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSite

`func (o *HotspotOperator) GetLastSite() string`

GetLastSite returns the LastSite field if non-nil, zero value otherwise.

### GetLastSiteOk

`func (o *HotspotOperator) GetLastSiteOk() (*string, bool)`

GetLastSiteOk returns a tuple with the LastSite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSite

`func (o *HotspotOperator) SetLastSite(v string)`

SetLastSite sets LastSite field to given value.

### HasLastSite

`func (o *HotspotOperator) HasLastSite() bool`

HasLastSite returns a boolean if a field has been set.

### GetName

`func (o *HotspotOperator) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *HotspotOperator) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *HotspotOperator) SetName(v string)`

SetName sets Name field to given value.


### GetNote

`func (o *HotspotOperator) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *HotspotOperator) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *HotspotOperator) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *HotspotOperator) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOperatorRoleType

`func (o *HotspotOperator) GetOperatorRoleType() int32`

GetOperatorRoleType returns the OperatorRoleType field if non-nil, zero value otherwise.

### GetOperatorRoleTypeOk

`func (o *HotspotOperator) GetOperatorRoleTypeOk() (*int32, bool)`

GetOperatorRoleTypeOk returns a tuple with the OperatorRoleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperatorRoleType

`func (o *HotspotOperator) SetOperatorRoleType(v int32)`

SetOperatorRoleType sets OperatorRoleType field to given value.

### HasOperatorRoleType

`func (o *HotspotOperator) HasOperatorRoleType() bool`

HasOperatorRoleType returns a boolean if a field has been set.

### GetPassword

`func (o *HotspotOperator) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *HotspotOperator) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *HotspotOperator) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetSelectedSites

`func (o *HotspotOperator) GetSelectedSites() []string`

GetSelectedSites returns the SelectedSites field if non-nil, zero value otherwise.

### GetSelectedSitesOk

`func (o *HotspotOperator) GetSelectedSitesOk() (*[]string, bool)`

GetSelectedSitesOk returns a tuple with the SelectedSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedSites

`func (o *HotspotOperator) SetSelectedSites(v []string)`

SetSelectedSites sets SelectedSites field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


