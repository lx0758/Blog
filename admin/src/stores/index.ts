import { ref } from 'vue'
import { createPinia, defineStore } from 'pinia'
import piniaPluginPersist from 'pinia-plugin-persist'

export const useUserStore = defineStore(
  'user-store',
  () => {
    const isLoggedIn = ref(false)

    function login() {
      isLoggedIn.value = true
    }
    function logout() {
      isLoggedIn.value = false
    }

    return { isLoggedIn, login, logout }
  },
  {
    persist: {
      enabled: true,
      strategies: [
        {
          storage: sessionStorage
        }
      ]
    }
  }
)

export const useXXXXStore = defineStore(
  'xxxx-store',
  () => {
    const a = ref(0)
    const b = ref(1)
    const c = ref(2)
    const d = ref(3)

    function method() { }

    return { a, b, c, d, method }
  },
  {
    persist: {
      enabled: true,
      strategies: [
        {
          storage: sessionStorage,
          paths: ['a', 'b']
        },
        {
          storage: localStorage,
          paths: ['c', 'd']
        }
      ]
    }
  }
)

const store = createPinia()
store.use(piniaPluginPersist)

export default store
