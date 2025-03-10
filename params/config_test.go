// Copyright 2017 The go-simplechain Authors
// This file is part of the go-simplechain library.
//
// The go-simplechain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-simplechain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-simplechain library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"math/big"
	"reflect"
	"testing"
)

func TestCheckCompatible(t *testing.T) {
	type test struct {
		stored, new *ChainConfig
		head        uint64
		wantErr     *ConfigCompatError
	}
	tests := []test{
		{stored: AllScryptProtocolChanges, new: AllScryptProtocolChanges, head: 0, wantErr: nil},
		{stored: AllScryptProtocolChanges, new: AllScryptProtocolChanges, head: 100, wantErr: nil},
		{
			stored:  &ChainConfig{SingularityBlock: big.NewInt(10)},
			new:     &ChainConfig{SingularityBlock: big.NewInt(20)},
			head:    9,
			wantErr: nil,
		},
		{
			stored: AllScryptProtocolChanges,
			new:    &ChainConfig{SingularityBlock: nil},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "SingularityBlock fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    nil,
				RewindTo:     0,
			},
		},
		{
			stored: AllScryptProtocolChanges,
			new:    &ChainConfig{SingularityBlock: big.NewInt(1)},
			head:   3,
			wantErr: &ConfigCompatError{
				What:         "SingularityBlock fork block",
				StoredConfig: big.NewInt(0),
				NewConfig:    big.NewInt(1),
				RewindTo:     0,
			},
		},
	}

	for _, test := range tests {
		err := test.stored.CheckCompatible(test.new, test.head)
		if !reflect.DeepEqual(err, test.wantErr) {
			t.Errorf("error mismatch:\nstored: %v\nnew: %v\nhead: %v\nerr: %v\nwant: %v", test.stored, test.new, test.head, err, test.wantErr)
		}
	}
}
