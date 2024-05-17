import ProfileSelectionBtn from "@/components/profile_button"

export default function Home() {

	return (
		<main className="flex flex-col min-h-screen items-center justify-center gap-8">
			<h1 className="text-6xl">Wer möchte gerade Kochen?</h1>
			<div className="grid grid-cols-2 gap-4 md:grid-cols-4">
				<ProfileSelectionBtn />
				<ProfileSelectionBtn />
				<ProfileSelectionBtn />
				<ProfileSelectionBtn />
			</div>
		</main>
	)
}
