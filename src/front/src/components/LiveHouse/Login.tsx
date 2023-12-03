"use client";

import React from "react";
import Button from "@/components/Button";
import { base64URLSafeToUint8Array } from "@/utils/base64";
import { toJSON } from "@/utils/publicKeyCredential";

const Login = () => {
  const [email, setEmail] = React.useState("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("submit", email);

    try {
      const res = await fetch(
        "/api/v1/live_house_account/credential/start_login",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({ emailAddress: email }),
        },
      );

      if (!res.ok) {
        throw new Error("Network response was not ok");
      }

      const { challenge, rpId, timeout, userVerification } = await res.json();

      const credential = (await navigator.credentials.get({
        publicKey: {
          challenge: base64URLSafeToUint8Array(challenge),
          rpId,
          timeout,
          userVerification,
          allowCredentials: [],
          // allowCredentials: [
          //   {
          //     id: Uint8Array.from("id", (c) => c.charCodeAt(0)),
          //     type: "public-key",
          //     transports: ["usb", "nfc", "ble"],
          //   },
          // ],
          // extensions: {},
        },
      })) as null | PublicKeyCredential;

      if (!credential) {
        throw new Error("Credential is null");
      }

      /*

		Email                       string
		CredentialAssertionResponse protocol.CredentialAssertionResponse
      */

      const res2 = await fetch(
        "/api/v1/live_house_account/credential/finish_login",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            Email: email,
            CredentialAssertionResponse: toJSON(credential),
          }),
        },
      );
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <form onSubmit={(e) => handleSubmit(e)}>
      <div className="mb-4">
        <label
          className="block text-gray-700 text-sm font-bold mb-2"
          htmlFor="email"
        >
          メールアドレス
        </label>
        <input
          className="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline"
          id="email"
          type="email"
          placeholder="hacobi@example.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />
        <Button type="submit">新規登録</Button>
      </div>
    </form>
  );
};
export default Login;
