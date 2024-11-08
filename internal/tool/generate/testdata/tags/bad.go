// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

//go:build !good

package tags

import (
	"context"
	"fmt"

	"github.com/ServiceWeaver/weaver"
)

type BadService interface {
	DoSomething(context.Context) error
}

type badServiceImpl struct {
	weaver.Implements[BadService]
}

func (b *badServiceImpl) DoSomething(context.Context) error {
	Some code that does not compile
}