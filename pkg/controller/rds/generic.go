package rds

import (
	. "github.com/crossplane-contrib/provider-aws/pkg/generics"

	"github.com/aws/aws-sdk-go/aws"
	svcsdk "github.com/aws/aws-sdk-go/service/rds"

	databasev1beta1 "github.com/crossplane-contrib/provider-aws/apis/database/v1beta1"
	rdsv1alpha1 "github.com/crossplane-contrib/provider-aws/apis/rds/v1alpha1"
)

// type constraints for inputs
type CustomResourceInputs interface {
	*rdsv1alpha1.DBCluster |
		*rdsv1alpha1.DBInstance
}

type CreateClusterInputs interface {
	*svcsdk.CreateDBClusterInput
}

type CreateInstanceInputs interface {
	*svcsdk.CreateDBInstanceInput
}

type CreateInputs interface {
	CreateClusterInputs |
		CreateInstanceInputs
}

type RestoreClusterInputs interface {
	*svcsdk.RestoreDBClusterFromSnapshotInput |
		*svcsdk.RestoreDBClusterFromS3Input |
		*svcsdk.RestoreDBClusterToPointInTimeInput
}

type RestoreInstanceInputs interface {
	*svcsdk.RestoreDBInstanceFromDBSnapshotInput |
		*svcsdk.RestoreDBInstanceFromS3Input |
		*svcsdk.RestoreDBInstanceToPointInTimeInput
}

type RestoreInputs interface {
	RestoreClusterInputs |
		RestoreInstanceInputs
}

type ModifyClusterInputs interface {
	*svcsdk.ModifyDBClusterInput
}

type ModifyInstanceInputs interface {
	*svcsdk.ModifyDBInstanceInput
}

type ModifyInputs interface {
	ModifyClusterInputs |
		ModifyInstanceInputs
}

type ClusterInputs interface {
	CreateClusterInputs |
		RestoreClusterInputs |
		ModifyClusterInputs
}

type InstanceInputs interface {
	CreateInstanceInputs |
		RestoreInstanceInputs |
		ModifyInstanceInputs
}

type Inputs interface {
	ClusterInputs |
		InstanceInputs
}

// Interface definitions for input setters

type WithSetAllocatedStorage[T any] interface {
	SetAllocatedStorage(v int64) T
}

func SetAllocatedStorage[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetAllocatedStorage[T]); ok {
		SetIfNonNil(with.SetAllocatedStorage, v)
	}
}

type WithSetAutoMinorVersionUpgrade[T any] interface {
	SetAutoMinorVersionUpgrade(v bool) T
}

func SetAutoMinorVersionUpgrade[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetAutoMinorVersionUpgrade[T]); ok {
		SetIfNonNil(with.SetAutoMinorVersionUpgrade, v)
	}
}

type WithSetAvailabilityZone[T any] interface {
	SetAvailabilityZone(v string) T
}

func SetAvailabilityZone[T any](i T, v *string) {
	if with, ok := any(i).(WithSetAvailabilityZone[T]); ok {
		SetIfNonNil(with.SetAvailabilityZone, v)
	}
}

type WithSetAvailabilityZones[T any] interface {
	SetAvailabilityZones(v []*string) T
}

func SetAvailabilityZones[T any](i T, v []*string) {
	if with, ok := any(i).(WithSetAvailabilityZones[T]); ok {
		SetIfNonNil(with.SetAvailabilityZones, &v)
	}
}

type WithSetBacktrackWindow[T any] interface {
	SetBacktrackWindow(v int64) T
}

func SetBacktrackWindow[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetBacktrackWindow[T]); ok {
		SetIfNonNil(with.SetBacktrackWindow, v)
	}
}

type WithSetBackupRetentionPeriod[T any] interface {
	SetBackupRetentionPeriod(v int64) T
}

func SetBackupRetentionPeriod[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetBackupRetentionPeriod[T]); ok {
		SetIfNonNil(with.SetBackupRetentionPeriod, v)
	}
}

type WithSetCharacterSetName[T any] interface {
	SetCharacterSetName(v string) T
}

func SetCharacterSetName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetCharacterSetName[T]); ok {
		SetIfNonNil(with.SetCharacterSetName, v)
	}
}

type WithSetCopyTagsToSnapshot[T any] interface {
	SetCopyTagsToSnapshot(v bool) T
}

func SetCopyTagsToSnapshot[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetCopyTagsToSnapshot[T]); ok {
		SetIfNonNil(with.SetCopyTagsToSnapshot, v)
	}
}

