# Splice Signer

`splicesign` is a Golang `crypto.Signer` that simply returns a constant signature and public key.  It is useful if:

* You already have a pre-existing signature for a given message and key.
* You do not have the private key.
* You are working with an API that expects a private key in the form of a `crypto.Signer`, e.g. `CreateCertificate` from `crypto/x509`.

## Credits

Thanks to Filippo Valsorda for suggesting this approach.

## Licence

Copyright (C) 2020 Namecoin Developers.

splicesign is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

splicesign is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with splicesign.  If not, see [https://www.gnu.org/licenses/](https://www.gnu.org/licenses/).
