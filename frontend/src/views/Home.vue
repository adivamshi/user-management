<template>
  <div>
    <el-input v-model="search" placeholder="Search users..." @input="fetchUsers" />
    <div v-if="users.length > 0">
      <el-table :data="users" style="width: 100%">
        <el-table-column prop="name" label="Name"></el-table-column>
        <el-table-column prop="email" label="Email"></el-table-column>
        <el-table-column prop="age" label="Age"></el-table-column>
        <el-table-column
            label="Operations"
            width="180">
          <template v-slot="scope">
            <el-button @click="viewDetails(scope.row)" type="primary" size="small">View Details</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
          v-model:current-page="currentPage"
          v-model:page-size="pageSize"
          :page-sizes="pageSizes"
          :small="false"
          :disabled="false"
          :background="false"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
      />
    </div>
    <div v-else>
      <el-text class="mx-1" type="warning">No Users Found, adjust search criteria or add users to display</el-text>
    </div>
  </div>
</template>

<script lang="ts">
import {defineComponent, ref, onMounted} from 'vue';
import {type LocationAsRelativeRaw, useRouter} from 'vue-router'
import type {GetUserDataResponse, User} from "@/utils/types";
import {GetUserData} from "@/utils/network";

export default defineComponent({
  name: 'UserList',
  setup() {
    const users = ref<User[]>([]);
    const router = useRouter()
    const search = ref('');
    const currentPage = ref(1);
    const pageSize = ref(10);
    const pageSizes = ref([10, 20, 50, 100, 200, 300, 400]);
    const total = ref(0);

    const fetchUsers = async () => {
      try {
        const response = await GetUserData(search.value, currentPage.value, pageSize.value);
        handleGetAllUserData(response)
      } catch (error) {
        console.error('Failed to fetch users:', error);
      }
    };

    onMounted(fetchUsers);
    const handleGetAllUserData = (data:GetUserDataResponse) =>{
      users.value = data.users
      total.value = data.total
    }

    const handleSizeChange = (val: number) => {
      GetUserData("", currentPage.value, val)
          .then((data) => handleGetAllUserData(data))
          .catch(err =>{
            console.log(err)
          })

    }
    const handleCurrentChange = (val: number) => {
      GetUserData("", val, pageSize.value)
          .then((data) => handleGetAllUserData(data))
          .catch(err =>{
            console.log(err)
          })
    }

    function viewDetails(user: any) {
      let route = {} as LocationAsRelativeRaw
      route.name = "UserDetails"
      // route.path= "/users/" + user.id;
      route.params = {
        id: user.id
      }
      router.push(route);
    }

    return {
      users,
      search,
      currentPage,
      pageSize,
      pageSizes,
      total,
      fetchUsers,
      handleSizeChange,
      handleCurrentChange,
      viewDetails
    };
  },
});</script>
router.push({name: 'UserDetail', params: {userId: user.id}});

