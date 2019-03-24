#!/usr/bin/env bash
set -euo pipefail

package=format-time
src_path=$(pwd)/src/
bin_path=$(pwd)/bin/
platforms=("windows/amd64" "linux/amd64" "darwin/amd64")

function help() {
  echo -e "Usage: go <command>"
  echo -e
  echo -e "    help               Print this help"
  echo -e "    run                Run locally without building binary"
  echo -e "    build              Build binary locally"
  echo -e "    test               Run local unit tests"
  echo -e "    build-docker       Bake docker image and run smoke tests locally"
  echo -e 
  exit 0
}

function run() {

    _console_msg "Running locally ..."
    echo

    src_path=$(pwd)/src/

    pushd "${src_path}" >/dev/null

    go run main.go "${@:-}"

    popd > /dev/null 

    _console_msg "Build complete"

}

function build() {

    _console_msg "Building locally ..."

    _assert_variables_set package src_path bin_path platforms

    pushd "${src_path}" >/dev/null

    for platform in "${platforms[@]}"; do

        platform_split=(${platform//\// })
        GOOS=${platform_split[0]}
        GOARCH=${platform_split[1]}

        output_name=${package}'-'${GOOS}'-'${GOARCH}
        if [ $GOOS = "windows" ]; then
            output_name+='.exe'
        fi  

        env GOOS=${GOOS} GOARCH=${GOARCH} go build -o ${bin_path}/${output_name}
        if [[ $? -ne 0 ]]; then 
            _console_msg "Error from 'go build' - aborting" ERROR true 
            exit 1
        fi

        if [[ ! -f ${bin_path}/${output_name} ]]; then 
            _console_msg "Binary for ${platform} not present -> ${bin_path}/${output_name}" ERROR true 
            exit 1
        else
            _console_msg "Built: ${bin_path}/${output_name}"
        fi 

    done

    _console_msg "Build complete"

}

function build-docker() {

    _console_msg "Creating docker image ..."

    pushd $(pwd) >/dev/null

    docker build -t format-time .

    _console_msg "Build complete, running smokes ..."

    docker run format-time 61

    popd > /dev/null

    _console_msg "Docker build complete. To use this program locally:

          docker run format-time <time-in-seconds>"

}

function test() {

    _console_msg "Running unit tests ..."

    pushd "${src_path}" >/dev/null

    go test -v 

    _console_msg "Build complete"

}

function _assert_variables_set() {
  local error=0
  local varname
  for varname in "$@"; do
    if [[ -z "${!varname-}" ]]; then
      echo "${varname} must be set" >&2
      error=1
    fi
  done
  if [[ ${error} = 1 ]]; then
    exit 1
  fi
}

function _console_msg() {
  local msg=${1}
  local level=${2:-}
  local ts=${3:-}
  if [[ -z ${level} ]]; then level=INFO; fi
  if [[ -n ${ts} ]]; then ts=" [$(date +"%Y-%m-%d %H:%M")]"; fi

  echo ""
  if [[ ${level} == "ERROR" ]] || [[ ${level} == "CRIT" ]] || [[ ${level} == "FATAL" ]]; then
    (echo 2>&1)
    (echo >&2 "-> [${level}]${ts} ${msg}")
  else 
    (echo "-> [${level}]${ts} ${msg}")
  fi
  echo ""

}

function ctrl_c() {
    if [ ! -z ${PID:-} ]; then
        kill ${PID}
    fi
    exit 1
}

trap ctrl_c INT

if [[ ${1:-} =~ ^(help|run|build|test|build-docker)$ ]]; then
  COMMAND=${1}
  shift
  $COMMAND "$@"
else
  help
  exit 1
fi
