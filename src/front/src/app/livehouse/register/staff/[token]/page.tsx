import React from "react";
import { Metadata } from "next";
import Register from "@/components/LiveHouse/Register";

export const metadata: Metadata = {
  title: "Staff Register",
};

type Props = {
  params: {
    token: string;
  };
};

const Page = ({params: {token}}: Props) => {

  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <Register token="token" />
    </main>
  );
};
export default Page;
