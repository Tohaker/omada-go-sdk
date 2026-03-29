# ExportGlobalDeviceListOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayColumns** | Pointer to **[]string** | In mode 1, the displayColumns cannot be empty. Allowed column name for input : [&#39;deviceName&#39;, &#39;serialNumber&#39;, &#39;macAddress&#39;, &#39;ipAddress&#39;, &#39;publicIpAddress&#39;,  &#39;ipv6Address&#39;, &#39;status&#39;, &#39;model&#39;, &#39;version&#39;, &#39;uplink&#39;, &#39;downlink&#39;, &#39;uptime&#39;, &#39;memoryUsage&#39;, &#39;cpuUsage&#39;,  &#39;clients&#39;, &#39;downloadTraffic&#39;, &#39;uploadTraffic&#39;, &#39;loopback&#39;, &#39;hop&#39;, &#39;wlanGroup&#39;, &#39;override&#39;,  &#39;radio2G&#39;, &#39;radio5G&#39;, &#39;radio6G&#39;, &#39;2gClients&#39;, &#39;5gClients&#39;, &#39;6gClients&#39;, &#39;bSsid&#39;, &#39;txRate&#39;, &#39;rxRate&#39;, &#39;channel&#39;, &#39;channelUtilization2G&#39;, &#39;channelUtilization5G&#39;, &#39;channelUtilization6G&#39;, &#39;tags&#39;, &#39;lastSeen&#39;,  &#39;configurationResult&#39;, &#39;licenseStatus&#39;(pro controller exclusive), &#39;dueTime&#39;(pro controller exclusive), &#39;series&#39;(pro controller exclusive), &#39;health&#39;(pro controller exclusive), &#39;stackGroup&#39;(site view exclusive),  &#39;uplinkDeviceName&#39;(site view exclusive), &#39;uplinkDevicePort&#39;(site view exclusive),  &#39;linkSpeed&#39;(site view exclusive), &#39;duplex&#39;(site view exclusive)] | [optional] 
**Format** | **int32** | Format should be a value as follows. 0 means CSV , 1 means XLSX. | 
**Mode** | **int32** | Value of mode should be 0 or 1 or 2. 0 : All columns, 1 : Customize the export columns, 2 : Default columns. | 
**QueryData** | Pointer to [**OpenApiQueryDataVO**](OpenApiQueryDataVO.md) |  | [optional] 

## Methods

### NewExportGlobalDeviceListOpenApiVO

`func NewExportGlobalDeviceListOpenApiVO(format int32, mode int32, ) *ExportGlobalDeviceListOpenApiVO`

NewExportGlobalDeviceListOpenApiVO instantiates a new ExportGlobalDeviceListOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportGlobalDeviceListOpenApiVOWithDefaults

`func NewExportGlobalDeviceListOpenApiVOWithDefaults() *ExportGlobalDeviceListOpenApiVO`

NewExportGlobalDeviceListOpenApiVOWithDefaults instantiates a new ExportGlobalDeviceListOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayColumns

`func (o *ExportGlobalDeviceListOpenApiVO) GetDisplayColumns() []string`

GetDisplayColumns returns the DisplayColumns field if non-nil, zero value otherwise.

### GetDisplayColumnsOk

`func (o *ExportGlobalDeviceListOpenApiVO) GetDisplayColumnsOk() (*[]string, bool)`

GetDisplayColumnsOk returns a tuple with the DisplayColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayColumns

`func (o *ExportGlobalDeviceListOpenApiVO) SetDisplayColumns(v []string)`

SetDisplayColumns sets DisplayColumns field to given value.

### HasDisplayColumns

`func (o *ExportGlobalDeviceListOpenApiVO) HasDisplayColumns() bool`

HasDisplayColumns returns a boolean if a field has been set.

### GetFormat

`func (o *ExportGlobalDeviceListOpenApiVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportGlobalDeviceListOpenApiVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportGlobalDeviceListOpenApiVO) SetFormat(v int32)`

SetFormat sets Format field to given value.


### GetMode

`func (o *ExportGlobalDeviceListOpenApiVO) GetMode() int32`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *ExportGlobalDeviceListOpenApiVO) GetModeOk() (*int32, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *ExportGlobalDeviceListOpenApiVO) SetMode(v int32)`

SetMode sets Mode field to given value.


### GetQueryData

`func (o *ExportGlobalDeviceListOpenApiVO) GetQueryData() OpenApiQueryDataVO`

GetQueryData returns the QueryData field if non-nil, zero value otherwise.

### GetQueryDataOk

`func (o *ExportGlobalDeviceListOpenApiVO) GetQueryDataOk() (*OpenApiQueryDataVO, bool)`

GetQueryDataOk returns a tuple with the QueryData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueryData

`func (o *ExportGlobalDeviceListOpenApiVO) SetQueryData(v OpenApiQueryDataVO)`

SetQueryData sets QueryData field to given value.

### HasQueryData

`func (o *ExportGlobalDeviceListOpenApiVO) HasQueryData() bool`

HasQueryData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


