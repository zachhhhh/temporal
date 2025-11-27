#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}🚀 Temporal Cloud Console - Authentication Setup${NC}\n"

# Check if .env exists
if [ ! -f ".env" ]; then
    echo -e "${RED}❌ .env file not found!${NC}"
    echo -e "${YELLOW}Please copy .env.example to .env and add your Google OAuth credentials${NC}"
    exit 1
fi

# Check if Google credentials are set
if grep -q "YOUR_GOOGLE_CLIENT_ID_HERE" .env; then
    echo -e "${YELLOW}⚠️  Google OAuth credentials not configured!${NC}"
    echo -e "${YELLOW}Please follow the guide in GOOGLE_OAUTH_SETUP.md${NC}"
    echo -e "${YELLOW}Or visit: https://console.cloud.google.com/apis/credentials${NC}\n"
    read -p "Press Enter to continue anyway (will fail at login) or Ctrl+C to exit..."
fi

# Check if database is running
if ! docker ps | grep -q "temporal-cloud-postgres"; then
    echo -e "${RED}❌ PostgreSQL container not running!${NC}"
    echo -e "${YELLOW}Starting PostgreSQL...${NC}"
    docker-compose -f docker-compose.dev.yaml up -d postgresql
    sleep 3
fi

# Check if migrations are needed
echo -e "${GREEN}✓ Database is running${NC}"

# Start the Cloud API
echo -e "\n${GREEN}Starting Cloud API on port 8081...${NC}"
echo -e "${YELLOW}Press Ctrl+C to stop${NC}\n"

go run ./cmd/cloud-api
