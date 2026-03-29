# ApInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualChannel2g** | Pointer to **string** | Actual 2.4G channel of the device. | [optional] 
**ActualChannel5g** | Pointer to **string** | Actual 5G channel of the device. | [optional] 
**ActualChannel6g** | Pointer to **string** | Actual 6G channel of the device. | [optional] 
**BandWidth2g** | Pointer to **string** | 2.4G bandWidth of the device. | [optional] 
**BandWidth5g** | Pointer to **string** | 5G bandWidth of the device. | [optional] 
**BandWidth6g** | Pointer to **string** | 6G bandWidth of the device. | [optional] 
**HasResult** | Pointer to **bool** | Whether the result exists. | [optional] 
**Mac** | Pointer to **string** | Mac address | [optional] 
**Model** | Pointer to **string** | Model, such as EAP225. | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device,for example:3.0 | [optional] 
**Name** | Pointer to **string** | Default uses the MAC address as the name. | [optional] 
**Support2g** | Pointer to **bool** | Whether the Ap device support 2g. | [optional] 
**Support5g** | Pointer to **bool** | Whether the Ap device support 5g. | [optional] 
**Support6g** | Pointer to **bool** | Whether the Ap device support 6g. | [optional] 
**SupportChannelUtil** | Pointer to **bool** | Whether the Ap device support channel util detection. | [optional] 
**SupportNonWifiInterf** | Pointer to **bool** | Whether the Ap device support non-wifi interference detection. | [optional] 
**SupportReturnBandWidth** | Pointer to **bool** | Whether the Ap device support get band width by interference detection. | [optional] 
**SupportWifiInterf** | Pointer to **bool** | Whether the Ap device support wifi interference detection. | [optional] 
**Time** | Pointer to **int64** | Last active time. | [optional] 
**Type** | Pointer to **string** | Device type:ap、gateway、switch、olt | [optional] 

## Methods

### NewApInfoOpenApiVO

`func NewApInfoOpenApiVO() *ApInfoOpenApiVO`

NewApInfoOpenApiVO instantiates a new ApInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApInfoOpenApiVOWithDefaults

`func NewApInfoOpenApiVOWithDefaults() *ApInfoOpenApiVO`

NewApInfoOpenApiVOWithDefaults instantiates a new ApInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualChannel2g

`func (o *ApInfoOpenApiVO) GetActualChannel2g() string`

GetActualChannel2g returns the ActualChannel2g field if non-nil, zero value otherwise.

### GetActualChannel2gOk

`func (o *ApInfoOpenApiVO) GetActualChannel2gOk() (*string, bool)`

GetActualChannel2gOk returns a tuple with the ActualChannel2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualChannel2g

`func (o *ApInfoOpenApiVO) SetActualChannel2g(v string)`

SetActualChannel2g sets ActualChannel2g field to given value.

### HasActualChannel2g

`func (o *ApInfoOpenApiVO) HasActualChannel2g() bool`

HasActualChannel2g returns a boolean if a field has been set.

### GetActualChannel5g

`func (o *ApInfoOpenApiVO) GetActualChannel5g() string`

GetActualChannel5g returns the ActualChannel5g field if non-nil, zero value otherwise.

### GetActualChannel5gOk

`func (o *ApInfoOpenApiVO) GetActualChannel5gOk() (*string, bool)`

GetActualChannel5gOk returns a tuple with the ActualChannel5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualChannel5g

`func (o *ApInfoOpenApiVO) SetActualChannel5g(v string)`

SetActualChannel5g sets ActualChannel5g field to given value.

### HasActualChannel5g

`func (o *ApInfoOpenApiVO) HasActualChannel5g() bool`

HasActualChannel5g returns a boolean if a field has been set.

### GetActualChannel6g

`func (o *ApInfoOpenApiVO) GetActualChannel6g() string`

GetActualChannel6g returns the ActualChannel6g field if non-nil, zero value otherwise.

### GetActualChannel6gOk

`func (o *ApInfoOpenApiVO) GetActualChannel6gOk() (*string, bool)`

GetActualChannel6gOk returns a tuple with the ActualChannel6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualChannel6g

`func (o *ApInfoOpenApiVO) SetActualChannel6g(v string)`

SetActualChannel6g sets ActualChannel6g field to given value.

### HasActualChannel6g

