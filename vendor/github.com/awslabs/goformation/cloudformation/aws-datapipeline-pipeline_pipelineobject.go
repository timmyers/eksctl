package cloudformation

// AWSDataPipelinePipeline_PipelineObject AWS CloudFormation Resource (AWS::DataPipeline::Pipeline.PipelineObject)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html
type AWSDataPipelinePipeline_PipelineObject struct {

	// Fields AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-fields
	Fields []AWSDataPipelinePipeline_Field `json:"Fields,omitempty"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-id
	Id *StringIntrinsic `json:"Id,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-pipelineobjects.html#cfn-datapipeline-pipeline-pipelineobjects-name
	Name *StringIntrinsic `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipeline_PipelineObject) AWSCloudFormationType() string {
	return "AWS::DataPipeline::Pipeline.PipelineObject"
}
