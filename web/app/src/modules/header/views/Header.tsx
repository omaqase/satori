"use client";

import { useState } from "react";
import { Menu, X } from "lucide-react";
import logo from "../../../assets/logo.svg";

export const Header = () => {
  const [isOpen, setIsOpen] = useState(false);

  const menuItems = [
    { title: "Курсы", href: "#" },
    { title: "О проекте", href: "#" },
    { title: "Контакты", href: "#" },
  ];

  return (
    <header className="w-11/12 lg:w-10/12">
      <div className="relative">
        <div className="absolute inset-0 bg-white/10 backdrop-blur-md rounded-full" />

        <div className="relative px-4 py-2 flex items-center justify-between">
          <a href="#" className="relative z-10 group" aria-label="На главную">
            <div className="absolute inset-0 bg-black/5 rounded-full blur-md transition-all duration-300 group-hover:blur-lg opacity-0 group-hover:opacity-100" />
            <img
              src={logo || "/placeholder.svg"}
              alt="Логотип"
              className="w-12 h-12 relative z-10 transition-transform duration-300"
            />
          </a>

          <nav className="hidden md:block">
            <ul className="flex gap-12 font-medium">
              {menuItems.map((item) => (
                <li key={item.title}>
                  <a
                    href={item.href}
                    className="relative py-2 transition-colors duration-200 hover:text-white/80 text-white"
                  >
                    <span>{item.title}</span>
                    <span className="absolute inset-x-0 -bottom-1 h-0.5 bg-black/80 scale-x-0 transition-transform duration-200 origin-left hover:scale-x-100" />
                  </a>
                </li>
              ))}
            </ul>
          </nav>

          <div className="hidden md:block">
            <a className="relative inline-flex group" href="#">
              <div className="absolute inset-0 bg-black rounded-full blur-sm transition-all duration-300 group-hover:blur-md" />
              <span className="relative bg-black px-6 py-2 text-white rounded-full cursor-pointer transition-all duration-200 hover:bg-black/90 hover:scale-105 active:scale-95">
                Войти
              </span>
            </a>
          </div>

          <button
            onClick={() => setIsOpen(!isOpen)}
            className="relative z-10 p-2 md:hidden"
            aria-expanded={isOpen}
            aria-label={isOpen ? "Закрыть меню" : "Открыть меню"}
          >
            <div
              className="transition-transform duration-200"
              style={{ transform: isOpen ? "scale(0.8)" : "scale(1)" }}
            >
              {isOpen ? <X size={24} /> : <Menu size={24} />}
            </div>
          </button>
        </div>

        {isOpen && (
          <div className="absolute top-full left-0 right-0 mt-2 p-4 bg-white/10 backdrop-blur-md rounded-2xl md:hidden">
            <nav>
              <ul className="flex flex-col gap-4">
                {menuItems.map((item, index) => (
                  <li
                    key={item.title}
                    style={{
                      opacity: 0,
                      transform: "translateX(-20px)",
                      animation: `fadeInSlide 0.3s ease forwards ${index * 0.1}s`,
                    }}
                  >
                    <a
                      href={item.href}
                      className="block py-2 px-4 rounded-xl transition-colors duration-200 hover:bg-black/5"
                      onClick={() => setIsOpen(false)}
                    >
                      {item.title}
                    </a>
                  </li>
                ))}
                <li
                  style={{
                    opacity: 0,
                    transform: "translateX(-20px)",
                    animation: `fadeInSlide 0.3s ease forwards ${menuItems.length * 0.1}s`,
                  }}
                >
                  <a
                    href="#"
                    className="block py-2 px-4 text-center bg-black text-white rounded-xl transition-colors duration-200 hover:bg-black/90"
                    onClick={() => setIsOpen(false)}
                  >
                    Войти
                  </a>
                </li>
              </ul>
            </nav>
          </div>
        )}
      </div>
    </header>
  );
};
