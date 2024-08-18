import type { ActionFunctionArgs } from "@remix-run/node"; // or cloudflare/deno
import { redirect } from "@remix-run/node"; // or cloudflare/deno
import { Form } from "@remix-run/react";

const genres = [
	{ id: 1, name: "Python" },
	{ id: 2, name: "JavaScript" },
	{ id: 3, name: "Java" },
	{ id: 4, name: "C++" },
	{ id: 5, name: "C#" },
	{ id: 6, name: "PHP" },
	{ id: 7, name: "TypeScript" },
	{ id: 8, name: "Swift" },
	{ id: 9, name: "Go" },
	{ id: 10, name: "Rust" },
];

export async function action({ request }: ActionFunctionArgs) {
	const body = await request.formData();
	const title = body.get("title");
	const genreIds = body
		.getAll("genre_ids")
		.map((id) => Number.parseInt(id as string, 10));

	if (!title) {
		return new Response("Title is required", {
			status: 400,
		});
	}

	const response = await fetch("http://localhost:8080/streams", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ title, genre_ids: genreIds }),
	});

	if (!response.ok) {
		return new Response(`Error creating stream: ${response.statusText}`, {
			status: response.status,
		});
	}
	return redirect("/hello");
}

export default function NewStream() {
	return (
		<div>
			<Form method="post">
				<input type="text" name="title" />
				<div>
					{genres.map((genre) => (
						<label key={genre.id}>
							<input type="checkbox" name="genre_ids" value={genre.id} />
							{genre.name}
						</label>
					))}
				</div>
				<button type="submit">Start Stream</button>
			</Form>
		</div>
	);
}
