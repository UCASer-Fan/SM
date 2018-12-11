/*
Copyright Suzhou Tongji Fintech Research Institute 2017 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sm2

import (
	"crypto/rand"
	"fmt"
	"log"
	"testing"
)


/*
func testKeyGeneration(t *testing.T, c elliptic.Curve, tag string) {
	priv, err := GenerateKey()
	if err != nil {
		t.Errorf("%s: error: %s", tag, err)
		return
	}
	if !c.IsOnCurve(priv.PublicKey.X, priv.PublicKey.Y) {
		t.Errorf("%s: public key invalid: %s", tag, err)
	}
}

func TestKeyGeneration(t *testing.T) {
	testKeyGeneration(t, elliptic.P224(), "p224")
	if testing.Short() {
		return
	}
	testKeyGeneration(t, elliptic.P256(), "p256")
	testKeyGeneration(t, elliptic.P384(), "p384")
	testKeyGeneration(t, elliptic.P521(), "p521")
}
*/

func TestSignandVerify(t *testing.T) {

	priv, err := GenerateKey() // 生成密钥对
	if err != nil {
		log.Fatal(err)
	}
	msg := []byte("test")

	sign, err := priv.Sign(rand.Reader, msg, nil) // 签名

	if err != nil {
		log.Fatal(err)
	}
	ok := priv.Verify(msg, sign) // 密钥验证
	if ok != true {
		fmt.Printf("Verify error\n")
	} else {
		fmt.Printf("Verify ok\n")
	}

}




