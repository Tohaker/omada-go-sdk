# IpsThreatOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activity** | Pointer to **string** | Activity. | [optional] 
**Archived** | Pointer to **bool** | Whether archived. | [optional] 
**Category** | Pointer to **int32** | Category. | [optional] 
**Classification** | Pointer to **string** | Classification. | [optional] 
**CreatTime** | Pointer to **int64** | Creat Time. | [optional] 
**DataUsage** | Pointer to **int64** | Data Usage. | [optional] 
**DstCountry** | Pointer to **string** | Destination Country. | [optional] 
**DstIp** | Pointer to **string** | Destination Ip. | [optional] 
**DstLatitude** | Pointer to **float64** | Destination Latitude(The precision is four decimal places). | [optional] 
**DstLongitude** | Pointer to **float64** | Destination Longitude(The precision is four decimal places). | [optional] 
**Id** | Pointer to **string** | ID. | [optional] 
**OmadacId** | Pointer to **string** | Omada ID | [optional] 
**Protocol** | Pointer to **string** | Protocol. | [optional] 
**Service** | Pointer to **string** | Service. | [optional] 
**Severity** | Pointer to **int32** | 0：critical, 1：major, 2：concerning 3：minor. | [optional] 
**Sid** | Pointer to **int64** | Sid. | [optional] 
**Signature** | Pointer to **string** | Signature. | [optional] 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**SiteName** | Pointer to **string** | Site Name. | [optional] 
**SrcCountry** | Pointer to **string** | Source Country. | [optional] 
**SrcIp** | Pointer to **string** | Source Ip. | [optional] 
**SrcLatitude** | Pointer to **float64** | Source Latitude(The precision is four decimal places). | [optional] 
**SrcLongitude** | Pointer to **float64** | Source Longitude(The precision is four decimal places). | [optional] 
**Time** | Pointer to **int64** | Time. | [optional] 

## Methods

### NewIpsThreatOpenApiVO

`func NewIpsThreatOpenApiVO() *IpsThreatOpenApiVO`

NewIpsThreatOpenApiVO instantiates a new IpsThreatOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsThreatOpenApiVOWithDefaults

`func NewIpsThreatOpenApiVOWithDefaults() *IpsThreatOpenApiVO`

NewIpsThreatOpenApiVOWithDefaults instantiates a new IpsThreatOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivity

`func (o *IpsThreatOpenApiVO) GetActivity() string`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *IpsThreatOpenApiVO) GetActivityOk() (*string, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *IpsThreatOpenApiVO) SetActivity(v string)`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *IpsThreatOpenApiVO) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetArchived

`func (o *IpsThreatOpenApiVO) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *IpsThreatOpenApiVO) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *IpsThreatOpenApiVO) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *IpsThreatOpenApiVO) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetCategory

`func (o *IpsThreatOpenApiVO) GetCategory() int32`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *IpsThreatOpenApiVO) GetCategoryOk() (*int32, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *IpsThreatOpenApiVO) SetCategory(v int32)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *IpsThreatOpenApiVO) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClassification

`func (o *IpsThreatOpenApiVO) GetClassification() string`

GetClassification returns the Classification field if non-nil, zero value otherwise.

### GetClassificationOk

`func (o *IpsThreatOpenApiVO) GetClassificationOk() (*string, bool)`

GetClassificationOk returns a tuple with the Classification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassification

`func (o *IpsThreatOpenApiVO) SetClassification(v string)`

SetClassification sets Classification field to given value.

### HasClassification

`func (o *IpsThreatOpenApiVO) HasClassification() bool`

HasClassification returns a boolean if a field has been set.

### GetCreatTime

`func (o *IpsThreatOpenApiVO) GetCreatTime() int64`

GetCreatTime returns the CreatTime field if non-nil, zero value otherwise.

### GetCreatTimeOk

`func (o *IpsThreatOpenApiVO) GetCreatTimeOk() (*int64, bool)`

GetCreatTimeOk returns a tuple with the CreatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatTime

`func (o *IpsThreatOpenApiVO) SetCreatTime(v int64)`

SetCreatTime sets CreatTime field to given value.

### HasCreatTime

`func (o *IpsThreatOpenApiVO) HasCreatTime() bool`

HasCreatTime returns a boolean if a field has been set.

