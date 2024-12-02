// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package toyblc

import genericoptions "github.com/rosas/onex/pkg/options"

func scheme(opts *genericoptions.TLSOptions) string {
	scheme := "http"
	if opts != nil && opts.UseTLS {
		scheme = "https"
	}

	return scheme
}
