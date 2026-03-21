# IspScanResultOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IspList** | Pointer to [**[]IspResultOpenApiVO**](IspResultOpenApiVO.md) | The list of the isp, contains the ISP name and the ID and the state of the ISP. | [optional] 
**Status** | Pointer to **int32** | The status of the band scan: 0 - Failed, 1 - Succeeded, 2 - Scanning. | [optional] 

## Methods

### NewIspScanResultOpenApiVO

`func NewIspScanResultOpenApiVO() *IspScanResultOpenApiVO`

NewIspScanResultOpenApiVO instantiates a new IspScanResultOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIspScanResultOpenApiVOWithDefaults

`func NewIspScanResultOpenApiVOWithDefaults() *IspScanResultOpenApiVO`

NewIspScanResultOpenApiVOWithDefaults instantiates a new IspScanResultOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIspList

`func (o *IspScanResultOpenApiVO) GetIspList() []IspResultOpenApiVO`

GetIspList returns the IspList field if non-nil, zero value otherwise.

### GetIspListOk

`func (o *IspScanResultOpenApiVO) GetIspListOk() (*[]IspResultOpenApiVO, bool)`

GetIspListOk returns a tuple with the IspList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspList

`func (o *IspScanResultOpenApiVO) SetIspList(v []IspResultOpenApiVO)`

SetIspList sets IspList field to given value.

### HasIspList

`func (o *IspScanResultOpenApiVO) HasIspList() bool`

HasIspList returns a boolean if a field has been set.

### GetStatus

`func (o *IspScanResultOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IspScanResultOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IspScanResultOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IspScanResultOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