### GetDataUsage

`func (o *IpsThreatOpenApiVO) GetDataUsage() int64`

GetDataUsage returns the DataUsage field if non-nil, zero value otherwise.

### GetDataUsageOk

`func (o *IpsThreatOpenApiVO) GetDataUsageOk() (*int64, bool)`

GetDataUsageOk returns a tuple with the DataUsage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataUsage

`func (o *IpsThreatOpenApiVO) SetDataUsage(v int64)`

SetDataUsage sets DataUsage field to given value.

### HasDataUsage

`func (o *IpsThreatOpenApiVO) HasDataUsage() bool`

HasDataUsage returns a boolean if a field has been set.

### GetDstCountry

`func (o *IpsThreatOpenApiVO) GetDstCountry() string`

GetDstCountry returns the DstCountry field if non-nil, zero value otherwise.

### GetDstCountryOk

`func (o *IpsThreatOpenApiVO) GetDstCountryOk() (*string, bool)`

GetDstCountryOk returns a tuple with the DstCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstCountry

`func (o *IpsThreatOpenApiVO) SetDstCountry(v string)`

SetDstCountry sets DstCountry field to given value.

### HasDstCountry

`func (o *IpsThreatOpenApiVO) HasDstCountry() bool`

HasDstCountry returns a boolean if a field has been set.

### GetDstIp

`func (o *IpsThreatOpenApiVO) GetDstIp() string`

GetDstIp returns the DstIp field if non-nil, zero value otherwise.

### GetDstIpOk

`func (o *IpsThreatOpenApiVO) GetDstIpOk() (*string, bool)`

GetDstIpOk returns a tuple with the DstIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstIp

`func (o *IpsThreatOpenApiVO) SetDstIp(v string)`

SetDstIp sets DstIp field to given value.

### HasDstIp

`func (o *IpsThreatOpenApiVO) HasDstIp() bool`

HasDstIp returns a boolean if a field has been set.

### GetDstLatitude

`func (o *IpsThreatOpenApiVO) GetDstLatitude() float64`

GetDstLatitude returns the DstLatitude field if non-nil, zero value otherwise.

### GetDstLatitudeOk

`func (o *IpsThreatOpenApiVO) GetDstLatitudeOk() (*float64, bool)`

GetDstLatitudeOk returns a tuple with the DstLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstLatitude

`func (o *IpsThreatOpenApiVO) SetDstLatitude(v float64)`

SetDstLatitude sets DstLatitude field to given value.

### HasDstLatitude

`func (o *IpsThreatOpenApiVO) HasDstLatitude() bool`

HasDstLatitude returns a boolean if a field has been set.

### GetDstLongitude

`func (o *IpsThreatOpenApiVO) GetDstLongitude() float64`

GetDstLongitude returns the DstLongitude field if non-nil, zero value otherwise.

### GetDstLongitudeOk

`func (o *IpsThreatOpenApiVO) GetDstLongitudeOk() (*float64, bool)`

GetDstLongitudeOk returns a tuple with the DstLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDstLongitude

`func (o *IpsThreatOpenApiVO) SetDstLongitude(v float64)`

SetDstLongitude sets DstLongitude field to given value.

### HasDstLongitude

`func (o *IpsThreatOpenApiVO) HasDstLongitude() bool`

HasDstLongitude returns a boolean if a field has been set.

### GetId

`func (o *IpsThreatOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IpsThreatOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IpsThreatOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IpsThreatOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOmadacId

`func (o *IpsThreatOpenApiVO) GetOmadacId() string`

GetOmadacId returns the OmadacId field if non-nil, zero value otherwise.

### GetOmadacIdOk

`func (o *IpsThreatOpenApiVO) GetOmadacIdOk() (*string, bool)`

GetOmadacIdOk returns a tuple with the OmadacId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOmadacId

`func (o *IpsThreatOpenApiVO) SetOmadacId(v string)`

SetOmadacId sets OmadacId field to given value.

### HasOmadacId

`func (o *IpsThreatOpenApiVO) HasOmadacId() bool`

HasOmadacId returns a boolean if a field has been set.

### GetProtocol

`func (o *IpsThreatOpenApiVO) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *IpsThreatOpenApiVO) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *IpsThreatOpenApiVO) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *IpsThreatOpenApiVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetService

