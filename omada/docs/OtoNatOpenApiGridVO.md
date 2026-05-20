# OtoNatOpenApiGridVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** | Current page number. | [optional] 
**CurrentSize** | Pointer to **int32** | Number of entries per page. | [optional] 
**Data** | Pointer to [**[]OtoNatInfoOpenApiVO**](OtoNatInfoOpenApiVO.md) |  | [optional] 
**SupportGeneralDialingTypeWan** | Pointer to **bool** | Whether to support Dynamic IP, Static IP, PPPoE, L2TP, PPTP dialing type wan port | [optional] 
**TotalRows** | Pointer to **int64** | Total rows of all items. | [optional] 

## Methods

### NewOtoNatOpenApiGridVO

`func NewOtoNatOpenApiGridVO() *OtoNatOpenApiGridVO`

NewOtoNatOpenApiGridVO instantiates a new OtoNatOpenApiGridVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOtoNatOpenApiGridVOWithDefaults

`func NewOtoNatOpenApiGridVOWithDefaults() *OtoNatOpenApiGridVO`

NewOtoNatOpenApiGridVOWithDefaults instantiates a new OtoNatOpenApiGridVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *OtoNatOpenApiGridVO) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *OtoNatOpenApiGridVO) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *OtoNatOpenApiGridVO) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *OtoNatOpenApiGridVO) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetCurrentSize

`func (o *OtoNatOpenApiGridVO) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *OtoNatOpenApiGridVO) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *OtoNatOpenApiGridVO) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.

### HasCurrentSize

`func (o *OtoNatOpenApiGridVO) HasCurrentSize() bool`

HasCurrentSize returns a boolean if a field has been set.

### GetData

`func (o *OtoNatOpenApiGridVO) GetData() []OtoNatInfoOpenApiVO`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *OtoNatOpenApiGridVO) GetDataOk() (*[]OtoNatInfoOpenApiVO, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *OtoNatOpenApiGridVO) SetData(v []OtoNatInfoOpenApiVO)`

SetData sets Data field to given value.

### HasData

`func (o *OtoNatOpenApiGridVO) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportGeneralDialingTypeWan

`func (o *OtoNatOpenApiGridVO) GetSupportGeneralDialingTypeWan() bool`

GetSupportGeneralDialingTypeWan returns the SupportGeneralDialingTypeWan field if non-nil, zero value otherwise.

### GetSupportGeneralDialingTypeWanOk

`func (o *OtoNatOpenApiGridVO) GetSupportGeneralDialingTypeWanOk() (*bool, bool)`

GetSupportGeneralDialingTypeWanOk returns a tuple with the SupportGeneralDialingTypeWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportGeneralDialingTypeWan

`func (o *OtoNatOpenApiGridVO) SetSupportGeneralDialingTypeWan(v bool)`

SetSupportGeneralDialingTypeWan sets SupportGeneralDialingTypeWan field to given value.

### HasSupportGeneralDialingTypeWan

`func (o *OtoNatOpenApiGridVO) HasSupportGeneralDialingTypeWan() bool`

HasSupportGeneralDialingTypeWan returns a boolean if a field has been set.

### GetTotalRows

`func (o *OtoNatOpenApiGridVO) GetTotalRows() int64`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *OtoNatOpenApiGridVO) GetTotalRowsOk() (*int64, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *OtoNatOpenApiGridVO) SetTotalRows(v int64)`

SetTotalRows sets TotalRows field to given value.

### HasTotalRows

`func (o *OtoNatOpenApiGridVO) HasTotalRows() bool`

HasTotalRows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


