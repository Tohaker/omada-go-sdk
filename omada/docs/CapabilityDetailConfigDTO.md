# CapabilityDetailConfigDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OmccVersion** | Pointer to **string** | OmccVersion should contain uppercase and lowercase letters, numbers, and the symbols -@_:/. | [optional] 
**TotalEthNumber** | Pointer to **int32** | The number of Ethernet ports on the ONU, totalEthNumber should be within the range of 0 to 24.  | [optional] 
**TotalGemPortNumber** | Pointer to **int32** | The number of configured GEM Ports,totalGemPortNumber should be within the range of 0 to 31. | [optional] 
**TotalTcontNumber** | Pointer to **int32** | The number of configured T-conts, totalTcontNumber should be within the range of 0 to 8. | [optional] 
**TotalVoipNumber** | Pointer to **int32** | The number of VoIP ports on the ONU, totalVoipNumber should be within the range of 0 to 4. | [optional] 

## Methods

### NewCapabilityDetailConfigDTO

`func NewCapabilityDetailConfigDTO() *CapabilityDetailConfigDTO`

NewCapabilityDetailConfigDTO instantiates a new CapabilityDetailConfigDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCapabilityDetailConfigDTOWithDefaults

`func NewCapabilityDetailConfigDTOWithDefaults() *CapabilityDetailConfigDTO`

NewCapabilityDetailConfigDTOWithDefaults instantiates a new CapabilityDetailConfigDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOmccVersion

`func (o *CapabilityDetailConfigDTO) GetOmccVersion() string`

GetOmccVersion returns the OmccVersion field if non-nil, zero value otherwise.

### GetOmccVersionOk

`func (o *CapabilityDetailConfigDTO) GetOmccVersionOk() (*string, bool)`

GetOmccVersionOk returns a tuple with the OmccVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmccVersion

`func (o *CapabilityDetailConfigDTO) SetOmccVersion(v string)`

SetOmccVersion sets OmccVersion field to given value.

### HasOmccVersion

`func (o *CapabilityDetailConfigDTO) HasOmccVersion() bool`

HasOmccVersion returns a boolean if a field has been set.

### GetTotalEthNumber

`func (o *CapabilityDetailConfigDTO) GetTotalEthNumber() int32`

GetTotalEthNumber returns the TotalEthNumber field if non-nil, zero value otherwise.

### GetTotalEthNumberOk

`func (o *CapabilityDetailConfigDTO) GetTotalEthNumberOk() (*int32, bool)`

GetTotalEthNumberOk returns a tuple with the TotalEthNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEthNumber

`func (o *CapabilityDetailConfigDTO) SetTotalEthNumber(v int32)`

SetTotalEthNumber sets TotalEthNumber field to given value.

### HasTotalEthNumber

`func (o *CapabilityDetailConfigDTO) HasTotalEthNumber() bool`

HasTotalEthNumber returns a boolean if a field has been set.

### GetTotalGemPortNumber

`func (o *CapabilityDetailConfigDTO) GetTotalGemPortNumber() int32`

GetTotalGemPortNumber returns the TotalGemPortNumber field if non-nil, zero value otherwise.

### GetTotalGemPortNumberOk

`func (o *CapabilityDetailConfigDTO) GetTotalGemPortNumberOk() (*int32, bool)`

GetTotalGemPortNumberOk returns a tuple with the TotalGemPortNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalGemPortNumber

`func (o *CapabilityDetailConfigDTO) SetTotalGemPortNumber(v int32)`

SetTotalGemPortNumber sets TotalGemPortNumber field to given value.

### HasTotalGemPortNumber

`func (o *CapabilityDetailConfigDTO) HasTotalGemPortNumber() bool`

HasTotalGemPortNumber returns a boolean if a field has been set.

### GetTotalTcontNumber

`func (o *CapabilityDetailConfigDTO) GetTotalTcontNumber() int32`

GetTotalTcontNumber returns the TotalTcontNumber field if non-nil, zero value otherwise.

### GetTotalTcontNumberOk

`func (o *CapabilityDetailConfigDTO) GetTotalTcontNumberOk() (*int32, bool)`

GetTotalTcontNumberOk returns a tuple with the TotalTcontNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTcontNumber

`func (o *CapabilityDetailConfigDTO) SetTotalTcontNumber(v int32)`

SetTotalTcontNumber sets TotalTcontNumber field to given value.

### HasTotalTcontNumber

`func (o *CapabilityDetailConfigDTO) HasTotalTcontNumber() bool`

HasTotalTcontNumber returns a boolean if a field has been set.

### GetTotalVoipNumber

`func (o *CapabilityDetailConfigDTO) GetTotalVoipNumber() int32`

GetTotalVoipNumber returns the TotalVoipNumber field if non-nil, zero value otherwise.

### GetTotalVoipNumberOk

`func (o *CapabilityDetailConfigDTO) GetTotalVoipNumberOk() (*int32, bool)`

GetTotalVoipNumberOk returns a tuple with the TotalVoipNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalVoipNumber

`func (o *CapabilityDetailConfigDTO) SetTotalVoipNumber(v int32)`

SetTotalVoipNumber sets TotalVoipNumber field to given value.

### HasTotalVoipNumber

`func (o *CapabilityDetailConfigDTO) HasTotalVoipNumber() bool`

HasTotalVoipNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


