<template>
  <el-card class="MovieAdvertisement">
    <el-row class="title">Classic Movie</el-row>
    <el-carousel :interval="3000" height="600px">
      <el-carousel-item v-for="movie in movieList" :key="movie.id">
        <div class="movie-item">
          <!-- 改用原生 img -->
          <img :src="movie.image" class="movie-img" />
          <div class="movie-info">
            <div class="movie-title">{{ movie.title }}</div>
            <div class="movie-description">{{ movie.description }}</div>
            <el-link :href="movie.link" target="_blank" type="primary" class="movie-link">点我跳转</el-link>
          </div>
        </div>
      </el-carousel-item>
    </el-carousel>
  </el-card>
</template>

<script setup lang="ts">
const movieList = [
  {id:1, title:"猫鼠游戏", link:"https://movie.douban.com/subject/1305487/", description:"弗兰克（莱昂纳多·迪卡普里奥 Leonardo DiCaprio饰）是FBI有史以来年龄最小的通缉犯。他的犯罪手段神通广大，伪装身份的能力超乎常人，全美各地几乎都留下他的犯罪足迹。乔装医生、律师、飞行员，统统都是他曾经的身份。当然，弗兰克凭着这个“本领”，骗取了高达几百万的金额。", image:"/image/movie_3.jpg"},
  {id:2, title:"机器人总动员", link:"https://movie.douban.com/subject/2131459/", description:"公元2805年，人类文明高度发展，却因污染和生活垃圾大量增加使得地球不再适于人类居住。地球人被迫乘坐飞船离开故乡，进行一次漫长无边的宇宙之旅。临行前他们委托Buynlarge的公司对地球垃圾进行清理，该公司开发了名为WALL·E（Waste Allocation Load Lifters – Earth 地球废品分装员）的机器人担当此重任......", image:"/image/movie_1.jpg"},
  {id:3, title:"让子弹飞", link:"https://movie.douban.com/subject/3742360/", description:"上任鹅城！", image:"/image/movie_4.jpg"},
  {id:4, title:"星际穿越", link:"https://movie.douban.com/subject/1889243/", description:"他们穿越遥远的星系银河，感受了一小时七年光阴的沧海桑田，窥见了未知星球和黑洞的壮伟与神秘。在浩瀚宇宙的绝望而孤独角落，总有一份超越了时空的笃定情怀将他们紧紧相连……", image:"/image/movie_2.jpg"}
];
</script>

<style scoped lang="scss">
.MovieAdvertisement {
  margin-bottom: 20px;

  .title {
    font-size: 24px;
    font-weight: bold;
    margin-bottom: 20px;
    display: flex;
    justify-content: center;
  }

  // Vue 3 scoped 样式需要使用 :deep() 才能渗透到 Element Plus 内部组件
  :deep(.el-carousel__item) {
    background-color: white;
  }

  .movie-item {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;

    .movie-img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      display: block; // 消除图片底部间隙
    }

    .movie-info {
      position: absolute;
      bottom: 0;
      left: 0;
      right: 0;
      padding: 15px;
      background: linear-gradient(to top, rgba(0,0,0,0.8), transparent);
      color: white;
      opacity: 0;
      transition: opacity 0.3s ease;

      .movie-title {
        font-size: 18px;
        margin-bottom: 8px;
        font-weight: bold;
      }

      .movie-description {
        font-size: 14px;
        margin-bottom: 10px;
        line-height: 1.4;
      }

      .movie-link {
        color: #409eff;
        font-size: 14px;
      }
    }

    &:hover .movie-info {
      opacity: 1;
    }
  }
}
</style>