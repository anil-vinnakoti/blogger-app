import { useLogout } from "../api/auth";
import { useAuth } from "../context/AuthContext";

export default function LogoutButton() {
  const logout = useLogout();
  const { setIsAuthenticated } = useAuth();

  const handleLogout = () => {
    logout.mutate(undefined, {
      onSuccess: () => setIsAuthenticated(false)
    });
  };

  return (
    <button
      onClick={handleLogout}
      className="bg-red-600 text-white px-4 py-2 rounded"
    >
      Logout
    </button>
  );
}
