# Google OAuth Setup Guide

## Step 1: Create Google Cloud Project

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Click on the project dropdown at the top
3. Click "New Project"
4. Enter project name: `Temporal Cloud Console` (or any name you prefer)
5. Click "Create"

## Step 2: Enable Google+ API

1. In the left sidebar, go to **APIs & Services** → **Library**
2. Search for "Google+ API"
3. Click on it and click "Enable"

## Step 3: Configure OAuth Consent Screen

1. Go to **APIs & Services** → **OAuth consent screen**
2. Select **External** (unless you have a Google Workspace)
3. Click "Create"
4. Fill in the required fields:
   - **App name**: `Temporal Cloud Console`
   - **User support email**: Your email
   - **Developer contact email**: Your email
5. Click "Save and Continue"
6. On the Scopes page, click "Save and Continue" (we'll add scopes in credentials)
7. On Test users page, click "Add Users" and add your Google email
8. Click "Save and Continue"

## Step 4: Create OAuth 2.0 Credentials

1. Go to **APIs & Services** → **Credentials**
2. Click "Create Credentials" → "OAuth 2.0 Client ID"
3. If prompted, configure the consent screen (follow Step 3)
4. Select **Application type**: "Web application"
5. Enter **Name**: `Temporal Cloud Console`
6. Under **Authorized redirect URIs**, click "Add URI" and enter:
   ```
   http://localhost:8081/auth/google/callback
   ```
7. Click "Create"
8. A dialog will appear with your credentials:
   - **Client ID**: Copy this (looks like `xxxxx.apps.googleusercontent.com`)
   - **Client Secret**: Copy this

## Step 5: Update .env File

1. Open `/Users/zach/Developments/temporal/cloud/.env`
2. Replace the placeholder values:
   ```bash
   GOOGLE_CLIENT_ID=paste-your-client-id-here.apps.googleusercontent.com
   GOOGLE_CLIENT_SECRET=paste-your-client-secret-here
   ```
3. Save the file

## Step 6: Verify Setup

Your `.env` file should now look like this:

```bash
# Google OAuth Configuration
GOOGLE_CLIENT_ID=123456789-abcdefg.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=GOCSPX-abc123def456
JWT_SECRET_KEY=e5LZ4YonE5a+I3l+/0L8e5TxoZpexpPQfLNx9fkhbgA=
JWT_ISSUER=temporal-cloud
JWT_AUDIENCE=temporal-cloud-console
CLOUD_API_PORT=8081
CONSOLE_URL=http://localhost:5174
DB_HOST=localhost
DB_PORT=5432
DB_USER=temporal
DB_PASSWORD=temporal
DB_NAME=temporal_cloud
```

## Step 7: Start the Services

### Terminal 1: Start Cloud API

```bash
cd /Users/zach/Developments/temporal/cloud
go run ./cmd/cloud-api
```

### Terminal 2: Start Console (if not already running)

```bash
cd /Users/zach/Developments/temporal/cloud/console
npm run dev
```

## Step 8: Test Login

1. Open browser to: http://localhost:5174/login
2. Click "Continue with Google"
3. Sign in with the Google account you added as a test user
4. You should be redirected to http://localhost:5174/console/namespaces

## Troubleshooting

### Error: "redirect_uri_mismatch"

- **Solution**: Make sure the redirect URI in Google Cloud Console exactly matches:
  ```
  http://localhost:8081/auth/google/callback
  ```
  (No trailing slash, correct port)

### Error: "Access blocked: This app's request is invalid"

- **Solution**: Make sure you've added your email as a test user in the OAuth consent screen

### Error: "missing state cookie"

- **Solution**:
  - Clear your browser cookies
  - Make sure both Cloud API (8081) and Console (5174) are running
  - Try in incognito mode

### Error: "Failed to initiate login"

- **Solution**:
  - Check that Cloud API is running on port 8081
  - Check browser console for CORS errors
  - Verify `.env` file is in the correct location

## Production Deployment

When deploying to production:

1. Update redirect URI in Google Cloud Console to your production URL:

   ```
   https://your-domain.com/auth/google/callback
   ```

2. Update `.env`:

   ```bash
   CONSOLE_URL=https://your-domain.com
   ```

3. Change OAuth consent screen from "Testing" to "In Production"

4. Use HTTPS for all endpoints (required by Google OAuth)
