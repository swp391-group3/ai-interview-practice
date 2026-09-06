import Link from "next/link";
import { routes } from "@/config/routes";
export default function NotFound() {
  return (
    <main className="mx-auto max-w-xl space-y-4 p-12">
      <h1 className="text-2xl font-semibold">Page not found</h1>
      <Link href={routes.home} className="underline">
        Return home
      </Link>
    </main>
  );
}
