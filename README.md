# Ahenk Go

**Liderahenk** is an open source project which provides solutions to manage, monitor and audit unlimited number of different systems and users on a network.

Ahenk-go is a Linux agent written in Go which enables Lider to manage & monitor clients remotely.

## Packaging
1. Install necessary packages (You dont need this step if you are using docker file [godeb](https://git.aliberksandikci.com.tr/Liderahenk/ahenk-docs/src/branch/main/dev/environment.md#creating-docker-development-environment))
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
4. *optional*, look for package errors and warnings
```sh
lintian
``` 

## Documentation & Changelog
- See comprehensively prepared [documentation](https://git.aliberksandikci.com.tr/Liderahenk/ahenk-docs/)
- See how to [setup development environment](https://git.aliberksandikci.com.tr/Liderahenk/ahenk-docs/src/branch/main/dev/environment.md)<!-- Web Documentation Link Here -->
- See [changelog](https://git.aliberksandikci.com.tr/Liderahenk/ahenk-docs/src/branch/main/admin/changelog.md)
<!-- Web Changelog Link Here -->

## Other Liderahenk Projects
- [Pardus-LiderAhenk/ahenk](https://github.com/Pardus-LiderAhenk/ahenk/): Current Python Implementation of Ahenk 
- [Lider UI](https://github.com/Pardus-LiderAhenk/liderui): Lider Administration User Interface
- [Lider API](https://github.com/Pardus-LiderAhenk/liderapi): Lider API
- [Liderahenk-ansible](https://github.com/Pardus-LiderAhenk/liderahenk-ansible): Liderahenk setup with ansible  
- [Ahenkdesk](https://github.com/Pardus-LiderAhenk/ahenkdesk): Ahenk User Interface
- See more in our [GitHub Page](https://github.com/Pardus-LiderAhenk)

## License
Liderahenk and its sub projects are licensed under the [LGPL v3](./LICENSE).

