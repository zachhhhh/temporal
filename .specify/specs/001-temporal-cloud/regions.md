# Supported Regions

## AWS Regions

| Region      | Code           | Cloud API Code     | PrivateLink | Same-Region HA | Multi-Region HA |
| ----------- | -------------- | ------------------ | ----------- | -------------- | --------------- |
| N. Virginia | us-east-1      | aws-us-east-1      | ✅          | ✅             | ✅              |
| Ohio        | us-east-2      | aws-us-east-2      | ✅          | ❌             | ✅              |
| Oregon      | us-west-2      | aws-us-west-2      | ✅          | ✅             | ✅              |
| Canada      | ca-central-1   | aws-ca-central-1   | ✅          | ❌             | ✅              |
| Ireland     | eu-west-1      | aws-eu-west-1      | ✅          | ✅             | ✅              |
| London      | eu-west-2      | aws-eu-west-2      | ✅          | ❌             | ✅              |
| Frankfurt   | eu-central-1   | aws-eu-central-1   | ✅          | ✅             | ✅              |
| Singapore   | ap-southeast-1 | aws-ap-southeast-1 | ✅          | ❌             | ✅              |
| Sydney      | ap-southeast-2 | aws-ap-southeast-2 | ✅          | ❌             | ✅              |
| Tokyo       | ap-northeast-1 | aws-ap-northeast-1 | ✅          | ❌             | ✅              |
| Seoul       | ap-northeast-2 | aws-ap-northeast-2 | ✅          | ❌             | ✅              |
| Mumbai      | ap-south-1     | aws-ap-south-1     | ✅          | ❌             | ✅              |
| Hyderabad   | ap-south-2     | aws-ap-south-2     | ✅          | ❌             | ✅              |
| São Paulo   | sa-east-1      | aws-sa-east-1      | ✅          | ❌             | ❌              |

## GCP Regions

| Region      | Code         | Cloud API Code   | PSC | Same-Region HA | Multi-Region HA |
| ----------- | ------------ | ---------------- | --- | -------------- | --------------- |
| Iowa        | us-central1  | gcp-us-central1  | ✅  | ❌             | ✅              |
| Oregon      | us-west1     | gcp-us-west1     | ✅  | ❌             | ✅              |
| N. Virginia | us-east4     | gcp-us-east4     | ✅  | ❌             | ✅              |
| Frankfurt   | europe-west3 | gcp-europe-west3 | ✅  | ❌             | ✅              |
| Mumbai      | asia-south1  | gcp-asia-south1  | ✅  | ❌             | ✅              |

## Multi-Region Replication Pairs

### US Regions

| Primary         | Secondary Options                              |
| --------------- | ---------------------------------------------- |
| aws-us-east-1   | aws-us-east-2, aws-us-west-2, aws-ca-central-1 |
| aws-us-west-2   | aws-us-east-1, aws-us-east-2, aws-ca-central-1 |
| gcp-us-central1 | gcp-us-west1, gcp-us-east4                     |

### EU Regions

| Primary          | Secondary Options               |
| ---------------- | ------------------------------- |
| aws-eu-west-1    | aws-eu-west-2, aws-eu-central-1 |
| aws-eu-central-1 | aws-eu-west-1, aws-eu-west-2    |

### APAC Regions

| Primary            | Secondary Options                                          |
| ------------------ | ---------------------------------------------------------- |
| aws-ap-northeast-1 | aws-ap-northeast-2, aws-ap-southeast-1, aws-ap-southeast-2 |
| aws-ap-south-1     | aws-ap-south-2, aws-ap-southeast-1                         |

## Multi-Cloud Replication

| AWS Region       | GCP Region                    |
| ---------------- | ----------------------------- |
| aws-us-east-1    | gcp-us-central1, gcp-us-east4 |
| aws-us-west-2    | gcp-us-west1                  |
| aws-eu-central-1 | gcp-europe-west3              |
| aws-ap-south-1   | gcp-asia-south1               |

## Regional Endpoints

### Format

```
<cloud-api-code>.region.tmprl.cloud
```

### Examples

```
aws-us-east-1.region.tmprl.cloud
gcp-us-central1.region.tmprl.cloud
```

## Latency Considerations

- Choose region closest to your workers
- Multi-region adds ~50-100ms latency for replication
- Cross-cloud replication may have higher latency
