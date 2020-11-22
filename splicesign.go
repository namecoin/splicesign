// Copyright 2020 Jeremy Rand.

// This file is part of splicesign.
//
// splicesign is free software: you can redistribute it and/or
// modify it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// splicesign is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with splicesign.  If not, see
// <https://www.gnu.org/licenses/>.

package splicesign

import (
	"crypto"
	"io"
)

type SpliceSigner struct {
	PublicKey crypto.PublicKey
	Signature []byte
}

func (signer *SpliceSigner) Public() crypto.PublicKey {
	return signer.PublicKey
}

func (signer *SpliceSigner) Sign(rand io.Reader, digest []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	return signer.Signature, nil
}
