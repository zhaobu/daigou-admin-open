<template>
  <el-form
    ref="formRef"
    :model="operaData"
    :rules="argsRules"
    label-width="80px"
    size="mini"
    label-position="left"
  >
    <el-divider content-position="center">参数信息</el-divider>
    <el-form-item label="测试参数" prop="shop_code">
      <el-input autocomplete="off"></el-input>
    </el-form-item>
  </el-form>
</template>

<script>
export default {
  name: "CloseShop",
  props: ["operaData"],
  data() {},
  methods: {
    checkArgs() {},
  },
};
</script>
<style lang="scss">
</style>