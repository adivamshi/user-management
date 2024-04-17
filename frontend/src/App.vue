<template>
  <el-container>
    <el-header style="background-color: #409EFF; color: white; text-align: center; padding: 20px 50px;">
      User Management Portal
    </el-header>
    <el-container>
      <el-aside width="200px" style="background-color: #333; color: white;">
        <!-- Navigation Menu -->
        <el-menu
            default-active="1"
            class="el-menu-vertical-demo"
            style="border-right: none;"
            background-color="#545c64"
            text-color="#fff"
            active-text-color="#ffd04b">
          <el-menu-item index="1" @click="goTo('/')">
            <span slot="title">Home</span>
          </el-menu-item>
          <el-menu-item index="2" @click="goTo('/add-user')">
            <span slot="title">Add User</span>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script lang="ts">
import { defineComponent } from 'vue';
import { useRoute, useRouter } from 'vue-router';

export default defineComponent({
  name: 'App',
  setup() {
    const router = useRouter();
    const route = useRoute();

    function goTo(path: string) {
      if (route.path !== path) {
        router.push(path);
      }
    }

    return {
      goTo
    };
  }
});
</script>

<style scoped>
.el-menu-vertical-demo:not(.el-menu--collapse) {
  width: 200px;
  min-height: 100vh;
}
</style>
