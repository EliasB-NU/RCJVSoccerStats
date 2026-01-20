import { createRouter, createWebHashHistory } from 'vue-router'

const router = createRouter({
  history: createWebHashHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/standings',
      name: 'standings',
      component: () => import('@/views/StandingsView.vue'),
    },
    {
      path: '/2v2matches',
      name: '2v2matches',
      component: () => import('@/views/2v2MatchesView.vue')
    },
    {
      path: '/entryMatches',
      name: 'entryMatches',
      component: () => import('@/views/EntryMatchesView.vue')
    },
    {
      path: '/lwlEntryMatches',
      name: 'lwlEntryMatches',
      component: () => import('@/views/LWLEntryMatchesView.vue')
    },
    {
      path: '/referee',
      name: 'referee',
      component: () => import('@/views/RefereeView.vue')
    }
  ],
})

export default router
