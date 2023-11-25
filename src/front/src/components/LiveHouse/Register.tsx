"use client";

import React from "react";
import Button from "@/components/Button";

type Props = {
  token: string;
}

const Register = ({token}: Props) => {

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>|React.MouseEvent) => {
    e.preventDefault();
    console.log("submit", token);
  };

  return (

    <Button type="button" onClick={(e) => handleSubmit(e)}>登録を開始</Button>
  );
};
export default Register;