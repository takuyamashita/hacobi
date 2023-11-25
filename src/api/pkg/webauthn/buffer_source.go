package webauthn

import "encoding/base64"

type BufferSource []byte

func (c BufferSource) URLSafeBase64() string {

	return base64.RawURLEncoding.EncodeToString(c)
}

func (c BufferSource) MarshalJSON() ([]byte, error) {

	return []byte(`"` + c.URLSafeBase64() + `"`), nil
}
