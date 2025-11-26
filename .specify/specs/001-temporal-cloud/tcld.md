# tcld CLI Reference

## Overview

The `tcld` (Temporal Cloud) CLI is the primary tool for managing Temporal Cloud resources. It interacts with the Cloud Ops API.

## Installation

```bash
# Homebrew
brew install temporalio/brew/tcld

# Curl
curl -sSf https://temporal.download/tcld/install.sh | sh
```

## Authentication

```bash
# Interactive login
tcld login

# API Key (CI/CD)
export TEMPORAL_CLOUD_API_KEY=temporal_ak_...
```

## Command Structure

`tcld <resource> <action> [flags]`

## Common Commands

### Account

```bash
tcld account get
tcld account users list
tcld account region list
```

### Namespace

```bash
# Create
tcld namespace create \
  --name my-ns \
  --region aws-us-east-1 \
  --retention-days 30 \
  --ca-certificate cert.pem

# Update
tcld namespace update \
  --namespace my-ns \
  --retention-days 7

# List
tcld namespace list

# Certificate Management
tcld namespace certificates add ...
tcld namespace certificates remove ...
```

### User

```bash
# Invite
tcld user invite \
  --email bob@example.com \
  --account-role developer

# Set Permissions
tcld user namespace-access set \
  --email bob@example.com \
  --namespace my-ns \
  --permission write
```

### Request

```bash
# Async operation status
tcld request get --request-id <id>
```

## Output Formats

All list commands support JSON output for scripting.

```bash
tcld namespace list --output json | jq '.[].name'
```

## Configuration

Config file stored at `~/.config/tcld/config.yml`.

```yaml
server: saas-api.tmprl.cloud:443
disable_version_check: false
default_namespace: my-ns
```
