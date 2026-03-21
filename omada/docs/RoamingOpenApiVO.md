# RoamingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiRoamingEnable** | **bool** | Whether to enable AI roaming, this configuration will take effect only when fast roaming is enabled | 
**DualBand11kReportEnable** | Pointer to **bool** | Whether to enable 802.11k report. It has been deprecated. | [optional] 
**FastRoamingEnable** | **bool** | Whether to enable fast roaming | 
**ForceDisassociationEnable** | Pointer to **bool** | Whether to enable forced disassociation. Note: This field will no longer be supported since Omada Controller 5.14.24.40. | [optional] 
**NonStickRoamingEnable** | Pointer to **bool** | Whether to enable non-stick roaming | [optional] 

## Methods

### NewRoamingOpenApiVO

`func NewRoamingOpenApiVO(aiRoamingEnable bool, fastRoamingEnable bool, ) *RoamingOpenApiVO`

NewRoamingOpenApiVO instantiates a new RoamingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoamingOpenApiVOWithDefaults

`func NewRoamingOpenApiVOWithDefaults() *RoamingOpenApiVO`

NewRoamingOpenApiVOWithDefaults instantiates a new RoamingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiRoamingEnable

`func (o *RoamingOpenApiVO) GetAiRoamingEnable() bool`

GetAiRoamingEnable returns the AiRoamingEnable field if non-nil, zero value otherwise.

### GetAiRoamingEnableOk

`func (o *RoamingOpenApiVO) GetAiRoamingEnableOk() (*bool, bool)`

GetAiRoamingEnableOk returns a tuple with the AiRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiRoamingEnable

`func (o *RoamingOpenApiVO) SetAiRoamingEnable(v bool)`

SetAiRoamingEnable sets AiRoamingEnable field to given value.


### GetDualBand11kReportEnable

`func (o *RoamingOpenApiVO) GetDualBand11kReportEnable() bool`

GetDualBand11kReportEnable returns the DualBand11kReportEnable field if non-nil, zero value otherwise.

### GetDualBand11kReportEnableOk

`func (o *RoamingOpenApiVO) GetDualBand11kReportEnableOk() (*bool, bool)`

GetDualBand11kReportEnableOk returns a tuple with the DualBand11kReportEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDualBand11kReportEnable

`func (o *RoamingOpenApiVO) SetDualBand11kReportEnable(v bool)`

SetDualBand11kReportEnable sets DualBand11kReportEnable field to given value.

### HasDualBand11kReportEnable

`func (o *RoamingOpenApiVO) HasDualBand11kReportEnable() bool`

HasDualBand11kReportEnable returns a boolean if a field has been set.

### GetFastRoamingEnable

`func (o *RoamingOpenApiVO) GetFastRoamingEnable() bool`

GetFastRoamingEnable returns the FastRoamingEnable field if non-nil, zero value otherwise.

### GetFastRoamingEnableOk

`func (o *RoamingOpenApiVO) GetFastRoamingEnableOk() (*bool, bool)`

GetFastRoamingEnableOk returns a tuple with the FastRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastRoamingEnable

`func (o *RoamingOpenApiVO) SetFastRoamingEnable(v bool)`

SetFastRoamingEnable sets FastRoamingEnable field to given value.


### GetForceDisassociationEnable

`func (o *RoamingOpenApiVO) GetForceDisassociationEnable() bool`

GetForceDisassociationEnable returns the ForceDisassociationEnable field if non-nil, zero value otherwise.

### GetForceDisassociationEnableOk

`func (o *RoamingOpenApiVO) GetForceDisassociationEnableOk() (*bool, bool)`

GetForceDisassociationEnableOk returns a tuple with the ForceDisassociationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceDisassociationEnable

`func (o *RoamingOpenApiVO) SetForceDisassociationEnable(v bool)`

SetForceDisassociationEnable sets ForceDisassociationEnable field to given value.

### HasForceDisassociationEnable

`func (o *RoamingOpenApiVO) HasForceDisassociationEnable() bool`

HasForceDisassociationEnable returns a boolean if a field has been set.

### GetNonStickRoamingEnable

`func (o *RoamingOpenApiVO) GetNonStickRoamingEnable() bool`

GetNonStickRoamingEnable returns the NonStickRoamingEnable field if non-nil, zero value otherwise.

### GetNonStickRoamingEnableOk

`func (o *RoamingOpenApiVO) GetNonStickRoamingEnableOk() (*bool, bool)`

GetNonStickRoamingEnableOk returns a tuple with the NonStickRoamingEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonStickRoamingEnable

`func (o *RoamingOpenApiVO) SetNonStickRoamingEnable(v bool)`

SetNonStickRoamingEnable sets NonStickRoamingEnable field to given value.

### HasNonStickRoamingEnable

`func (o *RoamingOpenApiVO) HasNonStickRoamingEnable() bool`

HasNonStickRoamingEnable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


