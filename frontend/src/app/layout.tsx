import type { Metadata } from 'next'
import './globals.css'

export const metadata: Metadata = {
	title: 'Foodiee',
	description: 'Foodiee Recipe App',
}

export default function RootLayout({
	children,
}: {
	children: React.ReactNode
}) {
	return (
		<html lang="de">
			<body className="bg-foodiee-black-primary text-foodiee-white">{children}</body>
		</html>
	)
}
