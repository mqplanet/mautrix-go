// Copyright (c) 2020 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package patch

import (
	"fmt"
	"strings"
)

var ThirdPartyIdEncrypt bool
var AsBotName string
var AsUserPrefix string
var XorKey string

func Parse(stateKey string) string {
	if ThirdPartyIdEncrypt && len(stateKey)>0 && stateKey != AsBotName {
		if strings.Index(stateKey, AsUserPrefix) == 1 {
			userIdArr := strings.Split(stateKey, AsUserPrefix)
			userIdArrLast := strings.Split(userIdArr[1], ":")
			stateKey = fmt.Sprintf("@%s%s:%s", AsUserPrefix, Dec(userIdArrLast[0]), userIdArrLast[1])
		}
	}
	return stateKey
}

func ParseLocalPart(localpart string, encrypt bool) string {
	if ThirdPartyIdEncrypt && len(localpart)>0 && localpart != AsBotName {
		if strings.Index(localpart, AsUserPrefix) == 0 {
			localpartArr := strings.Split(localpart, AsUserPrefix)
			if encrypt {
				localpart = AsUserPrefix + Enc(localpartArr[1])
			} else {
				localpart = AsUserPrefix + Dec(localpartArr[1])
			}
		}
	}
	return localpart
}

