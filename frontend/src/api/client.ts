import axios from 'axios'

const isProd = import.meta.env.PROD

export const apiURL = isProd ? window.location.origin : import.meta.env.VITE_BACKEND_BASE_URL

export const api = axios.create({
  baseURL: apiURL,
  headers: {
    'Content-Type': 'application/json',
  },
})
