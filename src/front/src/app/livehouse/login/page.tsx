import React from "react";
import { Metadata } from "next";
import Login from "@/components/LiveHouse/Login";

export const metadata: Metadata = {
  title: "Login",
};

const Page = () => {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Login />
    </main>
  );
};
export default Page;
