package cloudformation

// AWSS3Bucket_SseKmsEncryptedObjects AWS CloudFormation Resource (AWS::S3::Bucket.SseKmsEncryptedObjects)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html
type AWSS3Bucket_SseKmsEncryptedObjects struct {

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html#cfn-s3-bucket-ssekmsencryptedobjects-status
	Status string `json:"Status,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_SseKmsEncryptedObjects) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.SseKmsEncryptedObjects"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSS3Bucket_SseKmsEncryptedObjects) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
