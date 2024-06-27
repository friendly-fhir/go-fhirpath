#!/usr/bin/env bash

set -euo pipefail

repo_root=$(git rev-parse --show-toplevel)
readonly repo_root

function generate() {
  cd "${repo_root}/grammar"
  java                                                                         \
    -jar "${jar_path}"                                                         \
    -Dlanguage=Go -o ../internal/parser                                        \
    -package parser                                                            \
    -visitor                                                                   \
    -no-listener                                                               \
    $@
}

tools_dir="${repo_root}/tools"
jar_file=antlr-4.13.1-complete.jar
jar_path="${tools_dir}/${jar_file}"
readonly tools_dir jar_file jar_path

if [[ ! -f "${jar_path}" ]]; then
  curl https://www.antlr.org/download/antlr-4.13.1-complete.jar > "${jar_path}"
fi

generate "fhirpath.g4"
