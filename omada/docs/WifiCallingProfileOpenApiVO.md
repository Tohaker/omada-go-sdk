# WifiCallingProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierList** | Pointer to [**[]CarrierOpenApiVO**](CarrierOpenApiVO.md) | carrierList | [optional] 
**Description** | Pointer to **string** | The Description of Wi-Fi Calling Profile. It should contain 1 to 32 UTF-8 characters. | [optional] 
**Id** | Pointer to **string** | Wi-Fi Calling Profile ID | [optional] 
**Name** | Pointer to **string** | Wi-Fi Calling Profile Name. It should contain 1 to 32 UTF-8 characters. | [optional] 
**Resource** | Pointer to **int32** | The incident notifiction setting creation resource, such as: 0: new created, 1: from template, 2: override. | [optional] 

## Methods

### NewWifiCallingProfileOpenApiVO

`func NewWifiCallingProfileOpenApiVO() *WifiCallingProfileOpenApiVO`

NewWifiCallingProfileOpenApiVO instantiates a new WifiCallingProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWifiCallingProfileOpenApiVOWithDefaults

`func NewWifiCallingProfileOpenApiVOWithDefaults() *WifiCallingProfileOpenApiVO`

NewWifiCallingProfileOpenApiVOWithDefaults instantiates a new WifiCallingProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierList

`func (o *WifiCallingProfileOpenApiVO) GetCarrierList() []CarrierOpenApiVO`

GetCarrierList returns the CarrierList field if non-nil, zero value otherwise.

### GetCarrierListOk

`func (o *WifiCallingProfileOpenApiVO) GetCarrierListOk() (*[]CarrierOpenApiVO, bool)`

GetCarrierListOk returns a tuple with the CarrierList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierList

`func (o *WifiCallingProfileOpenApiVO) SetCarrierList(v []CarrierOpenApiVO)`

SetCarrierList sets CarrierList field to given value.

### HasCarrierList

`func (o *WifiCallingProfileOpenApiVO) HasCarrierList() bool`

HasCarrierList returns a boolean if a field has been set.

### GetDescription

`func (o *WifiCallingProfileOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WifiCallingProfileOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WifiCallingProfileOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WifiCallingProfileOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *WifiCallingProfileOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WifiCallingProfileOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WifiCallingProfileOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WifiCallingProfileOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *WifiCallingProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WifiCallingProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WifiCallingProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WifiCallingProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResource

`func (o *WifiCallingProfileOpenApiVO) GetResource() int32`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *WifiCallingProfileOpenApiVO) GetResourceOk() (*int32, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *WifiCallingProfileOpenApiVO) SetResource(v int32)`

SetResource sets Resource field to given value.

### HasResource

`func (o *WifiCallingProfileOpenApiVO) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


