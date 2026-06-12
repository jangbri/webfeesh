import axios from 'axios'

const isProd = import.meta.env.PROD

export const apiURL = window.location.origin + '/api'

export const api = axios.create({
  baseURL: apiURL,
  headers: {
    'Content-Type': 'application/json',
  },
})
