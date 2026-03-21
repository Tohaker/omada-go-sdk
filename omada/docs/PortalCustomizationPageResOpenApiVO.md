# PortalCustomizationPageResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImportedPortalPage** | Pointer to [**ImportedPortalPageResOpenApiVO**](ImportedPortalPageResOpenApiVO.md) |  | [optional] 
**PageType** | Pointer to **int32** | Page type, should be a value as follows: 1: Use default page, 2: use uploaded page. Required when portal auth type is not 4: External Portal Server | [optional] 
**PortalCustomize** | Pointer to [**PortalCustomizeResOpenApiVO**](PortalCustomizeResOpenApiVO.md) |  | [optional] 

## Methods

### NewPortalCustomizationPageResOpenApiVO

`func NewPortalCustomizationPageResOpenApiVO() *PortalCustomizationPageResOpenApiVO`

NewPortalCustomizationPageResOpenApiVO instantiates a new PortalCustomizationPageResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCustomizationPageResOpenApiVOWithDefaults

`func NewPortalCustomizationPageResOpenApiVOWithDefaults() *PortalCustomizationPageResOpenApiVO`

NewPortalCustomizationPageResOpenApiVOWithDefaults instantiates a new PortalCustomizationPageResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImportedPortalPage

`func (o *PortalCustomizationPageResOpenApiVO) GetImportedPortalPage() ImportedPortalPageResOpenApiVO`

GetImportedPortalPage returns the ImportedPortalPage field if non-nil, zero value otherwise.

### GetImportedPortalPageOk

`func (o *PortalCustomizationPageResOpenApiVO) GetImportedPortalPageOk() (*ImportedPortalPageResOpenApiVO, bool)`

GetImportedPortalPageOk returns a tuple with the ImportedPortalPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedPortalPage

`func (o *PortalCustomizationPageResOpenApiVO) SetImportedPortalPage(v ImportedPortalPageResOpenApiVO)`

SetImportedPortalPage sets ImportedPortalPage field to given value.

### HasImportedPortalPage

`func (o *PortalCustomizationPageResOpenApiVO) HasImportedPortalPage() bool`

HasImportedPortalPage returns a boolean if a field has been set.

### GetPageType

`func (o *PortalCustomizationPageResOpenApiVO) GetPageType() int32`

GetPageType returns the PageType field if non-nil, zero value otherwise.

### GetPageTypeOk

`func (o *PortalCustomizationPageResOpenApiVO) GetPageTypeOk() (*int32, bool)`

GetPageTypeOk returns a tuple with the PageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageType

`func (o *PortalCustomizationPageResOpenApiVO) SetPageType(v int32)`

SetPageType sets PageType field to given value.

### HasPageType

`func (o *PortalCustomizationPageResOpenApiVO) HasPageType() bool`

HasPageType returns a boolean if a field has been set.

### GetPortalCustomize

`func (o *PortalCustomizationPageResOpenApiVO) GetPortalCustomize() PortalCustomizeResOpenApiVO`

GetPortalCustomize returns the PortalCustomize field if non-nil, zero value otherwise.

### GetPortalCustomizeOk

`func (o *PortalCustomizationPageResOpenApiVO) GetPortalCustomizeOk() (*PortalCustomizeResOpenApiVO, bool)`

GetPortalCustomizeOk returns a tuple with the PortalCustomize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalCustomize

`func (o *PortalCustomizationPageResOpenApiVO) SetPortalCustomize(v PortalCustomizeResOpenApiVO)`

SetPortalCustomize sets PortalCustomize field to given value.

### HasPortalCustomize

`func (o *PortalCustomizationPageResOpenApiVO) HasPortalCustomize() bool`

HasPortalCustomize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


