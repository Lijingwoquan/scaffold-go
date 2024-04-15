import {
    createRouter,
    createWebHistory 
} from 'vue-router'

import Index from '~/pages/index.vue'
import NotFound from "~/pages/404.vue"

const routes = [
    { path: "/", component: Index },
    { path: '/:pathMatch(.*)*', name: 'NotFound', component: NotFound }
]



export const router = createRouter({
    history: createWebHistory(),
    routes
})

