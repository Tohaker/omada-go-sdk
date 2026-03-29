# CardInfoVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlertSummary** | Pointer to [**AlertSummaryTrendVO**](AlertSummaryTrendVO.md) |  | [optional] 
**ApHealth** | Pointer to [**[]DeviceHealthVO**](DeviceHealthVO.md) |  | [optional] 
**ApStatus** | Pointer to [**[]OnlineOfflineVO**](OnlineOfflineVO.md) |  | [optional] 
**AppCategories** | Pointer to [**AppCategoryTrafficsVO**](AppCategoryTrafficsVO.md) |  | [optional] 
**ClientConnectionTrend** | Pointer to [**ClientConnectionTrend**](ClientConnectionTrend.md) |  | [optional] 
**ClientHealthTrend** | Pointer to [**[]DeviceHealthVO**](DeviceHealthVO.md) |  | [optional] 
**ClientTraffic** | Pointer to [**ClientTrafficInfo**](ClientTrafficInfo.md) |  | [optional] 
**ClientsAssociationActivities** | Pointer to [**[]ClientAssociationActivities**](ClientAssociationActivities.md) | Clients association activities with time. | [optional] 
**ClientsOverview** | Pointer to [**ClientStatisticsOverview**](ClientStatisticsOverview.md) |  | [optional] 
**ClientsWithOnboardingTimes** | Pointer to [**ClientsWithOnBoardingTimes**](ClientsWithOnBoardingTimes.md) |  | [optional] 
**DeviceHealthTrend** | Pointer to [**[]DeviceHealthVO**](DeviceHealthVO.md) |  | [optional] 
**DeviceStatus** | Pointer to [**[]OnlineOfflineVO**](OnlineOfflineVO.md) |  | [optional] 
**GatewaySummary** | Pointer to [**GatewaySummaryVO**](GatewaySummaryVO.md) |  | [optional] 
**Internet** | Pointer to [**NetworkActivityVO**](NetworkActivityVO.md) |  | [optional] 
**IspLoad** | Pointer to [**[]IspLoadVO**](IspLoadVO.md) | ISP Load | [optional] 
**Network** | Pointer to [**NetworkVO**](NetworkVO.md) |  | [optional] 
**OverviewSummary** | Pointer to [**OverViewSummaryVO**](OverViewSummaryVO.md) |  | [optional] 
**PoePowerTrend** | Pointer to [**[]PoePowerTrendVO**](PoePowerTrendVO.md) |  | [optional] 
**SwitchAlertReboot** | Pointer to [**SwitchAlertRebootVO**](SwitchAlertRebootVO.md) |  | [optional] 
**SwitchHealthTrend** | Pointer to [**[]DeviceHealthVO**](DeviceHealthVO.md) |  | [optional] 
**SwitchStatus** | Pointer to [**[]OnlineOfflineVO**](OnlineOfflineVO.md) |  | [optional] 
**TopApByCpuAndMemory** | Pointer to [**ApUtilizationVO**](ApUtilizationVO.md) |  | [optional] 
**TopApByInterference** | Pointer to [**TopApByInterferenceVO**](TopApByInterferenceVO.md) |  | [optional] 
**TopApByRtAndDrop** | Pointer to [**TopApByRtDropVO**](TopApByRtDropVO.md) |  | [optional] 
**TopApByTrafficAndClient** | Pointer to [**TopApByTrafficAndClientVO**](TopApByTrafficAndClientVO.md) |  | [optional] 
**TopApplicationByTraffic** | Pointer to [**[]TopApplicationByTrafficVO**](TopApplicationByTrafficVO.md) |  | [optional] 
**TopClient** | Pointer to [**TopTrafficAndUptimeClients**](TopTrafficAndUptimeClients.md) |  | [optional] 
**TopSsidByTraffic** | Pointer to [**[]TopSsidTrafficVO**](TopSsidTrafficVO.md) |  | [optional] 
**TopSwitchByTrafficAndPoePower** | Pointer to [**TopSwitchVO**](TopSwitchVO.md) |  | [optional] 
**TopSwitchCpuMemory** | Pointer to [**SwitchUtilizationVO**](SwitchUtilizationVO.md) |  | [optional] 
**TrafficDistribution** | Pointer to [**TrafficDistributionVO**](TrafficDistributionVO.md) |  | [optional] 
**TrafficSummary** | Pointer to [**TrafficSummaryVO**](TrafficSummaryVO.md) |  | [optional] 
**WanHealthTrend** | Pointer to [**[]WanHealthTrendVO**](WanHealthTrendVO.md) |  | [optional] 
**WifiHealth** | Pointer to [**WifiHealthVO**](WifiHealthVO.md) |  | [optional] 
**WirelessTraffic** | Pointer to [**WirelessTrafficVO**](WirelessTrafficVO.md) |  | [optional] 
**WirelessTrafficSingle** | Pointer to [**WirelessTrafficSingleVO**](WirelessTrafficSingleVO.md) |  | [optional] 

