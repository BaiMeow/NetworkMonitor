import { ref } from 'vue'
const tokenRef = ref<string | null>(localStorage.getItem('auth_token'))
export function useAuth() {
  return {
    isAuthed: () => !!tokenRef.value,
    getToken: () => tokenRef.value,
    setToken: (t: string | null) => {
      tokenRef.value = t
      if (t) localStorage.setItem('auth_token', t)
      else localStorage.removeItem('auth_token')
    }
  }
}
