import type { ActionFunctionArgs } from "@remix-run/node"; // or cloudflare/deno
import { redirect } from "@remix-run/node"; // or cloudflare/deno
import { Form } from "@remix-run/react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Label } from "~/components/ui/label";

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
		<div className="p-4">
			<Form method="post">
				<div className="grid w-full max-w-sm items-center gap-1.5">
					<Label htmlFor="title">配信タイトル</Label>
					<Input
						type="text"
						name="title"
						id="title"
						placeholder="配信タイトルを入力してください"
					/>
				</div>
				{genres.map((genre) => (
					<label
						key={genre.id}
						className="inline-flex items-center px-3 py-1 rounded-full text-sm cursor-pointer transition-colors duration-200 bg-sky-400 text-white"
					>
						<input
							type="checkbox"
							name="genre_ids"
							value={genre.id}
							className="mr-2 hidden"
						/>
						{genre.name}
					</label>
				))}
				<Button type="submit">Start Stream</Button>
			</Form>
		</div>
	);
}
