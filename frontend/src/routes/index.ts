import { createWebHistory, createRouter } from 'vue-router'

import MailsView from '@/modules/mails/views/MailsView.vue'
import MailDetailsView from '@/modules/mails/views/MailsDetailsView.vue'
import { NamesRouter } from './routesNames'
const routes = [
  {
    path: '/',
    children:[
      {
        path: '/',
        component: MailsView,
        name:NamesRouter.Index,
      },
      {
        path: '/:id',
        component: MailDetailsView,
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
