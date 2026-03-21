# DisableNat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Disable Nat. | [optional] 
**Interface** | **string** | The wan of the Disable Nat. | 
**LanList** | **[]string** | A list of Lan of the Disable Nat. | 
**Name** | **string** | Name of the Disable Nat. | 
**Status** | **bool** | The Status of the Disable Nat. | 

## Methods

### NewDisableNat

`func NewDisableNat(interface_ string, lanList []string, name string, status bool, ) *DisableNat`

NewDisableNat instantiates a new DisableNat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableNatWithDefaults

`func NewDisableNatWithDefaults() *DisableNat`

NewDisableNatWithDefaults instantiates a new DisableNat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DisableNat) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DisableNat) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DisableNat) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DisableNat) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInterface

`func (o *DisableNat) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DisableNat) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DisableNat) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetLanList

`func (o *DisableNat) GetLanList() []string`

GetLanList returns the LanList field if non-nil, zero value otherwise.

### GetLanListOk

`func (o *DisableNat) GetLanListOk() (*[]string, bool)`

GetLanListOk returns a tuple with the LanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanList

`func (o *DisableNat) SetLanList(v []string)`

SetLanList sets LanList field to given value.


### GetName

`func (o *DisableNat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DisableNat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DisableNat) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DisableNat) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisableNat) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisableNat) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


