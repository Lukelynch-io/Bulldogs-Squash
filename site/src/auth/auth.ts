import { jwtDecode, type JwtPayload } from "jwt-decode";

export enum UserRole {
  Admin,
  Viewer
}

interface CustomJwtPayload extends JwtPayload {
  claims: Array<string>
}

export function CanLoggedInUserCreatePost(): boolean {
  let jwt = localStorage.getItem("UserJWT");
  if (jwt === null) {
    return false;
  }
  let headers = jwtDecode<CustomJwtPayload>(jwt)
  return headers.claims.includes("CREATE_BLOG");
}

export function StoreUserJWT(token: string) {
  localStorage.setItem("UserJWT", token);
}
