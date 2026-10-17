// Next.js needs a root layout — minimal on purpose.
export const metadata = { title: "stampede", description: "agent fleet control plane" };

export default function RootLayout({ children }: { children: React.ReactNode }) {
  return (
    <html lang="en">
      <body style={{ margin: 0, background: "#0a0a0a", color: "#eaeaea" }}>
        {children}
      </body>
    </html>
  );
}
