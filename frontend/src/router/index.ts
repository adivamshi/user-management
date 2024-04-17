import { createRouter, createWebHistory } from 'vue-router';
import Home from '../views/Home.vue';
import AddUser from '../views/AddUser.vue';
import UserDetail from '../views/UserDetail.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes:[
    {
      path: '/',
      name: 'Home',
      component: Home
    },
    {
      path: '/add-user',
      name: 'AddUser',
      component: AddUser
    },
    {
      path: '/user/:id',
      name: 'UserDetails',
      component: UserDetail,
      props: true
    }
  ]
});

export default router;
