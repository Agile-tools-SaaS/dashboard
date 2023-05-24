import {
  ReactNode,
  createContext,
  useContext,
  useEffect,
  useMemo,
  useState,
} from "react";

interface AuthContextType {
  // NEEDS PROPER TYPING
  user?: any;
  loading: boolean;
  login: (userParams: any, callback?: any) => void;
  logout: (callback?: any) => void;
  loggedIn: boolean;
}

const AuthContext = createContext<AuthContextType>({} as AuthContextType);

export function AuthProvider({
  children,
}: {
  children: ReactNode;
}): JSX.Element {
  // THIS NEEDS REFACTORED WITH THE RIGHT TYPES
  const [user, setUser] = useState<any>();
  const [loggedIn, setLoggedIn] = useState<boolean>(false);
  const [loading, setLoading] = useState<boolean>(false);
  const [loadingInitial, setLoadingInitial] = useState<boolean>(true);

  // NEEDS TYPED PROPERLY AND BOTH NEED FILLED OUT
  // - maybe also needs a check for if the user is still logged in. (not sure if this is needed as a 403 should have
  // logic to check if the user is supposed to be here and if they should be redirected to login or main page.)
  function login(
    { user, password }: any,
    callback?: ({ user, password }: any, callback?: any) => void
  ) {}
  function logout(callback?: Function) {}

  const memoedValue = useMemo(
    () => ({
      user,
      loading,
      login,
      logout,
      loggedIn,
    }),
    [user, loading, loggedIn]
  );

  return (
    <AuthContext.Provider value={memoedValue}>
      {!loadingInitial && children}
    </AuthContext.Provider>
  );
}

export default function useAuth() {
  return useContext(AuthContext);
}