`func (o *IpsThreatOpenApiVO) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *IpsThreatOpenApiVO) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *IpsThreatOpenApiVO) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *IpsThreatOpenApiVO) HasService() bool`

HasService returns a boolean if a field has been set.

### GetSeverity

`func (o *IpsThreatOpenApiVO) GetSeverity() int32`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *IpsThreatOpenApiVO) GetSeverityOk() (*int32, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *IpsThreatOpenApiVO) SetSeverity(v int32)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *IpsThreatOpenApiVO) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetSid

`func (o *IpsThreatOpenApiVO) GetSid() int64`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *IpsThreatOpenApiVO) GetSidOk() (*int64, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *IpsThreatOpenApiVO) SetSid(v int64)`

SetSid sets Sid field to given value.

### HasSid

`func (o *IpsThreatOpenApiVO) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetSignature

`func (o *IpsThreatOpenApiVO) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *IpsThreatOpenApiVO) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *IpsThreatOpenApiVO) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *IpsThreatOpenApiVO) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### GetSiteId

`func (o *IpsThreatOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *IpsThreatOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *IpsThreatOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *IpsThreatOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetSiteName

`func (o *IpsThreatOpenApiVO) GetSiteName() string`

GetSiteName returns the SiteName field if non-nil, zero value otherwise.

### GetSiteNameOk

`func (o *IpsThreatOpenApiVO) GetSiteNameOk() (*string, bool)`

GetSiteNameOk returns a tuple with the SiteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteName

`func (o *IpsThreatOpenApiVO) SetSiteName(v string)`

SetSiteName sets SiteName field to given value.

### HasSiteName

`func (o *IpsThreatOpenApiVO) HasSiteName() bool`

HasSiteName returns a boolean if a field has been set.

### GetSrcCountry

`func (o *IpsThreatOpenApiVO) GetSrcCountry() string`

GetSrcCountry returns the SrcCountry field if non-nil, zero value otherwise.

### GetSrcCountryOk

`func (o *IpsThreatOpenApiVO) GetSrcCountryOk() (*string, bool)`

GetSrcCountryOk returns a tuple with the SrcCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcCountry

`func (o *IpsThreatOpenApiVO) SetSrcCountry(v string)`

SetSrcCountry sets SrcCountry field to given value.

### HasSrcCountry

`func (o *IpsThreatOpenApiVO) HasSrcCountry() bool`

HasSrcCountry returns a boolean if a field has been set.

### GetSrcIp

`func (o *IpsThreatOpenApiVO) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *IpsThreatOpenApiVO) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *IpsThreatOpenApiVO) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *IpsThreatOpenApiVO) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetSrcLatitude

`func (o *IpsThreatOpenApiVO) GetSrcLatitude() float64`

GetSrcLatitude returns the SrcLatitude field if non-nil, zero value otherwise.

### GetSrcLatitudeOk

`func (o *IpsThreatOpenApiVO) GetSrcLatitudeOk() (*float64, bool)`

GetSrcLatitudeOk returns a tuple with the SrcLatitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcLatitude

`func (o *IpsThreatOpenApiVO) SetSrcLatitude(v float64)`

SetSrcLatitude sets SrcLatitude field to given value.

### HasSrcLatitude

`func (o *IpsThreatOpenApiVO) HasSrcLatitude() bool`

HasSrcLatitude returns a boolean if a field has been set.

### GetSrcLongitude

`func (o *IpsThreatOpenApiVO) GetSrcLongitude() float64`

GetSrcLongitude returns the SrcLongitude field if non-nil, zero value otherwise.

### GetSrcLongitudeOk

`func (o *IpsThreatOpenApiVO) GetSrcLongitudeOk() (*float64, bool)`

GetSrcLongitudeOk returns a tuple with the SrcLongitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcLongitude

`func (o *IpsThreatOpenApiVO) SetSrcLongitude(v float64)`

SetSrcLongitude sets SrcLongitude field to given value.

### HasSrcLongitude

`func (o *IpsThreatOpenApiVO) HasSrcLongitude() bool`

HasSrcLongitude returns a boolean if a field has been set.

### GetTime

`func (o *IpsThreatOpenApiVO) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *IpsThreatOpenApiVO) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *IpsThreatOpenApiVO) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *IpsThreatOpenApiVO) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


