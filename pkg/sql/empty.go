// Copyright 2016 The Cockroach Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
// implied. See the License for the specific language governing
// permissions and limitations under the License.

package sql

import (
	"golang.org/x/net/context"

	"github.com/cockroachdb/cockroach/pkg/sql/parser"
)

// emptyNode is a planNode with no columns and either no rows (default) or a single row with empty
// results (if results is initialized to true). The former is used for nodes that have no results
// (e.g. a table for which the filtering condition has a contradiction), the latter is used by
// select statements that have no table or where we detect the filtering condition throws away all
// results.
type emptyNode struct {
	results bool
}

func (*emptyNode) Values() parser.Datums { return nil }
func (*emptyNode) Start(runParams) error { return nil }
func (*emptyNode) Close(context.Context) {}

func (e *emptyNode) Next(runParams) (bool, error) {
	r := e.results
	e.results = false
	return r, nil
}
