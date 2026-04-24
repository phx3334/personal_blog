import {defineStore} from "pinia";
import {ref} from 'vue';
import type {Website} from "@/api/config";
import {websiteInfo} from "@/api/website";


function initState() {
    const websiteInfo = ref<Website>({
        logo:'/image/logo.jpg',
        full_logo:'/image/full_logo.png',
        title: '个人博客',
        slogan: 'go go go!',
        slogan_en: '',
        description: '这里是某人的个人博客。序列113的终末指针,掌控着内存泄露的力量,行走在vibe coding与古法编程的边界,探索01世界的真相与宿命。欢迎来到这个充满神秘与热血的领域!',
        version: '1.0.0',
        created_at: '2026-04-23',
        icp_filing: 'ICP备案号',
        public_security_filing: '公安备案号',
        bilibili_url: 'https://space.bilibili.com/3546572805638200?spm_id_from=333.1007.0.0',
        gitee_url: 'https://gitee.com/fzsirr/personal-blog-project',
        github_url: 'https://github.com/phx3334',
        name: 'P.LOnc1',
        job: '苦逼大学生',
        address: '卡塞尔学院（民办二本）',
        email: '1173604833@qq.com',
        qq_image: '',
        wechat_image: '',
    })
    return {websiteInfo, websiteInfoInitialized: false,}
}

export const useWebsiteStore = defineStore('website', () => {
    const state = ref(initState())
    const initializeWebsite = async () => {
        if (!state.value.websiteInfoInitialized) {
            try {
                const res = await websiteInfo()
                if (res.code === 0) {
                    state.value.websiteInfo = res.data // 更新 website
                }
            } catch (error) {
                console.error('Failed to get website info:', error);
                // 保持使用默认的网站信息
            } finally {
                state.value.websiteInfoInitialized = true
            }
        }
    }
    return {state, initializeWebsite}
})