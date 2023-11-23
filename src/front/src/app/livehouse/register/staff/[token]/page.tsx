import React from "react";
import { Metadata } from "next";

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
      {token}
    </main>
  );
};
export default Page;
