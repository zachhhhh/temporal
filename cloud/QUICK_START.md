# 🚀 Quick Start - Google OAuth Setup

## ✅ What's Already Done

- ✓ Database migrations applied
- ✓ `.env` file created with JWT secret
- ✓ Auth service implemented
- ✓ Login page created
- ✓ All backend code ready

## 📋 What You Need to Do (5 minutes)

### Option 1: Automated Setup (Recommended)

Run this command and follow the prompts:

```bash
cd /Users/zach/Developments/temporal/cloud
./setup-google-oauth.sh
```

The script will:

1. Open Google Cloud Console
2. Guide you through creating OAuth credentials
3. Automatically update your `.env` file
4. Start the services

### Option 2: Manual Setup

#### Step 1: Create OAuth Credentials

The Google Cloud Console should be open in your browser. If not, visit:
https://console.cloud.google.com/apis/credentials

1. **Sign in** with your Google account

2. **Select or create a project**:

   - Click the project dropdown at the top
   - Click "New Project"
   - Name: "Temporal Cloud Console"
   - Click "Create"

3. **Configure OAuth Consent Screen** (if first time):

   - Click "OAuth consent screen" in the left sidebar
   - Select "External"
   - Fill in:
     - App name: `Temporal Cloud Console`
     - User support email: Your email
     - Developer contact: Your email
   - Click "Save and Continue" through all steps
   - Add yourself as a test user

4. **Create OAuth 2.0 Client ID**:

   - Click "Credentials" in the left sidebar
   - Click "+ CREATE CREDENTIALS"
   - Select "OAuth 2.0 Client ID"
   - Application type: **Web application**
   - Name: `Temporal Cloud Console`
   - Authorized redirect URIs: Click "ADD URI"
     ```
     http://localhost:8081/auth/google/callback
     ```
   - Click "CREATE"

5. **Copy your credentials**:
   - A dialog will show your Client ID and Client Secret
   - Keep this dialog open or download the JSON

#### Step 2: Update .env File

Edit `/Users/zach/Developments/temporal/cloud/.env`:

```bash
GOOGLE_CLIENT_ID=paste-your-client-id-here.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=paste-your-client-secret-here
```

#### Step 3: Start the Services

**Terminal 1 - Cloud API:**

```bash
cd /Users/zach/Developments/temporal/cloud
go run ./cmd/cloud-api
```

**Terminal 2 - Console:**

```bash
cd /Users/zach/Developments/temporal/cloud/console
npm run dev
```

#### Step 4: Test Login

1. Open: http://localhost:5174/login
2. Click "Continue with Google"
3. Sign in with your Google account
4. You'll be redirected to: http://localhost:5174/console/namespaces

## 🎉 You're Done!

Your Temporal Cloud Console is now running with Google OAuth authentication!

## 🔧 Troubleshooting

### "redirect_uri_mismatch"

- Make sure the redirect URI is exactly: `http://localhost:8081/auth/google/callback`
- No trailing slash
- Correct port (8081, not 5174)

### "Access blocked"

- Add your email as a test user in OAuth consent screen

### "Failed to initiate login"

- Check that Cloud API is running on port 8081
- Check `.env` file has correct credentials

### Still having issues?

Run the diagnostic:

```bash
cd /Users/zach/Developments/temporal/cloud
./start-auth.sh
```

It will check your configuration and show any errors.