## Methods

### NewCardInfoVO

`func NewCardInfoVO() *CardInfoVO`

NewCardInfoVO instantiates a new CardInfoVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardInfoVOWithDefaults

`func NewCardInfoVOWithDefaults() *CardInfoVO`

NewCardInfoVOWithDefaults instantiates a new CardInfoVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlertSummary

`func (o *CardInfoVO) GetAlertSummary() AlertSummaryTrendVO`

GetAlertSummary returns the AlertSummary field if non-nil, zero value otherwise.

### GetAlertSummaryOk

`func (o *CardInfoVO) GetAlertSummaryOk() (*AlertSummaryTrendVO, bool)`

GetAlertSummaryOk returns a tuple with the AlertSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlertSummary

`func (o *CardInfoVO) SetAlertSummary(v AlertSummaryTrendVO)`

SetAlertSummary sets AlertSummary field to given value.

### HasAlertSummary

`func (o *CardInfoVO) HasAlertSummary() bool`

HasAlertSummary returns a boolean if a field has been set.

### GetApHealth

`func (o *CardInfoVO) GetApHealth() []DeviceHealthVO`

GetApHealth returns the ApHealth field if non-nil, zero value otherwise.

### GetApHealthOk

`func (o *CardInfoVO) GetApHealthOk() (*[]DeviceHealthVO, bool)`

GetApHealthOk returns a tuple with the ApHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApHealth

`func (o *CardInfoVO) SetApHealth(v []DeviceHealthVO)`

SetApHealth sets ApHealth field to given value.

### HasApHealth

`func (o *CardInfoVO) HasApHealth() bool`

HasApHealth returns a boolean if a field has been set.

### GetApStatus

`func (o *CardInfoVO) GetApStatus() []OnlineOfflineVO`

GetApStatus returns the ApStatus field if non-nil, zero value otherwise.

### GetApStatusOk

`func (o *CardInfoVO) GetApStatusOk() (*[]OnlineOfflineVO, bool)`

GetApStatusOk returns a tuple with the ApStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApStatus

`func (o *CardInfoVO) SetApStatus(v []OnlineOfflineVO)`

SetApStatus sets ApStatus field to given value.

### HasApStatus

`func (o *CardInfoVO) HasApStatus() bool`

HasApStatus returns a boolean if a field has been set.

### GetAppCategories

`func (o *CardInfoVO) GetAppCategories() AppCategoryTrafficsVO`

GetAppCategories returns the AppCategories field if non-nil, zero value otherwise.

### GetAppCategoriesOk

`func (o *CardInfoVO) GetAppCategoriesOk() (*AppCategoryTrafficsVO, bool)`

GetAppCategoriesOk returns a tuple with the AppCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppCategories

`func (o *CardInfoVO) SetAppCategories(v AppCategoryTrafficsVO)`

SetAppCategories sets AppCategories field to given value.

### HasAppCategories

`func (o *CardInfoVO) HasAppCategories() bool`

HasAppCategories returns a boolean if a field has been set.

### GetClientConnectionTrend

`func (o *CardInfoVO) GetClientConnectionTrend() ClientConnectionTrend`

GetClientConnectionTrend returns the ClientConnectionTrend field if non-nil, zero value otherwise.

