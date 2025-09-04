import {
  createContext,
  useContext,
  useState,
  ReactNode,
  useEffect
} from "react";
import { checkSession } from "../api/auth";

interface AuthContextType {
  isAuthenticated: boolean;
  user: any | null;
  setIsAuthenticated: (value: boolean) => void;
  setUser: (user: any | null) => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [user, setUser] = useState<any | null>(null);

  useEffect(() => {
    (async () => {
      try {
        const userData = await checkSession();
        setUser(userData);
        setIsAuthenticated(true);
      } catch {
        setUser(null);
        setIsAuthenticated(false);
      }
    })();
  }, []);

  return (
    <AuthContext.Provider
      value={{ isAuthenticated, user, setIsAuthenticated, setUser }}
    >
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within AuthProvider");
  return context;
}
