export const toJSON = (credential: PublicKeyCredential) => {
  const r = credential.response as
    | AuthenticatorAttestationResponse
    | AuthenticatorAssertionResponse;

  if (r instanceof AuthenticatorAssertionResponse) {
    /*


type CredentialAssertionResponse struct {
	PublicKeyCredential
	AssertionResponse AuthenticatorAssertionResponse `json:"response"`
}


type PublicKeyCredential struct {
	Credential
	RawID                   URLEncodedBase64                      `json:"rawId"`
	ClientExtensionResults  AuthenticationExtensionsClientOutputs `json:"clientExtensionResults,omitempty"`
	AuthenticatorAttachment string                                `json:"authenticatorAttachment,omitempty"`
}


type Credential struct {
	// ID is The credential’s identifier. The requirements for the
	// identifier are distinct for each type of credential. It might
	// represent a username for username/password tuples, for example.
	ID string `json:"id"`
	// Type is the value of the object’s interface object's [[type]] slot,
	// which specifies the credential type represented by this object.
	// This should be type "public-key" for Webauthn credentials.
	Type string `json:"type"`
}

type AuthenticatorResponse struct {
	// From the spec https://www.w3.org/TR/webauthn/#dom-authenticatorresponse-clientdatajson
	// This attribute contains a JSON serialization of the client data passed to the authenticator
	// by the client in its call to either create() or get().
	ClientDataJSON URLEncodedBase64 `json:"clientDataJSON"`
}
    */
    return {
      id: credential.id,
      rawId: btoa(String.fromCharCode(...new Uint8Array(credential.rawId)))
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=/g, ""),
      type: credential.type,
      authenticatorAttachment: credential.authenticatorAttachment,
      clientExtensionResults: credential.getClientExtensionResults(),
      response: {
        authenticatorData: btoa(
          String.fromCharCode(...new Uint8Array(r.authenticatorData)),
        )
          .replace(/\+/g, "-")
          .replace(/\//g, "_")
          .replace(/=/g, ""),
        clientDataJSON: btoa(
          String.fromCharCode(...new Uint8Array(r.clientDataJSON)),
        )
          .replace(/\+/g, "-")
          .replace(/\//g, "_")
          .replace(/=/g, ""),
        signature: btoa(String.fromCharCode(...new Uint8Array(r.signature)))
          .replace(/\+/g, "-")
          .replace(/\//g, "_")
          .replace(/=/g, ""),
        userHandle: r.userHandle
          ? btoa(String.fromCharCode(...new Uint8Array(r.userHandle)))
              .replace(/\+/g, "-")
              .replace(/\//g, "_")
              .replace(/=/g, "")
          : undefined,
      },
    };
  }

  return {
    id: credential.id,
    rawId: btoa(String.fromCharCode(...new Uint8Array(credential.rawId)))
      .replace(/\+/g, "-")
      .replace(/\//g, "_")
      .replace(/=/g, ""),
    type: credential.type,
    authenticatorAttachment: credential.authenticatorAttachment,
    clientExtensionResults: credential.getClientExtensionResults(),
    response: {
      attestationObject: btoa(
        String.fromCharCode(...new Uint8Array(r.attestationObject)),
      )
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=/g, ""),
      clientDataJSON: btoa(
        String.fromCharCode(...new Uint8Array(r.clientDataJSON)),
      )
        .replace(/\+/g, "-")
        .replace(/\//g, "_")
        .replace(/=/g, ""),
      transports: r.getTransports(),
      publicKeyAlgorithm: r.getPublicKeyAlgorithm(),
    },
  };
};