type WithSetDBClusterIdentifier[T any] interface {
	SetDBClusterIdentifier(v string) T
}

func SetDBClusterIdentifier[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBClusterIdentifier[T]); ok {
		SetIfNonNil(with.SetDBClusterIdentifier, v)
	}
}

type WithSetDBClusterParameterGroupName[T any] interface {
	SetDBClusterParameterGroupName(v string) T
}

func SetDBClusterParameterGroupName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBClusterParameterGroupName[T]); ok {
		SetIfNonNil(with.SetDBClusterParameterGroupName, v)
	}
}

type WithSetDBInstanceIdentifier[T any] interface {
	SetDBInstanceIdentifier(v string) T
}

func SetDBInstanceIdentifier[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBInstanceIdentifier[T]); ok {
		SetIfNonNil(with.SetDBInstanceIdentifier, v)
	}
}

type WithSetDBInstanceClass[T any] interface {
	SetDBInstanceClass(v string) T
}

func SetDBInstanceClass[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBInstanceClass[T]); ok {
		SetIfNonNil(with.SetDBInstanceClass, v)
	}
}

type WithSetDBName[T any] interface {
	SetDBName(v string) T
}

func SetDBName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBName[T]); ok {
		SetIfNonNil(with.SetDBName, v)
	}
}

type WithSetDBParameterGroupName[T any] interface {
	SetDBParameterGroupName(v string) T
}

func SetDBParameterGroupName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBParameterGroupName[T]); ok {
		SetIfNonNil(with.SetDBParameterGroupName, v)
	}
}

type WithSetDBSecurityGroups[T any] interface {
	SetDBSecurityGroups(v []*string) T
}

func SetDBSecurityGroups[T any](i T, v []*string) {
	if with, ok := any(i).(WithSetDBSecurityGroups[T]); ok {
		SetIfNonNil(with.SetDBSecurityGroups, &v)
	}
}

type WithSetDBSubnetGroupName[T any] interface {
	SetDBSubnetGroupName(v string) T
}

func SetDBSubnetGroupName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDBSubnetGroupName[T]); ok {
		SetIfNonNil(with.SetDBSubnetGroupName, v)
	}
}

type WithSetDeletionProtection[T any] interface {
	SetDeletionProtection(v bool) T
}

func SetDeletionProtection[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetDeletionProtection[T]); ok {
		SetIfNonNil(with.SetDeletionProtection, v)
	}
}

type WithSetDomain[T any] interface {
	SetDomain(v string) T
}

func SetDomain[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDomain[T]); ok {
		SetIfNonNil(with.SetDomain, v)
	}
}

type WithSetDomainIAMRoleName[T any] interface {
	SetDomainIAMRoleName(v string) T
}

func SetDomainIAMRoleName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetDomainIAMRoleName[T]); ok {
		SetIfNonNil(with.SetDomainIAMRoleName, v)
	}
}

type WithSetEnableCloudwatchLogsExports[T any] interface {
	SetEnableCloudwatchLogsExports(v []*string) T
}

func SetEnableCloudwatchLogsExports[T any](i T, v []*string) {
	if with, ok := any(i).(WithSetEnableCloudwatchLogsExports[T]); ok {
		SetIfNonNil(with.SetEnableCloudwatchLogsExports, &v)
	}
}

type WithSetEnableIAMDatabaseAuthentication[T any] interface {
	SetEnableIAMDatabaseAuthentication(v bool) T
}

func SetEnableIAMDatabaseAuthentication[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetEnableIAMDatabaseAuthentication[T]); ok {
		SetIfNonNil(with.SetEnableIAMDatabaseAuthentication, v)
	}
}

type WithSetEnablePerformanceInsights[T any] interface {
	SetEnablePerformanceInsights(v bool) T
}

func SetEnablePerformanceInsights[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetEnablePerformanceInsights[T]); ok {
		SetIfNonNil(with.SetEnablePerformanceInsights, v)
	}
}

type WithSetEngine[T any] interface {
	SetEngine(v string) T
}

func SetEngine[T any](i T, v *string) {
	if with, ok := any(i).(WithSetEngine[T]); ok {
		SetIfNonNil(with.SetEngine, v)
	}
}

type WithSetEngineMode[T any] interface {
	SetEngineMode(v string) T
}

func SetEngineMode[T any](i T, v *string) {
	if with, ok := any(i).(WithSetEngineMode[T]); ok {
		SetIfNonNil(with.SetEngineMode, v)
	}
}

type WithSetEngineVersion[T any] interface {
	SetEngineVersion(v string) T
}

func SetEngineVersion[T any](i T, v *string) {
	if with, ok := any(i).(WithSetEngineVersion[T]); ok {
		SetIfNonNil(with.SetEngineVersion, v)
	}
}

type WithSetIops[T any] interface {
	SetIops(v int64) T
}

func SetIops[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetIops[T]); ok {
		SetIfNonNil(with.SetIops, v)
	}
}

type WithSetKmsKeyId[T any] interface {
	SetKmsKeyId(v string) T
}

func SetKmsKeyId[T any](i T, v *string) {
	if with, ok := any(i).(WithSetKmsKeyId[T]); ok {
		SetIfNonNil(with.SetKmsKeyId, v)
	}
}

type WithSetLicenseModel[T any] interface {
	SetLicenseModel(v string) T
}

func SetLicenseModel[T any](i T, v *string) {
	if with, ok := any(i).(WithSetLicenseModel[T]); ok {
		SetIfNonNil(with.SetLicenseModel, v)
	}
}

type WithSetMasterUserPassword[T any] interface {
	SetMasterUserPassword(v string) T
}

func SetMasterUserPassword[T any](i T, v *string) {
	if with, ok := any(i).(WithSetMasterUserPassword[T]); ok {
		SetIfNonNil(with.SetMasterUserPassword, v)
	}
}

type WithSetMasterUsername[T any] interface {
	SetMasterUsername(v string) T
}

func SetMasterUsername[T any](i T, v *string) {
	if with, ok := any(i).(WithSetMasterUsername[T]); ok {
		SetIfNonNil(with.SetMasterUsername, v)
	}
}

type WithSetMaxAllocatedStorage[T any] interface {
	SetMaxAllocatedStorage(v int64) T
}

func SetMaxAllocatedStorage[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetMaxAllocatedStorage[T]); ok {
		SetIfNonNil(with.SetMaxAllocatedStorage, v)
	}
}

type WithSetMonitoringInterval[T any] interface {
	SetMonitoringInterval(v int64) T
}

func SetMonitoringInterval[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetMonitoringInterval[T]); ok {
		SetIfNonNil(with.SetMonitoringInterval, v)
	}
}

type WithSetMonitoringRoleArn[T any] interface {
	SetMonitoringRoleArn(v int64) T
}

func SetMonitoringRoleArn[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetMonitoringRoleArn[T]); ok {
		SetIfNonNil(with.SetMonitoringRoleArn, v)
	}
}

type WithSetMultiAZ[T any] interface {
	SetMultiAZ(v bool) T
}

func SetMultiAZ[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetMultiAZ[T]); ok {
		SetIfNonNil(with.SetMultiAZ, v)
	}
}

type WithSetOptionGroupName[T any] interface {
	SetOptionGroupName(v string) T
}

func SetOptionGroupName[T any](i T, v *string) {
	if with, ok := any(i).(WithSetOptionGroupName[T]); ok {
		SetIfNonNil(with.SetOptionGroupName, v)
	}
}

type WithSetPerformanceInsightsKMSKeyId[T any] interface {
	SetPerformanceInsightsKMSKeyId(v string) T
}

func SetPerformanceInsightsKMSKeyId[T any](i T, v *string) {
	if with, ok := any(i).(WithSetPerformanceInsightsKMSKeyId[T]); ok {
		SetIfNonNil(with.SetPerformanceInsightsKMSKeyId, v)
	}
}

type WithSetPerformanceInsightsRetentionPeriod[T any] interface {
	SetPerformanceInsightsRetentionPeriod(v int64) T
}

func SetPerformanceInsightsRetentionPeriod[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetPerformanceInsightsRetentionPeriod[T]); ok {
		SetIfNonNil(with.SetPerformanceInsightsRetentionPeriod, v)
	}
}

type WithSetPort[T any] interface {
	SetPort(v int64) T
}

func SetPort[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetPort[T]); ok {
		SetIfNonNil(with.SetPort, v)
	}
}

type WithSetPreferredBackupWindow[T any] interface {
	SetPreferredBackupWindow(v string) T
}

func SetPreferredBackupWindow[T any](i T, v *string) {
	if with, ok := any(i).(WithSetPreferredBackupWindow[T]); ok {
		SetIfNonNil(with.SetPreferredBackupWindow, v)
	}
}

type WithSetPreferredMaintenanceWindow[T any] interface {
	SetPreferredMaintenanceWindow(v string) T
}

