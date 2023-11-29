"use client";

import React from "react";
import Button from "../Button";

const RegisterEmail = () => {
  const [emailAddress, setEmailAddress] = React.useState<string>("");

  const handleSubmit = async (
    e: React.FormEvent<HTMLFormElement> | React.MouseEvent,
  ) => {
    e.preventDefault();

    const res = await fetch("/api/v1/send_live_house_staff_email_authorization", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        emailAddress,
      }),
    });
  };

  return (
    <div>
      <p>RegisterEmail</p>
      <form onSubmit={(e) => handleSubmit(e)}>
        <label htmlFor="email">Email</label>
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
