import { json, useLoaderData } from "@remix-run/react";
import type { Genre } from "~/types/types";

export const loader = async () => {
	try {
		const response = await fetch("http://backend:8080/genres");
		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const genreData: Genre[] = await response.json();
		return json({ genreData });
	} catch (error) {
		console.error("Failed to fetch genres:", error);
		return json({ genreData: [] }, { status: 500 });
	}
};

function GenreComponent() {
	const { genreData } = useLoaderData<typeof loader>();
	return (
		<div>
			<ul>
				{genreData.map((genre) => (
					<li key={genre.genre_id}>{genre.genre}</li>
				))}
			</ul>
		</div>
	);
}

export default GenreComponent;
