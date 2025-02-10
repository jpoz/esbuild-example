import React from "react";
import QuoteButton from "./components/QuoteButton";

function App() {
	return (
		<main className="flex min-h-screen flex-col items-center justify-center p-24">
			<h1 className="text-4xl font-bold mb-8">Quote Fetcher</h1>
			<QuoteButton />
		</main>
	)
}

export default App;
