# EditVoucherGroupPatternOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationEnable** | Pointer to **bool** | Whether to print duration of the voucher | [optional] 
**LimitEnable** | Pointer to **bool** | Whether to print limit information of the voucher | [optional] 
**LogoPictureId** | Pointer to **string** | Logo picture ID | [optional] 
**LogoSize** | Pointer to **int32** | Logo size. It should be within the range of 50-175 | [optional] 
**NetworkList** | Pointer to **[]string** | Network list on the pattern of the voucher | [optional] 
**PatternType** | Pointer to **int32** | 0: Logo, 1: Title, 2: Disable | [optional] 
**Position** | Pointer to **int32** | 0: Left, 1: Right, 2: Middle | [optional] 
**PrintComments** | Pointer to **string** | Comments to print on the voucher | [optional] 
**SsidList** | Pointer to **[]string** | SSID list on the pattern of the voucher | [optional] 
**SsidNetworkenable** | Pointer to **bool** | Whether to print SSIDs and networks on the pattern of the voucher | [optional] 
**Title** | Pointer to **string** | Voucher title | [optional] 
**TitleSize** | Pointer to **int32** | Voucher title size. It should be within the range of 12-18 | [optional] 

## Methods

### NewEditVoucherGroupPatternOpenApiVO

`func NewEditVoucherGroupPatternOpenApiVO() *EditVoucherGroupPatternOpenApiVO`

NewEditVoucherGroupPatternOpenApiVO instantiates a new EditVoucherGroupPatternOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditVoucherGroupPatternOpenApiVOWithDefaults

`func NewEditVoucherGroupPatternOpenApiVOWithDefaults() *EditVoucherGroupPatternOpenApiVO`

NewEditVoucherGroupPatternOpenApiVOWithDefaults instantiates a new EditVoucherGroupPatternOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationEnable

`func (o *EditVoucherGroupPatternOpenApiVO) GetDurationEnable() bool`

GetDurationEnable returns the DurationEnable field if non-nil, zero value otherwise.

### GetDurationEnableOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetDurationEnableOk() (*bool, bool)`

GetDurationEnableOk returns a tuple with the DurationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationEnable

`func (o *EditVoucherGroupPatternOpenApiVO) SetDurationEnable(v bool)`

SetDurationEnable sets DurationEnable field to given value.

### HasDurationEnable

`func (o *EditVoucherGroupPatternOpenApiVO) HasDurationEnable() bool`

HasDurationEnable returns a boolean if a field has been set.

### GetLimitEnable

`func (o *EditVoucherGroupPatternOpenApiVO) GetLimitEnable() bool`

GetLimitEnable returns the LimitEnable field if non-nil, zero value otherwise.

### GetLimitEnableOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetLimitEnableOk() (*bool, bool)`

GetLimitEnableOk returns a tuple with the LimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitEnable

`func (o *EditVoucherGroupPatternOpenApiVO) SetLimitEnable(v bool)`

SetLimitEnable sets LimitEnable field to given value.

### HasLimitEnable

`func (o *EditVoucherGroupPatternOpenApiVO) HasLimitEnable() bool`

HasLimitEnable returns a boolean if a field has been set.

### GetLogoPictureId

`func (o *EditVoucherGroupPatternOpenApiVO) GetLogoPictureId() string`

GetLogoPictureId returns the LogoPictureId field if non-nil, zero value otherwise.

### GetLogoPictureIdOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetLogoPictureIdOk() (*string, bool)`

GetLogoPictureIdOk returns a tuple with the LogoPictureId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPictureId

`func (o *EditVoucherGroupPatternOpenApiVO) SetLogoPictureId(v string)`

SetLogoPictureId sets LogoPictureId field to given value.

### HasLogoPictureId

`func (o *EditVoucherGroupPatternOpenApiVO) HasLogoPictureId() bool`

HasLogoPictureId returns a boolean if a field has been set.

### GetLogoSize

`func (o *EditVoucherGroupPatternOpenApiVO) GetLogoSize() int32`

GetLogoSize returns the LogoSize field if non-nil, zero value otherwise.

### GetLogoSizeOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetLogoSizeOk() (*int32, bool)`

GetLogoSizeOk returns a tuple with the LogoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSize

`func (o *EditVoucherGroupPatternOpenApiVO) SetLogoSize(v int32)`

SetLogoSize sets LogoSize field to given value.

### HasLogoSize

`func (o *EditVoucherGroupPatternOpenApiVO) HasLogoSize() bool`

HasLogoSize returns a boolean if a field has been set.

### GetNetworkList

`func (o *EditVoucherGroupPatternOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *EditVoucherGroupPatternOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *EditVoucherGroupPatternOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetPatternType

`func (o *EditVoucherGroupPatternOpenApiVO) GetPatternType() int32`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetPatternTypeOk() (*int32, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *EditVoucherGroupPatternOpenApiVO) SetPatternType(v int32)`

SetPatternType sets PatternType field to given value.

### HasPatternType

`func (o *EditVoucherGroupPatternOpenApiVO) HasPatternType() bool`

HasPatternType returns a boolean if a field has been set.

### GetPosition

`func (o *EditVoucherGroupPatternOpenApiVO) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *EditVoucherGroupPatternOpenApiVO) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *EditVoucherGroupPatternOpenApiVO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPrintComments

`func (o *EditVoucherGroupPatternOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *EditVoucherGroupPatternOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *EditVoucherGroupPatternOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetSsidList

`func (o *EditVoucherGroupPatternOpenApiVO) GetSsidList() []string`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetSsidListOk() (*[]string, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *EditVoucherGroupPatternOpenApiVO) SetSsidList(v []string)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *EditVoucherGroupPatternOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.

### GetSsidNetworkenable

`func (o *EditVoucherGroupPatternOpenApiVO) GetSsidNetworkenable() bool`

GetSsidNetworkenable returns the SsidNetworkenable field if non-nil, zero value otherwise.

### GetSsidNetworkenableOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetSsidNetworkenableOk() (*bool, bool)`

GetSsidNetworkenableOk returns a tuple with the SsidNetworkenable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNetworkenable

`func (o *EditVoucherGroupPatternOpenApiVO) SetSsidNetworkenable(v bool)`

SetSsidNetworkenable sets SsidNetworkenable field to given value.

### HasSsidNetworkenable

`func (o *EditVoucherGroupPatternOpenApiVO) HasSsidNetworkenable() bool`

HasSsidNetworkenable returns a boolean if a field has been set.

### GetTitle

`func (o *EditVoucherGroupPatternOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *EditVoucherGroupPatternOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *EditVoucherGroupPatternOpenApiVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleSize

`func (o *EditVoucherGroupPatternOpenApiVO) GetTitleSize() int32`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *EditVoucherGroupPatternOpenApiVO) GetTitleSizeOk() (*int32, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *EditVoucherGroupPatternOpenApiVO) SetTitleSize(v int32)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *EditVoucherGroupPatternOpenApiVO) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


