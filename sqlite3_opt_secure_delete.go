// Copyright (C) 2014 Yasuhiro Matsumoto <mattn.jp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.
// +build sqlite_secure_delete

package sqlite3

/*
#cgo CFLAGS: -DSQLITE_SECURE_DELETE
#cgo LDFLAGS: -lm
*/
import "C"