func SetPreferredMaintenanceWindow[T any](i T, v *string) {
	if with, ok := any(i).(WithSetPreferredMaintenanceWindow[T]); ok {
		SetIfNonNil(with.SetPreferredMaintenanceWindow, v)
	}
}

type WithSetProcessorFeatures[T any, PF ProcessorFeature] interface {
	SetProcessorFeatures(v []*svcsdk.ProcessorFeature) T
}

func SetProcessorFeatures[T any, PF ProcessorFeature](i T, v []PF) {
	if with, ok := any(i).(WithSetProcessorFeatures[T, PF]); ok {
		if v == nil {
			return
		}
		pfs := ConvertSlice(ToAWSProcessorFeature[PF], v)
		with.SetProcessorFeatures(pfs)
	}
}

type WithSetPromotionTier[T any] interface {
	SetPromotionTier(v int64) T
}

func SetPromotionTier[T any](i T, v *int64) {
	if with, ok := any(i).(WithSetPromotionTier[T]); ok {
		SetIfNonNil(with.SetPromotionTier, v)
	}
}

type WithSetPubliclyAccessible[T any] interface {
	SetPubliclyAccessible(v bool) T
}

func SetPubliclyAccessible[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetPubliclyAccessible[T]); ok {
		SetIfNonNil(with.SetPubliclyAccessible, v)
	}
}

type WithSetScalingConfiguration[T any] interface {
	SetScalingConfiguration(v *svcsdk.ScalingConfiguration) T
}

type ScalingConfiguration interface {
	*databasev1beta1.ScalingConfiguration |
		*rdsv1alpha1.ScalingConfiguration
}

func SetScalingConfiguration[T any, CONFIG ScalingConfiguration](i T, v CONFIG) {
	if v == nil {
		return
	}
	if with, ok := any(i).(WithSetScalingConfiguration[T]); ok {
		var sc svcsdk.ScalingConfiguration
		switch any(v).(type) {
		case *rdsv1alpha1.ScalingConfiguration:
			v := any(i).(*rdsv1alpha1.ScalingConfiguration)
			sc.AutoPause = v.AutoPause
			sc.MaxCapacity = v.MaxCapacity
			sc.MinCapacity = v.MinCapacity
			sc.SecondsBeforeTimeout = v.SecondsBeforeTimeout
			sc.SecondsUntilAutoPause = v.SecondsUntilAutoPause
			sc.TimeoutAction = v.TimeoutAction
		case *databasev1beta1.ScalingConfiguration:
			v := any(i).(*databasev1beta1.ScalingConfiguration)
			sc.AutoPause = v.AutoPause
			max := int64(aws.IntValue(v.MaxCapacity))
			sc.MaxCapacity = &max
			min := int64(aws.IntValue(v.MinCapacity))
			sc.MinCapacity = &min
			pause := int64(aws.IntValue(v.SecondsUntilAutoPause))
			sc.SecondsUntilAutoPause = &pause
		}
		with.SetScalingConfiguration(&sc)
	}
}

type WithSetTags[T any] interface {
	SetTags(v []*svcsdk.Tag) T
}

func SetTags[T any, TAG Tag](i T, v []TAG) {
	if with, ok := any(i).(WithSetTags[T]); ok {
		if v == nil {
			return
		}
		tags := ConvertSlice(ToAWSSDKTag[TAG], v)
		with.SetTags(tags)
	}
}

type WithSetVpcSecurityGroupIds[T any] interface {
	SetVpcSecurityGroupIds(v []*string) T
}

func SetVpcSecurityGroupIds[T any](i T, v []*string) {
	if with, ok := any(i).(WithSetVpcSecurityGroupIds[T]); ok {
		if v == nil {
			return
		}
		with.SetVpcSecurityGroupIds(v)
	}
}

type WithSetStorageEncrypted[T any] interface {
	SetStorageEncrypted(v bool) T
}

func SetStorageEncrypted[T any](i T, v *bool) {
	if with, ok := any(i).(WithSetStorageEncrypted[T]); ok {
		SetIfNonNil(with.SetStorageEncrypted, v)
	}
}

type WithSetStorageType[T any] interface {
	SetStorageType(v string) T
}

func SetStorageType[T any](i T, v *string) {
	if with, ok := any(i).(WithSetStorageType[T]); ok {
		SetIfNonNil(with.SetStorageType, v)
	}
}

type WithSetTimezone[T any] interface {
	SetTimezone(v string) T
}

