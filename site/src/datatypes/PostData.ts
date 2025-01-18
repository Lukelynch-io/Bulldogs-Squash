export interface PostData {
  id: string,
  title: string,
  description: string,
  imageUrl: string | null
}

export interface NewPost {
  title: string,
  description: string,
  imageUrl: string | null
}
