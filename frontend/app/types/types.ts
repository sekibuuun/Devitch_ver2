export type Hello = {
	id: number;
	content: string;
};

export type Genre = {
	genre_id: number;
	genre: string;
};

export type Stream = {
	stream_id: number;
	title: string;
	genre_ids: number[];
	comment_length: number;
	comments: Comment[];
	listener_length: number;
	listeners: User[];
};

export type Comment = {
	comment_id: number;
	stream_id: number;
	user_id: number;
	content: string;
	send_at: string;
};

export type User = {
	user_id: number;
	user_name: string;
	links: string[];
};
