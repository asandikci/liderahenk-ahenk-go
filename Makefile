REPO_NAME=ahenk-go
REPO_LINK=https://git.aliberksandikci.com.tr/liderahenk/${REPO_NAME}

DAEMON_NAME=ahenk-go
OLD_DAEMON_NAME=ahenkd-go
PYTHON_DAEMON_NAME=ahenkd

DATA_DIR=/etc/ahenk-go/
LIB_DIR=/usr/share/ahenk-go/
PLUGIN_DIR=${LIB_DIR}/plugins/
TEMP_DIR=$(mktemp -d)
MAIN_DIR=${TEMP_DIR}/${REPO_NAME}/

info:
	@echo "Made by Aliberk Sandıkçı - 2023"
	@echo "preclean: for cleaning old files, configurations"
	@echo "TODO test: Test go files"
	@echo "install: Build and install ahenk-go to DESTDIR"
	@echo "uninstall: Uninstall ahenk-go from DESTDIR"

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
	@# read -rp "(Y/N) " input
	
	@pgrep -x ${PYTHON_DAEMON_NAME} && (sudo killall "${PYTHON_DAEMON_NAME}" || sudo systemctl disable "${PYTHON_DAEMON_NAME}" || sudo systemctl stop "${PYTHON_DAEMON_NAME}") || echo "no ${PYTHON_DAEMON_NAME} found"
	
	sudo systemctl daemon-reload	
	sudo rm -rf ${DATA_DIR}
	@echo -e "PRE-CLENING DONE\n"
test:
	@echo -e "Testing go files not implemented yet!"
install:
	sudo go build -o ${DESTDIR}/usr/bin/${REPO_NAME} ./cmd/ahenk-go/
	@sudo mkdir -p "${DESTDIR}/${LIB_DIR}"
	@sudo mkdir -p "${DESTDIR}/${PLUGIN_DIR}"

	sudo go build -buildmode=plugin -o ${DESTDIR}/${PLUGIN_DIR}/resources.so ./plugins/resources
	@sudo mkdir -p "${DESTDIR}/${DATA_DIR}"

uninstall:
	@sudo rm -rf ${DESTDIR}/usr/bin/${REPO_NAME}
