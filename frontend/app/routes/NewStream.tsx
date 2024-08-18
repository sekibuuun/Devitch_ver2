import type { ActionFunctionArgs } from "@remix-run/node"; // or cloudflare/deno
import { redirect } from "@remix-run/node"; // or cloudflare/deno
import { Form } from "@remix-run/react";

export async function action({ request }: ActionFunctionArgs) {
	const body = await request.formData();
	const title = body.get("title");
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
		body: JSON.stringify({ title, genre_ids: [1, 6] }),
	});

	if (!response.ok) {
		throw new Error(`HTTP error! status: ${response.status}`);
	}
	return redirect("/hello");
}

export default function NewStream() {
	return (
		<div>
			<Form method="post">
				<input type="text" name="title" />
				<button type="submit">Start Stream</button>
			</Form>
		</div>
	);
}
