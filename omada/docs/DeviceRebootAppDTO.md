# DeviceRebootAppDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ControlType** | Pointer to **int32** | Types of user access to the Omada Controller, such as local web access, etc. | [optional] 
**Issu** | Pointer to **int32** | Perform non-disruptive software upgrades, first upgrading the backup control board, then switching the primary and backup roles, and subsequently upgrading the other control board. Issu should be a value as follows: 1:ENABLE;0:DISABLE, effective only on the primary control board. | [optional] 
**Mac** | **string** | List of device keys for devices to be restarted. | 
**RebootId** | Pointer to **string** | Generated device batch restart ID. | [optional] 
**SaveCurrentConfig** | **int32** | Whether to save the current configuration. | 
**Slots** | Pointer to **[]int32** | Slot ID list should be within the range of 1 to 4 | [optional] 
**Status** | Pointer to **int32** | Status | [optional] 
**StatusCategory** | Pointer to **int32** | Status category.StatusCategory should be a value as follows:0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated. | [optional] 
**Token** | Pointer to **string** | User token | [optional] 

## Methods

### NewDeviceRebootAppDTO

`func NewDeviceRebootAppDTO(mac string, saveCurrentConfig int32, ) *DeviceRebootAppDTO`

NewDeviceRebootAppDTO instantiates a new DeviceRebootAppDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceRebootAppDTOWithDefaults

`func NewDeviceRebootAppDTOWithDefaults() *DeviceRebootAppDTO`

NewDeviceRebootAppDTOWithDefaults instantiates a new DeviceRebootAppDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetControlType

`func (o *DeviceRebootAppDTO) GetControlType() int32`

GetControlType returns the ControlType field if non-nil, zero value otherwise.

### GetControlTypeOk

`func (o *DeviceRebootAppDTO) GetControlTypeOk() (*int32, bool)`

GetControlTypeOk returns a tuple with the ControlType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetControlType

`func (o *DeviceRebootAppDTO) SetControlType(v int32)`

SetControlType sets ControlType field to given value.

### HasControlType

`func (o *DeviceRebootAppDTO) HasControlType() bool`

HasControlType returns a boolean if a field has been set.

### GetIssu

`func (o *DeviceRebootAppDTO) GetIssu() int32`

GetIssu returns the Issu field if non-nil, zero value otherwise.

### GetIssuOk

`func (o *DeviceRebootAppDTO) GetIssuOk() (*int32, bool)`

GetIssuOk returns a tuple with the Issu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssu

`func (o *DeviceRebootAppDTO) SetIssu(v int32)`

SetIssu sets Issu field to given value.

### HasIssu

`func (o *DeviceRebootAppDTO) HasIssu() bool`

HasIssu returns a boolean if a field has been set.

### GetMac

`func (o *DeviceRebootAppDTO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *DeviceRebootAppDTO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *DeviceRebootAppDTO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetRebootId

`func (o *DeviceRebootAppDTO) GetRebootId() string`

GetRebootId returns the RebootId field if non-nil, zero value otherwise.

### GetRebootIdOk

`func (o *DeviceRebootAppDTO) GetRebootIdOk() (*string, bool)`

GetRebootIdOk returns a tuple with the RebootId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootId

`func (o *DeviceRebootAppDTO) SetRebootId(v string)`

SetRebootId sets RebootId field to given value.

### HasRebootId

`func (o *DeviceRebootAppDTO) HasRebootId() bool`

HasRebootId returns a boolean if a field has been set.

### GetSaveCurrentConfig

`func (o *DeviceRebootAppDTO) GetSaveCurrentConfig() int32`

GetSaveCurrentConfig returns the SaveCurrentConfig field if non-nil, zero value otherwise.

### GetSaveCurrentConfigOk

`func (o *DeviceRebootAppDTO) GetSaveCurrentConfigOk() (*int32, bool)`

GetSaveCurrentConfigOk returns a tuple with the SaveCurrentConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaveCurrentConfig

`func (o *DeviceRebootAppDTO) SetSaveCurrentConfig(v int32)`

SetSaveCurrentConfig sets SaveCurrentConfig field to given value.


### GetSlots

`func (o *DeviceRebootAppDTO) GetSlots() []int32`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *DeviceRebootAppDTO) GetSlotsOk() (*[]int32, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *DeviceRebootAppDTO) SetSlots(v []int32)`

SetSlots sets Slots field to given value.

### HasSlots

`func (o *DeviceRebootAppDTO) HasSlots() bool`

HasSlots returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceRebootAppDTO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceRebootAppDTO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceRebootAppDTO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceRebootAppDTO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *DeviceRebootAppDTO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *DeviceRebootAppDTO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *DeviceRebootAppDTO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *DeviceRebootAppDTO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetToken

`func (o *DeviceRebootAppDTO) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *DeviceRebootAppDTO) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *DeviceRebootAppDTO) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *DeviceRebootAppDTO) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


