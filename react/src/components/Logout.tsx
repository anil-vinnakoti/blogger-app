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
      className="font-bold px-3 py-1 border-2 rounded-lg"
    >
      <span className="align-[2px]">Logout</span>
    </button>
  );
}
