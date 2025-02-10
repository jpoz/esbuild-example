export async function getQuote() {
	const response = await fetch('/api/quote');
	if (!response.ok) {
		throw new Error('Failed to fetch quote');
	}
	const quote = await response.json();
	return quote;
}
