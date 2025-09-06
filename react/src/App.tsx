import { Routes, Route, Navigate } from "react-router-dom";
import { useAuth } from "./context/AuthContext";
import Home from "./pages/Home";
import Posts from "./pages/Posts";
import LoginOrSignup from "./pages/LoginOrSignup";
import { useCheckSession } from "./api/auth";
import { useEffect } from "react";

export default function App() {
  const {
    isAuthenticated,
    setUser,
    setIsSessionLoading,
    setIsAuthenticated,
    isSessionLoading
  } = useAuth();
  const sessionQuery = useCheckSession();

  useEffect(() => {
    if (sessionQuery.isSuccess) {
      setUser(sessionQuery.data);
      setIsAuthenticated(true);
      setIsSessionLoading(false);
    }
    if (sessionQuery.isError) {
      setIsSessionLoading(false);
      setIsAuthenticated(false);
    }
  }, [sessionQuery.isFetching]);

  if (isSessionLoading) {
    return (
      <div className="flex h-screen items-center justify-center">
        <div className="h-12 w-12 animate-spin rounded-full border-4 border-gray-300 border-t-blue-600"></div>
      </div>
    );
  }

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route path="/signup" element={<LoginOrSignup />} />
      <Route path="/login" element={<LoginOrSignup />} />

      <Route
        path="/posts"
        element={isAuthenticated ? <Posts /> : <Navigate to="/login" />}
      />
    </Routes>
  );
}
