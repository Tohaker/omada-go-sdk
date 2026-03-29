# IspScanStartOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortUuid** | **string** | The ID of the port. | 
**SimCard** | Pointer to **int32** | When the device supports Dual-SIM card, parameter [simCard] should not be null.1: SIM1; 2: SIM2. | [optional] 

## Methods

### NewIspScanStartOpenApiVO

`func NewIspScanStartOpenApiVO(portUuid string, ) *IspScanStartOpenApiVO`

NewIspScanStartOpenApiVO instantiates a new IspScanStartOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspScanStartOpenApiVOWithDefaults

`func NewIspScanStartOpenApiVOWithDefaults() *IspScanStartOpenApiVO`

NewIspScanStartOpenApiVOWithDefaults instantiates a new IspScanStartOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortUuid

`func (o *IspScanStartOpenApiVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *IspScanStartOpenApiVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *IspScanStartOpenApiVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.


### GetSimCard

`func (o *IspScanStartOpenApiVO) GetSimCard() int32`

GetSimCard returns the SimCard field if non-nil, zero value otherwise.

### GetSimCardOk

`func (o *IspScanStartOpenApiVO) GetSimCardOk() (*int32, bool)`

GetSimCardOk returns a tuple with the SimCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCard

`func (o *IspScanStartOpenApiVO) SetSimCard(v int32)`

SetSimCard sets SimCard field to given value.

### HasSimCard

`func (o *IspScanStartOpenApiVO) HasSimCard() bool`

HasSimCard returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


