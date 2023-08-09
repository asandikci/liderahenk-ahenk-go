#!/bin/bash

# see https://people.debian.org/~stapelberg/2015/07/27/dh-make-golang.html


GIT_HOST="git.aliberksandikci.com.tr"
GIT_ORG="Liderahenk"
GIT_REPO="ahenk-go"


mkdir /build
cd /build || exit
dh-make-golang make -allow_unknown_hoster "$GIT_HOST/$GIT_ORG/$GIT_REPO"

# TODO Add option to disable writing all these files all the time (-s --skip)
nano itp-ahenk-go.txt
sendmail -t < itp-ahenk-go.txt

cd ahenk-go || exit
grep --color=always -r TODO debian
echo -e "\nThese files needs review. Starting reviewing automatically in 10sec...\n"
sleep 10
# TODO Allow users to escape from automatic review and manually do it while in this script

# Edit files that contains TODO keyword until there is none of them
while [[ $(grep -r TODO debian | wc --lines) -ne 0 ]]
do
  nano "$(grep -r TODO debian | awk '{sub(/:.*/,"")} NR==1')"

  echo "Continuing from next file..."
  sleep 1
  grep --color=always -r TODO debian
  sleep 5

  # TODO Allow user to reviewing remaining files, choosing between them, optional skipping, and waiting for an answer (continue? [Y/n])
done
echo "ALL FILES DONE"

sleep 5 

head -100 debian/**/*

# TODO ask user to continue

git add debian && git commit -a -m 'Initial packaging'