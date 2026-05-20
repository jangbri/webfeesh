import type { Collection } from "@/types/collection";
import { api } from "./client";

export async function fetchCollections(): Promise<Collection[]> {
  const response = await api.get<Collection[]>(`/collection`);
  console.log(response.data);

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
  await api.delete<Promise<void>>(`/collection/${id}`);

  return fetchCollections();
}
