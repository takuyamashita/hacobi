import React, { ButtonHTMLAttributes } from "react";

type Props = {
  type?: ButtonHTMLAttributes<HTMLButtonElement>["type"];
  onClick?: () => void;
  children?: React.ReactNode;
};

const Button = ({ type, onClick, children }: Props) => {
  return (
    <div className="flex justify-center">
      <button
        type={type}
        onClick={onClick}
        className="bg-blue-500  text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
      >
        {children}
      </button>
    </div>
  );
};
export default Button;