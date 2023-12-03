import { arrayBufferToBase64URLSafe } from "@/utils/base64";

export const convertPublicKeyCredential = (credential: PublicKeyCredential) => {
  const r = credential.response as
    | AuthenticatorAttestationResponse
    | AuthenticatorAssertionResponse;

  if (r instanceof AuthenticatorAssertionResponse) {
    return {
      id: credential.id,
      rawId: arrayBufferToBase64URLSafe(credential.rawId),
      type: credential.type,
      authenticatorAttachment: credential.authenticatorAttachment,
      clientExtensionResults: credential.getClientExtensionResults(),
      response: {
        authenticatorData: arrayBufferToBase64URLSafe(r.authenticatorData),
        clientDataJSON: arrayBufferToBase64URLSafe(r.clientDataJSON),
        signature: arrayBufferToBase64URLSafe(r.signature),
        userHandle: r.userHandle
          ? arrayBufferToBase64URLSafe(r.userHandle)
          : undefined,
      },
    };
  }

  return {
    id: credential.id,
    rawId: arrayBufferToBase64URLSafe(credential.rawId),
    type: credential.type,
    authenticatorAttachment: credential.authenticatorAttachment,
    clientExtensionResults: credential.getClientExtensionResults(),
    response: {
      attestationObject: arrayBufferToBase64URLSafe(r.attestationObject),
      clientDataJSON: arrayBufferToBase64URLSafe(r.clientDataJSON),
      transports: r.getTransports(),
      publicKeyAlgorithm: r.getPublicKeyAlgorithm(),
    },
  };
};
