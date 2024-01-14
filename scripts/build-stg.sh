#!/bin/sh

set -x

ssh root@112.197.2.168 <<'ENDSSH'

cd ~/workspace/money-note-bot

WEBSITE='money-note-bot'

sh scripts/build-docker.sh

if [ $? -eq 0 ] ; then
    SUCCESS_MASSAGE="✅ SUCCESS: build $WEBSITE successfully!"
    sh scripts/notify.sh "${SUCCESS_MASSAGE}"
else
    ERROR_MESSAGE="❌ FAILED: Build $WEBSITE failed" 
    sh scripts/notify.sh "${ERROR_MESSAGE}"
fi


ENDSSH