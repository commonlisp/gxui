// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parts

import (
	"gaze/gxui"
)

type ParentableOuter interface{}

type Parentable struct {
	outer  ParentableOuter
	parent gxui.Container
}

func (p *Parentable) Init(outer ParentableOuter) {
	p.outer = outer
}

func (p *Parentable) Parent() gxui.Container {
	return p.parent
}

func (p *Parentable) SetParent(parent gxui.Container) {
	p.parent = parent
}