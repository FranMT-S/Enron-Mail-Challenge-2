import { createWebHistory, createRouter } from 'vue-router'

import MailsView from '@/modules/mails/views/MailsView.vue'
import MailDetailsView from '@/modules/mails/views/MailsDetailsView.vue'
import { NamesRouter } from './routesNames'
const routes = [
  { path: '/',
    component: MailsView,
    name:NamesRouter.Index,
    // children:[
    //   {
    //   path: '/mails',
    //   component: MailDetailsView,
    //   name:NamesRouter.MailDetail
    //   }
    // ]
  },
  {
    path: '/:id',
    component: MailDetailsView,
    name:NamesRouter.MailDetail
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
