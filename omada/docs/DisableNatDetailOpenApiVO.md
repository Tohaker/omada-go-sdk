# DisableNatDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the Disable Nat. | [optional] 
**Id** | Pointer to **string** | ID of the Disable Nat. | [optional] 
**Interface** | **string** | The wan of the Disable Nat. | 
**LanList** | **[]string** | A list of Lan of the Disable Nat. | 
**Name** | **string** | Name of the Disable Nat. | 
**Status** | **bool** | The Status of the Disable Nat. | 

## Methods

### NewDisableNatDetailOpenApiVO

`func NewDisableNatDetailOpenApiVO(interface_ string, lanList []string, name string, status bool, ) *DisableNatDetailOpenApiVO`

NewDisableNatDetailOpenApiVO instantiates a new DisableNatDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDisableNatDetailOpenApiVOWithDefaults

`func NewDisableNatDetailOpenApiVOWithDefaults() *DisableNatDetailOpenApiVO`

NewDisableNatDetailOpenApiVOWithDefaults instantiates a new DisableNatDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *DisableNatDetailOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DisableNatDetailOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DisableNatDetailOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DisableNatDetailOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *DisableNatDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DisableNatDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DisableNatDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DisableNatDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterface

`func (o *DisableNatDetailOpenApiVO) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *DisableNatDetailOpenApiVO) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *DisableNatDetailOpenApiVO) SetInterface(v string)`

SetInterface sets Interface field to given value.


### GetLanList

`func (o *DisableNatDetailOpenApiVO) GetLanList() []string`

GetLanList returns the LanList field if non-nil, zero value otherwise.

### GetLanListOk

`func (o *DisableNatDetailOpenApiVO) GetLanListOk() (*[]string, bool)`

GetLanListOk returns a tuple with the LanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanList

`func (o *DisableNatDetailOpenApiVO) SetLanList(v []string)`

SetLanList sets LanList field to given value.


### GetName

`func (o *DisableNatDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DisableNatDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DisableNatDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *DisableNatDetailOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DisableNatDetailOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DisableNatDetailOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


