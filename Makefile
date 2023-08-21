REPO_NAME="ahenk-go"
REPO_LINK="https://git.aliberksandikci.com.tr/liderahenk/${REPO_NAME}"

DAEMON_NAME="ahenk-go"
OLD_DAEMON_NAME="ahenkd-go"
PYTHON_DAEMON_NAME="ahenkd"

CONF_DIR="/etc/ahenk-go/"
TEMP_DIR="$(mktemp -d)"
MAIN_DIR="${TEMP_DIR}/${REPO_NAME}/"

info:
	@echo "Made by Aliberk Sandıkçı - 2023"
	@echo "preclean: for cleaning old files, configurations"
	@echo "TODO test: Test go files"
	@echo "install: Build and install ahenk-go"
	@echo "TODO uninstall: Uninstall ahenk-go"
	@echo "TODO clean: for postclean"

preclean:
	sudo rm -rf /usr/bin/$(DAEMON_NAME)
	sudo rm -rf /usr/bin/${OLD_DAEMON_NAME}
	sudo rm -rf /usr/bin/${PYTHON_DAEMON_NAME}
	
	@# TODO systemd control for both three process for docker 
	@# REVIEW are both killall and systemctl commands running?
	@pgrep -x ${DAEMON_NAME} && (sudo killall "${DAEMON_NAME}" || sudo systemctl disable "${DAEMON_NAME}" || sudo systemctl stop "${DAEMON_NAME}") || echo "no ${DAEMON_NAME} found"
	@pgrep -x ${OLD_DAEMON_NAME} && (sudo killall "${OLD_DAEMON_NAME}" || sudo systemctl disable "${OLD_DAEMON_NAME}" || sudo systemctl stop "${OLD_DAEMON_NAME}") || echo "no ${OLD_DAEMON_NAME} found"

	@# TODO
	@# echo -e "Do you want to remove python implementation of ahenk if installed in system?"
	# read -rp "(Y/N) " input
	
	@pgrep -x ${PYTHON_DAEMON_NAME} && (sudo killall "${PYTHON_DAEMON_NAME}" || sudo systemctl disable "${PYTHON_DAEMON_NAME}" || sudo systemctl stop "${PYTHON_DAEMON_NAME}") || echo "no ${PYTHON_DAEMON_NAME} found"
	
	sudo systemctl daemon-reload	
	sudo rm -rf ${CONF_DIR}
	@echo -e "PRE-CLENING DONE\n"
test:

install:
	go build -o ${DESTDIR}/usr/bin/${REPO_NAME} ./cmd/ahenk-go/
	@sudo mkdir -p "${CONF_DIR}"

uninstall:


clean: #postclean:

