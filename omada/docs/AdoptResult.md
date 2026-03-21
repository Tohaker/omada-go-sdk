# AdoptResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdoptErrorCode** | Pointer to **int32** | Adopt result should be a value as follows: 0: Adopt Device Success; -39002: Device adoption failed because the device does not respond to adopt commands; -39003: Failed to adopt the Device. The username or password  is incorrect; -39004: Failed to adopt device; -39005: Failed to adopt this device because the device is not connected; -39329: Failed to link to the uplink AP. | [optional] 
**AdoptFailedType** | Pointer to **int32** | Adopt failed type should be a value as follows: -1: No need print username or password;-2: Need print username or password. | [optional] 
**DeviceMac** | Pointer to **string** | Device MAC | [optional] 

## Methods

### NewAdoptResult

`func NewAdoptResult() *AdoptResult`

NewAdoptResult instantiates a new AdoptResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdoptResultWithDefaults

`func NewAdoptResultWithDefaults() *AdoptResult`

NewAdoptResultWithDefaults instantiates a new AdoptResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdoptErrorCode

`func (o *AdoptResult) GetAdoptErrorCode() int32`

GetAdoptErrorCode returns the AdoptErrorCode field if non-nil, zero value otherwise.

### GetAdoptErrorCodeOk

`func (o *AdoptResult) GetAdoptErrorCodeOk() (*int32, bool)`

GetAdoptErrorCodeOk returns a tuple with the AdoptErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptErrorCode

`func (o *AdoptResult) SetAdoptErrorCode(v int32)`

SetAdoptErrorCode sets AdoptErrorCode field to given value.

### HasAdoptErrorCode

`func (o *AdoptResult) HasAdoptErrorCode() bool`

HasAdoptErrorCode returns a boolean if a field has been set.

### GetAdoptFailedType

`func (o *AdoptResult) GetAdoptFailedType() int32`

GetAdoptFailedType returns the AdoptFailedType field if non-nil, zero value otherwise.

### GetAdoptFailedTypeOk

`func (o *AdoptResult) GetAdoptFailedTypeOk() (*int32, bool)`

GetAdoptFailedTypeOk returns a tuple with the AdoptFailedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdoptFailedType

`func (o *AdoptResult) SetAdoptFailedType(v int32)`

SetAdoptFailedType sets AdoptFailedType field to given value.

### HasAdoptFailedType

`func (o *AdoptResult) HasAdoptFailedType() bool`

HasAdoptFailedType returns a boolean if a field has been set.

### GetDeviceMac

`func (o *AdoptResult) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *AdoptResult) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *AdoptResult) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *AdoptResult) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


