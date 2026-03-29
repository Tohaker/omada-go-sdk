# BandResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Band** | Pointer to **string** | The name of the band. | [optional] 
**Belong** | Pointer to **bool** | Whether it belongs to the carrier to which the SIM card belongs. | [optional] 

## Methods

### NewBandResultOpenApiVO

`func NewBandResultOpenApiVO() *BandResultOpenApiVO`

NewBandResultOpenApiVO instantiates a new BandResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandResultOpenApiVOWithDefaults

`func NewBandResultOpenApiVOWithDefaults() *BandResultOpenApiVO`

NewBandResultOpenApiVOWithDefaults instantiates a new BandResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBand

`func (o *BandResultOpenApiVO) GetBand() string`

GetBand returns the Band field if non-nil, zero value otherwise.

### GetBandOk

`func (o *BandResultOpenApiVO) GetBandOk() (*string, bool)`

GetBandOk returns a tuple with the Band field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBand

`func (o *BandResultOpenApiVO) SetBand(v string)`

SetBand sets Band field to given value.

### HasBand

`func (o *BandResultOpenApiVO) HasBand() bool`

HasBand returns a boolean if a field has been set.

### GetBelong

`func (o *BandResultOpenApiVO) GetBelong() bool`

GetBelong returns the Belong field if non-nil, zero value otherwise.

### GetBelongOk

`func (o *BandResultOpenApiVO) GetBelongOk() (*bool, bool)`

GetBelongOk returns a tuple with the Belong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBelong

`func (o *BandResultOpenApiVO) SetBelong(v bool)`

SetBelong sets Belong field to given value.

### HasBelong

`func (o *BandResultOpenApiVO) HasBelong() bool`

HasBelong returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


