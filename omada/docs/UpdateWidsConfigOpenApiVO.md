# UpdateWidsConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detection** | Pointer to **[]int32** | Wireless IDS detection type, the value is returned only when level is custom(2); It should be a value as follows: 0: Signature_disassociation_broadcast; 1: Signature_deauth_broadcast; 2: Detect_apspoofing; 3: Detect_adhoc_using_valid_ssid; 4: Detect_malformed_large_duration; 5: Detect_overflow_eapol_key; 6: Detect_ap_impersonation; 7: Detect_ht_greenfield; 8: Detect_incomplete_ie; 9: Detect_malformed_htie; 10: Detect_malformed_frame_auth; 11: Detect_malformed_assoc_req; 12: Detect_valid_ssid_misuse; 13: Detect_adhoc_network; 14: Detect_client_flood; 15: Detect_hotspotter_attack; 16: Detect_power_save_dos_flood_attack; 17: Detect_violence_break. | [optional] 
**Level** | **int32** | Wireless IDS detection level; It should be a value as follows: 0:High; 1:Low; 2:Custom. | 
**Status** | **bool** | Wireless IDS config status; true:enable, false:disable. | 

## Methods

### NewUpdateWidsConfigOpenApiVO

`func NewUpdateWidsConfigOpenApiVO(level int32, status bool, ) *UpdateWidsConfigOpenApiVO`

NewUpdateWidsConfigOpenApiVO instantiates a new UpdateWidsConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWidsConfigOpenApiVOWithDefaults

`func NewUpdateWidsConfigOpenApiVOWithDefaults() *UpdateWidsConfigOpenApiVO`

NewUpdateWidsConfigOpenApiVOWithDefaults instantiates a new UpdateWidsConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetection

`func (o *UpdateWidsConfigOpenApiVO) GetDetection() []int32`

GetDetection returns the Detection field if non-nil, zero value otherwise.

### GetDetectionOk

`func (o *UpdateWidsConfigOpenApiVO) GetDetectionOk() (*[]int32, bool)`

GetDetectionOk returns a tuple with the Detection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetection

`func (o *UpdateWidsConfigOpenApiVO) SetDetection(v []int32)`

SetDetection sets Detection field to given value.

### HasDetection

`func (o *UpdateWidsConfigOpenApiVO) HasDetection() bool`

HasDetection returns a boolean if a field has been set.

### GetLevel

`func (o *UpdateWidsConfigOpenApiVO) GetLevel() int32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *UpdateWidsConfigOpenApiVO) GetLevelOk() (*int32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *UpdateWidsConfigOpenApiVO) SetLevel(v int32)`

SetLevel sets Level field to given value.


### GetStatus

`func (o *UpdateWidsConfigOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateWidsConfigOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateWidsConfigOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


