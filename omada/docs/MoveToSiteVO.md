# MoveToSiteVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Target site key | [optional] 
**DeviceMacs** | Pointer to **[]string** | Device macs.If only stacking is migrated, then it is empty. | [optional] 
**StackIds** | Pointer to **[]string** | The stack to be migrated; all stack members are migrated together during the migration. | [optional] 

## Methods

### NewMoveToSiteVO

`func NewMoveToSiteVO() *MoveToSiteVO`

NewMoveToSiteVO instantiates a new MoveToSiteVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveToSiteVOWithDefaults

`func NewMoveToSiteVOWithDefaults() *MoveToSiteVO`

NewMoveToSiteVOWithDefaults instantiates a new MoveToSiteVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *MoveToSiteVO) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MoveToSiteVO) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MoveToSiteVO) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *MoveToSiteVO) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetDeviceMacs

`func (o *MoveToSiteVO) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *MoveToSiteVO) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *MoveToSiteVO) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.

### HasDeviceMacs

`func (o *MoveToSiteVO) HasDeviceMacs() bool`

HasDeviceMacs returns a boolean if a field has been set.

### GetStackIds

`func (o *MoveToSiteVO) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *MoveToSiteVO) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *MoveToSiteVO) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.

### HasStackIds

`func (o *MoveToSiteVO) HasStackIds() bool`

HasStackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


