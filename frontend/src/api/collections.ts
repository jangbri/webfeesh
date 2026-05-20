import type { Collection } from "@/types/collection";
import type { Feed } from "@/types/feed";
import { api } from "./client";

export async function fetchCollections(): Promise<Collection[]> {
  const response = await api.get<Collection[]>(`/collection`);
  const collections = response.data;

  // sort the names alphabetically before displaying
  collections.sort((a, b) =>
    a.name.localeCompare(b.name, undefined, {
      sensitivity: "base",
    }),
  );

  return collections;
}

export async function fetchCollectionFeeds(id: number): Promise<Feed[]> {
  const response = await api.get<Feed[]>(`/collection/${id}`);

  return response.data;
}

export async function createCollection(data: Collection): Promise<Collection[]> {
  await api.post<Collection>(`/collection`, data);

  return fetchCollections();
}

export async function updateCollection(id: number, data: Collection): Promise<Collection[]> {
  await api.post<Collection>(`/collection/${id}`, data);

  return fetchCollections();
}

export async function deleteCollection(id: number): Promise<Collection[]> {
  await api.delete(`/collection/${id}`);

  return fetchCollections();
}
