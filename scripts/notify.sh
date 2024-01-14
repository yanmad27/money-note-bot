#!/bin/sh
# set -x

URL="https://api.telegram.org/bot6282193475:AAGZYnIStss6ZncBJUm7KOzwskixyfj6Da4/sendMessage?chat_id=-1001508349136"

# URL=$(sed "s/ /%20/g" <<< $URL)
# URL=$(sed "s/\[/%5B/g" <<< $URL)
# URL=$(sed "s/]/%7D/g" <<< $URL)
# URL=$(sed "s/{/%7B/g" <<< $URL)
# URL=$(sed "s/}/%7D/g" <<< $URL)
echo $1
curl $URL --data-urlencode "text=$1"
