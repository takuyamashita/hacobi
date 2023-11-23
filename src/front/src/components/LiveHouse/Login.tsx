"use client";

import React from "react";
import Button from "@/components/Button";

const Login = () => {
  const [email, setEmail] = React.useState("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("submit", email);
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
