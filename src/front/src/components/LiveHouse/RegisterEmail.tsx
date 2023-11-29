"use client";

import React from "react";
import Button from "../Button";

const RegisterEmail = () => {
  const [emailAddress, setEmailAddress] = React.useState<string>("");
  const [errorMessages, setErrorMessages] = React.useState<string[]>([]);

  const handleSubmit = async (
    e: React.FormEvent<HTMLFormElement> | React.MouseEvent,
  ) => {
    e.preventDefault();

    try {
      const res = await fetch(
        "/api/v1/send_live_house_staff_email_authorization",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            emailAddress,
          }),
        },
      );

      if (!res.ok) {
        throw res;
      }
    } catch (err) {
      if (err instanceof Response) {
        const json = await err.json();
        setErrorMessages((prev) => [...prev, json.message]);
      }
    }
  };

  return (
    <div>
      <p>RegisterEmail</p>
      <form onSubmit={(e) => handleSubmit(e)}>
        <label htmlFor="email">Email</label>
        {errorMessages.map((message) => (
          <p
            className="text-orange-500"
            key={message}
          >
            {message}
          </p>
        ))}
        <input
          type="email"
          id="email"
          value={emailAddress}
          onChange={(e) => setEmailAddress(e.target.value)}
        />
        <Button type="submit">Submit</Button>
      </form>
    </div>
  );
};

export default RegisterEmail;
