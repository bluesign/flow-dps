// Copyright 2021 Alvalor S.A.
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

package rosetta

import (
	"github.com/awfm9/flow-dps/models/identifier"
)

// Transaction contains an array of operations that are attributable to the same
// transaction identifier.
//
// Examples of metadata given in the Rosetta API documentation are "size" and
// "lockTime".
type Transaction struct {
	ID         identifier.Transaction `json:"transaction_identifier"`
	Operations []Operation            `json:"operations"`
}
