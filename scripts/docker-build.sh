#!/bin/bash

# see https://people.debian.org/~stapelberg/2015/07/27/dh-make-golang.html


GIT_HOST="git.aliberksandikci.com.tr"
GIT_ORG="Liderahenk"
GIT_REPO="ahenk-go"


mkdir /build
cd /build || exit
dh-make-golang make -allow_unknown_hoster "$GIT_HOST/$GIT_ORG/$GIT_REPO"

nano itp-ahenk-go.txt
sendmail -t < itp-ahenk-go.txt

cd ahenk-go || exit
grep -r TODO debian
head -100 debian/**/*


