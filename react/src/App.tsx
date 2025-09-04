import { Routes, Route, Navigate } from "react-router-dom";
import { useAuth } from "./context/AuthContext";
import Home from "./pages/Home";
import Posts from "./pages/Posts";
import Login from "./pages/Login";

export default function App() {
  const { isAuthenticated, user } = useAuth();

  if (user === null && !isAuthenticated) {
    return <Login />;
  }

  return (
    <Routes>
      <Route path="/" element={<Home />} />
      <Route
        path="/posts"
        element={isAuthenticated ? <Posts /> : <Navigate to="/login" />}
      />
      <Route path="/login" element={<Login />} />
    </Routes>
  );
}
