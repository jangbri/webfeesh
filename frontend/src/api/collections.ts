import type { Collection } from '@/types/collection'
import type { Feed } from '@/types/feed'
import { api } from './client'

export async function fetchCollections(): Promise<Collection[]> {
  const response = await api.get<Collection[]>(`/collections/`)
  const collections = response.data

  // sort the names alphabetically before displaying
  collections.sort((a, b) =>
    a.name.localeCompare(b.name, undefined, {
      sensitivity: 'base',
    }),
  )

  return collections
}

export async function fetchCollectionFeeds(id: number): Promise<Feed[]> {
  const response = await api.get<Feed[]>(`/collections/${id}`)
  const feeds = response.data

  // sort the names alphabetically before displaying
  feeds.sort((a, b) =>
    a.title.localeCompare(b.title, undefined, {
      sensitivity: 'base',
    }),
  )

  return feeds
}

export async function createCollection(data: Collection): Promise<void> {
  await api.post<Collection>(`/collections/`, data)
}

export async function updateCollection(data: Collection): Promise<void> {
  await api.post<Collection>(`/collections/${data.id}`, data)
}

export async function deleteCollection(data: Collection): Promise<void> {
  await api.delete(`/collections/${data.id}`)
}
