import React from "react";
import { Metadata } from "next";
import RegisterEmail from "@/components/LiveHouse/RegisterEmail";

export const metadata: Metadata = {
  title: "Staff Register",
};

const Page = () => {
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <RegisterEmail />
    </main>
  );
};
export default Page;
