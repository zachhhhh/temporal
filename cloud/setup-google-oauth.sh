#!/bin/bash

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

clear
echo -e "${BLUE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║${NC}  ${GREEN}Google OAuth Setup for Temporal Cloud Console${NC}          ${BLUE}║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════════════════════════════╝${NC}\n"

echo -e "${YELLOW}I'll guide you through setting up Google OAuth credentials.${NC}\n"

# Step 1: Open Google Cloud Console
echo -e "${GREEN}Step 1: Opening Google Cloud Console...${NC}"
echo -e "Opening: ${BLUE}https://console.cloud.google.com/apis/credentials${NC}\n"
open "https://console.cloud.google.com/apis/credentials" 2>/dev/null || xdg-open "https://console.cloud.google.com/apis/credentials" 2>/dev/null

echo -e "${YELLOW}Please follow these steps in the browser:${NC}"
echo -e "  1. Sign in to Google Cloud Console"
echo -e "  2. Select or create a project"
echo -e "  3. Click '+ CREATE CREDENTIALS' → 'OAuth 2.0 Client ID'"
echo -e "  4. If prompted, configure the consent screen:"
echo -e "     - App name: ${GREEN}Temporal Cloud Console${NC}"
echo -e "     - User support email: ${GREEN}Your email${NC}"
echo -e "     - Add your email as a test user"
echo -e "  5. Select Application type: ${GREEN}Web application${NC}"
echo -e "  6. Name: ${GREEN}Temporal Cloud Console${NC}"
echo -e "  7. Add Authorized redirect URI: ${GREEN}http://localhost:8081/auth/google/callback${NC}"
echo -e "  8. Click ${GREEN}CREATE${NC}\n"

read -p "Press Enter when you've created the OAuth credentials..."

# Step 2: Get credentials
echo -e "\n${GREEN}Step 2: Enter your OAuth credentials${NC}\n"
echo -e "Copy the ${YELLOW}Client ID${NC} from the dialog (looks like: xxxxx.apps.googleusercontent.com)"
read -p "Client ID: " CLIENT_ID

echo -e "\nCopy the ${YELLOW}Client Secret${NC} from the dialog"
read -p "Client Secret: " CLIENT_SECRET

# Step 3: Update .env file
echo -e "\n${GREEN}Step 3: Updating .env file...${NC}"

if [ -f ".env" ]; then
    # Update existing .env
    sed -i.bak "s|GOOGLE_CLIENT_ID=.*|GOOGLE_CLIENT_ID=$CLIENT_ID|" .env
    sed -i.bak "s|GOOGLE_CLIENT_SECRET=.*|GOOGLE_CLIENT_SECRET=$CLIENT_SECRET|" .env
    rm .env.bak
    echo -e "${GREEN}✓ Updated .env file${NC}"
else
    echo -e "${RED}✗ .env file not found!${NC}"
    exit 1
fi

# Step 4: Verify setup
echo -e "\n${GREEN}Step 4: Verifying setup...${NC}"

if grep -q "YOUR_GOOGLE_CLIENT_ID_HERE" .env; then
    echo -e "${RED}✗ Client ID not updated${NC}"
    exit 1
fi

if grep -q "YOUR_GOOGLE_CLIENT_SECRET_HERE" .env; then
    echo -e "${RED}✗ Client Secret not updated${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Google OAuth configured successfully!${NC}\n"

# Step 5: Start services
echo -e "${BLUE}╔══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${BLUE}║${NC}  ${GREEN}Setup Complete!${NC}                                          ${BLUE}║${NC}"
echo -e "${BLUE}╚══════════════════════════════════════════════════════════════╝${NC}\n"

echo -e "${YELLOW}Your configuration:${NC}"
echo -e "  Client ID: ${GREEN}${CLIENT_ID}${NC}"
echo -e "  Redirect URI: ${GREEN}http://localhost:8081/auth/google/callback${NC}\n"

echo -e "${YELLOW}Next steps:${NC}"
echo -e "  1. Start the Cloud API: ${GREEN}./start-auth.sh${NC}"
echo -e "  2. Start the Console: ${GREEN}cd console && npm run dev${NC}"
echo -e "  3. Visit: ${BLUE}http://localhost:5174/login${NC}\n"

read -p "Would you like to start the Cloud API now? (y/n) " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]; then
    echo -e "\n${GREEN}Starting Cloud API...${NC}\n"
    ./start-auth.sh
fi
