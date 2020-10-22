<template>
  <div>
    <el-form :model="formData" size="medium" label-position="left" label-width="90px">
      <el-form-item label="ip地址">
        <el-input v-model="formData.ip" readonly clearable :style="{width: '250px'}" />
      </el-form-item>
      <el-form-item label="端口">
        <el-input v-model="formData.port" readonly clearable :style="{width: '250px'}" />
      </el-form-item>
      <el-form-item label="用户名">
        <el-input v-model="formData.user_name" readonly clearable :style="{width: '250px'}" />
      </el-form-item>
      <el-form-item label="密码">
        <el-input
          v-model="formData.password"
          readonly
          clearable
          show-password:true
          :style="{width: '250px'}"
        />
      </el-form-item>
      <el-form-item label="数据库名称">
        <el-input v-model="formData.db_name" readonly clearable :style="{width: '250px'}" />
      </el-form-item>
      <el-form-item label="连接参数">
        <el-input
          type="textarea"
          v-model="formData.params"
          readonly
          clearable
          :style="{width: '250px'}"
        />
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  name: "Mysql",
  props: ["formData"],
};
</script>
<style>
</style>
