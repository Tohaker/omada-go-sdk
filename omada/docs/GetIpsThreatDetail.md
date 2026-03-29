# GetIpsThreatDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **int64** | Ips threat dataUsage, in Byte | [optional] 
**Category** | Pointer to **int32** | Ips threat category, Custom IDS/IPS categories list, if parameter[Dplevel] is 3, customCategories is needed.CustomCategories should be a list as follow: 1: Botcc, 2: Worm, 3: Malware, 4: Mobile_Malware, 6: P2P, 7: Tor, 8: Exploit, 9: Shellcode, 14: Activex, 15: DNS, 18: User Agents, 24: DShield | [optional] 
**Classification** | Pointer to **string** | Ips threat classification | [optional] 
**DstIp** | Pointer to **string** | IPS threat destination IP | [optional] 
**Id** | Pointer to **string** | IPS threat ID | [optional] 
**Protocol** | Pointer to **string** | Ips threat protocol | [optional] 
**Service** | Pointer to **string** | Ips threat description | [optional] 
**Severity** | Pointer to **int32** | Ips threat severity,  0：Critical, 1：Major, 2：Moderate 3：Minor 4:Low | [optional] 
**SrcCountry** | Pointer to **string** | IPS threat destination Country | [optional] 
**SrcIp** | Pointer to **string** | IPS threat source IP | [optional] 
**Time** | Pointer to **int64** | Timestamp, in seconds, such as 1682000000 | [optional] 

## Methods

### NewGetIpsThreatDetail

`func NewGetIpsThreatDetail() *GetIpsThreatDetail`

NewGetIpsThreatDetail instantiates a new GetIpsThreatDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIpsThreatDetailWithDefaults

`func NewGetIpsThreatDetailWithDefaults() *GetIpsThreatDetail`

NewGetIpsThreatDetailWithDefaults instantiates a new GetIpsThreatDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *GetIpsThreatDetail) GetActivity() int64`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *GetIpsThreatDetail) GetActivityOk() (*int64, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *GetIpsThreatDetail) SetActivity(v int64)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *GetIpsThreatDetail) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetCategory

`func (o *GetIpsThreatDetail) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetIpsThreatDetail) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetIpsThreatDetail) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetIpsThreatDetail) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClassification

`func (o *GetIpsThreatDetail) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *GetIpsThreatDetail) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *GetIpsThreatDetail) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *GetIpsThreatDetail) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetDstIp

`func (o *GetIpsThreatDetail) GetDstIp() string`

GetDstIp returns the DstIp field if non-nil, zero value otherwise.

### GetDstIpOk

`func (o *GetIpsThreatDetail) GetDstIpOk() (*string, bool)`

GetDstIpOk returns a tuple with the DstIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstIp

`func (o *GetIpsThreatDetail) SetDstIp(v string)`

SetDstIp sets DstIp field to given value.

### HasDstIp

`func (o *GetIpsThreatDetail) HasDstIp() bool`

HasDstIp returns a boolean if a field has been set.

### GetId

`func (o *GetIpsThreatDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIpsThreatDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIpsThreatDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetIpsThreatDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProtocol

`func (o *GetIpsThreatDetail) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GetIpsThreatDetail) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GetIpsThreatDetail) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GetIpsThreatDetail) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetService

`func (o *GetIpsThreatDetail) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetIpsThreatDetail) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetIpsThreatDetail) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetIpsThreatDetail) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSeverity

`func (o *GetIpsThreatDetail) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetIpsThreatDetail) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetIpsThreatDetail) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetIpsThreatDetail) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSrcCountry

`func (o *GetIpsThreatDetail) GetSrcCountry() string`

GetSrcCountry returns the SrcCountry field if non-nil, zero value otherwise.

### GetSrcCountryOk

`func (o *GetIpsThreatDetail) GetSrcCountryOk() (*string, bool)`

GetSrcCountryOk returns a tuple with the SrcCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCountry

`func (o *GetIpsThreatDetail) SetSrcCountry(v string)`

SetSrcCountry sets SrcCountry field to given value.

### HasSrcCountry

`func (o *GetIpsThreatDetail) HasSrcCountry() bool`

HasSrcCountry returns a boolean if a field has been set.

### GetSrcIp

`func (o *GetIpsThreatDetail) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *GetIpsThreatDetail) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *GetIpsThreatDetail) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *GetIpsThreatDetail) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetTime

`func (o *GetIpsThreatDetail) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *GetIpsThreatDetail) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *GetIpsThreatDetail) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *GetIpsThreatDetail) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


