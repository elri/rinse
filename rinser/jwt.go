package rinser

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWTHeader struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
	Kid       string `json:"kid"`
}

type JWTPayload struct {
	Issuer string `json:"iss"`
	Type   string `json:"typ"`
	// TODO add on more here
}

// decodeJWTStringToBytes decodes a JWT specific base64url encoding,
// and returns the bytes represented by the base64 string
func decodeJWTStringToBytes(str string) (b []byte) {
	var err error
	b, err = jwt.NewParser().DecodeSegment(str)
	if err != nil {
		fmt.Printf("could not decode segment: %v", err)
	}
	return
}

func ExtractHeaderPayloadSignature(jwtToken string) (header, payload, signature string) {
	tokenSplit := strings.Split(jwtToken, ".")
	header = tokenSplit[0]
	payload = tokenSplit[1]
	signature = tokenSplit[2]
	return
}

func (rns *Rinse) VerifyJWT(jwtToken string) {
	keyMap := rns.PublicJWTKeys

	h64, p64, s64 := ExtractHeaderPayloadSignature(jwtToken)
	var header JWTHeader
	json.Unmarshal(decodeJWTStringToBytes(h64), &header)

	//TODO check payload, ie is the issuer an approved one

	key := keyMap[header.Kid]

	signed := fmt.Sprintf("%s.%s", h64, p64)
	sig := decodeJWTStringToBytes(s64)

	/*
		TODO
		get kid from keyMap
		get alg from header and use that to fetch type of method from jwt
	*/

	err := jwt.SigningMethodRS256.Verify(signed, sig, key)
	if err != nil {
		log.Fatal(err)
	}

}
