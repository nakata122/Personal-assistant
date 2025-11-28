import React, { useState, useEffect } from "react";

const Header = () => {
  const [isVisible, setIsVisible] = useState(true);

//   useEffect(() => {
//     setIsVisible(true);
//   }, []);

  const navLinks = [
    { id: 1, name: "Home", href: "#" },
    { id: 2, name: "About", href: "#" },
    { id: 9, name: "Contact", href: "#" }
  ];

  return (
    <header
      className={`bg-white shadow-lg fixed w-full top-0 left-0 z-50 transition-transform duration-700 transform ${isVisible ? "translate-y-0" : "-translate-y-full"}`}
    >
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div className="flex justify-between items-center h-16">
          <div className="shrink-0">
            <h1 className="text-2xl font-bold text-gray-800">Personal Assistant</h1>
          </div>

          <nav className="hidden lg:flex space-x-6">
            {navLinks.map((link) => (
              <a
                key={link.id}
                href={link.href}
                className="text-gray-600 hover:text-gray-900 relative group py-2 text-sm"
              >
                {link.name}
                <span className="absolute bottom-0 left-0 w-full h-0.5 bg-blue-500 transform scale-x-0 transition-transform duration-300 group-hover:scale-x-100" />
              </a>
            ))}
          </nav>

        </div>
      </div>
    </header>
  );
};

export default Header;