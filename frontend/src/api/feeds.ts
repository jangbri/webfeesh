import type { Feed } from '@/types/feed'
import type { Item } from '@/types/item'
import { api } from './client'

export async function fetchFeedItems(data: Feed): Promise<Item[]> {
  const response = await api.get<Item[]>(`/feed/${data.id}`)

  return response.data
}

export async function createFeed(data: Feed): Promise<void> {
  await api.post<Feed>(`/feed`, data)
}

export async function updateFeed(data: Feed): Promise<void> {
  await api.post<Feed>(`/feed/${data.id}`, data)
}

export async function deleteFeed(data: Feed): Promise<void> {
  await api.delete(`/feed/${data.id}`)
}
