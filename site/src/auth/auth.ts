import { jwtDecode, type JwtPayload } from "jwt-decode";

export enum UserRole {
  Admin,
  Viewer
}

export class CurrentUser {
  readonly token: string
  readonly claims: Array<string>
  readonly username: string

  constructor(token: string, username: string) {
    this.token = token
    this.claims = jwtDecode<CustomJwtPayload>(this.token).claims;
    this.username = username;
  }


}

interface CustomJwtPayload extends JwtPayload {
  claims: Array<string>
}

const userJWT = "UserJWT"
export function IsUserLoggedIn(): boolean {
  if (localStorage.getItem(userJWT) === null) {
    return false;
  }
  return true;
}

export function CanLoggedInUserCreatePost(): boolean {
  let jwt = localStorage.getItem(userJWT);
  if (jwt === null) {
    return false;
  }
  let headers = jwtDecode<CustomJwtPayload>(jwt)
  return headers.claims.includes(userJWT);
}

export function StoreUserJWT(token: string) {
  localStorage.setItem(userJWT, token);
}
