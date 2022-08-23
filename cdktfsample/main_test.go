package main

import (
	"testing"

	"cdk.tf/go/stack/generated/aws/s3"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestShouldContainNewS3Bucket(t *testing.T) {
	stack := NewMyStack(cdktf.Testing_App(nil), "cdktfsample")
	synth := cdktf.Testing_Synth(stack)
	require.NotNil(t, synth)

	properties := &map[string]interface{}{
		"bucket": jsii.String("cdktf-sample-bucket"),
		"tags": &struct {
			Name string
		}{
			Name: "Bucket provisioned by CDKTF",
		},
	}

	bucketExists := cdktf.Testing_ToHaveResourceWithProperties(synth, s3.S3Bucket_TfResourceType(), properties)

	assert.True(t, *bucketExists)
}
