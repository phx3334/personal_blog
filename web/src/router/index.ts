import {createRouter, createWebHistory} from 'vue-router'
import {useUserStore} from "@/stores/user";
import {useLayoutStore} from "@/stores/layout";
import {ElMessageBox} from "element-plus"; // 添加 ElMessageBox 导入

const routes = [
    {
        path: '/',
        name: 'web',
        component: () => import('@/views/web/index.vue'),
        children: [
            {
                path: "/",
                name: "index",
                component: () => import('@/views/web/index/index.vue'),
                meta: {
                    title: "首页"
                }
            },
            {
                path: "search",
                name: "search",
                component: () => import('@/views/web/search/index.vue'),
                meta: {
                    title: "搜索"
                }
            },
            {
                path: "anime",
                name: "anime",
                component: () => import('@/views/web/anime/index.vue'),
                meta: {
                    title: "番剧推荐"
                }
            },
            {
                path: "game",
                name: "game",
                component: () => import('@/views/web/game/index.vue'),
                meta: {
                    title: "游戏推荐"
                }
            },
            {
                path: "about",
                name: "about",
                component: () => import('@/views/web/about/index.vue'),
                meta: {
                    title: "关于"
                }
            }
        ]
    },
    {
        path: "/login",
        name: "login",
        component: () => import('@/views/login/index.vue')
    },
    {
        path: "/article/:id",
        name: "article",
        component: () => import('@/views/web/article/index.vue')
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: () => import('@/views/dashboard/index.vue'),
        meta: {
            title: "控制面板",
            requiresAuth: true
        },
        children: [
            {
                path: "/dashboard/",
                name: "home",
                component: () => import('@/views/dashboard/home/index.vue'),
                meta: {
                    title: "主页"
                },
            },
            {
                path: "user-center",
                name: "user-center",
                meta: {
                    title: "个人中心"
                },
                children: [
                    {
                        path: "user-info",
                        name: "user-info",
                        component: () => import('@/views/dashboard/user-center/user-info.vue'),
                        meta: {
                            title: "我的信息"
                        }

                    },
                    {
                        path: "user-star",
                        name: "user-star",
                        component: () => import('@/views/dashboard/user-center/user-star.vue'),
                        meta: {
                            title: "我的收藏"
                        }
                    },
                    {
                        path: "user-comment",
                        name: "user-comment",
                        component: () => import('@/views/dashboard/user-center/user-comment.vue'),
                        meta: {
                            title: "我的评论"
                        }
                    },
                    {
                        path: "user-feedback",
                        name: "user-feedback",
                        component: () => import('@/views/dashboard/user-center/user-feedback.vue'),
                        meta: {
                            title: "我的反馈"
                        }
                    }
                ]
            },
            {
                path: "users",
                name: "users",
                meta: {
                    title: "用户管理",
                    requiresAdmin: true
                },
                children: [
                    {
                        path: "user-list",
                        name: "user-list",
                        component: () => import('@/views/dashboard/users/user-list.vue'),
                        meta: {
                            title: "用户列表"
                        }
                    }
                ]
            },
            {
                path: "articles",
                name: "articles",
                meta: {
                    title: "文章管理",
                    requiresAdmin: true
                },
                children: [
                    {
                        path: "article-publish",
                        name: "article-publish",
                        component: () => import('@/views/dashboard/articles/article-publish.vue'),
                        meta: {
                            title: "发布文章"
                        }
                    },
                    {
                        path: "comment-list",
                        name: "comment-list",
                        component: () => import('@/views/dashboard/articles/comment-list.vue'),
                        meta: {
                            title: "评论列表"
                        }
                    },
                    {
                        path: "article-list",
                        name: "article-list",
                        component: () => import('@/views/dashboard/articles/article-list.vue'),
                        meta: {
                            title: "文章列表"
                        }
                    }
                ]
            },
            {
                path: "images",
                name: "images",
                meta: {
                    title: "图片管理",
                    requiresAdmin: true
                },
                children: [
                    {
                        path: "image-list",
                        name: "image-list",
                        component: () => import('@/views/dashboard/images/image-list.vue'),
                        meta: {
                            title: "图片列表"
                        }
                    }
                ]
            },
            {
                path: "system",
                name: "system",
                meta: {
                    title: "系统管理",
                    requiresAdmin: true
                },
                children: [
                    {
                        path: "feedback-list",
                        name: "feedback-list",
                        component: () => import('@/views/dashboard/system/feedback-list.vue'),
                        meta: {
                            title: "反馈列表"
                        }
                    },
                    {
                        path: "login-logs",
                        name: "login-logs",
                        component: () => import('@/views/dashboard/system/login-logs.vue'),
                        meta: {
                            title: "登录日志"
                        }
                    },
                    {
                        path: "app-config",
                        name: "app-config",
                        redirect: "/dashboard/system/app-config/site-config",
                        component: () => import('@/views/dashboard/system/app-config.vue'),
                        meta: {
                            title: "应用配置"
                        },
                        children: [
                            {
                                path: "site-config",
                                name: "site-config",
                                component: () => import('@/views/dashboard/system/config/site-config.vue'),
                                meta: {
                                    title: "网站配置"
                                }
                            },
                            {
                                path: "system-config",
                                name: "system-config",
                                component: () => import('@/views/dashboard/system/config/system-config.vue'),
                                meta: {
                                    title: "系统配置"
                                }
                            },
                            {
                                path: "email-config",
                                name: "email-config",
                                component: () => import('@/views/dashboard/system/config/email-config.vue'),
                                meta: {
                                    title: "邮箱配置"
                                }
                            },
                            {
                                path: "qq-config",
                                name: "qq-config",
                                component: () => import('@/views/dashboard/system/config/qq-config.vue'),
                                meta: {
                                    title: "QQ登录配置"
                                }
                            },
                            {
                                path: "qiniu-config",
                                name: "qiniu-config",
                                component: () => import('@/views/dashboard/system/config/qiniu-config.vue'),
                                meta: {
                                    title: "七牛云配置"
                                }
                            },
                            {
                                path: "jwt-config",
                                name: "jwt-config",
                                component: () => import('@/views/dashboard/system/config/jwt-config.vue'),
                                meta: {
                                    title: "jwt配置"
                                }
                            },
                            {
                                path: "gaode-config",
                                name: "gaode-config",
                                component: () => import('@/views/dashboard/system/config/gaode-config.vue'),
                                meta: {
                                    title: "高德配置"
                                }
                            }
                        ]
                    }
                ]
            }
        ]
    },
    {
        path: "/404",
        name: "404",
        component: () => import('@/views/error/index.vue')
    },
    {
        path: "/:catchAll(.*)",
        component: () => import('@/views/error/index.vue')
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: routes,
})

