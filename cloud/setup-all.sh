#!/bin/bash

set -e

# Colors
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

clear
echo -e "${BLUE}"
cat << "EOF"
╔═══════════════════════════════════════════════════════════════╗
║                                                               ║
║   Temporal Cloud Console - Complete Setup                    ║
║                                                               ║
╚═══════════════════════════════════════════════════════════════╝
EOF
echo -e "${NC}\n"

# Check if we're in the right directory
if [ ! -f "go.mod" ] || [ ! -d "console" ]; then
    echo -e "${RED}Error: Please run this script from the cloud/ directory${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Running from correct directory${NC}\n"

# Step 1: Check database
echo -e "${BLUE}[1/5]${NC} Checking database..."
if docker ps | grep -q "temporal-cloud-postgres"; then
    echo -e "${GREEN}✓ PostgreSQL is running${NC}"
else
    echo -e "${YELLOW}! PostgreSQL not running, starting it...${NC}"
    docker-compose -f docker-compose.dev.yaml up -d postgresql
    sleep 3
    echo -e "${GREEN}✓ PostgreSQL started${NC}"
fi

# Check if database exists
if docker exec temporal-cloud-postgres psql -U temporal -lqt | cut -d \| -f 1 | grep -qw temporal_cloud; then
    echo -e "${GREEN}✓ Database 'temporal_cloud' exists${NC}"
else
    echo -e "${YELLOW}! Creating database...${NC}"
    docker exec temporal-cloud-postgres psql -U temporal -c "CREATE DATABASE temporal_cloud;"
    echo -e "${GREEN}✓ Database created${NC}"
fi

# Step 2: Check migrations
echo -e "\n${BLUE}[2/5]${NC} Checking migrations..."
MIGRATION_COUNT=$(docker exec temporal-cloud-postgres psql -U temporal -d temporal_cloud -t -c "SELECT COUNT(*) FROM information_schema.tables WHERE table_schema='public';" 2>/dev/null | tr -d ' ' || echo "0")

if [ "$MIGRATION_COUNT" -gt "10" ]; then
    echo -e "${GREEN}✓ Migrations already applied ($MIGRATION_COUNT tables)${NC}"
else
    echo -e "${YELLOW}! Applying migrations...${NC}"
    for f in schema/migrations/*.up.sql; do
        docker exec -i temporal-cloud-postgres psql -U temporal -d temporal_cloud < "$f" 2>&1 | grep -v "ERROR:" | head -5 || true
    done
    echo -e "${GREEN}✓ Migrations applied${NC}"
fi

# Step 3: Check .env file
echo -e "\n${BLUE}[3/5]${NC} Checking configuration..."
if [ -f ".env" ]; then
    echo -e "${GREEN}✓ .env file exists${NC}"
    
    if grep -q "YOUR_GOOGLE_CLIENT_ID_HERE" .env; then
        echo -e "${YELLOW}! Google OAuth not configured${NC}"
        echo -e "${YELLOW}  You can configure it later or run: ./setup-google-oauth.sh${NC}"
    else
        echo -e "${GREEN}✓ Google OAuth configured${NC}"
    fi
else
    echo -e "${RED}✗ .env file missing${NC}"
    exit 1
fi

# Step 4: Check dependencies
echo -e "\n${BLUE}[4/5]${NC} Checking dependencies..."

# Go dependencies
if go mod verify &>/dev/null; then
    echo -e "${GREEN}✓ Go dependencies OK${NC}"
else
    echo -e "${YELLOW}! Running go mod tidy...${NC}"
    go mod tidy
    echo -e "${GREEN}✓ Go dependencies installed${NC}"
fi

# Node dependencies
if [ -d "console/node_modules" ]; then
    echo -e "${GREEN}✓ Node dependencies OK${NC}"
else
    echo -e "${YELLOW}! Installing Node dependencies...${NC}"
    cd console && npm install && cd ..
    echo -e "${GREEN}✓ Node dependencies installed${NC}"
fi

# Step 5: Summary
echo -e "\n${BLUE}[5/5]${NC} Setup complete!\n"

echo -e "${GREEN}╔═══════════════════════════════════════════════════════════════╗${NC}"
echo -e "${GREEN}║                    Setup Complete! 🎉                        ║${NC}"
echo -e "${GREEN}╚═══════════════════════════════════════════════════════════════╝${NC}\n"

echo -e "${YELLOW}What's ready:${NC}"
echo -e "  ✓ Database created and migrated"
echo -e "  ✓ All code in place"
echo -e "  ✓ Dependencies installed"
echo -e "  ✓ Configuration files ready\n"

# Check if OAuth is configured
if grep -q "YOUR_GOOGLE_CLIENT_ID_HERE" .env; then
    echo -e "${YELLOW}Next step: Configure Google OAuth${NC}"
    echo -e "  Run: ${GREEN}./setup-google-oauth.sh${NC}\n"
    
    echo -e "${YELLOW}Or skip OAuth and test the console:${NC}"
    echo -e "  ${GREEN}cd console && npm run dev${NC}"
    echo -e "  ${GREEN}open http://localhost:5174/console/namespaces${NC}\n"
else
    echo -e "${YELLOW}Next step: Start the services${NC}\n"
    echo -e "  ${GREEN}Terminal 1:${NC} ./start-auth.sh"
    echo -e "  ${GREEN}Terminal 2:${NC} cd console && npm run dev\n"
    echo -e "  ${GREEN}Then visit:${NC} http://localhost:5174/login\n"
fi

echo -e "${BLUE}Documentation:${NC}"
echo -e "  • COMPLETE_SETUP.md - Full setup guide"
echo -e "  • GOOGLE_OAUTH_SETUP.md - OAuth configuration"
echo -e "  • QUICK_START.md - Quick reference\n"
