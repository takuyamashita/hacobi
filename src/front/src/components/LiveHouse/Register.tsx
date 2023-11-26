"use client";

import React from "react";
import Button from "@/components/Button";

type Props = {
  token: string;
};

const Register = ({ token }: Props) => {
  const handleSubmit = async (
    e: React.FormEvent<HTMLFormElement> | React.MouseEvent,
  ) => {
    e.preventDefault();

    const base64URLSafeToUint8Array = (base64URLSafe: string) => {

      if (base64URLSafe === undefined || base64URLSafe === null) {
        return new Uint8Array();
      }

      // padding
      const pad = (s: string) => {
        while (s.length % 4 !== 0) {
          s += '=';
        }
        return s;
      };

      // base64URLSafe to base64
      const base64 = pad(base64URLSafe)
        .replace(/\-/g, '+')
        .replace(/_/g, '/');

      // base64 to Uint8Array
      const raw = window.atob(base64);
      const rawLength = raw.length;
      const array = new Uint8Array(new ArrayBuffer(rawLength));
      for (let i = 0; i < rawLength; i++) {
        array[i] = raw.charCodeAt(i);
      }
      return array;
    }

    const res = await fetch("/api/v1/auth", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const data = await res.json();

    if (!window.PublicKeyCredential) return;

    const a = base64URLSafeToUint8Array(data.Challenge);

    const excludeCredentials = data.ExcludeCredentials.map( (credential: PublicKeyCredentialDescriptor) => {
      return {
        id: base64URLSafeToUint8Array(credential.id),
        type: credential.type,
        transports: credential.transports,
      };
    });

    const rp = 'id' in data.Rp ?  {
      ...data.Rp,
      id: base64URLSafeToUint8Array(data.Rp.id),
    }: data.Rp;

    const user = {
      ...data.User,
      id: base64URLSafeToUint8Array(data.User.id),
    };

    const pubKeyOptions: PublicKeyCredentialCreationOptions = {
      challenge: a,
      rp: rp,
      user: user,
      authenticatorSelection: data.AuthenticatorSelection,
      pubKeyCredParams: data.PubKeyCredParams,
      attestation: data.Attestation,
      timeout: data.Timeout,
      excludeCredentials: excludeCredentials,
      extensions: undefined,
    };


    const publickeyCredential = await navigator.credentials.create({
      publicKey: pubKeyOptions,
    }) as PublicKeyCredential;

    console.log(publickeyCredential);
    console.log(publickeyCredential.response.clientDataJSON);
    // decode clientDataJSON(Array Buffer) to JSON
    const clientDataJSON = JSON.parse(
      new TextDecoder().decode(publickeyCredential.response.clientDataJSON),
    );
    console.log(clientDataJSON);
    console.log(clientDataJSON.challenge);
    console.log(publickeyCredential.response);

    const r = publickeyCredential.response as AuthenticatorAttestationResponse

    const pubKey = r.getPublicKey();
    if (!pubKey) return;

    console.log("attestationObject", new TextDecoder().decode(r.attestationObject));
    console.log("clientDataJSON", new TextDecoder().decode(r.clientDataJSON));
    console.log("authenticatorData", new TextDecoder().decode(r.getAuthenticatorData()));
    console.log("transports", r.getTransports());
    console.log("publicKeyAlgorithm", r.getPublicKeyAlgorithm());
    console.log("signature", new TextDecoder().decode(pubKey));

    fetch("/api/v1/auth", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id : publickeyCredential.id,
        rawId: btoa(String.fromCharCode(...(new Uint8Array(publickeyCredential.rawId)))),
        type: publickeyCredential.type,
        authenticatorAttachment: publickeyCredential.authenticatorAttachment,
        clientExtensionResults: publickeyCredential.getClientExtensionResults(),
        response: {
          attestationObject: btoa(String.fromCharCode(...(new Uint8Array(r.attestationObject)))),
          clientDataJSON: btoa(String.fromCharCode(...(new Uint8Array(r.clientDataJSON)))),
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
