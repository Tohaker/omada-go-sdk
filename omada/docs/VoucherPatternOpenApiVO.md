# VoucherPatternOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationEnable** | Pointer to **bool** | Whether to print duration of the voucher | [optional] 
**LimitEnable** | Pointer to **bool** | Whether to print limit info of the voucher. Limit info is limited usage count or limited online user number depending on the setting of the voucher group. | [optional] 
**LogoPicture** | Pointer to [**VoucherLogoVO**](VoucherLogoVO.md) |  | [optional] 
**LogoSize** | Pointer to **int32** | Logo size of the voucher pattern, used when pattern type is 0 | [optional] 
**NetworkList** | Pointer to **[]string** | IDs of the networks on the pattern of the voucher | [optional] 
**PatternCode** | Pointer to **string** | Voucher code of the previewed voucher pattern | [optional] 
**PatternType** | Pointer to **int32** | 0: Logo, 1: Title, 2: Disable | [optional] 
**Position** | Pointer to **int32** | 0: Left, 1: Right, 2: Middle | [optional] 
**PrintComments** | Pointer to **string** | Comments to print on the voucher | [optional] 
**SsidList** | Pointer to **[]string** | IDs of the SSIDs on the pattern of the voucher | [optional] 
**SsidNetworkEnable** | Pointer to **bool** | Whether to print SSIDs and networks on the pattern of the voucher | [optional] 
**SsidNetworkNameList** | Pointer to **[]string** | SSID and network name list on the pattern of the voucher | [optional] 
**Title** | Pointer to **string** | Title of the voucher pattern, used when pattern type is 1 | [optional] 
**TitleSize** | Pointer to **int32** | Title size of the voucher pattern, used when pattern type is 1 | [optional] 
**Validity** | Pointer to **string** | Validity information to print on the voucher | [optional] 

## Methods

### NewVoucherPatternOpenApiVO

`func NewVoucherPatternOpenApiVO() *VoucherPatternOpenApiVO`

NewVoucherPatternOpenApiVO instantiates a new VoucherPatternOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoucherPatternOpenApiVOWithDefaults

`func NewVoucherPatternOpenApiVOWithDefaults() *VoucherPatternOpenApiVO`

NewVoucherPatternOpenApiVOWithDefaults instantiates a new VoucherPatternOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationEnable

`func (o *VoucherPatternOpenApiVO) GetDurationEnable() bool`

GetDurationEnable returns the DurationEnable field if non-nil, zero value otherwise.

### GetDurationEnableOk

`func (o *VoucherPatternOpenApiVO) GetDurationEnableOk() (*bool, bool)`

GetDurationEnableOk returns a tuple with the DurationEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationEnable

`func (o *VoucherPatternOpenApiVO) SetDurationEnable(v bool)`

SetDurationEnable sets DurationEnable field to given value.

### HasDurationEnable

`func (o *VoucherPatternOpenApiVO) HasDurationEnable() bool`

HasDurationEnable returns a boolean if a field has been set.

### GetLimitEnable

`func (o *VoucherPatternOpenApiVO) GetLimitEnable() bool`

GetLimitEnable returns the LimitEnable field if non-nil, zero value otherwise.

### GetLimitEnableOk

`func (o *VoucherPatternOpenApiVO) GetLimitEnableOk() (*bool, bool)`

GetLimitEnableOk returns a tuple with the LimitEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitEnable

`func (o *VoucherPatternOpenApiVO) SetLimitEnable(v bool)`

SetLimitEnable sets LimitEnable field to given value.

### HasLimitEnable

`func (o *VoucherPatternOpenApiVO) HasLimitEnable() bool`

HasLimitEnable returns a boolean if a field has been set.

### GetLogoPicture

`func (o *VoucherPatternOpenApiVO) GetLogoPicture() VoucherLogoVO`

GetLogoPicture returns the LogoPicture field if non-nil, zero value otherwise.

### GetLogoPictureOk

`func (o *VoucherPatternOpenApiVO) GetLogoPictureOk() (*VoucherLogoVO, bool)`

GetLogoPictureOk returns a tuple with the LogoPicture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoPicture

`func (o *VoucherPatternOpenApiVO) SetLogoPicture(v VoucherLogoVO)`

SetLogoPicture sets LogoPicture field to given value.

### HasLogoPicture

`func (o *VoucherPatternOpenApiVO) HasLogoPicture() bool`

HasLogoPicture returns a boolean if a field has been set.

### GetLogoSize

`func (o *VoucherPatternOpenApiVO) GetLogoSize() int32`

GetLogoSize returns the LogoSize field if non-nil, zero value otherwise.

### GetLogoSizeOk

`func (o *VoucherPatternOpenApiVO) GetLogoSizeOk() (*int32, bool)`

GetLogoSizeOk returns a tuple with the LogoSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoSize

`func (o *VoucherPatternOpenApiVO) SetLogoSize(v int32)`

SetLogoSize sets LogoSize field to given value.

### HasLogoSize

`func (o *VoucherPatternOpenApiVO) HasLogoSize() bool`

HasLogoSize returns a boolean if a field has been set.

### GetNetworkList

`func (o *VoucherPatternOpenApiVO) GetNetworkList() []string`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *VoucherPatternOpenApiVO) GetNetworkListOk() (*[]string, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *VoucherPatternOpenApiVO) SetNetworkList(v []string)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *VoucherPatternOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetPatternCode

