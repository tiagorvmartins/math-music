import { createRouter, createWebHistory } from 'vue-router'
import SingleView from '../views/SingleView.vue'
import Login from "../views/Login.vue";
import store from '@/store'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/link',
      name: 'Single Download',
      component: SingleView,      
    },
    {
      path: "/login",
      name: "Login",
      component: Login,
    },
    {
      path: "/",
      name: "Login",
      component: Login,
    }
  ]
})

router.beforeEach(async (to, from) => {  
  if (    
    !store.getters['isLoggedIn'] &&    
    to.name !== 'Login'
  ) {    
    return { name: 'Login' }
  }
})

export default router
