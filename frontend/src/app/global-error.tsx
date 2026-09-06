"use client";
export default function GlobalError({
  reset,
}: {
  error: Error & { digest?: string };
  reset: () => void;
}) {
  return (
    <html lang="en">
      <body>
        <main style={{ padding: "3rem", fontFamily: "sans-serif" }}>
          <h1>Something went wrong</h1>
          <p>The application could not load. Please try again.</p>
          <button onClick={reset}>Try again</button>
        </main>
      </body>
    </html>
  );
}
