import { createRouter, createWebHashHistory } from "vue-router";

import home from '../components/Homepage.vue'
import User from '../components/User.vue'

const routes = [
    {
        path: "/",
        name: "home",
        component: home
    },
    {
        path: "/user",
        name: "user",
        component: User
    }
]

const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})

export default router