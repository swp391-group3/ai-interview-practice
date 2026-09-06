export type ConnectionStatus =
  "disconnected" | "connecting" | "connected" | "reconnecting";
export type ParseResult<T> =
  { success: true; message: T } | { success: false; error: string };
/** Supply a Zod-backed parser once the backend message schema exists. */
export type ProtocolParser<T> = (raw: unknown) => ParseResult<T>;
/** Generic intentionally: no wire message shape or fake endpoint is established here. */
export interface InterviewTransport<Incoming, Outgoing> {
  connect(signal?: AbortSignal): Promise<void>;
  disconnect(): void;
  send(message: Outgoing): void;
  subscribe(listener: (message: Incoming) => void): () => void;
  subscribeStatus(listener: (status: ConnectionStatus) => void): () => void;
}
