# CreateWifiCallingProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierList** | Pointer to [**[]CarrierOpenApiVO**](CarrierOpenApiVO.md) | carrierList | [optional] 
**Description** | Pointer to **string** | Description of the Wi-Fi calling profile. It should contain 1 to 32 UTF-8 characters. | [optional] 
**Name** | Pointer to **string** | Wi-Fi Calling Profile Name. It should contain 1 to 32 UTF-8 characters. | [optional] 

## Methods

### NewCreateWifiCallingProfileOpenApiVO

`func NewCreateWifiCallingProfileOpenApiVO() *CreateWifiCallingProfileOpenApiVO`

NewCreateWifiCallingProfileOpenApiVO instantiates a new CreateWifiCallingProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWifiCallingProfileOpenApiVOWithDefaults

`func NewCreateWifiCallingProfileOpenApiVOWithDefaults() *CreateWifiCallingProfileOpenApiVO`

NewCreateWifiCallingProfileOpenApiVOWithDefaults instantiates a new CreateWifiCallingProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierList

`func (o *CreateWifiCallingProfileOpenApiVO) GetCarrierList() []CarrierOpenApiVO`

GetCarrierList returns the CarrierList field if non-nil, zero value otherwise.

### GetCarrierListOk

`func (o *CreateWifiCallingProfileOpenApiVO) GetCarrierListOk() (*[]CarrierOpenApiVO, bool)`

GetCarrierListOk returns a tuple with the CarrierList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierList

`func (o *CreateWifiCallingProfileOpenApiVO) SetCarrierList(v []CarrierOpenApiVO)`

SetCarrierList sets CarrierList field to given value.

### HasCarrierList

`func (o *CreateWifiCallingProfileOpenApiVO) HasCarrierList() bool`

HasCarrierList returns a boolean if a field has been set.

### GetDescription

`func (o *CreateWifiCallingProfileOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWifiCallingProfileOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWifiCallingProfileOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWifiCallingProfileOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CreateWifiCallingProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWifiCallingProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWifiCallingProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateWifiCallingProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