### GetClientConnectionTrendOk

`func (o *CardInfoVO) GetClientConnectionTrendOk() (*ClientConnectionTrend, bool)`

GetClientConnectionTrendOk returns a tuple with the ClientConnectionTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientConnectionTrend

`func (o *CardInfoVO) SetClientConnectionTrend(v ClientConnectionTrend)`

SetClientConnectionTrend sets ClientConnectionTrend field to given value.

### HasClientConnectionTrend

`func (o *CardInfoVO) HasClientConnectionTrend() bool`

HasClientConnectionTrend returns a boolean if a field has been set.

### GetClientHealthTrend

`func (o *CardInfoVO) GetClientHealthTrend() []DeviceHealthVO`

GetClientHealthTrend returns the ClientHealthTrend field if non-nil, zero value otherwise.

### GetClientHealthTrendOk

`func (o *CardInfoVO) GetClientHealthTrendOk() (*[]DeviceHealthVO, bool)`

GetClientHealthTrendOk returns a tuple with the ClientHealthTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientHealthTrend

`func (o *CardInfoVO) SetClientHealthTrend(v []DeviceHealthVO)`

SetClientHealthTrend sets ClientHealthTrend field to given value.

### HasClientHealthTrend

`func (o *CardInfoVO) HasClientHealthTrend() bool`

HasClientHealthTrend returns a boolean if a field has been set.

### GetClientTraffic

`func (o *CardInfoVO) GetClientTraffic() ClientTrafficInfo`

GetClientTraffic returns the ClientTraffic field if non-nil, zero value otherwise.

### GetClientTrafficOk

`func (o *CardInfoVO) GetClientTrafficOk() (*ClientTrafficInfo, bool)`

GetClientTrafficOk returns a tuple with the ClientTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientTraffic

`func (o *CardInfoVO) SetClientTraffic(v ClientTrafficInfo)`

SetClientTraffic sets ClientTraffic field to given value.

### HasClientTraffic

`func (o *CardInfoVO) HasClientTraffic() bool`

HasClientTraffic returns a boolean if a field has been set.

### GetClientsAssociationActivities

`func (o *CardInfoVO) GetClientsAssociationActivities() []ClientAssociationActivities`

GetClientsAssociationActivities returns the ClientsAssociationActivities field if non-nil, zero value otherwise.

### GetClientsAssociationActivitiesOk

`func (o *CardInfoVO) GetClientsAssociationActivitiesOk() (*[]ClientAssociationActivities, bool)`

GetClientsAssociationActivitiesOk returns a tuple with the ClientsAssociationActivities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsAssociationActivities

`func (o *CardInfoVO) SetClientsAssociationActivities(v []ClientAssociationActivities)`

SetClientsAssociationActivities sets ClientsAssociationActivities field to given value.

### HasClientsAssociationActivities

`func (o *CardInfoVO) HasClientsAssociationActivities() bool`

HasClientsAssociationActivities returns a boolean if a field has been set.

### GetClientsOverview

`func (o *CardInfoVO) GetClientsOverview() ClientStatisticsOverview`

GetClientsOverview returns the ClientsOverview field if non-nil, zero value otherwise.

### GetClientsOverviewOk

`func (o *CardInfoVO) GetClientsOverviewOk() (*ClientStatisticsOverview, bool)`

GetClientsOverviewOk returns a tuple with the ClientsOverview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsOverview

`func (o *CardInfoVO) SetClientsOverview(v ClientStatisticsOverview)`

SetClientsOverview sets ClientsOverview field to given value.

### HasClientsOverview

`func (o *CardInfoVO) HasClientsOverview() bool`

HasClientsOverview returns a boolean if a field has been set.

### GetClientsWithOnboardingTimes

`func (o *CardInfoVO) GetClientsWithOnboardingTimes() ClientsWithOnBoardingTimes`

GetClientsWithOnboardingTimes returns the ClientsWithOnboardingTimes field if non-nil, zero value otherwise.

### GetClientsWithOnboardingTimesOk

`func (o *CardInfoVO) GetClientsWithOnboardingTimesOk() (*ClientsWithOnBoardingTimes, bool)`

