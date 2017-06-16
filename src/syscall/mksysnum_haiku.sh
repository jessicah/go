#!/bin/sh
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

COMMAND="mksysnum_haiku.sh $@"

cat <<EOF
// $COMMAND
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

package syscall

const(
EOF

sed "s/^SYSCALL[0-9]\\+(\\([a-zA-Z0-9_]*\\), \\([0-9]*\\))/\\1=\\2/" < $1

cat <<EOF
)
EOF