`func (o *VoucherPatternOpenApiVO) GetPatternCode() string`

GetPatternCode returns the PatternCode field if non-nil, zero value otherwise.

### GetPatternCodeOk

`func (o *VoucherPatternOpenApiVO) GetPatternCodeOk() (*string, bool)`

GetPatternCodeOk returns a tuple with the PatternCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternCode

`func (o *VoucherPatternOpenApiVO) SetPatternCode(v string)`

SetPatternCode sets PatternCode field to given value.

### HasPatternCode

`func (o *VoucherPatternOpenApiVO) HasPatternCode() bool`

HasPatternCode returns a boolean if a field has been set.

### GetPatternType

`func (o *VoucherPatternOpenApiVO) GetPatternType() int32`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *VoucherPatternOpenApiVO) GetPatternTypeOk() (*int32, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *VoucherPatternOpenApiVO) SetPatternType(v int32)`

SetPatternType sets PatternType field to given value.

### HasPatternType

`func (o *VoucherPatternOpenApiVO) HasPatternType() bool`

HasPatternType returns a boolean if a field has been set.

### GetPosition

`func (o *VoucherPatternOpenApiVO) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *VoucherPatternOpenApiVO) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *VoucherPatternOpenApiVO) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *VoucherPatternOpenApiVO) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetPrintComments

`func (o *VoucherPatternOpenApiVO) GetPrintComments() string`

GetPrintComments returns the PrintComments field if non-nil, zero value otherwise.

### GetPrintCommentsOk

`func (o *VoucherPatternOpenApiVO) GetPrintCommentsOk() (*string, bool)`

GetPrintCommentsOk returns a tuple with the PrintComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintComments

`func (o *VoucherPatternOpenApiVO) SetPrintComments(v string)`

SetPrintComments sets PrintComments field to given value.

### HasPrintComments

`func (o *VoucherPatternOpenApiVO) HasPrintComments() bool`

HasPrintComments returns a boolean if a field has been set.

### GetSsidList

`func (o *VoucherPatternOpenApiVO) GetSsidList() []string`

GetSsidList returns the SsidList field if non-nil, zero value otherwise.

### GetSsidListOk

`func (o *VoucherPatternOpenApiVO) GetSsidListOk() (*[]string, bool)`

GetSsidListOk returns a tuple with the SsidList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidList

`func (o *VoucherPatternOpenApiVO) SetSsidList(v []string)`

SetSsidList sets SsidList field to given value.

### HasSsidList

`func (o *VoucherPatternOpenApiVO) HasSsidList() bool`

HasSsidList returns a boolean if a field has been set.

### GetSsidNetworkEnable

`func (o *VoucherPatternOpenApiVO) GetSsidNetworkEnable() bool`

GetSsidNetworkEnable returns the SsidNetworkEnable field if non-nil, zero value otherwise.

### GetSsidNetworkEnableOk

`func (o *VoucherPatternOpenApiVO) GetSsidNetworkEnableOk() (*bool, bool)`

GetSsidNetworkEnableOk returns a tuple with the SsidNetworkEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNetworkEnable

`func (o *VoucherPatternOpenApiVO) SetSsidNetworkEnable(v bool)`

SetSsidNetworkEnable sets SsidNetworkEnable field to given value.

### HasSsidNetworkEnable

`func (o *VoucherPatternOpenApiVO) HasSsidNetworkEnable() bool`

HasSsidNetworkEnable returns a boolean if a field has been set.

### GetSsidNetworkNameList

`func (o *VoucherPatternOpenApiVO) GetSsidNetworkNameList() []string`

GetSsidNetworkNameList returns the SsidNetworkNameList field if non-nil, zero value otherwise.

### GetSsidNetworkNameListOk

`func (o *VoucherPatternOpenApiVO) GetSsidNetworkNameListOk() (*[]string, bool)`

GetSsidNetworkNameListOk returns a tuple with the SsidNetworkNameList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsidNetworkNameList

`func (o *VoucherPatternOpenApiVO) SetSsidNetworkNameList(v []string)`

SetSsidNetworkNameList sets SsidNetworkNameList field to given value.

### HasSsidNetworkNameList

`func (o *VoucherPatternOpenApiVO) HasSsidNetworkNameList() bool`

HasSsidNetworkNameList returns a boolean if a field has been set.

### GetTitle

`func (o *VoucherPatternOpenApiVO) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VoucherPatternOpenApiVO) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VoucherPatternOpenApiVO) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VoucherPatternOpenApiVO) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTitleSize

`func (o *VoucherPatternOpenApiVO) GetTitleSize() int32`

GetTitleSize returns the TitleSize field if non-nil, zero value otherwise.

### GetTitleSizeOk

`func (o *VoucherPatternOpenApiVO) GetTitleSizeOk() (*int32, bool)`

GetTitleSizeOk returns a tuple with the TitleSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleSize

`func (o *VoucherPatternOpenApiVO) SetTitleSize(v int32)`

SetTitleSize sets TitleSize field to given value.

### HasTitleSize

`func (o *VoucherPatternOpenApiVO) HasTitleSize() bool`

HasTitleSize returns a boolean if a field has been set.

### GetValidity

`func (o *VoucherPatternOpenApiVO) GetValidity() string`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *VoucherPatternOpenApiVO) GetValidityOk() (*string, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *VoucherPatternOpenApiVO) SetValidity(v string)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *VoucherPatternOpenApiVO) HasValidity() bool`

HasValidity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


