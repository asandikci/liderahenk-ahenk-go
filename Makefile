SHELL := /bin/bash
STARSHIP_SHELL := /bin/bash
REPO_NAME=ahenk-go
REPO_LINK=https://git.aliberksandikci.com.tr/liderahenk/${REPO_NAME}

# These variables imported from pkg/confdir/confdir_linux.go
PATH_PROGRAM=/usr/bin/${REPO_NAME}
PATH_USER_CONFIG=$([ -n "${XDG_CONFIG_HOME}" ] && echo "${XDG_CONFIG_HOME}" || echo "${HOME}/.config")	# FIXME not working ?
PATH_DATA=/etc/${REPO_NAME}/
PATH_PLUGINS={PATH_DATA}plugins/
PATH_LOGS=${PATH_DATA}logs/
PATH_PID=/run/ahenk-go.pid
PATH_VERSION=${PATH_DATA}version
GO_ARCH=amd64

# Wine variables for users. STATIC! change if these paths not fits for you
WINE_USER=${USER}
WINE_C=~/.wine/drive_c/
WINE_PATH_PROGRAM=${WINE_C}Program\ Files/${REPO_NAME}/
WINE_PATH_EXE=${WINE_PATH_PROGRAM}${REPO_NAME}.exe
WINE_PATH_USER_CONFIG=${WINE_C}/users/${WINE_USER}/AppData/Local/
WINE_PATH_DATA=${WINE_C}ProgramData/${REPO_NAME}/
WINE_PATH_PLUGINS=${WINE_PATH_DATA}plugins/
WINE_PATH_LOGS=${WINE_PATH_DATA}logs/
WINE_PATH_VERSION=${PATH_DATA}version
WINE_GO_ARCH=amd64



info:
	@echo "Made by Aliberk Sandıkçı - 2023"
	@echo "-- GENERAL USAGE ---"
	@echo "install: install binary and copy plugins,packages,files"
	@echo "uninstall: remove files, program and daemon"

	@echo "-- LOCAL DEVELOPMENT ----"
	@echo "test: Test go files ,not implemented yet"
	@echo "local: install + systemd daemon"
	@echo "local_wine: install on wine with build windows&amd64"
	@echo "local_safeplugins: install with imported safeplugins, do not start systemd daemon"
	@echo "local_wine_safeplugins: install on wine with build windows&amd64 and imported safeplugins"

# 
# --- GENERAL USAGE ---
# 

install:
	sudo go build -o ${DESTDIR}${PATH_PROGRAM} ./cmd/${REPO_NAME}/
	@sudo mkdir -p "${DESTDIR}${PATH_DATA}"

	@echo "copying necessary plugins and packages"
	@sleep 1
	sudo cp -r ./plugins "${DESTDIR}${PATH_DATA}"
	sudo cp -r ./pkg "${DESTDIR}${PATH_DATA}"	

	@echo "copying necessary files"
	@sleep 1
	sudo cp go.mod "${DESTDIR}${PATH_DATA}"
	sudo cp go.sum "${DESTDIR}${PATH_DATA}"
	sudo cp version "${DESTDIR}${PATH_DATA}"

uninstall:
	@echo "removing systemd daemon"
	sudo systemctl stop ${REPO_NAME}
	sudo systemctl disable ${REPO_NAME}
	sudo rm -f "${DESTDIR}/etc/systemd/system/${REPO_NAME}.service"
	sudo systemctl daemon-reload

	@echo "WARNING: uninstall option removes all ${DESTDIR}${PATH_DATA} !!!"
	@echo "waiting 10 seconds, abort if you don't want to remove this dir"
	@sleep 10
	sudo rm -rf "${DESTDIR}${PATH_DATA}"
	sudo rm -rf "${DESTDIR}${PATH_PROGRAM}"


# 
# --- LOCAL DEVELOPMENT ---
# 

test:
	@echo -e "Testing go files not implemented yet!"

local: i_local_configure_paths i_local_build i_local_system_reload

local_wine: i_local_wine_configure_paths i_local_wine_build i_local_wine_link

local_safeplugins: i_local_configure_paths i_local_build_safeplugins

local_wine_safeplugins: i_local_wine_configure_paths i_local_wine_build_safeplugins i_local_wine_link



#
# --- IMPORTED BY OTHER ---
#

# --> i_local
i_local_configure_paths:
	sudo rm -rf "${DESTDIR}${PATH_DATA}"
	sudo rm -rf ${DESTDIR}${PATH_PROGRAM}

	sudo mkdir -p "${DESTDIR}${PATH_DATA}"

	sudo cp -r ./plugins "${DESTDIR}${PATH_DATA}"
	sudo cp -r ./pkg "${DESTDIR}${PATH_DATA}"	
	sudo cp go.mod "${DESTDIR}${PATH_DATA}"
	sudo cp go.sum "${DESTDIR}${PATH_DATA}"
	sudo cp version "${DESTDIR}${PATH_DATA}"

i_local_build:
	sudo env GOARCH=${GO_ARCH} go build  -o ${DESTDIR}${PATH_PROGRAM} ./cmd/${REPO_NAME}/

i_local_build_safeplugins:
	sudo env GOARCH=${GO_ARCH} go build -tags safeplugins -o ${DESTDIR}${PATH_PROGRAM} ./cmd/${REPO_NAME}/

i_local_system_reload:
	@echo "configuring systemd daemon"
	sudo cp ./debian/${REPO_NAME}.service /etc/systemd/system
	sudo systemctl daemon-reload
	sudo systemctl enable ${REPO_NAME}
	sudo systemctl start ${REPO_NAME}

# --> i_local_wine
i_local_wine_configure_paths:
	sudo rm -rf ${DESTDIR}${WINE_PATH_DATA}
	sudo rm -rf ${DESTDIR}${WINE_PATH_PROGRAM}
	sudo rm -rf ${DESTDIR}${PATH_PROGRAM}

	mkdir -p ${DESTDIR}${WINE_PATH_DATA}
	mkdir -p ${DESTDIR}${WINE_PATH_PROGRAM}

	cp -r ./plugins ${DESTDIR}${WINE_PATH_DATA}
	cp -r ./pkg ${DESTDIR}${WINE_PATH_DATA}	
	cp go.mod ${DESTDIR}${WINE_PATH_DATA}
	cp go.sum ${DESTDIR}${WINE_PATH_DATA}
	cp version ${DESTDIR}${WINE_PATH_DATA}

i_local_wine_build:
	sudo env GOOS=windows GOARCH=${WINE_GO_ARCH} go build -o ${DESTDIR}${WINE_PATH_EXE} ./cmd/${REPO_NAME}/

i_local_wine_build_safeplugins:
	sudo env GOOS=windows GOARCH=${WINE_GO_ARCH} go build -tags safeplugins -o ${DESTDIR}${WINE_PATH_EXE} ./cmd/${REPO_NAME}/

i_local_wine_link:
	sudo ln -svf ${DESTDIR}${WINE_PATH_EXE} ${DESTDIR}${PATH_PROGRAM}