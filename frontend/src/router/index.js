import { createRouter, createWebHashHistory } from 'vue-router'

import Explore from '../components/Explore.vue'
import User from '../components/User.vue'
import Question from '../components/Question.vue'
import Login from '../components/Login.vue'
import Register from '../components/Register.vue'
import Pay from '../views/pay.vue'
import Submit from '../views/submit-question.vue'
import Answerer from '../components/Answerer.vue'

const routes = [
  {
    path: '/',
    name: 'Explore',
    component: Explore,
    meta: {
      public: true,
    },
  },
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: {
      public: true,
    },
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: {
      public: true,
    },
  },
  {
    path: '/user',
    name: 'User',
    component: User,
  },
  {
    path: '/question',
    name: 'Question',
    component: Question,
  },
  {
    path: '/answerer',
    name: 'Answerer',
    component: Answerer,
  },

  {
    path: '/pay',
    name: 'Pay',
    component: Pay,
  },

  {
    path: '/submit',
    name: 'Submit',
    component: Submit,
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  routes: routes,
})
router.beforeEach((to, from, next) => {
  if (!to.matched.some((record) => record.meta.public)) {
    if (window.localStorage.getItem('token') == null) {
      next({ name: 'Login' })
    } else {
      next()
    }
  } else {
    next()
  }
})

export default router