`func (o *ApInfoOpenApiVO) HasActualChannel6g() bool`

HasActualChannel6g returns a boolean if a field has been set.

### GetBandWidth2g

`func (o *ApInfoOpenApiVO) GetBandWidth2g() string`

GetBandWidth2g returns the BandWidth2g field if non-nil, zero value otherwise.

### GetBandWidth2gOk

`func (o *ApInfoOpenApiVO) GetBandWidth2gOk() (*string, bool)`

GetBandWidth2gOk returns a tuple with the BandWidth2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth2g

`func (o *ApInfoOpenApiVO) SetBandWidth2g(v string)`

SetBandWidth2g sets BandWidth2g field to given value.

### HasBandWidth2g

`func (o *ApInfoOpenApiVO) HasBandWidth2g() bool`

HasBandWidth2g returns a boolean if a field has been set.

### GetBandWidth5g

`func (o *ApInfoOpenApiVO) GetBandWidth5g() string`

GetBandWidth5g returns the BandWidth5g field if non-nil, zero value otherwise.

### GetBandWidth5gOk

`func (o *ApInfoOpenApiVO) GetBandWidth5gOk() (*string, bool)`

GetBandWidth5gOk returns a tuple with the BandWidth5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth5g

`func (o *ApInfoOpenApiVO) SetBandWidth5g(v string)`

SetBandWidth5g sets BandWidth5g field to given value.

### HasBandWidth5g

`func (o *ApInfoOpenApiVO) HasBandWidth5g() bool`

HasBandWidth5g returns a boolean if a field has been set.

### GetBandWidth6g

`func (o *ApInfoOpenApiVO) GetBandWidth6g() string`

GetBandWidth6g returns the BandWidth6g field if non-nil, zero value otherwise.

### GetBandWidth6gOk

`func (o *ApInfoOpenApiVO) GetBandWidth6gOk() (*string, bool)`

GetBandWidth6gOk returns a tuple with the BandWidth6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandWidth6g

`func (o *ApInfoOpenApiVO) SetBandWidth6g(v string)`

SetBandWidth6g sets BandWidth6g field to given value.

### HasBandWidth6g

`func (o *ApInfoOpenApiVO) HasBandWidth6g() bool`

HasBandWidth6g returns a boolean if a field has been set.

### GetHasResult

`func (o *ApInfoOpenApiVO) GetHasResult() bool`

GetHasResult returns the HasResult field if non-nil, zero value otherwise.

### GetHasResultOk

`func (o *ApInfoOpenApiVO) GetHasResultOk() (*bool, bool)`

GetHasResultOk returns a tuple with the HasResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasResult

`func (o *ApInfoOpenApiVO) SetHasResult(v bool)`

SetHasResult sets HasResult field to given value.

### HasHasResult

`func (o *ApInfoOpenApiVO) HasHasResult() bool`

HasHasResult returns a boolean if a field has been set.

### GetMac

