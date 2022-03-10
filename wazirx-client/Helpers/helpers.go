package Helpers

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// ComputeHmac256Signature returns signature string
func ComputeHmac256Signature(message, secret string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(message))
	signingKey := fmt.Sprintf("%x", mac.Sum(nil))
	//fmt.Println(signingKey)
	return signingKey

}

// CurrentTimestamp returns formatted time as per wazrix's server time
func CurrentTimestamp() int64 {
	return FormatTimestamp(time.Now())
}

// FormatTimestamp formats a time into Unix timestamp in milliseconds
func FormatTimestamp(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// PrettyPrint to print struct in a readable way
func PrettyPrint(i interface{}) string {

	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
