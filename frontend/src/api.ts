// src/api.ts
export async function apiFetch(input: RequestInfo, init: RequestInit = {}) {
  const headers = new Headers(init.headers)

  // 从 localStorage / sessionStorage 取 token
  const token = localStorage.getItem('auth_token') || sessionStorage.getItem('auth_token')
  if (token) headers.set('Authorization', `Bearer ${token}`)

  const res = await fetch(input, { ...init, headers })

  if (res.status === 401) {
    // 登录过期处理：清掉 token，跳回登录页
    localStorage.removeItem('auth_token')
    sessionStorage.removeItem('auth_token')
    location.hash = '#/login'
    throw new Error('未授权，请重新登录')
  }

  return res
}
