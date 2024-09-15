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
        path: '/authorize',
        name: 'authorize',//用户授权
        component: () => import('@/view/authorize.vue'),
    },
    {
        path: '/setting',
        name: 'setting',//布局路由
        component: () => import('@/view/setting.vue'),
    },
];