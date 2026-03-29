# ApMeshStatisticsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CandidateParents** | Pointer to [**[]CandidateParentOpenApiVO**](CandidateParentOpenApiVO.md) | candidate parent aps info | [optional] 
**ChildAps** | Pointer to [**[]ChildApOpenApiVO**](ChildApOpenApiVO.md) | List of child aps | [optional] 
**ScanStatus** | Pointer to **int32** | 0-init; 1-scanning; 2-scan success; 3-fail | [optional] 
**Status** | Pointer to **int32** | ap status | [optional] 
**StatusCategory** | Pointer to **int32** | ap status category | [optional] 
**WirelessLinked** | Pointer to **bool** | whether ap is wireless linked | [optional] 
**WirelessUplink** | Pointer to [**ApWirelessUplink**](ApWirelessUplink.md) |  | [optional] 

## Methods

### NewApMeshStatisticsOpenApiVO

`func NewApMeshStatisticsOpenApiVO() *ApMeshStatisticsOpenApiVO`

NewApMeshStatisticsOpenApiVO instantiates a new ApMeshStatisticsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApMeshStatisticsOpenApiVOWithDefaults

`func NewApMeshStatisticsOpenApiVOWithDefaults() *ApMeshStatisticsOpenApiVO`

NewApMeshStatisticsOpenApiVOWithDefaults instantiates a new ApMeshStatisticsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidateParents

`func (o *ApMeshStatisticsOpenApiVO) GetCandidateParents() []CandidateParentOpenApiVO`

GetCandidateParents returns the CandidateParents field if non-nil, zero value otherwise.

### GetCandidateParentsOk

`func (o *ApMeshStatisticsOpenApiVO) GetCandidateParentsOk() (*[]CandidateParentOpenApiVO, bool)`

GetCandidateParentsOk returns a tuple with the CandidateParents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidateParents

`func (o *ApMeshStatisticsOpenApiVO) SetCandidateParents(v []CandidateParentOpenApiVO)`

SetCandidateParents sets CandidateParents field to given value.

### HasCandidateParents

`func (o *ApMeshStatisticsOpenApiVO) HasCandidateParents() bool`

HasCandidateParents returns a boolean if a field has been set.

### GetChildAps

`func (o *ApMeshStatisticsOpenApiVO) GetChildAps() []ChildApOpenApiVO`

GetChildAps returns the ChildAps field if non-nil, zero value otherwise.

### GetChildApsOk

`func (o *ApMeshStatisticsOpenApiVO) GetChildApsOk() (*[]ChildApOpenApiVO, bool)`

GetChildApsOk returns a tuple with the ChildAps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildAps

`func (o *ApMeshStatisticsOpenApiVO) SetChildAps(v []ChildApOpenApiVO)`

SetChildAps sets ChildAps field to given value.

### HasChildAps

`func (o *ApMeshStatisticsOpenApiVO) HasChildAps() bool`

HasChildAps returns a boolean if a field has been set.

### GetScanStatus

`func (o *ApMeshStatisticsOpenApiVO) GetScanStatus() int32`

GetScanStatus returns the ScanStatus field if non-nil, zero value otherwise.

### GetScanStatusOk

`func (o *ApMeshStatisticsOpenApiVO) GetScanStatusOk() (*int32, bool)`

GetScanStatusOk returns a tuple with the ScanStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScanStatus

`func (o *ApMeshStatisticsOpenApiVO) SetScanStatus(v int32)`

SetScanStatus sets ScanStatus field to given value.

### HasScanStatus

`func (o *ApMeshStatisticsOpenApiVO) HasScanStatus() bool`

HasScanStatus returns a boolean if a field has been set.

### GetStatus

`func (o *ApMeshStatisticsOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApMeshStatisticsOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApMeshStatisticsOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApMeshStatisticsOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *ApMeshStatisticsOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *ApMeshStatisticsOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *ApMeshStatisticsOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *ApMeshStatisticsOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetWirelessLinked

`func (o *ApMeshStatisticsOpenApiVO) GetWirelessLinked() bool`

GetWirelessLinked returns the WirelessLinked field if non-nil, zero value otherwise.

### GetWirelessLinkedOk

`func (o *ApMeshStatisticsOpenApiVO) GetWirelessLinkedOk() (*bool, bool)`

GetWirelessLinkedOk returns a tuple with the WirelessLinked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessLinked

`func (o *ApMeshStatisticsOpenApiVO) SetWirelessLinked(v bool)`

SetWirelessLinked sets WirelessLinked field to given value.

### HasWirelessLinked

`func (o *ApMeshStatisticsOpenApiVO) HasWirelessLinked() bool`

HasWirelessLinked returns a boolean if a field has been set.

### GetWirelessUplink

`func (o *ApMeshStatisticsOpenApiVO) GetWirelessUplink() ApWirelessUplink`

GetWirelessUplink returns the WirelessUplink field if non-nil, zero value otherwise.

### GetWirelessUplinkOk

`func (o *ApMeshStatisticsOpenApiVO) GetWirelessUplinkOk() (*ApWirelessUplink, bool)`

GetWirelessUplinkOk returns a tuple with the WirelessUplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessUplink

`func (o *ApMeshStatisticsOpenApiVO) SetWirelessUplink(v ApWirelessUplink)`

SetWirelessUplink sets WirelessUplink field to given value.

### HasWirelessUplink

`func (o *ApMeshStatisticsOpenApiVO) HasWirelessUplink() bool`

HasWirelessUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


