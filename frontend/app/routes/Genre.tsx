import { json, useLoaderData } from "@remix-run/react";
import type React from "react";
import type { Genre } from "~/types/types";

export const loader = async () => {
	const response = await fetch("http://localhost:8080/genres");
	const genreData: Genre[] = await response.json();

	return json({ genreData });
};

const Hello: React.FC = () => {
	const { genreData } = useLoaderData<typeof loader>();
	return (
		<div>
			<ul>
				{genreData.map((genre) => (
					<li key={genre.id}>{genre.genre}</li>
				))}
			</ul>
		</div>
	);
};

export default Hello;
