"use client";

import React from "react";

const Login = () => {

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    console.log("submit");
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
        />
      </div>
      <div className="flex justify-center">
        <button
          type="submit"
          className="bg-blue-500  text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        >
          新規登録
        </button>
      </div>
    </form>
  );
};
export default Login;