func SetTimezone[T any](i T, v *string) {
	if with, ok := any(i).(WithSetTimezone[T]); ok {
		SetIfNonNil(with.SetTimezone, v)
	}
}

// CommonParameters provides an aggregate of Interfaces which are supported across the contrained types
type CommonParameters[T any] interface {
	WithSetCopyTagsToSnapshot[T]
	WithSetDBSubnetGroupName[T]
	WithSetDeletionProtection[T]
	WithSetEnableCloudwatchLogsExports[T]
	WithSetEnableIAMDatabaseAuthentication[T]
	WithSetOptionGroupName[T]
	WithSetPort[T]
	WithSetTags[T]
	WithSetVpcSecurityGroupIds[T]
}

func SatisfiesCommonParameters[T CommonParameters[T]]() {}

// test that each concrete type satifies the type constrained CommonParameters interface
var _ = SatisfiesCommonParameters[*svcsdk.CreateDBClusterInput]
var _ = SatisfiesCommonParameters[*svcsdk.RestoreDBClusterFromS3Input]
var _ = SatisfiesCommonParameters[*svcsdk.RestoreDBClusterToPointInTimeInput]
var _ = SatisfiesCommonParameters[*svcsdk.CreateDBInstanceInput]
var _ = SatisfiesCommonParameters[*svcsdk.RestoreDBInstanceFromS3Input]
var _ = SatisfiesCommonParameters[*svcsdk.RestoreDBInstanceFromDBSnapshotInput]
var _ = SatisfiesCommonParameters[*svcsdk.RestoreDBInstanceToPointInTimeInput]

type ClusterParameters[T ClusterInputs] interface {
	ClusterInputs
	CommonParameters[T]
	WithSetBacktrackWindow[T]
	WithSetDBClusterIdentifier[T]
	WithSetDBClusterParameterGroupName[T]
	WithSetDomain[T]
	WithSetDomainIAMRoleName[T]
	WithSetKmsKeyId[T]
	// WithSetPreferredMaintenanceWindow[T]
}

func SatisfiesClusterParameters[T ClusterParameters[T]]() {}

// test that each concrete type satifies the type constrained ClusterParameters interface
var _ = SatisfiesClusterParameters[*svcsdk.CreateDBClusterInput]
var _ = SatisfiesClusterParameters[*svcsdk.RestoreDBClusterFromS3Input]
var _ = SatisfiesClusterParameters[*svcsdk.RestoreDBClusterFromSnapshotInput]
var _ = SatisfiesClusterParameters[*svcsdk.RestoreDBClusterToPointInTimeInput]

type InstanceParameters[T InstanceInputs] interface {
	InstanceInputs
	CommonParameters[T]
	// WithSetAllocatedStorage[T]
	WithSetAutoMinorVersionUpgrade[T]
	WithSetAvailabilityZone[T]
	// WithSetCharacterSetName[T]
	// WithSetBackupRetentionPeriod[T]
	// WithSetDBInstanceIdentifier[T]
	WithSetDBInstanceClass[T]
	WithSetDBName[T]
	WithSetDBParameterGroupName[T]
	// WithSetDBSecurityGroups[T]
	// WithSetEnablePerformanceInsights[T]
	WithSetEngine[T]
	// WithSetEngineVersion[T]
	WithSetIops[T]
	// WithSetKmsKeyId[T]
	WithSetLicenseModel[T]
	// WithSetMasterUserPassword[T]
	// WithSetMasterUsername[T]
	// WithSetMaxAllocatedStorage[T]
	// WithSetMonitoringInterval[T]
	WithSetMultiAZ[T]
	// WithSetPerformanceInsightsKMSKeyId[T]
	// WithSetPerformanceInsightsRetentionPeriod[T]
	// WithSetPreferredMaintenanceWindow[T]
	// WithSetPromotionTier[T]
	WithSetPubliclyAccessible[T]
	// WithSetStorageEncrypted[T]
	// WithSetTimezone[T]
}

func SatisfiesInstanceParameters[T InstanceParameters[T]]() {}

// test that each concrete type satifies the type constrained InstanceParameters interface
var _ = SatisfiesInstanceParameters[*svcsdk.CreateDBInstanceInput]
var _ = SatisfiesInstanceParameters[*svcsdk.RestoreDBInstanceFromS3Input]
var _ = SatisfiesInstanceParameters[*svcsdk.RestoreDBInstanceFromDBSnapshotInput]
var _ = SatisfiesInstanceParameters[*svcsdk.RestoreDBInstanceToPointInTimeInput]
