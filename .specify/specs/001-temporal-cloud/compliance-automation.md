# Compliance Automation

## Overview

Automated evidence collection and compliance checks for SOC 2 Type II, GDPR, and other regulatory requirements.

## Tagging Strategy

All Terraform resources must have compliance tags:

```hcl
locals {
  common_tags = {
    Environment        = var.environment
    Project            = "temporal-cloud"
    Compliance         = "SOC2"
    DataClassification = "Confidential"
    Owner              = "platform-team"
    ManagedBy          = "terraform"
  }
}

resource "aws_db_instance" "main" {
  # ... configuration ...
  tags = merge(local.common_tags, {
    DataClassification = "Restricted"
    BackupRequired     = "true"
  })
}
```

### Data Classification Levels

| Level        | Description        | Examples                   |
| ------------ | ------------------ | -------------------------- |
| Public       | No restrictions    | Marketing content          |
| Internal     | Internal use only  | Architecture docs          |
| Confidential | Business sensitive | Usage metrics              |
| Restricted   | Highly sensitive   | Customer data, credentials |

## SOC 2 Control Mapping

### CC6.1 - Logical Access Controls

| Control        | Implementation | Evidence                 |
| -------------- | -------------- | ------------------------ |
| Authentication | SAML SSO, MFA  | IAM policies, SSO config |
| Authorization  | RBAC           | Role assignments         |
| Access review  | Quarterly      | Access review reports    |

### CC6.6 - Encryption

| Control         | Implementation | Evidence                 |
| --------------- | -------------- | ------------------------ |
| Data at rest    | AES-256        | RDS/S3 encryption config |
| Data in transit | TLS 1.3        | ALB/cert config          |
| Key management  | AWS KMS        | KMS key policies         |

### CC7.2 - Monitoring

| Control             | Implementation        | Evidence        |
| ------------------- | --------------------- | --------------- |
| Security monitoring | CloudTrail, GuardDuty | Alert configs   |
| Anomaly detection   | CloudWatch Anomaly    | Detection rules |
| Incident response   | PagerDuty             | Runbooks        |

## Automated Compliance Checks

### Daily CI Job

```yaml
# .github/workflows/compliance-check.yaml
name: Compliance Check
on:
  schedule:
    - cron: "0 6 * * *" # Daily at 6 AM UTC
  workflow_dispatch:

jobs:
  compliance:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4

      - name: Setup AWS credentials
        uses: aws-actions/configure-aws-credentials@v4
        with:
          role-to-assume: ${{ secrets.COMPLIANCE_ROLE_ARN }}
          aws-region: us-east-1

      - name: Run compliance checks
        run: |
          python scripts/compliance-check.py \
            --output compliance-report.json

      - name: Upload report
        uses: actions/upload-artifact@v4
        with:
          name: compliance-report
          path: compliance-report.json

      - name: Fail on critical issues
        run: |
          CRITICAL=$(jq '.critical_count' compliance-report.json)
          if [ "$CRITICAL" -gt 0 ]; then
            echo "Critical compliance issues found!"
            exit 1
          fi
```

### Compliance Check Script

