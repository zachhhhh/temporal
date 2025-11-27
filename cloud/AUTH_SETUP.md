# Google OAuth Authentication Setup

## Prerequisites

1. **Google Cloud Console Setup**
   - Go to https://console.cloud.google.com/
   - Create a new project or select existing
   - Enable Google+ API
   - Go to "Credentials" → "Create Credentials" → "OAuth 2.0 Client ID"
   - Application type: "Web application"
   - Authorized redirect URIs: `http://localhost:8081/auth/google/callback`
   - Copy the Client ID and Client Secret

## Environment Variables

Create a `.env` file in the `cloud/` directory:

```bash
# Google OAuth
GOOGLE_CLIENT_ID=your-client-id-here.apps.googleusercontent.com
GOOGLE_CLIENT_SECRET=your-client-secret-here

# JWT Configuration
JWT_SECRET_KEY=your-super-secret-jwt-key-change-this-in-production
JWT_ISSUER=temporal-cloud
JWT_AUDIENCE=temporal-cloud-console

# Server Configuration
CLOUD_API_PORT=8081
CONSOLE_URL=http://localhost:5174
```

## Database Setup

Run the migrations to create the user tables:

```bash
cd cloud/schema/migrations

# Create the temporal_cloud database if it doesn't exist
psql -h localhost -U temporal -c "CREATE DATABASE temporal_cloud;"

# Run migrations
for f in *.up.sql; do
  psql -h localhost -U temporal -d temporal_cloud -f "$f"
done
```

## Running the Services

### 1. Start the Cloud API (Backend)

```bash
cd cloud
go run ./cmd/cloud-api
```

The API will start on `http://localhost:8081`

### 2. Start the Console (Frontend)

```bash
cd cloud/console
npm run dev
```

The console will start on `http://localhost:5174`

## Testing Authentication

1. Navigate to `http://localhost:5174/login`
2. Click "Continue with Google"
3. Sign in with your Google account
4. You'll be redirected back to `/console/namespaces`

## API Endpoints

- `GET /auth/google/login` - Initiate Google OAuth flow
- `GET /auth/google/callback` - Handle OAuth callback
- `POST /auth/refresh` - Refresh access token
- `POST /auth/logout` - Logout user

## Security Notes

- In production, use HTTPS for all endpoints
- Set `Secure: true` on cookies
- Use a strong JWT secret key
- Configure proper CORS settings
- Add rate limiting to auth endpoints

## Troubleshooting

### "redirect_uri_mismatch" error

- Ensure the redirect URI in Google Cloud Console exactly matches `http://localhost:8081/auth/google/callback`

### "missing state cookie" error

- Check that cookies are enabled in your browser
- Ensure the Cloud API and Console are on compatible domains

### Database connection errors

- Verify PostgreSQL is running: `docker ps | grep postgres`
- Check connection string in Cloud API configuration
