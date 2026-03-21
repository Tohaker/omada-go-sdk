# BandScanStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **int32** | Band mode. 1: 4G scan; 2: 5G scan. | [optional] 
**PortUuid** | **string** | The ID of the port. | 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCard] should not be null.1: SIM1; 2: SIM2. | [optional] 

## Methods

### NewBandScanStart

`func NewBandScanStart(portUuid string, ) *BandScanStart`

NewBandScanStart instantiates a new BandScanStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBandScanStartWithDefaults

`func NewBandScanStartWithDefaults() *BandScanStart`

NewBandScanStartWithDefaults instantiates a new BandScanStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *BandScanStart) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *BandScanStart) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *BandScanStart) SetMode(v int32)`

SetMode sets Mode field to given value.

### HasMode

`func (o *BandScanStart) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetPortUuid

`func (o *BandScanStart) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *BandScanStart) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *BandScanStart) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.


### GetSimCard

`func (o *BandScanStart) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *BandScanStart) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *BandScanStart) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *BandScanStart) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


