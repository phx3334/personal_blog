<template>
  <div class="user-info">
    <el-col :span="12">
      <div class="info">
        <div class="title">
          <el-row>用户信息</el-row>
        </div>
        <div class="content">
          <el-form
              ref="userChangeInfoForm"
              :model="userChangeInfoFormData"
              :rules="rules"
              :validate-on-rule-change="false"
              hide-required-asterisk
              label-width="auto"
              style="max-width: 400px"
          >
            <el-form-item label="头像">
              <el-upload
                  class="avatar-uploader"
                  action="#"
                  :show-file-list="false"
                  :on-change="handleAvatarChange"
                  :before-upload="beforeAvatarUpload"
              >
                <div class="avatar-container">
                  <el-image :src="userInfo.avatar || '/image/avatar.jpg'" alt="" class="avatar"/>
                </div>
              </el-upload>
            </el-form-item>
            <el-form-item label="uuid">
              {{ userInfo.uuid }}
            </el-form-item>
            <el-form-item label="用户名" prop="username">
              <el-input @change="updateUserInfo" v-model="userChangeInfoFormData.username"/>
            </el-form-item>
            <el-form-item label="地址" prop="address">
              <el-input @change="updateUserInfo" v-model="userChangeInfoFormData.address"/>
            </el-form-item>
            <el-form-item label="签名" prop="signature">
              <el-input @change="updateUserInfo" v-model="userChangeInfoFormData.signature" type="textarea" :rows="2"/>
            </el-form-item>
            <el-form-item label="邮箱">
              {{ userInfo.email }}
            </el-form-item>
            <el-form-item label="用户权限">
              {{ userInfo.role_id === 1 ? "普通用户" : "管理员" }}
            </el-form-item>
            <el-form-item label="注册来源">
              {{ userInfo.register }}
            </el-form-item>
          </el-form>
        </div>
      </div>
      <div class="operation" v-if="userStore.state.userInfo.register==='邮箱'">
        <div class="title">
          <el-row>操作</el-row>
        </div>
        <div class="content">
          <el-button @click="layoutStore.state.passwordResetVisible = true">修改密码</el-button>
        </div>
        <el-dialog
            v-model="passwordResetVisible"
            width="500"
            align-center
            destroy-on-close
            :before-close="passwordResetVisibleSynchronization"
        >
          <template #header>
            修改密码
          </template>
          <password-reset-form/>
          <template #footer>
          </template>
        </el-dialog>
      </div>
    </el-col>
    <el-col :span="12">
      <div class="card">
        <div class="title">
          <el-row>用户卡片</el-row>
        </div>
      </div>
      <div class="content">
        <user-card :key="cardKey" :uuid="userInfo.uuid" :user-card-info="null"/>
      </div>
    </el-col>
  </div>
</template>

<script setup lang="ts">
import {useUserStore} from "@/stores/user";
import {reactive, ref, watch} from "vue";
import {userChangeInfo, type UserChangeInfoRequest} from "@/api/user";
import {imageUpload, type ImageUploadResponse} from "@/api/image";
import type {FormInstance, FormRules} from "element-plus";
import UserCard from "@/components/widgets/UserCard.vue";
import {useLayoutStore} from "@/stores/layout";
import PasswordResetForm from "@/components/forms/PasswordResetForm.vue";
import {ElMessage} from "element-plus";
import { ca } from "element-plus/es/locales.mjs";


const userStore = useUserStore()
const layoutStore = useLayoutStore()

const userInfo = ref(userStore.state.userInfo)

const userChangeInfoForm = ref<FormInstance>()

const userChangeInfoFormData = reactive<UserChangeInfoRequest>({
  username: userInfo.value.username,
  address: userInfo.value.address,
  signature: userInfo.value.signature,
  avatar: userInfo.value.avatar,
})

const rules = reactive<FormRules<UserChangeInfoRequest>>({
  username: [{
    required:true,
    max:20,
    trigger:'blur',
    message:'用户名长度不应大于20位'
  }],
  address: [{
    max: 200,
    trigger: 'blur',
    message:'地址长度不应大于200位'
  }],
  signature: [{
    max: 320,
    trigger: 'blur',
    message:'签名长度不应大于320位'
  }]
})

const cardKey = ref(0)

const updateUserInfo = async () => {
  const isValid: boolean = await new Promise((resolve) => {
    userChangeInfoForm.value?.validate((valid: boolean) => {
      resolve(valid)
    })
  })

  if (isValid) {
    const res = await userChangeInfo(userChangeInfoFormData)
    if (res.code === 0) {
      cardKey.value += 1
    }
  }
}

const passwordResetVisible = ref(layoutStore.state.passwordResetVisible)
watch(
    () => layoutStore.state.passwordResetVisible,
    (newValue) => {
      passwordResetVisible.value = newValue
    }
)

const passwordResetVisibleSynchronization = () => {
  layoutStore.state.passwordResetVisible = false
}

const handleAvatarChange = async (file: any) => {
  try {
    const formData = new FormData();
    formData.append('image', file.raw);
    
    const res = await imageUpload(formData);
    if (res.code === 0) {
      const avatarUrl = res.data.url;
      
      // 更新表单数据中的头像字段
      userChangeInfoFormData.avatar = avatarUrl;
      
      // 调用 updateUserInfo 函数更新用户信息
      const isValid: boolean = await new Promise((resolve) => {
        userChangeInfoForm.value?.validate((valid: boolean) => {
          resolve(valid);
        });
      });
      
      if (isValid) {
        const updateRes = await userChangeInfo(userChangeInfoFormData);
        if (updateRes.code === 0) {
          // 重新获取最新用户信息，确保数据一致性
          try {
            // 假设 userChangeInfo 返回更新后的用户信息
           userInfo.value.avatar = avatarUrl;
           userStore.state.userInfo.avatar = avatarUrl
          } catch (error) {
            console.error('Failed to refresh user info:', error);
          }
          cardKey.value += 1;
          ElMessage.success('头像更新成功');
        } else {
          ElMessage.error('更新用户信息失败');
        }
      }
    } 
  } catch (error) {
    // 只在控制台打印错误，不显示错误信息
    console.error('Avatar upload failed:', error);
  }
};

const beforeAvatarUpload = (file: any) => {
  const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
  const isLt2M = file.size / 1024 / 1024 < 2;
  
  if (!isJpgOrPng) {
    ElMessage.error('只能上传 JPG/PNG 图片');
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB');
  }
  
  return isJpgOrPng && isLt2M;
};
</script>

<style scoped lang="scss">
.user-info {
  display: flex;

  .info {
    .title {
      border-left: 5px solid blue;
      padding-left: 10px;
    }

    .content {
      margin: 20px;

      .el-form {
        .el-form-item {
          .avatar-uploader {
            .avatar-container {
              position: relative;
              display: inline-block;
              cursor: pointer;
              transition: all 0.3s;
            }
            
            .avatar-container:hover {
              transform: scale(1.05);
            }
            
            .avatar {
              height: 100px;
              width: 100px;
              border-radius: 50%;
              transition: all 0.3s;
            }
          }
        }
      }
    }
  }

  .operation {
    .title {
      border-left: 5px solid blue;
      padding-left: 10px;
    }

    .content {
      margin: 20px;
    }
  }

  .card {
    .title {
      border-left: 5px solid blue;
      padding-left: 10px;
    }

    .content {
      margin: 20px;
    }
  }
}
</style>