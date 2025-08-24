import { ref } from 'vue'

const tokenRef = ref<string | null>(localStorage.getItem('auth_token'))

export function useAuth() {
  function isAuthed() {
    return !!tokenRef.value
  }
  function setToken(t: string | null) {
    tokenRef.value = t
    if (t) localStorage.setItem('auth_token', t)
    else localStorage.removeItem('auth_token')
  }
  function getToken() {
    return tokenRef.value
  }
  return { isAuthed, setToken, getToken }
}
