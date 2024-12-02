// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rosas/onex.
//

package initializer

import (
	"k8s.io/apiserver/pkg/admission"

	clientset "github.com/rosas/onex/pkg/generated/clientset/versioned"
	"github.com/rosas/onex/pkg/generated/informers"
)

// WantsInternalInformerFactory defines a function which sets InformerFactory for admission plugins that need it.
type WantsInternalInformerFactory interface {
	admission.InitializationValidator
	SetInternalInformerFactory(informers.SharedInformerFactory)
}

// WantsInternalClientSet defines a function which sets external ClientSet for admission plugins that need it.
type WantsInternalClientSet interface {
	admission.InitializationValidator
	SetInternalClientSet(clientset.Interface)
}
