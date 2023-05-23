export interface ApiResponse<T> {
  status: number;
  body: T;
  message: string;
}
