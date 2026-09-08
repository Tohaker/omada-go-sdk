# ApDensityInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | The mac of ap. | [optional] 
**Model** | Pointer to **string** | The model of ap. | [optional] 
**ModelVersion** | Pointer to **string** | The modelVersion of ap. | [optional] 
**Name** | Pointer to **string** | The name of ap. | [optional] 
**NeighborApNum** | Pointer to **int32** | The neighbor ap number of ap. | [optional] 
**NeighborApPercent** | Pointer to **int32** | The neighbor ap percent | [optional] 
**NoData** | Pointer to **bool** | This flag indicates whether data exists. | [optional] 
**NoDataReason** | Pointer to **int32** | The no data reason of ap. | [optional] 
**Type** | Pointer to **string** | The type of ap. | [optional] 

## Methods

### NewApDensityInfoOpenApiVO

`func NewApDensityInfoOpenApiVO() *ApDensityInfoOpenApiVO`

NewApDensityInfoOpenApiVO instantiates a new ApDensityInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApDensityInfoOpenApiVOWithDefaults

`func NewApDensityInfoOpenApiVOWithDefaults() *ApDensityInfoOpenApiVO`

NewApDensityInfoOpenApiVOWithDefaults instantiates a new ApDensityInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *ApDensityInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApDensityInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApDensityInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApDensityInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ApDensityInfoOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApDensityInfoOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApDensityInfoOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApDensityInfoOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApDensityInfoOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApDensityInfoOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApDensityInfoOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApDensityInfoOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApDensityInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApDensityInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApDensityInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApDensityInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeighborApNum

`func (o *ApDensityInfoOpenApiVO) GetNeighborApNum() int32`

GetNeighborApNum returns the NeighborApNum field if non-nil, zero value otherwise.

### GetNeighborApNumOk

`func (o *ApDensityInfoOpenApiVO) GetNeighborApNumOk() (*int32, bool)`

GetNeighborApNumOk returns a tuple with the NeighborApNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborApNum

`func (o *ApDensityInfoOpenApiVO) SetNeighborApNum(v int32)`

SetNeighborApNum sets NeighborApNum field to given value.

### HasNeighborApNum

`func (o *ApDensityInfoOpenApiVO) HasNeighborApNum() bool`

HasNeighborApNum returns a boolean if a field has been set.

### GetNeighborApPercent

`func (o *ApDensityInfoOpenApiVO) GetNeighborApPercent() int32`

GetNeighborApPercent returns the NeighborApPercent field if non-nil, zero value otherwise.

### GetNeighborApPercentOk

`func (o *ApDensityInfoOpenApiVO) GetNeighborApPercentOk() (*int32, bool)`

GetNeighborApPercentOk returns a tuple with the NeighborApPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeighborApPercent

`func (o *ApDensityInfoOpenApiVO) SetNeighborApPercent(v int32)`

SetNeighborApPercent sets NeighborApPercent field to given value.

### HasNeighborApPercent

`func (o *ApDensityInfoOpenApiVO) HasNeighborApPercent() bool`

HasNeighborApPercent returns a boolean if a field has been set.

### GetNoData

`func (o *ApDensityInfoOpenApiVO) GetNoData() bool`

GetNoData returns the NoData field if non-nil, zero value otherwise.

### GetNoDataOk

`func (o *ApDensityInfoOpenApiVO) GetNoDataOk() (*bool, bool)`

GetNoDataOk returns a tuple with the NoData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoData

`func (o *ApDensityInfoOpenApiVO) SetNoData(v bool)`

SetNoData sets NoData field to given value.

### HasNoData

`func (o *ApDensityInfoOpenApiVO) HasNoData() bool`

HasNoData returns a boolean if a field has been set.

### GetNoDataReason

`func (o *ApDensityInfoOpenApiVO) GetNoDataReason() int32`

GetNoDataReason returns the NoDataReason field if non-nil, zero value otherwise.

### GetNoDataReasonOk

`func (o *ApDensityInfoOpenApiVO) GetNoDataReasonOk() (*int32, bool)`

GetNoDataReasonOk returns a tuple with the NoDataReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDataReason

`func (o *ApDensityInfoOpenApiVO) SetNoDataReason(v int32)`

SetNoDataReason sets NoDataReason field to given value.

### HasNoDataReason

`func (o *ApDensityInfoOpenApiVO) HasNoDataReason() bool`

HasNoDataReason returns a boolean if a field has been set.

### GetType

`func (o *ApDensityInfoOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApDensityInfoOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApDensityInfoOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApDensityInfoOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


