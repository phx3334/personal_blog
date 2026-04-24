<template>
  <div class="index">
    <web-navbar/>
    <carousel/>
    <el-container class="main-content">
      <div class="container">
        <el-aside>
          <tolink/>
        </el-aside>
        <el-main>
          <movie-advertisement/>
          <stack-skills/>
        </el-main>
        <el-aside>
          <calendar/>
          <profile-card/>
          <tag-cloud/>
          <recent-comments/>
          <feedback/>
        </el-aside>
      </div>
    </el-container>
  </div>
</template>

<script setup lang="ts">
import Carousel from "@/components/layout/Carousel.vue";
import WebNavbar from "@/components/layout/WebNavbar.vue";
import MovieAdvertisement from "@/components/pages/MovieAdvertisement.vue";
import Calendar from "@/components/pages/Calendar.vue";
import ProfileCard from "@/components/pages/ProfileCard.vue";
import TagCloud from "@/components/pages/TagCloud.vue";
import RecentComments from "@/components/pages/RecentComments.vue";
import Feedback from "@/components/pages/Feedback.vue";
import Tolink from "@/components/pages/Tolink.vue";
import StackSkills from "@/components/pages/StackSkills.vue";
</script>

<style scoped lang="scss">
.index {
  .main-content {
    display: flex;
    justify-content: center;

    .container {
      display: flex;
      max-width: 1400px;
      width: 100%;

      .el-aside:first-child {
        width: 20%;
        padding: 20px 10px 20px 0;
      }

      .el-main {
        width: 55%;
      }

      .el-aside:last-child {
        width: 25%;
        padding: 20px;
      }
    }
  }
}
</style>
