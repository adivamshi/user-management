<template>
  <el-container style=" margin: 20px;">
    <el-form :model="user" :rules="rules" ref="userForm" label-width="120px">
      <el-form-item label="Name" prop="name">
        <el-input v-model="user.name"></el-input>
      </el-form-item>
      <el-form-item label="Email" prop="email">
        <el-input v-model="user.email"></el-input>
      </el-form-item>
      <el-form-item label="Age" prop="age">
        <el-input v-model.number="user.age" type="number"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm">Submit</el-button>
        <el-button @click="resetForm">Reset</el-button>
      </el-form-item>
    </el-form>
  </el-container>
</template>

<script lang="ts">
import { defineComponent, reactive, ref } from 'vue';
import { ElForm } from 'element-plus';

export default defineComponent({
  name: 'AddUser',
  setup() {
    const userForm = ref<InstanceType<typeof ElForm>>();
    const user = reactive({
      name: '',
      email: '',
      age: 0
    });

    const rules = {
      name: [
        { required: true, message: 'Please input the user name', trigger: 'blur' }
      ],
      email: [
        { required: true, message: 'Please input the email address', trigger: 'blur' },
        { type: 'email', message: 'Please input correct email address', trigger: ['blur', 'change'] }
      ],
      age: [
        { required: true, message: 'Please input the age', trigger: 'blur' },
        { type: 'number', min: 18, message: 'User must be at least 18 years old', trigger: 'blur' }
      ]
    };

    const submitForm = () => {
      userForm.value?.validate((valid) => {
        if (valid) {
          // Todo: Invoke Network call to add user
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    };

    const resetForm = () => {
      userForm.value?.resetFields();
    };

    return {
      user,
      rules,
      userForm,
      submitForm,
      resetForm
    };
  }
});
</script>
