"use client";

import React from "react";
import Button from "@/components/Button";
import {
  base64URLSafeToUint8Array,
  arrayBufferToBase64URLSafe,
} from "@/utils/base64";

type Props = {
  token: string;
};

const stringToUint8Array = (str: string) => {
  const array = new Uint8Array(str.length);
  for (let i = 0; i < str.length; i++) {
    array[i] = str.charCodeAt(i);
  }
  return array;
};

const Register = ({ token }: Props) => {
  const handleSubmit = async (
    e: React.FormEvent<HTMLFormElement> | React.MouseEvent,
  ) => {
    e.preventDefault();

    const res = await fetch(
      "/api/v1/live_house_account/credential/start_register",
      {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({
          token: token,
        }),
      },
    );

    const data = await res.json();

    if (!window.PublicKeyCredential) return;

    // decode challenge(Array Buffer) to Uint8Array
    const a = base64URLSafeToUint8Array(data.challenge);

    // decode excludeCredentials(Array Buffer) to Uint8Array
    const excludeCredentials = data.excludeCredentials
      ? data.ExcludeCredentials.map(
          (credential: { id: string; type: string; transports: string[] }) => {
            return {
              id: base64URLSafeToUint8Array(credential.id),
              type: credential.type,
              transports: credential.transports,
            };
          },
        )
      : [];

    const rp = {
      name: data.rp.name,
    };

    const user: PublicKeyCredentialUserEntity = {
      name: data.user.displayName,
      displayName: data.user.displayName,
      id: base64URLSafeToUint8Array(data.user.id),
    };

    const pubKeyOptions: PublicKeyCredentialCreationOptions = {
      challenge: a,
      rp: rp,
      user: user,
      authenticatorSelection: data.authenticatorSelection,
      pubKeyCredParams: data.pubKeyCredParams,
      attestation: data.attestation,
      timeout: data.timeout,
      excludeCredentials: excludeCredentials,
      extensions: undefined,
    };

    const publickeyCredential = (await navigator.credentials.create({
      publicKey: pubKeyOptions,
    })) as PublicKeyCredential;

    const r = publickeyCredential.response as AuthenticatorAttestationResponse;

    const pubKey = r.getPublicKey();
    if (!pubKey) return;

    fetch("/api/v1/live_house_account/credential/finish_register", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: publickeyCredential.id,
        rawId: arrayBufferToBase64URLSafe(publickeyCredential.rawId),
        type: publickeyCredential.type,
        authenticatorAttachment: publickeyCredential.authenticatorAttachment,
        clientExtensionResults: publickeyCredential.getClientExtensionResults(),
        response: {
          attestationObject: arrayBufferToBase64URLSafe(r.attestationObject),
          clientDataJSON: arrayBufferToBase64URLSafe(r.clientDataJSON),
          transports: r.getTransports(),
          publicKeyAlgorithm: r.getPublicKeyAlgorithm(),
        },
      }),
    });
  };

  return (
    <Button
      type="button"
      onClick={(e) => handleSubmit(e)}
    >
      登録を開始
    </Button>
  );
};
export default Register;
