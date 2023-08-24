# Ahenk Go

**Liderahenk** is an open source project which provides solutions to manage, monitor and audit unlimited number of different systems and users on a network.

Ahenk-go is a Linux agent written in Go which enables Lider to manage & monitor clients remotely.

### Creating Docker Development Environment (Optional)
1. Install and Create docker environment, [Quick Start](https://sulincix.github.io/sayfalar/html/docker-kullanimi.html)
2. Pull docker image
```sh
docker pull asandikci/godeb
```
> Or alternatively build Dockerfile with `docker build -t godeb:latest` code for lower data usage

3. Create container from image and attach to container, an example:
```sh
docker run -it -d --name build1 asandikci/godeb:latest "bin/bash"
docker attach build1 --detach-keys "ctrl-k"
```
## Packaging
1. Install necessary packages (You dont need this step if you are using docker file)
```sh
sudo apt install sudo dpkg-dev debhelper golang-any 
``` 
2. Clone the repository and move to main directory
```sh
git clone https://git.aliberksandikci.com.tr/Liderahenk/ahenk-go/
cd ahenk-go/
```
3. Build program and create binary/source packages
```sh
dpkg-buildpackage
```
> Refer to Makefile for more info

### Documentation
- See [Liderahenk/ahenk-docs](https://git.aliberksandikci.com.tr/Liderahenk/ahenk-docs/) for comprehensively prepared documentation.


## Other Liderahenk Projects
- [Pardus-LiderAhenk/ahenk](https://github.com/Pardus-LiderAhenk/ahenk/): Current Python Implementation of Ahenk 
- [Lider UI](https://github.com/Pardus-LiderAhenk/liderui): Lider Administration User Interface
- [Lider API](https://github.com/Pardus-LiderAhenk/liderapi): Lider API
- [Liderahenk-ansible](https://github.com/Pardus-LiderAhenk/liderahenk-ansible): Liderahenk setup with ansible  
- [Ahenkdesk](https://github.com/Pardus-LiderAhenk/ahenkdesk): Ahenk User Interface
- See more in our [GitHub Page](https://github.com/Pardus-LiderAhenk)


## Changelog
See [changelog](./debian/changelog) to learn what have been changed between releases or refer to [Releases](./releases) Tab above


## License
Lider Ahenk and its sub projects are licensed under the [LGPL v3](./LICENSE).

