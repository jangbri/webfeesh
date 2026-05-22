import type { Feed } from "@/types/feed";
import type { Item } from "@/types/item";
import { api } from "./client";
import { fetchCollectionFeeds } from "./collections";

export async function fetchFeedItems(data: Feed): Promise<Item[]> {
  const response = await api.get<Item[]>(`/feed/${data.id}`);

  return response.data;
}

export async function createFeed(data: Feed): Promise<Feed[]> {
  await api.post<Feed>(`/feed`, data);

  return fetchCollectionFeeds(data);
}

export async function updateFeed(data: Feed): Promise<Feed[]> {
  await api.post<Feed>(`/feed/${data.id}`, data);

  return fetchCollectionFeeds(data.collection_id);
}

export async function deleteFeed(data: Feed): Promise<Feed[]> {
  await api.delete(`/feed/${data.id}`);

  return fetchCollectionFeeds(data.collection_id);
}
