/*
 * Bitbucket API
 *
 * Code against the Bitbucket API to automate simple tasks, embed Bitbucket data into your own site, build mobile or desktop apps, or even add custom UI add-ons into Bitbucket itself using the Connect framework.
 *
 * API version: 2.0
 * Contact: support@bitbucket.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package bitbucket

import (
	"time"
)

type SshKey struct {

	Type_ string `json:"type"`

	// The SSH key's immutable ID.
	Uuid string `json:"uuid,omitempty"`

	// The SSH public key value in OpenSSH format.
	Key string `json:"key,omitempty"`

	// The comment parsed from the SSH key (if present)
	Comment string `json:"comment,omitempty"`

	// The user-defined label for the SSH key
	Label string `json:"label,omitempty"`

	CreatedOn time.Time `json:"created_on,omitempty"`

	LastUsed time.Time `json:"last_used,omitempty"`

	Links *VersionLinks `json:"links,omitempty"`
}
