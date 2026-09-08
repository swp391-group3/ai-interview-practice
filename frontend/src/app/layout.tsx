import type { Metadata } from "next";
import localFont from "next/font/local";
import "./globals.css";
import { AppProviders } from "@/providers/app-providers";

const albertSans = localFont({
  src: [
    {
      path: "../../design/exploration/datn/assets/fonts/open-design.ai/AlbertSans-VariableFont_wght-5ccc2f8f05.woff2",
      weight: "100 900",
      style: "normal",
    },
    {
      path: "../../design/exploration/datn/assets/fonts/open-design.ai/AlbertSans-Italic-VariableFont_wght-b6ad8b53d7.woff2",
      weight: "100 900",
      style: "italic",
    },
  ],
  display: "swap",
  variable: "--font-albert-sans",
});

export const metadata: Metadata = {
  title: "RoleCue — Technical interview practice",
  description: "Role-focused practice for technical interviews.",
};
export default function RootLayout({
  children,
}: Readonly<{ children: React.ReactNode }>) {
  return (
    <html lang="en">
      <body
        className={`${albertSans.variable} min-h-screen bg-background font-sans text-foreground antialiased`}
      >
        <AppProviders>{children}</AppProviders>
      </body>
    </html>
  );
}
