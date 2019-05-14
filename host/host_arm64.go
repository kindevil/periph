// Copyright 2016 The Periph Authors. All rights reserved.
// Use of this source code is governed under the Apache License, Version 2.0
// that can be found in the LICENSE file.

package host

import (
	// Make sure CPU and board drivers are registered.
	_ "github.com/kindevil/periph/host/allwinner"
	_ "github.com/kindevil/periph/host/bcm283x"
	_ "github.com/kindevil/periph/host/pine64"
	_ "github.com/kindevil/periph/host/rpi"
)
