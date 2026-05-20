# CarrierOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CarrierName** | Pointer to **string** | Carrier Name. | [optional] 
**Epdgs** | Pointer to [**[]EPDGOpenApiVO**](EPDGOpenApiVO.md) | epdgs | [optional] 

## Methods

### NewCarrierOpenApiVO

`func NewCarrierOpenApiVO() *CarrierOpenApiVO`

NewCarrierOpenApiVO instantiates a new CarrierOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierOpenApiVOWithDefaults

`func NewCarrierOpenApiVOWithDefaults() *CarrierOpenApiVO`

NewCarrierOpenApiVOWithDefaults instantiates a new CarrierOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCarrierName

`func (o *CarrierOpenApiVO) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *CarrierOpenApiVO) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *CarrierOpenApiVO) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *CarrierOpenApiVO) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetEpdgs

`func (o *CarrierOpenApiVO) GetEpdgs() []EPDGOpenApiVO`

GetEpdgs returns the Epdgs field if non-nil, zero value otherwise.

### GetEpdgsOk

`func (o *CarrierOpenApiVO) GetEpdgsOk() (*[]EPDGOpenApiVO, bool)`

GetEpdgsOk returns a tuple with the Epdgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpdgs

`func (o *CarrierOpenApiVO) SetEpdgs(v []EPDGOpenApiVO)`

SetEpdgs sets Epdgs field to given value.

### HasEpdgs

`func (o *CarrierOpenApiVO) HasEpdgs() bool`

HasEpdgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


