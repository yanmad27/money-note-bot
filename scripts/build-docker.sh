#!/bin/sh

CONTAINER='money-note-bot'

GREEN='\033[0;32m'
NC='\033[0m'

set -x

{
    echo ${GREEN} Build $CONTAINER starting ...${NC} &&
    git stash  &&
    git pull &&
    git config core.pager cat && 
    echo Last commit message: $(git log -1 --pretty=%B) &&
    docker compose up -d --build &&
    echo ${GREEN} Build $CONTAINER DONE${NC}
}
