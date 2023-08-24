# docker build -t IMAGE_NAME:VERSION .

# after cd ahenk-go
# docker run -it -d --name CONT_NAME --mount type=bind,source=".",target=/ahenk-go/ IMAGE_NAME:VERSION "/bin/bash"

FROM bitnami/minideb

RUN apt-get update
RUN apt-get upgrade -y

### INSTALL USEFUL UTILITIES ###
RUN apt-get install curl wget tree nano procps lintian -y 

### INSTALL BUILD DEPENDENCIES ###
RUN apt-get install sudo dpkg-dev debhelper golang-any -y

### CONFIGURATIONS ###
# Enable UTF8 Encoding
RUN echo "LC_ALL=en_US.UTF-8" >> /etc/environment
RUN echo "en_US.UTF-8 UTF-8" >> /etc/locale.gen
RUN echo "tr_TR.UTF-8 UTF-8" >> /etc/locale.gen
RUN echo "LANG=en_US.UTF-8" > /etc/locale.conf
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8  
RUN apt install locales -y
RUN locale-gen en_US.UTF-8