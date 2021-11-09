import { createRouter, createWebHashHistory } from 'vue-router';

import Explore from '../components/Explore.vue';
import User from '../components/User.vue';
import Question from '../components/Question.vue';
import Login from '../components/Login.vue';
import Register from '../components/Register.vue';
import Pay from '../views/pay.vue';
import Submit from '../views/submit-question.vue';
import Answerer from '../components/Answerer.vue';
import AdminLogin from '../components/AdminLogin.vue';
import AdminHomepage from '../components/AdminHomepage.vue';
import AdminUser from '../components/AdminUser.vue';
import Income from '../components/Income.vue';
import Review from '../components/Review.vue';

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
  {
    path: '/admin',
    name: 'AdminLogin',
    component: AdminLogin,
    meta: {
      public: true,
    },
  },
  {
    path: '/admin/homepage',
    name: 'AdminHomepage',
    component: AdminHomepage,
    meta: {
      public: true,
    },
  },
  {
    path: '/admin/user',
    name: 'AdminUser',
    component: AdminUser,
    meta: {
      public: true,
    },
  },
  {
    path: '/income',
    name: 'Income',
    component: Income,
  },
  {
    path: '/admin/review',
    name: 'Review',
    component: Review,
    meta: {
      public: true,
    },
  },
  // {
  //   path: '/admin/adminlist',
  //   name: 'AdminList',
  //   component: AdminList,
  //   meta: {
  //     public: true,
  //   }
  // },
  // {
  //   path: '/admin/sysconfig',
  //   name: 'SystemConfig',
  //   component: SystemConfig,
  //   meta: {
  //     public: true,
  //   }
  // },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});
router.beforeEach((to, from, next) => {
  if (!to.matched.some((record) => record.meta.public)) {
    if (window.localStorage.getItem('token') == null) {
      next({ name: 'Login' });
    } else {
      next();
    }
    if (window.localStorage.getItem('admintoken') == null) {
      next({ name: 'AdminLogin' });
    } else {
      next();
    }
  } else {
    next();
  }
});

export default router;
