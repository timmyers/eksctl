package cloudformation

// AWSEMRCluster_VolumeSpecification AWS CloudFormation Resource (AWS::EMR::Cluster.VolumeSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-volumespecification.html
type AWSEMRCluster_VolumeSpecification struct {

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-volumespecification.html#cfn-elasticmapreduce-cluster-volumespecification-iops
	Iops int `json:"Iops,omitempty"`

	// SizeInGB AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-volumespecification.html#cfn-elasticmapreduce-cluster-volumespecification-sizeingb
	SizeInGB int `json:"SizeInGB,omitempty"`

	// VolumeType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-volumespecification.html#cfn-elasticmapreduce-cluster-volumespecification-volumetype
	VolumeType *StringIntrinsic `json:"VolumeType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_VolumeSpecification) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.VolumeSpecification"
}
