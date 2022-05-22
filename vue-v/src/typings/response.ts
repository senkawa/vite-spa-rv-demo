export interface Response {
  data: Post[];
}

export interface Post {
  ID: string;
  CreatedAt: string;
  UpdatedAt: string;
  Body: string;
}
