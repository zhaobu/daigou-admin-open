<template>
  <div class="app-container">
    <el-form
      ref="elFormRef"
      :model="formData"
      size="medium"
      label-position="left"
      label-width="100px"
    >
      <el-form-item label="ip地址">
        <el-input v-model="formData.ip" readonly clearable :style="{width: '250px'}" />
      </el-form-item>
      <el-form-item label="端口">
        <el-input v-model="formData.port" readonly clearable :style="{width: '250px'}" />
      </el-form-item>
      <el-form-item label="数据库序号">
        <el-input v-model="formData.db" readonly clearable :style="{width: '250px'}" />
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
    </el-form>
  </div>
</template>
<script>
export default {
  name: "Redis",
  components: {},
  props: ["formData"],
};
</script>
<style>
</style>
