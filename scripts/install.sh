#!/bin/bash

### Ahenk Agent Local Installation Script

### --> Prerequisites 
# git
# go 1.20+
# systemd

# wget -qO- https://git.aliberksandikci.com.tr/liderahenk/ahenk-go/raw/branch/main/scripts/install.sh | bash

### --> Variable Set
REPO_NAME="ahenk-go"
REPO_LINK="https://git.aliberksandikci.com.tr/liderahenk/${REPO_NAME}"
DAEMON_NAME="ahenkd-go"
CONF_DIR="/etc/ahenk-go/"
TEMP_DIR="$(mktemp -d)"
MAIN_DIR="${TEMP_DIR}/${REPO_NAME}/"
echo -e "VARIABLES SET\n"


### --> Pre-Cleaning
sudo rm /usr/bin/${DAEMON_NAME}
sudo killall "$DAEMON_NAME"
sudo systemctl disable "$DAEMON_NAME"
sudo systemctl stop "$DAEMON_NAME"
sudo systemctl daemon-reload
sudo rm -rf "$CONF_DIR"
sudo mkdir -p "$CONF_DIR"
echo -e "PRE-CLENING DONE\n"


### --> Building
cd "$TEMP_DIR" || exit
echo "$TEMP_DIR"
git clone "$REPO_LINK"
cd "$MAIN_DIR" || exit

# cmd/ahenkd-go
cd cmd/ahenkd-go || exit 
sudo go build -o "/usr/bin/ahenkd-go"

echo -e "BUILDING DONE\n"

### --> Configurations
sudo touch "$CONF_DIR/user.info"
sudo cat << EOF | sudo tee -a "$CONF_DIR/user.info" > /dev/null
The current working directory is: $PWD
You are logged in as $(whoami)
$HOME $USER $SUDO_USER
EOF
sudo chown root:root "$CONF_DIR/user.info"
sudo chmod 640 "$CONF_DIR/user.info"
echo -e "CONFIGURATIONS DONE\n"

### --> Installing 
if [[ $(ps -p 1 | awk 'FNR == 2 {print $4} ') == "systemd" ]]
then
  sudo mv "${MAIN_DIR}build/package/deb/${DAEMON_NAME}.service" /etc/systemd/system/ 
fi
sudo systemctl daemon-reload
sudo systemctl enable "$DAEMON_NAME"
sudo systemctl start "$DAEMON_NAME"
echo -e "INSTALLING DONE\n"

### --> Post-Cleaning
sudo rm -rf "$TEMP_DIR"
echo -e "POST-CLEANING DONE\n"