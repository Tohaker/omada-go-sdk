# DropEap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApMac** | Pointer to **string** | AP MAC | [optional] 
**Avg** | Pointer to **float64** | Average current AP drop packet rate | [optional] 
**Dropouts** | Pointer to [**[]Drop**](Drop.md) | AP dropouts timing list | [optional] 
**Model** | Pointer to **string** | AP model | [optional] 
**ModelVersion** | Pointer to **string** | AP model Version | [optional] 
**Name** | Pointer to **string** | AP name | [optional] 
**Status** | Pointer to **int32** | AP status, 0: connected, 1: disconnected | [optional] 

## Methods

### NewDropEap

`func NewDropEap() *DropEap`

NewDropEap instantiates a new DropEap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDropEapWithDefaults

`func NewDropEapWithDefaults() *DropEap`

NewDropEapWithDefaults instantiates a new DropEap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApMac

`func (o *DropEap) GetApMac() string`

GetApMac returns the ApMac field if non-nil, zero value otherwise.

### GetApMacOk

`func (o *DropEap) GetApMacOk() (*string, bool)`

GetApMacOk returns a tuple with the ApMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApMac

`func (o *DropEap) SetApMac(v string)`

SetApMac sets ApMac field to given value.

### HasApMac

`func (o *DropEap) HasApMac() bool`

HasApMac returns a boolean if a field has been set.

### GetAvg

`func (o *DropEap) GetAvg() float64`

GetAvg returns the Avg field if non-nil, zero value otherwise.

### GetAvgOk

`func (o *DropEap) GetAvgOk() (*float64, bool)`

GetAvgOk returns a tuple with the Avg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvg

`func (o *DropEap) SetAvg(v float64)`

SetAvg sets Avg field to given value.

### HasAvg

`func (o *DropEap) HasAvg() bool`

HasAvg returns a boolean if a field has been set.

### GetDropouts

`func (o *DropEap) GetDropouts() []Drop`

GetDropouts returns the Dropouts field if non-nil, zero value otherwise.

### GetDropoutsOk

`func (o *DropEap) GetDropoutsOk() (*[]Drop, bool)`

GetDropoutsOk returns a tuple with the Dropouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropouts

`func (o *DropEap) SetDropouts(v []Drop)`

SetDropouts sets Dropouts field to given value.

### HasDropouts

`func (o *DropEap) HasDropouts() bool`

HasDropouts returns a boolean if a field has been set.

### GetModel

`func (o *DropEap) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DropEap) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DropEap) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DropEap) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *DropEap) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *DropEap) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *DropEap) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *DropEap) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *DropEap) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DropEap) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DropEap) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DropEap) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DropEap) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DropEap) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DropEap) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DropEap) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


