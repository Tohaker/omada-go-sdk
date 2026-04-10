# SnAddResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomerName** | Pointer to **string** | If device management records exist when adding devices by SN, you need to forget devices in this customer. Otherwise this field will have no value. | [optional] 
**DeviceKey** | Pointer to **string** | Device key(QR code) | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**Online** | Pointer to **bool** | Device online or offline | [optional] 
**SiteName** | Pointer to **string** | If device management records exist when adding devices by SN, you need to forget devices in this site. Otherwise this field will have no value. | [optional] 
**Sn** | Pointer to **string** | Serial number | [optional] 
**Status** | Pointer to **int32** | Device add status should be a value as follows: 0: success; 1: waiting to do; 2：importing; -51451：Device ID not found; -52201：Device is offline; -52202：Device is already bounded; -52208：Device is offline during processing; -52200：Device not exit; -53100：Invalid SN code; -53101：SN code already exists; -53102：Incorrect local Username/Password; -53103：Device error; -53113: Too many PRECONFIGURED devices in this account. Available licenses are not enough; -53114: This device has been added by another Controller; -53118: Omada Pro devices can only be added on the Omada Pro Controller; -53119: Omada devices can only be added on the Omada Controller; -39045: Failed to add the device due to duplicate SN codes. Please contact our TP-Link Support. | [optional] 
**SupportIppt** | Pointer to **bool** | Support ip pass-through or not。 | [optional] 

## Methods

### NewSnAddResultVO

`func NewSnAddResultVO() *SnAddResultVO`

NewSnAddResultVO instantiates a new SnAddResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnAddResultVOWithDefaults

`func NewSnAddResultVOWithDefaults() *SnAddResultVO`

NewSnAddResultVOWithDefaults instantiates a new SnAddResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomerName

`func (o *SnAddResultVO) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *SnAddResultVO) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *SnAddResultVO) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *SnAddResultVO) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetDeviceKey

`func (o *SnAddResultVO) GetDeviceKey() string`

GetDeviceKey returns the DeviceKey field if non-nil, zero value otherwise.

### GetDeviceKeyOk

`func (o *SnAddResultVO) GetDeviceKeyOk() (*string, bool)`

GetDeviceKeyOk returns a tuple with the DeviceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceKey

`func (o *SnAddResultVO) SetDeviceKey(v string)`

SetDeviceKey sets DeviceKey field to given value.

### HasDeviceKey

`func (o *SnAddResultVO) HasDeviceKey() bool`

HasDeviceKey returns a boolean if a field has been set.

### GetMac

`func (o *SnAddResultVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SnAddResultVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SnAddResultVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SnAddResultVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *SnAddResultVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SnAddResultVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SnAddResultVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SnAddResultVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnline

`func (o *SnAddResultVO) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *SnAddResultVO) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *SnAddResultVO) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *SnAddResultVO) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetSiteName

`func (o *SnAddResultVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *SnAddResultVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *SnAddResultVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *SnAddResultVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSn

`func (o *SnAddResultVO) GetSn() string`

GetSn returns the Sn field if non-nil, zero value otherwise.

### GetSnOk

`func (o *SnAddResultVO) GetSnOk() (*string, bool)`

GetSnOk returns a tuple with the Sn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSn

`func (o *SnAddResultVO) SetSn(v string)`

SetSn sets Sn field to given value.

### HasSn

`func (o *SnAddResultVO) HasSn() bool`

HasSn returns a boolean if a field has been set.

### GetStatus

`func (o *SnAddResultVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SnAddResultVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SnAddResultVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SnAddResultVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportIppt

`func (o *SnAddResultVO) GetSupportIppt() bool`

GetSupportIppt returns the SupportIppt field if non-nil, zero value otherwise.

### GetSupportIpptOk

`func (o *SnAddResultVO) GetSupportIpptOk() (*bool, bool)`

GetSupportIpptOk returns a tuple with the SupportIppt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportIppt

`func (o *SnAddResultVO) SetSupportIppt(v bool)`

SetSupportIppt sets SupportIppt field to given value.

### HasSupportIppt

`func (o *SnAddResultVO) HasSupportIppt() bool`

HasSupportIppt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


