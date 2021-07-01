// Copyright 2021 Optakt Labs OÜ
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package loader

import (
	"fmt"
	"os"

	"github.com/onflow/flow-go/ledger/complete/mtrie/flattener"
	"github.com/onflow/flow-go/ledger/complete/mtrie/trie"
	"github.com/onflow/flow-go/ledger/complete/wal"
)

type Loader struct {
	file *os.File
}

func New(file *os.File) *Loader {
	return &Loader{file: file}
}

func (l *Loader) Checkpoint() (*trie.MTrie, error) {

	checkpoint, err := wal.ReadCheckpoint(l.file)
	if err != nil {
		return nil, fmt.Errorf("could not read checkpoint: %w", err)
	}

	trees, err := flattener.RebuildTries(checkpoint)
	if err != nil {
		return nil, fmt.Errorf("could not rebuild tries: %w", err)
	}

	if len(trees) != 1 {
		return nil, fmt.Errorf("should only have one trie in root checkpoint (tries: %d)", len(trees))
	}

	return trees[0], nil
}