import type { ActionFunctionArgs } from "@remix-run/node";
import { json, redirect } from "@remix-run/node";
import { Form, useActionData, useNavigation } from "@remix-run/react";
import { useState } from "react";
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
		return json({ error: "タイトルは必須です" }, { status: 400 });
	}

	if (genreIds.length === 0) {
		return json(
			{ error: "ジャンルは1つ以上選択してください" },
			{ status: 400 },
		);
	}

	const response = await fetch("http://localhost:8080/streams", {
		method: "POST",
		headers: {
			"Content-Type": "application/json",
		},
		body: JSON.stringify({ title, genre_ids: genreIds }),
	});

	if (!response.ok) {
		return json(
			{ error: `配信の作成に失敗しました: ${response.statusText}` },
			{ status: response.status },
		);
	}
	return redirect("/hello");
}

export default function NewStream() {
	const [selectedGenres, setSelectedGenres] = useState<number[]>([]);
	const navigation = useNavigation();
	const error = useActionData<typeof action>();

	const handleGenreToggle = (genreId: number) => {
		setSelectedGenres((prev) =>
			prev.includes(genreId)
				? prev.filter((id) => id !== genreId)
				: [...prev, genreId],
		);
	};

	return (
		<div className="p-4">
			<Form method="post">
				<div className="grid w-full max-w-sm items-center gap-1.5 mb-4">
					<Label htmlFor="title">配信タイトル</Label>
					<Input
						type="text"
						name="title"
						id="title"
						placeholder="配信タイトルを入力してください"
					/>
				</div>
				<div className="flex flex-wrap gap-2 mb-4">
					<p>ジャンル</p>
					{genres.map((genre) => {
						const isSelected = selectedGenres.includes(genre.id);
						return (
							<label
								key={genre.id}
								className={`inline-flex items-center px-3 py-1 rounded-full text-sm cursor-pointer transition-colors duration-200 ${
									isSelected ? "bg-sky-800 text-white" : "bg-sky-400 text-white"
								}`}
							>
								<input
									type="checkbox"
									name="genre_ids"
									value={genre.id}
									className="mr-2 hidden"
									checked={isSelected}
									onChange={() => handleGenreToggle(genre.id)}
								/>
								{genre.name}
							</label>
						);
					})}
				</div>
				{error && <p className="text-red-500 py-2">{error.error}</p>}
				<Button type="submit" disabled={navigation.state === "submitting"}>
					{navigation.state === "submitting" ? "送信中..." : "配信開始"}
				</Button>
			</Form>
		</div>
	);
}
