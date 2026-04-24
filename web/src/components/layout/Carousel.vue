<template>
  <div class="carousel">
    <el-carousel interval="3000" trigger="click" height="700px">
      <el-carousel-item v-for="item in imgList" :key="item">
        <el-image fit="cover" :src=item alt=""></el-image>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script setup lang="ts">
import {ref} from "vue";
import {websiteCarousel} from "@/api/website";

const imgList = ref<string[]>([
  '/image/carousel_1.jpg',
  '/image/carousel_2.jpg',
  '/image/carousel_3.jpg',
  '/image/carousel_4.jpg',
])

const getWebsiteCarousel = async () => {
  try {
    const res = await websiteCarousel()
    if (res.code === 0 && res.data.length !== 0) {
      imgList.value = res.data
    }
  } catch (error) {
    console.error('Failed to get carousel images:', error)
    // 保持使用默认的图片路径
  }
}

getWebsiteCarousel()
</script>

<style scoped lang="scss">
.carousel {
  width: 100%;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;

  .el-carousel {
    width: 60%; /* 设置宽度为页面的60% */
    height: 700px; /* 保持高度不变 */
    
    .carousel-item {
      display: flex;
      justify-content: center;
      align-items: center;
      height: 100%;
      background-color: #f5f5f5; /* 添加背景色，使图片周围有边框感 */
    }
    
    .el-image {
      max-height: 100%;
      max-width: 100%;
      object-fit: contain; /* 确保图片完整显示，上下内容尽可能显示 */
    }
  }
}

/* 响应式调整 */
@media (max-width: 768px) {
  .carousel {
    .el-carousel {
      width: 80%;
      height: 600px;
    }
  }
}
</style>