```python
# scripts/compliance-check.py
import boto3
import json
from datetime import datetime, timedelta

class ComplianceChecker:
    def __init__(self):
        self.results = {
            "timestamp": datetime.utcnow().isoformat(),
            "checks": [],
            "critical_count": 0,
            "warning_count": 0
        }

    def check_rds_encryption(self):
        """CC6.6 - All RDS instances must be encrypted"""
        rds = boto3.client('rds')
        instances = rds.describe_db_instances()

        for db in instances['DBInstances']:
            encrypted = db.get('StorageEncrypted', False)
            self.add_result(
                control="CC6.6",
                resource=db['DBInstanceIdentifier'],
                check="RDS Encryption",
                passed=encrypted,
                severity="critical" if not encrypted else "pass"
            )

    def check_s3_encryption(self):
        """CC6.6 - All S3 buckets must have encryption"""
        s3 = boto3.client('s3')
        buckets = s3.list_buckets()

        for bucket in buckets['Buckets']:
            try:
                encryption = s3.get_bucket_encryption(Bucket=bucket['Name'])
                encrypted = True
            except:
                encrypted = False

            self.add_result(
                control="CC6.6",
                resource=bucket['Name'],
                check="S3 Encryption",
                passed=encrypted,
                severity="critical" if not encrypted else "pass"
            )

    def check_s3_public_access(self):
        """CC6.1 - No S3 buckets should be public"""
        s3 = boto3.client('s3')
        buckets = s3.list_buckets()

        for bucket in buckets['Buckets']:
            try:
                acl = s3.get_bucket_acl(Bucket=bucket['Name'])
                public = any(
                    grant['Grantee'].get('URI', '').endswith('AllUsers')
                    for grant in acl['Grants']
                )
            except:
                public = False

            self.add_result(
                control="CC6.1",
                resource=bucket['Name'],
                check="S3 Public Access",
                passed=not public,
                severity="critical" if public else "pass"
            )

    def check_backup_age(self):
        """CC7.1 - Backups must be recent"""
        rds = boto3.client('rds')
        snapshots = rds.describe_db_snapshots(SnapshotType='automated')

        for snapshot in snapshots['DBSnapshots']:
            age = datetime.utcnow() - snapshot['SnapshotCreateTime'].replace(tzinfo=None)
            recent = age < timedelta(hours=24)

            self.add_result(
                control="CC7.1",
                resource=snapshot['DBSnapshotIdentifier'],
                check="Backup Age",
                passed=recent,
                severity="warning" if not recent else "pass",
                details=f"Age: {age}"
            )

    def check_iam_mfa(self):
        """CC6.1 - All IAM users must have MFA"""
        iam = boto3.client('iam')
        users = iam.list_users()

        for user in users['Users']:
            mfa_devices = iam.list_mfa_devices(UserName=user['UserName'])
            has_mfa = len(mfa_devices['MFADevices']) > 0

            self.add_result(
                control="CC6.1",
                resource=user['UserName'],
                check="IAM MFA",
                passed=has_mfa,
                severity="critical" if not has_mfa else "pass"
            )

    def check_unused_credentials(self):
        """CC6.1 - No unused credentials > 90 days"""
        iam = boto3.client('iam')
        report = iam.generate_credential_report()
        # ... parse and check

    def add_result(self, control, resource, check, passed, severity, details=None):
        result = {
            "control": control,
            "resource": resource,
            "check": check,
            "passed": passed,
            "severity": severity,
            "details": details
        }
        self.results["checks"].append(result)

        if severity == "critical" and not passed:
            self.results["critical_count"] += 1
        elif severity == "warning" and not passed:
            self.results["warning_count"] += 1

    def run_all_checks(self):
        self.check_rds_encryption()
        self.check_s3_encryption()
        self.check_s3_public_access()
        self.check_backup_age()
        self.check_iam_mfa()
        return self.results

if __name__ == "__main__":
    checker = ComplianceChecker()
    results = checker.run_all_checks()
    print(json.dumps(results, indent=2))
```

## Evidence Collection

### Automated Evidence

| Evidence            | Collection Method | Frequency  |
| ------------------- | ----------------- | ---------- |
| IAM policies        | AWS Config        | Daily      |
| Encryption status   | Compliance script | Daily      |
| Access logs         | CloudTrail → S3   | Continuous |
| Change history      | Terraform state   | On change  |
| Vulnerability scans | Trivy/Snyk        | On build   |

### Manual Evidence

| Evidence         | Owner           | Frequency |
| ---------------- | --------------- | --------- |
| Access reviews   | Security team   | Quarterly |
| Penetration test | Third party     | Annual    |
| Policy reviews   | Compliance team | Annual    |
| Training records | HR              | Annual    |

## GDPR Compliance

### Data Subject Rights

| Right         | Implementation        |
| ------------- | --------------------- |
| Access        | Export API endpoint   |
| Erasure       | Delete workflow       |
| Portability   | Export in JSON format |
| Rectification | Update API            |

### Data Processing Records

```sql
CREATE TABLE data_processing_records (
    id UUID PRIMARY KEY,
    organization_id UUID NOT NULL,
    purpose VARCHAR(255) NOT NULL,
    legal_basis VARCHAR(100) NOT NULL,
    data_categories TEXT[],
    retention_period INTERVAL,
    created_at TIMESTAMPTZ NOT NULL
);
```

## Audit Trail

All compliance-relevant events are logged:

```json
{
  "event_type": "compliance_check",
  "timestamp": "2025-01-01T00:00:00Z",
  "check_type": "rds_encryption",
  "resource": "temporal-cloud-prod",
  "result": "pass",
  "control": "CC6.6",
  "evidence_url": "s3://compliance-evidence/2025/01/01/rds-encryption.json"
}
```

## Reporting

### Monthly Compliance Report

Generated automatically and sent to:

- Security team
- Compliance officer
- Engineering leadership

### Audit Preparation

Before SOC 2 audit:

1. Run full compliance check suite
2. Generate evidence package
3. Review and remediate findings
4. Prepare documentation index
