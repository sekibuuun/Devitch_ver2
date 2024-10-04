import { json, useLoaderData } from "@remix-run/react";
import type React from "react";
import type { Hello as HelloType } from "~/types/types";

export const loader = async () => {
	const hellosResponse = await fetch("http://backend:8080/hello");
	const hellos: HelloType[] = await hellosResponse.json();

	return json({ hellos });
};

const Hello: React.FC = () => {
	const { hellos } = useLoaderData<typeof loader>();
	return (
		<div>
			{hellos.map((hello) => (
				<p key={hello.id}>{hello.content}</p>
			))}
		</div>
	);
};

export default Hello;
