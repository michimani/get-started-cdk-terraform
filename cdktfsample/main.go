package main

import (
	"cdk.tf/go/stack/generated/aws"
	"cdk.tf/go/stack/generated/aws/s3"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)

	NewMyStack(app, "cdktfsample")

	app.Synth()
}

// NewMyStack は S3 Bucket を作成する Terraform Stack を返す
func NewMyStack(scope constructs.Construct, id string) cdktf.TerraformStack {
	// Terraform Stack の初期化
	stack := cdktf.NewTerraformStack(scope, &id)

	// Provider の設定
	aws.NewAwsProvider(stack, jsii.String("AWS"), &aws.AwsProviderConfig{
		Region: jsii.String("ap-northeast-1"),
	})

	// S3 Bucket を定義
	s3.NewS3Bucket(stack, jsii.String(id+"SampleBucket"), &s3.S3BucketConfig{
		Bucket: jsii.String("cdktf-sample-bucket"),
		Tags: &map[string]*string{
			"Name": jsii.String("Bucket provisioned by CDKTF"),
		},
	})

	return stack
}
