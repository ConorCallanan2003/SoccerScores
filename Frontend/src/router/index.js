import { createRouter, createWebHistory } from 'vue-router'
import LeagueTable from '../components/LeagueTable.vue'

const LeagueTable = () => import ('./views/LeagueTable.vue')

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/table',
      name: 'Table',
      component: LeagueTable
    },
  ]
})

export default router
