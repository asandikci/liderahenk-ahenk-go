ARG CODE_VERSION=latest
FROM debian:${CODE_VERSION}

RUN su -c "echo 'deb http://ftp.tr.debian.org/debian sid main' >> /etc/apt/sources.list.d/sid.list"

RUN apt-get update
RUN apt-get upgrade -y

### INSTALL USEFUL UTILITIES ###
RUN apt-get install curl wget tree nano -y

### INSTALL BUILD DEPENDENCIES ###
RUN apt-get install pbuilder -y
RUN apt-get install dh-make-golang -y
RUN apt-get install cowbuilder -y
RUN apt-get install golang-github-sevlyar-go-daemon-dev -y
RUN apt-get install golang-golang-x-exp-dev -y

### CONFIGURATIONS ###
RUN mkdir /var/cache/pbuilder/base.cow
# RUN cowbuilder --create
# FIXME Build Error

RUN echo "LC_ALL=en_US.UTF-8" >> /etc/environment
RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen
RUN echo "LANG=en_US.UTF-8" > /etc/locale.conf
RUN locale-gen en_US.UTF-8