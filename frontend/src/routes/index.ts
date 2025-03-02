import { createWebHistory, createRouter } from 'vue-router'
import { NamesRouter } from './routesNames'
const routes = [
  {
    path: '/',
    children:[
      {
        path: '/',
        component: () => import('@/modules/mails/views/MailsView.vue'),
        name:NamesRouter.Index,
      },
      {
        path: '/:id',
        component: () => import('@/modules/mails/views/MailsDetailsView.vue'),
        name:NamesRouter.MailDetail
      }
    ]
  },

]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

export default router
