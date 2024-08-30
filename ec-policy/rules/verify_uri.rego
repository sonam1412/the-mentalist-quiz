#
# METADATA
# title: Verify URI
# description: >-
#   This package is responsible for verifying the content of attestations
#
package verify_uri

import rego.v1

import data.lib

# METADATA
# title: Verify URI count
# description: Confirm there are entries (other than 1) in the predicate.materials
#   array of the attestation that contains the specific uri as seen in the below query
# custom:
#   short_name: verify_uri
#   failure_msg: Unexpected count of URI prefix matches
#   solution: Fix it
deny contains result if {
	match := [material |
		some material in materials
		material.uri == "git+https://github.com/dheerajodha/the-mentalist-quiz.git"
	]
	print(match)
	count(match) == 1
	result := lib.result_helper(rego.metadata.chain(), [])
}

materials contains material if {
	some attestation in input.attestations
		some material in attestation.statement.predicate.materials
			material.uri
			material.digest.sha1
}
