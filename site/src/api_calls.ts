import axios from "axios";
import type { PostData } from "./datatypes/PostData";

export async function GetPosts(): Promise<[PostData[], string]> {
  let returnArray: PostData[] = [];
  let returnError: string = "";
  await axios.get("/api/posts?trimmed=true")
    .then((response) => {
      let array = response.data as PostData[]
      returnArray = array
    })
    .catch((error) => {
      returnError = "Error fetching blog posts: " + error
    });
  return [returnArray, returnError]
}

export async function GetPost(postId: string): Promise<PostData> {
  return await axios.get(`/api/post/${postId}`)
    .then((response) => {
      return response.data as PostData
    })
}

export async function RequestUserToken(username: string, passwordHash: string): Promise<string> {
  return await axios.post("/api/auth/token", {
    username: username,
    passwordHash: passwordHash
  }, {
    headers: {
      'Content-Type': 'application/json'
    }
  }).then((response) => {
    return response.data.token
  })
}

export async function GetUsername(token: string): Promise<string> {
  try {
    return await axios.get("/api/auth/user", {
      headers: {
        "Authorization": "Bearer " + token
      }
    }).then((response) => response.data.username)
  } catch (_) {
    return "";
  }
}
