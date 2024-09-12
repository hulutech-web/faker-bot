export default [
    {
        path: '/',
        name: '/',//布局路由
        redirect:'/home',
    },
    {
        path: '/home',
        name: 'home',//布局路由
        component: () => import('@/view/home.vue'),
    },
    {
        path: '/category',
        name: 'category',//布局路由
        component: () => import('@/view/category.vue'),
    },
    {
        path: '/video',
        name: 'video',//布局路由
        component: () => import('@/view/video.vue'),
    },
    {
        path: '/setting',
        name: 'setting',//布局路由
        component: () => import('@/view/setting.vue'),
    },
];