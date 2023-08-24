#!/bin/bash

# see https://people.debian.org/~stapelberg/2015/07/27/dh-make-golang.html

GIT_HOST="git.aliberksandikci.com.tr"
GIT_ORG="liderahenk"
GIT_REPO="ahenk-go"
FULL_URL="https://$GIT_HOST/$GIT_ORG/$GIT_REPO"
DEV_NAME="Aliberk Sandıkçı"
GIT_MAIL="git@aliberksandikci.com.tr"

UPLOAD_URL="127.0.0.1"
UPLOAD_PORT="8000"

echo "Editing git configs..."
git config --global init.defaultBranch main
git config --global user.name "$DEV_NAME"
git config --global user.email $GIT_MAIL
git config --global color.status auto
git config --global color.branch auto
git config --global color.interactive auto
git config --global color.diff auto

mkdir /build
cd /build || exit

if [[ "$1" != "update" ]]; then

  dh-make-golang make -allow_unknown_hoster "$GIT_HOST/$GIT_ORG/$GIT_REPO"

  if [[ ! ("$1" != "notodo" || "$2" != "notodo") ]]; then

    # TODO Add option to disable writing all these files all the time (-s --skip)
    nano "itp-$GIT_REPO.txt"
    sendmail -t <"itp-$GIT_REPO.txt"

    cd $GIT_REPO || exit
    grep --color=always -r TODO debian
    echo -e "\nThese files needs review. Starting reviewing automatically in 10sec...\n"
    sleep 10
    # TODO Allow users to escape from automatic review and manually do it while in this script

    while [[ $(grep -r TODO debian | wc --lines) -ne 0 ]]; do
      nano "$(grep -r TODO debian | awk '{sub(/:.*/,"")} NR==1')"

      echo "Continuing from next file..."
      sleep 1
      grep --color=always -r TODO debian
      sleep 5

      # TODO Allow user to reviewing remaining files, choosing between them, optional skipping, and waiting for an answer (continue? [Y/n])
    done
    echo "ALL FILES DONE"

    dh_installsystemd
    echo """
#!/usr/bin/make -f

%:
	dh \$@ --builddirectory=_build --buildsystem=golang --with=golang --with-systemd
override_dh_auto_install:
	dh_auto_install -- --no-source
""" >debian/rules
    sleep 5
    # FIXME Not working!

  else
    cd $GIT_REPO || exit
  fi
else
  gbp clone "$FULL_URL"
  cd $GIT_REPO || exit
  git checkout debian/sid
  pwd
  if [[ "$1" == "nmu" ]]; then
    dch --nmu
  else
    dch
  fi
  echo "Succesfully Updated deb package information"
fi

# TODO ask user to continue

git add debian && git commit -a -m 'Initial packaging'

gbp buildpackage

cd ..
lintian -- *.changes

echo -e "SOLVE LINTIAN ERRORS / WARNINGS\nAFTER THAT, PUSH REPO"
