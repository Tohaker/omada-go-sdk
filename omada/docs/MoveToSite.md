# MoveToSite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Site** | Pointer to **string** | Target site key | [optional] 
**DeviceMacs** | Pointer to **[]string** | Device mac list | [optional] 
**StackIds** | Pointer to **[]string** | Stack id list | [optional] 

## Methods

### NewMoveToSite

`func NewMoveToSite() *MoveToSite`

NewMoveToSite instantiates a new MoveToSite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveToSiteWithDefaults

`func NewMoveToSiteWithDefaults() *MoveToSite`

NewMoveToSiteWithDefaults instantiates a new MoveToSite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSite

`func (o *MoveToSite) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *MoveToSite) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *MoveToSite) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *MoveToSite) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetDeviceMacs

`func (o *MoveToSite) GetDeviceMacs() []string`

GetDeviceMacs returns the DeviceMacs field if non-nil, zero value otherwise.

### GetDeviceMacsOk

`func (o *MoveToSite) GetDeviceMacsOk() (*[]string, bool)`

GetDeviceMacsOk returns a tuple with the DeviceMacs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMacs

`func (o *MoveToSite) SetDeviceMacs(v []string)`

SetDeviceMacs sets DeviceMacs field to given value.

### HasDeviceMacs

`func (o *MoveToSite) HasDeviceMacs() bool`

HasDeviceMacs returns a boolean if a field has been set.

### GetStackIds

`func (o *MoveToSite) GetStackIds() []string`

GetStackIds returns the StackIds field if non-nil, zero value otherwise.

### GetStackIdsOk

`func (o *MoveToSite) GetStackIdsOk() (*[]string, bool)`

GetStackIdsOk returns a tuple with the StackIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackIds

`func (o *MoveToSite) SetStackIds(v []string)`

SetStackIds sets StackIds field to given value.

### HasStackIds

`func (o *MoveToSite) HasStackIds() bool`

HasStackIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


