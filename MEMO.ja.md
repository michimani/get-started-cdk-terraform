CDK for Terraform 使ってみる
===

# CDK for Terraform とは

使い慣れたプログラミング言語でインフラの定義やプロビジョニングを行えるツール。HCL を学ぶ必要なく、 Terraform でインフラの管理が出来る。  
AWS CDK のコンセプトとライブラリを利用し、各言語で書いたコードを Terraform 用の設定ファイルとして変換する。

[CDK for Terraform | Terraform by HashiCorp](https://www.terraform.io/cdktf)

# 準備

[Install CDK for Terraform and Run a Quick Start Demo | Terraform - HashiCorp Learn](https://learn.hashicorp.com/tutorials/terraform/cdktf-install?in=terraform/cdktf)

## 必要なツール類

- Terraform CLI (v 1.1 以降)
  - [Install Terraform | Terraform - HashiCorp Learn](https://learn.hashicorp.com/tutorials/terraform/install-cli)
- Node.js, npm (v16 以降)

## CDKTF をインストール

```bash
npm install --global cdktf-cli@latest
```

```bash: バージョン確認
cdktf --version
```

## プロジェクトの作成

ディレクトリ作成。

```bash
mkdir cdktfsample
```

プロジェクトの初期化。今回は Go を使う。

```bash
cdktf init --template=go --local
```

作成されるファイル群。

```bash
.
├── cdktf.json
├── go.mod
├── go.sum
├── help
├── main.go
└── main_test.go
```

# AWS のリソースを管理する

[Build AWS Infrastructure with CDK for Terraform | Terraform - HashiCorp Learn](https://learn.hashicorp.com/tutorials/terraform/cdktf-build)

## AWS のプロバイダを追加

```json: cdktf.json
{
  "language": "go",
  "app": "go run main.go",
  "codeMakerOutput": "generated",
  "projectId": "928f31a4-f624-4a03-9dd1-fceefa5c0845",
  "sendCrashReports": "true",
  "terraformProviders": [
    "aws@~>4.0"
  ],
  "terraformModules": [],
  "context": {
    "excludeStackIdFromLogicalIds": "true",
    "allowSepCharsInLogicalIds": "true"
  }
}
```

```bash
cdktf get
```

他のプロバイダは下記参照。

[HashiCorp: cdktf-provider-](https://github.com/orgs/hashicorp/repositories?q=cdktf-provider-)

## S3 Bucket を定義

```go
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
	})

	return stack
}
```

## Terraform の定義ファイルを生成

```bash
cdktf synth
```

`cdktf.out` ディレクトリ以下に `.tf.json` ファイルが生成される。今回は下記のようなファイルが生成される。

```json
{
  "//": {
    "metadata": {
      "backend": "local",
      "stackName": "cdktfsample",
      "version": "0.12.0"
    },
    "outputs": {
    }
  },
  "provider": {
    "aws": [
      {
        "region": "ap-northeast-1"
      }
    ]
  },
  "resource": {
    "aws_s3_bucket": {
      "cdktfsampleSampleBucket": {
        "//": {
          "metadata": {
            "path": "cdktfsample/cdktfsampleSampleBucket",
            "uniqueId": "cdktfsampleSampleBucket"
          }
        },
        "bucket": "cdktf-sample-bucket"
      }
    }
  },
  "terraform": {
    "backend": {
      "local": {
        "path": "/path/to/get-started-cdk-terraform/cdktfsample/terraform.cdktfsample.tfstate"
      }
    },
    "required_providers": {
      "aws": {
        "source": "aws",
        "version": "4.27.0"
      }
    }
  }
}
```

## デプロイ

```bash
cdktf deploy
```

plan の結果が表示され、続けるかどうかを聞かれる。

```
...
             Plan: 1 to add, 0 to change, 0 to destroy.
             
             ─────────────────────────────────────────────────────────────────────────────
cdktfsample  Saved the plan to: plan

             To perform exactly these actions, run the following command to apply:
             terraform apply "plan"

Please review the diff output above for cdktfsample
❯ Approve  Applies the changes outlined in the plan.
  Dismiss
  Stop
```

`Approve` で apply される。

## 内容を変更

再度 `cdktf synth` を実行して、 `cdktf diff` で差分を確認。

```bash
cdktf diff
```

```
...
cdktfsample  Terraform used the selected providers to generate the following execution
             plan. Resource actions are indicated with the following symbols:
             ~ update in-place
             
             Terraform will perform the following actions:
cdktfsample    # aws_s3_bucket.cdktfsampleSampleBucket (cdktfsampleSampleBucket) will be updated in-place
               ~ resource "aws_s3_bucket" "cdktfsampleSampleBucket" {
             id                          = "cdktf-sample-bucket"
             ~ tags                        = {
             + "Name" = "Bucket provisioned by CDKTF"
             }
             ~ tags_all                    = {
             + "Name" = "Bucket provisioned by CDKTF"
             }
             # (9 unchanged attributes hidden)


             # (2 unchanged blocks hidden)
             }

             Plan: 0 to add, 1 to change, 0 to destroy.
```

デプロイ。

```bash
cdktf deploy
```