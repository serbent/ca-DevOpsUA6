# Create an App Registration
az ad app create \
  --display-name github-actions-app

  $APP_ID="xxxxxxx-1176-413d-8fxxxx30-b09974509c2d",

# Create a Service Principal

az ad sp create --id <APP_ID>

# Configure Federated Credentials (OIDC)

az ad app federated-credential create \
  --id <APP_ID> \
  --parameters '{
    "name": "github",
    "issuer": "https://token.actions.githubusercontent.com",
    "subject": "repo:OWNER/REPO:ref:refs/heads/main",
    "audiences": ["api://AzureADTokenExchange"]
  }'

# Assign a Role
az role assignment create \
  --assignee <APP_ID> \
  --role Contributor \
  --scope /subscriptions/<SUBSCRIPTION_ID>

# Get tennantID
az account show --query tenantId -o tsv

# Get subscriptionID
az account show --query id -o tsv


# Create a GitHub Secrets
AZURE_CLIENT_ID <- APP_ID
AZURE_SUBSCRIPTION_ID <- subscriptionID
AZURE_TENANT_ID <- tennantID

# Use in github actions
```yaml
- uses: azure/login@v2
  with:
    client-id: ${{ secrets.AZURE_CLIENT_ID }}
    tenant-id: ${{ secrets.AZURE_TENANT_ID }}
    subscription-id: ${{ secrets.AZURE_SUBSCRIPTION_ID }}
```