GetClientsWithOnboardingTimesOk returns a tuple with the ClientsWithOnboardingTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientsWithOnboardingTimes

`func (o *CardInfoVO) SetClientsWithOnboardingTimes(v ClientsWithOnBoardingTimes)`

SetClientsWithOnboardingTimes sets ClientsWithOnboardingTimes field to given value.

### HasClientsWithOnboardingTimes

`func (o *CardInfoVO) HasClientsWithOnboardingTimes() bool`

HasClientsWithOnboardingTimes returns a boolean if a field has been set.

### GetDeviceHealthTrend

`func (o *CardInfoVO) GetDeviceHealthTrend() []DeviceHealthVO`

GetDeviceHealthTrend returns the DeviceHealthTrend field if non-nil, zero value otherwise.

### GetDeviceHealthTrendOk

`func (o *CardInfoVO) GetDeviceHealthTrendOk() (*[]DeviceHealthVO, bool)`

GetDeviceHealthTrendOk returns a tuple with the DeviceHealthTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceHealthTrend

`func (o *CardInfoVO) SetDeviceHealthTrend(v []DeviceHealthVO)`

SetDeviceHealthTrend sets DeviceHealthTrend field to given value.

### HasDeviceHealthTrend

`func (o *CardInfoVO) HasDeviceHealthTrend() bool`

