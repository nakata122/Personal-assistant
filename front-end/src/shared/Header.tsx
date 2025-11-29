import React, { useState, useEffect } from "react";
import Login from "./Login";

const Header = () => {
  const [isVisible, setIsVisible] = useState<boolean>(true);
  const [login, setLogin] = useState<boolean>(false);

  function toggleLogin() {
    setLogin(!login);
  };

  return (
    <>
    <header
      className={`bg-white shadow-lg fixed w-full top-0 left-0 z-10 transition-transform duration-700 transform ${isVisible ? "translate-y-0" : "-translate-y-full"}`}
    >
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="shrink-0">
            <h1 className="text-2xl font-bold text-gray-800">Personal Assistant</h1>
          </div>

          <nav className="hidden lg:flex space-x-6">
              <button onClick={toggleLogin} className="text-white bg-gray-900 p-6 rounded-2xl hover:text-gray-200 relative group py-2 text-2xl cursor-pointer">
                Login
              </button>
          </nav>

        </div>
      </div>
    </header>
    {login ? <Login toggle={toggleLogin}/> : null}
    </>
  );
};

export default Header;