import { createRouter, createWebHashHistory } from "vue-router";

const home = () => import("../components/Homepage")
const user = () => import("../components/User")

const routes = [
    { path: "/", redirect: "/home" },
    {
        path: "/home",
        name: "home",
        component: home
    },
    {
        path: "/user",
        name: "user",
        component: user
    }
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes: routes
})