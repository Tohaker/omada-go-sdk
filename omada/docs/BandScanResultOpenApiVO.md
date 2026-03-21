# BandScanResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bands** | Pointer to [**[]BandResultOpenApiVO**](BandResultOpenApiVO.md) | The result of the band scan. | [optional] 
**Isp** | Pointer to **string** | Internet service provider. | [optional] 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCard] should not be null.1: SIM1; 2: SIM2. | [optional] 
**Status** | Pointer to **int32** | The status of the band scan: 0 - Failed, 1 - Succeeded, 2 - Scanning. | [optional] 

## Methods

### NewBandScanResultOpenApiVO

`func NewBandScanResultOpenApiVO() *BandScanResultOpenApiVO`

NewBandScanResultOpenApiVO instantiates a new BandScanResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandScanResultOpenApiVOWithDefaults

`func NewBandScanResultOpenApiVOWithDefaults() *BandScanResultOpenApiVO`

NewBandScanResultOpenApiVOWithDefaults instantiates a new BandScanResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBands

`func (o *BandScanResultOpenApiVO) GetBands() []BandResultOpenApiVO`

GetBands returns the Bands field if non-nil, zero value otherwise.

### GetBandsOk

`func (o *BandScanResultOpenApiVO) GetBandsOk() (*[]BandResultOpenApiVO, bool)`

GetBandsOk returns a tuple with the Bands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBands

`func (o *BandScanResultOpenApiVO) SetBands(v []BandResultOpenApiVO)`

SetBands sets Bands field to given value.

### HasBands

`func (o *BandScanResultOpenApiVO) HasBands() bool`

HasBands returns a boolean if a field has been set.

### GetIsp

`func (o *BandScanResultOpenApiVO) GetIsp() string`

GetIsp returns the Isp field if non-nil, zero value otherwise.

### GetIspOk

`func (o *BandScanResultOpenApiVO) GetIspOk() (*string, bool)`

GetIspOk returns a tuple with the Isp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsp

`func (o *BandScanResultOpenApiVO) SetIsp(v string)`

SetIsp sets Isp field to given value.

### HasIsp

`func (o *BandScanResultOpenApiVO) HasIsp() bool`

HasIsp returns a boolean if a field has been set.

### GetSimCard

`func (o *BandScanResultOpenApiVO) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *BandScanResultOpenApiVO) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *BandScanResultOpenApiVO) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *BandScanResultOpenApiVO) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.

### GetStatus

`func (o *BandScanResultOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BandScanResultOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BandScanResultOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BandScanResultOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


