"use client";

import React from "react";
import Button from "@/components/Button";

type Props = {
  token: string;
}

const Register = ({token}: Props) => {

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("submit", token);
  };

  return (

    <Button type="submit">登録を開始</Button>
  );
};
export default Register;