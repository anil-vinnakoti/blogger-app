// src/components/Header.tsx
import React from "react";
import { useTheme } from "../hooks/useTheme";
import { Moon, Sun } from "lucide-react";
import LogoutButton from "./Logout";
import { useAuth } from "../context/AuthContext";

export default function Header() {
  const { theme, toggleTheme } = useTheme();
  const { isAuthenticated } = useAuth();

  return (
    <header className="flex items-center justify-between px-6 py-4 shadow-md">
      <h1 className="text-xl font-bold ml-4">Blogger App</h1>

      <div className="flex items-center space-x-4">
        {/* Navigation Buttons */}
        {!isAuthenticated ? (
          <>
            <button className="px-3 py-1 rounded-lg border-2 transition font-bold">
              <span className="align-[2px]">Sign Up</span>
            </button>
            <button className="px-3 py-1 rounded-lg border-2 align-top transition font-bold">
              <span className="align-[2px]">Login</span>
            </button>
          </>
        ) : (
          <LogoutButton />
        )}

        {/* Dark Mode Toggle */}
        <button
          onClick={toggleTheme}
          className="p-2 rounded-lg bg-gray-700 hover:bg-gray-600 transition"
        >
          {theme === "light" ? (
            <Moon size={20} className="text-yellow-300" />
          ) : (
            <Sun size={20} className="text-yellow-400" />
          )}
        </button>
      </div>
    </header>
  );
}
