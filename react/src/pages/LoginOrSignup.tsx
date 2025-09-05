import { useState } from "react";
import { useLogin } from "../api/auth";
import { useAuth } from "../context/AuthContext";
import { useNavigate, useLocation, Link } from "react-router-dom";

export default function LoginOrSignup() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const login = useLogin();
  const { setIsAuthenticated } = useAuth();
  const navigate = useNavigate();
  const location = useLocation();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    login.mutate(
      { username, password },
      {
        onSuccess: () => {
          setIsAuthenticated(true);
          navigate("/"); // redirect after login
        }
      }
    );
  };

  return (
    <div className="p-6 max-w-md mx-auto">
      <h2 className="text-xl font-semibold mb-4">
        {location.pathname.includes("/login") ? "Login" : "Sign Up"}
      </h2>
      <form onSubmit={handleSubmit} className="space-y-3">
        <input
          type="text"
          placeholder="Username"
          value={username}
          onChange={(e) => setUsername(e.target.value)}
          className="border p-2 rounded w-full"
          required
        />
        <input
          type="password"
          placeholder="Password"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="border p-2 rounded w-full"
          required
        />
        <button className="bg-blue-600 text-white px-4 py-2 rounded w-full">
          {location.pathname.includes("/login") ? "Login" : "Sign Up"}
        </button>
        <div className="text-center">
          Want to &nbsp;
          <Link
            to={!location.pathname.includes("/login") ? "/login" : "/signup"}
            className="text-blue-700 "
          >
            {!location.pathname.includes("/login") ? "Login" : "Sign Up"}
          </Link>
        </div>
      </form>
      {login.isError && (
        <p className="text-red-500 mt-2">Invalid credentials</p>
      )}
    </div>
  );
}
