<template>
    <div class="tolink">
        <el-image class="top-image" src="/image/珂朵莉.jpg" fit="contain" />
        <div class="good_link">
            <el-card
                v-for="item in linkList"
                :key="item.id"
                class="link-card"
            >
            <div class="card-content">
                <span>{{ item.text }}</span>
                <br>
                <el-link :href="item.url" target="_blank" type="primary">
                    {{ item.LinkName }}
                </el-link>
            </div>
            </el-card>
        </div>
    </div>
</template>

<script setup lang="ts">
import { reactive } from 'vue';

    const linkList=reactive([
        {id:1,text: "mbti和sbti都测了，各位绅士们该试一下xpti了",url: "https://xpti.pages.dev/",LinkName:"xpti测试"},
        {id:2,text:"接下来是一款由无代码应用搭建平台制作的一个校园galgame,虽然已经断更很久了，就连作者自己都找不到原项目了，但链接还在",url:"https://app-7jrqkhu1vdoh.appmiaoda.com",LinkName:"校园galgame"},
        {id:3,text:"下面是我一个在米哈游实习的群友做的明日方舟同人视觉小说类型游戏，感兴趣的可以试玩一下",url:"https://loveark.penumbras.work/",LinkName:"恋与方舟"}
    ])

</script>

<style scoped>
.tolink{
    display: flex;
    flex-direction: column;
    gap: 20px;
    padding: 20px;
    background-color: #f5f5f5;
    border-radius: 8px;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);

}


</style>