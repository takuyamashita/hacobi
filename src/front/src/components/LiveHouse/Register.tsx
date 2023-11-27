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

    const stringToUint8Array = (str: string) => {
      const array = new Uint8Array(str.length);
      for (let i = 0; i < str.length; i++) {
        array[i] = str.charCodeAt(i);
      }
      return array;
    };

    const base64URLSafeToUint8Array = (base64URLSafe: string) => {
      if (base64URLSafe === undefined || base64URLSafe === null) {
        return new Uint8Array();
      }

      // padding
      const pad = (s: string) => {
        while (s.length % 4 !== 0) {
          s += "=";
        }
        return s;
      };

      // base64URLSafe to base64
      const base64 = pad(base64URLSafe).replace(/\-/g, "+").replace(/_/g, "/");

      // base64 to Uint8Array
      const raw = window.atob(base64);
      const rawLength = raw.length;
      const array = new Uint8Array(new ArrayBuffer(rawLength));
      for (let i = 0; i < rawLength; i++) {
        array[i] = raw.charCodeAt(i);
      }
      return array;
    };

    const res = await fetch("/api/v1/ceremony/start", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });

    const data = await res.json();

    if (!window.PublicKeyCredential) return;

    // decode challenge(Array Buffer) to Uint8Array
    const a = base64URLSafeToUint8Array(data.challenge);

    // decode excludeCredentials(Array Buffer) to Uint8Array
    const excludeCredentials = data.excludeCredentials
      ? data.ExcludeCredentials.map(
          (credential: { id: string; type: string; transports: string[] }) => {
            return {
              id: stringToUint8Array(credential.id),
              type: credential.type,
              transports: credential.transports,
            };
          },
        )
      : [];

    const rp = {
      name: data.rp.name,
    };

    const user = {
      ...data.user,
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

    console.log(publickeyCredential);
    console.log(publickeyCredential.response.clientDataJSON);
    // decode clientDataJSON(Array Buffer) to JSON
    const clientDataJSON = JSON.parse(
      new TextDecoder().decode(publickeyCredential.response.clientDataJSON),
    );
    console.log(clientDataJSON);
    console.log(clientDataJSON.challenge);
    console.log(publickeyCredential.response);

    const r = publickeyCredential.response as AuthenticatorAttestationResponse;

    const pubKey = r.getPublicKey();
    if (!pubKey) return;

    console.log(
      "attestationObject",
      new TextDecoder().decode(r.attestationObject),
    );
    console.log("clientDataJSON", new TextDecoder().decode(r.clientDataJSON));
    console.log(
      "authenticatorData",
      new TextDecoder().decode(r.getAuthenticatorData()),
    );
    console.log("transports", r.getTransports());
    console.log("publicKeyAlgorithm", r.getPublicKeyAlgorithm());
    console.log("signature", new TextDecoder().decode(pubKey));

    fetch("/api/v1/auth", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        id: publickeyCredential.id,
        rawId: btoa(
          String.fromCharCode(...new Uint8Array(publickeyCredential.rawId)),
        )
          .replace(/\+/g, "-")
          .replace(/\//g, "_")
          .replace(/=/g, ""),
        type: publickeyCredential.type,
        authenticatorAttachment: publickeyCredential.authenticatorAttachment,
        clientExtensionResults: publickeyCredential.getClientExtensionResults(),
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
