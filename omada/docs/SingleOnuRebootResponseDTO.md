# SingleOnuRebootResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Mac address of ONU | [optional] 
**Status** | Pointer to **int32** | ONU reboot status.Status should be a value as follows:0:success.1:fail | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 

## Methods

### NewSingleOnuRebootResponseDTO

`func NewSingleOnuRebootResponseDTO() *SingleOnuRebootResponseDTO`

NewSingleOnuRebootResponseDTO instantiates a new SingleOnuRebootResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleOnuRebootResponseDTOWithDefaults

`func NewSingleOnuRebootResponseDTOWithDefaults() *SingleOnuRebootResponseDTO`

NewSingleOnuRebootResponseDTOWithDefaults instantiates a new SingleOnuRebootResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *SingleOnuRebootResponseDTO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SingleOnuRebootResponseDTO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SingleOnuRebootResponseDTO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SingleOnuRebootResponseDTO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetStatus

`func (o *SingleOnuRebootResponseDTO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SingleOnuRebootResponseDTO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SingleOnuRebootResponseDTO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SingleOnuRebootResponseDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *SingleOnuRebootResponseDTO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *SingleOnuRebootResponseDTO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *SingleOnuRebootResponseDTO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *SingleOnuRebootResponseDTO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