`func (o *ApInfoOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ApInfoOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ApInfoOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ApInfoOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ApInfoOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ApInfoOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ApInfoOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ApInfoOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ApInfoOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ApInfoOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ApInfoOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ApInfoOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ApInfoOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApInfoOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApInfoOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApInfoOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSupport2g

`func (o *ApInfoOpenApiVO) GetSupport2g() bool`

GetSupport2g returns the Support2g field if non-nil, zero value otherwise.

### GetSupport2gOk

`func (o *ApInfoOpenApiVO) GetSupport2gOk() (*bool, bool)`

GetSupport2gOk returns a tuple with the Support2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport2g

`func (o *ApInfoOpenApiVO) SetSupport2g(v bool)`

SetSupport2g sets Support2g field to given value.

### HasSupport2g

`func (o *ApInfoOpenApiVO) HasSupport2g() bool`

HasSupport2g returns a boolean if a field has been set.

### GetSupport5g

`func (o *ApInfoOpenApiVO) GetSupport5g() bool`

GetSupport5g returns the Support5g field if non-nil, zero value otherwise.

### GetSupport5gOk

`func (o *ApInfoOpenApiVO) GetSupport5gOk() (*bool, bool)`

GetSupport5gOk returns a tuple with the Support5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5g

`func (o *ApInfoOpenApiVO) SetSupport5g(v bool)`

SetSupport5g sets Support5g field to given value.

### HasSupport5g

`func (o *ApInfoOpenApiVO) HasSupport5g() bool`

HasSupport5g returns a boolean if a field has been set.

### GetSupport6g

`func (o *ApInfoOpenApiVO) GetSupport6g() bool`

GetSupport6g returns the Support6g field if non-nil, zero value otherwise.

### GetSupport6gOk

`func (o *ApInfoOpenApiVO) GetSupport6gOk() (*bool, bool)`

GetSupport6gOk returns a tuple with the Support6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport6g

`func (o *ApInfoOpenApiVO) SetSupport6g(v bool)`

SetSupport6g sets Support6g field to given value.

### HasSupport6g

`func (o *ApInfoOpenApiVO) HasSupport6g() bool`

HasSupport6g returns a boolean if a field has been set.

### GetSupportChannelUtil

`func (o *ApInfoOpenApiVO) GetSupportChannelUtil() bool`

GetSupportChannelUtil returns the SupportChannelUtil field if non-nil, zero value otherwise.

### GetSupportChannelUtilOk

`func (o *ApInfoOpenApiVO) GetSupportChannelUtilOk() (*bool, bool)`

GetSupportChannelUtilOk returns a tuple with the SupportChannelUtil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportChannelUtil

`func (o *ApInfoOpenApiVO) SetSupportChannelUtil(v bool)`

SetSupportChannelUtil sets SupportChannelUtil field to given value.

### HasSupportChannelUtil

`func (o *ApInfoOpenApiVO) HasSupportChannelUtil() bool`

HasSupportChannelUtil returns a boolean if a field has been set.

### GetSupportNonWifiInterf

`func (o *ApInfoOpenApiVO) GetSupportNonWifiInterf() bool`

GetSupportNonWifiInterf returns the SupportNonWifiInterf field if non-nil, zero value otherwise.

### GetSupportNonWifiInterfOk

`func (o *ApInfoOpenApiVO) GetSupportNonWifiInterfOk() (*bool, bool)`

GetSupportNonWifiInterfOk returns a tuple with the SupportNonWifiInterf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportNonWifiInterf

`func (o *ApInfoOpenApiVO) SetSupportNonWifiInterf(v bool)`

SetSupportNonWifiInterf sets SupportNonWifiInterf field to given value.

### HasSupportNonWifiInterf

`func (o *ApInfoOpenApiVO) HasSupportNonWifiInterf() bool`

HasSupportNonWifiInterf returns a boolean if a field has been set.

### GetSupportReturnBandWidth

`func (o *ApInfoOpenApiVO) GetSupportReturnBandWidth() bool`

GetSupportReturnBandWidth returns the SupportReturnBandWidth field if non-nil, zero value otherwise.

### GetSupportReturnBandWidthOk

`func (o *ApInfoOpenApiVO) GetSupportReturnBandWidthOk() (*bool, bool)`

GetSupportReturnBandWidthOk returns a tuple with the SupportReturnBandWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportReturnBandWidth

`func (o *ApInfoOpenApiVO) SetSupportReturnBandWidth(v bool)`

SetSupportReturnBandWidth sets SupportReturnBandWidth field to given value.

### HasSupportReturnBandWidth

`func (o *ApInfoOpenApiVO) HasSupportReturnBandWidth() bool`

HasSupportReturnBandWidth returns a boolean if a field has been set.

### GetSupportWifiInterf

`func (o *ApInfoOpenApiVO) GetSupportWifiInterf() bool`

GetSupportWifiInterf returns the SupportWifiInterf field if non-nil, zero value otherwise.

### GetSupportWifiInterfOk

`func (o *ApInfoOpenApiVO) GetSupportWifiInterfOk() (*bool, bool)`

GetSupportWifiInterfOk returns a tuple with the SupportWifiInterf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportWifiInterf

`func (o *ApInfoOpenApiVO) SetSupportWifiInterf(v bool)`

SetSupportWifiInterf sets SupportWifiInterf field to given value.

### HasSupportWifiInterf

`func (o *ApInfoOpenApiVO) HasSupportWifiInterf() bool`

HasSupportWifiInterf returns a boolean if a field has been set.

### GetTime

`func (o *ApInfoOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApInfoOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApInfoOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ApInfoOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetType

`func (o *ApInfoOpenApiVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApInfoOpenApiVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApInfoOpenApiVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApInfoOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


