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

package convert

import (
	"math/big"
	"testing"

	"github.com/onflow/cadence"

	"github.com/stretchr/testify/assert"
)

func TestParseCadenceArgument(t *testing.T) {

	vectors := []struct {
		name     string
		param    string
		wantArg  cadence.Value
		checkErr assert.ErrorAssertionFunc
	}{
		{
			name:     "parse valid normal integer",
			param:    "Int16(1337)",
			wantArg:  cadence.Int16(1337),
			checkErr: assert.NoError,
		},
		{
			name:     "parse invalid normal integer",
			param:    "Int16(a337)",
			checkErr: assert.Error,
		},
		{
			name:     "parse invalid unsigned integer",
			param:    "UInt64(-1337)",
			checkErr: assert.Error,
		},
		{
			name:     "parse valid big integer",
			param:    "UInt256(1337)",
			wantArg:  cadence.UInt256{Value: big.NewInt(1337)},
			checkErr: assert.NoError,
		},
		{
			name:     "parse invalid big integer",
			param:    "Uint128(a337)",
			checkErr: assert.Error,
		},
		{
			name:     "parse valid fixed point",
			param:    "UFix64(13.37)",
			wantArg:  cadence.UFix64(1337000000),
			checkErr: assert.NoError,
		},
		{
			name:     "parse invalid fixed point",
			param:    "Fix64(13,37)",
			checkErr: assert.Error,
		},
		{
			name:     "parse valid address",
			param:    "Address(43AC64656E636521)",
			wantArg:  cadence.Address{0x43, 0xac, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x21},
			checkErr: assert.NoError,
		},
		{
			name:     "parse invalid address",
			param:    "Address(X3AC64656E636521)",
			checkErr: assert.Error,
		},
		{
			name:     "parse valid bytes",
			param:    "Bytes(43AC64656E636521)",
			wantArg:  cadence.Bytes{0x43, 0xac, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x21},
			checkErr: assert.NoError,
		},
		{
			name:     "parse invalid bytes",
			param:    "Bytes(X3AC64656E636521)",
			checkErr: assert.Error,
		},
		{
			name:     "parse valid string",
			param:    "String(MN7wrJh359Kx+J*#)",
			wantArg:  cadence.String("MN7wrJh359Kx+J*#"),
			checkErr: assert.NoError,
		},
	}

	for _, vector := range vectors {

		gotArg, err := ParseCadenceArgument(vector.param)
		vector.checkErr(t, err)

		if err == nil {
			assert.Equal(t, vector.wantArg, gotArg)
		}
	}
}