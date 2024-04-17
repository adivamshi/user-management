<template>
  <el-container style="margin: 20px;">
    <el-form :model="user" :rules="rules" ref="userForm" label-width="120px">
      <el-form-item label="Name" prop="name">
        <el-input v-model="user.name" :disabled="!isEditing"></el-input>
      </el-form-item>
      <el-form-item label="Email" prop="email">
        <el-input v-model="user.email" :disabled="!isEditing"></el-input>
      </el-form-item>
      <el-form-item label="Age" prop="age">
        <el-input v-model.number="user.age" type="number" :disabled="!isEditing"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button @click="toggleEdit" type="primary">{{ isEditing ? 'Save' : 'Edit' }}</el-button>
        <el-button @click="confirmDelete" type="danger">Delete</el-button>
      </el-form-item>
    </el-form>
  </el-container>
</template>

<script lang="ts">
import { defineComponent, reactive, ref, onMounted } from 'vue';
import { ElForm } from 'element-plus';
import {DeleteUser, GetUserById, UpdateUser} from "@/utils/network";
import {useRouter} from "vue-router";

export default defineComponent({
  name: 'UserDetail',
  props: {
    id: Number
  },
  setup(props) {
    const userForm = ref<InstanceType<typeof ElForm>>();
    const isEditing = ref(false);
    const router = useRouter()
    const userid = ref(Number(props.id))
    const user = reactive({
      name: '',
      email: '',
      age: 0
    });
    const rules = {
      name: [{ required: true, message: 'Name is required', trigger: 'blur' }],
      email: [{ required: true, message: 'Email is required', trigger: 'blur' }],
      age: [{ required: true, message: 'Age is required', trigger: 'blur' }]
    };

    onMounted(() => {
      userid.value = Number(props.id);
      fetchUserDetails();
    });

    const fetchUserDetails = () => {
      if(userid.value !== 0) {
        GetUserById(userid.value).then((response) => {
          user.name = response.name;
          user.email = response.email;
          user.age = response.age;
        })
      }
    };

    const toggleEdit = () => {
      if (isEditing.value) {
        userForm.value?.validate((valid) => {
          if (valid) {
            UpdateUser(userid.value, user.name,user.email, user.age)
                .then((response) => {
                  alert("Updated successfully");
                  isEditing.value = false;
                  user.age = response.updated.age;
                  user.email = response.updated.email;
                  user.name = response.updated.name;
                }).catch(err =>{
                  console.log(err);
            });
          }
        });
      } else {
        isEditing.value = true;
      }
    };
    const confirmDelete = () => {
      if (confirm('Are you sure you want to delete this user?') && userid.value !== 0) {
        DeleteUser(userid.value)
        .then((response) => {
          console.log("deleted successfully", response);
          alert('Deleted');
          router.push({name: 'Home' });
        }).catch((err)=>{
          console.log(err);
        });
      }
    };

    return {
      user,
      isEditing,
      rules,
      toggleEdit,
      confirmDelete,
      userForm
    };
  }
});
</script>
