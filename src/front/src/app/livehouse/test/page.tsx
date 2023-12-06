"use client";

import React from "react";

const Page = () => {
  const buttonHandler = async (e: React.MouseEvent<HTMLButtonElement>) => {
    await fetch("/api/v1/live_house_staff_account/enable/live_house_staff", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
      },
    });
  };

  return (
    <div>
      <h1>Test Page</h1>
      <button onClick={buttonHandler}>test</button>
    </div>
  );
};

export default Page;
