// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

package http

import (
	"net/http"

	"github.com/hashicorp/vault/vault"
)

// This endpoint is needed to answer queries before Vault unseals
// or becomes the leader.
func handleSysLeader(core *vault.Core, opt ...ListenerConfigOption) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			handleSysLeaderGet(core, w, opt...)
		default:
			respondError(w, http.StatusMethodNotAllowed, nil)
		}
	})
}

func handleSysLeaderGet(core *vault.Core, w http.ResponseWriter, opt ...ListenerConfigOption) {
	resp, err := core.GetLeaderStatus()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err)
		return
	}

	opts, err := getOpts(opt...)
	if opts.withRedactAddresses {
		resp.LeaderAddress = opts.withRedactionValue
		resp.LeaderClusterAddress = opts.withRedactionValue
	}

	respondOk(w, resp)
}