export default router

router.beforeEach((to, from, next) => {
    const userStore = useUserStore()
    const layoutStore = useLayoutStore()
    
    // 先处理不需要认证的路由
    if (!to.matched.some(record => record.meta.requiresAuth)) {
        next(); // 不需要登录，直接继续访问
        return;
    }
    
    // 处理需要认证的路由
    userStore.initializeUserInfo().then(() => {
        const isAuthenticated = userStore.isLoggedIn // 检查用户是否登录的逻辑
        const isAdmin = userStore.isAdmin // 检查用户是否为管理员的逻辑
        
        if (!isAuthenticated) {
            ElMessageBox.confirm(
                '登录已过期，需要重新登录，是否跳转到登录页？', 'Warning', {
                    cancelButtonText: '取消',
                    confirmButtonText: '确定',
                    type: 'warning',
                })
                .then(() => {
                    router.push({name: 'index', replace: true}).then(() => {
                        layoutStore.state.popoverVisible = true;
                        layoutStore.state.loginVisible = true;
                    });
                })
                .catch(() => {
                    router.push({name: from.name as string || 'index'}).then();
                });
        } else if (to.matched.some(record => record.meta.requiresAdmin) && !isAdmin) {
            ElMessageBox.confirm(
                '权限不足，请确认您的用户角色是否具备访问该页面的权限。', 'Warning', {
                    confirmButtonText: '确定',
                    type: 'warning',
                })
                .then(() => {
                    router.push({name: from.name as string || 'index'}).then();
                });
        } else {
            next(); // 继续访问
        }
    }).catch(error => {
        console.error('Error in router guard:', error);
        next(); // 即使出错，也继续访问，避免页面阻塞
    });
});