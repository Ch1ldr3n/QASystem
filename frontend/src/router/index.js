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
        name: "login",
        component: Login
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

export default router