HasDeviceHealthTrend returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *CardInfoVO) GetDeviceStatus() []OnlineOfflineVO`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *CardInfoVO) GetDeviceStatusOk() (*[]OnlineOfflineVO, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *CardInfoVO) SetDeviceStatus(v []OnlineOfflineVO)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *CardInfoVO) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetGatewaySummary

`func (o *CardInfoVO) GetGatewaySummary() GatewaySummaryVO`

GetGatewaySummary returns the GatewaySummary field if non-nil, zero value otherwise.

### GetGatewaySummaryOk

`func (o *CardInfoVO) GetGatewaySummaryOk() (*GatewaySummaryVO, bool)`

GetGatewaySummaryOk returns a tuple with the GatewaySummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGatewaySummary

`func (o *CardInfoVO) SetGatewaySummary(v GatewaySummaryVO)`

SetGatewaySummary sets GatewaySummary field to given value.

### HasGatewaySummary

`func (o *CardInfoVO) HasGatewaySummary() bool`

HasGatewaySummary returns a boolean if a field has been set.

### GetInternet

`func (o *CardInfoVO) GetInternet() NetworkActivityVO`

GetInternet returns the Internet field if non-nil, zero value otherwise.

### GetInternetOk

`func (o *CardInfoVO) GetInternetOk() (*NetworkActivityVO, bool)`

GetInternetOk returns a tuple with the Internet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternet

`func (o *CardInfoVO) SetInternet(v NetworkActivityVO)`

SetInternet sets Internet field to given value.

### HasInternet

`func (o *CardInfoVO) HasInternet() bool`

HasInternet returns a boolean if a field has been set.

### GetIspLoad

`func (o *CardInfoVO) GetIspLoad() []IspLoadVO`

GetIspLoad returns the IspLoad field if non-nil, zero value otherwise.

### GetIspLoadOk

`func (o *CardInfoVO) GetIspLoadOk() (*[]IspLoadVO, bool)`

GetIspLoadOk returns a tuple with the IspLoad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIspLoad

`func (o *CardInfoVO) SetIspLoad(v []IspLoadVO)`

SetIspLoad sets IspLoad field to given value.

### HasIspLoad

`func (o *CardInfoVO) HasIspLoad() bool`

HasIspLoad returns a boolean if a field has been set.

### GetNetwork

`func (o *CardInfoVO) GetNetwork() NetworkVO`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CardInfoVO) GetNetworkOk() (*NetworkVO, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CardInfoVO) SetNetwork(v NetworkVO)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *CardInfoVO) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOverviewSummary

`func (o *CardInfoVO) GetOverviewSummary() OverViewSummaryVO`

GetOverviewSummary returns the OverviewSummary field if non-nil, zero value otherwise.

### GetOverviewSummaryOk

`func (o *CardInfoVO) GetOverviewSummaryOk() (*OverViewSummaryVO, bool)`

GetOverviewSummaryOk returns a tuple with the OverviewSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverviewSummary

`func (o *CardInfoVO) SetOverviewSummary(v OverViewSummaryVO)`

SetOverviewSummary sets OverviewSummary field to given value.

### HasOverviewSummary

`func (o *CardInfoVO) HasOverviewSummary() bool`

HasOverviewSummary returns a boolean if a field has been set.

### GetPoePowerTrend

`func (o *CardInfoVO) GetPoePowerTrend() []PoePowerTrendVO`

GetPoePowerTrend returns the PoePowerTrend field if non-nil, zero value otherwise.

### GetPoePowerTrendOk

`func (o *CardInfoVO) GetPoePowerTrendOk() (*[]PoePowerTrendVO, bool)`

GetPoePowerTrendOk returns a tuple with the PoePowerTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePowerTrend

`func (o *CardInfoVO) SetPoePowerTrend(v []PoePowerTrendVO)`

SetPoePowerTrend sets PoePowerTrend field to given value.

### HasPoePowerTrend

`func (o *CardInfoVO) HasPoePowerTrend() bool`

HasPoePowerTrend returns a boolean if a field has been set.

### GetSwitchAlertReboot

`func (o *CardInfoVO) GetSwitchAlertReboot() SwitchAlertRebootVO`

GetSwitchAlertReboot returns the SwitchAlertReboot field if non-nil, zero value otherwise.

### GetSwitchAlertRebootOk

`func (o *CardInfoVO) GetSwitchAlertRebootOk() (*SwitchAlertRebootVO, bool)`

GetSwitchAlertRebootOk returns a tuple with the SwitchAlertReboot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchAlertReboot

`func (o *CardInfoVO) SetSwitchAlertReboot(v SwitchAlertRebootVO)`

SetSwitchAlertReboot sets SwitchAlertReboot field to given value.

### HasSwitchAlertReboot

`func (o *CardInfoVO) HasSwitchAlertReboot() bool`

HasSwitchAlertReboot returns a boolean if a field has been set.

### GetSwitchHealthTrend

`func (o *CardInfoVO) GetSwitchHealthTrend() []DeviceHealthVO`

GetSwitchHealthTrend returns the SwitchHealthTrend field if non-nil, zero value otherwise.

### GetSwitchHealthTrendOk

`func (o *CardInfoVO) GetSwitchHealthTrendOk() (*[]DeviceHealthVO, bool)`

GetSwitchHealthTrendOk returns a tuple with the SwitchHealthTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchHealthTrend

`func (o *CardInfoVO) SetSwitchHealthTrend(v []DeviceHealthVO)`

SetSwitchHealthTrend sets SwitchHealthTrend field to given value.

### HasSwitchHealthTrend

`func (o *CardInfoVO) HasSwitchHealthTrend() bool`

HasSwitchHealthTrend returns a boolean if a field has been set.

### GetSwitchStatus

`func (o *CardInfoVO) GetSwitchStatus() []OnlineOfflineVO`

GetSwitchStatus returns the SwitchStatus field if non-nil, zero value otherwise.

### GetSwitchStatusOk

`func (o *CardInfoVO) GetSwitchStatusOk() (*[]OnlineOfflineVO, bool)`

GetSwitchStatusOk returns a tuple with the SwitchStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchStatus

`func (o *CardInfoVO) SetSwitchStatus(v []OnlineOfflineVO)`

SetSwitchStatus sets SwitchStatus field to given value.

### HasSwitchStatus

`func (o *CardInfoVO) HasSwitchStatus() bool`

HasSwitchStatus returns a boolean if a field has been set.

### GetTopApByCpuAndMemory

`func (o *CardInfoVO) GetTopApByCpuAndMemory() ApUtilizationVO`

GetTopApByCpuAndMemory returns the TopApByCpuAndMemory field if non-nil, zero value otherwise.

### GetTopApByCpuAndMemoryOk

`func (o *CardInfoVO) GetTopApByCpuAndMemoryOk() (*ApUtilizationVO, bool)`

GetTopApByCpuAndMemoryOk returns a tuple with the TopApByCpuAndMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopApByCpuAndMemory

`func (o *CardInfoVO) SetTopApByCpuAndMemory(v ApUtilizationVO)`

SetTopApByCpuAndMemory sets TopApByCpuAndMemory field to given value.

### HasTopApByCpuAndMemory

`func (o *CardInfoVO) HasTopApByCpuAndMemory() bool`

HasTopApByCpuAndMemory returns a boolean if a field has been set.

### GetTopApByInterference

`func (o *CardInfoVO) GetTopApByInterference() TopApByInterferenceVO`

GetTopApByInterference returns the TopApByInterference field if non-nil, zero value otherwise.

### GetTopApByInterferenceOk

`func (o *CardInfoVO) GetTopApByInterferenceOk() (*TopApByInterferenceVO, bool)`

GetTopApByInterferenceOk returns a tuple with the TopApByInterference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopApByInterference

`func (o *CardInfoVO) SetTopApByInterference(v TopApByInterferenceVO)`

SetTopApByInterference sets TopApByInterference field to given value.

### HasTopApByInterference

`func (o *CardInfoVO) HasTopApByInterference() bool`

HasTopApByInterference returns a boolean if a field has been set.

### GetTopApByRtAndDrop

`func (o *CardInfoVO) GetTopApByRtAndDrop() TopApByRtDropVO`

GetTopApByRtAndDrop returns the TopApByRtAndDrop field if non-nil, zero value otherwise.

### GetTopApByRtAndDropOk

`func (o *CardInfoVO) GetTopApByRtAndDropOk() (*TopApByRtDropVO, bool)`

GetTopApByRtAndDropOk returns a tuple with the TopApByRtAndDrop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopApByRtAndDrop

`func (o *CardInfoVO) SetTopApByRtAndDrop(v TopApByRtDropVO)`

SetTopApByRtAndDrop sets TopApByRtAndDrop field to given value.

### HasTopApByRtAndDrop

`func (o *CardInfoVO) HasTopApByRtAndDrop() bool`

HasTopApByRtAndDrop returns a boolean if a field has been set.

### GetTopApByTrafficAndClient

`func (o *CardInfoVO) GetTopApByTrafficAndClient() TopApByTrafficAndClientVO`

GetTopApByTrafficAndClient returns the TopApByTrafficAndClient field if non-nil, zero value otherwise.

### GetTopApByTrafficAndClientOk

`func (o *CardInfoVO) GetTopApByTrafficAndClientOk() (*TopApByTrafficAndClientVO, bool)`

GetTopApByTrafficAndClientOk returns a tuple with the TopApByTrafficAndClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopApByTrafficAndClient

`func (o *CardInfoVO) SetTopApByTrafficAndClient(v TopApByTrafficAndClientVO)`

SetTopApByTrafficAndClient sets TopApByTrafficAndClient field to given value.

### HasTopApByTrafficAndClient

`func (o *CardInfoVO) HasTopApByTrafficAndClient() bool`

HasTopApByTrafficAndClient returns a boolean if a field has been set.

### GetTopApplicationByTraffic

`func (o *CardInfoVO) GetTopApplicationByTraffic() []TopApplicationByTrafficVO`

GetTopApplicationByTraffic returns the TopApplicationByTraffic field if non-nil, zero value otherwise.

### GetTopApplicationByTrafficOk

`func (o *CardInfoVO) GetTopApplicationByTrafficOk() (*[]TopApplicationByTrafficVO, bool)`

GetTopApplicationByTrafficOk returns a tuple with the TopApplicationByTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopApplicationByTraffic

`func (o *CardInfoVO) SetTopApplicationByTraffic(v []TopApplicationByTrafficVO)`

SetTopApplicationByTraffic sets TopApplicationByTraffic field to given value.

### HasTopApplicationByTraffic

`func (o *CardInfoVO) HasTopApplicationByTraffic() bool`

HasTopApplicationByTraffic returns a boolean if a field has been set.

### GetTopClient

`func (o *CardInfoVO) GetTopClient() TopTrafficAndUptimeClients`

GetTopClient returns the TopClient field if non-nil, zero value otherwise.

### GetTopClientOk

`func (o *CardInfoVO) GetTopClientOk() (*TopTrafficAndUptimeClients, bool)`

GetTopClientOk returns a tuple with the TopClient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopClient

`func (o *CardInfoVO) SetTopClient(v TopTrafficAndUptimeClients)`

SetTopClient sets TopClient field to given value.

### HasTopClient

`func (o *CardInfoVO) HasTopClient() bool`

HasTopClient returns a boolean if a field has been set.

### GetTopSsidByTraffic

`func (o *CardInfoVO) GetTopSsidByTraffic() []TopSsidTrafficVO`

GetTopSsidByTraffic returns the TopSsidByTraffic field if non-nil, zero value otherwise.

### GetTopSsidByTrafficOk

`func (o *CardInfoVO) GetTopSsidByTrafficOk() (*[]TopSsidTrafficVO, bool)`

GetTopSsidByTrafficOk returns a tuple with the TopSsidByTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSsidByTraffic

`func (o *CardInfoVO) SetTopSsidByTraffic(v []TopSsidTrafficVO)`

SetTopSsidByTraffic sets TopSsidByTraffic field to given value.

### HasTopSsidByTraffic

`func (o *CardInfoVO) HasTopSsidByTraffic() bool`

HasTopSsidByTraffic returns a boolean if a field has been set.

### GetTopSwitchByTrafficAndPoePower

`func (o *CardInfoVO) GetTopSwitchByTrafficAndPoePower() TopSwitchVO`

GetTopSwitchByTrafficAndPoePower returns the TopSwitchByTrafficAndPoePower field if non-nil, zero value otherwise.

### GetTopSwitchByTrafficAndPoePowerOk

`func (o *CardInfoVO) GetTopSwitchByTrafficAndPoePowerOk() (*TopSwitchVO, bool)`

GetTopSwitchByTrafficAndPoePowerOk returns a tuple with the TopSwitchByTrafficAndPoePower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSwitchByTrafficAndPoePower

`func (o *CardInfoVO) SetTopSwitchByTrafficAndPoePower(v TopSwitchVO)`

SetTopSwitchByTrafficAndPoePower sets TopSwitchByTrafficAndPoePower field to given value.

### HasTopSwitchByTrafficAndPoePower

`func (o *CardInfoVO) HasTopSwitchByTrafficAndPoePower() bool`

HasTopSwitchByTrafficAndPoePower returns a boolean if a field has been set.

### GetTopSwitchCpuMemory

`func (o *CardInfoVO) GetTopSwitchCpuMemory() SwitchUtilizationVO`

GetTopSwitchCpuMemory returns the TopSwitchCpuMemory field if non-nil, zero value otherwise.

### GetTopSwitchCpuMemoryOk

`func (o *CardInfoVO) GetTopSwitchCpuMemoryOk() (*SwitchUtilizationVO, bool)`

GetTopSwitchCpuMemoryOk returns a tuple with the TopSwitchCpuMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSwitchCpuMemory

`func (o *CardInfoVO) SetTopSwitchCpuMemory(v SwitchUtilizationVO)`

SetTopSwitchCpuMemory sets TopSwitchCpuMemory field to given value.

### HasTopSwitchCpuMemory

`func (o *CardInfoVO) HasTopSwitchCpuMemory() bool`

HasTopSwitchCpuMemory returns a boolean if a field has been set.

### GetTrafficDistribution

`func (o *CardInfoVO) GetTrafficDistribution() TrafficDistributionVO`

GetTrafficDistribution returns the TrafficDistribution field if non-nil, zero value otherwise.

### GetTrafficDistributionOk

`func (o *CardInfoVO) GetTrafficDistributionOk() (*TrafficDistributionVO, bool)`

GetTrafficDistributionOk returns a tuple with the TrafficDistribution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficDistribution

`func (o *CardInfoVO) SetTrafficDistribution(v TrafficDistributionVO)`

SetTrafficDistribution sets TrafficDistribution field to given value.

### HasTrafficDistribution

`func (o *CardInfoVO) HasTrafficDistribution() bool`

HasTrafficDistribution returns a boolean if a field has been set.

### GetTrafficSummary

`func (o *CardInfoVO) GetTrafficSummary() TrafficSummaryVO`

GetTrafficSummary returns the TrafficSummary field if non-nil, zero value otherwise.

### GetTrafficSummaryOk

`func (o *CardInfoVO) GetTrafficSummaryOk() (*TrafficSummaryVO, bool)`

GetTrafficSummaryOk returns a tuple with the TrafficSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficSummary

`func (o *CardInfoVO) SetTrafficSummary(v TrafficSummaryVO)`

SetTrafficSummary sets TrafficSummary field to given value.

### HasTrafficSummary

`func (o *CardInfoVO) HasTrafficSummary() bool`

HasTrafficSummary returns a boolean if a field has been set.

### GetWanHealthTrend

`func (o *CardInfoVO) GetWanHealthTrend() []WanHealthTrendVO`

GetWanHealthTrend returns the WanHealthTrend field if non-nil, zero value otherwise.

### GetWanHealthTrendOk

`func (o *CardInfoVO) GetWanHealthTrendOk() (*[]WanHealthTrendVO, bool)`

GetWanHealthTrendOk returns a tuple with the WanHealthTrend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanHealthTrend

`func (o *CardInfoVO) SetWanHealthTrend(v []WanHealthTrendVO)`

SetWanHealthTrend sets WanHealthTrend field to given value.

### HasWanHealthTrend

`func (o *CardInfoVO) HasWanHealthTrend() bool`

HasWanHealthTrend returns a boolean if a field has been set.

### GetWifiHealth

`func (o *CardInfoVO) GetWifiHealth() WifiHealthVO`

GetWifiHealth returns the WifiHealth field if non-nil, zero value otherwise.

### GetWifiHealthOk

`func (o *CardInfoVO) GetWifiHealthOk() (*WifiHealthVO, bool)`

GetWifiHealthOk returns a tuple with the WifiHealth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWifiHealth

`func (o *CardInfoVO) SetWifiHealth(v WifiHealthVO)`

SetWifiHealth sets WifiHealth field to given value.

### HasWifiHealth

`func (o *CardInfoVO) HasWifiHealth() bool`

HasWifiHealth returns a boolean if a field has been set.

### GetWirelessTraffic

`func (o *CardInfoVO) GetWirelessTraffic() WirelessTrafficVO`

GetWirelessTraffic returns the WirelessTraffic field if non-nil, zero value otherwise.

### GetWirelessTrafficOk

`func (o *CardInfoVO) GetWirelessTrafficOk() (*WirelessTrafficVO, bool)`

GetWirelessTrafficOk returns a tuple with the WirelessTraffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTraffic

`func (o *CardInfoVO) SetWirelessTraffic(v WirelessTrafficVO)`

SetWirelessTraffic sets WirelessTraffic field to given value.

### HasWirelessTraffic

`func (o *CardInfoVO) HasWirelessTraffic() bool`

HasWirelessTraffic returns a boolean if a field has been set.

### GetWirelessTrafficSingle

`func (o *CardInfoVO) GetWirelessTrafficSingle() WirelessTrafficSingleVO`

GetWirelessTrafficSingle returns the WirelessTrafficSingle field if non-nil, zero value otherwise.

### GetWirelessTrafficSingleOk

`func (o *CardInfoVO) GetWirelessTrafficSingleOk() (*WirelessTrafficSingleVO, bool)`

GetWirelessTrafficSingleOk returns a tuple with the WirelessTrafficSingle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWirelessTrafficSingle

`func (o *CardInfoVO) SetWirelessTrafficSingle(v WirelessTrafficSingleVO)`

SetWirelessTrafficSingle sets WirelessTrafficSingle field to given value.

### HasWirelessTrafficSingle

`func (o *CardInfoVO) HasWirelessTrafficSingle() bool`

HasWirelessTrafficSingle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


