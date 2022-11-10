package rds

import (
	. "github.com/crossplane-contrib/provider-aws/pkg/generics"

	svcsdk "github.com/aws/aws-sdk-go/service/rds"
	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/rds/v1alpha1"
)

func SetFieldsForCluster[T ClusterParameters[T]](cr *svcapitypes.DBCluster, obj *svcsdk.CreateDBClusterInput, input T) error {
	p := cr.Spec.ForProvider

	// SetIfNonNil(input.SetAllocatedStorage, p.BacktrackWindow)
	SetBacktrackWindow(input, p.BacktrackWindow)
	SetCopyTagsToSnapshot(input, p.CopyTagsToSnapshot)
	SetDBClusterIdentifier(input, obj.DBClusterIdentifier)
	SetDBClusterParameterGroupName(input, p.DBClusterParameterGroupName)
	SetDBSubnetGroupName(input, p.DBSubnetGroupName)
	SetDeletionProtection(input, p.DeletionProtection)
	SetDomain(input, p.Domain)
	SetDomainIAMRoleName(input, p.DomainIAMRoleName)
	SetEnableCloudwatchLogsExports(input, p.EnableCloudwatchLogsExports)
	SetEnableIAMDatabaseAuthentication(input, p.EnableIAMDatabaseAuthentication)
	SetKmsKeyId(input, p.KMSKeyID)
	SetOptionGroupName(input, p.OptionGroupName)
	SetPort(input, p.Port)
	SetTags(input, p.Tags)
	SetVpcSecurityGroupIds(input, obj.VpcSecurityGroupIds)
	return nil
}

func SetFieldsForInstance[T InstanceParameters[T]](cr *svcapitypes.DBInstance, obj *svcsdk.CreateDBInstanceInput, input T) error {
	p := cr.Spec.ForProvider
	SetAllocatedStorage(input, p.AllocatedStorage)
	SetAutoMinorVersionUpgrade(input, p.AutoMinorVersionUpgrade)
	SetAvailabilityZone(input, p.AvailabilityZone)
	SetBackupRetentionPeriod(input, p.BackupRetentionPeriod)
	SetCharacterSetName(input, p.CharacterSetName)
	SetCopyTagsToSnapshot(input, p.CopyTagsToSnapshot)
	SetDBClusterIdentifier(input, p.DBClusterIdentifier)
	SetDBInstanceIdentifier(input, obj.DBClusterIdentifier)
	SetDBInstanceClass(input, p.DBInstanceClass)
	SetDBName(input, p.DBName)
	SetDBParameterGroupName(input, p.DBParameterGroupName)
	SetDBSecurityGroups(
		input,
		ConvertSlice(
			func(i string) *string {
				return &i
			},
			p.DBSecurityGroups),
	)
	SetDBSubnetGroupName(input, p.DBSubnetGroupName)
	SetDeletionProtection(input, p.DeletionProtection)
	SetDomain(input, p.Domain)
	SetDomainIAMRoleName(input, p.DomainIAMRoleName)
	SetEnableCloudwatchLogsExports(input, p.EnableCloudwatchLogsExports)
	SetEnableIAMDatabaseAuthentication(input, p.EnableIAMDatabaseAuthentication)
	SetEnablePerformanceInsights(input, p.EnablePerformanceInsights)
	SetEngine(input, p.Engine)
	SetEngineVersion(input, p.EngineVersion)
	SetIops(input, p.IOPS)
	SetKmsKeyId(input, cr.Spec.ForProvider.KMSKeyID)
	SetLicenseModel(input, p.LicenseModel)
	SetMasterUserPassword(input, obj.MasterUserPassword)
	SetMasterUsername(input, p.MasterUsername)
	SetMaxAllocatedStorage(input, p.MaxAllocatedStorage)
	SetMonitoringInterval(input, p.MonitoringInterval)
	SetMonitoringRoleArn(input, p.MonitoringInterval)
	SetMultiAZ(input, p.MultiAZ)
	SetOptionGroupName(input, p.OptionGroupName)
	SetPerformanceInsightsKMSKeyId(input, p.PerformanceInsightsKMSKeyID)
	SetPerformanceInsightsRetentionPeriod(input, p.PerformanceInsightsRetentionPeriod)
	SetPort(input, p.Port)
	SetPreferredBackupWindow(input, p.PreferredBackupWindow)
	SetPreferredMaintenanceWindow(input, p.PreferredMaintenanceWindow)
	SetProcessorFeatures(input, p.ProcessorFeatures)
	SetPromotionTier(input, p.PromotionTier)
	SetPubliclyAccessible(input, p.PubliclyAccessible)
	//SetScalingConfiguration(input, p.ScalingConfiguration)
	SetTags(input, p.Tags)
	SetVpcSecurityGroupIds(input, obj.VpcSecurityGroupIds)
	SetStorageEncrypted(input, p.StorageEncrypted)
	SetStorageType(input, p.StorageType)
	SetTimezone(input, p.Timezone)
	return nil
}
