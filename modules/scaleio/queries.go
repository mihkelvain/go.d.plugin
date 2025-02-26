package scaleio

var (
	selectedStatisticsQuery = `
{
  "selectedStatisticsList": [
    {
      "type": "System",
      "properties": [
        "bckRebuildReadBwc",
        "bckRebuildWriteBwc",
        "fwdRebuildReadBwc",
        "fwdRebuildWriteBwc",
        "normRebuildReadBwc",
        "normRebuildWriteBwc",
        "rebalanceReadBwc",
        "rebalanceWriteBwc",
        "primaryReadBwc",
        "primaryWriteBwc",
        "secondaryReadBwc",
        "secondaryWriteBwc",
        "userDataReadBwc",
        "userDataWriteBwc",
        "totalReadBwc",
        "totalWriteBwc",
        "activeBckRebuildCapacityInKb",
        "activeFwdRebuildCapacityInKb",
        "activeMovingCapacityInKb",
        "activeNormRebuildCapacityInKb",
        "activeRebalanceCapacityInKb",
        "atRestCapacityInKb",
        "bckRebuildCapacityInKb",
        "capacityAvailableForVolumeAllocationInKb",
        "capacityInUseInKb",
        "capacityLimitInKb",
        "degradedFailedCapacityInKb",
        "degradedHealthyCapacityInKb",
        "failedCapacityInKb",
        "fwdRebuildCapacityInKb",
        "inMaintenanceCapacityInKb",
        "maxCapacityInKb",
        "movingCapacityInKb",
        "normRebuildCapacityInKb",
        "pendingBckRebuildCapacityInKb",
        "pendingFwdRebuildCapacityInKb",
        "pendingMovingCapacityInKb",
        "pendingNormRebuildCapacityInKb",
        "pendingRebalanceCapacityInKb",
        "protectedCapacityInKb",
        "rebalanceCapacityInKb",
        "semiProtectedCapacityInKb",
        "snapCapacityInUseInKb",
        "snapCapacityInUseOccupiedInKb",
        "spareCapacityInKb",
        "thickCapacityInUseInKb",
        "thinCapacityAllocatedInKb",
        "thinCapacityInUseInKb",
        "unreachableUnusedCapacityInKb",
        "unusedCapacityInKb",
        "numOfDevices",
        "numOfFaultSets",
        "numOfMappedToAllVolumes",
        "numOfProtectionDomains",
        "numOfRfcacheDevices",
	// Starting from version 3 of ScaleIO/VxFlex API numOfScsiInitiators property is removed.
	// Reference: VxFlex OS v3.x REST API Reference Guide.pdf
	// "numOfScsiInitiators",
        "numOfSdc",
        "numOfSds",
        "numOfSnapshots",
        "numOfStoragePools",
        "numOfThickBaseVolumes",
        "numOfThinBaseVolumes",
        "numOfUnmappedVolumes",
        "numOfVolumes",
        "numOfVolumesInDeletion",
        "numOfVtrees"
      ]
    }
  ]
}
`
)
