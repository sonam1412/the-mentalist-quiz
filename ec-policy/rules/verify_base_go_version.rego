#
# METADATA
# title: Verify Base Golang image version
# description: >-
#   This package is responsible for verifying the version of base Golang image to be 1.23.1
#
package verify_base_go_version

import rego.v1

import data.lib

# METADATA
# title: Verify Golang image version to match 1.23.1
# description: Confirm that the version of base go image used in the
#   Dockerfile of project is equal to 1.23.1
# custom:
#   short_name: verify_base_go_version
#   failure_msg: Golang base image's version isn't 1.23.1
#   solution: Update the version of Golang image in Dockerfile to 1.23.1
deny contains result if {
    required_version := "1.23.1"
    print(golang_version)
    not (required_version in golang_version)

    result := lib.result_helper(rego.metadata.chain(), [])
}

# Rule to extract the Go version from the environment variables
golang_version contains ver if {
    some env in input.image.config.Env
    print(env)
    startswith(env, "GOLANG_VERSION=")
    split(env, "=", parts)
    print(parts)
    ver = parts[1]
}
