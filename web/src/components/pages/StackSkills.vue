<template>
  <section class="skill-section">
    <!-- 顶部图片：绝对居中 -->
    <div class="skill-img-box">
      <img src="/image/Skills.jpg" class="skill-img" />
    </div>

    <!-- 技术块区域：强制2列布局 + 消除底部空白 -->
    <div class="skill-grid">
      <div class="stack-block">
        <div class="tech-label">cpp</div>
        <div class="progress-bar">
          <div class="progress-fill" style="width: 50%"></div>
        </div>
        <div class="overlay-text">熟悉cpp with stl和cpp的部分进阶用法,曾使用cpp写过ue5的GAS架构.</div>
      </div>
      <div class="stack-block">
        <div class="tech-label">go</div>
        <div class="progress-bar">
          <div class="progress-fill" style="width: 50%"></div>
        </div>
        <div class="overlay-text">较熟悉go语言和其网络编程,使用go写过博客项目的后端部分。</div>
      </div>
      <div class="stack-block">
        <div class="tech-label">vue3</div>
        <div class="progress-bar">
          <div class="progress-fill" style="width: 60%"></div>
        </div>
        <div class="overlay-text">熟悉vue3的语法和组件化开发,使用vue3写过博客项目的前端部分。</div>
      </div>
      <div class="stack-block">
        <div class="tech-label">ue5</div>
        <div class="progress-bar">
          <div class="progress-fill" style="width: 10%"></div>
        </div>
        <div class="overlay-text">了解ue5的界面的各种功能,比如场景编辑器,材质编辑器,动画编辑器.</div>
      </div>
      <div class="stack-block">
        <div class="tech-label">python</div>
        <div class="progress-bar">
          <div class="progress-fill" style="width: 10%"></div>
        </div>
        <div class="overlay-text">了解python的基本语法,用python写过一些简单的功能，比如随机生成密钥。</div>
      </div>
      <div class="stack-block">
        <div class="tech-label">javascript/typescript</div>
        <div class="progress-bar">
          <div class="progress-fill" style="width: 60%"></div>
        </div>
        <div class="overlay-text">较熟悉javascript/typescript的语法和typescript的类型系统,使用typescript写过博客项目的前端部分。</div>
      </div>
    </div>
     <div class="youlinaixu-img-box">
      <img src="/image/友利奈绪.jpg" class="youlinaixu-img" />
    </div>
</section>
</template>

<script setup lang="ts">
// StackSkills组件 - 技术栈展示组件
</script>

<style scoped lang="scss">
// 全局盒模型：防止所有元素因padding/margin撑出多余空间
* {
  box-sizing: border-box;
}

// 整体容器样式：消除不必要的外边距，自适应高度
.skill-section {
  width: 100%;
  padding: 0 20px;
  max-width: 1200px; 
  margin: 16px auto; // 仅保留上下少量间距，左右居中
  box-sizing: border-box; // 防止padding撑宽容器
}

// 图片容器：居中 + 白色背景
.skill-img-box {
  background: white;
  border-radius: 16px;
  padding: 32px;
  max-width: 600px; 
  margin: 0 auto 32px; // 仅保留和技术块的底部间距
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
}

// 图片样式
.skill-img {
  width: 100%;
  height: auto;
  max-height: 200px;
  object-fit: contain;
}

.youlinaixu-img-box {
  background: white;
  border-radius: 16px;
  padding: 12px;
  width: 100%;
  max-width: none;
  margin: 32px 0 32px 0;
  display: flex;
  justify-content: center;
  align-items: center;
  box-sizing: border-box;
  overflow: hidden; // 关键：裁剪超出圆角的部分
}

.youlinaixu-img {
  width: 100%;
  height: 100%; // 改为100%，强制填满容器高度
  max-height: none; // 取消高度限制
  object-fit: cover; // 覆盖式填充（裁剪多余部分，无留白）
  // 如果不想裁剪，想拉伸填满，就把object-fit: cover换成object-fit: fill;
}

// 技术块网格：强制2列 + 消除底部空白
.skill-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr); // 固定2列
  gap: 20px; // 块之间的间距
  width: 100%;
  min-height: 0; // 关键：消除grid默认的多余高度
  box-sizing: border-box;
}

// 技术块样式（核心修复交互冲突）
.stack-block {
  position: relative;
  height: 150px; // 固定高度，避免块高度不一致导致空白
  cursor: pointer;
  transition: all 0.3s ease;
  padding: 12px;
  border-radius: 12px;
  background-color: #faf8f5;
  box-sizing: border-box; // 防止padding撑高/撑宽
  overflow: hidden; // 关键1：限制内部元素不溢出本块
  z-index: 1; // 关键2：给每个块基础层级，避免跨块层级冲突

  .tech-label {
    position: absolute;
    top: 12px;
    left: 12px;
    font-size: 24px;
    font-weight: 800;
    color: #000;
    z-index: 2; // 层级高于背景，低于overlay-text
  }

  .progress-bar {
    position: absolute;
    bottom: 12px;
    left: 12px;
    right: 12px;
    height: 8px;
    background-color: #e0e0e0;
    border-radius: 4px;
    overflow: hidden;
    z-index: 2; // 层级高于背景，低于overlay-text

    .progress-fill {
      height: 100%;
      background-color: #d3d3d3;
      border-radius: 4px;
      transition: width 0.3s ease;
    }
  }

  .overlay-text {
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    display: flex;
    align-items: center;
    justify-content: center;
    background: #000;
    color: white;
    font-size: 16px;
    font-weight: 500;
    padding: 15px;
    text-align: center;
    z-index: 3; // 关键3：仅在本块内最高层级，不跨块
    transform: translateX(-100%);
    opacity: 0;
    transition: transform 0.3s ease, opacity 0.3s ease;
    box-sizing: border-box;
  }

  &:hover {
    .overlay-text {
      transform: translateX(0);
      opacity: 1;
    }
    z-index: 10; // 关键4：hover时提升本块层级，压过相邻块
  }
}
</style>