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