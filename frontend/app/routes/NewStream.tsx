import type { ActionFunctionArgs } from "@remix-run/node";
import { json, redirect } from "@remix-run/node";
import {
	Form,
	useActionData,
	useLoaderData,
	useNavigation,
} from "@remix-run/react";
import { useState } from "react";
import { Button } from "~/components/ui/button";
import { Input } from "~/components/ui/input";
import { Label } from "~/components/ui/label";
import type { Genre } from "~/types/types";

export const loader = async () => {
	try {
		const response = await fetch("http://localhost:8080/genres");
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
	const { genreData } = useLoaderData<typeof loader>();
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
					{genreData.map((genre) => {
						const isSelected = selectedGenres.includes(genre.genre_id);
						return (
							<label
								key={genre.genre_id}
								className={`inline-flex items-center px-3 py-1 rounded-full text-sm cursor-pointer transition-colors duration-200 ${
									isSelected ? "bg-sky-800 text-white" : "bg-sky-400 text-white"
								}`}
							>
								<input
									type="checkbox"
									name="genre_ids"
									value={genre.genre_id}
									className="mr-2 hidden"
									checked={isSelected}
									onChange={() => handleGenreToggle(genre.genre_id)}
								/>
								{genre.genre}
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
