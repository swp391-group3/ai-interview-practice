# Job description boundary

Own input/upload validation, analysis requests and extracted-skill review here. The wizard URL belongs to routing; future cross-step draft values belong to interview setup.

When the endpoint exists, split integration into api/requests.ts (transport + Zod decoder), api/keys.ts (domain keys), and api/mutations.ts (TanStack wrapper). Orchestration consumes that mutation and decides navigation; notifications stay at the use-case/UI boundary. No request or guessed response schema is shipped.

Forms use React Hook Form with zodResolver and z.infer<typeof schema>. Keep the schema beside the owning form. No input form is active while submission behavior is undefined.
