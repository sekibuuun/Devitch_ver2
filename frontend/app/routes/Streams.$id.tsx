import type { LoaderFunctionArgs } from "@remix-run/node";
import { useLoaderData } from "@remix-run/react";
import type React from "react";
import type { Genre, Stream } from "~/types/types";

export const loader = async ({ params }: LoaderFunctionArgs) => {
	const streamId = params.id;
	try {
		const response = await fetch(`http://backend:8080/streams/${streamId}`);
		if (!response.ok) {
			throw new Error(`HTTP error! status: ${response.status}`);
		}
		const streamData: Stream = await response.json();

		const genrePromises: Promise<Genre>[] = streamData.genre_ids.map(
			async (genre_id) => {
				const genreResponse = await fetch(
					`http://backend:8080/genres/${genre_id}`,
				);
				if (!genreResponse.ok) {
					throw new Error(`HTTP error! status: ${genreResponse.status}`);
				}
				return genreResponse.json();
			},
		);

		const genres = await Promise.all(genrePromises)
			.then((genres) => genres)
			.catch((error: Error) => {
				console.error("Failed to fetch genres:", error);
				return [];
			});

		return { streamData, genres };
	} catch {
		throw new Response("Failed to load stream data", { status: 500 });
	}
};

const Streams: React.FC = () => {
	const { streamData, genres } = useLoaderData<typeof loader>();
	return (
		<div>
			<h1>title: {streamData?.title}</h1>
			<h2>Genres</h2>
			<ul>
				{genres.map((genre) => (
					<li key={genre?.genre_id}>{genre?.genre}</li>
				))}
			</ul>
		</div>
	);
};

export default Streams;
