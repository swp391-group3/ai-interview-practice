export class ApiError extends Error {
  constructor(
    message: string,
    public readonly status: number,
    public readonly body: unknown,
  ) {
    super(message);
    this.name = "ApiError";
  }
}
export class ApiConfigurationError extends Error {
  constructor() {
    super("API base URL is not configured.");
    this.name = "ApiConfigurationError";
  }
}
