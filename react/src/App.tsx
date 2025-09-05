import { Routes, Route, Navigate } from "react-router-dom";
import { useAuth } from "./context/AuthContext";
import Home from "./pages/Home";
import Posts from "./pages/Posts";
import LoginOrSignup from "./pages/LoginOrSignup";

export default function App() {
  const { isAuthenticated, user } = useAuth();

  if (user === null && !isAuthenticated) {
    return <LoginOrSignup />;
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
      <Route path="/loginOrSignup" element={<LoginOrSignup />} />
    </Routes>
  );
}
