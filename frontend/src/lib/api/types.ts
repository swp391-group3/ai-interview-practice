/** Auth integration supplies headers without exposing token storage to features. */
export type AuthHeadersProvider = () => HeadersInit | Promise<HeadersInit>;
export type Fetcher = (
  input: RequestInfo | URL,
  init?: RequestInit,
) => Promise<Response>;
export type ResponseDecoder<T> = (value: unknown) => T;
export type ApiRequest<T> = Omit<RequestInit, "body"> & {
  body?: BodyInit | null;
  decode: ResponseDecoder<T>;
};
export interface ApiClient {
  request<T>(path: string, options: ApiRequest<T>): Promise<T>;
}
