import { createRouter, createWebHashHistory } from "vue-router";

import home from '../components/Homepage.vue'
import User from '../components/User.vue'
import Issues from '../components/Issues.vue'
import Login from '../components/Login.vue'

const routes = [
    {
        path: "/",
        name: "home",
        component: home
    },
    {
        path: "/login",
        name: "Login",
        component: Login,
        meta: {
            public: true
        }
    },
    {
        path: "/user",
        name: "user",
        component: User
    },
    {
        path: "/issues",
        name: "issues",
        component: Issues
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})
router.beforeEach((to, from, next) => {
    if (!to.matched.some(record => record.meta.public)) {
        if (window.localStorage.getItem("token") == null) {
            next({ name: 'Login' })
        } else {
            next()
        }
    } else {
        next()
    }
})

export default router