import { useState } from "react";
import { useLoginOrSignUp } from "../api/auth";
import { useAuth } from "../context/AuthContext";
import { useNavigate, useLocation, Link } from "react-router-dom";

export default function LoginOrSignup() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const location = useLocation();
  const loginOrSignup = useLoginOrSignUp(location.pathname);
  const { setIsAuthenticated } = useAuth();
  const navigate = useNavigate();

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    loginOrSignup.mutate(
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
    <div className="flex h-screen justify-center items-center">
      <div className="w-full max-w-md p-6 rounded-lg shadow-lg">
        <h2 className="text-2xl font-semibold mb-6 text-center">
          {location.pathname.includes("/login") ? "Login" : "Sign Up"}
        </h2>
        <form onSubmit={handleSubmit} className="space-y-4">
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
          <button className="bg-blue-600 px-4 py-2 rounded w-full hover:bg-blue-700">
            {location.pathname.includes("/login") ? "Login" : "Sign Up"}
          </button>
          <div className="text-center">
            Want to &nbsp;
            <Link
              to={!location.pathname.includes("/login") ? "/login" : "/signup"}
              className="text-blue-700 hover:underline"
            >
              {!location.pathname.includes("/login") ? "Login" : "Sign Up"}
            </Link>
          </div>
        </form>
        {loginOrSignup.isError && (
          <p className="text-red-500 mt-3">{loginOrSignup.error}</p>
        )}
      </div>
    </div>
  );
}
